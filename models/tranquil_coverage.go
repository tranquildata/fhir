
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Coverage) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Coverage) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "relationship":
		 return fhirVal.Relationship, true
 	case "period":
		 return fhirVal.Period, true
 	case "grouping":
		 return fhirVal.Grouping, true
 	case "dependent":
		 return fhirVal.Dependent, true
 	case "order":
		 return fhirVal.Order, true
 	case "subscriber":
		 return fhirVal.Subscriber, true
 	case "subscriberid":
		 return fhirVal.SubscriberId, true
 	case "beneficiary":
		 return fhirVal.Beneficiary, true
 	case "type":
		 return fhirVal.Type, true
 	case "policyholder":
		 return fhirVal.PolicyHolder, true
 	case "payor":
		 return fhirVal.Payor, true
 	case "network":
		 return fhirVal.Network, true
 	case "contract":
		 return fhirVal.Contract, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "sequence":
		 return fhirVal.Sequence, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Coverage) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"relationship": &FieldTypeSupport{"CodeableConcept", false, true},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"grouping": &FieldTypeSupport{"CoverageGroupComponent", false, true},
 		"dependent": &FieldTypeSupport{"string", false, false},
 		"order": &FieldTypeSupport{"uint32", false, true},
 		"subscriber": &FieldTypeSupport{"Reference", false, true},
 		"subscriberid": &FieldTypeSupport{"string", false, false},
 		"beneficiary": &FieldTypeSupport{"Reference", false, true},
 		"type": &FieldTypeSupport{"CodeableConcept", false, true},
 		"policyholder": &FieldTypeSupport{"Reference", false, true},
 		"payor": &FieldTypeSupport{"Reference", true, false},
 		"network": &FieldTypeSupport{"string", false, false},
 		"contract": &FieldTypeSupport{"Reference", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"sequence": &FieldTypeSupport{"string", false, false},
 
	 }
 }
 
 func (fhirVal *Coverage) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 