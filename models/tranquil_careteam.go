
 package models
 
 import (
	 "fmt"
 )
 
 
 func (fhirVal *CareTeam) TupleName() (id string, resourceType string) {
	 return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
 }
 
 func (fhirVal *CareTeam) FieldByLowerName(nameLower string) (interface{}, bool) {
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
 	case "context":
		 return fhirVal.Context, true
 	case "period":
		 return fhirVal.Period, true
 	case "participant":
		 return fhirVal.Participant, true
 	case "reasoncode":
		 return fhirVal.ReasonCode, true
 	case "managingorganization":
		 return fhirVal.ManagingOrganization, true
 	case "subject":
		 return fhirVal.Subject, true
 	case "status":
		 return fhirVal.Status, true
 	case "category":
		 return fhirVal.Category, true
 	case "name":
		 return fhirVal.Name, true
 	case "reasonreference":
		 return fhirVal.ReasonReference, true
 	case "note":
		 return fhirVal.Note, true
 	case "identifier":
		 return fhirVal.Identifier, true
 
	 default:
		 return nil, false
	 }
 }
 
 func (fhirVal *CareTeam) FieldsToTypes() map[string]*FieldTypeSupport {
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
 		"context": &FieldTypeSupport{"Reference", false, true},
 		"period": &FieldTypeSupport{"Period", false, true},
 		"participant": &FieldTypeSupport{"CareTeamParticipantComponent", true, false},
 		"reasoncode": &FieldTypeSupport{"CodeableConcept", true, false},
 		"managingorganization": &FieldTypeSupport{"Reference", true, false},
 		"subject": &FieldTypeSupport{"Reference", false, true},
 		"status": &FieldTypeSupport{"string", false, false},
 		"category": &FieldTypeSupport{"CodeableConcept", true, false},
 		"name": &FieldTypeSupport{"string", false, false},
 		"reasonreference": &FieldTypeSupport{"Reference", true, false},
 		"note": &FieldTypeSupport{"Annotation", true, false},
 		"identifier": &FieldTypeSupport{"Identifier", true, false},
 
	 }
 }
 
 func (fhirVal *CareTeam) Truncate(fieldToKeep string) (TupleSupport, error) {
	 return nil, fmt.Errorf("Not yet supported")
 }
 