
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *EligibilityRequest) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *EligibilityRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "benefitcategory":
		 return fhirVal.BenefitCategory, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "serviceddate":
		 return fhirVal.ServicedDate, true
 	case "organization":
		 return fhirVal.Organization, true
 	case "coverage":
		 return fhirVal.Coverage, true
 	case "businessarrangement":
		 return fhirVal.BusinessArrangement, true
 	case "benefitsubcategory":
		 return fhirVal.BenefitSubCategory, true
 	case "status":
		 return fhirVal.Status, true
 	case "priority":
		 return fhirVal.Priority, true
 	case "created":
		 return fhirVal.Created, true
 	case "enterer":
		 return fhirVal.Enterer, true
 	case "servicedperiod":
		 return fhirVal.ServicedPeriod, true
 	case "provider":
		 return fhirVal.Provider, true
 	case "insurer":
		 return fhirVal.Insurer, true
 	case "facility":
		 return fhirVal.Facility, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *EligibilityRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"benefitcategory": &FieldTypeSupport{"CodeableConcept", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"serviceddate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"organization": &FieldTypeSupport{"Reference", false, true},
 		"coverage": &FieldTypeSupport{"Reference", false, true},
 		"businessarrangement": &FieldTypeSupport{"string", false, false},
 		"benefitsubcategory": &FieldTypeSupport{"CodeableConcept", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"priority": &FieldTypeSupport{"CodeableConcept", false, true},
 		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"enterer": &FieldTypeSupport{"Reference", false, true},
 		"servicedperiod": &FieldTypeSupport{"Period", false, true},
 		"provider": &FieldTypeSupport{"Reference", false, true},
 		"insurer": &FieldTypeSupport{"Reference", false, true},
 		"facility": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *EligibilityRequest) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 