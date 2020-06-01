// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListAvailableManagementCidrRangesInput struct {
	_ struct{} `type:"structure"`

	// The IP address range to search. Specify an IP address range that is compatible
	// with your network and in CIDR notation (that is, specify the range as an
	// IPv4 CIDR block).
	//
	// ManagementCidrRangeConstraint is a required field
	ManagementCidrRangeConstraint *string `type:"string" required:"true"`

	// The maximum number of items to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListAvailableManagementCidrRangesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAvailableManagementCidrRangesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAvailableManagementCidrRangesInput"}

	if s.ManagementCidrRangeConstraint == nil {
		invalidParams.Add(aws.NewErrParamRequired("ManagementCidrRangeConstraint"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAvailableManagementCidrRangesOutput struct {
	_ struct{} `type:"structure"`

	// The list of available IP address ranges, specified as IPv4 CIDR blocks.
	ManagementCidrRanges []string `type:"list"`

	// The token to use to retrieve the next set of results, or null if no more
	// results are available.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListAvailableManagementCidrRangesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAvailableManagementCidrRanges = "ListAvailableManagementCidrRanges"

// ListAvailableManagementCidrRangesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Retrieves a list of IP address ranges, specified as IPv4 CIDR blocks, that
// you can use for the network management interface when you enable Bring Your
// Own License (BYOL).
//
// The management network interface is connected to a secure Amazon WorkSpaces
// management network. It is used for interactive streaming of the WorkSpace
// desktop to Amazon WorkSpaces clients, and to allow Amazon WorkSpaces to manage
// the WorkSpace.
//
//    // Example sending a request using ListAvailableManagementCidrRangesRequest.
//    req := client.ListAvailableManagementCidrRangesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ListAvailableManagementCidrRanges
func (c *Client) ListAvailableManagementCidrRangesRequest(input *ListAvailableManagementCidrRangesInput) ListAvailableManagementCidrRangesRequest {
	op := &aws.Operation{
		Name:       opListAvailableManagementCidrRanges,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAvailableManagementCidrRangesInput{}
	}

	req := c.newRequest(op, input, &ListAvailableManagementCidrRangesOutput{})

	return ListAvailableManagementCidrRangesRequest{Request: req, Input: input, Copy: c.ListAvailableManagementCidrRangesRequest}
}

// ListAvailableManagementCidrRangesRequest is the request type for the
// ListAvailableManagementCidrRanges API operation.
type ListAvailableManagementCidrRangesRequest struct {
	*aws.Request
	Input *ListAvailableManagementCidrRangesInput
	Copy  func(*ListAvailableManagementCidrRangesInput) ListAvailableManagementCidrRangesRequest
}

// Send marshals and sends the ListAvailableManagementCidrRanges API request.
func (r ListAvailableManagementCidrRangesRequest) Send(ctx context.Context) (*ListAvailableManagementCidrRangesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAvailableManagementCidrRangesResponse{
		ListAvailableManagementCidrRangesOutput: r.Request.Data.(*ListAvailableManagementCidrRangesOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAvailableManagementCidrRangesResponse is the response type for the
// ListAvailableManagementCidrRanges API operation.
type ListAvailableManagementCidrRangesResponse struct {
	*ListAvailableManagementCidrRangesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAvailableManagementCidrRanges request.
func (r *ListAvailableManagementCidrRangesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
