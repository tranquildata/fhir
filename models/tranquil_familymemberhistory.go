
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
		"gender": &FieldTypeSupport{"string", false, false},
		"bornperiod": &FieldTypeSupport{"Period", false, true},
		"ageage": &FieldTypeSupport{"Quantity", false, true},
		"agerange": &FieldTypeSupport{"Range", false, true},
		"estimatedage": &FieldTypeSupport{"bool", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"condition": &FieldTypeSupport{"FamilyMemberHistoryConditionComponent", true, false},
		"notdonereason": &FieldTypeSupport{"CodeableConcept", false, true},
		"deceaseddate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"note": &FieldTypeSupport{"Annotation", true, false},
		"bornstring": &FieldTypeSupport{"string", false, false},
		"deceasedboolean": &FieldTypeSupport{"bool", false, true},
		"definition": &FieldTypeSupport{"Reference", true, false},
		"notdone": &FieldTypeSupport{"bool", false, true},
		"borndate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"agestring": &FieldTypeSupport{"string", false, false},
		"deceasedage": &FieldTypeSupport{"Quantity", false, true},
		"deceasedrange": &FieldTypeSupport{"Range", false, true},
		"deceasedstring": &FieldTypeSupport{"string", false, false},
		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"relationship": &FieldTypeSupport{"CodeableConcept", false, true},
		"reasonreference": &FieldTypeSupport{"Reference", true, false},

	}
}
