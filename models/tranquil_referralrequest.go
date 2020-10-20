
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
	case "replaces":
		return fhirVal.Replaces, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "priority":
		return fhirVal.Priority, true
	case "requester":
		return fhirVal.Requester, true
	case "recipient":
		return fhirVal.Recipient, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "definition":
		return fhirVal.Definition, true
	case "subject":
		return fhirVal.Subject, true
	case "context":
		return fhirVal.Context, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "description":
		return fhirVal.Description, true
	case "supportinginfo":
		return fhirVal.SupportingInfo, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "type":
		return fhirVal.Type, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "note":
		return fhirVal.Note, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "intent":
		return fhirVal.Intent, true
	case "servicerequested":
		return fhirVal.ServiceRequested, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "specialty":
		return fhirVal.Specialty, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true

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
		"replaces": &FieldTypeSupport{"Reference", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"priority": &FieldTypeSupport{"string", false, false},
		"requester": &FieldTypeSupport{"ReferralRequestRequesterComponent", false, true},
		"recipient": &FieldTypeSupport{"Reference", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"supportinginfo": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"intent": &FieldTypeSupport{"string", false, false},
		"servicerequested": &FieldTypeSupport{"CodeableConcept", true, false},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"specialty": &FieldTypeSupport{"CodeableConcept", false, true},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},

	}
}
