
package models


func (fhirVal *ImagingStudy) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ImagingStudy) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *ImagingStudy) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"uid": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"started": &FieldTypeSupport{"FHIRDateTime", false, true},
		"numberofseries": &FieldTypeSupport{"uint32", false, true},
		"procedurereference": &FieldTypeSupport{"Reference", true, false},
		"reason": &FieldTypeSupport{"CodeableConcept", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"accession": &FieldTypeSupport{"Identifier", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"availability": &FieldTypeSupport{"string", false, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"interpreter": &FieldTypeSupport{"Reference", true, false},
		"numberofinstances": &FieldTypeSupport{"uint32", false, true},
		"procedurecode": &FieldTypeSupport{"CodeableConcept", true, false},
		"modalitylist": &FieldTypeSupport{"Coding", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"referrer": &FieldTypeSupport{"Reference", false, true},
		"endpoint": &FieldTypeSupport{"Reference", true, false},
		"series": &FieldTypeSupport{"ImagingStudySeriesComponent", true, false},

	}
}
