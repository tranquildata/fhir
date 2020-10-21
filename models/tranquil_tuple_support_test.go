package models

import "testing"

func TestTupleMethods(t *testing.T) {
	thingie := &ActivityDefinition{
		DomainResource: DomainResource{
			Resource: Resource{
				ResourceType: "blah",
				Id:           "not a uuid",
			},
		},
		Status: "foobar",
	}
	a, b := thingie.FieldByLowerName("foo")
	if a != nil {
		t.Errorf("found foo")
	}
	if b {
		t.Errorf("returned true")
	}
	a, b = thingie.FieldByLowerName("resourcetype")
	if !b {
		t.Errorf("returned false")
	}
	if a != "blah" {
		t.Errorf("Wrong resource type: %s", a)
	}
	a, b = thingie.FieldByLowerName("status")
	if a == nil || !b {
		t.Errorf("Returned false")
	}
	if a != "foobar" {
		t.Errorf("Incorrect status value: %v", a)
	}

	typeMap := thingie.FieldsToTypes()
	if typeMap == nil || len(typeMap) == 0 {
		t.Errorf("empty type map")
	}
	if len(typeMap) < 38 {
		t.Errorf("Too few fields: %d", len(typeMap))
	}
	ty, ok := typeMap["Url"]
	if ok {
		t.Errorf("Found field named Url")
	}
	if ty != nil {
		t.Errorf("NonNil ty: %v", ty)
	}
	ty, ok = typeMap["url"]
	if !ok {
		t.Errorf("No url field found")
	} else if ty.Array || ty.StructPointer || ty.Name != "string" {
		t.Errorf("Incorrect type")
	}
	ty, ok = typeMap["identifier"]
	if !ok {
		t.Errorf("No identifier field found")
	} else if !ty.Array || ty.StructPointer || ty.Name != "Identifier" {
		t.Errorf("Incorrect type: %v", ty)
	}
	ty, ok = typeMap["experimental"]
	if !ok {
		t.Errorf("No experimental field found")
	} else if ty.Array || !ty.StructPointer || ty.Name != "bool" {
		t.Errorf("Incorrect type: %v", ty)
	}
	ty, ok = typeMap["effectiveperiod"]
	if !ok {
		t.Errorf("No EffectivePeriod field found")
	} else if ty.Array || !ty.StructPointer || ty.Name != "Period" {
		t.Errorf("Incorrect type: %v", ty)
	}
	ty, ok = typeMap["resourcetype"]
	if !ok {
		t.Errorf("No resourcetype field found")
	} else if ty.Array || ty.StructPointer || ty.Name != "string" {
		t.Errorf("Incorrect type: %v", ty)
	}

	id, rTy := thingie.TupleName()
	if id != "not a uuid" {
		t.Errorf("Invalid id: %s", id)
	}
	if rTy != "blah" {
		t.Errorf("Invalid resource type: %s", rTy)
	}
}
