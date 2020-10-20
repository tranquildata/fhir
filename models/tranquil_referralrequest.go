
package models


func (fhirVal *ReferralRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ReferralRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "status":
		return fhirVal.Status, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "description":
		return fhirVal.Description, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "definition":
		return fhirVal.Definition, true
	case "subject":
		return fhirVal.Subject, true
	case "context":
		return fhirVal.Context, true
	case "recipient":
		return fhirVal.Recipient, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "note":
		return fhirVal.Note, true
	case "replaces":
		return fhirVal.Replaces, true
	case "priority":
		return fhirVal.Priority, true
	case "specialty":
		return fhirVal.Specialty, true
	case "type":
		return fhirVal.Type, true
	case "servicerequested":
		return fhirVal.ServiceRequested, true
	case "requester":
		return fhirVal.Requester, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "identifier":
		return fhirVal.Identifier, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "intent":
		return fhirVal.Intent, true

	default:
		return nil, false
	}
}

func (fhirVal *ReferralRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"recipient": &FieldTypeSupport{"Reference", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"priority": &FieldTypeSupport{"string", false, false},
		"specialty": &FieldTypeSupport{"CodeableConcept", false, true},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"servicerequested": &FieldTypeSupport{"CodeableConcept", true, false},
		"requester": &FieldTypeSupport{"ReferralRequestRequesterComponent", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"intent": &FieldTypeSupport{"string", false, false},

	}
}
