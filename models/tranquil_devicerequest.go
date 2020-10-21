
package models


func (fhirVal *DeviceRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *DeviceRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "priorrequest":
		return fhirVal.PriorRequest, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "status":
		return fhirVal.Status, true
	case "intent":
		return fhirVal.Intent, true
	case "priority":
		return fhirVal.Priority, true
	case "subject":
		return fhirVal.Subject, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "requester":
		return fhirVal.Requester, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "definition":
		return fhirVal.Definition, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "codereference":
		return fhirVal.CodeReference, true
	case "codecodeableconcept":
		return fhirVal.CodeCodeableConcept, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "identifier":
		return fhirVal.Identifier, true
	case "context":
		return fhirVal.Context, true
	case "performertype":
		return fhirVal.PerformerType, true
	case "performer":
		return fhirVal.Performer, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "note":
		return fhirVal.Note, true

	default:
		return nil, false
	}
}

func (fhirVal *DeviceRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"priorrequest": &FieldTypeSupport{"Reference", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"intent": &FieldTypeSupport{"CodeableConcept", false, true},
		"priority": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"requester": &FieldTypeSupport{"DeviceRequestRequesterComponent", false, true},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"codereference": &FieldTypeSupport{"Reference", false, true},
		"codecodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", false, true},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}
