
package models


func (fhirVal *RequestGroup) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *RequestGroup) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "reasoncodeableconcept":
		return fhirVal.ReasonCodeableConcept, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "status":
		return fhirVal.Status, true
	case "priority":
		return fhirVal.Priority, true
	case "author":
		return fhirVal.Author, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "action":
		return fhirVal.Action, true
	case "intent":
		return fhirVal.Intent, true
	case "subject":
		return fhirVal.Subject, true
	case "identifier":
		return fhirVal.Identifier, true
	case "definition":
		return fhirVal.Definition, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "replaces":
		return fhirVal.Replaces, true
	case "context":
		return fhirVal.Context, true
	case "note":
		return fhirVal.Note, true

	default:
		return nil, false
	}
}

func (fhirVal *RequestGroup) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"priority": &FieldTypeSupport{"string", false, false},
		"author": &FieldTypeSupport{"Reference", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", false, true},
		"action": &FieldTypeSupport{"RequestGroupActionComponent", true, false},
		"intent": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}
