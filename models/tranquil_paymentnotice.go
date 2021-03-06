
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *PaymentNotice) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *PaymentNotice) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "provider":
		 return fhirVal.Provider, true
 	case "organization":
		 return fhirVal.Organization, true
 	case "status":
		 return fhirVal.Status, true
 	case "request":
		 return fhirVal.Request, true
 	case "response":
		 return fhirVal.Response, true
 	case "statusdate":
		 return fhirVal.StatusDate, true
 	case "created":
		 return fhirVal.Created, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "target":
		 return fhirVal.Target, true
 	case "paymentstatus":
		 return fhirVal.PaymentStatus, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *PaymentNotice) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"provider": &FieldTypeSupport{"Reference", false, true},
 		"organization": &FieldTypeSupport{"Reference", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"request": &FieldTypeSupport{"Reference", false, true},
 		"response": &FieldTypeSupport{"Reference", false, true},
 		"statusdate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"target": &FieldTypeSupport{"Reference", false, true},
 		"paymentstatus": &FieldTypeSupport{"CodeableConcept", false, true},
 
	 }
 }
 
 func (fhirVal *PaymentNotice) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 