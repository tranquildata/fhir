
package models


func (fhirVal *ImplementationGuide) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ImplementationGuide) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "contact":
		return fhirVal.Contact, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "copyright":
		return fhirVal.Copyright, true
	case "package":
		return fhirVal.Package, true
	case "page":
		return fhirVal.Page, true
	case "global":
		return fhirVal.Global, true
	case "binary":
		return fhirVal.Binary, true
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "experimental":
		return fhirVal.Experimental, true
	case "fhirversion":
		return fhirVal.FhirVersion, true
	case "dependency":
		return fhirVal.Dependency, true
	case "name":
		return fhirVal.Name, true
	case "status":
		return fhirVal.Status, true
	case "date":
		return fhirVal.Date, true
	case "publisher":
		return fhirVal.Publisher, true
	case "description":
		return fhirVal.Description, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true

	default:
		return nil, false
	}
}

func (fhirVal *ImplementationGuide) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"package": &FieldTypeSupport{"ImplementationGuidePackageComponent", true, false},
		"page": &FieldTypeSupport{"ImplementationGuidePageComponent", false, true},
		"global": &FieldTypeSupport{"ImplementationGuideGlobalComponent", true, false},
		"binary": &FieldTypeSupport{"string", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"fhirversion": &FieldTypeSupport{"string", false, false},
		"dependency": &FieldTypeSupport{"ImplementationGuideDependencyComponent", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}