
package models


func (fhirVal *Provenance) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Provenance) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "activity":
		return fhirVal.Activity, true
	case "signature":
		return fhirVal.Signature, true
	case "target":
		return fhirVal.Target, true
	case "location":
		return fhirVal.Location, true
	case "policy":
		return fhirVal.Policy, true
	case "reason":
		return fhirVal.Reason, true
	case "agent":
		return fhirVal.Agent, true
	case "entity":
		return fhirVal.Entity, true
	case "period":
		return fhirVal.Period, true
	case "recorded":
		return fhirVal.Recorded, true

	default:
		return nil, false
	}
}

func (fhirVal *Provenance) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"activity": &FieldTypeSupport{"Coding", false, true},
		"signature": &FieldTypeSupport{"Signature", true, false},
		"target": &FieldTypeSupport{"Reference", true, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"policy": &FieldTypeSupport{"string", true, false},
		"reason": &FieldTypeSupport{"Coding", true, false},
		"agent": &FieldTypeSupport{"ProvenanceAgentComponent", true, false},
		"entity": &FieldTypeSupport{"ProvenanceEntityComponent", true, false},
		"period": &FieldTypeSupport{"Period", false, true},
		"recorded": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
