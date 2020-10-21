
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
	case "relatedartifact":
		return fhirVal.RelatedArtifact, true
	case "library":
		return fhirVal.Library, true
	case "code":
		return fhirVal.Code, true
	case "timingrange":
		return fhirVal.TimingRange, true
	case "url":
		return fhirVal.Url, true
	case "status":
		return fhirVal.Status, true
	case "purpose":
		return fhirVal.Purpose, true
	case "contact":
		return fhirVal.Contact, true
	case "experimental":
		return fhirVal.Experimental, true
	case "timingtiming":
		return fhirVal.TimingTiming, true
	case "dynamicvalue":
		return fhirVal.DynamicValue, true
	case "productcodeableconcept":
		return fhirVal.ProductCodeableConcept, true
	case "usage":
		return fhirVal.Usage, true
	case "effectiveperiod":
		return fhirVal.EffectivePeriod, true
	case "usecontext":
		return fhirVal.UseContext, true
	case "kind":
		return fhirVal.Kind, true
	case "version":
		return fhirVal.Version, true
	case "name":
		return fhirVal.Name, true
	case "lastreviewdate":
		return fhirVal.LastReviewDate, true
	case "location":
		return fhirVal.Location, true
	case "transform":
		return fhirVal.Transform, true
	case "identifier":
		return fhirVal.Identifier, true
	case "jurisdiction":
		return fhirVal.Jurisdiction, true
	case "topic":
		return fhirVal.Topic, true
	case "quantity":
		return fhirVal.Quantity, true
	case "publisher":
		return fhirVal.Publisher, true
	case "productreference":
		return fhirVal.ProductReference, true
	case "dosage":
		return fhirVal.Dosage, true
	case "timingperiod":
		return fhirVal.TimingPeriod, true
	case "participant":
		return fhirVal.Participant, true
	case "bodysite":
		return fhirVal.BodySite, true
	case "approvaldate":
		return fhirVal.ApprovalDate, true
	case "contributor":
		return fhirVal.Contributor, true
	case "copyright":
		return fhirVal.Copyright, true
	case "timingdatetime":
		return fhirVal.TimingDateTime, true
	case "title":
		return fhirVal.Title, true
	case "date":
		return fhirVal.Date, true
	case "description":
		return fhirVal.Description, true

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
		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
		"library": &FieldTypeSupport{"Reference", true, false},
		"code": &FieldTypeSupport{"CodeableConcept", false, true},
		"timingrange": &FieldTypeSupport{"Range", false, true},
		"url": &FieldTypeSupport{"string", false, false},
		"status": &FieldTypeSupport{"string", false, false},
		"purpose": &FieldTypeSupport{"string", false, false},
		"contact": &FieldTypeSupport{"ContactDetail", true, false},
		"experimental": &FieldTypeSupport{"bool", false, true},
		"timingtiming": &FieldTypeSupport{"Timing", false, true},
		"dynamicvalue": &FieldTypeSupport{"ActivityDefinitionDynamicValueComponent", true, false},
		"productcodeableconcept": &FieldTypeSupport{"CodeableConcept", false, true},
		"usage": &FieldTypeSupport{"string", false, false},
		"effectiveperiod": &FieldTypeSupport{"Period", false, true},
		"usecontext": &FieldTypeSupport{"UsageContext", true, false},
		"kind": &FieldTypeSupport{"string", false, false},
		"version": &FieldTypeSupport{"string", false, false},
		"name": &FieldTypeSupport{"string", false, false},
		"lastreviewdate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"location": &FieldTypeSupport{"Reference", false, true},
		"transform": &FieldTypeSupport{"Reference", false, true},
		"identifier": &FieldTypeSupport{"Identifier", true, false},
		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
		"topic": &FieldTypeSupport{"CodeableConcept", true, false},
		"quantity": &FieldTypeSupport{"Quantity", false, true},
		"publisher": &FieldTypeSupport{"string", false, false},
		"productreference": &FieldTypeSupport{"Reference", false, true},
		"dosage": &FieldTypeSupport{"Dosage", true, false},
		"timingperiod": &FieldTypeSupport{"Period", false, true},
		"participant": &FieldTypeSupport{"ActivityDefinitionParticipantComponent", true, false},
		"bodysite": &FieldTypeSupport{"CodeableConcept", true, false},
		"approvaldate": &FieldTypeSupport{"FHIRDateTime", false, true},
		"contributor": &FieldTypeSupport{"Contributor", true, false},
		"copyright": &FieldTypeSupport{"string", false, false},
		"timingdatetime": &FieldTypeSupport{"FHIRDateTime", false, true},
		"title": &FieldTypeSupport{"string", false, false},
		"date": &FieldTypeSupport{"FHIRDateTime", false, true},
		"description": &FieldTypeSupport{"string", false, false},

	}
}
