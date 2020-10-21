
package models


func (fhirVal *GraphDefinition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *GraphDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "date":
		return fhirVal.Date, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "status":
		return fhirVal.Status, true
	case "publisher":
		return fhirVal.Publisher, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "name":
		return fhirVal.Name, true
	case "experimental":
		return fhirVal.Experimental, true
	case "contact":
		return fhirVal.Contact, true
	case "description":
		return fhirVal.Description, true
	case "purpose":
		return fhirVal.Purpose, true
	case "start":
		return fhirVal.Start, true
	case "profile":
		return fhirVal.Profile, true
	case "link":
		return fhirVal.Link, true

	default:
		return nil, false
	}
}

func (fhirVal *GraphDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"start": &FieldTypeSupport{"string", false, false},
		"profile": &FieldTypeSupport{"string", false, false},
		"link": &FieldTypeSupport{"GraphDefinitionLinkComponent", true, false},

	}
}
