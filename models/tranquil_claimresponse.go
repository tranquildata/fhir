
package models


func (fhirVal *ClaimResponse) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ClaimResponse) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "insurance":
		return fhirVal.Insurance, true
	case "disposition":
		return fhirVal.Disposition, true
	case "item":
		return fhirVal.Item, true
	case "additem":
		return fhirVal.AddItem, true
	case "error":
		return fhirVal.Error, true
	case "status":
		return fhirVal.Status, true
	case "created":
		return fhirVal.Created, true
	case "insurer":
		return fhirVal.Insurer, true
	case "request":
		return fhirVal.Request, true
	case "payment":
		return fhirVal.Payment, true
	case "reserved":
		return fhirVal.Reserved, true
	case "unallocdeductable":
		return fhirVal.UnallocDeductable, true
	case "form":
		return fhirVal.Form, true
	case "identifier":
		return fhirVal.Identifier, true
	case "requestprovider":
		return fhirVal.RequestProvider, true
	case "requestorganization":
		return fhirVal.RequestOrganization, true
	case "totalcost":
		return fhirVal.TotalCost, true
	case "processnote":
		return fhirVal.ProcessNote, true
	case "communicationrequest":
		return fhirVal.CommunicationRequest, true
	case "patient":
		return fhirVal.Patient, true
	case "outcome":
		return fhirVal.Outcome, true
	case "payeetype":
		return fhirVal.PayeeType, true
	case "totalbenefit":
		return fhirVal.TotalBenefit, true

	default:
		return nil, false
	}
}

func (fhirVal *ClaimResponse) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"insurance": &FieldTypeSupport{"ClaimResponseInsuranceComponent", true, false},
		"disposition": &FieldTypeSupport{"string", false, false},
		"item": &FieldTypeSupport{"ClaimResponseItemComponent", true, false},
		"additem": &FieldTypeSupport{"ClaimResponseAddedItemComponent", true, false},
		"error": &FieldTypeSupport{"ClaimResponseErrorComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"request": &FieldTypeSupport{"Reference", false, true},
		"payment": &FieldTypeSupport{"ClaimResponsePaymentComponent", false, true},
		"reserved": &FieldTypeSupport{"Coding", false, true},
		"unallocdeductable": &FieldTypeSupport{"Quantity", false, true},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"requestprovider": &FieldTypeSupport{"Reference", false, true},
		"requestorganization": &FieldTypeSupport{"Reference", false, true},
		"totalcost": &FieldTypeSupport{"Quantity", false, true},
		"processnote": &FieldTypeSupport{"ClaimResponseNoteComponent", true, false},
		"communicationrequest": &FieldTypeSupport{"Reference", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"payeetype": &FieldTypeSupport{"CodeableConcept", false, true},
		"totalbenefit": &FieldTypeSupport{"Quantity", false, true},

	}
}
