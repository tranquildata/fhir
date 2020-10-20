
package models


func (fhirVal *AllergyIntolerance) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *AllergyIntolerance) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "verificationstatus":
		return fhirVal.VerificationStatus, true
	case "patient":
		return fhirVal.Patient, true
	case "onsetperiod":
		return fhirVal.OnsetPeriod, true
	case "clinicalstatus":
		return fhirVal.ClinicalStatus, true
	case "onsetdatetime":
		return fhirVal.OnsetDateTime, true
	case "onsetrange":
		return fhirVal.OnsetRange, true
	case "asserter":
		return fhirVal.Asserter, true
	case "criticality":
		return fhirVal.Criticality, true
	case "code":
		return fhirVal.Code, true
	case "onsetstring":
		return fhirVal.OnsetString, true
	case "asserteddate":
		return fhirVal.AssertedDate, true
	case "reaction":
		return fhirVal.Reaction, true
	case "category":
		return fhirVal.Category, true
	case "type":
		return fhirVal.Type, true
	case "onsetage":
		return fhirVal.OnsetAge, true
	case "recorder":
		return fhirVal.Recorder, true
	case "lastoccurrence":
		return fhirVal.LastOccurrence, true
	case "note":
		return fhirVal.Note, true
	case "identifier":
		return fhirVal.Identifier, true

	default:
		return nil, false
	}
}

func (fhirVal *AllergyIntolerance) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"verificationstatus": &FieldTypeSupport{"string", false, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"onsetperiod": &FieldTypeSupport{"Period", false, true},
		"clinicalstatus": &FieldTypeSupport{"string", false, false},
		"onsetdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"onsetrange": &FieldTypeSupport{"Range", false, true},
		"asserter": &FieldTypeSupport{"Reference", false, true},
		"criticality": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"onsetstring": &FieldTypeSupport{"string", false, false},
		"asserteddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reaction": &FieldTypeSupport{"AllergyIntoleranceReactionComponent", true, false},
		"category": &FieldTypeSupport{"string", true, false},
		"type": &FieldTypeSupport{"string", false, false},
		"onsetage": &FieldTypeSupport{"Quantity", false, true},
		"recorder": &FieldTypeSupport{"Reference", false, true},
		"lastoccurrence": &FieldTypeSupport{"FHIRDateTime", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},

	}
}
