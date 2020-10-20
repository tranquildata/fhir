
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
	case "experimental":
		return fhirVal.Experimental, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "purpose":
		return fhirVal.Purpose, true
	case "version":
		return fhirVal.Version, true
	case "status":
		return fhirVal.Status, true
	case "contact":
		return fhirVal.Contact, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "start":
		return fhirVal.Start, true
	case "date":
		return fhirVal.Date, true
	case "publisher":
		return fhirVal.Publisher, true
	case "profile":
		return fhirVal.Profile, true
	case "link":
		return fhirVal.Link, true
	case "url":
		return fhirVal.Url, true
	case "name":
		return fhirVal.Name, true
	case "description":
		return fhirVal.Description, true

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
		"experimental": &FieldTypeSupport{"bool", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"start": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"profile": &FieldTypeSupport{"string", false, false},
		"link": &FieldTypeSupport{"GraphDefinitionLinkComponent", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},

	}
}
