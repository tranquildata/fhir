
package models


func (fhirVal *DeviceComponent) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *DeviceComponent) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "type":
		return fhirVal.Type, true
	case "parent":
		return fhirVal.Parent, true
	case "operationalstatus":
		return fhirVal.OperationalStatus, true
	case "productionspecification":
		return fhirVal.ProductionSpecification, true
	case "identifier":
		return fhirVal.Identifier, true
	case "source":
		return fhirVal.Source, true
	case "parametergroup":
		return fhirVal.ParameterGroup, true
	case "measurementprinciple":
		return fhirVal.MeasurementPrinciple, true
	case "languagecode":
		return fhirVal.LanguageCode, true
	case "lastsystemchange":
		return fhirVal.LastSystemChange, true

	default:
		return nil, false
	}
}

func (fhirVal *DeviceComponent) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"parent": &FieldTypeSupport{"Reference", false, true},
		"operationalstatus": &FieldTypeSupport{"CodeableConcept", true, false},
		"productionspecification": &FieldTypeSupport{"DeviceComponentProductionSpecificationComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"source": &FieldTypeSupport{"Reference", false, true},
		"parametergroup": &FieldTypeSupport{"CodeableConcept", false, true},
		"measurementprinciple": &FieldTypeSupport{"string", false, false},
		"languagecode": &FieldTypeSupport{"CodeableConcept", false, true},
		"lastsystemchange": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
