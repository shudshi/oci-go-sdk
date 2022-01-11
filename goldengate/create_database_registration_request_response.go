// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
)

// CreateDatabaseRegistrationRequest wrapper for the CreateDatabaseRegistration operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreateDatabaseRegistration.go.html to see an example of how to use CreateDatabaseRegistrationRequest.
type CreateDatabaseRegistrationRequest struct {

	// Specification of the DatabaseRegistration to create.
	CreateDatabaseRegistrationDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried, in case of a timeout or server error, without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated before then due to conflicting operations. For example, if a resource has been deleted and purged from the system, then a retry of the original creation request is rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateDatabaseRegistrationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateDatabaseRegistrationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateDatabaseRegistrationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateDatabaseRegistrationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateDatabaseRegistrationResponse wrapper for the CreateDatabaseRegistration operation
type CreateDatabaseRegistrationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DatabaseRegistration instance
	DatabaseRegistration `presentIn:"body"`

	// A unique Oracle-assigned identifier for an asynchronous request. You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateDatabaseRegistrationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateDatabaseRegistrationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
