
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Bundle) TupleName() (id string, resourceType string) {
	 return fhirVal.Resource.Id, fhirVal.Resource.ResourceType
 }
 
 func (fhirVal *Bundle) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "signature":
		 return fhirVal.Signature, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "type":
		 return fhirVal.Type, true
 	case "total":
		 return fhirVal.Total, true
 	case "link":
		 return fhirVal.Link, true
 	case "entry":
		 return fhirVal.Entry, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Bundle) FieldsToTypes() map[string]*FieldTypeSupport {
	 return map[string]*FieldTypeSupport {
 
	 		"resourcetype": &FieldTypeSupport{"string", false, false},
		 "id": &FieldTypeSupport{"string", false, false},
		 "meta": &FieldTypeSupport{"Meta", false, true},
		 "implicitrules": &FieldTypeSupport{"string", false, false},
		 "language": &FieldTypeSupport{"string", false, false},
 		"signature": &FieldTypeSupport{"Signature", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"type": &FieldTypeSupport{"string", false, false},
 		"total": &FieldTypeSupport{"uint32", false, true},
 		"link": &FieldTypeSupport{"BundleLinkComponent", true, false},
 		"entry": &FieldTypeSupport{"BundleEntryComponent", true, false},
 
	 }
 }
 
 func (fhirVal *Bundle) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 