
package models


func (fhirVal *Practitioner) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Practitioner) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "photo":
		return fhirVal.Photo, true
	case "communication":
		return fhirVal.Communication, true
	case "identifier":
		return fhirVal.Identifier, true
	case "telecom":
		return fhirVal.Telecom, true
	case "address":
		return fhirVal.Address, true
	case "birthdate":
		return fhirVal.BirthDate, true
	case "active":
		return fhirVal.Active, true
	case "name":
		return fhirVal.Name, true
	case "gender":
		return fhirVal.Gender, true
	case "qualification":
		return fhirVal.Qualification, true

	default:
		return nil, false
	}
}

func (fhirVal *Practitioner) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"photo": &FieldTypeSupport{"Attachment", true, false},
		"communication": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"address": &FieldTypeSupport{"Address", true, false},
		"birthdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"active": &FieldTypeSupport{"bool", false, true},
		"name": &FieldTypeSupport{"HumanName", true, false},
		"gender": &FieldTypeSupport{"string", false, false},
		"qualification": &FieldTypeSupport{"PractitionerQualificationComponent", true, false},

	}
}
