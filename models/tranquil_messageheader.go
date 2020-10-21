
package models


func (fhirVal *MessageHeader) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *MessageHeader) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "author":
		return fhirVal.Author, true
	case "source":
		return fhirVal.Source, true
	case "responsible":
		return fhirVal.Responsible, true
	case "reason":
		return fhirVal.Reason, true
	case "destination":
		return fhirVal.Destination, true
	case "receiver":
		return fhirVal.Receiver, true
	case "sender":
		return fhirVal.Sender, true
	case "timestamp":
		return fhirVal.Timestamp, true
	case "enterer":
		return fhirVal.Enterer, true
	case "response":
		return fhirVal.Response, true
	case "focus":
		return fhirVal.Focus, true
	case "event":
		return fhirVal.Event, true

	default:
		return nil, false
	}
}

func (fhirVal *MessageHeader) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"author": &FieldTypeSupport{"Reference", false, true},
		"source": &FieldTypeSupport{"MessageHeaderMessageSourceComponent", false, true},
		"responsible": &FieldTypeSupport{"Reference", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", false, true},
		"destination": &FieldTypeSupport{"MessageHeaderMessageDestinationComponent", true, false},
		"receiver": &FieldTypeSupport{"Reference", false, true},
		"sender": &FieldTypeSupport{"Reference", false, true},
		"timestamp": &FieldTypeSupport{"FHIRDateTime", false, true},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"response": &FieldTypeSupport{"MessageHeaderResponseComponent", false, true},
		"focus": &FieldTypeSupport{"Reference", true, false},
		"event": &FieldTypeSupport{"Coding", false, true},

	}
}
