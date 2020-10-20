
package models


func (fhirVal *Goal) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Goal) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "status":
		return fhirVal.Status, true
	case "startcodeableconcept":
		return fhirVal.StartCodeableConcept, true
	case "target":
		return fhirVal.Target, true
	case "statusreason":
		return fhirVal.StatusReason, true
	case "outcomecode":
		return fhirVal.OutcomeCode, true
	case "identifier":
		return fhirVal.Identifier, true
	case "startdate":
		return fhirVal.StartDate, true
	case "priority":
		return fhirVal.Priority, true
	case "addresses":
		return fhirVal.Addresses, true
	case "category":
		return fhirVal.Category, true
	case "description":
		return fhirVal.Description, true
	case "subject":
		return fhirVal.Subject, true
	case "statusdate":
		return fhirVal.StatusDate, true
	case "expressedby":
		return fhirVal.ExpressedBy, true
	case "note":
		return fhirVal.Note, true
	case "outcomereference":
		return fhirVal.OutcomeReference, true

	default:
		return nil, false
	}
}

func (fhirVal *Goal) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"status": &FieldTypeSupport{"string", false, false},
		"startcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"target": &FieldTypeSupport{"GoalTargetComponent", false, true},
		"statusreason": &FieldTypeSupport{"string", false, false},
		"outcomecode": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"startdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"priority": &FieldTypeSupport{"CodeableConcept", false, true},
		"addresses": &FieldTypeSupport{"Reference", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"description": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"statusdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"expressedby": &FieldTypeSupport{"Reference", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"outcomereference": &FieldTypeSupport{"Reference", true, false},

	}
}
