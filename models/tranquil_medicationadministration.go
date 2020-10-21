
package models


func (fhirVal *MedicationAdministration) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *MedicationAdministration) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "performer":
		return fhirVal.Performer, true
	case "reasonnotgiven":
		return fhirVal.ReasonNotGiven, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "prescription":
		return fhirVal.Prescription, true
	case "identifier":
		return fhirVal.Identifier, true
	case "partof":
		return fhirVal.PartOf, true
	case "status":
		return fhirVal.Status, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "note":
		return fhirVal.Note, true
	case "dosage":
		return fhirVal.Dosage, true
	case "device":
		return fhirVal.Device, true
	case "medicationcodeableconcept":
		return fhirVal.MedicationCodeableConcept, true
	case "medicationreference":
		return fhirVal.MedicationReference, true
	case "supportinginformation":
		return fhirVal.SupportingInformation, true
	case "notgiven":
		return fhirVal.NotGiven, true
	case "category":
		return fhirVal.Category, true
	case "context":
		return fhirVal.Context, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "eventhistory":
		return fhirVal.EventHistory, true
	case "definition":
		return fhirVal.Definition, true
	case "subject":
		return fhirVal.Subject, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true

	default:
		return nil, false
	}
}

func (fhirVal *MedicationAdministration) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"performer": &FieldTypeSupport{"MedicationAdministrationPerformerComponent", true, false},
		"reasonnotgiven": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"dosage": &FieldTypeSupport{"MedicationAdministrationDosageComponent", false, true},
		"device": &FieldTypeSupport{"Reference", true, false},
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"notgiven": &FieldTypeSupport{"bool", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"eventhistory": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},

	}
}
