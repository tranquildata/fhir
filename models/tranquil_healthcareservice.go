
package models


func (fhirVal *HealthcareService) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *HealthcareService) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *HealthcareService) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"location": &FieldTypeSupport{"Reference", true, false},
		"comment": &FieldTypeSupport{"string", false, false},
		"extradetails": &FieldTypeSupport{"string", false, false},
		"eligibility": &FieldTypeSupport{"CodeableConcept", false, true},
		"appointmentrequired": &FieldTypeSupport{"bool", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"referralmethod": &FieldTypeSupport{"CodeableConcept", true, false},
		"notavailable": &FieldTypeSupport{"HealthcareServiceNotAvailableComponent", true, false},
		"endpoint": &FieldTypeSupport{"Reference", true, false},
		"photo": &FieldTypeSupport{"Attachment", false, true},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", true, false},
		"specialty": &FieldTypeSupport{"CodeableConcept", true, false},
		"serviceprovisioncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"eligibilitynote": &FieldTypeSupport{"string", false, false},
		"characteristic": &FieldTypeSupport{"CodeableConcept", true, false},
		"availabilityexceptions": &FieldTypeSupport{"string", false, false},
		"active": &FieldTypeSupport{"bool", false, true},
		"providedby": &FieldTypeSupport{"Reference", false, true},
		"programname": &FieldTypeSupport{"string", true, false},
		"availabletime": &FieldTypeSupport{"HealthcareServiceAvailableTimeComponent", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"coveragearea": &FieldTypeSupport{"Reference", true, false},

	}
}
