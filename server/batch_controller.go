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

// Handles batch and transaction requests
func (b *BatchController) Post(c *gin.Context) {
	bundleResource, err := FHIRBind(c, b.Config.ValidatorURL)
	if err != nil {
		abortWithErr(c, err)
		return
	}

	session := b.DAL.StartSession(c.GetHeader("Db"))
	defer session.Finish()

	bundle, err := bundleResource.AsShallowBundle(b.Config.FailedRequestsDir)
	if err != nil {
		abortWithErr(c, err)
		return
	}

	switch bundle.Type {
	case "transaction":
		err := session.StartTransaction()
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.Wrap(err, "error starting MongoDB transaction"))
			return
		}
	case "batch":
		// TODO: If type is batch, ensure there are no interdependent resources
	default:
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Bundle type is neither 'batch' nor 'transaction'"))
		return
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
		case "GET":
			if bundle.Entry[i].Request.Url == "" {
				c.AbortWithError(http.StatusBadRequest, errors.New("Batch GET must have a URL"))
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
				existingIds, err := session.FindIDs(query)
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
				entry.FullUrl = b.Config.responseURL(c.Request, entry.Request.Url, id).String()
			}

		} else if entry.Request.Method == "PUT" && isConditional(entry) {
			// We need to process conditionals referencing temp IDs in a second pass, so skip them here
			if hasTempID(entry.Request.Url) {
				// fmt.Printf("hasTempID! %s\n", entry.Request.Url)
				continue
			}

			if err := b.resolveConditionalPut(c.Request, session, i, entry, newIDs, refMap); err != nil {
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

			if err := b.resolveConditionalPut(c.Request, session, i, entry, newIDs, refMap); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	}

	// Process references
	references, err := bundle.GetAllReferences()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for _, reference := range references {
		
		if _, alreadyMapped := refMap[reference]; alreadyMapped {
			continue
		}

		// Conditional references
		// TODO: as per spec this needs to include in-bundle resources..
		queryPos := strings.Index(reference, "?")
		if queryPos >= 0 {

			if bundle.Type != "transaction" {
				c.AbortWithError(http.StatusBadRequest, errors.New("conditional references are only allowed in transactions, not batches"))
				return
			}

			resourceType := reference[0:queryPos]
			queryString := reference[queryPos+1:]
			searchQuery := search.Query{Resource: resourceType, Query: queryString }
			ids, err := session.FindIDs(searchQuery)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, errors.Wrapf(err, "lookup of conditional reference failed (%s)", reference))
				return
			}

			if len(ids) == 1 {
				refMap[reference] = resourceType + "/" + ids[0]
			} else if len(ids) == 0 {
				c.AbortWithError(http.StatusBadRequest, errors.Errorf("no matches for conditional reference (%s)", reference))
				return
			} else {
				c.AbortWithError(http.StatusBadRequest, errors.Errorf("multiple matches for conditional reference (%s)", reference))
				return
			}
		}
	}

	// When being converted to BSON references will be updated to reflect newly assigned or conditional IDs
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

				currentResource, err := session.Get(id, entry.Resource.ResourceType())
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

	// Make the changes in the database and update the entry responses
	if (proceed) {
		for i, entry := range entries {
			err = b.doRequest(c, session, i, entry, createStatus, newIDs)
			if err != nil {
				debug("  --> ERROR %+v", err)
			}
			if entry.Response != nil {
				debug("  --> %s", entry.Response.DebugString())
			} else {
				debug("  --> nil Response")
			}
			if entry.Resource != nil {
				debug("  --> %s", entry.Resource.JsonBytes())
			} else {
				debug("  --> nil Resource")
			}
			if err != nil {
				statusCode, outcome := ErrorToOpOutcome(err)
				if bundle.Type == "transaction" {
					sendReply(c, statusCode, outcome)
					return
				} else {
					entry.Resource = nil
					entry.Request = nil
					entry.Response = &models.BundleEntryResponseComponent{
						Status: strconv.Itoa(statusCode),
						Outcome: outcome,
					}
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

		session.CommmitIfTransaction()
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

func (b *BatchController) doRequest(c *gin.Context, session DataAccessSession, i int, entry *models2.ShallowBundleEntryComponent, createStatus []string, newIDs []string) error {
	debug("doRequest %s %s", entry.Request.Method, entry.Request.Url)
	if entry.Response != nil {
		// already handled (e.g. conditional update returned 409)
		debug("  already handled (%s)", entry.Response.DebugString())
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
			if _, err := session.Delete(parts[1], parts[0]); err != nil && err != ErrNotFound {
				return errors.Wrapf(err, "failed to delete %s", entry.Request.Url)
			}
		} else {
			// It's a conditional (query-based) delete
			parts := strings.SplitN(entry.Request.Url, "?", 2)
			query := search.Query{Resource: parts[0], Query: parts[1]}
			if _, err := session.ConditionalDelete(query); err != nil {
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
			err := session.PostWithID(newIDs[i], entry.Resource)
			if err != nil {
				return errors.Wrapf(err, "failed to create %s", entry.Request.Url)
			}
			updateEntryMeta(entry)
		} else if createStatus[i] == "200" {
			// have one existing resource
			components := strings.Split(entry.FullUrl, "/")
			existingId := components[len(components)-1]

			existingResource, err := session.Get(existingId, entry.Request.Url)
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
		entry.FullUrl = b.Config.responseURL(c.Request, entry.Request.Url).String()
		parts := strings.SplitN(entry.Request.Url, "/", 2)
		if len(parts) != 2 {
			return fmt.Errorf("Couldn't identify resource and id to put from %s", entry.Request.Url)
		}

		// Write
		createdNew, err := session.Put(parts[1], "", entry.Resource)
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
	case "GET":
		/*
		examples
			1 /Patient
			2 /Patient/_search
			2 /Patient/12345
			3 /Patient/12345/_history
			4 /Patient/12345/_history/55
		*/

		pathAndQuery := strings.SplitN(entry.Request.Url, "?", 2)
		path := pathAndQuery[0]
		var queryString string
		var err error
		if len(pathAndQuery) == 2 {
			queryString = pathAndQuery[1]
			if err != nil {
				return errors.Wrapf(err, "failed to parse query string: %s", entry.Request.Url)
			}
		}

		// remove leading /
		path = strings.TrimPrefix(path, "/")
		segments := strings.Split(path, "/")
		var id, vid string
		var historyRequest bool
		resourceType := segments [0]
		debug("  segments: %q (%d)", segments, len(segments))
		if len(segments) >= 2 {
			id = segments[1]
			if id == "_search" {
				id = ""
			}
			if id == "_history" {
				return errors.Errorf("resource-level history not supported in request: %s", entry.Request.Url)
			}
			
			if len(segments) >= 3 {
				op := segments[2]
				debug("  op = %s", op)
				if op != "_history" {
					return errors.Errorf("operation not supported in request: %s", entry.Request.Url)
				}

				if len(segments) == 3 {
					historyRequest = true
				} else if len(segments) == 4 {
					vid = segments[3]
				} else {
					return errors.Errorf("failed to parse request path: %s", entry.Request.Url)
				}
			}
		}

		if historyRequest {
			baseURL := b.Config.responseURL(c.Request, resourceType)
			bundle, err := session.History(*baseURL, resourceType, id)
			debug("  history request (%s/%s) --> err %+v", resourceType, id, err)
			if err != nil && err != ErrNotFound {
				return errors.Wrapf(err, "History request failed: %s", entry.Request.Url)
			}

			if err == ErrNotFound {
				entry.Response = &models.BundleEntryResponseComponent{
					Status: "404",
				}
			} else {
				entry.Response = &models.BundleEntryResponseComponent{
					Status: "200",
				}
				entry.Resource, err = bundle.ToResource()
				if err != nil {
					return errors.Wrapf(err, "bundle.ToResource failed for request: %s", entry.Request.Url)
				}
			}
		} else if id != "" {
			// /Patient/12345
			// /Patient/12345/_history/55

			entry.Response = &models.BundleEntryResponseComponent{}
			if vid == "" {
				entry.Resource, err = session.Get(id, resourceType)
			} else {
				entry.Resource, err = session.GetVersion(id, vid, resourceType)
			}
			debug("  get resource request (%s id=%s vid=%s) --> err %+v", resourceType, id, vid, err)

			switch (err) {
			case nil:
				lastUpdated := entry.Resource.LastUpdated()
				if lastUpdated != "" {
					// entry.Response.LastModified = entry.Resource.LastUpdatedTime().UTC().Format(http.TimeFormat)
					entry.Response.LastModified = &models.FHIRDateTime {
						Time: entry.Resource.LastUpdatedTime(),
						Precision: models.Timestamp,
					}
				}
				versionId := entry.Resource.VersionId()
				if versionId != "" {
					entry.Response.Etag = "W/\"" + versionId + "\""
				}
			case ErrNotFound:
				entry.Response.Status = "404"
			case ErrDeleted:
				entry.Response.Status = "410"
			default:
				return errors.Wrapf(err, "Get/GetVersion failed for %s", entry.Request.Url)
			}

		} else {
			// Search:
			// /Patient
			// /Patient/_search
			searchQuery := search.Query{Resource: resourceType, Query: queryString }
			baseURL := b.Config.responseURL(c.Request, resourceType)
			bundle, err := session.Search(*baseURL, searchQuery)
			debug("  search request (%s %s) --> err %#v", resourceType, queryString, err)
			if err != nil {
				return errors.Wrapf(err, "Search failed for %s", entry.Request.Url)
			}
			entry.Response = &models.BundleEntryResponseComponent{
				Status: "200",
			}
			entry.Resource, err = bundle.ToResource()
			if err != nil {
				return errors.Wrapf(err, "bundle.ToResource failed for request: %s", entry.Request.Url)
			}
		}
		entry.Request = nil
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

func (b *BatchController) resolveConditionalPut(request *http.Request, session DataAccessSession, entryIndex int, entry *models2.ShallowBundleEntryComponent, newIDs []string, refMap map[string]string) error {
	// Do a preflight to either get the existing ID, get a new ID, or detect multiple matches (not allowed)
	parts := strings.SplitN(entry.Request.Url, "?", 2)
	query := search.Query{Resource: parts[0], Query: parts[1]}

	var id string
	if IDs, err := session.FindIDs(query); err == nil {
		switch len(IDs) {
		case 0:
			id = bson.NewObjectId().Hex()
		case 1:
			id = IDs[0]
		default:
			return &ErrMultipleMatches{msg: fmt.Sprintf("Multiple matches for %s (%v)", entry.Request.Url, IDs)}
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
	entry.FullUrl = b.Config.responseURL(request, query.Resource, id).String()

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
