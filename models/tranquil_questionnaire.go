
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
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "identifier":
		return fhirVal.Identifier, true
	case "name":
		return fhirVal.Name, true
	case "status":
		return fhirVal.Status, true
	case "experimental":
		return fhirVal.Experimental, true
	case "description":
		return fhirVal.Description, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "copyright":
		return fhirVal.Copyright, true
	case "code":
		return fhirVal.Code, true
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "item":
		return fhirVal.Item, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "subjecttype":
		return fhirVal.SubjectType, true
	case "date":
		return fhirVal.Date, true
	case "publisher":
		return fhirVal.Publisher, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "contact":
		return fhirVal.Contact, true
	case "title":
		return fhirVal.Title, true
	case "purpose":
		return fhirVal.Purpose, true

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
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"copyright": &FieldTypeSupport{"string", false, false},
		"code": &FieldTypeSupport{"Coding", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"item": &FieldTypeSupport{"QuestionnaireItemComponent", true, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"subjecttype": &FieldTypeSupport{"string", true, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},

	}
}
