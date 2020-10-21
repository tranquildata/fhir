
package models


func (fhirVal *Subscription) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Subscription) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "channel":
		return fhirVal.Channel, true
	case "tag":
		return fhirVal.Tag, true
	case "status":
		return fhirVal.Status, true
	case "contact":
		return fhirVal.Contact, true
	case "end":
		return fhirVal.End, true
	case "reason":
		return fhirVal.Reason, true
	case "criteria":
		return fhirVal.Criteria, true
	case "error":
		return fhirVal.Error, true

	default:
		return nil, false
	}
}

func (fhirVal *Subscription) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"channel": &FieldTypeSupport{"SubscriptionChannelComponent", false, true},
		"tag": &FieldTypeSupport{"Coding", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactPoint", true, false},
		"end": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reason": &FieldTypeSupport{"string", false, false},
		"criteria": &FieldTypeSupport{"string", false, false},
		"error": &FieldTypeSupport{"string", false, false},

	}
}
