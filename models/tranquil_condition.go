
package models


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
	case "onsetdatetime":
		return fhirVal.OnsetDateTime, true
	case "onsetrange":
		return fhirVal.OnsetRange, true
	case "onsetstring":
		return fhirVal.OnsetString, true
	case "abatementdatetime":
		return fhirVal.AbatementDateTime, true
	case "abatementage":
		return fhirVal.AbatementAge, true
	case "identifier":
		return fhirVal.Identifier, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "context":
		return fhirVal.Context, true
	case "abatementperiod":
		return fhirVal.AbatementPeriod, true
	case "asserteddate":
		return fhirVal.AssertedDate, true
	case "abatementstring":
		return fhirVal.AbatementString, true
	case "asserter":
		return fhirVal.Asserter, true
	case "stage":
		return fhirVal.Stage, true
	case "note":
		return fhirVal.Note, true
	case "category":
		return fhirVal.Category, true
	case "onsetage":
		return fhirVal.OnsetAge, true
	case "onsetperiod":
		return fhirVal.OnsetPeriod, true
	case "abatementrange":
		return fhirVal.AbatementRange, true
	case "clinicalstatus":
		return fhirVal.ClinicalStatus, true
	case "code":
		return fhirVal.Code, true
	case "abatementboolean":
		return fhirVal.AbatementBoolean, true
	case "evidence":
		return fhirVal.Evidence, true
	case "verificationstatus":
		return fhirVal.VerificationStatus, true
	case "severity":
		return fhirVal.Severity, true
	case "subject":
		return fhirVal.Subject, true

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
		"onsetdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"onsetrange": &FieldTypeSupport{"Range", false, true},
		"onsetstring": &FieldTypeSupport{"string", false, false},
		"abatementdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"abatementage": &FieldTypeSupport{"Quantity", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"abatementperiod": &FieldTypeSupport{"Period", false, true},
		"asserteddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"abatementstring": &FieldTypeSupport{"string", false, false},
		"asserter": &FieldTypeSupport{"Reference", false, true},
		"stage": &FieldTypeSupport{"ConditionStageComponent", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"onsetage": &FieldTypeSupport{"Quantity", false, true},
		"onsetperiod": &FieldTypeSupport{"Period", false, true},
		"abatementrange": &FieldTypeSupport{"Range", false, true},
		"clinicalstatus": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"abatementboolean": &FieldTypeSupport{"bool", false, true},
		"evidence": &FieldTypeSupport{"ConditionEvidenceComponent", true, false},
		"verificationstatus": &FieldTypeSupport{"string", false, false},
		"severity": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},

	}
}
