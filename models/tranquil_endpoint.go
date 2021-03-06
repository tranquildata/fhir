
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Endpoint) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Endpoint) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "payloadtype":
		 return fhirVal.PayloadType, true
 	case "payloadmimetype":
		 return fhirVal.PayloadMimeType, true
 	case "address":
		 return fhirVal.Address, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "connectiontype":
		 return fhirVal.ConnectionType, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "period":
		 return fhirVal.Period, true
 	case "status":
		 return fhirVal.Status, true
 	case "name":
		 return fhirVal.Name, true
 	case "managingorganization":
		 return fhirVal.ManagingOrganization, true
 	case "header":
		 return fhirVal.Header, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Endpoint) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"payloadtype": &FieldTypeSupport{"CodeableConcept", true, false},
 		"payloadmimetype": &FieldTypeSupport{"string", true, false},
 		"address": &FieldTypeSupport{"string", false, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"connectiontype": &FieldTypeSupport{"Coding", false, true},
 		"contact": &FieldTypeSupport{"ContactPoint", true, false},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"managingorganization": &FieldTypeSupport{"Reference", false, true},
 		"header": &FieldTypeSupport{"string", true, false},
 
	 }
 }
 
 func (fhirVal *Endpoint) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 