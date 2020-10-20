
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
	case "appointmentrequired":
		return fhirVal.AppointmentRequired, true
	case "identifier":
		return fhirVal.Identifier, true
	case "providedby":
		return fhirVal.ProvidedBy, true
	case "location":
		return fhirVal.Location, true
	case "eligibilitynote":
		return fhirVal.EligibilityNote, true
	case "referralmethod":
		return fhirVal.ReferralMethod, true
	case "comment":
		return fhirVal.Comment, true
	case "coveragearea":
		return fhirVal.CoverageArea, true
	case "characteristic":
		return fhirVal.Characteristic, true
	case "endpoint":
		return fhirVal.Endpoint, true
	case "availabilityexceptions":
		return fhirVal.AvailabilityExceptions, true
	case "category":
		return fhirVal.Category, true
	case "specialty":
		return fhirVal.Specialty, true
	case "name":
		return fhirVal.Name, true
	case "photo":
		return fhirVal.Photo, true
	case "eligibility":
		return fhirVal.Eligibility, true
	case "programname":
		return fhirVal.ProgramName, true
	case "availabletime":
		return fhirVal.AvailableTime, true
	case "notavailable":
		return fhirVal.NotAvailable, true
	case "active":
		return fhirVal.Active, true
	case "type":
		return fhirVal.Type, true
	case "extradetails":
		return fhirVal.ExtraDetails, true
	case "telecom":
		return fhirVal.Telecom, true
	case "serviceprovisioncode":
		return fhirVal.ServiceProvisionCode, true

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
		"appointmentrequired": &FieldTypeSupport{"bool", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"providedby": &FieldTypeSupport{"Reference", false, true},
		"location": &FieldTypeSupport{"Reference", true, false},
		"eligibilitynote": &FieldTypeSupport{"string", false, false},
		"referralmethod": &FieldTypeSupport{"CodeableConcept", true, false},
		"comment": &FieldTypeSupport{"string", false, false},
		"coveragearea": &FieldTypeSupport{"Reference", true, false},
		"characteristic": &FieldTypeSupport{"CodeableConcept", true, false},
		"endpoint": &FieldTypeSupport{"Reference", true, false},
		"availabilityexceptions": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"specialty": &FieldTypeSupport{"CodeableConcept", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"photo": &FieldTypeSupport{"Attachment", false, true},
		"eligibility": &FieldTypeSupport{"CodeableConcept", false, true},
		"programname": &FieldTypeSupport{"string", true, false},
		"availabletime": &FieldTypeSupport{"HealthcareServiceAvailableTimeComponent", true, false},
		"notavailable": &FieldTypeSupport{"HealthcareServiceNotAvailableComponent", true, false},
		"active": &FieldTypeSupport{"bool", false, true},
		"type": &FieldTypeSupport{"CodeableConcept", true, false},
		"extradetails": &FieldTypeSupport{"string", false, false},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"serviceprovisioncode": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
