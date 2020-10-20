
package models


func (fhirVal *ExplanationOfBenefit) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ExplanationOfBenefit) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *ExplanationOfBenefit) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"employmentimpacted": &FieldTypeSupport{"Period", false, true},
		"additem": &FieldTypeSupport{"ExplanationOfBenefitAddedItemComponent", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"facility": &FieldTypeSupport{"Reference", false, true},
		"claimresponse": &FieldTypeSupport{"Reference", false, true},
		"precedence": &FieldTypeSupport{"uint32", false, true},
		"referral": &FieldTypeSupport{"Reference", false, true},
		"diagnosis": &FieldTypeSupport{"ExplanationOfBenefitDiagnosisComponent", true, false},
		"insurance": &FieldTypeSupport{"ExplanationOfBenefitInsuranceComponent", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"totalbenefit": &FieldTypeSupport{"Quantity", false, true},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"information": &FieldTypeSupport{"ExplanationOfBenefitSupportingInformationComponent", true, false},
		"procedure": &FieldTypeSupport{"ExplanationOfBenefitProcedureComponent", true, false},
		"accident": &FieldTypeSupport{"ExplanationOfBenefitAccidentComponent", false, true},
		"hospitalization": &FieldTypeSupport{"Period", false, true},
		"billableperiod": &FieldTypeSupport{"Period", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"disposition": &FieldTypeSupport{"string", false, false},
		"totalcost": &FieldTypeSupport{"Quantity", false, true},
		"unallocdeductable": &FieldTypeSupport{"Quantity", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"claim": &FieldTypeSupport{"Reference", false, true},
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"originalprescription": &FieldTypeSupport{"Reference", false, true},
		"careteam": &FieldTypeSupport{"ExplanationOfBenefitCareTeamComponent", true, false},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"related": &FieldTypeSupport{"ExplanationOfBenefitRelatedClaimComponent", true, false},
		"item": &FieldTypeSupport{"ExplanationOfBenefitItemComponent", true, false},
		"processnote": &FieldTypeSupport{"ExplanationOfBenefitNoteComponent", true, false},
		"benefitbalance": &FieldTypeSupport{"ExplanationOfBenefitBenefitBalanceComponent", true, false},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"provider": &FieldTypeSupport{"Reference", false, true},
		"payee": &FieldTypeSupport{"ExplanationOfBenefitPayeeComponent", false, true},
		"payment": &FieldTypeSupport{"ExplanationOfBenefitPaymentComponent", false, true},

	}
}
