/*
 * Copyright (c) 2020, Tranquil Data, Inc. All rights reserved.
 */

package utils

import (
	"io/ioutil"
	"testing"
)

func TestThatStructsCanBeParsed(t *testing.T) {
	st, err := ParseFhirModel("../models/activitydefinition.go")
	if err != nil {
		t.Error(err)
	}
	if st == nil {
		t.Errorf("Struct was nil")
	}
	emitted, err := EmitTupleCode(st, "models")
	if err != nil {
		t.Error(err)
	}
	if len(emitted) == 0 {
		t.Errorf("emitted was empty")
	}
	err = ioutil.WriteFile("../models/tranquil_activitydefinition.go", []byte(emitted), 0644)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

/*
func TestThatProcessDirWorks(t *testing.T) {
	processed, err := ProcessDirectory("../models", "models")
	if err != nil {
		t.Error(err)
	}
	if len(processed) < 100 {
		t.Errorf("Incorrect number of processed files: %d", len(processed))
	}
	fmt.Printf("processed: %v\n", processed)
}
*/
