
package models


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
	case "consentingparty":
		return fhirVal.ConsentingParty, true
	case "action":
		return fhirVal.Action, true
	case "sourcereference":
		return fhirVal.SourceReference, true
	case "category":
		return fhirVal.Category, true
	case "patient":
		return fhirVal.Patient, true
	case "actor":
		return fhirVal.Actor, true
	case "organization":
		return fhirVal.Organization, true
	case "identifier":
		return fhirVal.Identifier, true
	case "period":
		return fhirVal.Period, true
	case "sourceidentifier":
		return fhirVal.SourceIdentifier, true
	case "policy":
		return fhirVal.Policy, true
	case "policyrule":
		return fhirVal.PolicyRule, true
	case "securitylabel":
		return fhirVal.SecurityLabel, true
	case "purpose":
		return fhirVal.Purpose, true
	case "data":
		return fhirVal.Data, true
	case "status":
		return fhirVal.Status, true
	case "sourceattachment":
		return fhirVal.SourceAttachment, true
	case "except":
		return fhirVal.Except, true
	case "datetime":
		return fhirVal.DateTime, true
	case "dataperiod":
		return fhirVal.DataPeriod, true

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
		"consentingparty": &FieldTypeSupport{"Reference", true, false},
		"action": &FieldTypeSupport{"CodeableConcept", true, false},
		"sourcereference": &FieldTypeSupport{"Reference", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"actor": &FieldTypeSupport{"ConsentActorComponent", true, false},
		"organization": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"period": &FieldTypeSupport{"Period", false, true},
		"sourceidentifier": &FieldTypeSupport{"Identifier", false, true},
		"policy": &FieldTypeSupport{"ConsentPolicyComponent", true, false},
		"policyrule": &FieldTypeSupport{"string", false, false},
		"securitylabel": &FieldTypeSupport{"Coding", true, false},
		"purpose": &FieldTypeSupport{"Coding", true, false},
		"data": &FieldTypeSupport{"ConsentDataComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"sourceattachment": &FieldTypeSupport{"Attachment", false, true},
		"except": &FieldTypeSupport{"ConsentExceptComponent", true, false},
		"datetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"dataperiod": &FieldTypeSupport{"Period", false, true},

	}
}
