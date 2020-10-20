
package models


func (fhirVal *PlanDefinition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *PlanDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "contact":
		return fhirVal.Contact, true
	case "goal":
		return fhirVal.Goal, true
	case "action":
		return fhirVal.Action, true
	case "purpose":
		return fhirVal.Purpose, true
	case "usage":
		return fhirVal.Usage, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "contributor":
		return fhirVal.Contributor, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "library":
		return fhirVal.Library, true
	case "version":
		return fhirVal.Version, true
	case "status":
		return fhirVal.Status, true
	case "date":
		return fhirVal.Date, true
	case "description":
		return fhirVal.Description, true
	case "copyright":
		return fhirVal.Copyright, true
	case "identifier":
		return fhirVal.Identifier, true
	case "experimental":
		return fhirVal.Experimental, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "topic":
		return fhirVal.Topic, true
	case "publisher":
		return fhirVal.Publisher, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "url":
		return fhirVal.Url, true
	case "name":
		return fhirVal.Name, true
	case "title":
		return fhirVal.Title, true
	case "type":
		return fhirVal.Type, true

	default:
		return nil, false
	}
}

func (fhirVal *PlanDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"goal": &FieldTypeSupport{"PlanDefinitionGoalComponent", true, false},
		"action": &FieldTypeSupport{"PlanDefinitionActionComponent", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"library": &FieldTypeSupport{"Reference", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
