
package models


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
	case "status":
		return fhirVal.Status, true
	case "primarysource":
		return fhirVal.PrimarySource, true
	case "location":
		return fhirVal.Location, true
	case "lotnumber":
		return fhirVal.LotNumber, true
	case "expirationdate":
		return fhirVal.ExpirationDate, true
	case "practitioner":
		return fhirVal.Practitioner, true
	case "identifier":
		return fhirVal.Identifier, true
	case "patient":
		return fhirVal.Patient, true
	case "route":
		return fhirVal.Route, true
	case "dosequantity":
		return fhirVal.DoseQuantity, true
	case "explanation":
		return fhirVal.Explanation, true
	case "vaccinationprotocol":
		return fhirVal.VaccinationProtocol, true
	case "vaccinecode":
		return fhirVal.VaccineCode, true
	case "reportorigin":
		return fhirVal.ReportOrigin, true
	case "manufacturer":
		return fhirVal.Manufacturer, true
	case "site":
		return fhirVal.Site, true
	case "note":
		return fhirVal.Note, true
	case "reaction":
		return fhirVal.Reaction, true
	case "notgiven":
		return fhirVal.NotGiven, true
	case "encounter":
		return fhirVal.Encounter, true
	case "date":
		return fhirVal.Date, true

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
		"status": &FieldTypeSupport{"string", false, false},
		"primarysource": &FieldTypeSupport{"bool", false, true},
		"location": &FieldTypeSupport{"Reference", false, true},
		"lotnumber": &FieldTypeSupport{"string", false, false},
		"expirationdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"practitioner": &FieldTypeSupport{"ImmunizationPractitionerComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"route": &FieldTypeSupport{"CodeableConcept", false, true},
		"dosequantity": &FieldTypeSupport{"Quantity", false, true},
		"explanation": &FieldTypeSupport{"ImmunizationExplanationComponent", false, true},
		"vaccinationprotocol": &FieldTypeSupport{"ImmunizationVaccinationProtocolComponent", true, false},
		"vaccinecode": &FieldTypeSupport{"CodeableConcept", false, true},
		"reportorigin": &FieldTypeSupport{"CodeableConcept", false, true},
		"manufacturer": &FieldTypeSupport{"Reference", false, true},
		"site": &FieldTypeSupport{"CodeableConcept", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"reaction": &FieldTypeSupport{"ImmunizationReactionComponent", true, false},
		"notgiven": &FieldTypeSupport{"bool", false, true},
		"encounter": &FieldTypeSupport{"Reference", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
