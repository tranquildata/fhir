
package models


func (fhirVal *MessageDefinition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *MessageDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "focus":
		return fhirVal.Focus, true
	case "responserequired":
		return fhirVal.ResponseRequired, true
	case "allowedresponse":
		return fhirVal.AllowedResponse, true
	case "url":
		return fhirVal.Url, true
	case "date":
		return fhirVal.Date, true
	case "publisher":
		return fhirVal.Publisher, true
	case "contact":
		return fhirVal.Contact, true
	case "category":
		return fhirVal.Category, true
	case "experimental":
		return fhirVal.Experimental, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "copyright":
		return fhirVal.Copyright, true
	case "parent":
		return fhirVal.Parent, true
	case "identifier":
		return fhirVal.Identifier, true
	case "purpose":
		return fhirVal.Purpose, true
	case "base":
		return fhirVal.Base, true
	case "replaces":
		return fhirVal.Replaces, true
	case "event":
		return fhirVal.Event, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "version":
		return fhirVal.Version, true
	case "name":
		return fhirVal.Name, true
	case "title":
		return fhirVal.Title, true
	case "status":
		return fhirVal.Status, true
	case "description":
		return fhirVal.Description, true

	default:
		return nil, false
	}
}

func (fhirVal *MessageDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"focus": &FieldTypeSupport{"MessageDefinitionFocusComponent", true, false},
		"responserequired": &FieldTypeSupport{"bool", false, true},
		"allowedresponse": &FieldTypeSupport{"MessageDefinitionAllowedResponseComponent", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"category": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"parent": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"purpose": &FieldTypeSupport{"string", false, false},
		"base": &FieldTypeSupport{"Reference", false, true},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"event": &FieldTypeSupport{"Coding", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},

	}
}
