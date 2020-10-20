
package models


func (fhirVal *ProcessResponse) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ProcessResponse) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "outcome":
		return fhirVal.Outcome, true
	case "disposition":
		return fhirVal.Disposition, true
	case "processnote":
		return fhirVal.ProcessNote, true
	case "error":
		return fhirVal.Error, true
	case "created":
		return fhirVal.Created, true
	case "status":
		return fhirVal.Status, true
	case "organization":
		return fhirVal.Organization, true
	case "request":
		return fhirVal.Request, true
	case "requestprovider":
		return fhirVal.RequestProvider, true
	case "requestorganization":
		return fhirVal.RequestOrganization, true
	case "form":
		return fhirVal.Form, true
	case "communicationrequest":
		return fhirVal.CommunicationRequest, true
	case "identifier":
		return fhirVal.Identifier, true

	default:
		return nil, false
	}
}

func (fhirVal *ProcessResponse) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"outcome": &FieldTypeSupport{"CodeableConcept", false, true},
		"disposition": &FieldTypeSupport{"string", false, false},
		"processnote": &FieldTypeSupport{"ProcessResponseProcessNoteComponent", true, false},
		"error": &FieldTypeSupport{"CodeableConcept", true, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"request": &FieldTypeSupport{"Reference", false, true},
		"requestprovider": &FieldTypeSupport{"Reference", false, true},
		"requestorganization": &FieldTypeSupport{"Reference", false, true},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},
		"communicationrequest": &FieldTypeSupport{"Reference", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},

	}
}
