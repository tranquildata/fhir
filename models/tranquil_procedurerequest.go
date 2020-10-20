
package models


func (fhirVal *ProcedureRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ProcedureRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "donotperform":
		return fhirVal.DoNotPerform, true
	case "category":
		return fhirVal.Category, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "status":
		return fhirVal.Status, true
	case "context":
		return fhirVal.Context, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "specimen":
		return fhirVal.Specimen, true
	case "requisition":
		return fhirVal.Requisition, true
	case "code":
		return fhirVal.Code, true
	case "performer":
		return fhirVal.Performer, true
	case "priority":
		return fhirVal.Priority, true
	case "subject":
		return fhirVal.Subject, true
	case "asneededboolean":
		return fhirVal.AsNeededBoolean, true
	case "asneededcodeableconcept":
		return fhirVal.AsNeededCodeableConcept, true
	case "intent":
		return fhirVal.Intent, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "requester":
		return fhirVal.Requester, true
	case "performertype":
		return fhirVal.PerformerType, true
	case "note":
		return fhirVal.Note, true
	case "identifier":
		return fhirVal.Identifier, true
	case "definition":
		return fhirVal.Definition, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "replaces":
		return fhirVal.Replaces, true

	default:
		return nil, false
	}
}

func (fhirVal *ProcedureRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"donotperform": &FieldTypeSupport{"bool", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"specimen": &FieldTypeSupport{"Reference", true, false},
		"requisition": &FieldTypeSupport{"Identifier", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"priority": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"asneededboolean": &FieldTypeSupport{"bool", false, true},
		"asneededcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"intent": &FieldTypeSupport{"string", false, false},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"requester": &FieldTypeSupport{"ProcedureRequestRequesterComponent", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},

	}
}
