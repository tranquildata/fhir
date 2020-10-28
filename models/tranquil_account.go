
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Account) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Account) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "balance":
		 return fhirVal.Balance, true
 	case "coverage":
		 return fhirVal.Coverage, true
 	case "owner":
		 return fhirVal.Owner, true
 	case "description":
		 return fhirVal.Description, true
 	case "guarantor":
		 return fhirVal.Guarantor, true
 	case "type":
		 return fhirVal.Type, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "active":
		 return fhirVal.Active, true
 	case "period":
		 return fhirVal.Period, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "name":
		 return fhirVal.Name, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Account) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"balance": &FieldTypeSupport{"Quantity", false, true},
 		"coverage": &FieldTypeSupport{"AccountCoverageComponent", true, false},
 		"owner": &FieldTypeSupport{"Reference", false, true},
 		"description": &FieldTypeSupport{"string", false, false},
 		"guarantor": &FieldTypeSupport{"AccountGuarantorComponent", true, false},
 		"type": &FieldTypeSupport{"CodeableConcept", false, true},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"active": &FieldTypeSupport{"Period", false, true},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"name": &FieldTypeSupport{"string", false, false},
 
	 }
 }
 
 func (fhirVal *Account) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 