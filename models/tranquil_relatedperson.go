
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *RelatedPerson) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *RelatedPerson) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "active":
		 return fhirVal.Active, true
 	case "relationship":
		 return fhirVal.Relationship, true
 	case "address":
		 return fhirVal.Address, true
 	case "period":
		 return fhirVal.Period, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "name":
		 return fhirVal.Name, true
 	case "telecom":
		 return fhirVal.Telecom, true
 	case "gender":
		 return fhirVal.Gender, true
 	case "birthdate":
		 return fhirVal.BirthDate, true
 	case "photo":
		 return fhirVal.Photo, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *RelatedPerson) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"active": &FieldTypeSupport{"bool", false, true},
 		"relationship": &FieldTypeSupport{"CodeableConcept", false, true},
 		"address": &FieldTypeSupport{"Address", true, false},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"name": &FieldTypeSupport{"HumanName", true, false},
 		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
 		"gender": &FieldTypeSupport{"string", false, false},
 		"birthdate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"photo": &FieldTypeSupport{"Attachment", true, false},
 
	 }
 }
 
 func (fhirVal *RelatedPerson) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 