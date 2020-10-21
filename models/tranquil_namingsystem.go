
package models


func (fhirVal *NamingSystem) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *NamingSystem) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "date":
		return fhirVal.Date, true
	case "responsible":
		return fhirVal.Responsible, true
	case "type":
		return fhirVal.Type, true
	case "uniqueid":
		return fhirVal.UniqueId, true
	case "kind":
		return fhirVal.Kind, true
	case "name":
		return fhirVal.Name, true
	case "publisher":
		return fhirVal.Publisher, true
	case "description":
		return fhirVal.Description, true
	case "contact":
		return fhirVal.Contact, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "usage":
		return fhirVal.Usage, true
	case "replacedby":
		return fhirVal.ReplacedBy, true

	default:
		return nil, false
	}
}

func (fhirVal *NamingSystem) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"responsible": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"uniqueid": &FieldTypeSupport{"NamingSystemUniqueIdComponent", true, false},
		"kind": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"replacedby": &FieldTypeSupport{"Reference", false, true},

	}
}
