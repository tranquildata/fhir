
package models


func (fhirVal *ClinicalImpression) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ClinicalImpression) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "finding":
		return fhirVal.Finding, true
	case "prognosisreference":
		return fhirVal.PrognosisReference, true
	case "action":
		return fhirVal.Action, true
	case "description":
		return fhirVal.Description, true
	case "subject":
		return fhirVal.Subject, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "assessor":
		return fhirVal.Assessor, true
	case "summary":
		return fhirVal.Summary, true
	case "status":
		return fhirVal.Status, true
	case "date":
		return fhirVal.Date, true
	case "problem":
		return fhirVal.Problem, true
	case "code":
		return fhirVal.Code, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true
	case "investigation":
		return fhirVal.Investigation, true
	case "note":
		return fhirVal.Note, true
	case "identifier":
		return fhirVal.Identifier, true
	case "context":
		return fhirVal.Context, true
	case "previous":
		return fhirVal.Previous, true
	case "protocol":
		return fhirVal.Protocol, true
	case "prognosiscodeableconcept":
		return fhirVal.PrognosisCodeableConcept, true

	default:
		return nil, false
	}
}

func (fhirVal *ClinicalImpression) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"finding": &FieldTypeSupport{"ClinicalImpressionFindingComponent", true, false},
		"prognosisreference": &FieldTypeSupport{"Reference", true, false},
		"action": &FieldTypeSupport{"Reference", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"assessor": &FieldTypeSupport{"Reference", false, true},
		"summary": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"problem": &FieldTypeSupport{"Reference", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"investigation": &FieldTypeSupport{"ClinicalImpressionInvestigationComponent", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"previous": &FieldTypeSupport{"Reference", false, true},
		"protocol": &FieldTypeSupport{"string", true, false},
		"prognosiscodeableconcept": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
