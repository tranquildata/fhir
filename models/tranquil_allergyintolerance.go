
package models


func (fhirVal *AllergyIntolerance) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *AllergyIntolerance) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *AllergyIntolerance) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"criticality": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"recorder": &FieldTypeSupport{"Reference", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"onsetrange": &FieldTypeSupport{"Range", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"asserter": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"verificationstatus": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"string", false, false},
		"onsetdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"onsetage": &FieldTypeSupport{"Quantity", false, true},
		"onsetstring": &FieldTypeSupport{"string", false, false},
		"clinicalstatus": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"string", true, false},
		"onsetperiod": &FieldTypeSupport{"Period", false, true},
		"asserteddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"lastoccurrence": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reaction": &FieldTypeSupport{"AllergyIntoleranceReactionComponent", true, false},

	}
}
