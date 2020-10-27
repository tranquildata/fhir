
 package models
 
 import (
	 "fmt"
 )
 
 
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
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "code":
		 return fhirVal.Code, true
 	case "resource":
		 return fhirVal.Resource, true
 	case "instance":
		 return fhirVal.Instance, true
 	case "version":
		 return fhirVal.Version, true
 	case "status":
		 return fhirVal.Status, true
 	case "date":
		 return fhirVal.Date, true
 	case "description":
		 return fhirVal.Description, true
 	case "parameter":
		 return fhirVal.Parameter, true
 	case "comment":
		 return fhirVal.Comment, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "overload":
		 return fhirVal.Overload, true
 	case "url":
		 return fhirVal.Url, true
 	case "name":
		 return fhirVal.Name, true
 	case "kind":
		 return fhirVal.Kind, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "system":
		 return fhirVal.System, true
 	case "type":
		 return fhirVal.Type, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "idempotent":
		 return fhirVal.Idempotent, true
 	case "base":
		 return fhirVal.Base, true
 
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
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"code": &FieldTypeSupport{"string", false, false},
 		"resource": &FieldTypeSupport{"string", true, false},
 		"instance": &FieldTypeSupport{"bool", false, true},
 		"version": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"description": &FieldTypeSupport{"string", false, false},
 		"parameter": &FieldTypeSupport{"OperationDefinitionParameterComponent", true, false},
 		"comment": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"overload": &FieldTypeSupport{"OperationDefinitionOverloadComponent", true, false},
 		"url": &FieldTypeSupport{"string", false, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"kind": &FieldTypeSupport{"string", false, false},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"system": &FieldTypeSupport{"bool", false, true},
 		"type": &FieldTypeSupport{"bool", false, true},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"idempotent": &FieldTypeSupport{"bool", false, true},
 		"base": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *OperationDefinition) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 