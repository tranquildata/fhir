
 package models
 
 import (
	 "fmt"
 )
 
 
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
 	case "target":
		 return fhirVal.Target, true
 	case "comparator":
		 return fhirVal.Comparator, true
 	case "chain":
		 return fhirVal.Chain, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "base":
		 return fhirVal.Base, true
 	case "description":
		 return fhirVal.Description, true
 	case "xpath":
		 return fhirVal.Xpath, true
 	case "xpathusage":
		 return fhirVal.XpathUsage, true
 	case "version":
		 return fhirVal.Version, true
 	case "status":
		 return fhirVal.Status, true
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "type":
		 return fhirVal.Type, true
 	case "modifier":
		 return fhirVal.Modifier, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "date":
		 return fhirVal.Date, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "code":
		 return fhirVal.Code, true
 	case "derivedfrom":
		 return fhirVal.DerivedFrom, true
 	case "expression":
		 return fhirVal.Expression, true
 	case "component":
		 return fhirVal.Component, true
 	case "url":
		 return fhirVal.Url, true
 	case "name":
		 return fhirVal.Name, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 
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
 		"target": &FieldTypeSupport{"string", true, false},
 		"comparator": &FieldTypeSupport{"string", true, false},
 		"chain": &FieldTypeSupport{"string", true, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"base": &FieldTypeSupport{"string", true, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"xpath": &FieldTypeSupport{"string", false, false},
 		"xpathusage": &FieldTypeSupport{"string", false, false},
 		"version": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"type": &FieldTypeSupport{"string", false, false},
 		"modifier": &FieldTypeSupport{"string", true, false},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"code": &FieldTypeSupport{"string", false, false},
 		"derivedfrom": &FieldTypeSupport{"string", false, false},
 		"expression": &FieldTypeSupport{"string", false, false},
 		"component": &FieldTypeSupport{"SearchParameterComponentComponent", true, false},
 		"url": &FieldTypeSupport{"string", false, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 
	 }
 }
 
 func (fhirVal *SearchParameter) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 