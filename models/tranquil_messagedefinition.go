
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
	case "name":
		return fhirVal.Name, true
	case "title":
		return fhirVal.Title, true
	case "status":
		return fhirVal.Status, true
	case "experimental":
		return fhirVal.Experimental, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "parent":
		return fhirVal.Parent, true
	case "identifier":
		return fhirVal.Identifier, true
	case "responserequired":
		return fhirVal.ResponseRequired, true
	case "focus":
		return fhirVal.Focus, true
	case "version":
		return fhirVal.Version, true
	case "description":
		return fhirVal.Description, true
	case "purpose":
		return fhirVal.Purpose, true
	case "replaces":
		return fhirVal.Replaces, true
	case "event":
		return fhirVal.Event, true
	case "category":
		return fhirVal.Category, true
	case "url":
		return fhirVal.Url, true
	case "publisher":
		return fhirVal.Publisher, true
	case "contact":
		return fhirVal.Contact, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "copyright":
		return fhirVal.Copyright, true
	case "base":
		return fhirVal.Base, true
	case "allowedresponse":
		return fhirVal.AllowedResponse, true
	case "date":
		return fhirVal.Date, true

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
		"name": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"parent": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"responserequired": &FieldTypeSupport{"bool", false, true},
		"focus": &FieldTypeSupport{"MessageDefinitionFocusComponent", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"event": &FieldTypeSupport{"Coding", false, true},
		"category": &FieldTypeSupport{"string", false, false},
		"url": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"base": &FieldTypeSupport{"Reference", false, true},
		"allowedresponse": &FieldTypeSupport{"MessageDefinitionAllowedResponseComponent", true, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
