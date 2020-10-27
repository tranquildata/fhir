
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Questionnaire) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Questionnaire) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "subjecttype":
		 return fhirVal.SubjectType, true
 	case "description":
		 return fhirVal.Description, true
 	case "date":
		 return fhirVal.Date, true
 	case "effectiveperiod":
		 return fhirVal.EffectivePeriod, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "version":
		 return fhirVal.Version, true
 	case "name":
		 return fhirVal.Name, true
 	case "title":
		 return fhirVal.Title, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "approvaldate":
		 return fhirVal.ApprovalDate, true
 	case "lastreviewdate":
		 return fhirVal.LastReviewDate, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "url":
		 return fhirVal.Url, true
 	case "item":
		 return fhirVal.Item, true
 	case "code":
		 return fhirVal.Code, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Questionnaire) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"subjecttype": &FieldTypeSupport{"string", true, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"version": &FieldTypeSupport{"string", false, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"title": &FieldTypeSupport{"string", false, false},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"url": &FieldTypeSupport{"string", false, false},
 		"item": &FieldTypeSupport{"QuestionnaireItemComponent", true, false},
 		"code": &FieldTypeSupport{"Coding", true, false},
 
	 }
 }
 
 func (fhirVal *Questionnaire) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 