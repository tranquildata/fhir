package models2


import (
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestConversion(t *testing.T) {

	os.Setenv("GOFHIR_ENCRYPTION_KEY_ID", "testKey1")
	os.Setenv("GOFHIR_ENCRYPTION_KEY_BASE64", "9PM6OC+IlpvbIaS7VSxk1Q4kwe3RH4p5XertTYej46k=")

	for _, encrypt := range []bool { false, true } {

		encryptionEnable  := WhatToEncrypt { PatientDetails: encrypt }

		t.Run(fmt.Sprintf("%+v", encrypt), func(t *testing.T) {
			jsonBytes, err := ioutil.ReadFile("../fixtures/json-edge-cases.json")
			assert.Nil(t, err)

			transformReferencesMap := map[string]string{}
			bson, err := ConvertJsonToGoFhirBSON(jsonBytes, encryptionEnable, transformReferencesMap)
			assert.Nil(t, err)

			if encrypt {
				for _, field := range bson {
					if shouldEncryptField(field.Name) {
						assert.Failf(t, "field should have been encrypted", "field: %s", field.Name)
					}
				}
			}

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
		})
	}
}

func printBSON(bsonDoc *bson.D) {
	bsonBytes, err := bson.Marshal(bsonDoc)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("/tmp/tst2.bson", bsonBytes, 0777)
}