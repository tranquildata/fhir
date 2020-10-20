
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
		"priority": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"statusreason": &FieldTypeSupport{"CodeableConcept", false, true},
		"restriction": &FieldTypeSupport{"TaskRestrictionComponent", false, true},
		"input": &FieldTypeSupport{"TaskParameterComponent", true, false},
		"output": &FieldTypeSupport{"TaskOutputComponent", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"executionperiod": &FieldTypeSupport{"Period", false, true},
		"lastmodified": &FieldTypeSupport{"FHIRDateTime", false, true},
		"owner": &FieldTypeSupport{"Reference", false, true},
		"requester": &FieldTypeSupport{"TaskRequesterComponent", false, true},
		"performertype": &FieldTypeSupport{"CodeableConcept", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"relevanthistory": &FieldTypeSupport{"Reference", true, false},
		"definitionreference": &FieldTypeSupport{"Reference", false, true},
		"intent": &FieldTypeSupport{"string", false, false},
		"definitionuri": &FieldTypeSupport{"string", false, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"businessstatus": &FieldTypeSupport{"CodeableConcept", false, true},
		"focus": &FieldTypeSupport{"Reference", false, true},
		"for": &FieldTypeSupport{"Reference", false, true},

	}
}
