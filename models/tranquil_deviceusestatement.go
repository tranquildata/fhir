
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *DeviceUseStatement) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *DeviceUseStatement) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "timingtiming":
		 return fhirVal.TimingTiming, true
 	case "timingperiod":
		 return fhirVal.TimingPeriod, true
 	case "device":
		 return fhirVal.Device, true
 	case "bodysite":
		 return fhirVal.BodySite, true
 	case "note":
		 return fhirVal.Note, true
 	case "status":
		 return fhirVal.Status, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "whenused":
		 return fhirVal.WhenUsed, true
 	case "source":
		 return fhirVal.Source, true
 	case "indication":
		 return fhirVal.Indication, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "timingdatetime":
		 return fhirVal.TimingDateTime, true
 	case "recordedon":
		 return fhirVal.RecordedOn, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *DeviceUseStatement) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"timingtiming": &FieldTypeSupport{"Timing", false, true},
 		"timingperiod": &FieldTypeSupport{"Period", false, true},
 		"device": &FieldTypeSupport{"Reference", false, true},
 		"bodysite": &FieldTypeSupport{"CodeableConcept", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"whenused": &FieldTypeSupport{"Period", false, true},
 		"source": &FieldTypeSupport{"Reference", false, true},
 		"indication": &FieldTypeSupport{"CodeableConcept", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"timingdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"recordedon": &FieldTypeSupport{"FHIRDateTime", false, true},
 
	 }
 }
 
 func (fhirVal *DeviceUseStatement) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 