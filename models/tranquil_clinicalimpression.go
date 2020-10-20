
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
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"investigation": &FieldTypeSupport{"ClinicalImpressionInvestigationComponent", true, false},
		"finding": &FieldTypeSupport{"ClinicalImpressionFindingComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"assessor": &FieldTypeSupport{"Reference", false, true},
		"protocol": &FieldTypeSupport{"string", true, false},
		"summary": &FieldTypeSupport{"string", false, false},
		"prognosiscodeableconcept": &FieldTypeSupport{"CodeableConcept", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"problem": &FieldTypeSupport{"Reference", true, false},
		"action": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"previous": &FieldTypeSupport{"Reference", false, true},
		"prognosisreference": &FieldTypeSupport{"Reference", true, false},

	}
}
