
package models


func (fhirVal *Procedure) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Procedure) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "performer":
		return fhirVal.Performer, true
	case "status":
		return fhirVal.Status, true
	case "context":
		return fhirVal.Context, true
	case "note":
		return fhirVal.Note, true
	case "notdone":
		return fhirVal.NotDone, true
	case "notdonereason":
		return fhirVal.NotDoneReason, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "location":
		return fhirVal.Location, true
	case "followup":
		return fhirVal.FollowUp, true
	case "usedreference":
		return fhirVal.UsedReference, true
	case "report":
		return fhirVal.Report, true
	case "complication":
		return fhirVal.Complication, true
	case "partof":
		return fhirVal.PartOf, true
	case "code":
		return fhirVal.Code, true
	case "performeddatetime":
		return fhirVal.PerformedDateTime, true
	case "outcome":
		return fhirVal.Outcome, true
	case "complicationdetail":
		return fhirVal.ComplicationDetail, true
	case "identifier":
		return fhirVal.Identifier, true
	case "definition":
		return fhirVal.Definition, true
	case "performedperiod":
		return fhirVal.PerformedPeriod, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "category":
		return fhirVal.Category, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "focaldevice":
		return fhirVal.FocalDevice, true
	case "usedcode":
		return fhirVal.UsedCode, true

	default:
		return nil, false
	}
}

func (fhirVal *Procedure) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"performer": &FieldTypeSupport{"ProcedurePerformerComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"followup": &FieldTypeSupport{"CodeableConcept", true, false},
		"usedreference": &FieldTypeSupport{"Reference", true, false},
		"report": &FieldTypeSupport{"Reference", true, false},
		"complication": &FieldTypeSupport{"CodeableConcept", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"performeddatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"complicationdetail": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"performedperiod": &FieldTypeSupport{"Period", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"focaldevice": &FieldTypeSupport{"ProcedureFocalDeviceComponent", true, false},
		"usedcode": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
