
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ImmunizationRecommendation) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ImmunizationRecommendation) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "patient":
		 return fhirVal.Patient, true
 	case "recommendation":
		 return fhirVal.Recommendation, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ImmunizationRecommendation) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"recommendation": &FieldTypeSupport{"ImmunizationRecommendationRecommendationComponent", true, false},
 
	 }
 }
 
 func (fhirVal *ImmunizationRecommendation) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 