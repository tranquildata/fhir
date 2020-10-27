
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Contract) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Contract) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "applies":
		 return fhirVal.Applies, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "authority":
		 return fhirVal.Authority, true
 	case "type":
		 return fhirVal.Type, true
 	case "decisiontype":
		 return fhirVal.DecisionType, true
 	case "contentderivative":
		 return fhirVal.ContentDerivative, true
 	case "securitylabel":
		 return fhirVal.SecurityLabel, true
 	case "issued":
		 return fhirVal.Issued, true
 	case "signer":
		 return fhirVal.Signer, true
 	case "rule":
		 return fhirVal.Rule, true
 	case "action":
		 return fhirVal.Action, true
 	case "bindingattachment":
		 return fhirVal.BindingAttachment, true
 	case "legal":
		 return fhirVal.Legal, true
 	case "subtype":
		 return fhirVal.SubType, true
 	case "status":
		 return fhirVal.Status, true
 	case "topic":
		 return fhirVal.Topic, true
 	case "domain":
		 return fhirVal.Domain, true
 	case "actionreason":
		 return fhirVal.ActionReason, true
 	case "agent":
		 return fhirVal.Agent, true
 	case "valueditem":
		 return fhirVal.ValuedItem, true
 	case "term":
		 return fhirVal.Term, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "friendly":
		 return fhirVal.Friendly, true
 	case "bindingreference":
		 return fhirVal.BindingReference, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Contract) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"applies": &FieldTypeSupport{"Period", false, true},
 		"subject": &FieldTypeSupport{"Reference", true, false},
 		"authority": &FieldTypeSupport{"Reference", true, false},
 		"type": &FieldTypeSupport{"CodeableConcept", false, true},
 		"decisiontype": &FieldTypeSupport{"CodeableConcept", false, true},
 		"contentderivative": &FieldTypeSupport{"CodeableConcept", false, true},
 		"securitylabel": &FieldTypeSupport{"Coding", true, false},
 		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"signer": &FieldTypeSupport{"ContractSignatoryComponent", true, false},
 		"rule": &FieldTypeSupport{"ContractComputableLanguageComponent", true, false},
 		"action": &FieldTypeSupport{"CodeableConcept", true, false},
 		"bindingattachment": &FieldTypeSupport{"Attachment", false, true},
 		"legal": &FieldTypeSupport{"ContractLegalLanguageComponent", true, false},
 		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"topic": &FieldTypeSupport{"Reference", true, false},
 		"domain": &FieldTypeSupport{"Reference", true, false},
 		"actionreason": &FieldTypeSupport{"CodeableConcept", true, false},
 		"agent": &FieldTypeSupport{"ContractAgentComponent", true, false},
 		"valueditem": &FieldTypeSupport{"ContractValuedItemComponent", true, false},
 		"term": &FieldTypeSupport{"ContractTermComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"friendly": &FieldTypeSupport{"ContractFriendlyLanguageComponent", true, false},
 		"bindingreference": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *Contract) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 