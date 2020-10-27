
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *PaymentReconciliation) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *PaymentReconciliation) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "outcome":
		 return fhirVal.Outcome, true
 	case "form":
		 return fhirVal.Form, true
 	case "processnote":
		 return fhirVal.ProcessNote, true
 	case "organization":
		 return fhirVal.Organization, true
 	case "status":
		 return fhirVal.Status, true
 	case "created":
		 return fhirVal.Created, true
 	case "requestorganization":
		 return fhirVal.RequestOrganization, true
 	case "detail":
		 return fhirVal.Detail, true
 	case "total":
		 return fhirVal.Total, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "period":
		 return fhirVal.Period, true
 	case "request":
		 return fhirVal.Request, true
 	case "disposition":
		 return fhirVal.Disposition, true
 	case "requestprovider":
		 return fhirVal.RequestProvider, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *PaymentReconciliation) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
 		"form": &FieldTypeSupport{"CodeableConcept", false, true},
 		"processnote": &FieldTypeSupport{"PaymentReconciliationNotesComponent", true, false},
 		"organization": &FieldTypeSupport{"Reference", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"requestorganization": &FieldTypeSupport{"Reference", false, true},
 		"detail": &FieldTypeSupport{"PaymentReconciliationDetailsComponent", true, false},
 		"total": &FieldTypeSupport{"Quantity", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"request": &FieldTypeSupport{"Reference", false, true},
 		"disposition": &FieldTypeSupport{"string", false, false},
 		"requestprovider": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *PaymentReconciliation) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 