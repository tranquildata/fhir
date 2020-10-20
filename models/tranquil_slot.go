
package models


func (fhirVal *Slot) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Slot) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "specialty":
		return fhirVal.Specialty, true
	case "status":
		return fhirVal.Status, true
	case "overbooked":
		return fhirVal.Overbooked, true
	case "identifier":
		return fhirVal.Identifier, true
	case "servicecategory":
		return fhirVal.ServiceCategory, true
	case "servicetype":
		return fhirVal.ServiceType, true
	case "appointmenttype":
		return fhirVal.AppointmentType, true
	case "schedule":
		return fhirVal.Schedule, true
	case "start":
		return fhirVal.Start, true
	case "end":
		return fhirVal.End, true
	case "comment":
		return fhirVal.Comment, true

	default:
		return nil, false
	}
}

func (fhirVal *Slot) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"specialty": &FieldTypeSupport{"CodeableConcept", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"overbooked": &FieldTypeSupport{"bool", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"servicecategory": &FieldTypeSupport{"CodeableConcept", false, true},
		"servicetype": &FieldTypeSupport{"CodeableConcept", true, false},
		"appointmenttype": &FieldTypeSupport{"CodeableConcept", false, true},
		"schedule": &FieldTypeSupport{"Reference", false, true},
		"start": &FieldTypeSupport{"FHIRDateTime", false, true},
		"end": &FieldTypeSupport{"FHIRDateTime", false, true},
		"comment": &FieldTypeSupport{"string", false, false},

	}
}
