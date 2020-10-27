
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ActivityDefinition) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ActivityDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "productcodeableconcept":
		 return fhirVal.ProductCodeableConcept, true
 	case "quantity":
		 return fhirVal.Quantity, true
 	case "dosage":
		 return fhirVal.Dosage, true
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "transform":
		 return fhirVal.Transform, true
 	case "kind":
		 return fhirVal.Kind, true
 	case "timingtiming":
		 return fhirVal.TimingTiming, true
 	case "productreference":
		 return fhirVal.ProductReference, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "effectiveperiod":
		 return fhirVal.EffectivePeriod, true
 	case "contributor":
		 return fhirVal.Contributor, true
 	case "timingrange":
		 return fhirVal.TimingRange, true
 	case "url":
		 return fhirVal.Url, true
 	case "version":
		 return fhirVal.Version, true
 	case "lastreviewdate":
		 return fhirVal.LastReviewDate, true
 	case "code":
		 return fhirVal.Code, true
 	case "timingperiod":
		 return fhirVal.TimingPeriod, true
 	case "dynamicvalue":
		 return fhirVal.DynamicValue, true
 	case "title":
		 return fhirVal.Title, true
 	case "description":
		 return fhirVal.Description, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "topic":
		 return fhirVal.Topic, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "timingdatetime":
		 return fhirVal.TimingDateTime, true
 	case "location":
		 return fhirVal.Location, true
 	case "participant":
		 return fhirVal.Participant, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "date":
		 return fhirVal.Date, true
 	case "approvaldate":
		 return fhirVal.ApprovalDate, true
 	case "library":
		 return fhirVal.Library, true
 	case "bodysite":
		 return fhirVal.BodySite, true
 	case "name":
		 return fhirVal.Name, true
 	case "status":
		 return fhirVal.Status, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "usage":
		 return fhirVal.Usage, true
 	case "copyright":
		 return fhirVal.Copyright, true
 	case "relatedartifact":
		 return fhirVal.RelatedArtifact, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ActivityDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"productcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
 		"quantity": &FieldTypeSupport{"Quantity", false, true},
 		"dosage": &FieldTypeSupport{"Dosage", true, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"transform": &FieldTypeSupport{"Reference", false, true},
 		"kind": &FieldTypeSupport{"string", false, false},
 		"timingtiming": &FieldTypeSupport{"Timing", false, true},
 		"productreference": &FieldTypeSupport{"Reference", false, true},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
 		"contributor": &FieldTypeSupport{"Contributor", true, false},
 		"timingrange": &FieldTypeSupport{"Range", false, true},
 		"url": &FieldTypeSupport{"string", false, false},
 		"version": &FieldTypeSupport{"string", false, false},
 		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"timingperiod": &FieldTypeSupport{"Period", false, true},
 		"dynamicvalue": &FieldTypeSupport{"ActivityDefinitionDynamicValueComponent", true, false},
 		"title": &FieldTypeSupport{"string", false, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"timingdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"location": &FieldTypeSupport{"Reference", false, true},
 		"participant": &FieldTypeSupport{"ActivityDefinitionParticipantComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"library": &FieldTypeSupport{"Reference", true, false},
 		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"usage": &FieldTypeSupport{"string", false, false},
 		"copyright": &FieldTypeSupport{"string", false, false},
 		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
 
	 }
 }
 
 func (fhirVal *ActivityDefinition) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 