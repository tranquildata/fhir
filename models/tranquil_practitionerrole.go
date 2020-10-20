
package models


func (fhirVal *PractitionerRole) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *PractitionerRole) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *PractitionerRole) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"active": &FieldTypeSupport{"bool", false, true},
		"notavailable": &FieldTypeSupport{"PractitionerRoleNotAvailableComponent", true, false},
		"availabilityexceptions": &FieldTypeSupport{"string", false, false},
		"period": &FieldTypeSupport{"Period", false, true},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", true, false},
		"availabletime": &FieldTypeSupport{"PractitionerRoleAvailableTimeComponent", true, false},
		"location": &FieldTypeSupport{"Reference", true, false},
		"healthcareservice": &FieldTypeSupport{"Reference", true, false},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"practitioner": &FieldTypeSupport{"Reference", false, true},
		"specialty": &FieldTypeSupport{"CodeableConcept", true, false},
		"endpoint": &FieldTypeSupport{"Reference", true, false},

	}
}
