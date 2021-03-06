
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *VisionPrescription) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *VisionPrescription) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "patient":
		 return fhirVal.Patient, true
 	case "encounter":
		 return fhirVal.Encounter, true
 	case "datewritten":
		 return fhirVal.DateWritten, true
 	case "dispense":
		 return fhirVal.Dispense, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "status":
		 return fhirVal.Status, true
 	case "prescriber":
		 return fhirVal.Prescriber, true
 	case "reasoncodeableconcept":
		 return fhirVal.ReasonCodeableConcept, true
 	case "reasonreference":
		 return fhirVal.ReasonReference, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *VisionPrescription) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"encounter": &FieldTypeSupport{"Reference", false, true},
 		"datewritten": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"dispense": &FieldTypeSupport{"VisionPrescriptionDispenseComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"prescriber": &FieldTypeSupport{"Reference", false, true},
 		"reasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
 		"reasonreference": &FieldTypeSupport{"Reference", false, true},
 
	 }
 }
 
 func (fhirVal *VisionPrescription) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 