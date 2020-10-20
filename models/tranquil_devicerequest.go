
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
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "priorrequest":
		return fhirVal.PriorRequest, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "codecodeableconcept":
		return fhirVal.CodeCodeableConcept, true
	case "subject":
		return fhirVal.Subject, true
	case "performer":
		return fhirVal.Performer, true
	case "requester":
		return fhirVal.Requester, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "note":
		return fhirVal.Note, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "priority":
		return fhirVal.Priority, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "definition":
		return fhirVal.Definition, true
	case "status":
		return fhirVal.Status, true
	case "context":
		return fhirVal.Context, true
	case "occurrencetiming":
		return fhirVal.OccurrenceTiming, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "identifier":
		return fhirVal.Identifier, true
	case "intent":
		return fhirVal.Intent, true
	case "codereference":
		return fhirVal.CodeReference, true
	case "performertype":
		return fhirVal.PerformerType, true

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
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"priorrequest": &FieldTypeSupport{"Reference", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"codecodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"requester": &FieldTypeSupport{"DeviceRequestRequesterComponent", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"priority": &FieldTypeSupport{"string", false, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"occurrencetiming": &FieldTypeSupport{"Timing", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"intent": &FieldTypeSupport{"CodeableConcept", false, true},
		"codereference": &FieldTypeSupport{"Reference", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
