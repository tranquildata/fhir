
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
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"valuequantity": &FieldTypeSupport{"Quantity", false, true},
		"valueratio": &FieldTypeSupport{"Ratio", false, true},
		"valueattachment": &FieldTypeSupport{"Attachment", false, true},
		"valuedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"related": &FieldTypeSupport{"ObservationRelatedComponent", true, false},
		"component": &FieldTypeSupport{"ObservationComponentComponent", true, false},
		"performer": &FieldTypeSupport{"Reference", true, false},
		"valuerange": &FieldTypeSupport{"Range", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"valuestring": &FieldTypeSupport{"string", false, false},
		"comment": &FieldTypeSupport{"string", false, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"valueperiod": &FieldTypeSupport{"Period", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"interpretation": &FieldTypeSupport{"CodeableConcept", false, true},
		"method": &FieldTypeSupport{"CodeableConcept", false, true},
		"specimen": &FieldTypeSupport{"Reference", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"valuecodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"valueboolean": &FieldTypeSupport{"bool", false, true},
		"valuesampleddata": &FieldTypeSupport{"SampledData", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", false, true},
		"device": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"valuetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"dataabsentreason": &FieldTypeSupport{"CodeableConcept", false, true},
		"referencerange": &FieldTypeSupport{"ObservationReferenceRangeComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},

	}
}
