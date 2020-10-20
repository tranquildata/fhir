
package models


func (fhirVal *AdverseEvent) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *AdverseEvent) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "date":
		return fhirVal.Date, true
	case "description":
		return fhirVal.Description, true
	case "category":
		return fhirVal.Category, true
	case "reaction":
		return fhirVal.Reaction, true
	case "location":
		return fhirVal.Location, true
	case "outcome":
		return fhirVal.Outcome, true
	case "recorder":
		return fhirVal.Recorder, true
	case "suspectentity":
		return fhirVal.SuspectEntity, true
	case "subjectmedicalhistory":
		return fhirVal.SubjectMedicalHistory, true
	case "type":
		return fhirVal.Type, true
	case "subject":
		return fhirVal.Subject, true
	case "seriousness":
		return fhirVal.Seriousness, true
	case "study":
		return fhirVal.Study, true
	case "identifier":
		return fhirVal.Identifier, true
	case "eventparticipant":
		return fhirVal.EventParticipant, true
	case "referencedocument":
		return fhirVal.ReferenceDocument, true

	default:
		return nil, false
	}
}

func (fhirVal *AdverseEvent) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"string", false, false},
		"reaction": &FieldTypeSupport{"Reference", true, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"recorder": &FieldTypeSupport{"Reference", false, true},
		"suspectentity": &FieldTypeSupport{"AdverseEventSuspectEntityComponent", true, false},
		"subjectmedicalhistory": &FieldTypeSupport{"Reference", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"seriousness": &FieldTypeSupport{"CodeableConcept", false, true},
		"study": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"eventparticipant": &FieldTypeSupport{"Reference", false, true},
		"referencedocument": &FieldTypeSupport{"Reference", true, false},

	}
}
