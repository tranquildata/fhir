
package models


func (fhirVal *DetectedIssue) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *DetectedIssue) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "date":
		return fhirVal.Date, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "category":
		return fhirVal.Category, true
	case "severity":
		return fhirVal.Severity, true
	case "reference":
		return fhirVal.Reference, true
	case "mitigation":
		return fhirVal.Mitigation, true
	case "patient":
		return fhirVal.Patient, true
	case "author":
		return fhirVal.Author, true
	case "implicated":
		return fhirVal.Implicated, true
	case "detail":
		return fhirVal.Detail, true

	default:
		return nil, false
	}
}

func (fhirVal *DetectedIssue) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"category": &FieldTypeSupport{"CodeableConcept", false, true},
		"severity": &FieldTypeSupport{"string", false, false},
		"reference": &FieldTypeSupport{"string", false, false},
		"mitigation": &FieldTypeSupport{"DetectedIssueMitigationComponent", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"author": &FieldTypeSupport{"Reference", false, true},
		"implicated": &FieldTypeSupport{"Reference", true, false},
		"detail": &FieldTypeSupport{"string", false, false},

	}
}
