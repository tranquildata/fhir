
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *Immunization) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *Immunization) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "notgiven":
		 return fhirVal.NotGiven, true
 	case "vaccinecode":
		 return fhirVal.VaccineCode, true
 	case "site":
		 return fhirVal.Site, true
 	case "route":
		 return fhirVal.Route, true
 	case "practitioner":
		 return fhirVal.Practitioner, true
 	case "reaction":
		 return fhirVal.Reaction, true
 	case "vaccinationprotocol":
		 return fhirVal.VaccinationProtocol, true
 	case "primarysource":
		 return fhirVal.PrimarySource, true
 	case "dosequantity":
		 return fhirVal.DoseQuantity, true
 	case "status":
		 return fhirVal.Status, true
 	case "encounter":
		 return fhirVal.Encounter, true
 	case "reportorigin":
		 return fhirVal.ReportOrigin, true
 	case "location":
		 return fhirVal.Location, true
 	case "explanation":
		 return fhirVal.Explanation, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "date":
		 return fhirVal.Date, true
 	case "manufacturer":
		 return fhirVal.Manufacturer, true
 	case "lotnumber":
		 return fhirVal.LotNumber, true
 	case "expirationdate":
		 return fhirVal.ExpirationDate, true
 	case "note":
		 return fhirVal.Note, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *Immunization) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"notgiven": &FieldTypeSupport{"bool", false, true},
 		"vaccinecode": &FieldTypeSupport{"CodeableConcept", false, true},
 		"site": &FieldTypeSupport{"CodeableConcept", false, true},
 		"route": &FieldTypeSupport{"CodeableConcept", false, true},
 		"practitioner": &FieldTypeSupport{"ImmunizationPractitionerComponent", true, false},
 		"reaction": &FieldTypeSupport{"ImmunizationReactionComponent", true, false},
 		"vaccinationprotocol": &FieldTypeSupport{"ImmunizationVaccinationProtocolComponent", true, false},
 		"primarysource": &FieldTypeSupport{"bool", false, true},
 		"dosequantity": &FieldTypeSupport{"Quantity", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"encounter": &FieldTypeSupport{"Reference", false, true},
 		"reportorigin": &FieldTypeSupport{"CodeableConcept", false, true},
 		"location": &FieldTypeSupport{"Reference", false, true},
 		"explanation": &FieldTypeSupport{"ImmunizationExplanationComponent", false, true},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"manufacturer": &FieldTypeSupport{"Reference", false, true},
 		"lotnumber": &FieldTypeSupport{"string", false, false},
 		"expirationdate": &FieldTypeSupport{"FHIRDateTime", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 
	 }
 }
 
 func (fhirVal *Immunization) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 