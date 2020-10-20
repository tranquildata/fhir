
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
	case "purpose":
		return fhirVal.Purpose, true
	case "metadata":
		return fhirVal.Metadata, true
	case "profile":
		return fhirVal.Profile, true
	case "rule":
		return fhirVal.Rule, true
	case "description":
		return fhirVal.Description, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "origin":
		return fhirVal.Origin, true
	case "fixture":
		return fhirVal.Fixture, true
	case "url":
		return fhirVal.Url, true
	case "title":
		return fhirVal.Title, true
	case "date":
		return fhirVal.Date, true
	case "contact":
		return fhirVal.Contact, true
	case "setup":
		return fhirVal.Setup, true
	case "test":
		return fhirVal.Test, true
	case "copyright":
		return fhirVal.Copyright, true
	case "variable":
		return fhirVal.Variable, true
	case "teardown":
		return fhirVal.Teardown, true
	case "identifier":
		return fhirVal.Identifier, true
	case "version":
		return fhirVal.Version, true
	case "status":
		return fhirVal.Status, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "ruleset":
		return fhirVal.Ruleset, true
	case "name":
		return fhirVal.Name, true
	case "experimental":
		return fhirVal.Experimental, true
	case "publisher":
		return fhirVal.Publisher, true
	case "destination":
		return fhirVal.Destination, true

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
		"purpose": &FieldTypeSupport{"string", false, false},
		"metadata": &FieldTypeSupport{"TestScriptMetadataComponent", false, true},
		"profile": &FieldTypeSupport{"Reference", true, false},
		"rule": &FieldTypeSupport{"TestScriptRuleComponent", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"origin": &FieldTypeSupport{"TestScriptOriginComponent", true, false},
		"fixture": &FieldTypeSupport{"TestScriptFixtureComponent", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"setup": &FieldTypeSupport{"TestScriptSetupComponent", false, true},
		"test": &FieldTypeSupport{"TestScriptTestComponent", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"variable": &FieldTypeSupport{"TestScriptVariableComponent", true, false},
		"teardown": &FieldTypeSupport{"TestScriptTeardownComponent", false, true},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"version": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"ruleset": &FieldTypeSupport{"TestScriptRulesetComponent", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"destination": &FieldTypeSupport{"TestScriptDestinationComponent", true, false},

	}
}
