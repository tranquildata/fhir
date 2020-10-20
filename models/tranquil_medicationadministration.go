
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
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "eventhistory":
		return fhirVal.EventHistory, true
	case "medicationcodeableconcept":
		return fhirVal.MedicationCodeableConcept, true
	case "reasonnotgiven":
		return fhirVal.ReasonNotGiven, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true
	case "definition":
		return fhirVal.Definition, true
	case "status":
		return fhirVal.Status, true
	case "category":
		return fhirVal.Category, true
	case "medicationreference":
		return fhirVal.MedicationReference, true
	case "supportinginformation":
		return fhirVal.SupportingInformation, true
	case "performer":
		return fhirVal.Performer, true
	case "device":
		return fhirVal.Device, true
	case "identifier":
		return fhirVal.Identifier, true
	case "dosage":
		return fhirVal.Dosage, true
	case "subject":
		return fhirVal.Subject, true
	case "context":
		return fhirVal.Context, true
	case "notgiven":
		return fhirVal.NotGiven, true
	case "prescription":
		return fhirVal.Prescription, true
	case "partof":
		return fhirVal.PartOf, true

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
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"eventhistory": &FieldTypeSupport{"Reference", true, false},
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasonnotgiven": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"performer": &FieldTypeSupport{"MedicationAdministrationPerformerComponent", true, false},
		"device": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"dosage": &FieldTypeSupport{"MedicationAdministrationDosageComponent", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"notgiven": &FieldTypeSupport{"bool", false, true},
		"prescription": &FieldTypeSupport{"Reference", false, true},
		"partof": &FieldTypeSupport{"Reference", true, false},

	}
}
