
package models


func (fhirVal *ExpansionProfile) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ExpansionProfile) FieldByLowerName(nameLower string) (interface{}, bool) {
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

func (fhirVal *ExpansionProfile) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"excludenested": &FieldTypeSupport{"bool", false, true},
		"excludenotforui": &FieldTypeSupport{"bool", false, true},
		"excludepostcoordinated": &FieldTypeSupport{"bool", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"identifier": &FieldTypeSupport{"Identifier", false, true},
		"excludedsystem": &FieldTypeSupport{"ExpansionProfileExcludedSystemComponent", false, true},
		"activeonly": &FieldTypeSupport{"bool", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"publisher": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"description": &FieldTypeSupport{"string", false, false},
		"designation": &FieldTypeSupport{"ExpansionProfileDesignationComponent", false, true},
		"includedefinition": &FieldTypeSupport{"bool", false, true},
		"displaylanguage": &FieldTypeSupport{"string", false, false},
		"limitedexpansion": &FieldTypeSupport{"bool", false, true},
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"fixedversion": &FieldTypeSupport{"ExpansionProfileFixedVersionComponent", true, false},
		"includedesignations": &FieldTypeSupport{"bool", false, true},
		"name": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},

	}
}
