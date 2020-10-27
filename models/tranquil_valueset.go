
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ValueSet) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ValueSet) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "title":
		 return fhirVal.Title, true
 	case "date":
		 return fhirVal.Date, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "copyright":
		 return fhirVal.Copyright, true
 	case "compose":
		 return fhirVal.Compose, true
 	case "url":
		 return fhirVal.Url, true
 	case "status":
		 return fhirVal.Status, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "immutable":
		 return fhirVal.Immutable, true
 	case "extensible":
		 return fhirVal.Extensible, true
 	case "expansion":
		 return fhirVal.Expansion, true
 	case "version":
		 return fhirVal.Version, true
 	case "description":
		 return fhirVal.Description, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ValueSet) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"title": &FieldTypeSupport{"string", false, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"copyright": &FieldTypeSupport{"string", false, false},
 		"compose": &FieldTypeSupport{"ValueSetComposeComponent", false, true},
 		"url": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"immutable": &FieldTypeSupport{"bool", false, true},
 		"extensible": &FieldTypeSupport{"bool", false, true},
 		"expansion": &FieldTypeSupport{"ValueSetExpansionComponent", false, true},
 		"version": &FieldTypeSupport{"string", false, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 
	 }
 }
 
 func (fhirVal *ValueSet) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 