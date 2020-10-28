
 package models
 
 import (
	 "encoding/json"
	 "fmt"
	 "strings"
 )
 
 
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
 	case "benefitbalance":
		 return fhirVal.BenefitBalance, true
 	case "claimresponse":
		 return fhirVal.ClaimResponse, true
 	case "payee":
		 return fhirVal.Payee, true
 	case "careteam":
		 return fhirVal.CareTeam, true
 	case "procedure":
		 return fhirVal.Procedure, true
 	case "hospitalization":
		 return fhirVal.Hospitalization, true
 	case "item":
		 return fhirVal.Item, true
 	case "payment":
		 return fhirVal.Payment, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "outcome":
		 return fhirVal.Outcome, true
 	case "originalprescription":
		 return fhirVal.OriginalPrescription, true
 	case "employmentimpacted":
		 return fhirVal.EmploymentImpacted, true
 	case "totalcost":
		 return fhirVal.TotalCost, true
 	case "unallocdeductable":
		 return fhirVal.UnallocDeductable, true
 	case "subtype":
		 return fhirVal.SubType, true
 	case "facility":
		 return fhirVal.Facility, true
 	case "disposition":
		 return fhirVal.Disposition, true
 	case "additem":
		 return fhirVal.AddItem, true
 	case "insurance":
		 return fhirVal.Insurance, true
 	case "totalbenefit":
		 return fhirVal.TotalBenefit, true
 	case "provider":
		 return fhirVal.Provider, true
 	case "information":
		 return fhirVal.Information, true
 	case "type":
		 return fhirVal.Type, true
 	case "organization":
		 return fhirVal.Organization, true
 	case "referral":
		 return fhirVal.Referral, true
 	case "claim":
		 return fhirVal.Claim, true
 	case "form":
		 return fhirVal.Form, true
 	case "processnote":
		 return fhirVal.ProcessNote, true
 	case "status":
		 return fhirVal.Status, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "billableperiod":
		 return fhirVal.BillablePeriod, true
 	case "created":
		 return fhirVal.Created, true
 	case "enterer":
		 return fhirVal.Enterer, true
 	case "accident":
		 return fhirVal.Accident, true
 	case "insurer":
		 return fhirVal.Insurer, true
 	case "related":
		 return fhirVal.Related, true
 	case "prescription":
		 return fhirVal.Prescription, true
 	case "diagnosis":
		 return fhirVal.Diagnosis, true
 	case "precedence":
		 return fhirVal.Precedence, true
 
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
 		"benefitbalance": &FieldTypeSupport{"ExplanationOfBenefitBenefitBalanceComponent", true, false},
 		"claimresponse": &FieldTypeSupport{"Reference", false, true},
 		"payee": &FieldTypeSupport{"ExplanationOfBenefitPayeeComponent", false, true},
 		"careteam": &FieldTypeSupport{"ExplanationOfBenefitCareTeamComponent", true, false},
 		"procedure": &FieldTypeSupport{"ExplanationOfBenefitProcedureComponent", true, false},
 		"hospitalization": &FieldTypeSupport{"Period", false, true},
 		"item": &FieldTypeSupport{"ExplanationOfBenefitItemComponent", true, false},
 		"payment": &FieldTypeSupport{"ExplanationOfBenefitPaymentComponent", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"outcome": &FieldTypeSupport{"string", false, false},
 		"originalprescription": &FieldTypeSupport{"Reference", false, true},
 		"employmentimpacted": &FieldTypeSupport{"Period", false, true},
 		"totalcost": &FieldTypeSupport{"Quantity", false, true},
 		"unallocdeductable": &FieldTypeSupport{"Quantity", false, true},
 		"subtype": &FieldTypeSupport{"CodeableConcept", true, false},
 		"facility": &FieldTypeSupport{"Reference", false, true},
 		"disposition": &FieldTypeSupport{"string", false, false},
 		"additem": &FieldTypeSupport{"ExplanationOfBenefitAddedItemComponent", true, false},
 		"insurance": &FieldTypeSupport{"ExplanationOfBenefitInsuranceComponent", true, false},
 		"totalbenefit": &FieldTypeSupport{"Quantity", false, true},
 		"provider": &FieldTypeSupport{"Reference", false, true},
 		"information": &FieldTypeSupport{"ExplanationOfBenefitSupportingInformationComponent", true, false},
 		"type": &FieldTypeSupport{"CodeableConcept", false, true},
 		"organization": &FieldTypeSupport{"Reference", false, true},
 		"referral": &FieldTypeSupport{"Reference", false, true},
 		"claim": &FieldTypeSupport{"Reference", false, true},
 		"form": &FieldTypeSupport{"CodeableConcept", false, true},
 		"processnote": &FieldTypeSupport{"ExplanationOfBenefitNoteComponent", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"billableperiod": &FieldTypeSupport{"Period", false, true},
 		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"enterer": &FieldTypeSupport{"Reference", false, true},
 		"accident": &FieldTypeSupport{"ExplanationOfBenefitAccidentComponent", false, true},
 		"insurer": &FieldTypeSupport{"Reference", false, true},
 		"related": &FieldTypeSupport{"ExplanationOfBenefitRelatedClaimComponent", true, false},
 		"prescription": &FieldTypeSupport{"Reference", false, true},
 		"diagnosis": &FieldTypeSupport{"ExplanationOfBenefitDiagnosisComponent", true, false},
 		"precedence": &FieldTypeSupport{"uint32", false, true},
 
	 }
 }
 
 func (fhirVal *ExplanationOfBenefit) Truncate(fieldToKeep string) (TupleSupport, error) {
	 if strings.ToLower(fieldToKeep) == "benefitbalance" {
		 fieldToKeep = "{.benefitBalance}"
	 } else if !strings.HasPrefix(fieldToKeep, "{.benefitBalance[") {
		 return nil, fmt.Errorf("Not yet supported")
	 }
	 results, err := ApplyFieldName(fieldToKeep, fhirVal)
	 if err != nil {
		 return nil, err
	 }
	 structPtr := &ExplanationOfBenefit {
		 DomainResource: DomainResource{
			 Resource: Resource{
				 ResourceType: fhirVal.DomainResource.Resource.ResourceType,
				 Id: fhirVal.DomainResource.Resource.Id,
			 },
		 },
	 }
	 benefitBal := []ExplanationOfBenefitBenefitBalanceComponent{}
	 for _, benefit := range results {
		 var asBytes []byte
		 asBennies := []ExplanationOfBenefitBenefitBalanceComponent{}
		 switch asType := benefit.(type) {
		 case map[string]interface{}:
			 if asBytes, err = json.Marshal(asType); err != nil {
				 return nil, err
			 } else {
				 asBennies = append(asBennies, ExplanationOfBenefitBenefitBalanceComponent{})
				 if err = json.Unmarshal(asBytes, &asBennies[0]); err != nil {
					 return nil, err
				 }
			 }
		 case []interface{}:
			 if asBytes, err = json.Marshal(asType); err != nil {
				 return nil, err
			 } else if err = json.Unmarshal(asBytes, &asBennies); err != nil {
				 return nil, err
			 }
		 default:
			 return nil, fmt.Errorf("invalid type encountered")
		 }
		 benefitBal = append(benefitBal, asBennies...)
	 }
	 structPtr.BenefitBalance = benefitBal
	 return structPtr, nil
 }
 