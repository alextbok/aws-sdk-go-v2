// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStacksInput struct {
	_ struct{} `type:"structure"`

	// The names of the stacks to describe.
	Names []string `type:"list"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeStacksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStacksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStacksInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStacksOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string `min:"1" type:"string"`

	// Information about the stacks.
	Stacks []Stack `type:"list"`
}

// String returns the string representation
func (s DescribeStacksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStacks = "DescribeStacks"

// DescribeStacksRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Retrieves a list that describes one or more specified stacks, if the stack
// names are provided. Otherwise, all stacks in the account are described.
//
//    // Example sending a request using DescribeStacksRequest.
//    req := client.DescribeStacksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeStacks
func (c *Client) DescribeStacksRequest(input *DescribeStacksInput) DescribeStacksRequest {
	op := &aws.Operation{
		Name:       opDescribeStacks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStacksInput{}
	}

	req := c.newRequest(op, input, &DescribeStacksOutput{})

	return DescribeStacksRequest{Request: req, Input: input, Copy: c.DescribeStacksRequest}
}

// DescribeStacksRequest is the request type for the
// DescribeStacks API operation.
type DescribeStacksRequest struct {
	*aws.Request
	Input *DescribeStacksInput
	Copy  func(*DescribeStacksInput) DescribeStacksRequest
}

// Send marshals and sends the DescribeStacks API request.
func (r DescribeStacksRequest) Send(ctx context.Context) (*DescribeStacksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStacksResponse{
		DescribeStacksOutput: r.Request.Data.(*DescribeStacksOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStacksResponse is the response type for the
// DescribeStacks API operation.
type DescribeStacksResponse struct {
	*DescribeStacksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStacks request.
func (r *DescribeStacksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
