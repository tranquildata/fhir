/*
 * Copyright (c) 2020, Tranquil Data, Inc. All rights reserved.
 */

package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//Format with the package name and the original type's import string
const FileHeaderMaterial = `
package %s

`

const TupleNameDomainResource = `
func (fhirVal *%s) TupleName() (id string, resourceType string) {
	return fhirVal.DomainResource.Resource.Id, fhirVal.DomainResource.Resource.ResourceType
}
`

//Need to format with the actual struct type name
const FieldToTypesPreamble = `
func (fhirVal *%s) FieldsToTypes() map[string]*FieldTypeSupport {
	return map[string]*FieldTypeSupport {

	`

const FieldToTypesPostamble = `
	}
}
`

//Need to format with lowercase fieldname, type name, array flag (bool), struct pointer flag (bool)
const FieldToTypesEntry = `		"%s": &FieldTypeSupport{"%s", %t, %t},
`

//Need to format with the actual struct type name
const FieldByLowerNamePreamble = `
func (fhirVal *%s) FieldByLowerName(nameLower string) (interface{}, bool) {
	switch nameLower {
`

const FieldByLowerNamePostamble = `
	default:
		return nil, false
	}
}
`

const ResourceFields = `
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
`

const ResourceFieldMappings = `		"resourcetype": &FieldTypeSupport{"string", false, false},
		"id": &FieldTypeSupport{"string", false, false},
		"meta": &FieldTypeSupport{"Meta", false, true},
		"implicitrules": &FieldTypeSupport{"string", false, false},
		"language": &FieldTypeSupport{"string", false, false},
`

const DomainResourceFields = ResourceFields + `
	case "text":
		return fhirVal.Text, true
	case "contained":
		return fhirVal.Contained, true
	case "extension":
		return fhirVal.Extension, true
	case "modifierextension":
		return fhirVal.ModifierExtension, true
`

const DomainResourceFieldMappings = ResourceFieldMappings + `
		"text": &FieldTypeSupport{"Narrative", false, true},
		"Contained": &FieldTypeSupport{"Containedresources", false, false},
		"extension": &FieldTypeSupport{"Extension", true, false},
		"modifierextension": &FieldTypeSupport{"Extension", true, false},						
`

type fhirStruct struct {
	name              string
	fieldNamesToTypes map[string]ast.Expr
	includedStructs   []string
}

func extractTypeName(filename string) string {
	_, file := filepath.Split(filename)
	if len(file) > 0 {
		return strings.ToLower(strings.TrimSuffix(file, ".go"))
	}
	return ""
}

func toFieldNameAndArrayStructFlags(expr ast.Expr) (string, bool, bool, error) {
	var err error
	switch ex := expr.(type) {
	case *ast.StarExpr: //X is usually an Ident
		if name, _, _, err := toFieldNameAndArrayStructFlags(ex.X); err == nil {
			return name, false, true, nil
		}
	case *ast.Ident: //including another struct (DomainResource or Resource)
		return ex.Name, false, false, nil
	case *ast.ArrayType: //arrays of idents mostly
		if name, _, _, err := toFieldNameAndArrayStructFlags(ex.Elt); err == nil {
			return name, true, false, nil
		}
	}
	return "", false, false, err
}

func newFhirStruct(name string, fields *ast.FieldList) (*fhirStruct, error) {
	if fields == nil || fields.List == nil || len(fields.List) == 0 {
		return nil, fmt.Errorf("No fields")
	}
	tranqFields := map[string]ast.Expr{}
	includes := []string{}
	for _, astField := range fields.List {
		//dealing with included structs
		if astField.Names == nil || len(astField.Names) == 0 {
			//DomainResource, Resource, BackboneElement, Element, Quantity
			inlineName := astField.Type.(*ast.Ident).Name
			includes = append(includes, inlineName)
		} else {
			for _, nm := range astField.Names {
				tranqFields[nm.Name] = astField.Type
			}
		}
	}
	return &fhirStruct{
		name:              name,
		fieldNamesToTypes: tranqFields,
		includedStructs:   includes,
	}, nil
}

func ParseFhirModel(filename string) (*fhirStruct, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, err
	}
	typenameLower := extractTypeName(filename)
	if typenameLower == "" {
		return nil, fmt.Errorf("Could not extract type name for: %s", filename)
	}
	var theStruct *fhirStruct = nil
	var foundStruct **fhirStruct = &theStruct
	ast.Inspect(f, func(nd ast.Node) bool {
		switch node := nd.(type) {
		case *ast.TypeSpec:
			switch typeDec := node.Type.(type) {
			case *ast.StructType:
				if nm := node.Name; nm != nil {
					if strings.ToLower(nm.Name) == typenameLower {
						if st, err := newFhirStruct(nm.Name, typeDec.Fields); err == nil {
							*foundStruct = st
						}
					}
				}
			}
		}
		return true
	})
	if theStruct == nil {
		return nil, fmt.Errorf("Could not find a fhir type for: %s", filename)
	}
	return theStruct, nil
}

func EmitFileHeader(packageName string) string {
	return fmt.Sprintf(FileHeaderMaterial, packageName)
}

func EmitFieldByLowerNamePreamble(fhirType *fhirStruct) string {
	return fmt.Sprintf(FieldByLowerNamePreamble, fhirType.name)
}

func EmitInlinedStructCode(inlinedStructName string) string {
	switch inlinedStructName {
	case "Resource":
		return ResourceFields
	case "DomainResource":
		return DomainResourceFields
	default:
		return ""
	}
}

func EmitInlinedStructFieldEntries(inlinedStructName string) string {
	switch inlinedStructName {
	case "Resource":
		return ResourceFieldMappings
	case "DomainResource":
		return DomainResourceFieldMappings
	default:
		return ""
	}
}

func EmitFieldToTypesPreamble(fhirType *fhirStruct) string {
	return fmt.Sprintf(FieldToTypesPreamble, fhirType.name)
}

func EmitFieldToTypesEntry(fieldName string, fieldType ast.Expr) (string, error) {
	name, isArray, isStructPtr, err := toFieldNameAndArrayStructFlags(fieldType)
	if err != nil {
		return "", err
	}
	return EmitFieldToTypesEntryString(fieldName, name, isArray, isStructPtr), nil
}

func EmitFieldToTypesEntryString(fieldName string, typeName string, isArray bool, isStructPtr bool) string {
	return fmt.Sprintf(FieldToTypesEntry, strings.ToLower(fieldName), typeName, isArray, isStructPtr)
}

func EmitTupleNameDomainResource(fhirType *fhirStruct) string {
	return fmt.Sprintf(TupleNameDomainResource, fhirType.name)
}

func EmitTupleName(fhirType *fhirStruct) (string, error) {
	for _, includedStruct := range fhirType.includedStructs {
		switch strings.ToLower(includedStruct) {
		case "domainresource":
			return EmitTupleNameDomainResource(fhirType), nil
		}
	}
	return "", fmt.Errorf("No tuple name function found: %s", fhirType.name)
}

func EmitTupleCode(fhirType *fhirStruct, destPackage string) (string, error) {
	bldr := strings.Builder{}
	fieldToTypesBuilder := strings.Builder{}
	fieldByLowerBuilder := strings.Builder{}
	//@TODO: need to add inlined fields to fieldToTypes map
	bldr.WriteString(EmitFileHeader(destPackage))
	tupleName, err := EmitTupleName(fhirType)
	if err != nil {
		return "", err
	}
	bldr.WriteString(tupleName)
	fieldByLowerBuilder.WriteString(EmitFieldByLowerNamePreamble(fhirType))
	fieldToTypesBuilder.WriteString(EmitFieldToTypesPreamble(fhirType))
	for _, includedStruct := range fhirType.includedStructs {
		fieldByLowerBuilder.WriteString(EmitInlinedStructCode(includedStruct))
		fieldToTypesBuilder.WriteString(EmitInlinedStructFieldEntries(includedStruct))
	}
	fieldByLowerBuilder.WriteString(FieldByLowerNamePostamble)
	for fieldName, fieldTy := range fhirType.fieldNamesToTypes {
		//add field type accessor stuff
		if entry, err := EmitFieldToTypesEntry(fieldName, fieldTy); err == nil {
			fieldToTypesBuilder.WriteString(entry)
		} else {
			return "", err
		}
	}
	fieldToTypesBuilder.WriteString(FieldToTypesPostamble)
	bldr.WriteString(fieldByLowerBuilder.String())
	bldr.WriteString(fieldToTypesBuilder.String())
	return bldr.String(), nil
}

func ProcessDirectory(directory string, targetPackage string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	processed := []string{}
	if err != nil {
		return processed, err
	}
	for _, file := range files {
		if file.IsDir() ||
			strings.HasPrefix(file.Name(), "tranquil") ||
			strings.HasSuffix(file.Name(), "_test.go") {
			fmt.Printf("Skipping: %s\n", file.Name())
			continue
		}
		fmt.Printf("Processing: %s\n", file.Name())
		path := filepath.Join(directory, file.Name())
		st, err := ParseFhirModel(path)
		if err != nil {
			fmt.Printf("Error encountered, not emitting code for: %s. %s", file.Name(), err.Error())
			continue
		}
		emitted, err := EmitTupleCode(st, targetPackage)
		if err != nil {
			fmt.Printf("Error encountered, not emitting code for: %s. %s", file.Name(), err.Error())
			continue
		}
		wpath := filepath.Join(directory, "tranquil_"+file.Name())
		err = ioutil.WriteFile(wpath, []byte(emitted), 0644)
		if err != nil {
			return processed, err
		}
		processed = append(processed, file.Name())
	}
	return processed, nil
}
