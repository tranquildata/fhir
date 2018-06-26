package server

import (
	"strings"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/eug48/fhir/models"
	"github.com/eug48/fhir/search"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MasterSession manages a master session connected to the Mongo database.
// The user is responsible for creating and closing this master session.
// Each request to the database should first obtain a new WorkerSession,
// perform the desired operation(s), then close the WorkerSession.
type MasterSession struct {
	session *mgo.Session
	dbname  string
}

// WorkerSession is obtained from MasterSession.GetWorkerSession() and manages
// a temporary copy of the master session. A WorkerSession should be used for all
// requests to the database. When a WorkerSession is no longer needed it must
// be closed.
type WorkerSession struct {
	session *mgo.Session
	dbname  string
}

// NewMasterSession returns a new MasterSession object with an established
// session and database. Once instantiated the MasterSession object cannot
// be changed.
func NewMasterSession(session *mgo.Session, dbname string) (master *MasterSession) {
	return &MasterSession{
		session: session,
		dbname:  dbname,
	}
}

// GetWorkerSession returns a new WorkerSession with a copy of the master
// mongo session.
func (ms *MasterSession) GetWorkerSession() (worker *WorkerSession) {

	if ms.session == nil {
		return nil
	}

	return &WorkerSession{
		session: ms.session.Copy(),
		dbname:  ms.dbname,
	}
}

// SetTimeout sets the session timeout for requests made using
// the worker's session. The default timeout is 1 minute.
func (ws *WorkerSession) SetTimeout(d time.Duration) {
	if ws.session != nil {
		ws.session.SetSocketTimeout(d)
	}
}

// DB returns the mongo database available on the current session.
func (ws *WorkerSession) DB() (db *mgo.Database) {
	if ws.session != nil && ws.dbname != "" {
		db = ws.session.DB(ws.dbname)
	}
	return
}

func (ws *WorkerSession) CurrentVersionCollection(resourceType string) *mgo.Collection {
	return ws.DB().C(models.PluralizeLowerResourceName(resourceType))
}
func (ws *WorkerSession) PreviousVersionsCollection(resourceType string) *mgo.Collection {
	return ws.DB().C(models.PluralizeLowerResourceName(resourceType) + "_prev")
}

// Close closes the master session copy used by WorkerSession
func (ws *WorkerSession) Close() {
	if ws.session != nil {
		ws.session.Close()
	}
}

// NewMongoDataAccessLayer returns an implementation of DataAccessLayer that is backed by a Mongo database
func NewMongoDataAccessLayer(ms *MasterSession, interceptors map[string]InterceptorList, config Config) DataAccessLayer {
	return &mongoDataAccessLayer{
		MasterSession:     ms,
		Interceptors:      interceptors,
		countTotalResults: config.CountTotalResults,
		enableCISearches:  config.EnableCISearches,
		enableHistory:     config.EnableHistory,
		readonly:          config.ReadOnly,
	}
}

type mongoDataAccessLayer struct {
	MasterSession     *MasterSession
	Interceptors      map[string]InterceptorList
	countTotalResults bool
	enableCISearches  bool
	enableHistory     bool
	readonly          bool
}

// InterceptorList is a list of interceptors registered for a given database operation
type InterceptorList []Interceptor

// Interceptor optionally executes functions on a specified resource type before and after
// a database operation involving that resource. To register an interceptor for ALL resource
// types use a "*" as the resourceType.
type Interceptor struct {
	ResourceType string
	Handler      InterceptorHandler
}

// InterceptorHandler is an interface that defines three methods that are executed on a resource
// before the database operation, after the database operation SUCCEEDS, and after the database
// operation FAILS.
type InterceptorHandler interface {
	Before(resource interface{})
	After(resource interface{})
	OnError(err error, resource interface{})
}

// invokeInterceptorsBefore invokes the interceptor list for the given resource type before a database
// operation occurs.
func (dal *mongoDataAccessLayer) invokeInterceptorsBefore(op, resourceType string, resource interface{}) {

	for _, interceptor := range dal.Interceptors[op] {
		if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
			interceptor.Handler.Before(resource)
		}
	}
}

// invokeInterceptorsAfter invokes the interceptor list for the given resource type after a database
// operation occurs and succeeds.
func (dal *mongoDataAccessLayer) invokeInterceptorsAfter(op, resourceType string, resource interface{}) {

	for _, interceptor := range dal.Interceptors[op] {
		if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
			interceptor.Handler.After(resource)
		}
	}
}

// invokeInterceptorsOnError invokes the interceptor list for the given resource type after a database
// operation occurs and fails.
func (dal *mongoDataAccessLayer) invokeInterceptorsOnError(op, resourceType string, err error, resource interface{}) {

	for _, interceptor := range dal.Interceptors[op] {
		if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
			interceptor.Handler.OnError(err, resource)
		}
	}
}

// hasInterceptorsForOpAndType checks if any interceptors are registered for a particular database operation AND resource type
func (dal *mongoDataAccessLayer) hasInterceptorsForOpAndType(op, resourceType string) bool {

	if len(dal.Interceptors[op]) > 0 {
		for _, interceptor := range dal.Interceptors[op] {
			if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
				// At least 1 interceptor is registered for this database operation and resource type
				return true
			}
		}
	}
	return false
}

func (dal *mongoDataAccessLayer) Get(id, resourceType string) (result interface{}, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return nil, convertMongoErr(err)
	}

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	collection := worker.CurrentVersionCollection(resourceType)
	result = models.NewStructForResourceName(resourceType)
	err = collection.FindId(bsonID.Hex()).One(result)
	if err == mgo.ErrNotFound && dal.enableHistory {
		// check whether this is a deleted record
		prevCollection := worker.PreviousVersionsCollection(resourceType)
		prevQuery := bson.M{
			"_id._id":      bsonID.Hex(),
			"_id._deleted": 1,
		}
		var deletionRecord bson.D
		err = prevCollection.Find(prevQuery).Select(bson.M{"_id": 1}).Limit(1).One(&deletionRecord)
		if (err == nil) {
			return nil, ErrDeleted
		}
		return nil, convertMongoErr(err)
	}

	if err != nil {
		return nil, convertMongoErr(err)
	}
	return
}

func (dal *mongoDataAccessLayer) GetVersion(id, versionIdStr, resourceType string) (result interface{}, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return nil, convertMongoErr(err)
	}

	versionIdInt, err := strconv.Atoi(versionIdStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert versionId to an integer (%s)", versionIdStr)
	}

	result = models.NewStructForResourceName(resourceType)

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	// First assume versionId is for the current version
	curQuery := bson.M{
		"_id":            bsonID.Hex(),
		"meta.versionId": versionIdStr,
	}
	curCollection := worker.CurrentVersionCollection(resourceType)
	err = curCollection.Find(curQuery).One(result)
	// fmt.Printf("GetVersion: curQuery=%+v; err=%+v\n", curQuery, err)
	if err == mgo.ErrNotFound {
		// try to search for previous versions
		prevQuery := bson.M{
			"_id._id":      bsonID.Hex(),
			"_id._version": versionIdInt,
		}
		prevCollection := worker.PreviousVersionsCollection(resourceType)

		// have to extract resource id from vermongo formatted id field
		// first load as a bson.D
		var asBSON bson.D
		err = prevCollection.Find(prevQuery).One(&asBSON)
		if err == mgo.ErrNotFound {
			fmt.Printf("not found: %#v\n", prevQuery)
			return nil, convertMongoErr(err)
		} else if err != nil {
			return nil, errors.Wrap(err, "failed to search for a previous version")
		}

		deleted, err := unmarshalPreviousVersion(asBSON, result)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal previous version")
		}
		if deleted {
			return nil, ErrDeleted
		}
	}
	if err != nil {
		return nil, errors.Wrap(convertMongoErr(err), "failed to search for current version")
	}
	return
}

// Convert document stored in one of the _prev collections into a resource
func unmarshalPreviousVersion(asBSON bson.D, result interface{}) (deleted bool, err error) {
	// first we have to parse the vermongo-style id
	if len(asBSON) == 0 {
		return false, fmt.Errorf("unmarshalPreviousVersion: input empty")
	}
	idItem := asBSON[0]
	if idItem.Name != "_id" {
		return false, fmt.Errorf("unmarshalPreviousVersion: first element not an _id")
	}

	switch val := idItem.Value.(type) {
	case bson.D:
		if len(val) < 2 {
			return false, fmt.Errorf("unmarshalPreviousVersion: _id value size < 2")
		}
		actualId := val[0]
		if actualId.Name != "_id" {
			return false, fmt.Errorf("unmarshalPreviousVersion: _id value without first inner _id field")
		}

		switch idVal := actualId.Value.(type) {
		case string:
			// check if actually deleted
			deleted := false
			for _, idElt := range val {
				if idElt.Name == "_deleted" {
					deletedInt, isInt := idElt.Value.(int)
					if isInt {
						deleted = deletedInt > 0
					} else {
						return false, fmt.Errorf("unmarshalPreviousVersion: _id._deleted is not an integer")
					}
				}
			}
			if deleted {
				return true, nil
			}

			// convert back to a binary so we can re-unmarshal as our DTOs
			// this seems to be a limitation of mgo, hopefully will be fixed in upcoming mongo-go-driver
			asBinary, err := bson.Marshal(asBSON)
			if err != nil {
				return false, errors.Wrap(err, "failed marshalling back to bytes")
			}

			// convert to our DTO
			err = bson.Unmarshal(asBinary, result)
			if err != nil {
				return false, errors.Wrap(err, "failed to unmarshal bytes to DTO")
			}

			// set actual Id
			idField := reflect.ValueOf(result).Elem().FieldByName("Id")
			idField.Set(reflect.ValueOf(idVal))
			return false, nil
		default:
			return false, fmt.Errorf("unmarshalPreviousVersion: _id._id not a string")
		}
	default:
		return false, fmt.Errorf("unmarshalPreviousVersion: _id not a bson dictionary")
	}
}

func (dal *mongoDataAccessLayer) Post(resource interface{}) (id string, err error) {
	id = bson.NewObjectId().Hex()
	err = convertMongoErr(dal.PostWithID(id, resource))
	return
}

func (dal *mongoDataAccessLayer) ConditionalPost(query search.Query, resource interface{}) (httpStatus int, id string, outputResource interface{}, err error) {
	existingIds, err := dal.FindIDs(query)
	if err != nil {
		return
	}

	if len(existingIds) == 0 {
		httpStatus = 201
		id = bson.NewObjectId().Hex()
		err = convertMongoErr(dal.PostWithID(id, resource))
		if err == nil {
			outputResource = resource
		}

	} else if len(existingIds) == 1 {
		httpStatus = 200
		id = existingIds[0]
		outputResource, err = dal.Get(id, query.Resource)

	} else if len(existingIds) > 1 {
		httpStatus = 412
	}

	return
}

func (dal *mongoDataAccessLayer) PostWithID(id string, resource interface{}) error {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return convertMongoErr(err)
	}

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	reflect.ValueOf(resource).Elem().FieldByName("Id").SetString(bsonID.Hex())
	resourceType := reflect.TypeOf(resource).Elem().Name()
	curCollection := worker.CurrentVersionCollection(resourceType)

	if dal.enableHistory {
		versionIdInt := 1
		updateResourceMeta(resource, &versionIdInt)
	} else {
		updateResourceMeta(resource, nil)
	}

	dal.invokeInterceptorsBefore("Create", resourceType, resource)

	err = curCollection.Insert(resource)

	if err == nil {
		dal.invokeInterceptorsAfter("Create", resourceType, resource)
	} else {
		dal.invokeInterceptorsOnError("Create", resourceType, err, resource)
	}

	return convertMongoErr(err)
}

func (dal *mongoDataAccessLayer) Put(id string, resource interface{}) (createdNew bool, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return false, convertMongoErr(err)
	}

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	resourceType := reflect.TypeOf(resource).Elem().Name()
	curCollection := worker.CurrentVersionCollection(resourceType)
	reflect.ValueOf(resource).Elem().FieldByName("Id").SetString(bsonID.Hex())

	var curVersionId *int = nil
	var newVersionId *int = nil
	if dal.enableHistory {

		// get current version of this document
		var currentDoc *bson.D = nil
		if err = curCollection.FindId(bsonID.Hex()).One(&currentDoc); err != nil && err != mgo.ErrNotFound {
			return false, errors.Wrap(convertMongoErr(err), "Put handler: error retrieving current version")
		}

		// extract current version
		if currentDoc != nil {
			hasVersionId, curVersionIdTemp := getVersionIdFromResource(currentDoc)
			var newVersionIdTemp int
			if hasVersionId {
				newVersionIdTemp = curVersionIdTemp + 1
			} else {
				// for documents created by previous versions not supporting versioning or if it was disabled
				newVersionIdTemp = 1
				curVersionIdTemp = 0
			}
			newVersionId = &newVersionIdTemp
			curVersionId = &curVersionIdTemp

			// store current document in the previous version collection, adding its versionId to
			// its mongo _id like in vermongo (https://github.com/thiloplanz/v7files/wiki/Vermongo)
			//   i.e. { "_id" : { "_id" : ObjectId("4c78da..."), "_version" : "2" }
			setVermongoId(currentDoc, curVersionIdTemp)
			// NOTE: currentDoc._id modified in-place

			prevCollection := worker.PreviousVersionsCollection(resourceType)
			err = prevCollection.Insert(currentDoc) // TODO: do concurrently with the main update
			if mgo.IsDup(err) {
				/* Ignore duplicate key error - these can happen if we successfully
				insert into the prev collection but then fail to update the
				current version - they are harmless in this case since the correct
				data should have been written to prev.
				Should be fixed by MongoDB 4.0 transactions. */
				fmt.Printf("Duplicate key error in prev version collection for PUT %s/%s (likely harmless & due to prev error)\n", resourceType, id)
			} else if err != nil {
				return false, errors.Wrap(convertMongoErr(err), "failed to store previous version")
			}
		}
	}

	updateResourceMeta(resource, newVersionId)

	if dal.hasInterceptorsForOpAndType("Update", resourceType) {
		oldResource, getError := dal.Get(id, resourceType)
		if getError == nil {
			dal.invokeInterceptorsBefore("Update", resourceType, oldResource)
		}
	}

	var updated int
	if curVersionId == nil {
		var info *mgo.ChangeInfo
		selector := bson.M{"_id": bsonID.Hex()}
		info, err = curCollection.Upsert(selector, resource)
		if err != nil {
			err = errors.Wrap(err, "PUT handler: failed to upsert new document")
		}
		updated = info.Updated
	} else {
		// atomic chec-then-update
		selector := bson.M{
			"_id":            bsonID.Hex(),
			"meta.versionId": strconv.Itoa(*curVersionId),
		}
		if *curVersionId == 0 {
			// cur doc won't actually have a versionId field
			selector["meta.versionId"] = bson.M{"$exists": false}
		}
		err = curCollection.Update(selector, resource)
		if err == mgo.ErrNotFound {
			// "If the session is in safe mode (see SetSafe) a ErrNotFound error is
			// returned if a document isn't found, or a value of type *LastError
			// when some other error is detected."
			return false, fmt.Errorf("conflicting update for %+v", selector)
		} else {
			err = errors.Wrap(err, "PUT handler: failed to update current document")
		}
		updated = 1
	}

	if err == nil {
		createdNew = (updated == 0)
		if createdNew {
			dal.invokeInterceptorsAfter("Create", resourceType, resource)
		} else {
			dal.invokeInterceptorsAfter("Update", resourceType, resource)
		}
	} else {
		dal.invokeInterceptorsOnError("Update", resourceType, err, resource)
	}

	return createdNew, convertMongoErr(err)
}

func getVersionIdFromResource(doc *bson.D) (hasVersionId bool, versionId int) {
	for _, item := range *doc {
		if item.Name == "meta" {
			meta := item.Value.(bson.D)
			for _, item := range meta {
				if item.Name == "versionId" {
					hasVersionId = true
					versionIdStr, isString := item.Value.(string)
					if !isString {
						panic(fmt.Errorf("meta.versionId is not a BSON string"))
					}
					var err error
					versionId, err = strconv.Atoi(versionIdStr)
					if err == nil {
						return
					} else {
						panic(fmt.Errorf("meta.versionId BSON string is not an integer: %s", versionIdStr))
					}
				}
			}
		}
	}
	return false, -1
}

// Updates the doc to use a vermongo-like _id (_id: current_id, _version: versionId)
func setVermongoId(doc *bson.D, versionIdInt int) {
	idItem := (*doc)[0]
	if idItem.Name != "_id" {
		panic("_id field not first in bson document")
	}

	newId := bson.D{
		bson.DocElem{Name: "_id", Value: idItem.Value},
		bson.DocElem{Name: "_version", Value: versionIdInt},
	}

	(*doc)[0].Value = newId
}

func (dal *mongoDataAccessLayer) ConditionalPut(query search.Query, resource interface{}) (id string, createdNew bool, err error) {
	if IDs, err := dal.FindIDs(query); err == nil {
		switch len(IDs) {
		case 0:
			id = bson.NewObjectId().Hex()
		case 1:
			id = IDs[0]
		default:
			return "", false, ErrMultipleMatches
		}
	} else {
		return "", false, err
	}

	createdNew, err = dal.Put(id, resource)
	return id, createdNew, err
}

func (dal *mongoDataAccessLayer) Delete(id, resourceType string) (newVersionId string, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return "", convertMongoErr(err)
	}

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()
	curCollection := worker.CurrentVersionCollection(resourceType)
	prevCollection := worker.PreviousVersionsCollection(resourceType)

	if dal.enableHistory {
		newVersionId, err = saveDeletionIntoHistory(resourceType, bsonID.Hex(), curCollection, prevCollection)
		if err != nil {
			return "", errors.Wrap(err, "failed to save deletion into history")
		}
	}

	var resource interface{}
	var getError error
	hasInterceptor := dal.hasInterceptorsForOpAndType("Delete", resourceType)
	if hasInterceptor {
		// Although this is a delete operation we need to get the resource first so we can
		// run any interceptors on the resource before it's deleted.
		resource, getError = dal.Get(id, resourceType)
		dal.invokeInterceptorsBefore("Delete", resourceType, resource)
	}

	err = curCollection.RemoveId(bsonID.Hex())

	if hasInterceptor {
		if err == nil && getError == nil {
			dal.invokeInterceptorsAfter("Delete", resourceType, resource)
		} else {
			dal.invokeInterceptorsOnError("Delete", resourceType, err, resource)
		}
	}

	err = convertMongoErr(err)
	return
}

func saveDeletionIntoHistory(resourceType string, id string, curCollection *mgo.Collection, prevCollection *mgo.Collection) (newVersionIdStr string, err error) {
	// get current version of this document
	var currentDoc *bson.D = nil
	if err = curCollection.FindId(id).One(&currentDoc); err != nil && err != mgo.ErrNotFound {
		return "", errors.Wrap(convertMongoErr(err), "saveDeletionIntoHistory: error retrieving current version")
	}

	// extract current version
	if currentDoc != nil {
		hasVersionId, curVersionId := getVersionIdFromResource(currentDoc)
		var newVersionId int
		if hasVersionId {
			newVersionId = curVersionId + 1
		} else {
			// document created by previous versions not supporting versioning or if it was disabled
			newVersionId = 1
			curVersionId = 0
		}
		newVersionIdStr = strconv.Itoa(newVersionId)

		// store current document in the previous version collection, adding its versionId to
		// its mongo _id like in vermongo (https://github.com/thiloplanz/v7files/wiki/Vermongo)
		//   i.e. { "_id" : { "_id" : ObjectId("4c78da..."), "_version" : "2" }
		setVermongoId(currentDoc, curVersionId)
		// NOTE: currentDoc._id modified in-place

		// create a deletion record
		now := &models.FHIRDateTime{Time: time.Now(), Precision: models.Timestamp}
		deletionRecord := bson.M{
			"_id": bson.M{
				"_id":      id,
				"_version": newVersionId,
				"_deleted": 1,
			},
			"meta": bson.M{
				"versionId":   newVersionIdStr,
				"lastUpdated": now,
			},
		}

		err = prevCollection.Insert(currentDoc, deletionRecord) // TODO: do concurrently with the main deletion
		if mgo.IsDup(err) {
			/* Ignore duplicate key error - these can happen if we successfully
			insert into the prev collection but then fail to update the
			current version - they are harmless in this case since the correct
			data should have been written to prev.
			Should be fixed by MongoDB 4.0 transactions. */
			fmt.Printf("Duplicate key error in prev version collection for DELETE %s/%s (likely harmless & due to prev error)\n", resourceType, id)
		} else if err != nil {
			return "", errors.Wrap(convertMongoErr(err), "Delete handler: failed to store previous version")
		}
	}
	return
}

func (dal *mongoDataAccessLayer) ConditionalDelete(query search.Query) (count int, err error) {

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	IDsToDelete, err := dal.FindIDs(query)
	if err != nil {
		return 0, err
	}
	// There is the potential here for the delete to fail if the slice of IDs
	// is too large (exceeding Mongo's 16MB document size limit).
	deleteQuery := bson.M{"_id": bson.M{"$in": IDsToDelete}}
	resourceType := query.Resource
	curCollection := worker.CurrentVersionCollection(resourceType)
	prevCollection := worker.PreviousVersionsCollection(resourceType)

	hasInterceptors := dal.hasInterceptorsForOpAndType("Delete", resourceType)

	if hasInterceptors || dal.enableHistory {
		/* Interceptors for a conditional delete are tricky since an interceptor is only run
		   AFTER the database operation and only on resources that were SUCCESSFULLY deleted. We use
		   the following approach:
		   1. Bulk delete those resources by ID
		   2. Search again using the SAME query, to verify that those resources were in fact deleted
		   3. Run the interceptor(s) on all resources that ARE NOT in the second search (since they were truly deleted)
		*/

		// get the resources that are about to be deleted
		var bundle *models.Bundle
		bundle, err = dal.Search(url.URL{}, query) // the baseURL argument here does not matter

		if err == nil {
			for _, elem := range bundle.Entry {
				if (hasInterceptors) {
					dal.invokeInterceptorsBefore("Delete", resourceType, elem.Resource)
				}
			}

			for _, elem := range bundle.Entry {
				if dal.enableHistory {
					idField := reflect.ValueOf(elem.Resource).Elem().FieldByName("Id")
					idKind := idField.Kind()
					id := idField.String()
					if idKind.String() != "string" {
						return count, errors.Wrapf(err, "failed to save deletion into history (%s/%s) - id not a string!", resourceType, id)
					}
					_, err = saveDeletionIntoHistory(resourceType, id, curCollection, prevCollection)
					if err != nil {
						return count, errors.Wrapf(err, "failed to save deletion into history (%s/%s)", resourceType, id)
					}
				}
			}

			// Do the bulk delete by ID.
			info, err := curCollection.RemoveAll(deleteQuery)
			deletedIds := make([]string, len(IDsToDelete))
			if info != nil {
				count = info.Removed
			}

			if err != nil {
				if hasInterceptors {
					for _, elem := range bundle.Entry {
						dal.invokeInterceptorsOnError("Delete", resourceType, err, elem.Resource)
					}
				}
				return count, convertMongoErr(err)
			} else if (hasInterceptors == false) {
				return count, nil
			}

			var searchErr error

			if count < len(IDsToDelete) {
				// Some but not all resources were removed, so use the original search query
				// to see which resources are left.
				var failBundle *models.Bundle
				failBundle, searchErr = dal.Search(url.URL{}, query)
				deletedIds = setDiff(IDsToDelete, getResourceIdsFromBundle(failBundle))
			} else {
				// All resources were successfully removed
				deletedIds = IDsToDelete
			}

			if searchErr == nil {
				for _, elem := range bundle.Entry {
					id := reflect.ValueOf(elem.Resource).Elem().FieldByName("Id").String()

					if elementInSlice(id, deletedIds) {
						// This resource was confirmed deleted
						dal.invokeInterceptorsAfter("Delete", resourceType, elem.Resource)
					} else {
						// This resource was not confirmed deleted, which is an error
						resourceErr := fmt.Errorf("ConditionalDelete: failed to delete resource %s with ID %s", resourceType, id)
						dal.invokeInterceptorsOnError("Delete", resourceType, resourceErr, elem.Resource)
					}
				}
			}
		}
		return count, convertMongoErr(err)
	} else {
		// do the bulk delete the usual way
		info, err := curCollection.RemoveAll(deleteQuery)
		if info != nil {
			count = info.Removed
		}
		return count, convertMongoErr(err)
	}
}

func (dal *mongoDataAccessLayer) History(baseURL url.URL, resourceType string, id string) (result *models.Bundle, err error) {
	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	baseURLstr := baseURL.String()
	if !strings.HasSuffix(baseURLstr, "/") {
		baseURLstr = baseURLstr + "/"
	}
	fullUrl := baseURLstr + id

	prevCollection := worker.PreviousVersionsCollection(resourceType)
	curCollection := worker.CurrentVersionCollection(resourceType)

	var entryList []models.BundleEntryComponent
	method := "POST"
	makeEntryRequest := func(actualMethod string) (*models.BundleEntryRequestComponent) {
		url := ""
		if method == "POST" {
			url = resourceType
		} else {
			url = resourceType + "/" + id
		}
		return &models.BundleEntryRequestComponent{Url: url, Method: actualMethod}
	}

	// sort - oldest versions last
	iter := prevCollection.Find(bson.M{"_id._id": id}).Sort("_id._version").Iter()

	var prevDocBson bson.D
	for iter.Next(&prevDocBson) {

		var entry models.BundleEntryComponent
		entry.FullUrl = fullUrl

		resource := models.NewStructForResourceName(resourceType)
		deleted, err := unmarshalPreviousVersion(prevDocBson, resource)
		if err != nil {
			return nil, err
		}
		if deleted {
			entry.Request = makeEntryRequest("DELETE")
		} else {
			entry.Resource = resource
			entry.Request = makeEntryRequest(method)
			method = "PUT"
		}

		entryList = append(entryList, entry)
	}
	if err := iter.Close(); err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	// add current version
	var curDoc interface{}
	curDoc = models.NewStructForResourceName(resourceType)
	err = curCollection.Find(bson.M{"_id": id}).One(curDoc)
	if err == nil {
		var entry models.BundleEntryComponent
		entry.FullUrl = fullUrl
		entry.Resource = &curDoc
		entry.Request = makeEntryRequest(method)
		entryList = append(entryList, entry)
	} else if err != mgo.ErrNotFound {
		return nil, err
	}

	// output a Bundle
	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "history"
	bundle.Entry = entryList
	totalDocs := uint32(len(entryList))
	bundle.Total = &totalDocs

	// TODO: use paging
	// bundle.Link = dal.generatePagingLinks(baseURL, searchQuery, total, uint32(numResults))

	return &bundle, nil
}

func (dal *mongoDataAccessLayer) Search(baseURL url.URL, searchQuery search.Query) (*models.Bundle, error) {

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	searcher := search.NewMongoSearcher(worker.DB(), dal.countTotalResults, dal.enableCISearches, dal.readonly)

	result, total, err := searcher.Search(searchQuery)
	if err != nil {
		return nil, convertMongoErr(err)
	}

	includesMap := make(map[string]interface{})
	var entryList []models.BundleEntryComponent
	resultVal := reflect.ValueOf(result).Elem()
	numResults := resultVal.Len()
	baseURLstr := baseURL.String()
	if !strings.HasSuffix(baseURLstr, "/") {
		baseURLstr = baseURLstr + "/"
	}

	for i := 0; i < numResults; i++ {
		var entry models.BundleEntryComponent
		entry.Resource = resultVal.Index(i).Addr().Interface()
		entry.FullUrl = baseURLstr + resultVal.Index(i).FieldByName("Id").String()
		entry.Search = &models.BundleEntrySearchComponent{Mode: "match"}
		entryList = append(entryList, entry)

		if searchQuery.UsesIncludes() || searchQuery.UsesRevIncludes() {
			rpi, ok := entry.Resource.(ResourcePlusRelatedResources)
			if ok {
				for k, v := range rpi.GetIncludedAndRevIncludedResources() {
					includesMap[k] = v
				}
			}
		}
	}

	for _, v := range includesMap {
		var entry models.BundleEntryComponent
		entry.Resource = v
		entry.Search = &models.BundleEntrySearchComponent{Mode: "include"}
		entryList = append(entryList, entry)
	}

	var bundle models.Bundle
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Type = "searchset"
	bundle.Entry = entryList

	// Only include the total if counts are enabled, or if _summary=count was applied.
	if dal.countTotalResults || searchQuery.Options().Summary == "count" {
		bundle.Total = &total
	}

	bundle.Link = dal.generatePagingLinks(baseURL, searchQuery, total, uint32(numResults))

	return &bundle, nil
}

func (dal *mongoDataAccessLayer) FindIDs(searchQuery search.Query) (IDs []string, err error) {

	worker := dal.MasterSession.GetWorkerSession()
	defer worker.Close()

	// First create a new query with the unsupported query options filtered out
	oldParams := searchQuery.URLQueryParameters(false)
	newParams := search.URLQueryParameters{}
	for _, param := range oldParams.All() {
		switch param.Key {
		case search.ContainedParam, search.ContainedTypeParam, search.ElementsParam, search.IncludeParam,
			search.RevIncludeParam, search.SummaryParam:
			continue
		default:
			newParams.Add(param.Key, param.Value)
		}
	}
	newQuery := search.Query{Resource: searchQuery.Resource, Query: newParams.Encode()}

	// Now search on that query, unmarshaling to a temporary struct and converting results to []string
	searcher := search.NewMongoSearcher(worker.DB(), dal.countTotalResults, dal.enableCISearches, dal.readonly)
	results, _, err := searcher.Search(newQuery)
	if err != nil {
		return nil, convertMongoErr(err)
	}

	resultsVal := reflect.ValueOf(results).Elem()
	IDs = make([]string, resultsVal.Len())

	for i := 0; i < resultsVal.Len(); i++ {
		IDs[i] = resultsVal.Index(i).FieldByName("Id").String()
	}

	return IDs, nil
}

func (dal *mongoDataAccessLayer) generatePagingLinks(baseURL url.URL, query search.Query, total uint32, numResults uint32) []models.BundleLinkComponent {

	links := make([]models.BundleLinkComponent, 0, 5)
	params := query.URLQueryParameters(true)
	offset := 0
	if pOffset := params.Get(search.OffsetParam); pOffset != "" {
		offset, _ = strconv.Atoi(pOffset)
		if offset < 0 {
			offset = 0
		}
	}
	count := search.NewQueryOptions().Count
	if pCount := params.Get(search.CountParam); pCount != "" {
		count, _ = strconv.Atoi(pCount)
		if count < 1 {
			count = search.NewQueryOptions().Count
		}
	}

	// For queries that don't support paging, only return the "self" link created directly from the original query.
	if !query.SupportsPaging() {
		links = append(links, newRawSelfLink(baseURL, query))
		return links
	}

	// Self link
	links = append(links, newLink("self", baseURL, params, offset, count))

	// First link
	links = append(links, newLink("first", baseURL, params, 0, count))

	// Previous link
	if offset > 0 {
		prevOffset := offset - count
		// Handle case where paging is uneven (e.g., count=10&offset=5)
		if prevOffset < 0 {
			prevOffset = 0
		}
		prevCount := offset - prevOffset
		links = append(links, newLink("previous", baseURL, params, prevOffset, prevCount))
	}

	// If counts are enabled, the total is accurate and can be used to compute the links.
	if dal.countTotalResults {
		// Next Link
		if total > uint32(offset+count) {
			nextOffset := offset + count
			links = append(links, newLink("next", baseURL, params, nextOffset, count))
		}

		// Last Link
		remainder := (int(total) - offset) % count
		if int(total) < offset {
			remainder = 0
		}
		newOffset := int(total) - remainder
		if remainder == 0 && int(total) > count {
			newOffset = int(total) - count
		}
		links = append(links, newLink("last", baseURL, params, newOffset, count))

	} else {
		// Otherwise, we can only use the number of results returned by the search, and compare
		// it to the expected paging count to determine if we've exhaused the search results or not.

		// Next Link
		if int(numResults) == count {
			nextOffset := offset + count
			links = append(links, newLink("next", baseURL, params, nextOffset, count))
		}

		// Last Link
		// Without a total there is no way to compute the Last link. However, this still conforms
		// to RFC 5005 (https://tools.ietf.org/html/rfc5005).
	}

	return links
}

// ResourcePlusRelatedResources is an interface to capture those structs that implement the functions for
// getting included and rev-included resources
type ResourcePlusRelatedResources interface {
	GetIncludedAndRevIncludedResources() map[string]interface{}
	GetIncludedResources() map[string]interface{}
	GetRevIncludedResources() map[string]interface{}
}

func newRawSelfLink(baseURL url.URL, query search.Query) models.BundleLinkComponent {
	queryString := ""
	if len(query.Query) > 0 {
		queryString = "?" + query.Query
	}

	return models.BundleLinkComponent{
		Relation: "self",
		Url:      baseURL.String() + queryString,
	}
}

func newLink(relation string, baseURL url.URL, params search.URLQueryParameters, offset int, count int) models.BundleLinkComponent {
	params.Set(search.OffsetParam, strconv.Itoa(offset))
	params.Set(search.CountParam, strconv.Itoa(count))
	baseURL.RawQuery = params.Encode()
	return models.BundleLinkComponent{Relation: relation, Url: baseURL.String()}
}

func convertIDToBsonID(id string) (bson.ObjectId, error) {
	if bson.IsObjectIdHex(id) {
		return bson.ObjectIdHex(id), nil
	}
	return bson.ObjectId(""), models.NewOperationOutcome("fatal", "exception", "Id must be a valid BSON ObjectId")
}

func updateResourceMeta(resource interface{}, versionId *int) {
	m := reflect.ValueOf(resource).Elem().FieldByName("Meta")
	if m.IsNil() {
		newMeta := &models.Meta{}
		m.Set(reflect.ValueOf(newMeta))
	}
	now := &models.FHIRDateTime{Time: time.Now(), Precision: models.Timestamp}
	m.Elem().FieldByName("LastUpdated").Set(reflect.ValueOf(now))

	if versionId != nil {
		versionIdStr := strconv.Itoa(*versionId)
		m.Elem().FieldByName("VersionId").SetString(versionIdStr)
	}
}

func convertMongoErr(err error) error {
	switch err {
	default:
		return err
	case mgo.ErrNotFound:
		return ErrNotFound
	}
}

// getResourceIdsFromBundle parses a slice of BSON resource IDs from a valid
// bundle of resources (typically returned from a search operation). Order is
// preserved.
func getResourceIdsFromBundle(bundle *models.Bundle) []string {
	resourceIds := make([]string, int(*bundle.Total))
	for i, elem := range bundle.Entry {
		resourceIds[i] = reflect.ValueOf(elem.Resource).Elem().FieldByName("Id").String()
	}
	return resourceIds
}

// setDiff returns all the elements in slice X that are not in slice Y
func setDiff(X, Y []string) []string {
	m := make(map[string]int)

	for _, y := range Y {
		m[y]++
	}

	var ret []string
	for _, x := range X {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}

// elementInSlice tests if a string element is in a larger slice of strings
func elementInSlice(element string, slice []string) bool {
	for _, el := range slice {
		if element == el {
			return true
		}
	}
	return false
}
