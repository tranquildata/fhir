
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
	case "type":
		return fhirVal.Type, true
	case "insurer":
		return fhirVal.Insurer, true
	case "outcome":
		return fhirVal.Outcome, true
	case "related":
		return fhirVal.Related, true
	case "unallocdeductable":
		return fhirVal.UnallocDeductable, true
	case "payment":
		return fhirVal.Payment, true
	case "status":
		return fhirVal.Status, true
	case "claimresponse":
		return fhirVal.ClaimResponse, true
	case "disposition":
		return fhirVal.Disposition, true
	case "payee":
		return fhirVal.Payee, true
	case "careteam":
		return fhirVal.CareTeam, true
	case "totalbenefit":
		return fhirVal.TotalBenefit, true
	case "enterer":
		return fhirVal.Enterer, true
	case "facility":
		return fhirVal.Facility, true
	case "employmentimpacted":
		return fhirVal.EmploymentImpacted, true
	case "additem":
		return fhirVal.AddItem, true
	case "patient":
		return fhirVal.Patient, true
	case "organization":
		return fhirVal.Organization, true
	case "precedence":
		return fhirVal.Precedence, true
	case "benefitbalance":
		return fhirVal.BenefitBalance, true
	case "identifier":
		return fhirVal.Identifier, true
	case "information":
		return fhirVal.Information, true
	case "insurance":
		return fhirVal.Insurance, true
	case "hospitalization":
		return fhirVal.Hospitalization, true
	case "subtype":
		return fhirVal.SubType, true
	case "claim":
		return fhirVal.Claim, true
	case "prescription":
		return fhirVal.Prescription, true
	case "accident":
		return fhirVal.Accident, true
	case "billableperiod":
		return fhirVal.BillablePeriod, true
	case "created":
		return fhirVal.Created, true
	case "referral":
		return fhirVal.Referral, true
	case "originalprescription":
		return fhirVal.OriginalPrescription, true
	case "diagnosis":
		return fhirVal.Diagnosis, true
	case "procedure":
		return fhirVal.Procedure, true
	case "item":
		return fhirVal.Item, true
	case "totalcost":
		return fhirVal.TotalCost, true
	case "processnote":
		return fhirVal.ProcessNote, true
	case "provider":
		return fhirVal.Provider, true
	case "form":
		return fhirVal.Form, true

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
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"outcome": &FieldTypeSupport{"string", false, false},
		"related": &FieldTypeSupport{"ExplanationOfBenefitRelatedClaimComponent", true, false},
		"unallocdeductable": &FieldTypeSupport{"Quantity", false, true},
		"payment": &FieldTypeSupport{"ExplanationOfBenefitPaymentComponent", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"claimresponse": &FieldTypeSupport{"Reference", false, true},
		"disposition": &FieldTypeSupport{"string", false, false},
		"payee": &FieldTypeSupport{"ExplanationOfBenefitPayeeComponent", false, true},
		"careteam": &FieldTypeSupport{"ExplanationOfBenefitCareTeamComponent", true, false},
		"totalbenefit": &FieldTypeSupport{"Quantity", false, true},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"facility": &FieldTypeSupport{"Reference", false, true},
		"employmentimpacted": &FieldTypeSupport{"Period", false, true},
		"additem": &FieldTypeSupport{"ExplanationOfBenefitAddedItemComponent", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"precedence": &FieldTypeSupport{"uint32", false, true},
		"benefitbalance": &FieldTypeSupport{"ExplanationOfBenefitBenefitBalanceComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"information": &FieldTypeSupport{"ExplanationOfBenefitSupportingInformationComponent", true, false},
		"insurance": &FieldTypeSupport{"ExplanationOfBenefitInsuranceComponent", true, false},
		"hospitalization": &FieldTypeSupport{"Period", false, true},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"claim": &FieldTypeSupport{"Reference", false, true},
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"accident": &FieldTypeSupport{"ExplanationOfBenefitAccidentComponent", false, true},
		"billableperiod": &FieldTypeSupport{"Period", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"referral": &FieldTypeSupport{"Reference", false, true},
		"originalprescription": &FieldTypeSupport{"Reference", false, true},
		"diagnosis": &FieldTypeSupport{"ExplanationOfBenefitDiagnosisComponent", true, false},
		"procedure": &FieldTypeSupport{"ExplanationOfBenefitProcedureComponent", true, false},
		"item": &FieldTypeSupport{"ExplanationOfBenefitItemComponent", true, false},
		"totalcost": &FieldTypeSupport{"Quantity", false, true},
		"processnote": &FieldTypeSupport{"ExplanationOfBenefitNoteComponent", true, false},
		"provider": &FieldTypeSupport{"Reference", false, true},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
