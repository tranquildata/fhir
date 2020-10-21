
package models


func (fhirVal *Medication) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Medication) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "code":
		return fhirVal.Code, true
	case "isbrand":
		return fhirVal.IsBrand, true
	case "isoverthecounter":
		return fhirVal.IsOverTheCounter, true
	case "image":
		return fhirVal.Image, true
	case "package":
		return fhirVal.Package, true
	case "status":
		return fhirVal.Status, true
	case "manufacturer":
		return fhirVal.Manufacturer, true
	case "form":
		return fhirVal.Form, true
	case "ingredient":
		return fhirVal.Ingredient, true

	default:
		return nil, false
	}
}

func (fhirVal *Medication) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"isbrand": &FieldTypeSupport{"bool", false, true},
		"isoverthecounter": &FieldTypeSupport{"bool", false, true},
		"image": &FieldTypeSupport{"Attachment", true, false},
		"package": &FieldTypeSupport{"MedicationPackageComponent", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"manufacturer": &FieldTypeSupport{"Reference", false, true},
		"form": &FieldTypeSupport{"CodeableConcept", false, true},
		"ingredient": &FieldTypeSupport{"MedicationIngredientComponent", true, false},

	}
}
