
package models


func (fhirVal *Sequence) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Sequence) FieldByLowerName(nameLower string) (interface{}, bool) {
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

	default:
		return nil, false
	}
}

func (fhirVal *Sequence) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"readcoverage": &FieldTypeSupport{"int32", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"device": &FieldTypeSupport{"Reference", false, true},
		"referenceseq": &FieldTypeSupport{"SequenceReferenceSeqComponent", false, true},
		"pointer": &FieldTypeSupport{"Reference", true, false},
		"coordinatesystem": &FieldTypeSupport{"int32", false, true},
		"variant": &FieldTypeSupport{"SequenceVariantComponent", true, false},
		"repository": &FieldTypeSupport{"SequenceRepositoryComponent", true, false},
		"quality": &FieldTypeSupport{"SequenceQualityComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"type": &FieldTypeSupport{"string", false, false},
		"performer": &FieldTypeSupport{"Reference", false, true},
		"specimen": &FieldTypeSupport{"Reference", false, true},
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"observedseq": &FieldTypeSupport{"string", false, false},

	}
}
