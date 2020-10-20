
package models


func (fhirVal *EpisodeOfCare) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *EpisodeOfCare) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "patient":
		return fhirVal.Patient, true
	case "period":
		return fhirVal.Period, true
	case "caremanager":
		return fhirVal.CareManager, true
	case "team":
		return fhirVal.Team, true
	case "identifier":
		return fhirVal.Identifier, true
	case "statushistory":
		return fhirVal.StatusHistory, true
	case "diagnosis":
		return fhirVal.Diagnosis, true
	case "managingorganization":
		return fhirVal.ManagingOrganization, true
	case "referralrequest":
		return fhirVal.ReferralRequest, true
	case "account":
		return fhirVal.Account, true
	case "status":
		return fhirVal.Status, true
	case "type":
		return fhirVal.Type, true

	default:
		return nil, false
	}
}

func (fhirVal *EpisodeOfCare) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"patient": &FieldTypeSupport{"Reference", false, true},
		"period": &FieldTypeSupport{"Period", false, true},
		"caremanager": &FieldTypeSupport{"Reference", false, true},
		"team": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"statushistory": &FieldTypeSupport{"EpisodeOfCareStatusHistoryComponent", true, false},
		"diagnosis": &FieldTypeSupport{"EpisodeOfCareDiagnosisComponent", true, false},
		"managingorganization": &FieldTypeSupport{"Reference", false, true},
		"referralrequest": &FieldTypeSupport{"Reference", true, false},
		"account": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
