// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Ingestion API
//
// Use the Logging Ingestion API to ingest your application logs.
//

package loggingingestion

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// PutLogsDetails The request body for the PutLogs request.
type PutLogsDetails struct {

	// Required for identifying the version of the data format being used.
	// Permitted values include: "1.0"
	Specversion *string `mandatory:"true" json:"specversion"`

	// List of log-batches. Each batch has a single source, type and subject.
	LogEntryBatches []LogEntryBatch `mandatory:"true" json:"logEntryBatches"`
}

func (m PutLogsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PutLogsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
