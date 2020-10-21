
package models


func (fhirVal *Organization) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Organization) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "active":
		return fhirVal.Active, true
	case "type":
		return fhirVal.Type, true
	case "alias":
		return fhirVal.Alias, true
	case "address":
		return fhirVal.Address, true
	case "partof":
		return fhirVal.PartOf, true
	case "contact":
		return fhirVal.Contact, true
	case "identifier":
		return fhirVal.Identifier, true
	case "name":
		return fhirVal.Name, true
	case "telecom":
		return fhirVal.Telecom, true
	case "endpoint":
		return fhirVal.Endpoint, true

	default:
		return nil, false
	}
}

func (fhirVal *Organization) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"active": &FieldTypeSupport{"bool", false, true},
		"type": &FieldTypeSupport{"CodeableConcept", true, false},
		"alias": &FieldTypeSupport{"string", true, false},
		"address": &FieldTypeSupport{"Address", true, false},
		"partof": &FieldTypeSupport{"Reference", false, true},
		"contact": &FieldTypeSupport{"OrganizationContactComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"endpoint": &FieldTypeSupport{"Reference", true, false},

	}
}
