
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *DeviceMetric) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *DeviceMetric) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "source":
		 return fhirVal.Source, true
 	case "parent":
		 return fhirVal.Parent, true
 	case "color":
		 return fhirVal.Color, true
 	case "category":
		 return fhirVal.Category, true
 	case "calibration":
		 return fhirVal.Calibration, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "unit":
		 return fhirVal.Unit, true
 	case "operationalstatus":
		 return fhirVal.OperationalStatus, true
 	case "measurementperiod":
		 return fhirVal.MeasurementPeriod, true
 	case "type":
		 return fhirVal.Type, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *DeviceMetric) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"source": &FieldTypeSupport{"Reference", false, true},
 		"parent": &FieldTypeSupport{"Reference", false, true},
 		"color": &FieldTypeSupport{"string", false, false},
 		"category": &FieldTypeSupport{"string", false, false},
 		"calibration": &FieldTypeSupport{"DeviceMetricCalibrationComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", false, true},
 		"unit": &FieldTypeSupport{"CodeableConcept", false, true},
 		"operationalstatus": &FieldTypeSupport{"string", false, false},
 		"measurementperiod": &FieldTypeSupport{"Timing", false, true},
 		"type": &FieldTypeSupport{"CodeableConcept", false, true},
 
	 }
 }
 
 func (fhirVal *DeviceMetric) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 