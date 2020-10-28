
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Device) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Device) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "udi":
		 return fhirVal.Udi, true
 	case "type":
		 return fhirVal.Type, true
 	case "lotnumber":
		 return fhirVal.LotNumber, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "status":
		 return fhirVal.Status, true
 	case "manufacturer":
		 return fhirVal.Manufacturer, true
 	case "expirationdate":
		 return fhirVal.ExpirationDate, true
 	case "owner":
		 return fhirVal.Owner, true
 	case "url":
		 return fhirVal.Url, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "manufacturedate":
		 return fhirVal.ManufactureDate, true
 	case "version":
		 return fhirVal.Version, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "safety":
		 return fhirVal.Safety, true
 	case "model":
		 return fhirVal.Model, true
 	case "location":
		 return fhirVal.Location, true
 	case "note":
		 return fhirVal.Note, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Device) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"udi": &FieldTypeSupport{"DeviceUdiComponent", false, true},
 		"type": &FieldTypeSupport{"CodeableConcept", false, true},
 		"lotnumber": &FieldTypeSupport{"string", false, false},
 		"contact": &FieldTypeSupport{"ContactPoint", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"manufacturer": &FieldTypeSupport{"string", false, false},
 		"expirationdate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"owner": &FieldTypeSupport{"Reference", false, true},
 		"url": &FieldTypeSupport{"string", false, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"manufacturedate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"version": &FieldTypeSupport{"string", false, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"safety": &FieldTypeSupport{"CodeableConcept", true, false},
 		"model": &FieldTypeSupport{"string", false, false},
 		"location": &FieldTypeSupport{"Reference", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 
	 }
 }
 
 func (fhirVal *Device) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 