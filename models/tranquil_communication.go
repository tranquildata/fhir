
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
	case "medium":
		return fhirVal.Medium, true
	case "subject":
		return fhirVal.Subject, true
	case "status":
		return fhirVal.Status, true
	case "notdone":
		return fhirVal.NotDone, true
	case "notdonereason":
		return fhirVal.NotDoneReason, true
	case "received":
		return fhirVal.Received, true
	case "payload":
		return fhirVal.Payload, true
	case "definition":
		return fhirVal.Definition, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "recipient":
		return fhirVal.Recipient, true
	case "context":
		return fhirVal.Context, true
	case "sender":
		return fhirVal.Sender, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true
	case "identifier":
		return fhirVal.Identifier, true
	case "category":
		return fhirVal.Category, true
	case "sent":
		return fhirVal.Sent, true
	case "partof":
		return fhirVal.PartOf, true
	case "topic":
		return fhirVal.Topic, true

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
		"medium": &FieldTypeSupport{"CodeableConcept", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"received": &FieldTypeSupport{"FHIRDateTime", false, true},
		"payload": &FieldTypeSupport{"CommunicationPayloadComponent", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"recipient": &FieldTypeSupport{"Reference", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"sender": &FieldTypeSupport{"Reference", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"sent": &FieldTypeSupport{"FHIRDateTime", false, true},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"topic": &FieldTypeSupport{"Reference", true, false},

	}
}
