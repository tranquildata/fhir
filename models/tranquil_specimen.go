
package models


func (fhirVal *Specimen) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *Specimen) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "collection":
		return fhirVal.Collection, true
	case "identifier":
		return fhirVal.Identifier, true
	case "type":
		return fhirVal.Type, true
	case "subject":
		return fhirVal.Subject, true
	case "receivedtime":
		return fhirVal.ReceivedTime, true
	case "parent":
		return fhirVal.Parent, true
	case "request":
		return fhirVal.Request, true
	case "accessionidentifier":
		return fhirVal.AccessionIdentifier, true
	case "status":
		return fhirVal.Status, true
	case "processing":
		return fhirVal.Processing, true
	case "container":
		return fhirVal.Container, true
	case "note":
		return fhirVal.Note, true

	default:
		return nil, false
	}
}

func (fhirVal *Specimen) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"collection": &FieldTypeSupport{"SpecimenCollectionComponent", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"type": &FieldTypeSupport{"CodeableConcept", false, true},
		"subject": &FieldTypeSupport{"Reference", false, true},
		"receivedtime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"parent": &FieldTypeSupport{"Reference", true, false},
		"request": &FieldTypeSupport{"Reference", true, false},
		"accessionidentifier": &FieldTypeSupport{"Identifier", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"processing": &FieldTypeSupport{"SpecimenProcessingComponent", true, false},
		"container": &FieldTypeSupport{"SpecimenContainerComponent", true, false},
		"note": &FieldTypeSupport{"Annotation", true, false},

	}
}
