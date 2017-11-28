// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

type ObjectStorageClient struct {
	common.BaseClient
}

//Create a new default ObjectStorage client for a given region
func NewObjectStorageClientForRegion(region common.Region) (client ObjectStorageClient) {
	client = ObjectStorageClient{BaseClient: common.NewClientForRegion(region)}

	client.Host = fmt.Sprintf(common.DefaultHostUrlTemplate, "objectstorage", string(region))
	client.BasePath = "20160918"
	return
}

// Aborts an in-progress multipart upload and deletes all parts that have been uploaded.
func (client ObjectStorageClient) AbortMultipartUpload(ctx context.Context, request AbortMultipartUploadRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Commits a multipart upload, which involves checking part numbers and ETags of the parts, to create an aggregate object.
func (client ObjectStorageClient) CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Creates a bucket in the given namespace with a bucket name and optional user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) CreateBucket(ctx context.Context, request CreateBucketRequest) (response CreateBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Starts a new multipart upload to a specific object in the given bucket in the given namespace.
func (client ObjectStorageClient) CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest) (response CreateMultipartUploadResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Create a pre-authenticated request specific to the bucket
func (client ObjectStorageClient) CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest) (response CreatePreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/p/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Deletes a bucket if it is already empty. If the bucket is not empty, use DeleteObject first.
func (client ObjectStorageClient) DeleteBucket(ctx context.Context, request DeleteBucketRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes an object.
func (client ObjectStorageClient) DeleteObject(ctx context.Context, request DeleteObjectRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the bucket level pre-authenticateted request
func (client ObjectStorageClient) DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/p/{parId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Gets the current representation of the given bucket in the given namespace.
func (client ObjectStorageClient) GetBucket(ctx context.Context, request GetBucketRequest) (response GetBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the name of the namespace for the user making the request. An account name must be unique, must start with a
// letter, and can have up to 15 lowercase letters and numbers. You cannot use spaces or special characters.
func (client ObjectStorageClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the metadata and body of an object.
func (client ObjectStorageClient) GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Get the bucket level pre-authenticateted request
func (client ObjectStorageClient) GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest) (response GetPreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/{parId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Efficiently checks if a bucket exists and gets the current ETag for the bucket.
func (client ObjectStorageClient) HeadBucket(ctx context.Context, request HeadBucketRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Gets the user-defined metadata and entity tag for an object.
func (client ObjectStorageClient) HeadObject(ctx context.Context, request HeadObjectRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Gets a list of all `BucketSummary`s in a compartment. A `BucketSummary` contains only summary fields for the bucket
// and does not contain fields like the user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListBuckets(ctx context.Context, request ListBucketsRequest) (response ListBucketsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the parts of an in-progress multipart upload.
func (client ObjectStorageClient) ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest) (response ListMultipartUploadPartsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists all in-progress multipart uploads for the given bucket in the given namespace.
func (client ObjectStorageClient) ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest) (response ListMultipartUploadsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the objects in a bucket.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListObjects(ctx context.Context, request ListObjectsRequest) (response ListObjectsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// List pre-authenticated requests for the bucket
func (client ObjectStorageClient) ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest) (response ListPreauthenticatedRequestsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a new object or overwrites an existing one.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) PutObject(ctx context.Context, request PutObjectRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Performs a partial or full update of a bucket's user-defined metadata.
func (client ObjectStorageClient) UpdateBucket(ctx context.Context, request UpdateBucketRequest) (response UpdateBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Uploads a single part of a multipart upload.
func (client ObjectStorageClient) UploadPart(ctx context.Context, request UploadPartRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}
