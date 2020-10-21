
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
	case "status":
		return fhirVal.Status, true
	case "experimental":
		return fhirVal.Experimental, true
	case "purpose":
		return fhirVal.Purpose, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "identifier":
		return fhirVal.Identifier, true
	case "title":
		return fhirVal.Title, true
	case "date":
		return fhirVal.Date, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "contact":
		return fhirVal.Contact, true
	case "copyright":
		return fhirVal.Copyright, true
	case "trigger":
		return fhirVal.Trigger, true
	case "datarequirement":
		return fhirVal.DataRequirement, true
	case "version":
		return fhirVal.Version, true
	case "publisher":
		return fhirVal.Publisher, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "contributor":
		return fhirVal.Contributor, true
	case "operationdefinition":
		return fhirVal.OperationDefinition, true
	case "url":
		return fhirVal.Url, true
	case "name":
		return fhirVal.Name, true
	case "description":
		return fhirVal.Description, true
	case "usage":
		return fhirVal.Usage, true
	case "topic":
		return fhirVal.Topic, true

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
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"purpose": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"trigger": &FieldTypeSupport{"TriggerDefinition", true, false},
		"datarequirement": &FieldTypeSupport{"DataRequirement", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"operationdefinition": &FieldTypeSupport{"Reference", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
