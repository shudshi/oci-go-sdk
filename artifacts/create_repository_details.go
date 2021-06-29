// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Images API
//
// API covering the Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as container images and repositories.
//

package artifacts

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v43/common"
)

// CreateRepositoryDetails Parameters needed to create an artifact repository.
type CreateRepositoryDetails interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the repository's compartment.
	GetCompartmentId() *string

	// Whether to make the repository immutable. The artifacts of an immutable repository cannot be overwritten.
	GetIsImmutable() *bool

	// A user-friendly display name for the repository. If not present, will be auto-generated. It can be modified later. Avoid entering confidential information.
	GetDisplayName() *string

	// A short description of the repository. It can be updated later.
	GetDescription() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createrepositorydetails struct {
	JsonData       []byte
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	IsImmutable    *bool                             `mandatory:"true" json:"isImmutable"`
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	Description    *string                           `mandatory:"false" json:"description"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	RepositoryType string                            `json:"repositoryType"`
}

// UnmarshalJSON unmarshals json
func (m *createrepositorydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreaterepositorydetails createrepositorydetails
	s := struct {
		Model Unmarshalercreaterepositorydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.IsImmutable = s.Model.IsImmutable
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.RepositoryType = s.Model.RepositoryType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createrepositorydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RepositoryType {
	case "GENERIC":
		mm := CreateGenericRepositoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCompartmentId returns CompartmentId
func (m createrepositorydetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetIsImmutable returns IsImmutable
func (m createrepositorydetails) GetIsImmutable() *bool {
	return m.IsImmutable
}

//GetDisplayName returns DisplayName
func (m createrepositorydetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m createrepositorydetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m createrepositorydetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createrepositorydetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createrepositorydetails) String() string {
	return common.PointerString(m)
}
