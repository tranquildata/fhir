
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
		"status": &FieldTypeSupport{"string", false, false},
		"specimen": &FieldTypeSupport{"Reference", true, false},
		"result": &FieldTypeSupport{"Reference", true, false},
		"imagingstudy": &FieldTypeSupport{"Reference", true, false},
		"codeddiagnosis": &FieldTypeSupport{"CodeableConcept", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"image": &FieldTypeSupport{"DiagnosticReportImageComponent", true, false},
		"conclusion": &FieldTypeSupport{"string", false, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"presentedform": &FieldTypeSupport{"Attachment", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"performer": &FieldTypeSupport{"DiagnosticReportPerformerComponent", true, false},

	}
}
