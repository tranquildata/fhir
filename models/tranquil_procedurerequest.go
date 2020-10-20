
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
	case "asneededcodeableconcept":
		return fhirVal.AsNeededCodeableConcept, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "category":
		return fhirVal.Category, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "note":
		return fhirVal.Note, true
	case "status":
		return fhirVal.Status, true
	case "code":
		return fhirVal.Code, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "definition":
		return fhirVal.Definition, true
	case "subject":
		return fhirVal.Subject, true
	case "donotperform":
		return fhirVal.DoNotPerform, true
	case "context":
		return fhirVal.Context, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "performertype":
		return fhirVal.PerformerType, true
	case "performer":
		return fhirVal.Performer, true
	case "requisition":
		return fhirVal.Requisition, true
	case "intent":
		return fhirVal.Intent, true
	case "specimen":
		return fhirVal.Specimen, true
	case "asneededboolean":
		return fhirVal.AsNeededBoolean, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "replaces":
		return fhirVal.Replaces, true
	case "priority":
		return fhirVal.Priority, true
	case "requester":
		return fhirVal.Requester, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "identifier":
		return fhirVal.Identifier, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true

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
		"asneededcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"donotperform": &FieldTypeSupport{"bool", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", false, true},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"requisition": &FieldTypeSupport{"Identifier", false, true},
		"intent": &FieldTypeSupport{"string", false, false},
		"specimen": &FieldTypeSupport{"Reference", true, false},
		"asneededboolean": &FieldTypeSupport{"bool", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"priority": &FieldTypeSupport{"string", false, false},
		"requester": &FieldTypeSupport{"ProcedureRequestRequesterComponent", false, true},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
