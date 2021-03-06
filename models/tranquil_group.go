
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Group) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Group) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "name":
		 return fhirVal.Name, true
 	case "quantity":
		 return fhirVal.Quantity, true
 	case "member":
		 return fhirVal.Member, true
 	case "active":
		 return fhirVal.Active, true
 	case "type":
		 return fhirVal.Type, true
 	case "code":
		 return fhirVal.Code, true
 	case "characteristic":
		 return fhirVal.Characteristic, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "actual":
		 return fhirVal.Actual, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Group) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"name": &FieldTypeSupport{"string", false, false},
 		"quantity": &FieldTypeSupport{"uint32", false, true},
 		"member": &FieldTypeSupport{"GroupMemberComponent", true, false},
 		"active": &FieldTypeSupport{"bool", false, true},
 		"type": &FieldTypeSupport{"string", false, false},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"characteristic": &FieldTypeSupport{"GroupCharacteristicComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"actual": &FieldTypeSupport{"bool", false, true},
 
	 }
 }
 
 func (fhirVal *Group) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 