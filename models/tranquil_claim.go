
package models


func (fhirVal *Claim) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Claim) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *Claim) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"employmentimpacted": &FieldTypeSupport{"Period", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"careteam": &FieldTypeSupport{"ClaimCareTeamComponent", true, false},
		"payee": &FieldTypeSupport{"ClaimPayeeComponent", false, true},
		"information": &FieldTypeSupport{"ClaimSpecialConditionComponent", true, false},
		"total": &FieldTypeSupport{"Quantity", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"hospitalization": &FieldTypeSupport{"Period", false, true},
		"related": &FieldTypeSupport{"ClaimRelatedClaimComponent", true, false},
		"procedure": &FieldTypeSupport{"ClaimProcedureComponent", true, false},
		"item": &FieldTypeSupport{"ClaimItemComponent", true, false},
		"use": &FieldTypeSupport{"string", false, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"provider": &FieldTypeSupport{"Reference", false, true},
		"originalprescription": &FieldTypeSupport{"Reference", false, true},
		"facility": &FieldTypeSupport{"Reference", false, true},
		"insurance": &FieldTypeSupport{"ClaimInsuranceComponent", true, false},
		"accident": &FieldTypeSupport{"ClaimAccidentComponent", false, true},
		"referral": &FieldTypeSupport{"Reference", false, true},
		"diagnosis": &FieldTypeSupport{"ClaimDiagnosisComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"billableperiod": &FieldTypeSupport{"Period", false, true},
		"priority": &FieldTypeSupport{"CodeableConcept", false, true},
		"fundsreserve": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
