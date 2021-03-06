
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *EligibilityResponse) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *EligibilityResponse) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "insurer":
		 return fhirVal.Insurer, true
 	case "insurance":
		 return fhirVal.Insurance, true
 	case "form":
		 return fhirVal.Form, true
 	case "error":
		 return fhirVal.Error, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "requestorganization":
		 return fhirVal.RequestOrganization, true
 	case "disposition":
		 return fhirVal.Disposition, true
 	case "request":
		 return fhirVal.Request, true
 	case "outcome":
		 return fhirVal.Outcome, true
 	case "inforce":
		 return fhirVal.Inforce, true
 	case "status":
		 return fhirVal.Status, true
 	case "created":
		 return fhirVal.Created, true
 	case "requestprovider":
		 return fhirVal.RequestProvider, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *EligibilityResponse) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"insurer": &FieldTypeSupport{"Reference", false, true},
 		"insurance": &FieldTypeSupport{"EligibilityResponseInsuranceComponent", true, false},
 		"form": &FieldTypeSupport{"CodeableConcept", false, true},
 		"error": &FieldTypeSupport{"EligibilityResponseErrorsComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"requestorganization": &FieldTypeSupport{"Reference", false, true},
 		"disposition": &FieldTypeSupport{"string", false, false},
 		"request": &FieldTypeSupport{"Reference", false, true},
 		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
 		"inforce": &FieldTypeSupport{"bool", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"requestprovider": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *EligibilityResponse) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 