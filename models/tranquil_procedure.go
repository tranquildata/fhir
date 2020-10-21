
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
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "report":
		return fhirVal.Report, true
	case "followup":
		return fhirVal.FollowUp, true
	case "location":
		return fhirVal.Location, true
	case "complicationdetail":
		return fhirVal.ComplicationDetail, true
	case "performedperiod":
		return fhirVal.PerformedPeriod, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "outcome":
		return fhirVal.Outcome, true
	case "context":
		return fhirVal.Context, true
	case "usedreference":
		return fhirVal.UsedReference, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "subject":
		return fhirVal.Subject, true
	case "code":
		return fhirVal.Code, true
	case "performer":
		return fhirVal.Performer, true
	case "note":
		return fhirVal.Note, true
	case "focaldevice":
		return fhirVal.FocalDevice, true
	case "usedcode":
		return fhirVal.UsedCode, true
	case "status":
		return fhirVal.Status, true
	case "notdonereason":
		return fhirVal.NotDoneReason, true
	case "category":
		return fhirVal.Category, true
	case "partof":
		return fhirVal.PartOf, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "performeddatetime":
		return fhirVal.PerformedDateTime, true
	case "complication":
		return fhirVal.Complication, true
	case "identifier":
		return fhirVal.Identifier, true
	case "definition":
		return fhirVal.Definition, true
	case "notdone":
		return fhirVal.NotDone, true

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
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"report": &FieldTypeSupport{"Reference", true, false},
		"followup": &FieldTypeSupport{"CodeableConcept", true, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"complicationdetail": &FieldTypeSupport{"Reference", true, false},
		"performedperiod": &FieldTypeSupport{"Period", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"usedreference": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"performer": &FieldTypeSupport{"ProcedurePerformerComponent", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"focaldevice": &FieldTypeSupport{"ProcedureFocalDeviceComponent", true, false},
		"usedcode": &FieldTypeSupport{"CodeableConcept", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"performeddatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"complication": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},

	}
}
