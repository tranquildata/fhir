
package models


func (fhirVal *Encounter) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Encounter) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *Encounter) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"hospitalization": &FieldTypeSupport{"EncounterHospitalizationComponent", false, true},
		"location": &FieldTypeSupport{"EncounterLocationComponent", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", true, false},
		"priority": &FieldTypeSupport{"CodeableConcept", false, true},
		"appointment": &FieldTypeSupport{"Reference", false, true},
		"account": &FieldTypeSupport{"Reference", true, false},
		"participant": &FieldTypeSupport{"EncounterParticipantComponent", true, false},
		"period": &FieldTypeSupport{"Period", false, true},
		"length": &FieldTypeSupport{"Quantity", false, true},
		"diagnosis": &FieldTypeSupport{"EncounterDiagnosisComponent", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"statushistory": &FieldTypeSupport{"EncounterStatusHistoryComponent", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"incomingreferral": &FieldTypeSupport{"Reference", true, false},
		"partof": &FieldTypeSupport{"Reference", false, true},
		"serviceprovider": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"class": &FieldTypeSupport{"Coding", false, true},
		"episodeofcare": &FieldTypeSupport{"Reference", true, false},
		"reason": &FieldTypeSupport{"CodeableConcept", true, false},
		"classhistory": &FieldTypeSupport{"EncounterClassHistoryComponent", true, false},

	}
}
