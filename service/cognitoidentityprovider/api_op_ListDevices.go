// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to list the devices.
type ListDevicesInput struct {
	_ struct{} `type:"structure"`

	// The access tokens for the request to list devices.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true" sensitive:"true"`

	// The limit of the device request.
	Limit *int64 `type:"integer"`

	// The pagination token for the list request.
	PaginationToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDevicesInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}
	if s.PaginationToken != nil && len(*s.PaginationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PaginationToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response to list devices.
type ListDevicesOutput struct {
	_ struct{} `type:"structure"`

	// The devices returned in the list devices response.
	Devices []DeviceType `type:"list"`

	// The pagination token for the list device response.
	PaginationToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDevicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDevices = "ListDevices"

// ListDevicesRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the devices.
//
//    // Example sending a request using ListDevicesRequest.
//    req := client.ListDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListDevices
func (c *Client) ListDevicesRequest(input *ListDevicesInput) ListDevicesRequest {
	op := &aws.Operation{
		Name:       opListDevices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDevicesInput{}
	}

	req := c.newRequest(op, input, &ListDevicesOutput{})
	return ListDevicesRequest{Request: req, Input: input, Copy: c.ListDevicesRequest}
}

// ListDevicesRequest is the request type for the
// ListDevices API operation.
type ListDevicesRequest struct {
	*aws.Request
	Input *ListDevicesInput
	Copy  func(*ListDevicesInput) ListDevicesRequest
}

// Send marshals and sends the ListDevices API request.
func (r ListDevicesRequest) Send(ctx context.Context) (*ListDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDevicesResponse{
		ListDevicesOutput: r.Request.Data.(*ListDevicesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDevicesResponse is the response type for the
// ListDevices API operation.
type ListDevicesResponse struct {
	*ListDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDevices request.
func (r *ListDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
