
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Consent) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Consent) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "category":
		 return fhirVal.Category, true
 	case "datetime":
		 return fhirVal.DateTime, true
 	case "action":
		 return fhirVal.Action, true
 	case "policy":
		 return fhirVal.Policy, true
 	case "actor":
		 return fhirVal.Actor, true
 	case "sourcereference":
		 return fhirVal.SourceReference, true
 	case "period":
		 return fhirVal.Period, true
 	case "sourceattachment":
		 return fhirVal.SourceAttachment, true
 	case "sourceidentifier":
		 return fhirVal.SourceIdentifier, true
 	case "policyrule":
		 return fhirVal.PolicyRule, true
 	case "dataperiod":
		 return fhirVal.DataPeriod, true
 	case "except":
		 return fhirVal.Except, true
 	case "data":
		 return fhirVal.Data, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "consentingparty":
		 return fhirVal.ConsentingParty, true
 	case "organization":
		 return fhirVal.Organization, true
 	case "securitylabel":
		 return fhirVal.SecurityLabel, true
 	case "purpose":
		 return fhirVal.Purpose, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Consent) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"category": &FieldTypeSupport{"CodeableConcept", true, false},
 		"datetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"action": &FieldTypeSupport{"CodeableConcept", true, false},
 		"policy": &FieldTypeSupport{"ConsentPolicyComponent", true, false},
 		"actor": &FieldTypeSupport{"ConsentActorComponent", true, false},
 		"sourcereference": &FieldTypeSupport{"Reference", false, true},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"sourceattachment": &FieldTypeSupport{"Attachment", false, true},
 		"sourceidentifier": &FieldTypeSupport{"Identifier", false, true},
 		"policyrule": &FieldTypeSupport{"string", false, false},
 		"dataperiod": &FieldTypeSupport{"Period", false, true},
 		"except": &FieldTypeSupport{"ConsentExceptComponent", true, false},
 		"data": &FieldTypeSupport{"ConsentDataComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"consentingparty": &FieldTypeSupport{"Reference", true, false},
 		"organization": &FieldTypeSupport{"Reference", true, false},
 		"securitylabel": &FieldTypeSupport{"Coding", true, false},
 		"purpose": &FieldTypeSupport{"Coding", true, false},
 
	 }
 }
 
 func (fhirVal *Consent) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 