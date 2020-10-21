
package models


func (fhirVal *ValueSet) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ValueSet) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "purpose":
		return fhirVal.Purpose, true
	case "extensible":
		return fhirVal.Extensible, true
	case "expansion":
		return fhirVal.Expansion, true
	case "url":
		return fhirVal.Url, true
	case "experimental":
		return fhirVal.Experimental, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "contact":
		return fhirVal.Contact, true
	case "description":
		return fhirVal.Description, true
	case "identifier":
		return fhirVal.Identifier, true
	case "name":
		return fhirVal.Name, true
	case "status":
		return fhirVal.Status, true
	case "immutable":
		return fhirVal.Immutable, true
	case "version":
		return fhirVal.Version, true
	case "title":
		return fhirVal.Title, true
	case "date":
		return fhirVal.Date, true
	case "publisher":
		return fhirVal.Publisher, true
	case "copyright":
		return fhirVal.Copyright, true
	case "compose":
		return fhirVal.Compose, true

	default:
		return nil, false
	}
}

func (fhirVal *ValueSet) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"extensible": &FieldTypeSupport{"bool", false, true},
		"expansion": &FieldTypeSupport{"ValueSetExpansionComponent", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"immutable": &FieldTypeSupport{"bool", false, true},
		"version": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"compose": &FieldTypeSupport{"ValueSetComposeComponent", false, true},

	}
}
