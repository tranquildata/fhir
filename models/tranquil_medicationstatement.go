
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
	case "medicationcodeableconcept":
		return fhirVal.MedicationCodeableConcept, true
	case "medicationreference":
		return fhirVal.MedicationReference, true
	case "dateasserted":
		return fhirVal.DateAsserted, true
	case "reasonnottaken":
		return fhirVal.ReasonNotTaken, true
	case "subject":
		return fhirVal.Subject, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "status":
		return fhirVal.Status, true
	case "category":
		return fhirVal.Category, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true
	case "dosage":
		return fhirVal.Dosage, true
	case "identifier":
		return fhirVal.Identifier, true
	case "partof":
		return fhirVal.PartOf, true
	case "taken":
		return fhirVal.Taken, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "context":
		return fhirVal.Context, true
	case "effectivedatetime":
		return fhirVal.EffectiveDateTime, true
	case "informationsource":
		return fhirVal.InformationSource, true
	case "derivedfrom":
		return fhirVal.DerivedFrom, true

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
		"medicationcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"medicationreference": &FieldTypeSupport{"Reference", false, true},
		"dateasserted": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reasonnottaken": &FieldTypeSupport{"CodeableConcept", true, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"dosage": &FieldTypeSupport{"Dosage", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"partof": &FieldTypeSupport{"Reference", true, false},
		"taken": &FieldTypeSupport{"string", false, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"effectivedatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"informationsource": &FieldTypeSupport{"Reference", false, true},
		"derivedfrom": &FieldTypeSupport{"Reference", true, false},

	}
}
