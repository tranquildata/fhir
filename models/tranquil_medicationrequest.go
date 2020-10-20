
package models


func (fhirVal *MedicationRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *MedicationRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "subject":
		return fhirVal.Subject, true
	case "context":
		return fhirVal.Context, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "dosageinstruction":
		return fhirVal.DosageInstruction, true
	case "identifier":
		return fhirVal.Identifier, true
	case "intent":
		return fhirVal.Intent, true
	case "priority":
		return fhirVal.Priority, true
	case "priorprescription":
		return fhirVal.PriorPrescription, true
	case "detectedissue":
		return fhirVal.DetectedIssue, true
	case "supportinginformation":
		return fhirVal.SupportingInformation, true
	case "eventhistory":
		return fhirVal.EventHistory, true
	case "definition":
		return fhirVal.Definition, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "medicationcodeableconcept":
		return fhirVal.MedicationCodeableConcept, true
	case "note":
		return fhirVal.Note, true
	case "dispenserequest":
		return fhirVal.DispenseRequest, true
	case "substitution":
		return fhirVal.Substitution, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "status":
		return fhirVal.Status, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "recorder":
		return fhirVal.Recorder, true
	case "category":
		return fhirVal.Category, true
	case "medicationreference":
		return fhirVal.MedicationReference, true
	case "requester":
		return fhirVal.Requester, true

	default:
		return nil, false
	}
}

func (fhirVal *MedicationRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"subject": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"dosageinstruction": &FieldTypeSupport{"Dosage", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"intent": &FieldTypeSupport{"string", false, false},
		"priority": &FieldTypeSupport{"string", false, false},
		"priorprescription": &FieldTypeSupport{"Reference", false, true},
		"detectedissue": &FieldTypeSupport{"Reference", true, false},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"eventhistory": &FieldTypeSupport{"Reference", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"dispenserequest": &FieldTypeSupport{"MedicationRequestDispenseRequestComponent", false, true},
		"substitution": &FieldTypeSupport{"MedicationRequestSubstitutionComponent", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"recorder": &FieldTypeSupport{"Reference", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"requester": &FieldTypeSupport{"MedicationRequestRequesterComponent", false, true},

	}
}
