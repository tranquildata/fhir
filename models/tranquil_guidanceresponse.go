
package models


func (fhirVal *GuidanceResponse) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *GuidanceResponse) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "subject":
		return fhirVal.Subject, true
	case "outputparameters":
		return fhirVal.OutputParameters, true
	case "requestid":
		return fhirVal.RequestId, true
	case "identifier":
		return fhirVal.Identifier, true
	case "module":
		return fhirVal.Module, true
	case "reasoncodeableconcept":
		return fhirVal.ReasonCodeableConcept, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true
	case "evaluationmessage":
		return fhirVal.EvaluationMessage, true
	case "status":
		return fhirVal.Status, true
	case "context":
		return fhirVal.Context, true
	case "result":
		return fhirVal.Result, true
	case "datarequirement":
		return fhirVal.DataRequirement, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "performer":
		return fhirVal.Performer, true

	default:
		return nil, false
	}
}

func (fhirVal *GuidanceResponse) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"subject": &FieldTypeSupport{"Reference", false, true},
		"outputparameters": &FieldTypeSupport{"Reference", false, true},
		"requestid": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"module": &FieldTypeSupport{"Reference", false, true},
		"reasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"evaluationmessage": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"result": &FieldTypeSupport{"Reference", false, true},
		"datarequirement": &FieldTypeSupport{"DataRequirement", true, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performer": &FieldTypeSupport{"Reference", false, true},

	}
}
