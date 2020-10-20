
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
		"bindingreference": &FieldTypeSupport{"Reference", false, true},
		"friendly": &FieldTypeSupport{"ContractFriendlyLanguageComponent", true, false},
		"legal": &FieldTypeSupport{"ContractLegalLanguageComponent", true, false},
		"rule": &FieldTypeSupport{"ContractComputableLanguageComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"bindingattachment": &FieldTypeSupport{"Attachment", false, true},
		"action": &FieldTypeSupport{"CodeableConcept", true, false},
		"actionreason": &FieldTypeSupport{"CodeableConcept", true, false},
		"decisiontype": &FieldTypeSupport{"CodeableConcept", false, true},
		"contentderivative": &FieldTypeSupport{"CodeableConcept", false, true},
		"agent": &FieldTypeSupport{"ContractAgentComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", true, false},
		"domain": &FieldTypeSupport{"Reference", true, false},
		"signer": &FieldTypeSupport{"ContractSignatoryComponent", true, false},
		"topic": &FieldTypeSupport{"Reference", true, false},
		"authority": &FieldTypeSupport{"Reference", true, false},
		"securitylabel": &FieldTypeSupport{"Coding", true, false},
		"valueditem": &FieldTypeSupport{"ContractValuedItemComponent", true, false},
		"term": &FieldTypeSupport{"ContractTermComponent", true, false},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"applies": &FieldTypeSupport{"Period", false, true},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
