/*
 * Copyright (c) 2020, Tranquil Data, Inc. All rights reserved.
 */

package models2

func FieldNameToFHIRType(resourceType string, fieldName string) (string, bool) {
	key := resourceType + "." + fieldName
	if ty, found := fhirTypes[key]; found {
		return ty, true
	} else {
		return "", false
	}
}
