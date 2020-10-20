
package models


func (fhirVal *ImagingStudy) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ImagingStudy) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "identifier":
		return fhirVal.Identifier, true
	case "context":
		return fhirVal.Context, true
	case "basedon":
		return fhirVal.BasedOn, true
	case "endpoint":
		return fhirVal.Endpoint, true
	case "accession":
		return fhirVal.Accession, true
	case "started":
		return fhirVal.Started, true
	case "numberofseries":
		return fhirVal.NumberOfSeries, true
	case "procedurereference":
		return fhirVal.ProcedureReference, true
	case "procedurecode":
		return fhirVal.ProcedureCode, true
	case "series":
		return fhirVal.Series, true
	case "patient":
		return fhirVal.Patient, true
	case "referrer":
		return fhirVal.Referrer, true
	case "interpreter":
		return fhirVal.Interpreter, true
	case "uid":
		return fhirVal.Uid, true
	case "availability":
		return fhirVal.Availability, true
	case "modalitylist":
		return fhirVal.ModalityList, true
	case "numberofinstances":
		return fhirVal.NumberOfInstances, true
	case "reason":
		return fhirVal.Reason, true
	case "description":
		return fhirVal.Description, true

	default:
		return nil, false
	}
}

func (fhirVal *ImagingStudy) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"endpoint": &FieldTypeSupport{"Reference", true, false},
		"accession": &FieldTypeSupport{"Identifier", false, true},
		"started": &FieldTypeSupport{"FHIRDateTime", false, true},
		"numberofseries": &FieldTypeSupport{"uint32", false, true},
		"procedurereference": &FieldTypeSupport{"Reference", true, false},
		"procedurecode": &FieldTypeSupport{"CodeableConcept", true, false},
		"series": &FieldTypeSupport{"ImagingStudySeriesComponent", true, false},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"referrer": &FieldTypeSupport{"Reference", false, true},
		"interpreter": &FieldTypeSupport{"Reference", true, false},
		"uid": &FieldTypeSupport{"string", false, false},
		"availability": &FieldTypeSupport{"string", false, false},
		"modalitylist": &FieldTypeSupport{"Coding", true, false},
		"numberofinstances": &FieldTypeSupport{"uint32", false, true},
		"reason": &FieldTypeSupport{"CodeableConcept", false, true},
		"description": &FieldTypeSupport{"string", false, false},

	}
}
