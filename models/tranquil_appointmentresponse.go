
package models


func (fhirVal *AppointmentResponse) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *AppointmentResponse) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "identifier":
		return fhirVal.Identifier, true
	case "appointment":
		return fhirVal.Appointment, true
	case "start":
		return fhirVal.Start, true
	case "end":
		return fhirVal.End, true
	case "participanttype":
		return fhirVal.ParticipantType, true
	case "actor":
		return fhirVal.Actor, true
	case "participantstatus":
		return fhirVal.ParticipantStatus, true
	case "comment":
		return fhirVal.Comment, true

	default:
		return nil, false
	}
}

func (fhirVal *AppointmentResponse) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"appointment": &FieldTypeSupport{"Reference", false, true},
		"start": &FieldTypeSupport{"FHIRDateTime", false, true},
		"end": &FieldTypeSupport{"FHIRDateTime", false, true},
		"participanttype": &FieldTypeSupport{"CodeableConcept", true, false},
		"actor": &FieldTypeSupport{"Reference", false, true},
		"participantstatus": &FieldTypeSupport{"string", false, false},
		"comment": &FieldTypeSupport{"string", false, false},

	}
}
