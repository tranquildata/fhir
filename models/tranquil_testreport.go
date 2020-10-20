
package models


func (fhirVal *TestReport) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *TestReport) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "testscript":
		return fhirVal.TestScript, true
	case "tester":
		return fhirVal.Tester, true
	case "participant":
		return fhirVal.Participant, true
	case "test":
		return fhirVal.Test, true
	case "teardown":
		return fhirVal.Teardown, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "result":
		return fhirVal.Result, true
	case "score":
		return fhirVal.Score, true
	case "issued":
		return fhirVal.Issued, true
	case "setup":
		return fhirVal.Setup, true
	case "name":
		return fhirVal.Name, true

	default:
		return nil, false
	}
}

func (fhirVal *TestReport) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"testscript": &FieldTypeSupport{"Reference", false, true},
		"tester": &FieldTypeSupport{"string", false, false},
		"participant": &FieldTypeSupport{"TestReportParticipantComponent", true, false},
		"test": &FieldTypeSupport{"TestReportTestComponent", true, false},
		"teardown": &FieldTypeSupport{"TestReportTeardownComponent", false, true},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"result": &FieldTypeSupport{"string", false, false},
		"score": &FieldTypeSupport{"float64", false, true},
		"issued": &FieldTypeSupport{"FHIRDateTime", false, true},
		"setup": &FieldTypeSupport{"TestReportSetupComponent", false, true},
		"name": &FieldTypeSupport{"string", false, false},

	}
}
