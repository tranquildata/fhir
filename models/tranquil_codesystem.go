
package models


func (fhirVal *CodeSystem) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *CodeSystem) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "casesensitive":
		return fhirVal.CaseSensitive, true
	case "valueset":
		return fhirVal.ValueSet, true
	case "property":
		return fhirVal.Property, true
	case "url":
		return fhirVal.Url, true
	case "date":
		return fhirVal.Date, true
	case "description":
		return fhirVal.Description, true
	case "purpose":
		return fhirVal.Purpose, true
	case "versionneeded":
		return fhirVal.VersionNeeded, true
	case "filter":
		return fhirVal.Filter, true
	case "concept":
		return fhirVal.Concept, true
	case "name":
		return fhirVal.Name, true
	case "publisher":
		return fhirVal.Publisher, true
	case "copyright":
		return fhirVal.Copyright, true
	case "compositional":
		return fhirVal.Compositional, true
	case "content":
		return fhirVal.Content, true
	case "count":
		return fhirVal.Count, true
	case "experimental":
		return fhirVal.Experimental, true
	case "contact":
		return fhirVal.Contact, true
	case "title":
		return fhirVal.Title, true
	case "status":
		return fhirVal.Status, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "hierarchymeaning":
		return fhirVal.HierarchyMeaning, true
	case "identifier":
		return fhirVal.Identifier, true
	case "version":
		return fhirVal.Version, true

	default:
		return nil, false
	}
}

func (fhirVal *CodeSystem) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"casesensitive": &FieldTypeSupport{"bool", false, true},
		"valueset": &FieldTypeSupport{"string", false, false},
		"property": &FieldTypeSupport{"CodeSystemPropertyComponent", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"versionneeded": &FieldTypeSupport{"bool", false, true},
		"filter": &FieldTypeSupport{"CodeSystemFilterComponent", true, false},
		"concept": &FieldTypeSupport{"CodeSystemConceptDefinitionComponent", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"compositional": &FieldTypeSupport{"bool", false, true},
		"content": &FieldTypeSupport{"string", false, false},
		"count": &FieldTypeSupport{"uint32", false, true},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"hierarchymeaning": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"version": &FieldTypeSupport{"string", false, false},

	}
}
