
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ImagingManifest) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ImagingManifest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "authoringtime":
		 return fhirVal.AuthoringTime, true
 	case "author":
		 return fhirVal.Author, true
 	case "description":
		 return fhirVal.Description, true
 	case "study":
		 return fhirVal.Study, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ImagingManifest) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"authoringtime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"author": &FieldTypeSupport{"Reference", false, true},
 		"description": &FieldTypeSupport{"string", false, false},
 		"study": &FieldTypeSupport{"ImagingManifestStudyComponent", true, false},
 
	 }
 }
 
 func (fhirVal *ImagingManifest) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 