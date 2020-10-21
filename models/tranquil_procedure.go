
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
	case "category":
		return fhirVal.Category, true
	case "context":
		return fhirVal.Context, true
	case "complicationdetail":
		return fhirVal.ComplicationDetail, true
	case "followup":
		return fhirVal.FollowUp, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "notdonereason":
		return fhirVal.NotDoneReason, true
	case "note":
		return fhirVal.Note, true
	case "location":
		return fhirVal.Location, true
	case "focaldevice":
		return fhirVal.FocalDevice, true
	case "partof":
		return fhirVal.PartOf, true
	case "performer":
		return fhirVal.Performer, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "outcome":
		return fhirVal.Outcome, true
	case "usedcode":
		return fhirVal.UsedCode, true
	case "definition":
		return fhirVal.Definition, true
	case "status":
		return fhirVal.Status, true
	case "performeddatetime":
		return fhirVal.PerformedDateTime, true
	case "performedperiod":
		return fhirVal.PerformedPeriod, true
	case "report":
		return fhirVal.Report, true
	case "complication":
		return fhirVal.Complication, true
	case "usedreference":
		return fhirVal.UsedReference, true
	case "code":
		return fhirVal.Code, true
	case "subject":
		return fhirVal.Subject, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "identifier":
		return fhirVal.Identifier, true
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
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"complicationdetail": &FieldTypeSupport{"Reference", true, false},
		"followup": &FieldTypeSupport{"CodeableConcept", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"focaldevice": &FieldTypeSupport{"ProcedureFocalDeviceComponent", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"performer": &FieldTypeSupport{"ProcedurePerformerComponent", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"usedcode": &FieldTypeSupport{"CodeableConcept", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"performeddatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performedperiod": &FieldTypeSupport{"Period", false, true},
		"report": &FieldTypeSupport{"Reference", true, false},
		"complication": &FieldTypeSupport{"CodeableConcept", true, false},
		"usedreference": &FieldTypeSupport{"Reference", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},

	}
}
