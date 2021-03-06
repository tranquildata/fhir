
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Organization) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Organization) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "partof":
		 return fhirVal.PartOf, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "endpoint":
		 return fhirVal.Endpoint, true
 	case "active":
		 return fhirVal.Active, true
 	case "alias":
		 return fhirVal.Alias, true
 	case "telecom":
		 return fhirVal.Telecom, true
 	case "address":
		 return fhirVal.Address, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "type":
		 return fhirVal.Type, true
 	case "name":
		 return fhirVal.Name, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Organization) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"partof": &FieldTypeSupport{"Reference", false, true},
 		"contact": &FieldTypeSupport{"OrganizationContactComponent", true, false},
 		"endpoint": &FieldTypeSupport{"Reference", true, false},
 		"active": &FieldTypeSupport{"bool", false, true},
 		"alias": &FieldTypeSupport{"string", true, false},
 		"telecom": &FieldTypeSupport{"ContactPoint", true, false},
 		"address": &FieldTypeSupport{"Address", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"type": &FieldTypeSupport{"CodeableConcept", true, false},
 		"name": &FieldTypeSupport{"string", false, false},
 
	 }
 }
 
 func (fhirVal *Organization) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 