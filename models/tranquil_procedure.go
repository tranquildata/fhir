
package models


func (fhirVal *Procedure) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Procedure) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *Procedure) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"performer": &FieldTypeSupport{"ProcedurePerformerComponent", true, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"complication": &FieldTypeSupport{"CodeableConcept", true, false},
		"usedcode": &FieldTypeSupport{"CodeableConcept", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"complicationdetail": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"report": &FieldTypeSupport{"Reference", true, false},
		"usedreference": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"performedperiod": &FieldTypeSupport{"Period", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"focaldevice": &FieldTypeSupport{"ProcedureFocalDeviceComponent", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"performeddatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"followup": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
