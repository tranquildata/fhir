
package models


func (fhirVal *ChargeItem) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ChargeItem) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "performingorganization":
		return fhirVal.PerformingOrganization, true
	case "requestingorganization":
		return fhirVal.RequestingOrganization, true
	case "priceoverride":
		return fhirVal.PriceOverride, true
	case "enterer":
		return fhirVal.Enterer, true
	case "service":
		return fhirVal.Service, true
	case "code":
		return fhirVal.Code, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "quantity":
		return fhirVal.Quantity, true
	case "bodysite":
		return fhirVal.Bodysite, true
	case "overridereason":
		return fhirVal.OverrideReason, true
	case "reason":
		return fhirVal.Reason, true
	case "note":
		return fhirVal.Note, true
	case "identifier":
		return fhirVal.Identifier, true
	case "context":
		return fhirVal.Context, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "participant":
		return fhirVal.Participant, true
	case "supportinginformation":
		return fhirVal.SupportingInformation, true
	case "definition":
		return fhirVal.Definition, true
	case "status":
		return fhirVal.Status, true
	case "factoroverride":
		return fhirVal.FactorOverride, true
	case "entereddate":
		return fhirVal.EnteredDate, true
	case "account":
		return fhirVal.Account, true
	case "partof":
		return fhirVal.PartOf, true
	case "subject":
		return fhirVal.Subject, true

	default:
		return nil, false
	}
}

func (fhirVal *ChargeItem) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"performingorganization": &FieldTypeSupport{"Reference", false, true},
		"requestingorganization": &FieldTypeSupport{"Reference", false, true},
		"priceoverride": &FieldTypeSupport{"Quantity", false, true},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"service": &FieldTypeSupport{"Reference", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"overridereason": &FieldTypeSupport{"string", false, false},
		"reason": &FieldTypeSupport{"CodeableConcept", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"participant": &FieldTypeSupport{"ChargeItemParticipantComponent", true, false},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"string", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"factoroverride": &FieldTypeSupport{"float64", false, true},
		"entereddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"account": &FieldTypeSupport{"Reference", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},

	}
}
