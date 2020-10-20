
package models


func (fhirVal *ResearchStudy) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ResearchStudy) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *ResearchStudy) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"focus": &FieldTypeSupport{"CodeableConcept", true, false},
		"keyword": &FieldTypeSupport{"CodeableConcept", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"sponsor": &FieldTypeSupport{"Reference", false, true},
		"title": &FieldTypeSupport{"string", false, false},
		"protocol": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"site": &FieldTypeSupport{"Reference", true, false},
		"reasonstopped": &FieldTypeSupport{"CodeableConcept", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"principalinvestigator": &FieldTypeSupport{"Reference", false, true},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"enrollment": &FieldTypeSupport{"Reference", true, false},
		"period": &FieldTypeSupport{"Period", false, true},
		"arm": &FieldTypeSupport{"ResearchStudyArmComponent", true, false},

	}
}
