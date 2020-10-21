
package models


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
	case "experimental":
		return fhirVal.Experimental, true
	case "date":
		return fhirVal.Date, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "subjecttype":
		return fhirVal.SubjectType, true
	case "item":
		return fhirVal.Item, true
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "name":
		return fhirVal.Name, true
	case "description":
		return fhirVal.Description, true
	case "purpose":
		return fhirVal.Purpose, true
	case "copyright":
		return fhirVal.Copyright, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "contact":
		return fhirVal.Contact, true
	case "code":
		return fhirVal.Code, true
	case "identifier":
		return fhirVal.Identifier, true
	case "title":
		return fhirVal.Title, true
	case "status":
		return fhirVal.Status, true
	case "publisher":
		return fhirVal.Publisher, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true

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
		"experimental": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"subjecttype": &FieldTypeSupport{"string", true, false},
		"item": &FieldTypeSupport{"QuestionnaireItemComponent", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"code": &FieldTypeSupport{"Coding", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
