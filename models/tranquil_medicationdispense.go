
package models


func (fhirVal *MedicationDispense) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *MedicationDispense) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "quantity":
		return fhirVal.Quantity, true
	case "whenprepared":
		return fhirVal.WhenPrepared, true
	case "whenhandedover":
		return fhirVal.WhenHandedOver, true
	case "receiver":
		return fhirVal.Receiver, true
	case "dosageinstruction":
		return fhirVal.DosageInstruction, true
	case "notdonereasoncodeableconcept":
		return fhirVal.NotDoneReasonCodeableConcept, true
	case "performer":
		return fhirVal.Performer, true
	case "context":
		return fhirVal.Context, true
	case "note":
		return fhirVal.Note, true
	case "medicationreference":
		return fhirVal.MedicationReference, true
	case "category":
		return fhirVal.Category, true
	case "dayssupply":
		return fhirVal.DaysSupply, true
	case "substitution":
		return fhirVal.Substitution, true
	case "notdone":
		return fhirVal.NotDone, true
	case "eventhistory":
		return fhirVal.EventHistory, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "medicationcodeableconcept":
		return fhirVal.MedicationCodeableConcept, true
	case "subject":
		return fhirVal.Subject, true
	case "supportinginformation":
		return fhirVal.SupportingInformation, true
	case "authorizingprescription":
		return fhirVal.AuthorizingPrescription, true
	case "type":
		return fhirVal.Type, true
	case "destination":
		return fhirVal.Destination, true
	case "partof":
		return fhirVal.PartOf, true
	case "notdonereasonreference":
		return fhirVal.NotDoneReasonReference, true
	case "detectedissue":
		return fhirVal.DetectedIssue, true

	default:
		return nil, false
	}
}

func (fhirVal *MedicationDispense) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"whenprepared": &FieldTypeSupport{"FHIRDateTime", false, true},
		"whenhandedover": &FieldTypeSupport{"FHIRDateTime", false, true},
		"receiver": &FieldTypeSupport{"Reference", true, false},
		"dosageinstruction": &FieldTypeSupport{"Dosage", true, false},
		"notdonereasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"performer": &FieldTypeSupport{"MedicationDispensePerformerComponent", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"dayssupply": &FieldTypeSupport{"Quantity", false, true},
		"substitution": &FieldTypeSupport{"MedicationDispenseSubstitutionComponent", false, true},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"eventhistory": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"authorizingprescription": &FieldTypeSupport{"Reference", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"destination": &FieldTypeSupport{"Reference", false, true},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"notdonereasonreference": &FieldTypeSupport{"Reference", false, true},
		"detectedissue": &FieldTypeSupport{"Reference", true, false},

	}
}
