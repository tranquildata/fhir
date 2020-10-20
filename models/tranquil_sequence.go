
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
	case "performer":
		return fhirVal.Performer, true
	case "quality":
		return fhirVal.Quality, true
	case "readcoverage":
		return fhirVal.ReadCoverage, true
	case "device":
		return fhirVal.Device, true
	case "quantity":
		return fhirVal.Quantity, true
	case "referenceseq":
		return fhirVal.ReferenceSeq, true
	case "variant":
		return fhirVal.Variant, true
	case "identifier":
		return fhirVal.Identifier, true
	case "patient":
		return fhirVal.Patient, true
	case "specimen":
		return fhirVal.Specimen, true
	case "repository":
		return fhirVal.Repository, true
	case "type":
		return fhirVal.Type, true
	case "coordinatesystem":
		return fhirVal.CoordinateSystem, true
	case "observedseq":
		return fhirVal.ObservedSeq, true
	case "pointer":
		return fhirVal.Pointer, true

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
		"performer": &FieldTypeSupport{"Reference", false, true},
		"quality": &FieldTypeSupport{"SequenceQualityComponent", true, false},
		"readcoverage": &FieldTypeSupport{"int32", false, true},
		"device": &FieldTypeSupport{"Reference", false, true},
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"referenceseq": &FieldTypeSupport{"SequenceReferenceSeqComponent", false, true},
		"variant": &FieldTypeSupport{"SequenceVariantComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"specimen": &FieldTypeSupport{"Reference", false, true},
		"repository": &FieldTypeSupport{"SequenceRepositoryComponent", true, false},
		"type": &FieldTypeSupport{"string", false, false},
		"coordinatesystem": &FieldTypeSupport{"int32", false, true},
		"observedseq": &FieldTypeSupport{"string", false, false},
		"pointer": &FieldTypeSupport{"Reference", true, false},

	}
}
