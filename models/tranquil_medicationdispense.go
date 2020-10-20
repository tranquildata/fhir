
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
		"subject": &FieldTypeSupport{"Reference", false, true},
		"supportinginformation": &FieldTypeSupport{"Reference", true, false},
		"authorizingprescription": &FieldTypeSupport{"Reference", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"detectedissue": &FieldTypeSupport{"Reference", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"notdonereasonreference": &FieldTypeSupport{"Reference", false, true},
		"performer": &FieldTypeSupport{"MedicationDispensePerformerComponent", true, false},
		"whenhandedover": &FieldTypeSupport{"FHIRDateTime", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"dosageinstruction": &FieldTypeSupport{"Dosage", true, false},
		"eventhistory": &FieldTypeSupport{"Reference", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"context": &FieldTypeSupport{"Reference", false, true},
		"dayssupply": &FieldTypeSupport{"Quantity", false, true},
		"destination": &FieldTypeSupport{"Reference", false, true},
		"receiver": &FieldTypeSupport{"Reference", true, false},
		"substitution": &FieldTypeSupport{"MedicationDispenseSubstitutionComponent", false, true},
		"notdonereasoncodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"whenprepared": &FieldTypeSupport{"FHIRDateTime", false, true},
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"quantity": &FieldTypeSupport{"Quantity", false, true},

	}
}
