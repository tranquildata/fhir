
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
	case "referral":
		return fhirVal.Referral, true
	case "information":
		return fhirVal.Information, true
	case "unallocdeductable":
		return fhirVal.UnallocDeductable, true
	case "organization":
		return fhirVal.Organization, true
	case "related":
		return fhirVal.Related, true
	case "careteam":
		return fhirVal.CareTeam, true
	case "procedure":
		return fhirVal.Procedure, true
	case "totalcost":
		return fhirVal.TotalCost, true
	case "totalbenefit":
		return fhirVal.TotalBenefit, true
	case "billableperiod":
		return fhirVal.BillablePeriod, true
	case "outcome":
		return fhirVal.Outcome, true
	case "diagnosis":
		return fhirVal.Diagnosis, true
	case "identifier":
		return fhirVal.Identifier, true
	case "payee":
		return fhirVal.Payee, true
	case "created":
		return fhirVal.Created, true
	case "patient":
		return fhirVal.Patient, true
	case "provider":
		return fhirVal.Provider, true
	case "originalprescription":
		return fhirVal.OriginalPrescription, true
	case "insurance":
		return fhirVal.Insurance, true
	case "benefitbalance":
		return fhirVal.BenefitBalance, true
	case "type":
		return fhirVal.Type, true
	case "facility":
		return fhirVal.Facility, true
	case "item":
		return fhirVal.Item, true
	case "form":
		return fhirVal.Form, true
	case "processnote":
		return fhirVal.ProcessNote, true
	case "subtype":
		return fhirVal.SubType, true
	case "payment":
		return fhirVal.Payment, true
	case "status":
		return fhirVal.Status, true
	case "insurer":
		return fhirVal.Insurer, true
	case "claim":
		return fhirVal.Claim, true
	case "claimresponse":
		return fhirVal.ClaimResponse, true
	case "disposition":
		return fhirVal.Disposition, true
	case "prescription":
		return fhirVal.Prescription, true
	case "precedence":
		return fhirVal.Precedence, true
	case "accident":
		return fhirVal.Accident, true
	case "enterer":
		return fhirVal.Enterer, true
	case "hospitalization":
		return fhirVal.Hospitalization, true
	case "additem":
		return fhirVal.AddItem, true
	case "employmentimpacted":
		return fhirVal.EmploymentImpacted, true

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
		"referral": &FieldTypeSupport{"Reference", false, true},
		"information": &FieldTypeSupport{"ExplanationOfBenefitSupportingInformationComponent", true, false},
		"unallocdeductable": &FieldTypeSupport{"Quantity", false, true},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"related": &FieldTypeSupport{"ExplanationOfBenefitRelatedClaimComponent", true, false},
		"careteam": &FieldTypeSupport{"ExplanationOfBenefitCareTeamComponent", true, false},
		"procedure": &FieldTypeSupport{"ExplanationOfBenefitProcedureComponent", true, false},
		"totalcost": &FieldTypeSupport{"Quantity", false, true},
		"totalbenefit": &FieldTypeSupport{"Quantity", false, true},
		"billableperiod": &FieldTypeSupport{"Period", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"diagnosis": &FieldTypeSupport{"ExplanationOfBenefitDiagnosisComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"payee": &FieldTypeSupport{"ExplanationOfBenefitPayeeComponent", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"provider": &FieldTypeSupport{"Reference", false, true},
		"originalprescription": &FieldTypeSupport{"Reference", false, true},
		"insurance": &FieldTypeSupport{"ExplanationOfBenefitInsuranceComponent", false, true},
		"benefitbalance": &FieldTypeSupport{"ExplanationOfBenefitBenefitBalanceComponent", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"facility": &FieldTypeSupport{"Reference", false, true},
		"item": &FieldTypeSupport{"ExplanationOfBenefitItemComponent", true, false},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},
		"processnote": &FieldTypeSupport{"ExplanationOfBenefitNoteComponent", true, false},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"payment": &FieldTypeSupport{"ExplanationOfBenefitPaymentComponent", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"claim": &FieldTypeSupport{"Reference", false, true},
		"claimresponse": &FieldTypeSupport{"Reference", false, true},
		"disposition": &FieldTypeSupport{"string", false, false},
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"precedence": &FieldTypeSupport{"uint32", false, true},
		"accident": &FieldTypeSupport{"ExplanationOfBenefitAccidentComponent", false, true},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"hospitalization": &FieldTypeSupport{"Period", false, true},
		"additem": &FieldTypeSupport{"ExplanationOfBenefitAddedItemComponent", true, false},
		"employmentimpacted": &FieldTypeSupport{"Period", false, true},

	}
}
