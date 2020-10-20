
package models


func (fhirVal *Task) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Task) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "definitionuri":
		return fhirVal.DefinitionUri, true
	case "partof":
		return fhirVal.PartOf, true
	case "statusreason":
		return fhirVal.StatusReason, true
	case "focus":
		return fhirVal.Focus, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "executionperiod":
		return fhirVal.ExecutionPeriod, true
	case "requester":
		return fhirVal.Requester, true
	case "reason":
		return fhirVal.Reason, true
	case "output":
		return fhirVal.Output, true
	case "status":
		return fhirVal.Status, true
	case "businessstatus":
		return fhirVal.BusinessStatus, true
	case "description":
		return fhirVal.Description, true
	case "for":
		return fhirVal.For, true
	case "note":
		return fhirVal.Note, true
	case "restriction":
		return fhirVal.Restriction, true
	case "input":
		return fhirVal.Input, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "lastmodified":
		return fhirVal.LastModified, true
	case "performertype":
		return fhirVal.PerformerType, true
	case "context":
		return fhirVal.Context, true
	case "owner":
		return fhirVal.Owner, true
	case "definitionreference":
		return fhirVal.DefinitionReference, true
	case "intent":
		return fhirVal.Intent, true
	case "relevanthistory":
		return fhirVal.RelevantHistory, true
	case "identifier":
		return fhirVal.Identifier, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "priority":
		return fhirVal.Priority, true
	case "code":
		return fhirVal.Code, true

	default:
		return nil, false
	}
}

func (fhirVal *Task) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"definitionuri": &FieldTypeSupport{"string", false, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"statusreason": &FieldTypeSupport{"CodeableConcept", false, true},
		"focus": &FieldTypeSupport{"Reference", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"executionperiod": &FieldTypeSupport{"Period", false, true},
		"requester": &FieldTypeSupport{"TaskRequesterComponent", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", false, true},
		"output": &FieldTypeSupport{"TaskOutputComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"businessstatus": &FieldTypeSupport{"CodeableConcept", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"for": &FieldTypeSupport{"Reference", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"restriction": &FieldTypeSupport{"TaskRestrictionComponent", false, true},
		"input": &FieldTypeSupport{"TaskParameterComponent", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"lastmodified": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"owner": &FieldTypeSupport{"Reference", false, true},
		"definitionreference": &FieldTypeSupport{"Reference", false, true},
		"intent": &FieldTypeSupport{"string", false, false},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"priority": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
