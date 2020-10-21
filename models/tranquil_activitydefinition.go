
package models


func (fhirVal *ActivityDefinition) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}

func (fhirVal *ActivityDefinition) FieldByLowerName(nameLower string) (interface{}, bool) {
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
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "kind":
		return fhirVal.Kind, true
	case "timingtiming":
		return fhirVal.TimingTiming, true
	case "code":
		return fhirVal.Code, true
	case "timingdatetime":
		return fhirVal.TimingDateTime, true
	case "productreference":
		return fhirVal.ProductReference, true
	case "experimental":
		return fhirVal.Experimental, true
	case "purpose":
		return fhirVal.Purpose, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "contact":
		return fhirVal.Contact, true
	case "title":
		return fhirVal.Title, true
	case "quantity":
		return fhirVal.Quantity, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "topic":
		return fhirVal.Topic, true
	case "contributor":
		return fhirVal.Contributor, true
	case "location":
		return fhirVal.Location, true
	case "productcodeableconcept":
		return fhirVal.ProductCodeableConcept, true
	case "dynamicvalue":
		return fhirVal.DynamicValue, true
	case "identifier":
		return fhirVal.Identifier, true
	case "version":
		return fhirVal.Version, true
	case "description":
		return fhirVal.Description, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "timingperiod":
		return fhirVal.TimingPeriod, true
	case "timingrange":
		return fhirVal.TimingRange, true
	case "dosage":
		return fhirVal.Dosage, true
	case "date":
		return fhirVal.Date, true
	case "publisher":
		return fhirVal.Publisher, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "url":
		return fhirVal.Url, true
	case "usage":
		return fhirVal.Usage, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "transform":
		return fhirVal.Transform, true
	case "participant":
		return fhirVal.Participant, true
	case "name":
		return fhirVal.Name, true
	case "status":
		return fhirVal.Status, true
	case "copyright":
		return fhirVal.Copyright, true
	case "library":
		return fhirVal.Library, true

	default:
		return nil, false
	}
}

func (fhirVal *ActivityDefinition) FieldsToTypes() map[string]*FieldTypeSupport {
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
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"kind": &FieldTypeSupport{"string", false, false},
		"timingtiming": &FieldTypeSupport{"Timing", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"timingdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"productreference": &FieldTypeSupport{"Reference", false, true},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"purpose": &FieldTypeSupport{"string", false, false},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"title": &FieldTypeSupport{"string", false, false},
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"productcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"dynamicvalue": &FieldTypeSupport{"ActivityDefinitionDynamicValueComponent", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"version": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"timingperiod": &FieldTypeSupport{"Period", false, true},
		"timingrange": &FieldTypeSupport{"Range", false, true},
		"dosage": &FieldTypeSupport{"Dosage", true, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"url": &FieldTypeSupport{"string", false, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"transform": &FieldTypeSupport{"Reference", false, true},
		"participant": &FieldTypeSupport{"ActivityDefinitionParticipantComponent", true, false},
		"name": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"library": &FieldTypeSupport{"Reference", true, false},

	}
}
