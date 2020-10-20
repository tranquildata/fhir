
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
	case "deceaseddatetime":
		return fhirVal.DeceasedDateTime, true
	case "multiplebirthboolean":
		return fhirVal.MultipleBirthBoolean, true
	case "contact":
		return fhirVal.Contact, true
	case "managingorganization":
		return fhirVal.ManagingOrganization, true
	case "identifier":
		return fhirVal.Identifier, true
	case "multiplebirthinteger":
		return fhirVal.MultipleBirthInteger, true
	case "generalpractitioner":
		return fhirVal.GeneralPractitioner, true
	case "telecom":
		return fhirVal.Telecom, true
	case "maritalstatus":
		return fhirVal.MaritalStatus, true
	case "animal":
		return fhirVal.Animal, true
	case "communication":
		return fhirVal.Communication, true
	case "name":
		return fhirVal.Name, true
	case "gender":
		return fhirVal.Gender, true
	case "birthdate":
		return fhirVal.BirthDate, true
	case "deceasedboolean":
		return fhirVal.DeceasedBoolean, true
	case "address":
		return fhirVal.Address, true
	case "photo":
		return fhirVal.Photo, true
	case "link":
		return fhirVal.Link, true
	case "active":
		return fhirVal.Active, true

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
		"deceaseddatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"multiplebirthboolean": &FieldTypeSupport{"bool", false, true},
		"contact": &FieldTypeSupport{"PatientContactComponent", true, false},
		"managingorganization": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"multiplebirthinteger": &FieldTypeSupport{"int32", false, true},
		"generalpractitioner": &FieldTypeSupport{"Reference", true, false},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"maritalstatus": &FieldTypeSupport{"CodeableConcept", false, true},
		"animal": &FieldTypeSupport{"PatientAnimalComponent", false, true},
		"communication": &FieldTypeSupport{"PatientCommunicationComponent", true, false},
		"name": &FieldTypeSupport{"HumanName", true, false},
		"gender": &FieldTypeSupport{"string", false, false},
		"birthdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"deceasedboolean": &FieldTypeSupport{"bool", false, true},
		"address": &FieldTypeSupport{"Address", true, false},
		"photo": &FieldTypeSupport{"Attachment", true, false},
		"link": &FieldTypeSupport{"PatientLinkComponent", true, false},
		"active": &FieldTypeSupport{"bool", false, true},

	}
}
