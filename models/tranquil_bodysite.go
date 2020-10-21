
package models


func (fhirVal *BodySite) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *BodySite) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "code":
		return fhirVal.Code, true
	case "qualifier":
		return fhirVal.Qualifier, true
	case "description":
		return fhirVal.Description, true
	case "image":
		return fhirVal.Image, true
	case "patient":
		return fhirVal.Patient, true
	case "identifier":
		return fhirVal.Identifier, true
	case "active":
		return fhirVal.Active, true

	default:
		return nil, false
	}
}

func (fhirVal *BodySite) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"qualifier": &FieldTypeSupport{"CodeableConcept", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"image": &FieldTypeSupport{"Attachment", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"active": &FieldTypeSupport{"bool", false, true},

	}
}
