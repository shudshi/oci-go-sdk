// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
// Use the table of contents and search tool to explore the  OS Management Hub API.
//

package osmanagementhub

import (
	"strings"
)

// MirrorTypeEnum Enum with underlying type: string
type MirrorTypeEnum string

// Set of constants representing the allowable values for MirrorTypeEnum
const (
	MirrorTypeCustom    MirrorTypeEnum = "CUSTOM"
	MirrorTypeVendor    MirrorTypeEnum = "VENDOR"
	MirrorTypeVersioned MirrorTypeEnum = "VERSIONED"
)

var mappingMirrorTypeEnum = map[string]MirrorTypeEnum{
	"CUSTOM":    MirrorTypeCustom,
	"VENDOR":    MirrorTypeVendor,
	"VERSIONED": MirrorTypeVersioned,
}

var mappingMirrorTypeEnumLowerCase = map[string]MirrorTypeEnum{
	"custom":    MirrorTypeCustom,
	"vendor":    MirrorTypeVendor,
	"versioned": MirrorTypeVersioned,
}

// GetMirrorTypeEnumValues Enumerates the set of values for MirrorTypeEnum
func GetMirrorTypeEnumValues() []MirrorTypeEnum {
	values := make([]MirrorTypeEnum, 0)
	for _, v := range mappingMirrorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMirrorTypeEnumStringValues Enumerates the set of values in String for MirrorTypeEnum
func GetMirrorTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"VENDOR",
		"VERSIONED",
	}
}

// GetMappingMirrorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMirrorTypeEnum(val string) (MirrorTypeEnum, bool) {
	enum, ok := mappingMirrorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
