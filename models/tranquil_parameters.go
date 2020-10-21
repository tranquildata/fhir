
package models


func (fhirVal *Parameters) TupleName() (id string, resourceType string) {
	return fhirVal.Resource.Id, fhirVal.Resource.ResourceType
}

func (fhirVal *Parameters) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "parameter":
		return fhirVal.Parameter, true

	default:
		return nil, false
	}
}

func (fhirVal *Parameters) FieldsToTypes() map[string]*FieldTypeSupport {
	return map[string]*FieldTypeSupport {

			"resourcetype": &FieldTypeSupport{"string", false, false},
		"id": &FieldTypeSupport{"string", false, false},
		"meta": &FieldTypeSupport{"Meta", false, true},
		"implicitrules": &FieldTypeSupport{"string", false, false},
		"language": &FieldTypeSupport{"string", false, false},
		"parameter": &FieldTypeSupport{"ParametersParameterComponent", true, false},

	}
}
