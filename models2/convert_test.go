package models2


import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestConversion(t *testing.T) {

	jsonBytes, err := ioutil.ReadFile("../fixtures/json-edge-cases.json")
	assert.Nil(t, err)

	transformReferencesMap := map[string]string{}
	bson, err := ConvertJsonToGoFhirBSON(jsonBytes, transformReferencesMap)
	assert.Nil(t, err)

	backToJson, included, err := ConvertGoFhirBSONToJSON(bson)
	assert.Nil(t, err)
	assert.Nil(t, included)

	// fmt.Printf("%#v\n", bson)
	// fmt.Printf("Back to JSON: %s\n", string(backToJson))

	var indentedOrig bytes.Buffer
	var indentedBack bytes.Buffer
	err = json.Indent(&indentedOrig, jsonBytes, "", "  "); assert.Nil(t, err)
	err = json.Indent(&indentedBack, backToJson, "", "  "); assert.Nil(t, err)

	ioutil.WriteFile("/tmp/json_orig", jsonBytes, 0777)
	ioutil.WriteFile("/tmp/json_back", backToJson, 0777)

	// assert.Equal(t, indentedOrig.String(), indentedBack.String(), "json roundtrip")
	assert.JSONEq(t, string(jsonBytes), string(backToJson), "should get back original json")

	printBSON(&bson)
}

func printBSON(bsonDoc *bson.D) {
	bsonBytes, err := bson.Marshal(bsonDoc)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("/tmp/tst2.bson", bsonBytes, 0777)
}