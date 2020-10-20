
package models


func (fhirVal *DiagnosticReport) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *DiagnosticReport) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "issued":
		return fhirVal.Issued, true
	case "code":
		return fhirVal.Code, true
	case "performer":
		return fhirVal.Performer, true
	case "specimen":
		return fhirVal.Specimen, true
	case "imagingstudy":
		return fhirVal.ImagingStudy, true
	case "identifier":
		return fhirVal.Identifier, true
	case "subject":
		return fhirVal.Subject, true
	case "context":
		return fhirVal.Context, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "result":
		return fhirVal.Result, true
	case "conclusion":
		return fhirVal.Conclusion, true
	case "codeddiagnosis":
		return fhirVal.CodedDiagnosis, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "status":
		return fhirVal.Status, true
	case "category":
		return fhirVal.Category, true
	case "image":
		return fhirVal.Image, true
	case "presentedform":
		return fhirVal.PresentedForm, true

	default:
		return nil, false
	}
}

func (fhirVal *DiagnosticReport) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"performer": &FieldTypeSupport{"DiagnosticReportPerformerComponent", true, false},
		"specimen": &FieldTypeSupport{"Reference", true, false},
		"imagingstudy": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"result": &FieldTypeSupport{"Reference", true, false},
		"conclusion": &FieldTypeSupport{"string", false, false},
		"codeddiagnosis": &FieldTypeSupport{"CodeableConcept", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"image": &FieldTypeSupport{"DiagnosticReportImageComponent", true, false},
		"presentedform": &FieldTypeSupport{"Attachment", true, false},

	}
}
