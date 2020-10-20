
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
	case "usage":
		return fhirVal.Usage, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "topic":
		return fhirVal.Topic, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "version":
		return fhirVal.Version, true
	case "title":
		return fhirVal.Title, true
	case "type":
		return fhirVal.Type, true
	case "purpose":
		return fhirVal.Purpose, true
	case "goal":
		return fhirVal.Goal, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "copyright":
		return fhirVal.Copyright, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "date":
		return fhirVal.Date, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "url":
		return fhirVal.Url, true
	case "experimental":
		return fhirVal.Experimental, true
	case "contributor":
		return fhirVal.Contributor, true
	case "action":
		return fhirVal.Action, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "contact":
		return fhirVal.Contact, true
	case "library":
		return fhirVal.Library, true
	case "name":
		return fhirVal.Name, true
	case "publisher":
		return fhirVal.Publisher, true
	case "description":
		return fhirVal.Description, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true

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
		"usage": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"purpose": &FieldTypeSupport{"string", false, false},
		"goal": &FieldTypeSupport{"PlanDefinitionGoalComponent", true, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"copyright": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"action": &FieldTypeSupport{"PlanDefinitionActionComponent", true, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"library": &FieldTypeSupport{"Reference", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
