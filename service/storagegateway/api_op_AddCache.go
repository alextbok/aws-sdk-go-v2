// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddCacheInput struct {
	_ struct{} `type:"structure"`

	// An array of strings that identify disks that are to be configured as working
	// storage. Each string has a minimum length of 1 and maximum length of 300.
	// You can get the disk IDs from the ListLocalDisks API.
	//
	// DiskIds is a required field
	DiskIds []string `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s AddCacheInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddCacheInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddCacheInput"}

	if s.DiskIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskIds"))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddCacheOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s AddCacheOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddCache = "AddCache"

// AddCacheRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Configures one or more gateway local disks as cache for a gateway. This operation
// is only supported in the cached volume, tape and file gateway type (see Storage
// Gateway Concepts (https://docs.aws.amazon.com/storagegateway/latest/userguide/StorageGatewayConcepts.html)).
//
// In the request, you specify the gateway Amazon Resource Name (ARN) to which
// you want to add cache, and one or more disk IDs that you want to configure
// as cache.
//
//    // Example sending a request using AddCacheRequest.
//    req := client.AddCacheRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/AddCache
func (c *Client) AddCacheRequest(input *AddCacheInput) AddCacheRequest {
	op := &aws.Operation{
		Name:       opAddCache,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddCacheInput{}
	}

	req := c.newRequest(op, input, &AddCacheOutput{})

	return AddCacheRequest{Request: req, Input: input, Copy: c.AddCacheRequest}
}

// AddCacheRequest is the request type for the
// AddCache API operation.
type AddCacheRequest struct {
	*aws.Request
	Input *AddCacheInput
	Copy  func(*AddCacheInput) AddCacheRequest
}

// Send marshals and sends the AddCache API request.
func (r AddCacheRequest) Send(ctx context.Context) (*AddCacheResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddCacheResponse{
		AddCacheOutput: r.Request.Data.(*AddCacheOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddCacheResponse is the response type for the
// AddCache API operation.
type AddCacheResponse struct {
	*AddCacheOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddCache request.
func (r *AddCacheResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
