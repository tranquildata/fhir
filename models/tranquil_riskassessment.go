
package models


func (fhirVal *RiskAssessment) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *RiskAssessment) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "context":
		return fhirVal.Context, true
	case "condition":
		return fhirVal.Condition, true
	case "mitigation":
		return fhirVal.Mitigation, true
	case "performer":
		return fhirVal.Performer, true
	case "comment":
		return fhirVal.Comment, true
	case "identifier":
		return fhirVal.Identifier, true
	case "parent":
		return fhirVal.Parent, true
	case "subject":
		return fhirVal.Subject, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "status":
		return fhirVal.Status, true
	case "code":
		return fhirVal.Code, true
	case "reasoncodeableconcept":
		return fhirVal.ReasonCodeableConcept, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "basis":
		return fhirVal.Basis, true
	case "prediction":
		return fhirVal.Prediction, true
	case "method":
		return fhirVal.Method, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true

	default:
		return nil, false
	}
}

func (fhirVal *RiskAssessment) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"context": &FieldTypeSupport{"Reference", false, true},
		"condition": &FieldTypeSupport{"Reference", false, true},
		"mitigation": &FieldTypeSupport{"string", false, false},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"comment": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"parent": &FieldTypeSupport{"Reference", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"basedon": &FieldTypeSupport{"Reference", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", false, true},
		"basis": &FieldTypeSupport{"Reference", true, false},
		"prediction": &FieldTypeSupport{"RiskAssessmentPredictionComponent", true, false},
		"method": &FieldTypeSupport{"CodeableConcept", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},

	}
}
