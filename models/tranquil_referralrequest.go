
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ReferralRequest) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ReferralRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "note":
		 return fhirVal.Note, true
 	case "recipient":
		 return fhirVal.Recipient, true
 	case "relevanthistory":
		 return fhirVal.RelevantHistory, true
 	case "replaces":
		 return fhirVal.Replaces, true
 	case "intent":
		 return fhirVal.Intent, true
 	case "servicerequested":
		 return fhirVal.ServiceRequested, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "requester":
		 return fhirVal.Requester, true
 	case "specialty":
		 return fhirVal.Specialty, true
 	case "description":
		 return fhirVal.Description, true
 	case "supportinginfo":
		 return fhirVal.SupportingInfo, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "basedon":
		 return fhirVal.BasedOn, true
 	case "status":
		 return fhirVal.Status, true
 	case "priority":
		 return fhirVal.Priority, true
 	case "occurrenceperiod":
		 return fhirVal.OccurrencePeriod, true
 	case "reasonreference":
		 return fhirVal.ReasonReference, true
 	case "reasoncode":
		 return fhirVal.ReasonCode, true
 	case "definition":
		 return fhirVal.Definition, true
 	case "groupidentifier":
		 return fhirVal.GroupIdentifier, true
 	case "type":
		 return fhirVal.Type, true
 	case "context":
		 return fhirVal.Context, true
 	case "occurrencedatetime":
		 return fhirVal.OccurrenceDateTime, true
 	case "authoredon":
		 return fhirVal.AuthoredOn, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ReferralRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"recipient": &FieldTypeSupport{"Reference", true, false},
 		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
 		"replaces": &FieldTypeSupport{"Reference", true, false},
 		"intent": &FieldTypeSupport{"string", false, false},
 		"servicerequested": &FieldTypeSupport{"CodeableConcept", true, false},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"requester": &FieldTypeSupport{"ReferralRequestRequesterComponent", false, true},
 		"specialty": &FieldTypeSupport{"CodeableConcept", false, true},
 		"description": &FieldTypeSupport{"string", false, false},
 		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"basedon": &FieldTypeSupport{"Reference", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"priority": &FieldTypeSupport{"string", false, false},
 		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
 		"reasonreference": &FieldTypeSupport{"Reference", true, false},
 		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
 		"definition": &FieldTypeSupport{"Reference", true, false},
 		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
 		"type": &FieldTypeSupport{"CodeableConcept", false, true},
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
 
	 }
 }
 
 func (fhirVal *ReferralRequest) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 