
package models


func (fhirVal *Observation) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Observation) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "valuequantity":
		return fhirVal.ValueQuantity, true
	case "valuecodeableconcept":
		return fhirVal.ValueCodeableConcept, true
	case "valueboolean":
		return fhirVal.ValueBoolean, true
	case "component":
		return fhirVal.Component, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "issued":
		return fhirVal.Issued, true
	case "performer":
		return fhirVal.Performer, true
	case "specimen":
		return fhirVal.Specimen, true
	case "valuestring":
		return fhirVal.ValueString, true
	case "valueattachment":
		return fhirVal.ValueAttachment, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "valuesampleddata":
		return fhirVal.ValueSampledData, true
	case "valueperiod":
		return fhirVal.ValuePeriod, true
	case "comment":
		return fhirVal.Comment, true
	case "context":
		return fhirVal.Context, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true
	case "valueratio":
		return fhirVal.ValueRatio, true
	case "subject":
		return fhirVal.Subject, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "method":
		return fhirVal.Method, true
	case "dataabsentreason":
		return fhirVal.DataAbsentReason, true
	case "code":
		return fhirVal.Code, true
	case "valuerange":
		return fhirVal.ValueRange, true
	case "referencerange":
		return fhirVal.ReferenceRange, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "valuetime":
		return fhirVal.ValueTime, true
	case "device":
		return fhirVal.Device, true
	case "related":
		return fhirVal.Related, true
	case "category":
		return fhirVal.Category, true
	case "valuedatetime":
		return fhirVal.ValueDateTime, true
	case "interpretation":
		return fhirVal.Interpretation, true

	default:
		return nil, false
	}
}

func (fhirVal *Observation) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"valuequantity": &FieldTypeSupport{"Quantity", false, true},
		"valuecodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"valueboolean": &FieldTypeSupport{"bool", false, true},
		"component": &FieldTypeSupport{"ObservationComponentComponent", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performer": &FieldTypeSupport{"Reference", true, false},
		"specimen": &FieldTypeSupport{"Reference", false, true},
		"valuestring": &FieldTypeSupport{"string", false, false},
		"valueattachment": &FieldTypeSupport{"Attachment", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", false, true},
		"valuesampleddata": &FieldTypeSupport{"SampledData", false, true},
		"valueperiod": &FieldTypeSupport{"Period", false, true},
		"comment": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"valueratio": &FieldTypeSupport{"Ratio", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"method": &FieldTypeSupport{"CodeableConcept", false, true},
		"dataabsentreason": &FieldTypeSupport{"CodeableConcept", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"valuerange": &FieldTypeSupport{"Range", false, true},
		"referencerange": &FieldTypeSupport{"ObservationReferenceRangeComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"valuetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"device": &FieldTypeSupport{"Reference", false, true},
		"related": &FieldTypeSupport{"ObservationRelatedComponent", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"valuedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"interpretation": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
