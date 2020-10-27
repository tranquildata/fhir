
 package models
 
 import (
	 "fmt"
 )
 
 
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
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "referenceseq":
		 return fhirVal.ReferenceSeq, true
 	case "repository":
		 return fhirVal.Repository, true
 	case "quality":
		 return fhirVal.Quality, true
 	case "type":
		 return fhirVal.Type, true
 	case "patient":
		 return fhirVal.Patient, true
 	case "performer":
		 return fhirVal.Performer, true
 	case "quantity":
		 return fhirVal.Quantity, true
 	case "variant":
		 return fhirVal.Variant, true
 	case "observedseq":
		 return fhirVal.ObservedSeq, true
 	case "coordinatesystem":
		 return fhirVal.CoordinateSystem, true
 	case "device":
		 return fhirVal.Device, true
 	case "specimen":
		 return fhirVal.Specimen, true
 	case "readcoverage":
		 return fhirVal.ReadCoverage, true
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
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"referenceseq": &FieldTypeSupport{"SequenceReferenceSeqComponent", false, true},
 		"repository": &FieldTypeSupport{"SequenceRepositoryComponent", true, false},
 		"quality": &FieldTypeSupport{"SequenceQualityComponent", true, false},
 		"type": &FieldTypeSupport{"string", false, false},
 		"patient": &FieldTypeSupport{"Reference", false, true},
 		"performer": &FieldTypeSupport{"Reference", false, true},
 		"quantity": &FieldTypeSupport{"Quantity", false, true},
 		"variant": &FieldTypeSupport{"SequenceVariantComponent", true, false},
 		"observedseq": &FieldTypeSupport{"string", false, false},
 		"coordinatesystem": &FieldTypeSupport{"int32", false, true},
 		"device": &FieldTypeSupport{"Reference", false, true},
 		"specimen": &FieldTypeSupport{"Reference", false, true},
 		"readcoverage": &FieldTypeSupport{"int32", false, true},
 		"pointer": &FieldTypeSupport{"Reference", true, false},
 
	 }
 }
 
 func (fhirVal *Sequence) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 