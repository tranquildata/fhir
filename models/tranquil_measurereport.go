
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *MeasureReport) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *MeasureReport) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "measure":
		 return fhirVal.Measure, true
 	case "date":
		 return fhirVal.Date, true
 	case "evaluatedresources":
		 return fhirVal.EvaluatedResources, true
 	case "status":
		 return fhirVal.Status, true
 	case "type":
		 return fhirVal.Type, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "reportingorganization":
		 return fhirVal.ReportingOrganization, true
 	case "period":
		 return fhirVal.Period, true
 	case "group":
		 return fhirVal.Group, true
 	case "identifier":
		 return fhirVal.Identifier, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *MeasureReport) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"measure": &FieldTypeSupport{"Reference", false, true},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"evaluatedresources": &FieldTypeSupport{"Reference", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"type": &FieldTypeSupport{"string", false, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"reportingorganization": &FieldTypeSupport{"Reference", false, true},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"group": &FieldTypeSupport{"MeasureReportGroupComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 
	 }
 }
 
 func (fhirVal *MeasureReport) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 