
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
	case "indexed":
		return fhirVal.Indexed, true
	case "author":
		return fhirVal.Author, true
	case "context":
		return fhirVal.Context, true
	case "securitylabel":
		return fhirVal.SecurityLabel, true
	case "content":
		return fhirVal.Content, true
	case "masteridentifier":
		return fhirVal.MasterIdentifier, true
	case "type":
		return fhirVal.Type, true
	case "subject":
		return fhirVal.Subject, true
	case "custodian":
		return fhirVal.Custodian, true
	case "class":
		return fhirVal.Class, true
	case "created":
		return fhirVal.Created, true
	case "relatesto":
		return fhirVal.RelatesTo, true
	case "description":
		return fhirVal.Description, true
	case "identifier":
		return fhirVal.Identifier, true
	case "status":
		return fhirVal.Status, true
	case "docstatus":
		return fhirVal.DocStatus, true
	case "authenticator":
		return fhirVal.Authenticator, true

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
		"indexed": &FieldTypeSupport{"FHIRDateTime", false, true},
		"author": &FieldTypeSupport{"Reference", true, false},
		"context": &FieldTypeSupport{"DocumentReferenceContextComponent", false, true},
		"securitylabel": &FieldTypeSupport{"CodeableConcept", true, false},
		"content": &FieldTypeSupport{"DocumentReferenceContentComponent", true, false},
		"masteridentifier": &FieldTypeSupport{"Identifier", false, true},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"custodian": &FieldTypeSupport{"Reference", false, true},
		"class": &FieldTypeSupport{"CodeableConcept", false, true},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"relatesto": &FieldTypeSupport{"DocumentReferenceRelatesToComponent", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"status": &FieldTypeSupport{"string", false, false},
		"docstatus": &FieldTypeSupport{"string", false, false},
		"authenticator": &FieldTypeSupport{"Reference", false, true},

	}
}
