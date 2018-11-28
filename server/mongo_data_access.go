package server

import (
	"context"
	"fmt"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/eug48/fhir/models"
	"github.com/eug48/fhir/models2"
	"github.com/eug48/fhir/search"
	"github.com/pkg/errors"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/mongodb/mongo-go-driver/mongo/replaceopt"
)

type mongoDataAccessLayer struct {
	client            *mongo.Client
	defaultDbName     string
	enableMultiDB     bool
	dbSuffix          string
	Interceptors      map[string]InterceptorList
	countTotalResults bool
	enableCISearches  bool
	enableHistory     bool
	readonly          bool
}

type mongoSession struct {
	session       *mongo.Session
	db            *mongo.Database
	dal           *mongoDataAccessLayer
	inTransaction bool
}

func (dal *mongoDataAccessLayer) StartSession(customDbName string) DataAccessSession {
	session, err := dal.client.StartSession()
	if err != nil {
		panic(errors.Wrap(err, "StartSession failed"))
	}

	var dbName string
	if dal.enableMultiDB && customDbName != "" {
		if dal.dbSuffix != "" && !strings.HasSuffix(customDbName, dal.dbSuffix) {
			panic(errors.Wrapf(err, "database name (%s) doesn't end with suffix (%s)", customDbName, dal.dbSuffix))
		}
		dbName = customDbName
	} else {
		dbName = dal.defaultDbName
	}

	db := dal.client.Database(dbName)
	if db == nil {
		panic(errors.Wrap(err, "client.Database failed"))
	}

	return &mongoSession{
		session:       session,
		db:            db,
		inTransaction: false,
		dal:           dal,
	}
}

func (ms *mongoSession) CurrentVersionCollection(resourceType string) *mongo.Collection {
	return ms.db.Collection(models.PluralizeLowerResourceName(resourceType))
}
func (ms *mongoSession) PreviousVersionsCollection(resourceType string) *mongo.Collection {
	return ms.db.Collection(models.PluralizeLowerResourceName(resourceType) + "_prev")
}

func (ms *mongoSession) StartTransaction() error {
	if ms.inTransaction {
		// sucess if already in a transaction
		return nil
	}

	err := ms.session.StartTransaction()
	ms.debug("StartTransaction")
	if err == nil {
		ms.inTransaction = true
	}
	return errors.Wrap(err, "mongoSession.StartTransaction")
}
func (ms *mongoSession) CommmitIfTransaction() error {
	if ms.inTransaction {
		err := ms.session.CommitTransaction(context.TODO())
		ms.debug("CommmitTransaction")
		if err == nil {
			ms.inTransaction = false
		}
		return errors.Wrap(err, "mongoSession.CommmitIfTransaction")
	} else {
		return nil
	}
}
func (ms *mongoSession) Finish() {
	var err error
	if ms.inTransaction {
		err = ms.session.AbortTransaction(context.TODO())
		ms.debug("AbortTransaction called from mongoSession.Finish")
		if err == nil {
			ms.inTransaction = false
		}
	}
	ms.session.EndSession(context.TODO())
	if err != nil {
		panic(errors.Wrap(err, "session.Finish error"))
	}
}

// NewMongoDataAccessLayer returns an implementation of DataAccessLayer that is backed by a Mongo database
func NewMongoDataAccessLayer(client *mongo.Client, defaultDbName string, enableMultiDB bool, dbSuffix string, interceptors map[string]InterceptorList, config Config) DataAccessLayer {
	return &mongoDataAccessLayer{
		client:            client,
		defaultDbName:     defaultDbName,
		enableMultiDB:     enableMultiDB,
		dbSuffix:          dbSuffix,
		Interceptors:      interceptors,
		countTotalResults: config.CountTotalResults,
		enableCISearches:  config.EnableCISearches,
		enableHistory:     config.EnableHistory,
		readonly:          config.ReadOnly,
	}
}

func (ms *mongoSession) debug(format string, a ...interface{}) {
	// fmt.Printf("[mongo] " + format + "\n", a...)
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
func (ms *mongoSession) invokeInterceptorsBefore(op, resourceType string, resource interface{}) {

	for _, interceptor := range ms.dal.Interceptors[op] {
		if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
			interceptor.Handler.Before(resource)
		}
	}
}

// invokeInterceptorsAfter invokes the interceptor list for the given resource type after a database
// operation occurs and succeeds.
func (ms *mongoSession) invokeInterceptorsAfter(op, resourceType string, resource interface{}) {

	for _, interceptor := range ms.dal.Interceptors[op] {
		if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
			interceptor.Handler.After(resource)
		}
	}
}

// invokeInterceptorsOnError invokes the interceptor list for the given resource type after a database
// operation occurs and fails.
func (ms *mongoSession) invokeInterceptorsOnError(op, resourceType string, err error, resource interface{}) {

	for _, interceptor := range ms.dal.Interceptors[op] {
		if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
			interceptor.Handler.OnError(err, resource)
		}
	}
}

// hasInterceptorsForOpAndType checks if any interceptors are registered for a particular database operation AND resource type
func (ms *mongoSession) hasInterceptorsForOpAndType(op, resourceType string) bool {

	if len(ms.dal.Interceptors[op]) > 0 {
		for _, interceptor := range ms.dal.Interceptors[op] {
			if interceptor.ResourceType == resourceType || interceptor.ResourceType == "*" {
				// At least 1 interceptor is registered for this database operation and resource type
				return true
			}
		}
	}
	return false
}

func (ms *mongoSession) Get(id, resourceType string) (resource *models2.Resource, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return nil, ErrNotFound
	}

	collection := ms.CurrentVersionCollection(resourceType)
	filter := bson.NewDocument(bson.EC.String("_id", bsonID.Hex()))
	var doc bson.Document
	err = collection.FindOne(context.TODO(), filter, ms.session).Decode(&doc)
	ms.debug("Get %s/%s --> %s (err %+v)", resourceType, id, doc.String(), err)
	if err == mongo.ErrNoDocuments && ms.dal.enableHistory {
		// check whether this is a deleted record
		prevCollection := ms.PreviousVersionsCollection(resourceType)
		prevQuery := bson.NewDocument(
			bson.EC.String("_id._id", bsonID.Hex()),
			bson.EC.Int32("_id._deleted", 1),
		)
		idOnly := bson.NewDocument(bson.EC.Int32("_id", 1))
		cursor, err := prevCollection.Find(context.TODO(), prevQuery, ms.session, findopt.Limit(1), findopt.Projection(idOnly))
		if err != nil {
			return nil, errors.Wrap(err, "Get --> prevCollection.Find")
		}

		deleted := cursor.Next(context.TODO())
		err = cursor.Err()
		ms.debug("   deleted version: %t (err %+v)", deleted, err)
		if err != nil {
			return nil, errors.Wrap(err, "Get --> prevCollection.Find --> cursor error")
		}

		if deleted {
			return nil, ErrDeleted
		}
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, convertMongoErr(err)
	}

	resource, err = models2.NewResourceFromBSON2(&doc)
	return
}

func (ms *mongoSession) GetVersion(id, versionIdStr, resourceType string) (resource *models2.Resource, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return nil, ErrNotFound
	}

	versionIdInt, err := strconv.Atoi(versionIdStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert versionId to an integer (%s)", versionIdStr)
	}

	// First assume versionId is for the current version
	curQuery := bson.NewDocument(
		bson.EC.String("_id", bsonID.Hex()),
		bson.EC.String("meta.versionId", versionIdStr),
	)
	curCollection := ms.CurrentVersionCollection(resourceType)
	var result bson.Document
	err = curCollection.FindOne(context.TODO(), curQuery, ms.session).Decode(&result)
	// fmt.Printf("GetVersion: curQuery=%+v; err=%+v\n", curQuery, err)

	if err == mongo.ErrNoDocuments {
		// try to search for previous versions
		prevQuery := bson.NewDocument(
			bson.EC.String("_id._id", bsonID.Hex()),
			bson.EC.Int32("_id._version", int32(versionIdInt)),
		)
		prevCollection := ms.PreviousVersionsCollection(resourceType)
		cur, err := prevCollection.Find(context.TODO(), prevQuery, ms.session, findopt.Limit(1))
		if err != nil {
			return nil, errors.Wrap(err, "GetVersion --> prevCollection.Find")
		}

		if cur.Next(context.TODO()) {

			var prevDoc bson.Document
			err = cur.Decode(&prevDoc)
			if err != nil {
				return nil, errors.Wrap(err, "GetVersion --> prevCollection.Find --> Decode")
			}

			var deleted bool
			deleted, resource, err = unmarshalPreviousVersion(&prevDoc)
			if err != nil {
				return nil, errors.Wrap(err, "failed to unmarshal previous version")
			}
			if deleted {
				return nil, ErrDeleted
			} else {
				return resource, nil
			}
		}
		if err := cur.Err(); err != nil {
			return nil, errors.Wrap(err, "GetVersion --> prevCollection.Find --> cursor error")
		}

		return nil, ErrNotFound

	} else if err != nil {
		return nil, errors.Wrap(convertMongoErr(err), "failed to search for current version")
	} else {
		resource, err = models2.NewResourceFromBSON2(&result)
	}

	return
}

// Convert document stored in one of the _prev collections into a resource
func unmarshalPreviousVersion(asBSON *bson.Document) (deleted bool, resource *models2.Resource, err error) {
	// fmt.Printf("[unmarshalPreviousVersion] %+v\n", asBSON)
	// first we have to parse the vermongo-style id
	if asBSON.Len() == 0 {
		return false, nil, fmt.Errorf("unmarshalPreviousVersion: input empty")
	}
	idItem := asBSON.ElementAt(0)
	if idItem.Key() != "_id" {
		return false, nil, fmt.Errorf("unmarshalPreviousVersion: first element not an _id")
	}
	val := idItem.Value()
	switch val.Type() {
	case bson.TypeEmbeddedDocument:
		idDoc := val.MutableDocument()
		if idDoc.Len() < 2 {
			return false, nil, fmt.Errorf("unmarshalPreviousVersion: _id value size < 2")
		}
		actualId := idDoc.ElementAt(0)
		// fmt.Printf("[unmarshalPreviousVersion] ACTUAL %+v\n", actualId)
		if actualId.Key() != "_id" {
			return false, nil, fmt.Errorf("unmarshalPreviousVersion: _id value without first inner _id field")
		}

		switch idVal := actualId.Value().Interface().(type) {
		case string:
			// check if actually deleted
			deleted := false
			for i := 0; i < idDoc.Len(); i++ {
				idElt := idDoc.ElementAt(uint(i))
				if idElt.Key() == "_deleted" {
					deletedInt, isInt := idElt.Value().Int32OK()
					if isInt {
						deleted = deletedInt > 0
					} else {
						return false, nil, fmt.Errorf("unmarshalPreviousVersion: _id._deleted is not an integer")
					}
				}
			}
			if deleted {
				return true, nil, nil
			}

			// create a BSON doc with a string id
			stringIdElem := bson.EC.String("_id", idVal)
			bsonWithStringId := asBSON.Copy().Set(stringIdElem)

			// convert to JSON
			resource, err = models2.NewResourceFromBSON2(bsonWithStringId)
			if err != nil {
				return false, nil, errors.Wrap(err, "unmarshalPreviousVersion: NewResourceFromBSON failed")
			}

			return false, resource, nil
		default:
			return false, nil, fmt.Errorf("unmarshalPreviousVersion: _id._id not a string")
		}
	default:
		return false, nil, fmt.Errorf("unmarshalPreviousVersion: _id not a bson dictionary")
	}
}

func (ms *mongoSession) Post(resource *models2.Resource) (id string, err error) {
	id = objectid.New().Hex()
	err = convertMongoErr(ms.PostWithID(id, resource))
	return
}

func (ms *mongoSession) ConditionalPost(query search.Query, resource *models2.Resource) (httpStatus int, id string, outputResource *models2.Resource, err error) {
	existingIds, err := ms.FindIDs(query)
	if err != nil {
		return
	}

	if len(existingIds) == 0 {
		httpStatus = 201
		id = objectid.New().Hex()
		err = convertMongoErr(ms.PostWithID(id, resource))
		if err == nil {
			outputResource = resource
		}

	} else if len(existingIds) == 1 {
		httpStatus = 200
		id = existingIds[0]
		outputResource, err = ms.Get(id, query.Resource)

	} else if len(existingIds) > 1 {
		httpStatus = 412
	}

	return
}

func (ms *mongoSession) PostWithID(id string, resource *models2.Resource) error {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return convertMongoErr(err)
	}

	ms.debug("PostWithID: updating %s", resource)
	resource.SetId(bsonID.Hex())
	updateResourceMeta(resource, 1)
	resourceType := resource.ResourceType()
	curCollection := ms.CurrentVersionCollection(resourceType)

	ms.invokeInterceptorsBefore("Create", resourceType, resource)

	ms.debug("PostWithID: inserting %s", resource)
	_, err = curCollection.InsertOne(context.TODO(), resource, ms.session)

	if err == nil {
		ms.invokeInterceptorsAfter("Create", resourceType, resource)
	} else {
		ms.invokeInterceptorsOnError("Create", resourceType, err, resource)
	}

	return convertMongoErr(err)
}

func (ms *mongoSession) Put(id string, conditionalVersionId string, resource *models2.Resource) (createdNew bool, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return false, convertMongoErr(err)
	}

	resourceType := resource.ResourceType()
	curCollection := ms.CurrentVersionCollection(resourceType)
	resource.SetId(bsonID.Hex())
	if conditionalVersionId != "" {
		ms.debug("PUT %s/%s (If-Match %s)", resourceType, resource.Id(), conditionalVersionId)
	} else {
		ms.debug("PUT %s/%s", resourceType, resource.Id())
	}

	var curVersionId *int = nil
	var newVersionId = 1
	if ms.dal.enableHistory == false {
		if conditionalVersionId != "" {
			return false, errors.Errorf("If-Match specified for a conditional put, but version histories are disabled")
		}
		ms.debug("  versionIds: history disabled; new %d", newVersionId)
	} else {

		// get current version of this document
		var currentDoc bson.Document
		currentDocQuery := bson.NewDocument(bson.EC.String("_id", bsonID.Hex()))
		if err = curCollection.FindOne(context.TODO(), currentDocQuery, ms.session).Decode(&currentDoc); err != nil && err != mongo.ErrNoDocuments {
			return false, errors.Wrap(convertMongoErr(err), "Put handler: error retrieving current version")
		}

		if err == mongo.ErrNoDocuments {
			if conditionalVersionId != "" {
				return false, ErrConflict{msg: "If-Match specified for a resource that doesn't exist"}
			}
			ms.debug("  versionIds: no current; new %d", newVersionId)
		} else {
			hasVersionId, curVersionIdTemp, curVersionIdStr := getVersionIdFromResource(&currentDoc)
			if hasVersionId {
				newVersionId = curVersionIdTemp + 1
			} else {
				// for documents created by previous versions not supporting versioning or if it was disabled
				newVersionId = 1
				curVersionIdTemp = 0
			}
			curVersionId = &curVersionIdTemp
			ms.debug("  versionIds: current %d; new %d", *curVersionId, newVersionId)

			if conditionalVersionId != "" && conditionalVersionId != curVersionIdStr {
				return false, ErrConflict{msg: "If-Match doesn't match current versionId"}
			}

			// store current document in the previous version collection, adding its versionId to
			// its mongo _id like in vermongo (https://github.com/thiloplanz/v7files/wiki/Vermongo)
			//   i.e. { "_id" : { "_id" : ObjectId("4c78da..."), "_version" : "2" }
			setVermongoId(&currentDoc, curVersionIdTemp)
			// NOTE: currentDoc._id modified in-place

			prevCollection := ms.PreviousVersionsCollection(resourceType)
			_, err := prevCollection.InsertOne(context.TODO(), &currentDoc, ms.session) // TODO: do concurrently with the main update
			if err != nil && strings.Contains(err.Error(), "duplicate key") {
				/* Ignore duplicate key error - these can happen if we successfully
				insert into the prev collection but then fail to update the
				current version - they are harmless in this case since the correct
				data should have been written to prev.
				Should be fixed by MongoDB 4.0 transactions. */
				ms.debug("Duplicate key error in prev version collection for PUT %s/%s (likely harmless & due to prev error)\n", resourceType, id)
				fmt.Printf("Duplicate key error in prev version collection for PUT %s/%s (likely harmless & due to prev error)\n", resourceType, id)
			} else if err != nil {
				return false, errors.Wrap(convertMongoErr(err), "failed to store previous version")
			}
		}
	}

	updateResourceMeta(resource, newVersionId)

	if ms.hasInterceptorsForOpAndType("Update", resourceType) {
		oldResource, getError := ms.Get(id, resourceType)
		if getError == nil {
			ms.invokeInterceptorsBefore("Update", resourceType, oldResource)
		}
	}

	var updated int64
	if curVersionId == nil {
		var info *mongo.UpdateResult
		selector := bson.NewDocument(bson.EC.String("_id", bsonID.Hex()))
		info, err = curCollection.ReplaceOne(context.TODO(), selector, resource, replaceopt.Upsert(true), ms.session)
		ms.debug("   upsert %#v", selector)
		if err != nil {
			bson, err2 := resource.GetBSON()
			if err2 != nil {
				panic(err2)
			}
			err = errors.Wrapf(err, "PUT handler: failed to upsert new document: %#v --> %s %#v", selector, resource.JsonBytes(), bson)
		} else {
			updated = info.ModifiedCount
		}
	} else {
		// atomic check-then-update
		selector := bson.NewDocument(
			bson.EC.String("_id", bsonID.Hex()),
			bson.EC.String("meta.versionId", strconv.Itoa(*curVersionId)),
		)
		if *curVersionId == 0 {
			// cur doc won't actually have a versionId field
			selector.Set(bson.EC.SubDocumentFromElements("meta.versionId", bson.EC.Boolean("$exists", false)))
		}
		var updateOneInfo *mongo.UpdateResult
		updateOneInfo, err = curCollection.ReplaceOne(context.TODO(), selector, resource, ms.session)
		ms.debug("   update %#v --> %#v (err %#v)", selector.String(), updateOneInfo, err)
		if err != nil {
			err = errors.Wrap(err, "PUT handler: failed to update current document")
		} else if updateOneInfo.ModifiedCount == 0 {
			return false, fmt.Errorf("conflicting update for %+v", selector)
		}
		updated = 1
	}
	if updated == 0 {
		ms.debug("      created new")
	} else {
		ms.debug("      updated %d", updated)
	}

	if err == nil {
		createdNew = (updated == 0)
		if createdNew {
			ms.invokeInterceptorsAfter("Create", resourceType, resource)
		} else {
			ms.invokeInterceptorsAfter("Update", resourceType, resource)
		}
	} else {
		ms.invokeInterceptorsOnError("Update", resourceType, err, resource)
	}

	return createdNew, convertMongoErr(err)
}

func getVersionIdFromResource(doc *bson.Document) (hasVersionId bool, versionIdInt int, versionIdStr string) {
	versionId, err := doc.LookupErr("meta", "versionId")
	if err == bson.ErrElementNotFound {
		return false, -1, ""
	} else if err != nil {
		panic(errors.Wrap(err, "getVersionIdFromResource LookupErr failed"))
	}

	hasVersionId = true
	var isString bool
	versionIdStr, isString = versionId.StringValueOK()
	if !isString {
		panic(errors.Errorf("meta.versionId is not a BSON string"))
	}
	versionIdInt, err = strconv.Atoi(versionIdStr)
	if err == nil {
		return
	} else {
		panic(errors.Errorf("meta.versionId BSON string is not an integer: %s", versionIdStr))
	}
}

// Updates the doc to use a vermongo-like _id (_id: current_id, _version: versionId)
func setVermongoId(doc *bson.Document, versionIdInt int) {
	idItem := doc.ElementAt(0)
	if idItem == nil || idItem.Key() != "_id" {
		panic("_id field not first in bson document")
	}

	newId := bson.NewDocument(
		bson.EC.Interface("_id", idItem.Value().Interface()),
		bson.EC.Int32("_version", int32(versionIdInt)),
	)

	doc.Set(bson.EC.SubDocument("_id", newId))
}

func (ms *mongoSession) ConditionalPut(query search.Query, conditionalVersionId string, resource *models2.Resource) (id string, createdNew bool, err error) {
	if IDs, err := ms.FindIDs(query); err == nil {
		switch len(IDs) {
		case 0:
			id = objectid.New().Hex()
		case 1:
			id = IDs[0]
		default:
			return "", false, &ErrMultipleMatches{msg: fmt.Sprintf("Multiple matches for %s?%s", query.Resource, query.Query)}
		}
	} else {
		return "", false, err
	}

	createdNew, err = ms.Put(id, conditionalVersionId, resource)
	return id, createdNew, err
}

func (ms *mongoSession) Delete(id, resourceType string) (newVersionId string, err error) {
	bsonID, err := convertIDToBsonID(id)
	if err != nil {
		return "", ErrNotFound
	}

	curCollection := ms.CurrentVersionCollection(resourceType)
	prevCollection := ms.PreviousVersionsCollection(resourceType)

	if ms.dal.enableHistory {
		newVersionId, err = saveDeletionIntoHistory(resourceType, bsonID.Hex(), curCollection, prevCollection, ms.session)
		if err == mongo.ErrNoDocuments {
			return "", ErrNotFound
		} else if err != nil {
			return "", errors.Wrap(err, "failed to save deletion into history")
		}
	}

	var resource interface{}
	var getError error
	hasInterceptor := ms.hasInterceptorsForOpAndType("Delete", resourceType)
	if hasInterceptor {
		// Although this is a delete operation we need to get the resource first so we can
		// run any interceptors on the resource before it's deleted.
		resource, getError = ms.Get(id, resourceType)
		ms.invokeInterceptorsBefore("Delete", resourceType, resource)
	}

	filter := bson.NewDocument(bson.EC.String("_id", bsonID.Hex()))
	deleteInfo, err := curCollection.DeleteOne(context.TODO(), filter, ms.session)
	ms.debug("   deleteInfo: %+v (err %+v)", deleteInfo, err)
	if deleteInfo.DeletedCount == 0 && err == nil {
		err = mongo.ErrNoDocuments
	}

	if hasInterceptor {
		if err == nil && getError == nil {
			ms.invokeInterceptorsAfter("Delete", resourceType, resource)
		} else {
			ms.invokeInterceptorsOnError("Delete", resourceType, err, resource)
		}
	}

	err = convertMongoErr(err)
	return
}

func saveDeletionIntoHistory(resourceType string, id string, curCollection *mongo.Collection, prevCollection *mongo.Collection, session *mongo.Session) (newVersionIdStr string, err error) {
	// get current version of this document
	var currentDoc bson.Document
	currentDocQuery := bson.NewDocument(bson.EC.String("_id", id))
	err = curCollection.FindOne(context.TODO(), currentDocQuery, session).Decode(&currentDoc)

	if err == mongo.ErrNoDocuments {
		return "", err
	} else if err != nil {
		return "", errors.Wrap(convertMongoErr(err), "saveDeletionIntoHistory: error retrieving current version")
	} else {
		// extract current version
		hasVersionId, curVersionId, _ := getVersionIdFromResource(&currentDoc)
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
		setVermongoId(&currentDoc, curVersionId)
		// NOTE: currentDoc._id modified in-place

		// create a deletion record
		now := time.Now()
		deletionRecord := bson.NewDocument(
			bson.EC.SubDocumentFromElements("_id",
				bson.EC.String("_id", id),
				bson.EC.Int32("_version", int32(newVersionId)),
				bson.EC.Int32("_deleted", 1),
			),
			bson.EC.SubDocumentFromElements("meta",
				bson.EC.String("versionId", newVersionIdStr),
				bson.EC.Time("lastUpdated", now),
			),
		)

		_, err = prevCollection.InsertMany(context.TODO(), []interface{}{&currentDoc, deletionRecord}, session) // TODO: do concurrently with the main deletion
		if err != nil && strings.Contains(err.Error(), "duplicate key") {
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

func (ms *mongoSession) ConditionalDelete(query search.Query) (count int64, err error) {

	IDsToDelete, err := ms.FindIDs(query)
	if err != nil {
		return 0, err
	}
	// There is the potential here for the delete to fail if the slice of IDs
	// is too large (exceeding Mongo's 16MB document size limit).
	IDsToDeleteValues := make([]*bson.Value, len(IDsToDelete))
	for i, id := range IDsToDelete {
		IDsToDeleteValues[i] = bson.VC.String(id)
	}
	deleteQuery := bson.NewDocument(bson.EC.SubDocumentFromElements("_id", bson.EC.ArrayFromElements("$in", IDsToDeleteValues...)))
	resourceType := query.Resource
	curCollection := ms.CurrentVersionCollection(resourceType)
	prevCollection := ms.PreviousVersionsCollection(resourceType)

	hasInterceptors := ms.hasInterceptorsForOpAndType("Delete", resourceType)

	if hasInterceptors || ms.dal.enableHistory {
		/* Interceptors for a conditional delete are tricky since an interceptor is only run
		   AFTER the database operation and only on resources that were SUCCESSFULLY deleted. We use
		   the following approach:
		   1. Bulk delete those resources by ID
		   2. Search again using the SAME query, to verify that those resources were in fact deleted
		   3. Run the interceptor(s) on all resources that ARE NOT in the second search (since they were truly deleted)
		*/

		// get the resources that are about to be deleted
		bundle, err := ms.Search(url.URL{}, query) // the baseURL argument here does not matter

		if err == nil {
			for _, elem := range bundle.Entry {
				if hasInterceptors {
					ms.invokeInterceptorsBefore("Delete", resourceType, elem.Resource)
				}
			}

			for _, elem := range bundle.Entry {
				if ms.dal.enableHistory {
					id := elem.Resource.Id()
					_, err = saveDeletionIntoHistory(resourceType, id, curCollection, prevCollection, ms.session)
					if err != nil {
						return count, errors.Wrapf(err, "failed to save deletion into history (%s/%s)", resourceType, id)
					}
				}
			}

			// Do the bulk delete by ID.
			info, err := curCollection.DeleteMany(context.TODO(), deleteQuery, ms.session)
			deletedIds := make([]string, len(IDsToDelete))
			if info != nil {
				count = info.DeletedCount
			}

			if err != nil {
				if hasInterceptors {
					for _, elem := range bundle.Entry {
						ms.invokeInterceptorsOnError("Delete", resourceType, err, elem.Resource)
					}
				}
				return count, convertMongoErr(err)
			} else if hasInterceptors == false {
				return count, nil
			}

			var searchErr error

			if count < int64(len(IDsToDelete)) {
				// Some but not all resources were removed, so use the original search query
				// to see which resources are left.
				var failBundle *models2.ShallowBundle
				failBundle, searchErr = ms.Search(url.URL{}, query)
				deletedIds = setDiff(IDsToDelete, getResourceIdsFromBundle(failBundle))
			} else {
				// All resources were successfully removed
				deletedIds = IDsToDelete
			}

			if searchErr == nil {
				for _, elem := range bundle.Entry {
					id := elem.Resource.Id()

					if elementInSlice(id, deletedIds) {
						// This resource was confirmed deleted
						ms.invokeInterceptorsAfter("Delete", resourceType, elem.Resource)
					} else {
						// This resource was not confirmed deleted, which is an error
						resourceErr := fmt.Errorf("ConditionalDelete: failed to delete resource %s with ID %s", resourceType, id)
						ms.invokeInterceptorsOnError("Delete", resourceType, resourceErr, elem.Resource)
					}
				}
			}
		}
		return count, convertMongoErr(err)
	} else {
		// do the bulk delete the usual way
		info, err := curCollection.DeleteMany(context.TODO(), deleteQuery, ms.session)
		if info != nil {
			count = info.DeletedCount
		}
		return count, convertMongoErr(err)
	}
}

func (ms *mongoSession) History(baseURL url.URL, resourceType string, id string) (bundle *models2.ShallowBundle, err error) {

	// check id
	_, err = convertIDToBsonID(id)
	if err != nil {
		return nil, ErrNotFound
	}

	baseURLstr := baseURL.String()
	if !strings.HasSuffix(baseURLstr, "/") {
		baseURLstr = baseURLstr + "/"
	}
	fullUrl := baseURLstr + id

	curCollection := ms.CurrentVersionCollection(resourceType)
	prevCollection := ms.PreviousVersionsCollection(resourceType)

	var entryList []models2.ShallowBundleEntryComponent
	makeEntryRequest := func(method string) *models.BundleEntryRequestComponent {
		return &models.BundleEntryRequestComponent{
			Url:    resourceType + "/" + id,
			Method: method,
		}
	}

	// add current version
	var curDoc bson.Document
	curDocQuery := bson.NewDocument(bson.EC.String("_id", id))
	err = curCollection.FindOne(context.TODO(), curDocQuery, ms.session).Decode(&curDoc)
	if err == nil {
		var entry models2.ShallowBundleEntryComponent
		entry.FullUrl = fullUrl
		entry.Resource, err = models2.NewResourceFromBSON2(&curDoc)
		if err != nil {
			return nil, errors.Wrap(err, "History: NewResourceFromBSON failed")
		}
		entry.Request = makeEntryRequest("PUT")
		entryList = append(entryList, entry)
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	// sort - oldest versions last
	prevDocsQuery := bson.NewDocument(bson.EC.String("_id._id", id))
	prevDocsSort := findopt.Sort(bson.NewDocument(bson.EC.Int32("_id._version", -1)))
	cursor, err := prevCollection.Find(context.TODO(), prevDocsQuery, prevDocsSort, ms.session)
	if err != nil {
		return nil, errors.Wrap(err, "History: prevCollection.Find failed")
	}

	for cursor.Next(context.TODO()) {

		var prevDocBson bson.Document
		err = cursor.Decode(&prevDocBson)
		ms.debug("History: decoded prev document: %s", prevDocBson.String())
		if err != nil {
			return nil, errors.Wrap(err, "History: cursor.Decode failed")
		}

		var entry models2.ShallowBundleEntryComponent
		entry.FullUrl = fullUrl

		deleted, resource, err := unmarshalPreviousVersion(&prevDocBson)
		if err != nil {
			return nil, errors.Wrap(err, "History: unmarshalPreviousVersion failed")
		}
		if deleted {
			entry.Request = makeEntryRequest("DELETE")
		} else {
			entry.Resource = resource
			entry.Request = makeEntryRequest("PUT")
		}

		entryList = append(entryList, entry)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "History: MongoDB query for previous versions failed")
	}

	totalDocs := uint32(len(entryList))
	if totalDocs == 0 {
		return nil, ErrNotFound
	}

	// last entry should be a POST
	entryList[len(entryList)-1].Request.Method = "POST"
	entryList[len(entryList)-1].Request.Url = resourceType

	// output a Bundle
	bundle = &models2.ShallowBundle{
		Id:    objectid.New().Hex(),
		Type:  "history",
		Entry: entryList,
		Total: &totalDocs,
	}

	// TODO: use paging
	// bundle.Link = dal.generatePagingLinks(baseURL, searchQuery, total, uint32(numResults))

	return bundle, nil
}

func (ms *mongoSession) Search(baseURL url.URL, searchQuery search.Query) (*models2.ShallowBundle, error) {

	searcher := search.NewMongoSearcher(ms.db, ms.session, ms.dal.countTotalResults, ms.dal.enableCISearches, ms.dal.readonly)

	resources, total, err := searcher.Search(searchQuery)
	if err != nil {
		return nil, convertMongoErr(err)
	}

	includesMap := make(map[string]*models2.Resource)
	var entryList []models2.ShallowBundleEntryComponent
	numResults := len(resources)
	baseURLstr := baseURL.String()
	if !strings.HasSuffix(baseURLstr, "/") {
		baseURLstr = baseURLstr + "/"
	}

	for i := 0; i < numResults; i++ {
		var entry models2.ShallowBundleEntryComponent
		entry.Resource = resources[i]
		entry.FullUrl = baseURLstr + resources[i].Id()
		entry.Search = &models.BundleEntrySearchComponent{Mode: "match"}
		entryList = append(entryList, entry)

		if searchQuery.UsesIncludes() || searchQuery.UsesRevIncludes() {

			for _, included := range entry.Resource.SearchIncludes() {
				includesMap[included.ResourceType()+"/"+included.Id()] = included
			}

		}
	}

	for _, v := range includesMap {
		ms.debug("includesMap: %#v\n", v)
		var entry models2.ShallowBundleEntryComponent
		entry.Resource = v
		entry.Search = &models.BundleEntrySearchComponent{Mode: "include"}
		entryList = append(entryList, entry)
	}

	bundle := models2.ShallowBundle{
		Id:    objectid.New().Hex(),
		Type:  "searchset",
		Entry: entryList,
	}

	// Only include the total if counts are enabled, or if _summary=count was applied.
	if ms.dal.countTotalResults || searchQuery.Options().Summary == "count" {
		bundle.Total = &total
	}

	bundle.Link = ms.generatePagingLinks(baseURL, searchQuery, total, uint32(numResults))

	return &bundle, nil
}

func (ms *mongoSession) FindIDs(searchQuery search.Query) (IDs []string, err error) {

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
	searcher := search.NewMongoSearcher(ms.db, ms.session, ms.dal.countTotalResults, ms.dal.enableCISearches, ms.dal.readonly)
	results, _, err := searcher.Search(newQuery)
	if err != nil {
		return nil, convertMongoErr(err)
	}

	IDs = make([]string, len(results))
	for i, result := range results {
		IDs[i] = result.Id()
	}

	return IDs, nil
}

func (ms *mongoSession) generatePagingLinks(baseURL url.URL, query search.Query, total uint32, numResults uint32) []models.BundleLinkComponent {

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
	if ms.dal.countTotalResults {
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

func convertIDToBsonID(id string) (objectid.ObjectID, error) {
	objId, err := objectid.FromHex(id)
	if err == nil {
		return objId, nil
	}
	return objectid.NilObjectID, models.NewOperationOutcome("fatal", "exception", "Id must be a valid BSON ObjectId")
}

func updateResourceMeta(resource *models2.Resource, versionId int) {
	now := time.Now()
	resource.SetLastUpdatedTime(now)
	resource.SetVersionId(versionId)
}

func convertMongoErr(err error) error {
	switch err {
	case mongo.ErrNoDocuments:
		return ErrNotFound
	default:
		_, filename, lineno, _ := runtime.Caller(1)
		return errors.Wrapf(err, "MongoDB operation error (%s:%d)", filename, lineno)
	}
}

// getResourceIdsFromBundle parses a slice of BSON resource IDs from a valid
// bundle of resources (typically returned from a search operation). Order is
// preserved.
func getResourceIdsFromBundle(bundle *models2.ShallowBundle) []string {
	resourceIds := make([]string, int(*bundle.Total))
	for i, elem := range bundle.Entry {
		resourceIds[i] = elem.Resource.Id()
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
