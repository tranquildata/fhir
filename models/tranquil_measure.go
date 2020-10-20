
package models


func (fhirVal *Measure) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Measure) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "riskadjustment":
		return fhirVal.RiskAdjustment, true
	case "supplementaldata":
		return fhirVal.SupplementalData, true
	case "identifier":
		return fhirVal.Identifier, true
	case "title":
		return fhirVal.Title, true
	case "status":
		return fhirVal.Status, true
	case "usage":
		return fhirVal.Usage, true
	case "type":
		return fhirVal.Type, true
	case "improvementnotation":
		return fhirVal.ImprovementNotation, true
	case "set":
		return fhirVal.Set, true
	case "purpose":
		return fhirVal.Purpose, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "copyright":
		return fhirVal.Copyright, true
	case "rationale":
		return fhirVal.Rationale, true
	case "topic":
		return fhirVal.Topic, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "guidance":
		return fhirVal.Guidance, true
	case "rateaggregation":
		return fhirVal.RateAggregation, true
	case "definition":
		return fhirVal.Definition, true
	case "experimental":
		return fhirVal.Experimental, true
	case "date":
		return fhirVal.Date, true
	case "contributor":
		return fhirVal.Contributor, true
	case "contact":
		return fhirVal.Contact, true
	case "library":
		return fhirVal.Library, true
	case "name":
		return fhirVal.Name, true
	case "publisher":
		return fhirVal.Publisher, true
	case "scoring":
		return fhirVal.Scoring, true
	case "version":
		return fhirVal.Version, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "compositescoring":
		return fhirVal.CompositeScoring, true
	case "group":
		return fhirVal.Group, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "disclaimer":
		return fhirVal.Disclaimer, true
	case "url":
		return fhirVal.Url, true
	case "description":
		return fhirVal.Description, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "clinicalrecommendationstatement":
		return fhirVal.ClinicalRecommendationStatement, true

	default:
		return nil, false
	}
}

func (fhirVal *Measure) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"riskadjustment": &FieldTypeSupport{"string", false, false},
		"supplementaldata": &FieldTypeSupport{"MeasureSupplementalDataComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", true, false},
		"improvementnotation": &FieldTypeSupport{"string", false, false},
		"set": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"rationale": &FieldTypeSupport{"string", false, false},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"guidance": &FieldTypeSupport{"string", false, false},
		"rateaggregation": &FieldTypeSupport{"string", false, false},
		"definition": &FieldTypeSupport{"string", true, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"library": &FieldTypeSupport{"Reference", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"scoring": &FieldTypeSupport{"CodeableConcept", false, true},
		"version": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"compositescoring": &FieldTypeSupport{"CodeableConcept", false, true},
		"group": &FieldTypeSupport{"MeasureGroupComponent", true, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"disclaimer": &FieldTypeSupport{"string", false, false},
		"url": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"clinicalrecommendationstatement": &FieldTypeSupport{"string", false, false},

	}
}
