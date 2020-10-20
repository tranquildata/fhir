
package models


func (fhirVal *Appointment) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Appointment) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *Appointment) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"slot": &FieldTypeSupport{"Reference", true, false},
		"appointmenttype": &FieldTypeSupport{"CodeableConcept", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", true, false},
		"indication": &FieldTypeSupport{"Reference", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"comment": &FieldTypeSupport{"string", false, false},
		"participant": &FieldTypeSupport{"AppointmentParticipantComponent", true, false},
		"incomingreferral": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"servicecategory": &FieldTypeSupport{"CodeableConcept", false, true},
		"servicetype": &FieldTypeSupport{"CodeableConcept", true, false},
		"priority": &FieldTypeSupport{"uint32", false, true},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"start": &FieldTypeSupport{"FHIRDateTime", false, true},
		"requestedperiod": &FieldTypeSupport{"Period", true, false},
		"end": &FieldTypeSupport{"FHIRDateTime", false, true},
		"minutesduration": &FieldTypeSupport{"uint32", false, true},

	}
}
