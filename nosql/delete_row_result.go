// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// DeleteRowResult The result of a DeleteRow operation.
type DeleteRowResult struct {

	// Convey the success or failure of the operation.
	IsSuccess *bool `mandatory:"false" json:"isSuccess"`

	// The version string associated with the existing row.
	// Returned if the delete fails due to options setting in the
	// request.
	ExistingVersion *string `mandatory:"false" json:"existingVersion"`

	// The map of values from a row.
	ExistingValue map[string]interface{} `mandatory:"false" json:"existingValue"`

	Usage *RequestUsage `mandatory:"false" json:"usage"`
}

func (m DeleteRowResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeleteRowResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
