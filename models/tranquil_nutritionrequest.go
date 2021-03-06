
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *NutritionRequest) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *NutritionRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "patient":
		 return fhirVal.Patient, true
 	case "datetime":
		 return fhirVal.DateTime, true
 	case "orderer":
		 return fhirVal.Orderer, true
 	case "foodpreferencemodifier":
		 return fhirVal.FoodPreferenceModifier, true
 	case "enteralformula":
		 return fhirVal.EnteralFormula, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "encounter":
		 return fhirVal.Encounter, true
 	case "allergyintolerance":
		 return fhirVal.AllergyIntolerance, true
 	case "excludefoodmodifier":
		 return fhirVal.ExcludeFoodModifier, true
 	case "oraldiet":
		 return fhirVal.OralDiet, true
 	case "supplement":
		 return fhirVal.Supplement, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *NutritionRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"datetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"orderer": &FieldTypeSupport{"Reference", false, true},
 		"foodpreferencemodifier": &FieldTypeSupport{"CodeableConcept", true, false},
 		"enteralformula": &FieldTypeSupport{"NutritionRequestEnteralFormulaComponent", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"encounter": &FieldTypeSupport{"Reference", false, true},
 		"allergyintolerance": &FieldTypeSupport{"Reference", true, false},
 		"excludefoodmodifier": &FieldTypeSupport{"CodeableConcept", true, false},
 		"oraldiet": &FieldTypeSupport{"NutritionRequestOralDietComponent", false, true},
 		"supplement": &FieldTypeSupport{"NutritionRequestSupplementComponent", true, false},
 
	 }
 }
 
 func (fhirVal *NutritionRequest) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 