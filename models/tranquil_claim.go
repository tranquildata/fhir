
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
	case "originalprescription":
		return fhirVal.OriginalPrescription, true
	case "insurance":
		return fhirVal.Insurance, true
	case "patient":
		return fhirVal.Patient, true
	case "billableperiod":
		return fhirVal.BillablePeriod, true
	case "created":
		return fhirVal.Created, true
	case "referral":
		return fhirVal.Referral, true
	case "diagnosis":
		return fhirVal.Diagnosis, true
	case "identifier":
		return fhirVal.Identifier, true
	case "subtype":
		return fhirVal.SubType, true
	case "enterer":
		return fhirVal.Enterer, true
	case "provider":
		return fhirVal.Provider, true
	case "payee":
		return fhirVal.Payee, true
	case "hospitalization":
		return fhirVal.Hospitalization, true
	case "status":
		return fhirVal.Status, true
	case "fundsreserve":
		return fhirVal.FundsReserve, true
	case "information":
		return fhirVal.Information, true
	case "insurer":
		return fhirVal.Insurer, true
	case "related":
		return fhirVal.Related, true
	case "prescription":
		return fhirVal.Prescription, true
	case "procedure":
		return fhirVal.Procedure, true
	case "careteam":
		return fhirVal.CareTeam, true
	case "accident":
		return fhirVal.Accident, true
	case "employmentimpacted":
		return fhirVal.EmploymentImpacted, true
	case "item":
		return fhirVal.Item, true
	case "type":
		return fhirVal.Type, true
	case "use":
		return fhirVal.Use, true
	case "organization":
		return fhirVal.Organization, true
	case "priority":
		return fhirVal.Priority, true
	case "facility":
		return fhirVal.Facility, true
	case "total":
		return fhirVal.Total, true

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
		"originalprescription": &FieldTypeSupport{"Reference", false, true},
		"insurance": &FieldTypeSupport{"ClaimInsuranceComponent", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"billableperiod": &FieldTypeSupport{"Period", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"referral": &FieldTypeSupport{"Reference", false, true},
		"diagnosis": &FieldTypeSupport{"ClaimDiagnosisComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"provider": &FieldTypeSupport{"Reference", false, true},
		"payee": &FieldTypeSupport{"ClaimPayeeComponent", false, true},
		"hospitalization": &FieldTypeSupport{"Period", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"fundsreserve": &FieldTypeSupport{"CodeableConcept", false, true},
		"information": &FieldTypeSupport{"ClaimSpecialConditionComponent", true, false},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"related": &FieldTypeSupport{"ClaimRelatedClaimComponent", true, false},
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"procedure": &FieldTypeSupport{"ClaimProcedureComponent", true, false},
		"careteam": &FieldTypeSupport{"ClaimCareTeamComponent", true, false},
		"accident": &FieldTypeSupport{"ClaimAccidentComponent", false, true},
		"employmentimpacted": &FieldTypeSupport{"Period", false, true},
		"item": &FieldTypeSupport{"ClaimItemComponent", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"use": &FieldTypeSupport{"string", false, false},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"priority": &FieldTypeSupport{"CodeableConcept", false, true},
		"facility": &FieldTypeSupport{"Reference", false, true},
		"total": &FieldTypeSupport{"Quantity", false, true},

	}
}
