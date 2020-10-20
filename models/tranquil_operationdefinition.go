
package models


func (fhirVal *OperationDefinition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *OperationDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "comment":
		return fhirVal.Comment, true
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "kind":
		return fhirVal.Kind, true
	case "publisher":
		return fhirVal.Publisher, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "purpose":
		return fhirVal.Purpose, true
	case "code":
		return fhirVal.Code, true
	case "name":
		return fhirVal.Name, true
	case "status":
		return fhirVal.Status, true
	case "experimental":
		return fhirVal.Experimental, true
	case "date":
		return fhirVal.Date, true
	case "instance":
		return fhirVal.Instance, true
	case "overload":
		return fhirVal.Overload, true
	case "contact":
		return fhirVal.Contact, true
	case "description":
		return fhirVal.Description, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "idempotent":
		return fhirVal.Idempotent, true
	case "base":
		return fhirVal.Base, true
	case "system":
		return fhirVal.System, true
	case "type":
		return fhirVal.Type, true
	case "resource":
		return fhirVal.Resource, true
	case "parameter":
		return fhirVal.Parameter, true

	default:
		return nil, false
	}
}

func (fhirVal *OperationDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"comment": &FieldTypeSupport{"string", false, false},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"kind": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"instance": &FieldTypeSupport{"bool", false, true},
		"overload": &FieldTypeSupport{"OperationDefinitionOverloadComponent", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"idempotent": &FieldTypeSupport{"bool", false, true},
		"base": &FieldTypeSupport{"Reference", false, true},
		"system": &FieldTypeSupport{"bool", false, true},
		"type": &FieldTypeSupport{"bool", false, true},
		"resource": &FieldTypeSupport{"string", true, false},
		"parameter": &FieldTypeSupport{"OperationDefinitionParameterComponent", true, false},

	}
}
