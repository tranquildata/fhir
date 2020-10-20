
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
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"enterer": &FieldTypeSupport{"Reference", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"performingorganization": &FieldTypeSupport{"Reference", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"entereddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"priceoverride": &FieldTypeSupport{"Quantity", false, true},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"string", true, false},
		"participant": &FieldTypeSupport{"ChargeItemParticipantComponent", true, false},
		"requestingorganization": &FieldTypeSupport{"Reference", false, true},
		"factoroverride": &FieldTypeSupport{"float64", false, true},
		"overridereason": &FieldTypeSupport{"string", false, false},
		"service": &FieldTypeSupport{"Reference", true, false},
		"account": &FieldTypeSupport{"Reference", true, false},

	}
}
