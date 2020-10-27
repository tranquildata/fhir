
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *List) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *List) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "note":
		 return fhirVal.Note, true
 	case "mode":
		 return fhirVal.Mode, true
 	case "title":
		 return fhirVal.Title, true
 	case "code":
		 return fhirVal.Code, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "date":
		 return fhirVal.Date, true
 	case "source":
		 return fhirVal.Source, true
 	case "orderedby":
		 return fhirVal.OrderedBy, true
 	case "entry":
		 return fhirVal.Entry, true
 	case "emptyreason":
		 return fhirVal.EmptyReason, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "encounter":
		 return fhirVal.Encounter, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *List) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"mode": &FieldTypeSupport{"string", false, false},
 		"title": &FieldTypeSupport{"string", false, false},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"source": &FieldTypeSupport{"Reference", false, true},
 		"orderedby": &FieldTypeSupport{"CodeableConcept", false, true},
 		"entry": &FieldTypeSupport{"ListEntryComponent", true, false},
 		"emptyreason": &FieldTypeSupport{"CodeableConcept", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"encounter": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *List) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 