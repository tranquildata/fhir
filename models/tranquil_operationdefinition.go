
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
	case "code":
		return fhirVal.Code, true
	case "instance":
		return fhirVal.Instance, true
	case "name":
		return fhirVal.Name, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "status":
		return fhirVal.Status, true
	case "experimental":
		return fhirVal.Experimental, true
	case "contact":
		return fhirVal.Contact, true
	case "purpose":
		return fhirVal.Purpose, true
	case "system":
		return fhirVal.System, true
	case "type":
		return fhirVal.Type, true
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "resource":
		return fhirVal.Resource, true
	case "idempotent":
		return fhirVal.Idempotent, true
	case "comment":
		return fhirVal.Comment, true
	case "publisher":
		return fhirVal.Publisher, true
	case "description":
		return fhirVal.Description, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "base":
		return fhirVal.Base, true
	case "parameter":
		return fhirVal.Parameter, true
	case "overload":
		return fhirVal.Overload, true
	case "kind":
		return fhirVal.Kind, true
	case "date":
		return fhirVal.Date, true

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
		"code": &FieldTypeSupport{"string", false, false},
		"instance": &FieldTypeSupport{"bool", false, true},
		"name": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"system": &FieldTypeSupport{"bool", false, true},
		"type": &FieldTypeSupport{"bool", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"resource": &FieldTypeSupport{"string", true, false},
		"idempotent": &FieldTypeSupport{"bool", false, true},
		"comment": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"base": &FieldTypeSupport{"Reference", false, true},
		"parameter": &FieldTypeSupport{"OperationDefinitionParameterComponent", true, false},
		"overload": &FieldTypeSupport{"OperationDefinitionOverloadComponent", true, false},
		"kind": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
