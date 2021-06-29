// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v43/common"
	"net/http"
)

// ListDataFlowValidationsRequest wrapper for the ListDataFlowValidations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDataFlowValidations.go.html to see an example of how to use ListDataFlowValidationsRequest.
type ListDataFlowValidationsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Used to filter by the key of the object.
	Key *string `mandatory:"false" contributesTo:"query" name:"key"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier *string `mandatory:"false" contributesTo:"query" name:"identifier"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListDataFlowValidationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDataFlowValidationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataFlowValidationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataFlowValidationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataFlowValidationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataFlowValidationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataFlowValidationsResponse wrapper for the ListDataFlowValidations operation
type ListDataFlowValidationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataFlowValidationSummaryCollection instances
	DataFlowValidationSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListDataFlowValidationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataFlowValidationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataFlowValidationsSortByEnum Enum with underlying type: string
type ListDataFlowValidationsSortByEnum string

// Set of constants representing the allowable values for ListDataFlowValidationsSortByEnum
const (
	ListDataFlowValidationsSortByTimeCreated ListDataFlowValidationsSortByEnum = "TIME_CREATED"
	ListDataFlowValidationsSortByDisplayName ListDataFlowValidationsSortByEnum = "DISPLAY_NAME"
)

var mappingListDataFlowValidationsSortBy = map[string]ListDataFlowValidationsSortByEnum{
	"TIME_CREATED": ListDataFlowValidationsSortByTimeCreated,
	"DISPLAY_NAME": ListDataFlowValidationsSortByDisplayName,
}

// GetListDataFlowValidationsSortByEnumValues Enumerates the set of values for ListDataFlowValidationsSortByEnum
func GetListDataFlowValidationsSortByEnumValues() []ListDataFlowValidationsSortByEnum {
	values := make([]ListDataFlowValidationsSortByEnum, 0)
	for _, v := range mappingListDataFlowValidationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDataFlowValidationsSortOrderEnum Enum with underlying type: string
type ListDataFlowValidationsSortOrderEnum string

// Set of constants representing the allowable values for ListDataFlowValidationsSortOrderEnum
const (
	ListDataFlowValidationsSortOrderAsc  ListDataFlowValidationsSortOrderEnum = "ASC"
	ListDataFlowValidationsSortOrderDesc ListDataFlowValidationsSortOrderEnum = "DESC"
)

var mappingListDataFlowValidationsSortOrder = map[string]ListDataFlowValidationsSortOrderEnum{
	"ASC":  ListDataFlowValidationsSortOrderAsc,
	"DESC": ListDataFlowValidationsSortOrderDesc,
}

// GetListDataFlowValidationsSortOrderEnumValues Enumerates the set of values for ListDataFlowValidationsSortOrderEnum
func GetListDataFlowValidationsSortOrderEnumValues() []ListDataFlowValidationsSortOrderEnum {
	values := make([]ListDataFlowValidationsSortOrderEnum, 0)
	for _, v := range mappingListDataFlowValidationsSortOrder {
		values = append(values, v)
	}
	return values
}
