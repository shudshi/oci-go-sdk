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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnableModuleStreamOnManagedInstanceGroupDetails The work request details for the module stream operation on the managed instance group.
type EnableModuleStreamOnManagedInstanceGroupDetails struct {

	// The name of a module.
	ModuleName *string `mandatory:"false" json:"moduleName"`

	// The name of a stream of the specified module.
	StreamName *string `mandatory:"false" json:"streamName"`

	WorkRequestDetails *WorkRequestDetails `mandatory:"false" json:"workRequestDetails"`
}

func (m EnableModuleStreamOnManagedInstanceGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableModuleStreamOnManagedInstanceGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
