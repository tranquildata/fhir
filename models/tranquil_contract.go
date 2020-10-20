
package models


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
	case "actionreason":
		return fhirVal.ActionReason, true
	case "signer":
		return fhirVal.Signer, true
	case "legal":
		return fhirVal.Legal, true
	case "rule":
		return fhirVal.Rule, true
	case "identifier":
		return fhirVal.Identifier, true
	case "subject":
		return fhirVal.Subject, true
	case "topic":
		return fhirVal.Topic, true
	case "action":
		return fhirVal.Action, true
	case "bindingattachment":
		return fhirVal.BindingAttachment, true
	case "bindingreference":
		return fhirVal.BindingReference, true
	case "status":
		return fhirVal.Status, true
	case "applies":
		return fhirVal.Applies, true
	case "decisiontype":
		return fhirVal.DecisionType, true
	case "contentderivative":
		return fhirVal.ContentDerivative, true
	case "friendly":
		return fhirVal.Friendly, true
	case "issued":
		return fhirVal.Issued, true
	case "authority":
		return fhirVal.Authority, true
	case "domain":
		return fhirVal.Domain, true
	case "agent":
		return fhirVal.Agent, true
	case "term":
		return fhirVal.Term, true
	case "type":
		return fhirVal.Type, true
	case "subtype":
		return fhirVal.SubType, true
	case "securitylabel":
		return fhirVal.SecurityLabel, true
	case "valueditem":
		return fhirVal.ValuedItem, true

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
		"actionreason": &FieldTypeSupport{"CodeableConcept", true, false},
		"signer": &FieldTypeSupport{"ContractSignatoryComponent", true, false},
		"legal": &FieldTypeSupport{"ContractLegalLanguageComponent", true, false},
		"rule": &FieldTypeSupport{"ContractComputableLanguageComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"subject": &FieldTypeSupport{"Reference", true, false},
		"topic": &FieldTypeSupport{"Reference", true, false},
		"action": &FieldTypeSupport{"CodeableConcept", true, false},
		"bindingattachment": &FieldTypeSupport{"Attachment", false, true},
		"bindingreference": &FieldTypeSupport{"Reference", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"applies": &FieldTypeSupport{"Period", false, true},
		"decisiontype": &FieldTypeSupport{"CodeableConcept", false, true},
		"contentderivative": &FieldTypeSupport{"CodeableConcept", false, true},
		"friendly": &FieldTypeSupport{"ContractFriendlyLanguageComponent", true, false},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"authority": &FieldTypeSupport{"Reference", true, false},
		"domain": &FieldTypeSupport{"Reference", true, false},
		"agent": &FieldTypeSupport{"ContractAgentComponent", true, false},
		"term": &FieldTypeSupport{"ContractTermComponent", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"securitylabel": &FieldTypeSupport{"Coding", true, false},
		"valueditem": &FieldTypeSupport{"ContractValuedItemComponent", true, false},

	}
}
