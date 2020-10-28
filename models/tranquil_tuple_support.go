/*
 * Copyright (c) 2020, Tranquil Data, Inc. All rights reserved.
 */

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"k8s.io/client-go/util/jsonpath"
	//"github.com/kubernetes/client-go/util/jsonpath"
)

type FieldTypeSupport struct {
	Name          string
	Array         bool
	StructPointer bool
}

type TupleSupport interface {
	FieldByLowerName(nameLower string) (interface{}, bool)
	FieldsToTypes() map[string]*FieldTypeSupport
	TupleName() (id string, resourceType string)
	//Truncate takes a set of field names (lowercase) and returns a new TupleSupport that contains only
	// those fields. The returned value should be used, and the caller should not assume that the
	// original receiver is side-effected. The field to keep string can be either a top-level field name
	// or full jsonPath strings (which start with a '$', '.', or '{')
	Truncate(fieldToKeep string) (TupleSupport, error)
}

//Returns an array of results
func ApplyFieldName(fieldToKeep string, fhirVal interface{}) ([]interface{}, error) {
	if len(fieldToKeep) < 1 {
		return nil, fmt.Errorf("No field specified")
	}
	jsonPathString := ""
	switch {
	case strings.HasPrefix(fieldToKeep, "."),
		strings.HasPrefix(fieldToKeep, "$"),
		strings.HasPrefix(fieldToKeep, "{"):
		jsonPathString = fieldToKeep
	default:
		jsonPathString = "{." + fieldToKeep + "}"
	}
	jsonPth := jsonpath.New("td testing")
	jsonPth.EnableJSONOutput(true)
	if err := jsonPth.Parse(jsonPathString); err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	var bort interface{}
	jsonBytez, err := json.Marshal(fhirVal)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytez, &bort)
	if err != nil {
		return nil, err
	}
	err = jsonPth.Execute(buf, bort)
	if err != nil {
		return nil, err
	}
	rezzes := []interface{}{}
	err = json.Unmarshal(buf.Bytes(), &rezzes)
	if err != nil {
		return nil, err
	}
	return rezzes, nil
}
