
package models


func (fhirVal *Library) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Library) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "title":
		return fhirVal.Title, true
	case "publisher":
		return fhirVal.Publisher, true
	case "usage":
		return fhirVal.Usage, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "topic":
		return fhirVal.Topic, true
	case "datarequirement":
		return fhirVal.DataRequirement, true
	case "identifier":
		return fhirVal.Identifier, true
	case "version":
		return fhirVal.Version, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "contributor":
		return fhirVal.Contributor, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "content":
		return fhirVal.Content, true
	case "description":
		return fhirVal.Description, true
	case "purpose":
		return fhirVal.Purpose, true
	case "type":
		return fhirVal.Type, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "contact":
		return fhirVal.Contact, true
	case "parameter":
		return fhirVal.Parameter, true
	case "url":
		return fhirVal.Url, true
	case "experimental":
		return fhirVal.Experimental, true
	case "date":
		return fhirVal.Date, true
	case "copyright":
		return fhirVal.Copyright, true
	case "name":
		return fhirVal.Name, true
	case "status":
		return fhirVal.Status, true

	default:
		return nil, false
	}
}

func (fhirVal *Library) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"title": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"datarequirement": &FieldTypeSupport{"DataRequirement", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"content": &FieldTypeSupport{"Attachment", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"parameter": &FieldTypeSupport{"ParameterDefinition", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"copyright": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},

	}
}
