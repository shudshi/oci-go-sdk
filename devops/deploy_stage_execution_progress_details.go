// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DeployStageExecutionProgressDetails Details about stage execution for each target environment.
type DeployStageExecutionProgressDetails struct {

	// The function ID, instance ID or the cluster ID. For Wait stage it will be the stage ID.
	TargetId *string `mandatory:"false" json:"targetId"`

	// Group for the target environment for example, the batch number for an Instance Group deployment.
	TargetGroup *string `mandatory:"false" json:"targetGroup"`

	// Details about all the steps for one target environment.
	Steps []DeployStageExecutionStep `mandatory:"false" json:"steps"`

	// Details about all the rollback steps for one target environment.
	RollbackSteps []DeployStageExecutionStep `mandatory:"false" json:"rollbackSteps"`
}

func (m DeployStageExecutionProgressDetails) String() string {
	return common.PointerString(m)
}
