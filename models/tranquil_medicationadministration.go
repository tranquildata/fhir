
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *MedicationAdministration) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *MedicationAdministration) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "device":
		 return fhirVal.Device, true
 	case "medicationreference":
		 return fhirVal.MedicationReference, true
 	case "context":
		 return fhirVal.Context, true
 	case "performer":
		 return fhirVal.Performer, true
 	case "notgiven":
		 return fhirVal.NotGiven, true
 	case "reasoncode":
		 return fhirVal.ReasonCode, true
 	case "reasonreference":
		 return fhirVal.ReasonReference, true
 	case "definition":
		 return fhirVal.Definition, true
 	case "partof":
		 return fhirVal.PartOf, true
 	case "category":
		 return fhirVal.Category, true
 	case "effectivedatetime":
		 return fhirVal.EffectiveDateTime, true
 	case "effectiveperiod":
		 return fhirVal.EffectivePeriod, true
 	case "prescription":
		 return fhirVal.Prescription, true
 	case "note":
		 return fhirVal.Note, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "medicationcodeableconcept":
		 return fhirVal.MedicationCodeableConcept, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "supportinginformation":
		 return fhirVal.SupportingInformation, true
 	case "reasonnotgiven":
		 return fhirVal.ReasonNotGiven, true
 	case "dosage":
		 return fhirVal.Dosage, true
 	case "eventhistory":
		 return fhirVal.EventHistory, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *MedicationAdministration) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"device": &FieldTypeSupport{"Reference", true, false},
 		"medicationreference": &FieldTypeSupport{"Reference", false, true},
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"performer": &FieldTypeSupport{"MedicationAdministrationPerformerComponent", true, false},
 		"notgiven": &FieldTypeSupport{"bool", false, true},
 		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
 		"reasonreference": &FieldTypeSupport{"Reference", true, false},
 		"definition": &FieldTypeSupport{"Reference", true, false},
 		"partof": &FieldTypeSupport{"Reference", true, false},
 		"category": &FieldTypeSupport{"CodeableConcept", false, true},
 		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
 		"prescription": &FieldTypeSupport{"Reference", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
 		"reasonnotgiven": &FieldTypeSupport{"CodeableConcept", true, false},
 		"dosage": &FieldTypeSupport{"MedicationAdministrationDosageComponent", false, true},
 		"eventhistory": &FieldTypeSupport{"Reference", true, false},
 
	 }
 }
 
 func (fhirVal *MedicationAdministration) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 