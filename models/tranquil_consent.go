
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
		"organization": &FieldTypeSupport{"Reference", true, false},
		"data": &FieldTypeSupport{"ConsentDataComponent", true, false},
		"except": &FieldTypeSupport{"ConsentExceptComponent", true, false},
		"datetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"sourceidentifier": &FieldTypeSupport{"Identifier", false, true},
		"securitylabel": &FieldTypeSupport{"Coding", true, false},
		"purpose": &FieldTypeSupport{"Coding", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"actor": &FieldTypeSupport{"ConsentActorComponent", true, false},
		"sourcereference": &FieldTypeSupport{"Reference", false, true},
		"dataperiod": &FieldTypeSupport{"Period", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"period": &FieldTypeSupport{"Period", false, true},
		"sourceattachment": &FieldTypeSupport{"Attachment", false, true},
		"policy": &FieldTypeSupport{"ConsentPolicyComponent", true, false},
		"policyrule": &FieldTypeSupport{"string", false, false},

	}
}
