
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
	case "category":
		return fhirVal.Category, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "specimen":
		return fhirVal.Specimen, true
	case "codeddiagnosis":
		return fhirVal.CodedDiagnosis, true
	case "presentedform":
		return fhirVal.PresentedForm, true
	case "identifier":
		return fhirVal.Identifier, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "code":
		return fhirVal.Code, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true
	case "performer":
		return fhirVal.Performer, true
	case "status":
		return fhirVal.Status, true
	case "issued":
		return fhirVal.Issued, true
	case "subject":
		return fhirVal.Subject, true
	case "context":
		return fhirVal.Context, true
	case "result":
		return fhirVal.Result, true
	case "imagingstudy":
		return fhirVal.ImagingStudy, true
	case "image":
		return fhirVal.Image, true
	case "conclusion":
		return fhirVal.Conclusion, true

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
		"category": &FieldTypeSupport{"CodeableConcept", true, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"specimen": &FieldTypeSupport{"Reference", true, false},
		"codeddiagnosis": &FieldTypeSupport{"CodeableConcept", true, false},
		"presentedform": &FieldTypeSupport{"Attachment", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"performer": &FieldTypeSupport{"DiagnosticReportPerformerComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"result": &FieldTypeSupport{"Reference", true, false},
		"imagingstudy": &FieldTypeSupport{"Reference", true, false},
		"image": &FieldTypeSupport{"DiagnosticReportImageComponent", true, false},
		"conclusion": &FieldTypeSupport{"string", false, false},

	}
}
