
package models


func (fhirVal *FamilyMemberHistory) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *FamilyMemberHistory) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "bornstring":
		return fhirVal.BornString, true
	case "agestring":
		return fhirVal.AgeString, true
	case "deceaseddate":
		return fhirVal.DeceasedDate, true
	case "condition":
		return fhirVal.Condition, true
	case "relationship":
		return fhirVal.Relationship, true
	case "gender":
		return fhirVal.Gender, true
	case "bornperiod":
		return fhirVal.BornPeriod, true
	case "deceasedboolean":
		return fhirVal.DeceasedBoolean, true
	case "status":
		return fhirVal.Status, true
	case "name":
		return fhirVal.Name, true
	case "patient":
		return fhirVal.Patient, true
	case "borndate":
		return fhirVal.BornDate, true
	case "deceasedstring":
		return fhirVal.DeceasedString, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "identifier":
		return fhirVal.Identifier, true
	case "notdone":
		return fhirVal.NotDone, true
	case "date":
		return fhirVal.Date, true
	case "ageage":
		return fhirVal.AgeAge, true
	case "agerange":
		return fhirVal.AgeRange, true
	case "estimatedage":
		return fhirVal.EstimatedAge, true
	case "deceasedage":
		return fhirVal.DeceasedAge, true
	case "deceasedrange":
		return fhirVal.DeceasedRange, true
	case "definition":
		return fhirVal.Definition, true
	case "notdonereason":
		return fhirVal.NotDoneReason, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true

	default:
		return nil, false
	}
}

func (fhirVal *FamilyMemberHistory) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"bornstring": &FieldTypeSupport{"string", false, false},
		"agestring": &FieldTypeSupport{"string", false, false},
		"deceaseddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"condition": &FieldTypeSupport{"FamilyMemberHistoryConditionComponent", true, false},
		"relationship": &FieldTypeSupport{"CodeableConcept", false, true},
		"gender": &FieldTypeSupport{"string", false, false},
		"bornperiod": &FieldTypeSupport{"Period", false, true},
		"deceasedboolean": &FieldTypeSupport{"bool", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"borndate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"deceasedstring": &FieldTypeSupport{"string", false, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"ageage": &FieldTypeSupport{"Quantity", false, true},
		"agerange": &FieldTypeSupport{"Range", false, true},
		"estimatedage": &FieldTypeSupport{"bool", false, true},
		"deceasedage": &FieldTypeSupport{"Quantity", false, true},
		"deceasedrange": &FieldTypeSupport{"Range", false, true},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}