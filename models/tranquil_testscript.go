
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
	case "publisher":
		return fhirVal.Publisher, true
	case "contact":
		return fhirVal.Contact, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "origin":
		return fhirVal.Origin, true
	case "metadata":
		return fhirVal.Metadata, true
	case "status":
		return fhirVal.Status, true
	case "name":
		return fhirVal.Name, true
	case "experimental":
		return fhirVal.Experimental, true
	case "date":
		return fhirVal.Date, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "copyright":
		return fhirVal.Copyright, true
	case "setup":
		return fhirVal.Setup, true
	case "version":
		return fhirVal.Version, true
	case "purpose":
		return fhirVal.Purpose, true
	case "variable":
		return fhirVal.Variable, true
	case "rule":
		return fhirVal.Rule, true
	case "description":
		return fhirVal.Description, true
	case "identifier":
		return fhirVal.Identifier, true
	case "title":
		return fhirVal.Title, true
	case "destination":
		return fhirVal.Destination, true
	case "fixture":
		return fhirVal.Fixture, true
	case "profile":
		return fhirVal.Profile, true
	case "ruleset":
		return fhirVal.Ruleset, true
	case "test":
		return fhirVal.Test, true
	case "url":
		return fhirVal.Url, true
	case "teardown":
		return fhirVal.Teardown, true

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
		"publisher": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"origin": &FieldTypeSupport{"TestScriptOriginComponent", true, false},
		"metadata": &FieldTypeSupport{"TestScriptMetadataComponent", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"setup": &FieldTypeSupport{"TestScriptSetupComponent", false, true},
		"version": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"variable": &FieldTypeSupport{"TestScriptVariableComponent", true, false},
		"rule": &FieldTypeSupport{"TestScriptRuleComponent", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"title": &FieldTypeSupport{"string", false, false},
		"destination": &FieldTypeSupport{"TestScriptDestinationComponent", true, false},
		"fixture": &FieldTypeSupport{"TestScriptFixtureComponent", true, false},
		"profile": &FieldTypeSupport{"Reference", true, false},
		"ruleset": &FieldTypeSupport{"TestScriptRulesetComponent", true, false},
		"test": &FieldTypeSupport{"TestScriptTestComponent", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"teardown": &FieldTypeSupport{"TestScriptTeardownComponent", false, true},

	}
}
