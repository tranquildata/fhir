
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
	case "basedon":
		return fhirVal.BasedOn, true
	case "referrer":
		return fhirVal.Referrer, true
	case "endpoint":
		return fhirVal.Endpoint, true
	case "uid":
		return fhirVal.Uid, true
	case "availability":
		return fhirVal.Availability, true
	case "context":
		return fhirVal.Context, true
	case "description":
		return fhirVal.Description, true
	case "series":
		return fhirVal.Series, true
	case "procedurecode":
		return fhirVal.ProcedureCode, true
	case "reason":
		return fhirVal.Reason, true
	case "started":
		return fhirVal.Started, true
	case "interpreter":
		return fhirVal.Interpreter, true
	case "numberofinstances":
		return fhirVal.NumberOfInstances, true
	case "patient":
		return fhirVal.Patient, true
	case "numberofseries":
		return fhirVal.NumberOfSeries, true
	case "procedurereference":
		return fhirVal.ProcedureReference, true
	case "accession":
		return fhirVal.Accession, true
	case "identifier":
		return fhirVal.Identifier, true
	case "modalitylist":
		return fhirVal.ModalityList, true

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
		"basedon": &FieldTypeSupport{"Reference", true, false},
		"referrer": &FieldTypeSupport{"Reference", false, true},
		"endpoint": &FieldTypeSupport{"Reference", true, false},
		"uid": &FieldTypeSupport{"string", false, false},
		"availability": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"Reference", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"series": &FieldTypeSupport{"ImagingStudySeriesComponent", true, false},
		"procedurecode": &FieldTypeSupport{"CodeableConcept", true, false},
		"reason": &FieldTypeSupport{"CodeableConcept", false, true},
		"started": &FieldTypeSupport{"FHIRDateTime", false, true},
		"interpreter": &FieldTypeSupport{"Reference", true, false},
		"numberofinstances": &FieldTypeSupport{"uint32", false, true},
		"patient": &FieldTypeSupport{"Reference", false, true},
		"numberofseries": &FieldTypeSupport{"uint32", false, true},
		"procedurereference": &FieldTypeSupport{"Reference", true, false},
		"accession": &FieldTypeSupport{"Identifier", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"modalitylist": &FieldTypeSupport{"Coding", true, false},

	}
}
