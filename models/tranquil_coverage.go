
package models


func (fhirVal *Coverage) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Coverage) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *Coverage) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"status": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"subscriberid": &FieldTypeSupport{"string", false, false},
		"sequence": &FieldTypeSupport{"string", false, false},
		"contract": &FieldTypeSupport{"Reference", true, false},
		"policyholder": &FieldTypeSupport{"Reference", false, true},
		"grouping": &FieldTypeSupport{"CoverageGroupComponent", false, true},
		"order": &FieldTypeSupport{"uint32", false, true},
		"subscriber": &FieldTypeSupport{"Reference", false, true},
		"beneficiary": &FieldTypeSupport{"Reference", false, true},
		"dependent": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"relationship": &FieldTypeSupport{"CodeableConcept", false, true},
		"period": &FieldTypeSupport{"Period", false, true},
		"payor": &FieldTypeSupport{"Reference", true, false},
		"network": &FieldTypeSupport{"string", false, false},

	}
}
