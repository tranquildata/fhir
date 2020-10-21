
package models


func (fhirVal *CapabilityStatement) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *CapabilityStatement) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "status":
		return fhirVal.Status, true
	case "date":
		return fhirVal.Date, true
	case "software":
		return fhirVal.Software, true
	case "acceptunknown":
		return fhirVal.AcceptUnknown, true
	case "profile":
		return fhirVal.Profile, true
	case "name":
		return fhirVal.Name, true
	case "instantiates":
		return fhirVal.Instantiates, true
	case "fhirversion":
		return fhirVal.FhirVersion, true
	case "document":
		return fhirVal.Document, true
	case "url":
		return fhirVal.Url, true
	case "version":
		return fhirVal.Version, true
	case "title":
		return fhirVal.Title, true
	case "publisher":
		return fhirVal.Publisher, true
	case "contact":
		return fhirVal.Contact, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "kind":
		return fhirVal.Kind, true
	case "implementation":
		return fhirVal.Implementation, true
	case "rest":
		return fhirVal.Rest, true
	case "purpose":
		return fhirVal.Purpose, true
	case "copyright":
		return fhirVal.Copyright, true
	case "experimental":
		return fhirVal.Experimental, true
	case "description":
		return fhirVal.Description, true
	case "format":
		return fhirVal.Format, true
	case "patchformat":
		return fhirVal.PatchFormat, true
	case "implementationguide":
		return fhirVal.ImplementationGuide, true
	case "messaging":
		return fhirVal.Messaging, true

	default:
		return nil, false
	}
}

func (fhirVal *CapabilityStatement) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"status": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"software": &FieldTypeSupport{"CapabilityStatementSoftwareComponent", false, true},
		"acceptunknown": &FieldTypeSupport{"string", false, false},
		"profile": &FieldTypeSupport{"Reference", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"instantiates": &FieldTypeSupport{"string", true, false},
		"fhirversion": &FieldTypeSupport{"string", false, false},
		"document": &FieldTypeSupport{"CapabilityStatementDocumentComponent", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"kind": &FieldTypeSupport{"string", false, false},
		"implementation": &FieldTypeSupport{"CapabilityStatementImplementationComponent", false, true},
		"rest": &FieldTypeSupport{"CapabilityStatementRestComponent", true, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"description": &FieldTypeSupport{"string", false, false},
		"format": &FieldTypeSupport{"string", true, false},
		"patchformat": &FieldTypeSupport{"string", true, false},
		"implementationguide": &FieldTypeSupport{"string", true, false},
		"messaging": &FieldTypeSupport{"CapabilityStatementMessagingComponent", true, false},

	}
}