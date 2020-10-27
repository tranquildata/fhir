
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *ResearchStudy) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *ResearchStudy) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "title":
		 return fhirVal.Title, true
 	case "partof":
		 return fhirVal.PartOf, true
 	case "focus":
		 return fhirVal.Focus, true
 	case "relatedartifact":
		 return fhirVal.RelatedArtifact, true
 	case "jurisdiction":
		 return fhirVal.Jurisdiction, true
 	case "sponsor":
		 return fhirVal.Sponsor, true
 	case "reasonstopped":
		 return fhirVal.ReasonStopped, true
 	case "note":
		 return fhirVal.Note, true
 	case "protocol":
		 return fhirVal.Protocol, true
 	case "status":
		 return fhirVal.Status, true
 	case "keyword":
		 return fhirVal.Keyword, true
 	case "description":
		 return fhirVal.Description, true
 	case "enrollment":
		 return fhirVal.Enrollment, true
 	case "principalinvestigator":
		 return fhirVal.PrincipalInvestigator, true
 	case "category":
		 return fhirVal.Category, true
 	case "contact":
		 return fhirVal.Contact, true
 	case "period":
		 return fhirVal.Period, true
 	case "arm":
		 return fhirVal.Arm, true
 	case "identifier":
		 return fhirVal.Identifier, true
 	case "site":
		 return fhirVal.Site, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *ResearchStudy) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"title": &FieldTypeSupport{"string", false, false},
 		"partof": &FieldTypeSupport{"Reference", true, false},
 		"focus": &FieldTypeSupport{"CodeableConcept", true, false},
 		"relatedartifact": &FieldTypeSupport{"RelatedArtifact", true, false},
 		"jurisdiction": &FieldTypeSupport{"CodeableConcept", true, false},
 		"sponsor": &FieldTypeSupport{"Reference", false, true},
 		"reasonstopped": &FieldTypeSupport{"CodeableConcept", false, true},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"protocol": &FieldTypeSupport{"Reference", true, false},
 		"status": &FieldTypeSupport{"string", false, false},
 		"keyword": &FieldTypeSupport{"CodeableConcept", true, false},
 		"description": &FieldTypeSupport{"string", false, false},
 		"enrollment": &FieldTypeSupport{"Reference", true, false},
 		"principalinvestigator": &FieldTypeSupport{"Reference", false, true},
 		"category": &FieldTypeSupport{"CodeableConcept", true, false},
 		"contact": &FieldTypeSupport{"ContactDetail", true, false},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"arm": &FieldTypeSupport{"ResearchStudyArmComponent", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 		"site": &FieldTypeSupport{"Reference", true, false},
 
	 }
 }
 
 func (fhirVal *ResearchStudy) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 