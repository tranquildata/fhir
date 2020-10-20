
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
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "reportorigin":
		return fhirVal.ReportOrigin, true
	case "dosequantity":
		return fhirVal.DoseQuantity, true
	case "explanation":
		return fhirVal.Explanation, true
	case "reaction":
		return fhirVal.Reaction, true
	case "vaccinecode":
		return fhirVal.VaccineCode, true
	case "date":
		return fhirVal.Date, true
	case "primarysource":
		return fhirVal.PrimarySource, true
	case "location":
		return fhirVal.Location, true
	case "route":
		return fhirVal.Route, true
	case "note":
		return fhirVal.Note, true
	case "vaccinationprotocol":
		return fhirVal.VaccinationProtocol, true
	case "notgiven":
		return fhirVal.NotGiven, true
	case "lotnumber":
		return fhirVal.LotNumber, true
	case "expirationdate":
		return fhirVal.ExpirationDate, true
	case "site":
		return fhirVal.Site, true
	case "practitioner":
		return fhirVal.Practitioner, true
	case "patient":
		return fhirVal.Patient, true
	case "encounter":
		return fhirVal.Encounter, true
	case "manufacturer":
		return fhirVal.Manufacturer, true

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
		"status": &FieldTypeSupport{"string", false, false},
		"reportorigin": &FieldTypeSupport{"CodeableConcept", false, true},
		"dosequantity": &FieldTypeSupport{"Quantity", false, true},
		"explanation": &FieldTypeSupport{"ImmunizationExplanationComponent", false, true},
		"reaction": &FieldTypeSupport{"ImmunizationReactionComponent", true, false},
		"vaccinecode": &FieldTypeSupport{"CodeableConcept", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"primarysource": &FieldTypeSupport{"bool", false, true},
		"location": &FieldTypeSupport{"Reference", false, true},
		"route": &FieldTypeSupport{"CodeableConcept", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"vaccinationprotocol": &FieldTypeSupport{"ImmunizationVaccinationProtocolComponent", true, false},
		"notgiven": &FieldTypeSupport{"bool", false, true},
		"lotnumber": &FieldTypeSupport{"string", false, false},
		"expirationdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"site": &FieldTypeSupport{"CodeableConcept", false, true},
		"practitioner": &FieldTypeSupport{"ImmunizationPractitionerComponent", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"encounter": &FieldTypeSupport{"Reference", false, true},
		"manufacturer": &FieldTypeSupport{"Reference", false, true},

	}
}
