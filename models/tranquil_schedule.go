
package models


func (fhirVal *Schedule) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Schedule) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "identifier":
		return fhirVal.Identifier, true
	case "active":
		return fhirVal.Active, true
	case "servicecategory":
		return fhirVal.ServiceCategory, true
	case "servicetype":
		return fhirVal.ServiceType, true
	case "specialty":
		return fhirVal.Specialty, true
	case "actor":
		return fhirVal.Actor, true
	case "planninghorizon":
		return fhirVal.PlanningHorizon, true
	case "comment":
		return fhirVal.Comment, true

	default:
		return nil, false
	}
}

func (fhirVal *Schedule) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"active": &FieldTypeSupport{"bool", false, true},
		"servicecategory": &FieldTypeSupport{"CodeableConcept", false, true},
		"servicetype": &FieldTypeSupport{"CodeableConcept", true, false},
		"specialty": &FieldTypeSupport{"CodeableConcept", true, false},
		"actor": &FieldTypeSupport{"Reference", true, false},
		"planninghorizon": &FieldTypeSupport{"Period", false, true},
		"comment": &FieldTypeSupport{"string", false, false},

	}
}
