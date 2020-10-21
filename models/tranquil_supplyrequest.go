
package models


func (fhirVal *SupplyRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *SupplyRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "requester":
		return fhirVal.Requester, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "identifier":
		return fhirVal.Identifier, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "reasoncodeableconcept":
		return fhirVal.ReasonCodeableConcept, true
	case "deliverfrom":
		return fhirVal.DeliverFrom, true
	case "category":
		return fhirVal.Category, true
	case "priority":
		return fhirVal.Priority, true
	case "ordereditem":
		return fhirVal.OrderedItem, true
	case "status":
		return fhirVal.Status, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "supplier":
		return fhirVal.Supplier, true
	case "deliverto":
		return fhirVal.DeliverTo, true

	default:
		return nil, false
	}
}

func (fhirVal *SupplyRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"requester": &FieldTypeSupport{"SupplyRequestRequesterComponent", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"reasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"deliverfrom": &FieldTypeSupport{"Reference", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"priority": &FieldTypeSupport{"string", false, false},
		"ordereditem": &FieldTypeSupport{"SupplyRequestOrderedItemComponent", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"supplier": &FieldTypeSupport{"Reference", true, false},
		"deliverto": &FieldTypeSupport{"Reference", false, true},

	}
}
