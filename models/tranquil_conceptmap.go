
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ConceptMap) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ConceptMap) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "version":
		 return fhirVal.Version, true
 	case "name":
		 return fhirVal.Name, true
 	case "description":
		 return fhirVal.Description, true
 	case "targetreference":
		 return fhirVal.TargetReference, true
 	case "group":
		 return fhirVal.Group, true
 	case "sourcereference":
		 return fhirVal.SourceReference, true
 	case "title":
		 return fhirVal.Title, true
 	case "status":
		 return fhirVal.Status, true
 	case "date":
		 return fhirVal.Date, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "sourceuri":
		 return fhirVal.SourceUri, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "copyright":
		 return fhirVal.Copyright, true
 	case "targeturi":
		 return fhirVal.TargetUri, true
 	case "url":
		 return fhirVal.Url, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "contact":
		 return fhirVal.Contact, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ConceptMap) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"version": &FieldTypeSupport{"string", false, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"targetreference": &FieldTypeSupport{"Reference", false, true},
 		"group": &FieldTypeSupport{"ConceptMapGroupComponent", true, false},
 		"sourcereference": &FieldTypeSupport{"Reference", false, true},
 		"title": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"sourceuri": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"copyright": &FieldTypeSupport{"string", false, false},
 		"targeturi": &FieldTypeSupport{"string", false, false},
 		"url": &FieldTypeSupport{"string", false, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 
	 }
 }
 
 func (fhirVal *ConceptMap) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 