
package models


func (fhirVal *Device) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Device) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "lotnumber":
		return fhirVal.LotNumber, true
	case "model":
		return fhirVal.Model, true
	case "contact":
		return fhirVal.Contact, true
	case "safety":
		return fhirVal.Safety, true
	case "manufacturer":
		return fhirVal.Manufacturer, true
	case "expirationdate":
		return fhirVal.ExpirationDate, true
	case "location":
		return fhirVal.Location, true
	case "udi":
		return fhirVal.Udi, true
	case "status":
		return fhirVal.Status, true
	case "manufacturedate":
		return fhirVal.ManufactureDate, true
	case "version":
		return fhirVal.Version, true
	case "owner":
		return fhirVal.Owner, true
	case "identifier":
		return fhirVal.Identifier, true
	case "type":
		return fhirVal.Type, true
	case "patient":
		return fhirVal.Patient, true
	case "url":
		return fhirVal.Url, true
	case "note":
		return fhirVal.Note, true

	default:
		return nil, false
	}
}

func (fhirVal *Device) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"lotnumber": &FieldTypeSupport{"string", false, false},
		"model": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactPoint", true, false},
		"safety": &FieldTypeSupport{"CodeableConcept", true, false},
		"manufacturer": &FieldTypeSupport{"string", false, false},
		"expirationdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"location": &FieldTypeSupport{"Reference", false, true},
		"udi": &FieldTypeSupport{"DeviceUdiComponent", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"manufacturedate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"version": &FieldTypeSupport{"string", false, false},
		"owner": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}
