
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Flag) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Flag) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "category":
		 return fhirVal.Category, true
 	case "code":
		 return fhirVal.Code, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "period":
		 return fhirVal.Period, true
 	case "encounter":
		 return fhirVal.Encounter, true
 	case "author":
		 return fhirVal.Author, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Flag) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"category": &FieldTypeSupport{"CodeableConcept", false, true},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"encounter": &FieldTypeSupport{"Reference", false, true},
 		"author": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *Flag) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 