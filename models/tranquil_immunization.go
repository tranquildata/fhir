
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
		"notgiven": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"site": &FieldTypeSupport{"CodeableConcept", false, true},
		"expirationdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"route": &FieldTypeSupport{"CodeableConcept", false, true},
		"explanation": &FieldTypeSupport{"ImmunizationExplanationComponent", false, true},
		"vaccinationprotocol": &FieldTypeSupport{"ImmunizationVaccinationProtocolComponent", true, false},
		"primarysource": &FieldTypeSupport{"bool", false, true},
		"reportorigin": &FieldTypeSupport{"CodeableConcept", false, true},
		"lotnumber": &FieldTypeSupport{"string", false, false},
		"dosequantity": &FieldTypeSupport{"Quantity", false, true},
		"practitioner": &FieldTypeSupport{"ImmunizationPractitionerComponent", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"reaction": &FieldTypeSupport{"ImmunizationReactionComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"encounter": &FieldTypeSupport{"Reference", false, true},
		"manufacturer": &FieldTypeSupport{"Reference", false, true},
		"location": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"vaccinecode": &FieldTypeSupport{"CodeableConcept", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},

	}
}
