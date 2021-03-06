
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *CompartmentDefinition) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *CompartmentDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "url":
		 return fhirVal.Url, true
 	case "status":
		 return fhirVal.Status, true
 	case "description":
		 return fhirVal.Description, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "code":
		 return fhirVal.Code, true
 	case "search":
		 return fhirVal.Search, true
 	case "resource":
		 return fhirVal.Resource, true
 	case "name":
		 return fhirVal.Name, true
 	case "title":
		 return fhirVal.Title, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "date":
		 return fhirVal.Date, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "purpose":
		 return fhirVal.Purpose, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *CompartmentDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"url": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"code": &FieldTypeSupport{"string", false, false},
 		"search": &FieldTypeSupport{"bool", false, true},
 		"resource": &FieldTypeSupport{"CompartmentDefinitionResourceComponent", true, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"title": &FieldTypeSupport{"string", false, false},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 
	 }
 }
 
 func (fhirVal *CompartmentDefinition) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 