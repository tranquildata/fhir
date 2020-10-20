
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
	case "usecontext":
		return fhirVal.UseContext, true
	case "experimental":
		return fhirVal.Experimental, true
	case "description":
		return fhirVal.Description, true
	case "type":
		return fhirVal.Type, true
	case "context":
		return fhirVal.Context, true
	case "name":
		return fhirVal.Name, true
	case "publisher":
		return fhirVal.Publisher, true
	case "copyright":
		return fhirVal.Copyright, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "contextinvariant":
		return fhirVal.ContextInvariant, true
	case "derivation":
		return fhirVal.Derivation, true
	case "snapshot":
		return fhirVal.Snapshot, true
	case "identifier":
		return fhirVal.Identifier, true
	case "date":
		return fhirVal.Date, true
	case "kind":
		return fhirVal.Kind, true
	case "purpose":
		return fhirVal.Purpose, true
	case "fhirversion":
		return fhirVal.FhirVersion, true
	case "mapping":
		return fhirVal.Mapping, true
	case "version":
		return fhirVal.Version, true
	case "title":
		return fhirVal.Title, true
	case "status":
		return fhirVal.Status, true
	case "contact":
		return fhirVal.Contact, true
	case "basedefinition":
		return fhirVal.BaseDefinition, true
	case "contexttype":
		return fhirVal.ContextType, true
	case "differential":
		return fhirVal.Differential, true
	case "url":
		return fhirVal.Url, true
	case "keyword":
		return fhirVal.Keyword, true
	case "abstract":
		return fhirVal.Abstract, true

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
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"type": &FieldTypeSupport{"string", false, false},
		"context": &FieldTypeSupport{"string", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"contextinvariant": &FieldTypeSupport{"string", true, false},
		"derivation": &FieldTypeSupport{"string", false, false},
		"snapshot": &FieldTypeSupport{"StructureDefinitionSnapshotComponent", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"kind": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"fhirversion": &FieldTypeSupport{"string", false, false},
		"mapping": &FieldTypeSupport{"StructureDefinitionMappingComponent", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"basedefinition": &FieldTypeSupport{"string", false, false},
		"contexttype": &FieldTypeSupport{"string", false, false},
		"differential": &FieldTypeSupport{"StructureDefinitionDifferentialComponent", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"keyword": &FieldTypeSupport{"Coding", true, false},
		"abstract": &FieldTypeSupport{"bool", false, true},

	}
}
