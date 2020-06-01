// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListIpRoutesInput struct {
	_ struct{} `type:"structure"`

	// Identifier (ID) of the directory for which you want to retrieve the IP addresses.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// Maximum number of items to return. If this value is zero, the maximum number
	// of items is specified by the limitations of the operation.
	Limit *int64 `type:"integer"`

	// The ListIpRoutes.NextToken value from a previous call to ListIpRoutes. Pass
	// null if this is the first call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListIpRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIpRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIpRoutesInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListIpRoutesOutput struct {
	_ struct{} `type:"structure"`

	// A list of IpRoutes.
	IpRoutesInfo []IpRouteInfo `type:"list"`

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to ListIpRoutes to retrieve the next set of
	// items.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListIpRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListIpRoutes = "ListIpRoutes"

// ListIpRoutesRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Lists the address blocks that you have added to a directory.
//
//    // Example sending a request using ListIpRoutesRequest.
//    req := client.ListIpRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/ListIpRoutes
func (c *Client) ListIpRoutesRequest(input *ListIpRoutesInput) ListIpRoutesRequest {
	op := &aws.Operation{
		Name:       opListIpRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListIpRoutesInput{}
	}

	req := c.newRequest(op, input, &ListIpRoutesOutput{})

	return ListIpRoutesRequest{Request: req, Input: input, Copy: c.ListIpRoutesRequest}
}

// ListIpRoutesRequest is the request type for the
// ListIpRoutes API operation.
type ListIpRoutesRequest struct {
	*aws.Request
	Input *ListIpRoutesInput
	Copy  func(*ListIpRoutesInput) ListIpRoutesRequest
}

// Send marshals and sends the ListIpRoutes API request.
func (r ListIpRoutesRequest) Send(ctx context.Context) (*ListIpRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIpRoutesResponse{
		ListIpRoutesOutput: r.Request.Data.(*ListIpRoutesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIpRoutesResponse is the response type for the
// ListIpRoutes API operation.
type ListIpRoutesResponse struct {
	*ListIpRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIpRoutes request.
func (r *ListIpRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
