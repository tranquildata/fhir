
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *DiagnosticRequest) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *DiagnosticRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "replaces":
		 return fhirVal.Replaces, true
 	case "status":
		 return fhirVal.Status, true
 	case "code":
		 return fhirVal.Code, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "performertype":
		 return fhirVal.PerformerType, true
 	case "reasoncode":
		 return fhirVal.ReasonCode, true
 	case "note":
		 return fhirVal.Note, true
 	case "basedon":
		 return fhirVal.BasedOn, true
 	case "relevanthistory":
		 return fhirVal.RelevantHistory, true
 	case "intent":
		 return fhirVal.Intent, true
 	case "priority":
		 return fhirVal.Priority, true
 	case "occurrencedatetime":
		 return fhirVal.OccurrenceDateTime, true
 	case "authoredon":
		 return fhirVal.AuthoredOn, true
 	case "supportinginformation":
		 return fhirVal.SupportingInformation, true
 	case "requisition":
		 return fhirVal.Requisition, true
 	case "requester":
		 return fhirVal.Requester, true
 	case "performer":
		 return fhirVal.Performer, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "context":
		 return fhirVal.Context, true
 	case "occurrenceperiod":
		 return fhirVal.OccurrencePeriod, true
 	case "occurrencetiming":
		 return fhirVal.OccurrenceTiming, true
 	case "reasonreference":
		 return fhirVal.ReasonReference, true
 	case "definition":
		 return fhirVal.Definition, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *DiagnosticRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"replaces": &FieldTypeSupport{"Reference", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"performertype": &FieldTypeSupport{"CodeableConcept", false, true},
 		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"basedon": &FieldTypeSupport{"Reference", true, false},
 		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
 		"intent": &FieldTypeSupport{"string", false, false},
 		"priority": &FieldTypeSupport{"string", false, false},
 		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
 		"requisition": &FieldTypeSupport{"Identifier", false, true},
 		"requester": &FieldTypeSupport{"Reference", false, true},
 		"performer": &FieldTypeSupport{"Reference", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
 		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
 		"reasonreference": &FieldTypeSupport{"Reference", true, false},
 		"definition": &FieldTypeSupport{"Reference", true, false},
 
	 }
 }
 
 func (fhirVal *DiagnosticRequest) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 