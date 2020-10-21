
package models


func (fhirVal *Composition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Composition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "class":
		return fhirVal.Class, true
	case "subject":
		return fhirVal.Subject, true
	case "date":
		return fhirVal.Date, true
	case "author":
		return fhirVal.Author, true
	case "title":
		return fhirVal.Title, true
	case "confidentiality":
		return fhirVal.Confidentiality, true
	case "custodian":
		return fhirVal.Custodian, true
	case "event":
		return fhirVal.Event, true
	case "encounter":
		return fhirVal.Encounter, true
	case "attester":
		return fhirVal.Attester, true
	case "section":
		return fhirVal.Section, true
	case "status":
		return fhirVal.Status, true
	case "type":
		return fhirVal.Type, true
	case "relatesto":
		return fhirVal.RelatesTo, true

	default:
		return nil, false
	}
}

func (fhirVal *Composition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"class": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"author": &FieldTypeSupport{"Reference", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"confidentiality": &FieldTypeSupport{"string", false, false},
		"custodian": &FieldTypeSupport{"Reference", false, true},
		"event": &FieldTypeSupport{"CompositionEventComponent", true, false},
		"encounter": &FieldTypeSupport{"Reference", false, true},
		"attester": &FieldTypeSupport{"CompositionAttesterComponent", true, false},
		"section": &FieldTypeSupport{"CompositionSectionComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"relatesto": &FieldTypeSupport{"CompositionRelatesToComponent", true, false},

	}
}
