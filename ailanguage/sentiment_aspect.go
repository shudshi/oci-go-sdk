// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Infrastructure Artificial Intelligence Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately using our proven,
//     pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
//     This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"github.com/oracle/oci-go-sdk/v43/common"
)

// SentimentAspect Aspect object.
type SentimentAspect struct {

	// The number of Unicode code points preceding this entity in the submitted text.
	Offset *int `mandatory:"false" json:"offset"`

	// Length of text.
	Length *int `mandatory:"false" json:"length"`

	// Aspect text.
	Text *string `mandatory:"false" json:"text"`

	// The highest-score sentiment for the aspect text.
	Sentiment *string `mandatory:"false" json:"sentiment"`

	// Scores for each sentiment.
	// Example: "{\"positive\": 1.0, \"negative\": 0.0}"
	Scores map[string]float64 `mandatory:"false" json:"scores"`
}

func (m SentimentAspect) String() string {
	return common.PointerString(m)
}
