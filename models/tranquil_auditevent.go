
package models


func (fhirVal *AuditEvent) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *AuditEvent) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "subtype":
		return fhirVal.Subtype, true
	case "action":
		return fhirVal.Action, true
	case "recorded":
		return fhirVal.Recorded, true
	case "outcome":
		return fhirVal.Outcome, true
	case "purposeofevent":
		return fhirVal.PurposeOfEvent, true
	case "source":
		return fhirVal.Source, true
	case "type":
		return fhirVal.Type, true
	case "agent":
		return fhirVal.Agent, true
	case "entity":
		return fhirVal.Entity, true
	case "outcomedesc":
		return fhirVal.OutcomeDesc, true

	default:
		return nil, false
	}
}

func (fhirVal *AuditEvent) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"subtype": &FieldTypeSupport{"Coding", true, false},
		"action": &FieldTypeSupport{"string", false, false},
		"recorded": &FieldTypeSupport{"FHIRDateTime", false, true},
		"outcome": &FieldTypeSupport{"string", false, false},
		"purposeofevent": &FieldTypeSupport{"CodeableConcept", true, false},
		"source": &FieldTypeSupport{"AuditEventSourceComponent", false, true},
		"type": &FieldTypeSupport{"Coding", false, true},
		"agent": &FieldTypeSupport{"AuditEventAgentComponent", true, false},
		"entity": &FieldTypeSupport{"AuditEventEntityComponent", true, false},
		"outcomedesc": &FieldTypeSupport{"string", false, false},

	}
}
