
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
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "copyright":
		return fhirVal.Copyright, true
	case "status":
		return fhirVal.Status, true
	case "experimental":
		return fhirVal.Experimental, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "contributor":
		return fhirVal.Contributor, true
	case "timingtiming":
		return fhirVal.TimingTiming, true
	case "timingdatetime":
		return fhirVal.TimingDateTime, true
	case "quantity":
		return fhirVal.Quantity, true
	case "url":
		return fhirVal.Url, true
	case "dynamicvalue":
		return fhirVal.DynamicValue, true
	case "date":
		return fhirVal.Date, true
	case "contact":
		return fhirVal.Contact, true
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "productreference":
		return fhirVal.ProductReference, true
	case "dosage":
		return fhirVal.Dosage, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "identifier":
		return fhirVal.Identifier, true
	case "usage":
		return fhirVal.Usage, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "topic":
		return fhirVal.Topic, true
	case "timingperiod":
		return fhirVal.TimingPeriod, true
	case "productcodeableconcept":
		return fhirVal.ProductCodeableConcept, true
	case "publisher":
		return fhirVal.Publisher, true
	case "title":
		return fhirVal.Title, true
	case "description":
		return fhirVal.Description, true
	case "kind":
		return fhirVal.Kind, true
	case "timingrange":
		return fhirVal.TimingRange, true
	case "name":
		return fhirVal.Name, true
	case "purpose":
		return fhirVal.Purpose, true
	case "library":
		return fhirVal.Library, true
	case "participant":
		return fhirVal.Participant, true
	case "transform":
		return fhirVal.Transform, true
	case "version":
		return fhirVal.Version, true
	case "location":
		return fhirVal.Location, true
	case "code":
		return fhirVal.Code, true

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
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"copyright": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"timingtiming": &FieldTypeSupport{"Timing", false, true},
		"timingdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"dynamicvalue": &FieldTypeSupport{"ActivityDefinitionDynamicValueComponent", true, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"productreference": &FieldTypeSupport{"Reference", false, true},
		"dosage": &FieldTypeSupport{"Dosage", true, false},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"usage": &FieldTypeSupport{"string", false, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"timingperiod": &FieldTypeSupport{"Period", false, true},
		"productcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"title": &FieldTypeSupport{"string", false, false},
		"description": &FieldTypeSupport{"string", false, false},
		"kind": &FieldTypeSupport{"string", false, false},
		"timingrange": &FieldTypeSupport{"Range", false, true},
		"name": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"library": &FieldTypeSupport{"Reference", true, false},
		"participant": &FieldTypeSupport{"ActivityDefinitionParticipantComponent", true, false},
		"transform": &FieldTypeSupport{"Reference", false, true},
		"version": &FieldTypeSupport{"string", false, false},
		"location": &FieldTypeSupport{"Reference", false, true},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},

	}
}
