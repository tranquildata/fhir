
package models


func (fhirVal *Media) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Media) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "type":
		return fhirVal.Type, true
	case "subject":
		return fhirVal.Subject, true
	case "operator":
		return fhirVal.Operator, true
	case "duration":
		return fhirVal.Duration, true
	case "content":
		return fhirVal.Content, true
	case "occurrenceperiod":
		return fhirVal.OccurrencePeriod, true
	case "height":
		return fhirVal.Height, true
	case "note":
		return fhirVal.Note, true
	case "identifier":
		return fhirVal.Identifier, true
	case "view":
		return fhirVal.View, true
	case "context":
		return fhirVal.Context, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "width":
		return fhirVal.Width, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "subtype":
		return fhirVal.Subtype, true
	case "occurrencedatetime":
		return fhirVal.OccurrenceDateTime, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "device":
		return fhirVal.Device, true
	case "frames":
		return fhirVal.Frames, true

	default:
		return nil, false
	}
}

func (fhirVal *Media) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"type": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"operator": &FieldTypeSupport{"Reference", false, true},
		"duration": &FieldTypeSupport{"uint32", false, true},
		"content": &FieldTypeSupport{"Attachment", false, true},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"height": &FieldTypeSupport{"uint32", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"view": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"width": &FieldTypeSupport{"uint32", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"subtype": &FieldTypeSupport{"CodeableConcept", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", false, true},
		"device": &FieldTypeSupport{"Reference", false, true},
		"frames": &FieldTypeSupport{"uint32", false, true},

	}
}
