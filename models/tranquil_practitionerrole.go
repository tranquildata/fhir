
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
	case "active":
		return fhirVal.Active, true
	case "period":
		return fhirVal.Period, true
	case "organization":
		return fhirVal.Organization, true
	case "healthcareservice":
		return fhirVal.HealthcareService, true
	case "telecom":
		return fhirVal.Telecom, true
	case "availabletime":
		return fhirVal.AvailableTime, true
	case "endpoint":
		return fhirVal.Endpoint, true
	case "identifier":
		return fhirVal.Identifier, true
	case "availabilityexceptions":
		return fhirVal.AvailabilityExceptions, true
	case "notavailable":
		return fhirVal.NotAvailable, true
	case "code":
		return fhirVal.Code, true
	case "specialty":
		return fhirVal.Specialty, true
	case "location":
		return fhirVal.Location, true
	case "practitioner":
		return fhirVal.Practitioner, true

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
		"period": &FieldTypeSupport{"Period", false, true},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"healthcareservice": &FieldTypeSupport{"Reference", true, false},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"availabletime": &FieldTypeSupport{"PractitionerRoleAvailableTimeComponent", true, false},
		"endpoint": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"availabilityexceptions": &FieldTypeSupport{"string", false, false},
		"notavailable": &FieldTypeSupport{"PractitionerRoleNotAvailableComponent", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", true, false},
		"specialty": &FieldTypeSupport{"CodeableConcept", true, false},
		"location": &FieldTypeSupport{"Reference", true, false},
		"practitioner": &FieldTypeSupport{"Reference", false, true},

	}
}
