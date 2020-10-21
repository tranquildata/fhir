
package models


func (fhirVal *NutritionOrder) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *NutritionOrder) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "supplement":
		return fhirVal.Supplement, true
	case "enteralformula":
		return fhirVal.EnteralFormula, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "orderer":
		return fhirVal.Orderer, true
	case "excludefoodmodifier":
		return fhirVal.ExcludeFoodModifier, true
	case "foodpreferencemodifier":
		return fhirVal.FoodPreferenceModifier, true
	case "oraldiet":
		return fhirVal.OralDiet, true
	case "patient":
		return fhirVal.Patient, true
	case "encounter":
		return fhirVal.Encounter, true
	case "datetime":
		return fhirVal.DateTime, true
	case "allergyintolerance":
		return fhirVal.AllergyIntolerance, true

	default:
		return nil, false
	}
}

func (fhirVal *NutritionOrder) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"supplement": &FieldTypeSupport{"NutritionOrderSupplementComponent", true, false},
		"enteralformula": &FieldTypeSupport{"NutritionOrderEnteralFormulaComponent", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"orderer": &FieldTypeSupport{"Reference", false, true},
		"excludefoodmodifier": &FieldTypeSupport{"CodeableConcept", true, false},
		"foodpreferencemodifier": &FieldTypeSupport{"CodeableConcept", true, false},
		"oraldiet": &FieldTypeSupport{"NutritionOrderOralDietComponent", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"encounter": &FieldTypeSupport{"Reference", false, true},
		"datetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"allergyintolerance": &FieldTypeSupport{"Reference", true, false},

	}
}
