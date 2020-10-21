
package models


func (fhirVal *MedicationStatement) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *MedicationStatement) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "note":
		return fhirVal.Note, true
	case "category":
		return fhirVal.Category, true
	case "medicationreference":
		return fhirVal.MedicationReference, true
	case "dosage":
		return fhirVal.Dosage, true
	case "reasonnottaken":
		return fhirVal.ReasonNotTaken, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true
	case "dateasserted":
		return fhirVal.DateAsserted, true
	case "subject":
		return fhirVal.Subject, true
	case "taken":
		return fhirVal.Taken, true
	case "derivedfrom":
		return fhirVal.DerivedFrom, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "identifier":
		return fhirVal.Identifier, true
	case "partof":
		return fhirVal.PartOf, true
	case "context":
		return fhirVal.Context, true
	case "medicationcodeableconcept":
		return fhirVal.MedicationCodeableConcept, true
	case "informationsource":
		return fhirVal.InformationSource, true

	default:
		return nil, false
	}
}

func (fhirVal *MedicationStatement) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"dosage": &FieldTypeSupport{"Dosage", true, false},
		"reasonnottaken": &FieldTypeSupport{"CodeableConcept", true, false},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"dateasserted": &FieldTypeSupport{"FHIRDateTime", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"taken": &FieldTypeSupport{"string", false, false},
		"derivedfrom": &FieldTypeSupport{"Reference", true, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"informationsource": &FieldTypeSupport{"Reference", false, true},

	}
}
