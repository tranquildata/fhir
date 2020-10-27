
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *DomainResource) TupleName() (id string, resourceType string) {
	 return fhirVal.Resource.Id, fhirVal.Resource.ResourceType
 }
 
 func (fhirVal *DomainResource) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *DomainResource) FieldsToTypes() map[string]*FieldTypeSupport {
	 return map[string]*FieldTypeSupport {
 
	 		"resourcetype": &FieldTypeSupport{"string", false, false},
		 "id": &FieldTypeSupport{"string", false, false},
		 "meta": &FieldTypeSupport{"Meta", false, true},
		 "implicitrules": &FieldTypeSupport{"string", false, false},
		 "language": &FieldTypeSupport{"string", false, false},
 		"text": &FieldTypeSupport{"Narrative", false, true},
 		"contained": &FieldTypeSupport{"ContainedResources", false, false},
 		"extension": &FieldTypeSupport{"Extension", true, false},
 		"modifierextension": &FieldTypeSupport{"Extension", true, false},
 
	 }
 }
 
 func (fhirVal *DomainResource) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 