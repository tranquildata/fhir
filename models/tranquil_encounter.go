
package models


func (fhirVal *Encounter) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Encounter) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "episodeofcare":
		return fhirVal.EpisodeOfCare, true
	case "appointment":
		return fhirVal.Appointment, true
	case "length":
		return fhirVal.Length, true
	case "diagnosis":
		return fhirVal.Diagnosis, true
	case "partof":
		return fhirVal.PartOf, true
	case "classhistory":
		return fhirVal.ClassHistory, true
	case "type":
		return fhirVal.Type, true
	case "incomingreferral":
		return fhirVal.IncomingReferral, true
	case "location":
		return fhirVal.Location, true
	case "class":
		return fhirVal.Class, true
	case "priority":
		return fhirVal.Priority, true
	case "period":
		return fhirVal.Period, true
	case "reason":
		return fhirVal.Reason, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "participant":
		return fhirVal.Participant, true
	case "account":
		return fhirVal.Account, true
	case "hospitalization":
		return fhirVal.Hospitalization, true
	case "serviceprovider":
		return fhirVal.ServiceProvider, true
	case "statushistory":
		return fhirVal.StatusHistory, true
	case "subject":
		return fhirVal.Subject, true

	default:
		return nil, false
	}
}

func (fhirVal *Encounter) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"episodeofcare": &FieldTypeSupport{"Reference", true, false},
		"appointment": &FieldTypeSupport{"Reference", false, true},
		"length": &FieldTypeSupport{"Quantity", false, true},
		"diagnosis": &FieldTypeSupport{"EncounterDiagnosisComponent", true, false},
		"partof": &FieldTypeSupport{"Reference", false, true},
		"classhistory": &FieldTypeSupport{"EncounterClassHistoryComponent", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", true, false},
		"incomingreferral": &FieldTypeSupport{"Reference", true, false},
		"location": &FieldTypeSupport{"EncounterLocationComponent", true, false},
		"class": &FieldTypeSupport{"Coding", false, true},
		"priority": &FieldTypeSupport{"CodeableConcept", false, true},
		"period": &FieldTypeSupport{"Period", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"participant": &FieldTypeSupport{"EncounterParticipantComponent", true, false},
		"account": &FieldTypeSupport{"Reference", true, false},
		"hospitalization": &FieldTypeSupport{"EncounterHospitalizationComponent", false, true},
		"serviceprovider": &FieldTypeSupport{"Reference", false, true},
		"statushistory": &FieldTypeSupport{"EncounterStatusHistoryComponent", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},

	}
}
