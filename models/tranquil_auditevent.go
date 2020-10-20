
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
	case "outcomedesc":
		return fhirVal.OutcomeDesc, true
	case "purposeofevent":
		return fhirVal.PurposeOfEvent, true
	case "agent":
		return fhirVal.Agent, true
	case "entity":
		return fhirVal.Entity, true
	case "type":
		return fhirVal.Type, true
	case "outcome":
		return fhirVal.Outcome, true
	case "recorded":
		return fhirVal.Recorded, true
	case "source":
		return fhirVal.Source, true
	case "subtype":
		return fhirVal.Subtype, true
	case "action":
		return fhirVal.Action, true

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
		"outcomedesc": &FieldTypeSupport{"string", false, false},
		"purposeofevent": &FieldTypeSupport{"CodeableConcept", true, false},
		"agent": &FieldTypeSupport{"AuditEventAgentComponent", true, false},
		"entity": &FieldTypeSupport{"AuditEventEntityComponent", true, false},
		"type": &FieldTypeSupport{"Coding", false, true},
		"outcome": &FieldTypeSupport{"string", false, false},
		"recorded": &FieldTypeSupport{"FHIRDateTime", false, true},
		"source": &FieldTypeSupport{"AuditEventSourceComponent", false, true},
		"subtype": &FieldTypeSupport{"Coding", true, false},
		"action": &FieldTypeSupport{"string", false, false},

	}
}
