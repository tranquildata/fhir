
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *DataElement) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *DataElement) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "status":
		 return fhirVal.Status, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "title":
		 return fhirVal.Title, true
 	case "version":
		 return fhirVal.Version, true
 	case "date":
		 return fhirVal.Date, true
 	case "name":
		 return fhirVal.Name, true
 	case "mapping":
		 return fhirVal.Mapping, true
 	case "element":
		 return fhirVal.Element, true
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
 	case "copyright":
		 return fhirVal.Copyright, true
 	case "stringency":
		 return fhirVal.Stringency, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *DataElement) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"status": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"title": &FieldTypeSupport{"string", false, false},
 		"version": &FieldTypeSupport{"string", false, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"name": &FieldTypeSupport{"string", false, false},
 		"mapping": &FieldTypeSupport{"DataElementMappingComponent", true, false},
 		"element": &FieldTypeSupport{"ElementDefinition", true, false},
 		"url": &FieldTypeSupport{"string", false, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"copyright": &FieldTypeSupport{"string", false, false},
 		"stringency": &FieldTypeSupport{"string", false, false},
 
	 }
 }
 
 func (fhirVal *DataElement) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 