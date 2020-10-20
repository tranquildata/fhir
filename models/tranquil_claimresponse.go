
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
		"requestprovider": &FieldTypeSupport{"Reference", false, true},
		"error": &FieldTypeSupport{"ClaimResponseErrorComponent", true, false},
		"totalbenefit": &FieldTypeSupport{"Quantity", false, true},
		"reserved": &FieldTypeSupport{"Coding", false, true},
		"additem": &FieldTypeSupport{"ClaimResponseAddedItemComponent", true, false},
		"totalcost": &FieldTypeSupport{"Quantity", false, true},
		"payment": &FieldTypeSupport{"ClaimResponsePaymentComponent", false, true},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"disposition": &FieldTypeSupport{"string", false, false},
		"unallocdeductable": &FieldTypeSupport{"Quantity", false, true},
		"requestorganization": &FieldTypeSupport{"Reference", false, true},
		"request": &FieldTypeSupport{"Reference", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"item": &FieldTypeSupport{"ClaimResponseItemComponent", true, false},
		"communicationrequest": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"payeetype": &FieldTypeSupport{"CodeableConcept", false, true},
		"processnote": &FieldTypeSupport{"ClaimResponseNoteComponent", true, false},

	}
}
