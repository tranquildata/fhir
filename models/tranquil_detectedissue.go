
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *DetectedIssue) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *DetectedIssue) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "date":
		 return fhirVal.Date, true
 	case "author":
		 return fhirVal.Author, true
 	case "implicated":
		 return fhirVal.Implicated, true
 	case "status":
		 return fhirVal.Status, true
 	case "severity":
		 return fhirVal.Severity, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "detail":
		 return fhirVal.Detail, true
 	case "reference":
		 return fhirVal.Reference, true
 	case "mitigation":
		 return fhirVal.Mitigation, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "category":
		 return fhirVal.Category, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *DetectedIssue) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"author": &FieldTypeSupport{"Reference", false, true},
 		"implicated": &FieldTypeSupport{"Reference", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"severity": &FieldTypeSupport{"string", false, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"detail": &FieldTypeSupport{"string", false, false},
 		"reference": &FieldTypeSupport{"string", false, false},
 		"mitigation": &FieldTypeSupport{"DetectedIssueMitigationComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"category": &FieldTypeSupport{"CodeableConcept", false, true},
 
	 }
 }
 
 func (fhirVal *DetectedIssue) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 