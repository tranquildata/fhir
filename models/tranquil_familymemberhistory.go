
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
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "date":
		return fhirVal.Date, true
	case "bornperiod":
		return fhirVal.BornPeriod, true
	case "bornstring":
		return fhirVal.BornString, true
	case "deceasedage":
		return fhirVal.DeceasedAge, true
	case "condition":
		return fhirVal.Condition, true
	case "agerange":
		return fhirVal.AgeRange, true
	case "deceasedboolean":
		return fhirVal.DeceasedBoolean, true
	case "reasonreference":
		return fhirVal.ReasonReference, true
	case "note":
		return fhirVal.Note, true
	case "patient":
		return fhirVal.Patient, true
	case "name":
		return fhirVal.Name, true
	case "relationship":
		return fhirVal.Relationship, true
	case "ageage":
		return fhirVal.AgeAge, true
	case "estimatedage":
		return fhirVal.EstimatedAge, true
	case "deceasedrange":
		return fhirVal.DeceasedRange, true
	case "deceaseddate":
		return fhirVal.DeceasedDate, true
	case "reasoncode":
		return fhirVal.ReasonCode, true
	case "definition":
		return fhirVal.Definition, true
	case "notdone":
		return fhirVal.NotDone, true
	case "notdonereason":
		return fhirVal.NotDoneReason, true
	case "gender":
		return fhirVal.Gender, true
	case "borndate":
		return fhirVal.BornDate, true
	case "agestring":
		return fhirVal.AgeString, true
	case "deceasedstring":
		return fhirVal.DeceasedString, true

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
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"bornperiod": &FieldTypeSupport{"Period", false, true},
		"bornstring": &FieldTypeSupport{"string", false, false},
		"deceasedage": &FieldTypeSupport{"Quantity", false, true},
		"condition": &FieldTypeSupport{"FamilyMemberHistoryConditionComponent", true, false},
		"agerange": &FieldTypeSupport{"Range", false, true},
		"deceasedboolean": &FieldTypeSupport{"bool", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"name": &FieldTypeSupport{"string", false, false},
		"relationship": &FieldTypeSupport{"CodeableConcept", false, true},
		"ageage": &FieldTypeSupport{"Quantity", false, true},
		"estimatedage": &FieldTypeSupport{"bool", false, true},
		"deceasedrange": &FieldTypeSupport{"Range", false, true},
		"deceaseddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"gender": &FieldTypeSupport{"string", false, false},
		"borndate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"agestring": &FieldTypeSupport{"string", false, false},
		"deceasedstring": &FieldTypeSupport{"string", false, false},

	}
}
