
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
	case "authority":
		return fhirVal.Authority, true
	case "contentderivative":
		return fhirVal.ContentDerivative, true
	case "securitylabel":
		return fhirVal.SecurityLabel, true
	case "bindingreference":
		return fhirVal.BindingReference, true
	case "identifier":
		return fhirVal.Identifier, true
	case "issued":
		return fhirVal.Issued, true
	case "actionreason":
		return fhirVal.ActionReason, true
	case "decisiontype":
		return fhirVal.DecisionType, true
	case "valueditem":
		return fhirVal.ValuedItem, true
	case "bindingattachment":
		return fhirVal.BindingAttachment, true
	case "friendly":
		return fhirVal.Friendly, true
	case "applies":
		return fhirVal.Applies, true
	case "topic":
		return fhirVal.Topic, true
	case "type":
		return fhirVal.Type, true
	case "action":
		return fhirVal.Action, true
	case "agent":
		return fhirVal.Agent, true
	case "signer":
		return fhirVal.Signer, true
	case "legal":
		return fhirVal.Legal, true
	case "status":
		return fhirVal.Status, true
	case "subject":
		return fhirVal.Subject, true
	case "domain":
		return fhirVal.Domain, true
	case "subtype":
		return fhirVal.SubType, true
	case "term":
		return fhirVal.Term, true
	case "rule":
		return fhirVal.Rule, true

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
		"authority": &FieldTypeSupport{"Reference", true, false},
		"contentderivative": &FieldTypeSupport{"CodeableConcept", false, true},
		"securitylabel": &FieldTypeSupport{"Coding", true, false},
		"bindingreference": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"actionreason": &FieldTypeSupport{"CodeableConcept", true, false},
		"decisiontype": &FieldTypeSupport{"CodeableConcept", false, true},
		"valueditem": &FieldTypeSupport{"ContractValuedItemComponent", true, false},
		"bindingattachment": &FieldTypeSupport{"Attachment", false, true},
		"friendly": &FieldTypeSupport{"ContractFriendlyLanguageComponent", true, false},
		"applies": &FieldTypeSupport{"Period", false, true},
		"topic": &FieldTypeSupport{"Reference", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"action": &FieldTypeSupport{"CodeableConcept", true, false},
		"agent": &FieldTypeSupport{"ContractAgentComponent", true, false},
		"signer": &FieldTypeSupport{"ContractSignatoryComponent", true, false},
		"legal": &FieldTypeSupport{"ContractLegalLanguageComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", true, false},
		"domain": &FieldTypeSupport{"Reference", true, false},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"term": &FieldTypeSupport{"ContractTermComponent", true, false},
		"rule": &FieldTypeSupport{"ContractComputableLanguageComponent", true, false},

	}
}
