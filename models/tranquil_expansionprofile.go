
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ExpansionProfile) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ExpansionProfile) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "excludenested":
		 return fhirVal.ExcludeNested, true
 	case "version":
		 return fhirVal.Version, true
 	case "status":
		 return fhirVal.Status, true
 	case "publisher":
		 return fhirVal.Publisher, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "includedesignations":
		 return fhirVal.IncludeDesignations, true
 	case "designation":
		 return fhirVal.Designation, true
 	case "displaylanguage":
		 return fhirVal.DisplayLanguage, true
 	case "name":
		 return fhirVal.Name, true
 	case "experimental":
		 return fhirVal.Experimental, true
 	case "usecontext":
		 return fhirVal.UseContext, true
 	case "activeonly":
		 return fhirVal.ActiveOnly, true
 	case "excludenotforui":
		 return fhirVal.ExcludeNotForUI, true
 	case "excludepostcoordinated":
		 return fhirVal.ExcludePostCoordinated, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "limitedexpansion":
		 return fhirVal.LimitedExpansion, true
 	case "url":
		 return fhirVal.Url, true
 	case "date":
		 return fhirVal.Date, true
 	case "description":
		 return fhirVal.Description, true
 	case "fixedversion":
		 return fhirVal.FixedVersion, true
 	case "excludedsystem":
		 return fhirVal.ExcludedSystem, true
 	case "includedefinition":
		 return fhirVal.IncludeDefinition, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ExpansionProfile) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"excludenested": &FieldTypeSupport{"bool", false, true},
 		"version": &FieldTypeSupport{"string", false, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"publisher": &FieldTypeSupport{"string", false, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"includedesignations": &FieldTypeSupport{"bool", false, true},
 		"designation": &FieldTypeSupport{"ExpansionProfileDesignationComponent", false, true},
 		"displaylanguage": &FieldTypeSupport{"string", false, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"experimental": &FieldTypeSupport{"bool", false, true},
 		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
 		"activeonly": &FieldTypeSupport{"bool", false, true},
 		"excludenotforui": &FieldTypeSupport{"bool", false, true},
 		"excludepostcoordinated": &FieldTypeSupport{"bool", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"limitedexpansion": &FieldTypeSupport{"bool", false, true},
 		"url": &FieldTypeSupport{"string", false, false},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"description": &FieldTypeSupport{"string", false, false},
 		"fixedversion": &FieldTypeSupport{"ExpansionProfileFixedVersionComponent", true, false},
 		"excludedsystem": &FieldTypeSupport{"ExpansionProfileExcludedSystemComponent", false, true},
 		"includedefinition": &FieldTypeSupport{"bool", false, true},
 
	 }
 }
 
 func (fhirVal *ExpansionProfile) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 