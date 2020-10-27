
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *FamilyMemberHistory) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *FamilyMemberHistory) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "status":
		 return fhirVal.Status, true
 	case "notdone":
		 return fhirVal.NotDone, true
 	case "relationship":
		 return fhirVal.Relationship, true
 	case "deceasedage":
		 return fhirVal.DeceasedAge, true
 	case "deceasedrange":
		 return fhirVal.DeceasedRange, true
 	case "note":
		 return fhirVal.Note, true
 	case "condition":
		 return fhirVal.Condition, true
 	case "date":
		 return fhirVal.Date, true
 	case "deceasedboolean":
		 return fhirVal.DeceasedBoolean, true
 	case "deceaseddate":
		 return fhirVal.DeceasedDate, true
 	case "agestring":
		 return fhirVal.AgeString, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "definition":
		 return fhirVal.Definition, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "bornperiod":
		 return fhirVal.BornPeriod, true
 	case "borndate":
		 return fhirVal.BornDate, true
 	case "ageage":
		 return fhirVal.AgeAge, true
 	case "agerange":
		 return fhirVal.AgeRange, true
 	case "estimatedage":
		 return fhirVal.EstimatedAge, true
 	case "reasoncode":
		 return fhirVal.ReasonCode, true
 	case "reasonreference":
		 return fhirVal.ReasonReference, true
 	case "notdonereason":
		 return fhirVal.NotDoneReason, true
 	case "name":
		 return fhirVal.Name, true
 	case "gender":
		 return fhirVal.Gender, true
 	case "bornstring":
		 return fhirVal.BornString, true
 	case "deceasedstring":
		 return fhirVal.DeceasedString, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *FamilyMemberHistory) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"status": &FieldTypeSupport{"string", false, false},
 		"notdone": &FieldTypeSupport{"bool", false, true},
 		"relationship": &FieldTypeSupport{"CodeableConcept", false, true},
 		"deceasedage": &FieldTypeSupport{"Quantity", false, true},
 		"deceasedrange": &FieldTypeSupport{"Range", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"condition": &FieldTypeSupport{"FamilyMemberHistoryConditionComponent", true, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"deceasedboolean": &FieldTypeSupport{"bool", false, true},
 		"deceaseddate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"agestring": &FieldTypeSupport{"string", false, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"definition": &FieldTypeSupport{"Reference", true, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"bornperiod": &FieldTypeSupport{"Period", false, true},
 		"borndate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"ageage": &FieldTypeSupport{"Quantity", false, true},
 		"agerange": &FieldTypeSupport{"Range", false, true},
 		"estimatedage": &FieldTypeSupport{"bool", false, true},
 		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
 		"reasonreference": &FieldTypeSupport{"Reference", true, false},
 		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
 		"name": &FieldTypeSupport{"string", false, false},
 		"gender": &FieldTypeSupport{"string", false, false},
 		"bornstring": &FieldTypeSupport{"string", false, false},
 		"deceasedstring": &FieldTypeSupport{"string", false, false},
 
	 }
 }
 
 func (fhirVal *FamilyMemberHistory) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 