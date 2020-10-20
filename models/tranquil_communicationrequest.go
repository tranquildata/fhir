
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
	case "identifier":
		return fhirVal.Identifier, true
	case "priority":
		return fhirVal.Priority, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "status":
		return fhirVal.Status, true
	case "recipient":
		return fhirVal.Recipient, true
	case "payload":
		return fhirVal.Payload, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "sender":
		return fhirVal.Sender, true
	case "context":
		return fhirVal.Context, true
	case "requester":
		return fhirVal.Requester, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "category":
		return fhirVal.Category, true
	case "medium":
		return fhirVal.Medium, true
	case "subject":
		return fhirVal.Subject, true
	case "topic":
		return fhirVal.Topic, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "replaces":
		return fhirVal.Replaces, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true

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
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"priority": &FieldTypeSupport{"string", false, false},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"recipient": &FieldTypeSupport{"Reference", true, false},
		"payload": &FieldTypeSupport{"CommunicationRequestPayloadComponent", true, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"sender": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"requester": &FieldTypeSupport{"CommunicationRequestRequesterComponent", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"medium": &FieldTypeSupport{"CodeableConcept", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"topic": &FieldTypeSupport{"Reference", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}
