
package models


func (fhirVal *DocumentReference) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *DocumentReference) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "subject":
		return fhirVal.Subject, true
	case "created":
		return fhirVal.Created, true
	case "indexed":
		return fhirVal.Indexed, true
	case "authenticator":
		return fhirVal.Authenticator, true
	case "custodian":
		return fhirVal.Custodian, true
	case "securitylabel":
		return fhirVal.SecurityLabel, true
	case "masteridentifier":
		return fhirVal.MasterIdentifier, true
	case "relatesto":
		return fhirVal.RelatesTo, true
	case "identifier":
		return fhirVal.Identifier, true
	case "type":
		return fhirVal.Type, true
	case "class":
		return fhirVal.Class, true
	case "author":
		return fhirVal.Author, true
	case "description":
		return fhirVal.Description, true
	case "context":
		return fhirVal.Context, true
	case "docstatus":
		return fhirVal.DocStatus, true
	case "content":
		return fhirVal.Content, true

	default:
		return nil, false
	}
}

func (fhirVal *DocumentReference) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"subject": &FieldTypeSupport{"Reference", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"indexed": &FieldTypeSupport{"FHIRDateTime", false, true},
		"authenticator": &FieldTypeSupport{"Reference", false, true},
		"custodian": &FieldTypeSupport{"Reference", false, true},
		"securitylabel": &FieldTypeSupport{"CodeableConcept", true, false},
		"masteridentifier": &FieldTypeSupport{"Identifier", false, true},
		"relatesto": &FieldTypeSupport{"DocumentReferenceRelatesToComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"class": &FieldTypeSupport{"CodeableConcept", false, true},
		"author": &FieldTypeSupport{"Reference", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"DocumentReferenceContextComponent", false, true},
		"docstatus": &FieldTypeSupport{"string", false, false},
		"content": &FieldTypeSupport{"DocumentReferenceContentComponent", true, false},

	}
}
