
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *RequestGroup) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *RequestGroup) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "basedon":
		 return fhirVal.BasedOn, true
 	case "groupidentifier":
		 return fhirVal.GroupIdentifier, true
 	case "note":
		 return fhirVal.Note, true
 	case "action":
		 return fhirVal.Action, true
 	case "replaces":
		 return fhirVal.Replaces, true
 	case "author":
		 return fhirVal.Author, true
 	case "reasonreference":
		 return fhirVal.ReasonReference, true
 	case "status":
		 return fhirVal.Status, true
 	case "priority":
		 return fhirVal.Priority, true
 	case "reasoncodeableconcept":
		 return fhirVal.ReasonCodeableConcept, true
 	case "definition":
		 return fhirVal.Definition, true
 	case "intent":
		 return fhirVal.Intent, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "context":
		 return fhirVal.Context, true
 	case "authoredon":
		 return fhirVal.AuthoredOn, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *RequestGroup) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"basedon": &FieldTypeSupport{"Reference", true, false},
 		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"action": &FieldTypeSupport{"RequestGroupActionComponent", true, false},
 		"replaces": &FieldTypeSupport{"Reference", true, false},
 		"author": &FieldTypeSupport{"Reference", false, true},
 		"reasonreference": &FieldTypeSupport{"Reference", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"priority": &FieldTypeSupport{"string", false, false},
 		"reasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
 		"definition": &FieldTypeSupport{"Reference", true, false},
 		"intent": &FieldTypeSupport{"string", false, false},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
 
	 }
 }
 
 func (fhirVal *RequestGroup) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 