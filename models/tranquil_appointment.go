
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
	case "servicetype":
		return fhirVal.ServiceType, true
	case "incomingreferral":
		return fhirVal.IncomingReferral, true
	case "identifier":
		return fhirVal.Identifier, true
	case "specialty":
		return fhirVal.Specialty, true
	case "indication":
		return fhirVal.Indication, true
	case "description":
		return fhirVal.Description, true
	case "supportinginformation":
		return fhirVal.SupportingInformation, true
	case "minutesduration":
		return fhirVal.MinutesDuration, true
	case "participant":
		return fhirVal.Participant, true
	case "appointmenttype":
		return fhirVal.AppointmentType, true
	case "reason":
		return fhirVal.Reason, true
	case "end":
		return fhirVal.End, true
	case "slot":
		return fhirVal.Slot, true
	case "created":
		return fhirVal.Created, true
	case "comment":
		return fhirVal.Comment, true
	case "status":
		return fhirVal.Status, true
	case "servicecategory":
		return fhirVal.ServiceCategory, true
	case "priority":
		return fhirVal.Priority, true
	case "start":
		return fhirVal.Start, true
	case "requestedperiod":
		return fhirVal.RequestedPeriod, true

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
		"servicetype": &FieldTypeSupport{"CodeableConcept", true, false},
		"incomingreferral": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"specialty": &FieldTypeSupport{"CodeableConcept", true, false},
		"indication": &FieldTypeSupport{"Reference", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"minutesduration": &FieldTypeSupport{"uint32", false, true},
		"participant": &FieldTypeSupport{"AppointmentParticipantComponent", true, false},
		"appointmenttype": &FieldTypeSupport{"CodeableConcept", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", true, false},
		"end": &FieldTypeSupport{"FHIRDateTime", false, true},
		"slot": &FieldTypeSupport{"Reference", true, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"comment": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"servicecategory": &FieldTypeSupport{"CodeableConcept", false, true},
		"priority": &FieldTypeSupport{"uint32", false, true},
		"start": &FieldTypeSupport{"FHIRDateTime", false, true},
		"requestedperiod": &FieldTypeSupport{"Period", true, false},

	}
}
