
package models


func (fhirVal *Communication) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Communication) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "subject":
		return fhirVal.Subject, true
	case "context":
		return fhirVal.Context, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "identifier":
		return fhirVal.Identifier, true
	case "definition":
		return fhirVal.Definition, true
	case "notdonereason":
		return fhirVal.NotDoneReason, true
	case "recipient":
		return fhirVal.Recipient, true
	case "topic":
		return fhirVal.Topic, true
	case "payload":
		return fhirVal.Payload, true
	case "note":
		return fhirVal.Note, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "status":
		return fhirVal.Status, true
	case "medium":
		return fhirVal.Medium, true
	case "received":
		return fhirVal.Received, true
	case "sender":
		return fhirVal.Sender, true
	case "sent":
		return fhirVal.Sent, true
	case "partof":
		return fhirVal.PartOf, true
	case "notdone":
		return fhirVal.NotDone, true
	case "category":
		return fhirVal.Category, true

	default:
		return nil, false
	}
}

func (fhirVal *Communication) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"recipient": &FieldTypeSupport{"Reference", true, false},
		"topic": &FieldTypeSupport{"Reference", true, false},
		"payload": &FieldTypeSupport{"CommunicationPayloadComponent", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"medium": &FieldTypeSupport{"CodeableConcept", true, false},
		"received": &FieldTypeSupport{"FHIRDateTime", false, true},
		"sender": &FieldTypeSupport{"Reference", false, true},
		"sent": &FieldTypeSupport{"FHIRDateTime", false, true},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
