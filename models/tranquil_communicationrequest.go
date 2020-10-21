
package models


func (fhirVal *CommunicationRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *CommunicationRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "priority":
		return fhirVal.Priority, true
	case "medium":
		return fhirVal.Medium, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "category":
		return fhirVal.Category, true
	case "context":
		return fhirVal.Context, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "note":
		return fhirVal.Note, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "recipient":
		return fhirVal.Recipient, true
	case "replaces":
		return fhirVal.Replaces, true
	case "status":
		return fhirVal.Status, true
	case "subject":
		return fhirVal.Subject, true
	case "topic":
		return fhirVal.Topic, true
	case "payload":
		return fhirVal.Payload, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "identifier":
		return fhirVal.Identifier, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "requester":
		return fhirVal.Requester, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "sender":
		return fhirVal.Sender, true

	default:
		return nil, false
	}
}

func (fhirVal *CommunicationRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"priority": &FieldTypeSupport{"string", false, false},
		"medium": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"recipient": &FieldTypeSupport{"Reference", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"topic": &FieldTypeSupport{"Reference", true, false},
		"payload": &FieldTypeSupport{"CommunicationRequestPayloadComponent", true, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"requester": &FieldTypeSupport{"CommunicationRequestRequesterComponent", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"sender": &FieldTypeSupport{"Reference", false, true},

	}
}
