
package models


func (fhirVal *QuestionnaireResponse) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *QuestionnaireResponse) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "context":
		return fhirVal.Context, true
	case "authored":
		return fhirVal.Authored, true
	case "author":
		return fhirVal.Author, true
	case "questionnaire":
		return fhirVal.Questionnaire, true
	case "subject":
		return fhirVal.Subject, true
	case "parent":
		return fhirVal.Parent, true
	case "status":
		return fhirVal.Status, true
	case "source":
		return fhirVal.Source, true
	case "item":
		return fhirVal.Item, true
	case "identifier":
		return fhirVal.Identifier, true
	case "basedon":
		return fhirVal.BasedOn, true

	default:
		return nil, false
	}
}

func (fhirVal *QuestionnaireResponse) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"context": &FieldTypeSupport{"Reference", false, true},
		"authored": &FieldTypeSupport{"FHIRDateTime", false, true},
		"author": &FieldTypeSupport{"Reference", false, true},
		"questionnaire": &FieldTypeSupport{"Reference", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"parent": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"source": &FieldTypeSupport{"Reference", false, true},
		"item": &FieldTypeSupport{"QuestionnaireResponseItemComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},

	}
}
