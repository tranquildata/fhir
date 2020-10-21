
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
	case "definition":
		return fhirVal.Definition, true
	case "supportinginformation":
		return fhirVal.SupportingInformation, true
	case "authoredon":
		return fhirVal.AuthoredOn, true
	case "priorprescription":
		return fhirVal.PriorPrescription, true
	case "dispenserequest":
		return fhirVal.DispenseRequest, true
	case "groupidentifier":
		return fhirVal.GroupIdentifier, true
	case "intent":
		return fhirVal.Intent, true
	case "medicationreference":
		return fhirVal.MedicationReference, true
	case "context":
		return fhirVal.Context, true
	case "note":
		return fhirVal.Note, true
	case "recorder":
		return fhirVal.Recorder, true
	case "dosageinstruction":
		return fhirVal.DosageInstruction, true
	case "eventhistory":
		return fhirVal.EventHistory, true
	case "identifier":
		return fhirVal.Identifier, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "status":
		return fhirVal.Status, true
	case "priority":
		return fhirVal.Priority, true
	case "subject":
		return fhirVal.Subject, true
	case "substitution":
		return fhirVal.Substitution, true
	case "detectedissue":
		return fhirVal.DetectedIssue, true
	case "category":
		return fhirVal.Category, true
	case "medicationcodeableconcept":
		return fhirVal.MedicationCodeableConcept, true
	case "requester":
		return fhirVal.Requester, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "reasonreference":
		return fhirVal.ReasonReference, true

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
		"definition": &FieldTypeSupport{"Reference", true, false},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"authoredon": &FieldTypeSupport{"FHIRDateTime", false, true},
		"priorprescription": &FieldTypeSupport{"Reference", false, true},
		"dispenserequest": &FieldTypeSupport{"MedicationRequestDispenseRequestComponent", false, true},
		"groupidentifier": &FieldTypeSupport{"Identifier", false, true},
		"intent": &FieldTypeSupport{"string", false, false},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"recorder": &FieldTypeSupport{"Reference", false, true},
		"dosageinstruction": &FieldTypeSupport{"Dosage", true, false},
		"eventhistory": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"priority": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"substitution": &FieldTypeSupport{"MedicationRequestSubstitutionComponent", false, true},
		"detectedissue": &FieldTypeSupport{"Reference", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"requester": &FieldTypeSupport{"MedicationRequestRequesterComponent", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},

	}
}
