
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ChargeItem) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ChargeItem) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "account":
		 return fhirVal.Account, true
 	case "definition":
		 return fhirVal.Definition, true
 	case "partof":
		 return fhirVal.PartOf, true
 	case "code":
		 return fhirVal.Code, true
 	case "quantity":
		 return fhirVal.Quantity, true
 	case "factoroverride":
		 return fhirVal.FactorOverride, true
 	case "overridereason":
		 return fhirVal.OverrideReason, true
 	case "occurrencedatetime":
		 return fhirVal.OccurrenceDateTime, true
 	case "occurrenceperiod":
		 return fhirVal.OccurrencePeriod, true
 	case "occurrencetiming":
		 return fhirVal.OccurrenceTiming, true
 	case "enterer":
		 return fhirVal.Enterer, true
 	case "service":
		 return fhirVal.Service, true
 	case "participant":
		 return fhirVal.Participant, true
 	case "bodysite":
		 return fhirVal.Bodysite, true
 	case "priceoverride":
		 return fhirVal.PriceOverride, true
 	case "supportinginformation":
		 return fhirVal.SupportingInformation, true
 	case "entereddate":
		 return fhirVal.EnteredDate, true
 	case "reason":
		 return fhirVal.Reason, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "context":
		 return fhirVal.Context, true
 	case "performingorganization":
		 return fhirVal.PerformingOrganization, true
 	case "requestingorganization":
		 return fhirVal.RequestingOrganization, true
 	case "note":
		 return fhirVal.Note, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ChargeItem) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"account": &FieldTypeSupport{"Reference", true, false},
 		"definition": &FieldTypeSupport{"string", true, false},
 		"partof": &FieldTypeSupport{"Reference", true, false},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"quantity": &FieldTypeSupport{"Quantity", false, true},
 		"factoroverride": &FieldTypeSupport{"float64", false, true},
 		"overridereason": &FieldTypeSupport{"string", false, false},
 		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
 		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
 		"enterer": &FieldTypeSupport{"Reference", false, true},
 		"service": &FieldTypeSupport{"Reference", true, false},
 		"participant": &FieldTypeSupport{"ChargeItemParticipantComponent", true, false},
 		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
 		"priceoverride": &FieldTypeSupport{"Quantity", false, true},
 		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
 		"entereddate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"reason": &FieldTypeSupport{"CodeableConcept", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"performingorganization": &FieldTypeSupport{"Reference", false, true},
 		"requestingorganization": &FieldTypeSupport{"Reference", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 
	 }
 }
 
 func (fhirVal *ChargeItem) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 