// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSqlJobDetails The details specific to the SQL job request.
type UpdateSqlJobDetails struct {

	// The description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The job timeout duration, which is expressed like "1h 10m 15s".
	Timeout *string `mandatory:"false" json:"timeout"`

	ResultLocation JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`

	ScheduleDetails *JobScheduleDetails `mandatory:"false" json:"scheduleDetails"`

	// The SQL text to be executed as part of the job.
	SqlText *string `mandatory:"false" json:"sqlText"`

	InBinds *JobInBindsDetails `mandatory:"false" json:"inBinds"`

	OutBinds *JobOutBindsDetails `mandatory:"false" json:"outBinds"`

	// The database user name used to execute the SQL job. If the job is being executed on a
	// Managed Database Group, then the user name should exist on all the databases in the
	// group with the same password.
	UserName *string `mandatory:"false" json:"userName"`

	// The password for the database user name used to execute the SQL job.
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Named Credentials containing password secret.
	NamedCredentialId *string `mandatory:"false" json:"namedCredentialId"`

	SqlType SqlJobSqlTypeEnum `mandatory:"false" json:"sqlType,omitempty"`

	// The role of the database user. Indicates whether the database user is a normal user or sysdba.
	Role SqlJobRoleEnum `mandatory:"false" json:"role,omitempty"`
}

// GetDescription returns Description
func (m UpdateSqlJobDetails) GetDescription() *string {
	return m.Description
}

// GetTimeout returns Timeout
func (m UpdateSqlJobDetails) GetTimeout() *string {
	return m.Timeout
}

// GetResultLocation returns ResultLocation
func (m UpdateSqlJobDetails) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

// GetScheduleDetails returns ScheduleDetails
func (m UpdateSqlJobDetails) GetScheduleDetails() *JobScheduleDetails {
	return m.ScheduleDetails
}

func (m UpdateSqlJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSqlJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlJobSqlTypeEnum(string(m.SqlType)); !ok && m.SqlType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlType: %s. Supported values are: %s.", m.SqlType, strings.Join(GetSqlJobSqlTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlJobRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetSqlJobRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateSqlJobDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSqlJobDetails UpdateSqlJobDetails
	s := struct {
		DiscriminatorParam string `json:"jobType"`
		MarshalTypeUpdateSqlJobDetails
	}{
		"SQL",
		(MarshalTypeUpdateSqlJobDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSqlJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description       *string                    `json:"description"`
		Timeout           *string                    `json:"timeout"`
		ResultLocation    jobexecutionresultlocation `json:"resultLocation"`
		ScheduleDetails   *JobScheduleDetails        `json:"scheduleDetails"`
		SqlText           *string                    `json:"sqlText"`
		InBinds           *JobInBindsDetails         `json:"inBinds"`
		OutBinds          *JobOutBindsDetails        `json:"outBinds"`
		SqlType           SqlJobSqlTypeEnum          `json:"sqlType"`
		UserName          *string                    `json:"userName"`
		Password          *string                    `json:"password"`
		SecretId          *string                    `json:"secretId"`
		NamedCredentialId *string                    `json:"namedCredentialId"`
		Role              SqlJobRoleEnum             `json:"role"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Timeout = model.Timeout

	nn, e = model.ResultLocation.UnmarshalPolymorphicJSON(model.ResultLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocation = nn.(JobExecutionResultLocation)
	} else {
		m.ResultLocation = nil
	}

	m.ScheduleDetails = model.ScheduleDetails

	m.SqlText = model.SqlText

	m.InBinds = model.InBinds

	m.OutBinds = model.OutBinds

	m.SqlType = model.SqlType

	m.UserName = model.UserName

	m.Password = model.Password

	m.SecretId = model.SecretId

	m.NamedCredentialId = model.NamedCredentialId

	m.Role = model.Role

	return
}
