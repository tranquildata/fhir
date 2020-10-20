
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
		"action": &FieldTypeSupport{"PlanDefinitionActionComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"goal": &FieldTypeSupport{"PlanDefinitionGoalComponent", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"library": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"copyright": &FieldTypeSupport{"string", false, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},

	}
}
