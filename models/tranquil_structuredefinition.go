
package models


func (fhirVal *StructureDefinition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *StructureDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "copyright":
		return fhirVal.Copyright, true
	case "fhirversion":
		return fhirVal.FhirVersion, true
	case "identifier":
		return fhirVal.Identifier, true
	case "title":
		return fhirVal.Title, true
	case "type":
		return fhirVal.Type, true
	case "experimental":
		return fhirVal.Experimental, true
	case "publisher":
		return fhirVal.Publisher, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "derivation":
		return fhirVal.Derivation, true
	case "snapshot":
		return fhirVal.Snapshot, true
	case "date":
		return fhirVal.Date, true
	case "contact":
		return fhirVal.Contact, true
	case "description":
		return fhirVal.Description, true
	case "kind":
		return fhirVal.Kind, true
	case "version":
		return fhirVal.Version, true
	case "status":
		return fhirVal.Status, true
	case "abstract":
		return fhirVal.Abstract, true
	case "basedefinition":
		return fhirVal.BaseDefinition, true
	case "contextinvariant":
		return fhirVal.ContextInvariant, true
	case "differential":
		return fhirVal.Differential, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "keyword":
		return fhirVal.Keyword, true
	case "mapping":
		return fhirVal.Mapping, true
	case "context":
		return fhirVal.Context, true
	case "url":
		return fhirVal.Url, true
	case "name":
		return fhirVal.Name, true
	case "purpose":
		return fhirVal.Purpose, true
	case "contexttype":
		return fhirVal.ContextType, true

	default:
		return nil, false
	}
}

func (fhirVal *StructureDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"copyright": &FieldTypeSupport{"string", false, false},
		"fhirversion": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"derivation": &FieldTypeSupport{"string", false, false},
		"snapshot": &FieldTypeSupport{"StructureDefinitionSnapshotComponent", false, true},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"kind": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"abstract": &FieldTypeSupport{"bool", false, true},
		"basedefinition": &FieldTypeSupport{"string", false, false},
		"contextinvariant": &FieldTypeSupport{"string", true, false},
		"differential": &FieldTypeSupport{"StructureDefinitionDifferentialComponent", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"keyword": &FieldTypeSupport{"Coding", true, false},
		"mapping": &FieldTypeSupport{"StructureDefinitionMappingComponent", true, false},
		"context": &FieldTypeSupport{"string", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"contexttype": &FieldTypeSupport{"string", false, false},

	}
}
