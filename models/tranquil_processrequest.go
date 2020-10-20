
package models


func (fhirVal *ProcessRequest) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ProcessRequest) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "response":
		return fhirVal.Response, true
	case "nullify":
		return fhirVal.Nullify, true
	case "status":
		return fhirVal.Status, true
	case "request":
		return fhirVal.Request, true
	case "item":
		return fhirVal.Item, true
	case "include":
		return fhirVal.Include, true
	case "exclude":
		return fhirVal.Exclude, true
	case "period":
		return fhirVal.Period, true
	case "identifier":
		return fhirVal.Identifier, true
	case "action":
		return fhirVal.Action, true
	case "target":
		return fhirVal.Target, true
	case "organization":
		return fhirVal.Organization, true
	case "reference":
		return fhirVal.Reference, true
	case "created":
		return fhirVal.Created, true
	case "provider":
		return fhirVal.Provider, true

	default:
		return nil, false
	}
}

func (fhirVal *ProcessRequest) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"response": &FieldTypeSupport{"Reference", false, true},
		"nullify": &FieldTypeSupport{"bool", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"request": &FieldTypeSupport{"Reference", false, true},
		"item": &FieldTypeSupport{"ProcessRequestItemsComponent", true, false},
		"include": &FieldTypeSupport{"string", true, false},
		"exclude": &FieldTypeSupport{"string", true, false},
		"period": &FieldTypeSupport{"Period", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"action": &FieldTypeSupport{"string", false, false},
		"target": &FieldTypeSupport{"Reference", false, true},
		"organization": &FieldTypeSupport{"Reference", false, true},
		"reference": &FieldTypeSupport{"string", false, false},
		"created": &FieldTypeSupport{"FHIRDateTime", false, true},
		"provider": &FieldTypeSupport{"Reference", false, true},

	}
}
