
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
		"subtype": &FieldTypeSupport{"CodeableConcept", false, true},
		"operator": &FieldTypeSupport{"Reference", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"device": &FieldTypeSupport{"Reference", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"type": &FieldTypeSupport{"string", false, false},
		"occurrenceperiod": &FieldTypeSupport{"Period", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"height": &FieldTypeSupport{"uint32", false, true},
		"width": &FieldTypeSupport{"uint32", false, true},
		"frames": &FieldTypeSupport{"uint32", false, true},
		"content": &FieldTypeSupport{"Attachment", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", false, true},
		"duration": &FieldTypeSupport{"uint32", false, true},
		"view": &FieldTypeSupport{"CodeableConcept", false, true},
		"occurrencedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
