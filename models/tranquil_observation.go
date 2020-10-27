
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Observation) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Observation) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "valuerange":
		 return fhirVal.ValueRange, true
 	case "valuedatetime":
		 return fhirVal.ValueDateTime, true
 	case "specimen":
		 return fhirVal.Specimen, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "context":
		 return fhirVal.Context, true
 	case "effectivedatetime":
		 return fhirVal.EffectiveDateTime, true
 	case "component":
		 return fhirVal.Component, true
 	case "effectiveperiod":
		 return fhirVal.EffectivePeriod, true
 	case "valuestring":
		 return fhirVal.ValueString, true
 	case "valueperiod":
		 return fhirVal.ValuePeriod, true
 	case "bodysite":
		 return fhirVal.BodySite, true
 	case "category":
		 return fhirVal.Category, true
 	case "valuesampleddata":
		 return fhirVal.ValueSampledData, true
 	case "comment":
		 return fhirVal.Comment, true
 	case "valueboolean":
		 return fhirVal.ValueBoolean, true
 	case "valuetime":
		 return fhirVal.ValueTime, true
 	case "method":
		 return fhirVal.Method, true
 	case "valueratio":
		 return fhirVal.ValueRatio, true
 	case "dataabsentreason":
		 return fhirVal.DataAbsentReason, true
 	case "status":
		 return fhirVal.Status, true
 	case "issued":
		 return fhirVal.Issued, true
 	case "valuequantity":
		 return fhirVal.ValueQuantity, true
 	case "related":
		 return fhirVal.Related, true
 	case "basedon":
		 return fhirVal.BasedOn, true
 	case "interpretation":
		 return fhirVal.Interpretation, true
 	case "referencerange":
		 return fhirVal.ReferenceRange, true
 	case "valuecodeableconcept":
		 return fhirVal.ValueCodeableConcept, true
 	case "valueattachment":
		 return fhirVal.ValueAttachment, true
 	case "device":
		 return fhirVal.Device, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "code":
		 return fhirVal.Code, true
 	case "performer":
		 return fhirVal.Performer, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Observation) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"valuerange": &FieldTypeSupport{"Range", false, true},
 		"valuedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"specimen": &FieldTypeSupport{"Reference", false, true},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"component": &FieldTypeSupport{"ObservationComponentComponent", true, false},
 		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
 		"valuestring": &FieldTypeSupport{"string", false, false},
 		"valueperiod": &FieldTypeSupport{"Period", false, true},
 		"bodysite": &FieldTypeSupport{"CodeableConcept", false, true},
 		"category": &FieldTypeSupport{"CodeableConcept", true, false},
 		"valuesampleddata": &FieldTypeSupport{"SampledData", false, true},
 		"comment": &FieldTypeSupport{"string", false, false},
 		"valueboolean": &FieldTypeSupport{"bool", false, true},
 		"valuetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"method": &FieldTypeSupport{"CodeableConcept", false, true},
 		"valueratio": &FieldTypeSupport{"Ratio", false, true},
 		"dataabsentreason": &FieldTypeSupport{"CodeableConcept", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"valuequantity": &FieldTypeSupport{"Quantity", false, true},
 		"related": &FieldTypeSupport{"ObservationRelatedComponent", true, false},
 		"basedon": &FieldTypeSupport{"Reference", true, false},
 		"interpretation": &FieldTypeSupport{"CodeableConcept", false, true},
 		"referencerange": &FieldTypeSupport{"ObservationReferenceRangeComponent", true, false},
 		"valuecodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
 		"valueattachment": &FieldTypeSupport{"Attachment", false, true},
 		"device": &FieldTypeSupport{"Reference", false, true},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"code": &FieldTypeSupport{"CodeableConcept", false, true},
 		"performer": &FieldTypeSupport{"Reference", true, false},
 
	 }
 }
 
 func (fhirVal *Observation) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 