
package models


func (fhirVal *Binary) TupleName() (id string, resourceType string) {
	return fhirVal.Resource.Id, fhirVal.Resource.ResourceType
}

func (fhirVal *Binary) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "contenttype":
		return fhirVal.ContentType, true
	case "securitycontext":
		return fhirVal.SecurityContext, true
	case "content":
		return fhirVal.Content, true

	default:
		return nil, false
	}
}

func (fhirVal *Binary) FieldsToTypes() map[string]*FieldTypeSupport {
	return map[string]*FieldTypeSupport {

			"resourcetype": &FieldTypeSupport{"string", false, false},
		"id": &FieldTypeSupport{"string", false, false},
		"meta": &FieldTypeSupport{"Meta", false, true},
		"implicitrules": &FieldTypeSupport{"string", false, false},
		"language": &FieldTypeSupport{"string", false, false},
		"contenttype": &FieldTypeSupport{"string", false, false},
		"securitycontext": &FieldTypeSupport{"Reference", false, true},
		"content": &FieldTypeSupport{"string", false, false},

	}
}
