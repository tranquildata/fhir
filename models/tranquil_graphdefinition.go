
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
	case "name":
		return fhirVal.Name, true
	case "description":
		return fhirVal.Description, true
	case "purpose":
		return fhirVal.Purpose, true
	case "link":
		return fhirVal.Link, true
	case "status":
		return fhirVal.Status, true
	case "publisher":
		return fhirVal.Publisher, true
	case "contact":
		return fhirVal.Contact, true
	case "start":
		return fhirVal.Start, true
	case "url":
		return fhirVal.Url, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "profile":
		return fhirVal.Profile, true
	case "version":
		return fhirVal.Version, true
	case "experimental":
		return fhirVal.Experimental, true
	case "date":
		return fhirVal.Date, true

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
		"name": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"link": &FieldTypeSupport{"GraphDefinitionLinkComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"start": &FieldTypeSupport{"string", false, false},
		"url": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"profile": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
