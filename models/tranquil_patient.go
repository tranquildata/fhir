
package models


func (fhirVal *Patient) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Patient) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *Patient) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"gender": &FieldTypeSupport{"string", false, false},
		"deceasedboolean": &FieldTypeSupport{"bool", false, true},
		"deceaseddatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"multiplebirthinteger": &FieldTypeSupport{"int32", false, true},
		"communication": &FieldTypeSupport{"PatientCommunicationComponent", true, false},
		"link": &FieldTypeSupport{"PatientLinkComponent", true, false},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"birthdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"maritalstatus": &FieldTypeSupport{"CodeableConcept", false, true},
		"photo": &FieldTypeSupport{"Attachment", true, false},
		"generalpractitioner": &FieldTypeSupport{"Reference", true, false},
		"active": &FieldTypeSupport{"bool", false, true},
		"address": &FieldTypeSupport{"Address", true, false},
		"contact": &FieldTypeSupport{"PatientContactComponent", true, false},
		"managingorganization": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"multiplebirthboolean": &FieldTypeSupport{"bool", false, true},
		"animal": &FieldTypeSupport{"PatientAnimalComponent", false, true},
		"name": &FieldTypeSupport{"HumanName", true, false},

	}
}
