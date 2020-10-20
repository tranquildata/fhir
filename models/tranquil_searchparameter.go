
package models


func (fhirVal *SearchParameter) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *SearchParameter) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "publisher":
		return fhirVal.Publisher, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "code":
		return fhirVal.Code, true
	case "xpath":
		return fhirVal.Xpath, true
	case "xpathusage":
		return fhirVal.XpathUsage, true
	case "url":
		return fhirVal.Url, true
	case "experimental":
		return fhirVal.Experimental, true
	case "contact":
		return fhirVal.Contact, true
	case "purpose":
		return fhirVal.Purpose, true
	case "version":
		return fhirVal.Version, true
	case "status":
		return fhirVal.Status, true
	case "expression":
		return fhirVal.Expression, true
	case "target":
		return fhirVal.Target, true
	case "chain":
		return fhirVal.Chain, true
	case "name":
		return fhirVal.Name, true
	case "base":
		return fhirVal.Base, true
	case "type":
		return fhirVal.Type, true
	case "derivedfrom":
		return fhirVal.DerivedFrom, true
	case "description":
		return fhirVal.Description, true
	case "comparator":
		return fhirVal.Comparator, true
	case "modifier":
		return fhirVal.Modifier, true
	case "component":
		return fhirVal.Component, true

	default:
		return nil, false
	}
}

func (fhirVal *SearchParameter) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"publisher": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"code": &FieldTypeSupport{"string", false, false},
		"xpath": &FieldTypeSupport{"string", false, false},
		"xpathusage": &FieldTypeSupport{"string", false, false},
		"url": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"expression": &FieldTypeSupport{"string", false, false},
		"target": &FieldTypeSupport{"string", true, false},
		"chain": &FieldTypeSupport{"string", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"base": &FieldTypeSupport{"string", true, false},
		"type": &FieldTypeSupport{"string", false, false},
		"derivedfrom": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"comparator": &FieldTypeSupport{"string", true, false},
		"modifier": &FieldTypeSupport{"string", true, false},
		"component": &FieldTypeSupport{"SearchParameterComponentComponent", true, false},

	}
}
