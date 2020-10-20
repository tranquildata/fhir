
package models


func (fhirVal *Condition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Condition) FieldByLowerName(nameLower string) (interface{}, bool) {
	switch nameLower {

	case "resourcetype":
		return fhirVal.ResourceType, true
	case "id":
		return fhirVal.Id, true
	case "meta":
		return fhirVal.Meta, true
	case "implicitrules":
		return fhirVal.ImplicitRules, true
	case "language":
		return fhirVal.Language, true

	case "text":
		return fhirVal.Text, true
	case "contained":
		return fhirVal.Contained, true
	case "extension":
		return fhirVal.Extension, true
	case "modifierextension":
		return fhirVal.ModifierExtension, true

	default:
		return nil, false
	}
}

func (fhirVal *Condition) FieldsToTypes() map[string]*FieldTypeSupport {
	return map[string]*FieldTypeSupport {

			"resourcetype": &FieldTypeSupport{"string", false, false},
		"id": &FieldTypeSupport{"string", false, false},
		"meta": &FieldTypeSupport{"Meta", false, true},
		"implicitrules": &FieldTypeSupport{"string", false, false},
		"language": &FieldTypeSupport{"string", false, false},

		"text": &FieldTypeSupport{"Narrative", false, true},
		"Contained": &FieldTypeSupport{"Containedresources", false, false},
		"extension": &FieldTypeSupport{"Extension", true, false},
		"modifierextension": &FieldTypeSupport{"Extension", true, false},						
		"verificationstatus": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"onsetstring": &FieldTypeSupport{"string", false, false},
		"abatementage": &FieldTypeSupport{"Quantity", false, true},
		"abatementboolean": &FieldTypeSupport{"bool", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"severity": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"onsetage": &FieldTypeSupport{"Quantity", false, true},
		"onsetperiod": &FieldTypeSupport{"Period", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"clinicalstatus": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"onsetrange": &FieldTypeSupport{"Range", false, true},
		"abatementdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"abatementperiod": &FieldTypeSupport{"Period", false, true},
		"asserter": &FieldTypeSupport{"Reference", false, true},
		"evidence": &FieldTypeSupport{"ConditionEvidenceComponent", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"onsetdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"abatementrange": &FieldTypeSupport{"Range", false, true},
		"abatementstring": &FieldTypeSupport{"string", false, false},
		"asserteddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"stage": &FieldTypeSupport{"ConditionStageComponent", false, true},

	}
}
