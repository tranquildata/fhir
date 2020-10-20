/*
 * Copyright (c) 2020, Tranquil Data, Inc. All rights reserved.
 */

package models

type FieldTypeSupport struct {
	Name          string
	Array         bool
	StructPointer bool
}

type TupleSupport interface {
	FieldByLowerName(nameLower string) (interface{}, bool)
	FieldsToTypes() map[string]*FieldTypeSupport
	TupleName() (id string, resourceType string)
}
