package server

import (
	"strconv"
	"github.com/pkg/errors"
	"github.com/eug48/fhir/utils"
	"time"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/eug48/fhir/models"
	"github.com/eug48/fhir/models2"
	"github.com/eug48/fhir/search"
)

// BatchController handles FHIR batch operations via input bundles
type BatchController struct {
	DAL    DataAccessLayer
	Config Config
}

// NewBatchController creates a new BatchController based on the passed in DAL
func NewBatchController(dal DataAccessLayer, config Config) *BatchController {
	return &BatchController{
		DAL:    dal,
		Config: config,
	}
}

func abortWithErr(c *gin.Context, err error) {
	outcome := models.CreateOpOutcome("fatal", "exception", "", err.Error())
	c.AbortWithStatusJSON(http.StatusBadRequest, outcome)
}

// Post processes and incoming batch request
func (b *BatchController) Post(c *gin.Context) {
	bundleResource, err := FHIRBind(c, b.Config.ValidatorURL)
	if err != nil {
		abortWithErr(c, err)
		return
	}

	bundle, err := bundleResource.AsShallowBundle()
	if err != nil {
		abortWithErr(c, err)
		return
	}

	switch bundle.Type {
	case "transaction":
	case "batch":
		// TODO: If type is batch, ensure there are no interdependent resources
	default:
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Bundle type is neither 'batch' nor 'transaction'"))
	}

	// Loop through the entries, ensuring they have a request and that we support the method,
	// while also creating a new entries array that can be sorted by method.
	entries := make([]*models2.ShallowBundleEntryComponent, len(bundle.Entry))
	for i := range bundle.Entry {
		if bundle.Entry[i].Request == nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("Entries in a batch operation require a request"))
			return
		}

		switch bundle.Entry[i].Request.Method {
		default:
			c.AbortWithError(http.StatusNotImplemented,
				errors.New("Operation currently unsupported in batch requests: "+bundle.Entry[i].Request.Method))
			return
		case "DELETE":
			if bundle.Entry[i].Request.Url == "" {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch DELETE must have a URL"))
				return
			}
		case "POST":
			if bundle.Entry[i].Resource == nil {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch POST must have a resource body"))
				return
			}
		case "PUT":
			if bundle.Entry[i].Resource == nil {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch PUT must have a resource body"))
				return
			}
			if !strings.Contains(bundle.Entry[i].Request.Url, "/") && !strings.Contains(bundle.Entry[i].Request.Url, "?") {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch PUT url must have an id or a condition"))
				return
			}
		}
		entries[i] = &bundle.Entry[i]
	}

	sort.Sort(byRequestMethod(entries))

	// Now loop through the entries, assigning new IDs to those that are POST or Conditional PUT and fixing any
	// references to reference the new ID.
	refMap := make(map[string]string)
	newIDs := make([]string, len(entries))
	createStatus := make([]string, len(entries))
	for i, entry := range entries {
		if entry.Request.Method == "POST" {

			id := ""

			if len(entry.Request.IfNoneExist) > 0 {
				// Conditional Create
				query := search.Query{Resource: entry.Request.Url, Query: entry.Request.IfNoneExist}
				existingIds, err := b.DAL.FindIDs(query)
				if err != nil {
					c.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				if len(existingIds) == 0 {
					createStatus[i] = "201"
				} else if len(existingIds) == 1 {
					createStatus[i] = "200"
					id = existingIds[0]
				} else if len(existingIds) > 1 {
					createStatus[i] = "412" // HTTP 412 - Precondition Failed
				}
			} else {
				// Unconditional create
				createStatus[i] = "201"
			}

			if createStatus[i] == "201" {
				// Create a new ID
				id = bson.NewObjectId().Hex()
				newIDs[i] = id
			}

			if len(id) > 0 {
				// Add id to the reference map
				refMap[entry.FullUrl] = entry.Request.Url + "/" + id
				// Rewrite the FullUrl using the new ID
				entry.FullUrl = responseURL(c.Request, b.Config, entry.Request.Url, id).String()
			}

		} else if entry.Request.Method == "PUT" && isConditional(entry) {
			// We need to process conditionals referencing temp IDs in a second pass, so skip them here
			if hasTempID(entry.Request.Url) {
				continue
			}

			if err := b.resolveConditionalPut(c.Request, i, entry, newIDs, refMap); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	}

	// Second pass to take care of conditionals referencing temporary IDs.  Known limitation: if a conditional
	// references a temp ID also defined by a conditional, we error out if it hasn't been resolved yet -- too many
	// rabbit holes.
	for i, entry := range entries {
		if entry.Request.Method == "PUT" && isConditional(entry) {
			// Use a regex to swap out the temp IDs with the new IDs
			for oldID, ref := range refMap {
				re := regexp.MustCompile("([=,])(" + oldID + "|" + url.QueryEscape(oldID) + ")(&|,|$)")
				entry.Request.Url = re.ReplaceAllString(entry.Request.Url, "${1}"+ref+"${3}")
			}

			if hasTempID(entry.Request.Url) {
				c.AbortWithError(http.StatusNotImplemented,
					errors.New("Cannot resolve conditionals referencing other conditionals"))
				return
			}

			if err := b.resolveConditionalPut(c.Request, i, entry, newIDs, refMap); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	}

	// When being converted to BSON references will be updated to reflect newly assigned IDs
	bundle.SetTransformReferencesMap(refMap)

	// Handle If-Match
	for _, entry := range entries {
		switch entry.Request.Method {
		case "PUT":
			if entry.Request.IfMatch != "" {
				parts := strings.SplitN(entry.Request.Url, "/", 2)
				if len(parts) != 2 { // TODO: refactor
					c.AbortWithError(http.StatusBadRequest,
						fmt.Errorf("Couldn't identify resource and id to put from %s", entry.Request.Url))
					return
				}
				id := parts[1]

				conditionalVersionId, err := utils.ETagToVersionId(entry.Request.IfMatch)
				if err != nil {
					c.AbortWithError(http.StatusBadRequest, err)
					return
				}

				currentResource, err := b.DAL.Get(id, entry.Resource.ResourceType())
				if err == ErrNotFound {
					entry.Response = &models.BundleEntryResponseComponent{
						Status: "404",
						Outcome: models.CreateOpOutcome("error", "not-found", "", "Existing resource not found when handling If-Match"),
					}
					entry.Resource = nil
				} else if err != nil {
					err = errors.Wrapf(err, "failed to get current resource while processing If-Match for %s", entry.Request.Url)
					c.AbortWithError(http.StatusInternalServerError, err)
					return
				} else if conditionalVersionId != currentResource.VersionId() {
					entry.Response = &models.BundleEntryResponseComponent{
						Status: "409",
						Outcome: models.CreateOpOutcome("error", "conflict", "", fmt.Sprintf("Version mismatch when handling If-Match (current=%s wanted=%s)", currentResource.VersionId(), conditionalVersionId)),
					}
					entry.Resource = nil
				}
			}
		}
	}

	// If have an error for a transaction, do not proceeed
	proceed := true
	if bundle.Type == "transaction" {
		for _, entry := range entries {
			if entry.Response != nil && entry.Response.Outcome != nil {
				// FIXME: ensure it is a "failed" outcome

				proceed = false
				break
			}
		}
	}

	// Then make the changes in the database and update the entry responses
	if (proceed) {
		for i, entry := range entries {
			err = b.makeUpdates(c, i, entry, createStatus, newIDs)
			if err != nil {
				statusCode, outcome := ErrorToOpOutcome(err)
				if bundle.Type == "transaction" {
					sendReply(c, statusCode, outcome)
					return
				} else {
					entry.Resource = nil
					entry.Request = nil
					entry.Response.Status = strconv.Itoa(statusCode)
					entry.Response.Outcome = outcome
				}
			}
		}
	}

	// For failing transactions return a single operation-outcome
	if bundle.Type == "transaction" {
		for _, entry := range entries {
			if entry.Response != nil && entry.Response.Outcome != nil {
				// FIXME: ensure it is a "failed" outcome

				status, err := strconv.Atoi(entry.Response.Status)
				if err != nil {
					panic(fmt.Errorf("bad Response.Status (%s)", entry.Response.Status))
				}

				sendReply(c, status, entry.Response.Outcome)
				return
			}
		}
	}

	if proceed {
		total := uint32(len(entries))
		bundle.Total = &total
		bundle.Type = fmt.Sprintf("%s-response", bundle.Type)

		c.Set("Bundle", bundle)
		c.Set("Resource", "Bundle")
		c.Set("Action", "batch")

		sendReply(c, http.StatusOK, bundle)
	}
}

func sendReply(c *gin.Context, httpStatus int, reply interface{}) {
	if c.GetBool("SendXML") {
		converterInt := c.MustGet("FhirFormatConverter")
		converter := converterInt.(*FhirFormatConverter)
		converter.SendXML(httpStatus, reply, c)
	} else {
		c.JSON(httpStatus, reply)
	}
}

func (b *BatchController) makeUpdates(c *gin.Context, i int, entry *models2.ShallowBundleEntryComponent, createStatus []string, newIDs []string) error {
	if entry.Response != nil {
		// already handled (e.g. conditional update returned 409)
		return nil
	}

	switch entry.Request.Method {
	case "DELETE":
		if !isConditional(entry) {
			// It's a normal DELETE
			parts := strings.SplitN(entry.Request.Url, "/", 2)
			if len(parts) != 2 {
				return fmt.Errorf("Couldn't identify resource and id to delete from %s", entry.Request.Url)
			}
			if _, err := b.DAL.Delete(parts[1], parts[0]); err != nil && err != ErrNotFound {
				return errors.Wrapf(err, "failed to delete %s", entry.Request.Url)
			}
		} else {
			// It's a conditional (query-based) delete
			parts := strings.SplitN(entry.Request.Url, "?", 2)
			query := search.Query{Resource: parts[0], Query: parts[1]}
			if _, err := b.DAL.ConditionalDelete(query); err != nil {
				return errors.Wrapf(err, "failed to conditional-delete %s", entry.Request.Url)
			}
		}

		entry.Request = nil
		entry.Response = &models.BundleEntryResponseComponent{
			Status: "204",
		}
	case "POST":

		entry.Response = &models.BundleEntryResponseComponent{
			Status:   createStatus[i],
			Location: entry.FullUrl,
		}

		if createStatus[i] == "201" {
			// creating
			err := b.DAL.PostWithID(newIDs[i], entry.Resource)
			if err != nil {
				return errors.Wrapf(err, "failed to create %s", entry.Request.Url)
			}
			updateEntryMeta(entry)
		} else if createStatus[i] == "200" {
			// have one existing resource
			components := strings.Split(entry.FullUrl, "/")
			existingId := components[len(components)-1]

			existingResource, err := b.DAL.Get(existingId, entry.Request.Url)
			if err != nil {
				return errors.Wrapf(err, "failed to get existing resource during conditional create of %s", entry.Request.Url)
			}
			entry.Resource = existingResource
			updateEntryMeta(entry)
		} else if createStatus[i] == "412" {
			entry.Response.Outcome = models.CreateOpOutcome("error", "duplicate", "", "search criteria were not selective enough")
			entry.Resource = nil
		}
		entry.Request = nil

	case "PUT":
		// Because we pre-process conditional PUTs, we know this is always a normal PUT operation
		entry.FullUrl = responseURL(c.Request, b.Config, entry.Request.Url).String()
		parts := strings.SplitN(entry.Request.Url, "/", 2)
		if len(parts) != 2 {
			return fmt.Errorf("Couldn't identify resource and id to put from %s", entry.Request.Url)
		}

		// Write
		createdNew, err := b.DAL.Put(parts[1], "", entry.Resource)
		if err != nil {
			return errors.Wrapf(err, "failed to update %s", entry.Request.Url)
		}

		// Response
		entry.Request = nil
		entry.Response = new(models.BundleEntryResponseComponent)
		entry.Response.Location = entry.FullUrl
		if createdNew {
			entry.Response.Status = "201"
		} else {
			entry.Response.Status = "200"
		}
		updateEntryMeta(entry)
	}
	return nil
}

func updateEntryMeta(entry *models2.ShallowBundleEntryComponent) {
	
	// TODO: keep LastModified as a string
	lastUpdated := entry.Resource.LastUpdated()
	if lastUpdated != "" {
		t := time.Time{}
		err := t.UnmarshalJSON([]byte("\"" + lastUpdated + "\""))
		if err != nil {
			panic(fmt.Errorf("failed to parse LastUpdated String: %s", lastUpdated))
		}
		entry.Response.LastModified = &models.FHIRDateTime{ Time: t, Precision: "timestamp" }
	}

	versionId := entry.Resource.VersionId()
	if versionId != "" {
		entry.Response.Etag = "W/\"" + versionId + "\""
	}
}

func (b *BatchController) resolveConditionalPut(request *http.Request, entryIndex int, entry *models2.ShallowBundleEntryComponent, newIDs []string, refMap map[string]string) error {
	// Do a preflight to either get the existing ID, get a new ID, or detect multiple matches (not allowed)
	parts := strings.SplitN(entry.Request.Url, "?", 2)
	query := search.Query{Resource: parts[0], Query: parts[1]}

	var id string
	if IDs, err := b.DAL.FindIDs(query); err == nil {
		switch len(IDs) {
		case 0:
			id = bson.NewObjectId().Hex()
		case 1:
			id = IDs[0]
		default:
			return ErrMultipleMatches
		}
	} else {
		return err
	}

	// Rewrite the PUT as a normal (non-conditional) PUT
	entry.Request.Url = query.Resource + "/" + id

	// Add the new ID to the reference map
	newIDs[entryIndex] = id
	refMap[entry.FullUrl] = entry.Request.Url

	// Rewrite the FullUrl using the new ID
	entry.FullUrl = responseURL(request, b.Config, query.Resource, id).String()

	return nil
}

func isConditional(entry *models2.ShallowBundleEntryComponent) bool {
	if entry.Request == nil {
		return false
	} else if entry.Request.Method != "PUT" && entry.Request.Method != "DELETE" {
		return false
	}
	return !strings.Contains(entry.Request.Url, "/") || strings.Contains(entry.Request.Url, "?")
}

func hasTempID(str string) bool {

	// do not match URLs like Patient?identifier=urn:oid:0.1.2.3.4.5.6.7|urn:uuid:6002c2ab-9571-4db7-9a79-87163475b071
	tempIdRegexp := regexp.MustCompile("([=,])(urn:uuid:|urn%3Auuid%3A)[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}(&|,|$)")
	matches := tempIdRegexp.MatchString(str)

	// hasPrefix := strings.HasPrefix(str, "urn:uuid:") || strings.HasPrefix(str, "urn%3Auuid%3A")
	// contains := strings.Contains(str, "urn:uuid:") || strings.Contains(str, "urn%3Auuid%3A")
	// if matches != contains {
	// fmt.Printf("re != contains (re = %t): %s\n", matches, str)
	// }

	return matches
}

// Support sorting by request method, as defined in the spec
type byRequestMethod []*models2.ShallowBundleEntryComponent

func (e byRequestMethod) Len() int {
	return len(e)
}
func (e byRequestMethod) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e byRequestMethod) Less(i, j int) bool {
	methodMap := map[string]int{"DELETE": 0, "POST": 1, "PUT": 2, "GET": 3}
	return methodMap[e[i].Request.Method] < methodMap[e[j].Request.Method]
}
