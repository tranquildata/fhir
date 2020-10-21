
package models


func (fhirVal *EligibilityResponse) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *EligibilityResponse) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "insurer":
		return fhirVal.Insurer, true
	case "inforce":
		return fhirVal.Inforce, true
	case "insurance":
		return fhirVal.Insurance, true
	case "identifier":
		return fhirVal.Identifier, true
	case "created":
		return fhirVal.Created, true
	case "request":
		return fhirVal.Request, true
	case "disposition":
		return fhirVal.Disposition, true
	case "form":
		return fhirVal.Form, true
	case "error":
		return fhirVal.Error, true
	case "status":
		return fhirVal.Status, true
	case "requestprovider":
		return fhirVal.RequestProvider, true
	case "requestorganization":
		return fhirVal.RequestOrganization, true
	case "outcome":
		return fhirVal.Outcome, true

	default:
		return nil, false
	}
}

func (fhirVal *EligibilityResponse) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"insurer": &FieldTypeSupport{"Reference", false, true},
		"inforce": &FieldTypeSupport{"bool", false, true},
		"insurance": &FieldTypeSupport{"EligibilityResponseInsuranceComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"request": &FieldTypeSupport{"Reference", false, true},
		"disposition": &FieldTypeSupport{"string", false, false},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},
		"error": &FieldTypeSupport{"EligibilityResponseErrorsComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"requestprovider": &FieldTypeSupport{"Reference", false, true},
		"requestorganization": &FieldTypeSupport{"Reference", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
