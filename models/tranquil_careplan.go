
package models


func (fhirVal *CarePlan) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *CarePlan) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "addresses":
		return fhirVal.Addresses, true
	case "goal":
		return fhirVal.Goal, true
	case "activity":
		return fhirVal.Activity, true
	case "description":
		return fhirVal.Description, true
	case "context":
		return fhirVal.Context, true
	case "period":
		return fhirVal.Period, true
	case "author":
		return fhirVal.Author, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "partof":
		return fhirVal.PartOf, true
	case "category":
		return fhirVal.Category, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "careteam":
		return fhirVal.CareTeam, true
	case "definition":
		return fhirVal.Definition, true
	case "status":
		return fhirVal.Status, true
	case "intent":
		return fhirVal.Intent, true
	case "subject":
		return fhirVal.Subject, true
	case "identifier":
		return fhirVal.Identifier, true
	case "replaces":
		return fhirVal.Replaces, true
	case "title":
		return fhirVal.Title, true
	case "note":
		return fhirVal.Note, true

	default:
		return nil, false
	}
}

func (fhirVal *CarePlan) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"addresses": &FieldTypeSupport{"Reference", true, false},
		"goal": &FieldTypeSupport{"Reference", true, false},
		"activity": &FieldTypeSupport{"CarePlanActivityComponent", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"period": &FieldTypeSupport{"Period", false, true},
		"author": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"careteam": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"intent": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}
