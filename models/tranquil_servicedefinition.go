
package models


func (fhirVal *ServiceDefinition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ServiceDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "contributor":
		return fhirVal.Contributor, true
	case "datarequirement":
		return fhirVal.DataRequirement, true
	case "identifier":
		return fhirVal.Identifier, true
	case "title":
		return fhirVal.Title, true
	case "status":
		return fhirVal.Status, true
	case "date":
		return fhirVal.Date, true
	case "description":
		return fhirVal.Description, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "topic":
		return fhirVal.Topic, true
	case "copyright":
		return fhirVal.Copyright, true
	case "experimental":
		return fhirVal.Experimental, true
	case "publisher":
		return fhirVal.Publisher, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "operationdefinition":
		return fhirVal.OperationDefinition, true
	case "name":
		return fhirVal.Name, true
	case "purpose":
		return fhirVal.Purpose, true
	case "usage":
		return fhirVal.Usage, true
	case "contact":
		return fhirVal.Contact, true
	case "trigger":
		return fhirVal.Trigger, true

	default:
		return nil, false
	}
}

func (fhirVal *ServiceDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"datarequirement": &FieldTypeSupport{"DataRequirement", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"operationdefinition": &FieldTypeSupport{"Reference", false, true},
		"name": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"trigger": &FieldTypeSupport{"TriggerDefinition", true, false},

	}
}
