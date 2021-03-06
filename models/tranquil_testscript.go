
 package models
 
 import (
	 "fmt"
 )
 
 
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
 	case "status":
		 return fhirVal.Status, true
 	case "copyright":
		 return fhirVal.Copyright, true
 	case "version":
		 return fhirVal.Version, true
 	case "description":
		 return fhirVal.Description, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "profile":
		 return fhirVal.Profile, true
 	case "setup":
		 return fhirVal.Setup, true
 	case "teardown":
		 return fhirVal.Teardown, true
 	case "name":
		 return fhirVal.Name, true
 	case "title":
		 return fhirVal.Title, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "purpose":
		 return fhirVal.Purpose, true
 	case "metadata":
		 return fhirVal.Metadata, true
 	case "fixture":
		 return fhirVal.Fixture, true
 	case "rule":
		 return fhirVal.Rule, true
 	case "test":
		 return fhirVal.Test, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "date":
		 return fhirVal.Date, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "origin":
		 return fhirVal.Origin, true
 	case "destination":
		 return fhirVal.Destination, true
 	case "variable":
		 return fhirVal.Variable, true
 	case "ruleset":
		 return fhirVal.Ruleset, true
 	case "url":
		 return fhirVal.Url, true
 
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
 		"status": &FieldTypeSupport{"string", false, false},
 		"copyright": &FieldTypeSupport{"string", false, false},
 		"version": &FieldTypeSupport{"string", false, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"profile": &FieldTypeSupport{"Reference", true, false},
 		"setup": &FieldTypeSupport{"TestScriptSetupComponent", false, true},
 		"teardown": &FieldTypeSupport{"TestScriptTeardownComponent", false, true},
 		"name": &FieldTypeSupport{"string", false, false},
 		"title": &FieldTypeSupport{"string", false, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"purpose": &FieldTypeSupport{"string", false, false},
 		"metadata": &FieldTypeSupport{"TestScriptMetadataComponent", false, true},
 		"fixture": &FieldTypeSupport{"TestScriptFixtureComponent", true, false},
 		"rule": &FieldTypeSupport{"TestScriptRuleComponent", true, false},
 		"test": &FieldTypeSupport{"TestScriptTestComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"origin": &FieldTypeSupport{"TestScriptOriginComponent", true, false},
 		"destination": &FieldTypeSupport{"TestScriptDestinationComponent", true, false},
 		"variable": &FieldTypeSupport{"TestScriptVariableComponent", true, false},
 		"ruleset": &FieldTypeSupport{"TestScriptRulesetComponent", true, false},
 		"url": &FieldTypeSupport{"string", false, false},
 
	 }
 }
 
 func (fhirVal *TestScript) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 