
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ResearchSubject) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ResearchSubject) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "individual":
		 return fhirVal.Individual, true
 	case "assignedarm":
		 return fhirVal.AssignedArm, true
 	case "actualarm":
		 return fhirVal.ActualArm, true
 	case "consent":
		 return fhirVal.Consent, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "period":
		 return fhirVal.Period, true
 	case "study":
		 return fhirVal.Study, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ResearchSubject) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"individual": &FieldTypeSupport{"Reference", false, true},
 		"assignedarm": &FieldTypeSupport{"string", false, false},
 		"actualarm": &FieldTypeSupport{"string", false, false},
 		"consent": &FieldTypeSupport{"Reference", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"study": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *ResearchSubject) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 