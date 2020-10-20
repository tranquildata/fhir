
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
	case "note":
		return fhirVal.Note, true
	case "devicecodeableconcept":
		return fhirVal.DeviceCodeableConcept, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "authored":
		return fhirVal.Authored, true
	case "performertype":
		return fhirVal.PerformerType, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "definition":
		return fhirVal.Definition, true
	case "requisition":
		return fhirVal.Requisition, true
	case "subject":
		return fhirVal.Subject, true
	case "performer":
		return fhirVal.Performer, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "identifier":
		return fhirVal.Identifier, true
	case "replaces":
		return fhirVal.Replaces, true
	case "status":
		return fhirVal.Status, true
	case "stage":
		return fhirVal.Stage, true
	case "devicereference":
		return fhirVal.DeviceReference, true
	case "context":
		return fhirVal.Context, true
	case "requester":
		return fhirVal.Requester, true

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
		"note": &FieldTypeSupport{"Annotation", true, false},
		"devicecodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"authored": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"requisition": &FieldTypeSupport{"Identifier", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"stage": &FieldTypeSupport{"CodeableConcept", false, true},
		"devicereference": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"requester": &FieldTypeSupport{"Reference", false, true},

	}
}
