
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Condition) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Condition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "abatementstring":
		 return fhirVal.AbatementString, true
 	case "asserteddate":
		 return fhirVal.AssertedDate, true
 	case "evidence":
		 return fhirVal.Evidence, true
 	case "note":
		 return fhirVal.Note, true
 	case "context":
		 return fhirVal.Context, true
 	case "onsetage":
		 return fhirVal.OnsetAge, true
 	case "onsetrange":
		 return fhirVal.OnsetRange, true
 	case "abatementperiod":
		 return fhirVal.AbatementPeriod, true
 	case "abatementboolean":
		 return fhirVal.AbatementBoolean, true
 	case "severity":
		 return fhirVal.Severity, true
 	case "bodysite":
		 return fhirVal.BodySite, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "onsetstring":
		 return fhirVal.OnsetString, true
 	case "onsetdatetime":
		 return fhirVal.OnsetDateTime, true
 	case "onsetperiod":
		 return fhirVal.OnsetPeriod, true
 	case "abatementdatetime":
		 return fhirVal.AbatementDateTime, true
 	case "abatementage":
		 return fhirVal.AbatementAge, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "clinicalstatus":
		 return fhirVal.ClinicalStatus, true
 	case "verificationstatus":
		 return fhirVal.VerificationStatus, true
 	case "code":
		 return fhirVal.Code, true
 	case "stage":
		 return fhirVal.Stage, true
 	case "category":
		 return fhirVal.Category, true
 	case "abatementrange":
		 return fhirVal.AbatementRange, true
 	case "asserter":
		 return fhirVal.Asserter, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Condition) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"abatementstring": &FieldTypeSupport{"string", false, false},
 		"asserteddate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"evidence": &FieldTypeSupport{"ConditionEvidenceComponent", true, false},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"onsetage": &FieldTypeSupport{"Quantity", false, true},
 		"onsetrange": &FieldTypeSupport{"Range", false, true},
 		"abatementperiod": &FieldTypeSupport{"Period", false, true},
 		"abatementboolean": &FieldTypeSupport{"bool", false, true},
 		"severity": &FieldTypeSupport{"CodeableConcept", false, true},
 		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"onsetstring": &FieldTypeSupport{"string", false, false},
 		"onsetdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"onsetperiod": &FieldTypeSupport{"Period", false, true},
 		"abatementdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"abatementage": &FieldTypeSupport{"Quantity", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"clinicalstatus": &FieldTypeSupport{"CodeableConcept", false, true},
 		"verificationstatus": &FieldTypeSupport{"CodeableConcept", false, true},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"stage": &FieldTypeSupport{"ConditionStageComponent", false, true},
 		"category": &FieldTypeSupport{"CodeableConcept", true, false},
 		"abatementrange": &FieldTypeSupport{"Range", false, true},
 		"asserter": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *Condition) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 