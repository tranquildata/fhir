
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
	case "action":
		return fhirVal.Action, true
	case "sourcereference":
		return fhirVal.SourceReference, true
	case "policyrule":
		return fhirVal.PolicyRule, true
	case "except":
		return fhirVal.Except, true
	case "status":
		return fhirVal.Status, true
	case "datetime":
		return fhirVal.DateTime, true
	case "actor":
		return fhirVal.Actor, true
	case "policy":
		return fhirVal.Policy, true
	case "sourceattachment":
		return fhirVal.SourceAttachment, true
	case "identifier":
		return fhirVal.Identifier, true
	case "patient":
		return fhirVal.Patient, true
	case "period":
		return fhirVal.Period, true
	case "organization":
		return fhirVal.Organization, true
	case "purpose":
		return fhirVal.Purpose, true
	case "dataperiod":
		return fhirVal.DataPeriod, true
	case "data":
		return fhirVal.Data, true
	case "category":
		return fhirVal.Category, true
	case "consentingparty":
		return fhirVal.ConsentingParty, true
	case "sourceidentifier":
		return fhirVal.SourceIdentifier, true
	case "securitylabel":
		return fhirVal.SecurityLabel, true

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
		"action": &FieldTypeSupport{"CodeableConcept", true, false},
		"sourcereference": &FieldTypeSupport{"Reference", false, true},
		"policyrule": &FieldTypeSupport{"string", false, false},
		"except": &FieldTypeSupport{"ConsentExceptComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"datetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"actor": &FieldTypeSupport{"ConsentActorComponent", true, false},
		"policy": &FieldTypeSupport{"ConsentPolicyComponent", true, false},
		"sourceattachment": &FieldTypeSupport{"Attachment", false, true},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"period": &FieldTypeSupport{"Period", false, true},
		"organization": &FieldTypeSupport{"Reference", true, false},
		"purpose": &FieldTypeSupport{"Coding", true, false},
		"dataperiod": &FieldTypeSupport{"Period", false, true},
		"data": &FieldTypeSupport{"ConsentDataComponent", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"consentingparty": &FieldTypeSupport{"Reference", true, false},
		"sourceidentifier": &FieldTypeSupport{"Identifier", false, true},
		"securitylabel": &FieldTypeSupport{"Coding", true, false},

	}
}
