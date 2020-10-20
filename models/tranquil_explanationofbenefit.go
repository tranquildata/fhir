
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
	case "billableperiod":
		return fhirVal.BillablePeriod, true
	case "careteam":
		return fhirVal.CareTeam, true
	case "patient":
		return fhirVal.Patient, true
	case "insurance":
		return fhirVal.Insurance, true
	case "payment":
		return fhirVal.Payment, true
	case "precedence":
		return fhirVal.Precedence, true
	case "diagnosis":
		return fhirVal.Diagnosis, true
	case "unallocdeductable":
		return fhirVal.UnallocDeductable, true
	case "enterer":
		return fhirVal.Enterer, true
	case "totalbenefit":
		return fhirVal.TotalBenefit, true
	case "benefitbalance":
		return fhirVal.BenefitBalance, true
	case "claimresponse":
		return fhirVal.ClaimResponse, true
	case "outcome":
		return fhirVal.Outcome, true
	case "procedure":
		return fhirVal.Procedure, true
	case "processnote":
		return fhirVal.ProcessNote, true
	case "created":
		return fhirVal.Created, true
	case "facility":
		return fhirVal.Facility, true
	case "claim":
		return fhirVal.Claim, true
	case "payee":
		return fhirVal.Payee, true
	case "accident":
		return fhirVal.Accident, true
	case "item":
		return fhirVal.Item, true
	case "status":
		return fhirVal.Status, true
	case "subtype":
		return fhirVal.SubType, true
	case "insurer":
		return fhirVal.Insurer, true
	case "provider":
		return fhirVal.Provider, true
	case "related":
		return fhirVal.Related, true
	case "prescription":
		return fhirVal.Prescription, true
	case "originalprescription":
		return fhirVal.OriginalPrescription, true
	case "totalcost":
		return fhirVal.TotalCost, true
	case "type":
		return fhirVal.Type, true
	case "organization":
		return fhirVal.Organization, true
	case "referral":
		return fhirVal.Referral, true
	case "disposition":
		return fhirVal.Disposition, true
	case "information":
		return fhirVal.Information, true
	case "employmentimpacted":
		return fhirVal.EmploymentImpacted, true
	case "hospitalization":
		return fhirVal.Hospitalization, true
	case "additem":
		return fhirVal.AddItem, true
	case "identifier":
		return fhirVal.Identifier, true
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
		"billableperiod": &FieldTypeSupport{"Period", false, true},
		"careteam": &FieldTypeSupport{"ExplanationOfBenefitCareTeamComponent", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"insurance": &FieldTypeSupport{"ExplanationOfBenefitInsuranceComponent", false, true},
		"payment": &FieldTypeSupport{"ExplanationOfBenefitPaymentComponent", false, true},
		"precedence": &FieldTypeSupport{"uint32", false, true},
		"diagnosis": &FieldTypeSupport{"ExplanationOfBenefitDiagnosisComponent", true, false},
		"unallocdeductable": &FieldTypeSupport{"Quantity", false, true},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"totalbenefit": &FieldTypeSupport{"Quantity", false, true},
		"benefitbalance": &FieldTypeSupport{"ExplanationOfBenefitBenefitBalanceComponent", true, false},
		"claimresponse": &FieldTypeSupport{"Reference", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"procedure": &FieldTypeSupport{"ExplanationOfBenefitProcedureComponent", true, false},
		"processnote": &FieldTypeSupport{"ExplanationOfBenefitNoteComponent", true, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"facility": &FieldTypeSupport{"Reference", false, true},
		"claim": &FieldTypeSupport{"Reference", false, true},
		"payee": &FieldTypeSupport{"ExplanationOfBenefitPayeeComponent", false, true},
		"accident": &FieldTypeSupport{"ExplanationOfBenefitAccidentComponent", false, true},
		"item": &FieldTypeSupport{"ExplanationOfBenefitItemComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"provider": &FieldTypeSupport{"Reference", false, true},
		"related": &FieldTypeSupport{"ExplanationOfBenefitRelatedClaimComponent", true, false},
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"originalprescription": &FieldTypeSupport{"Reference", false, true},
		"totalcost": &FieldTypeSupport{"Quantity", false, true},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"referral": &FieldTypeSupport{"Reference", false, true},
		"disposition": &FieldTypeSupport{"string", false, false},
		"information": &FieldTypeSupport{"ExplanationOfBenefitSupportingInformationComponent", true, false},
		"employmentimpacted": &FieldTypeSupport{"Period", false, true},
		"hospitalization": &FieldTypeSupport{"Period", false, true},
		"additem": &FieldTypeSupport{"ExplanationOfBenefitAddedItemComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
