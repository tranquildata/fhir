
package models


func (fhirVal *Location) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Location) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "telecom":
		return fhirVal.Telecom, true
	case "address":
		return fhirVal.Address, true
	case "identifier":
		return fhirVal.Identifier, true
	case "type":
		return fhirVal.Type, true
	case "physicaltype":
		return fhirVal.PhysicalType, true
	case "managingorganization":
		return fhirVal.ManagingOrganization, true
	case "partof":
		return fhirVal.PartOf, true
	case "endpoint":
		return fhirVal.Endpoint, true
	case "name":
		return fhirVal.Name, true
	case "operationalstatus":
		return fhirVal.OperationalStatus, true
	case "alias":
		return fhirVal.Alias, true
	case "description":
		return fhirVal.Description, true
	case "mode":
		return fhirVal.Mode, true
	case "position":
		return fhirVal.Position, true
	case "status":
		return fhirVal.Status, true

	default:
		return nil, false
	}
}

func (fhirVal *Location) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
		"address": &FieldTypeSupport{"Address", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"physicaltype": &FieldTypeSupport{"CodeableConcept", false, true},
		"managingorganization": &FieldTypeSupport{"Reference", false, true},
		"partof": &FieldTypeSupport{"Reference", false, true},
		"endpoint": &FieldTypeSupport{"Reference", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"operationalstatus": &FieldTypeSupport{"Coding", false, true},
		"alias": &FieldTypeSupport{"string", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"mode": &FieldTypeSupport{"string", false, false},
		"position": &FieldTypeSupport{"LocationPositionComponent", false, true},
		"status": &FieldTypeSupport{"string", false, false},

	}
}
