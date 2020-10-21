
package models


func (fhirVal *DeviceUseRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *DeviceUseRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "basedon":
		return fhirVal.BasedOn, true
	case "stage":
		return fhirVal.Stage, true
	case "context":
		return fhirVal.Context, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "requester":
		return fhirVal.Requester, true
	case "replaces":
		return fhirVal.Replaces, true
	case "requisition":
		return fhirVal.Requisition, true
	case "devicereference":
		return fhirVal.DeviceReference, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "performertype":
		return fhirVal.PerformerType, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "devicecodeableconcept":
		return fhirVal.DeviceCodeableConcept, true
	case "authored":
		return fhirVal.Authored, true
	case "performer":
		return fhirVal.Performer, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "definition":
		return fhirVal.Definition, true
	case "subject":
		return fhirVal.Subject, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true

	default:
		return nil, false
	}
}

func (fhirVal *DeviceUseRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"stage": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"requester": &FieldTypeSupport{"Reference", false, true},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"requisition": &FieldTypeSupport{"Identifier", false, true},
		"devicereference": &FieldTypeSupport{"Reference", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", false, true},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"devicecodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"authored": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}
