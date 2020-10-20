
package models


func (fhirVal *SupplyDelivery) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *SupplyDelivery) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "identifier":
		return fhirVal.Identifier, true
	case "patient":
		return fhirVal.Patient, true
	case "supplieditem":
		return fhirVal.SuppliedItem, true
	case "supplier":
		return fhirVal.Supplier, true
	case "receiver":
		return fhirVal.Receiver, true
	case "destination":
		return fhirVal.Destination, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "partof":
		return fhirVal.PartOf, true
	case "status":
		return fhirVal.Status, true
	case "type":
		return fhirVal.Type, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true

	default:
		return nil, false
	}
}

func (fhirVal *SupplyDelivery) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"patient": &FieldTypeSupport{"Reference", false, true},
		"supplieditem": &FieldTypeSupport{"SupplyDeliverySuppliedItemComponent", false, true},
		"supplier": &FieldTypeSupport{"Reference", false, true},
		"receiver": &FieldTypeSupport{"Reference", true, false},
		"destination": &FieldTypeSupport{"Reference", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},

	}
}
