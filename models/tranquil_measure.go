
 package models
 
 import (
	 "fmt"
 )
 
 
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
 	case "copyright":
		 return fhirVal.Copyright, true
 	case "version":
		 return fhirVal.Version, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "topic":
		 return fhirVal.Topic, true
 	case "rateaggregation":
		 return fhirVal.RateAggregation, true
 	case "rationale":
		 return fhirVal.Rationale, true
 	case "set":
		 return fhirVal.Set, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "lastreviewdate":
		 return fhirVal.LastReviewDate, true
 	case "relatedartifact":
		 return fhirVal.RelatedArtifact, true
 	case "title":
		 return fhirVal.Title, true
 	case "description":
		 return fhirVal.Description, true
 	case "contributor":
		 return fhirVal.Contributor, true
 	case "clinicalrecommendationstatement":
		 return fhirVal.ClinicalRecommendationStatement, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "name":
		 return fhirVal.Name, true
 	case "approvaldate":
		 return fhirVal.ApprovalDate, true
 	case "library":
		 return fhirVal.Library, true
 	case "compositescoring":
		 return fhirVal.CompositeScoring, true
 	case "definition":
		 return fhirVal.Definition, true
 	case "url":
		 return fhirVal.Url, true
 	case "status":
		 return fhirVal.Status, true
 	case "date":
		 return fhirVal.Date, true
 	case "effectiveperiod":
		 return fhirVal.EffectivePeriod, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "disclaimer":
		 return fhirVal.Disclaimer, true
 	case "type":
		 return fhirVal.Type, true
 	case "supplementaldata":
		 return fhirVal.SupplementalData, true
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "usage":
		 return fhirVal.Usage, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "riskadjustment":
		 return fhirVal.RiskAdjustment, true
 	case "improvementnotation":
		 return fhirVal.ImprovementNotation, true
 	case "guidance":
		 return fhirVal.Guidance, true
 	case "scoring":
		 return fhirVal.Scoring, true
 	case "group":
		 return fhirVal.Group, true
 
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
 		"copyright": &FieldTypeSupport{"string", false, false},
 		"version": &FieldTypeSupport{"string", false, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
 		"rateaggregation": &FieldTypeSupport{"string", false, false},
 		"rationale": &FieldTypeSupport{"string", false, false},
 		"set": &FieldTypeSupport{"string", false, false},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
 		"title": &FieldTypeSupport{"string", false, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"contributor": &FieldTypeSupport{"Contributor", true, false},
 		"clinicalrecommendationstatement": &FieldTypeSupport{"string", false, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"library": &FieldTypeSupport{"Reference", true, false},
 		"compositescoring": &FieldTypeSupport{"CodeableConcept", false, true},
 		"definition": &FieldTypeSupport{"string", true, false},
 		"url": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"disclaimer": &FieldTypeSupport{"string", false, false},
 		"type": &FieldTypeSupport{"CodeableConcept", true, false},
 		"supplementaldata": &FieldTypeSupport{"MeasureSupplementalDataComponent", true, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"usage": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"riskadjustment": &FieldTypeSupport{"string", false, false},
 		"improvementnotation": &FieldTypeSupport{"string", false, false},
 		"guidance": &FieldTypeSupport{"string", false, false},
 		"scoring": &FieldTypeSupport{"CodeableConcept", false, true},
 		"group": &FieldTypeSupport{"MeasureGroupComponent", true, false},
 
	 }
 }
 
 func (fhirVal *Measure) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 