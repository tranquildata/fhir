package models

import (
	"encoding/json"
	"testing"
)

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

	results, err := ApplyFieldName("status", thingie)
	if err != nil {
		t.Error(err)
	}
	if len(results) != 1 || results[0].(string) != "foobar" {
		t.Errorf("Wrong value for status: %v", results)
	}
}

func Test_thatSpecialEOBTruncateLogicWorks(t *testing.T) {
	eob := &ExplanationOfBenefit{
		DomainResource: DomainResource{
			Resource: Resource{
				ResourceType: "blah blah blah",
				Id:           "not a uuid2",
			},
		},
		BenefitBalance: []ExplanationOfBenefitBenefitBalanceComponent{
			ExplanationOfBenefitBenefitBalanceComponent{
				Name: "balance 1",
				Financial: []ExplanationOfBenefitBenefitComponent{
					ExplanationOfBenefitBenefitComponent{
						BackboneElement: BackboneElement{
							Element: Element{
								Id: "fin 1",
							},
						},
					},
				},
			},
			ExplanationOfBenefitBenefitBalanceComponent{
				Name: "balance 2",
				Financial: []ExplanationOfBenefitBenefitComponent{
					ExplanationOfBenefitBenefitComponent{
						BackboneElement: BackboneElement{
							Element: Element{
								Id: "fin 2",
							},
						},
					},
				},
			},
		},
	}

	results, err := ApplyFieldName("{.benefitBalance[1]}", eob)
	if err != nil {
		t.Error(err)
	}
	structPtr := &ExplanationOfBenefit{
		DomainResource: DomainResource{
			Resource: Resource{
				ResourceType: eob.DomainResource.Resource.ResourceType,
				Id:           eob.DomainResource.Resource.Id,
			},
		},
	}
	benefitBal := []ExplanationOfBenefitBenefitBalanceComponent{}
	for _, benefit := range results {
		var asBytes []byte
		asBennies := []ExplanationOfBenefitBenefitBalanceComponent{}
		switch asType := benefit.(type) {
		case map[string]interface{}:
			if asBytes, err = json.Marshal(asType); err != nil {
				t.Error(err)
			} else {
				asBennies = append(asBennies, ExplanationOfBenefitBenefitBalanceComponent{})
				if err = json.Unmarshal(asBytes, &asBennies[0]); err != nil {
					t.Error(err)
				}
			}
		case []interface{}:
			if asBytes, err = json.Marshal(asType); err != nil {
				t.Error(err)
			} else if err = json.Unmarshal(asBytes, &asBennies); err != nil {
				t.Error(err)
			}
		default:
			t.Errorf("invalid type encountered")
		}
		benefitBal = append(benefitBal, asBennies...)
	}
	structPtr.BenefitBalance = benefitBal
}

func Test_thatSpecialEOBHandlingWorks(t *testing.T) {
	eob := &ExplanationOfBenefit{
		DomainResource: DomainResource{
			Resource: Resource{
				ResourceType: "blah blah blah",
				Id:           "not a uuid2",
			},
		},
		BenefitBalance: []ExplanationOfBenefitBenefitBalanceComponent{
			ExplanationOfBenefitBenefitBalanceComponent{
				Name: "balance 1",
				Financial: []ExplanationOfBenefitBenefitComponent{
					ExplanationOfBenefitBenefitComponent{
						BackboneElement: BackboneElement{
							Element: Element{
								Id: "fin 1",
							},
						},
					},
				},
			},
			ExplanationOfBenefitBenefitBalanceComponent{
				Name: "balance 2",
				Financial: []ExplanationOfBenefitBenefitComponent{
					ExplanationOfBenefitBenefitComponent{
						BackboneElement: BackboneElement{
							Element: Element{
								Id: "fin 2",
							},
						},
					},
				},
			},
		},
	}

	nu, err := eob.Truncate("{.benefitBalance[1]}")
	if err != nil {
		t.Error(err)
	}
	if asEob, ok := nu.(*ExplanationOfBenefit); !ok {
		t.Errorf("invalid return type")
	} else if len(asEob.BenefitBalance) != 1 {
		t.Errorf("incorrect number of benefits: %d", len(asEob.BenefitBalance))
	}
}
