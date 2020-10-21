
package models


func (fhirVal *DocumentManifest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *DocumentManifest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "type":
		return fhirVal.Type, true
	case "created":
		return fhirVal.Created, true
	case "author":
		return fhirVal.Author, true
	case "source":
		return fhirVal.Source, true
	case "description":
		return fhirVal.Description, true
	case "masteridentifier":
		return fhirVal.MasterIdentifier, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "subject":
		return fhirVal.Subject, true
	case "recipient":
		return fhirVal.Recipient, true
	case "content":
		return fhirVal.Content, true
	case "related":
		return fhirVal.Related, true

	default:
		return nil, false
	}
}

func (fhirVal *DocumentManifest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"author": &FieldTypeSupport{"Reference", true, false},
		"source": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"masteridentifier": &FieldTypeSupport{"Identifier", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"recipient": &FieldTypeSupport{"Reference", true, false},
		"content": &FieldTypeSupport{"DocumentManifestContentComponent", true, false},
		"related": &FieldTypeSupport{"DocumentManifestRelatedComponent", true, false},

	}
}
