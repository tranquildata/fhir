
package models


func (fhirVal *TestScript) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *TestScript) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *TestScript) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"profile": &FieldTypeSupport{"Reference", true, false},
		"rule": &FieldTypeSupport{"TestScriptRuleComponent", true, false},
		"test": &FieldTypeSupport{"TestScriptTestComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"name": &FieldTypeSupport{"string", false, false},
		"fixture": &FieldTypeSupport{"TestScriptFixtureComponent", true, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"origin": &FieldTypeSupport{"TestScriptOriginComponent", true, false},
		"destination": &FieldTypeSupport{"TestScriptDestinationComponent", true, false},
		"variable": &FieldTypeSupport{"TestScriptVariableComponent", true, false},
		"ruleset": &FieldTypeSupport{"TestScriptRulesetComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"title": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"metadata": &FieldTypeSupport{"TestScriptMetadataComponent", false, true},
		"setup": &FieldTypeSupport{"TestScriptSetupComponent", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"teardown": &FieldTypeSupport{"TestScriptTeardownComponent", false, true},

	}
}
