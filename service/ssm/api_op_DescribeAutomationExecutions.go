// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAutomationExecutionsInput struct {
	_ struct{} `type:"structure"`

	// Filters used to limit the scope of executions that are requested.
	Filters []AutomationExecutionFilter `min:"1" type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutomationExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAutomationExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAutomationExecutionsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAutomationExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// The list of details about each automation execution which has occurred which
	// matches the filter specification, if any.
	AutomationExecutionMetadataList []AutomationExecutionMetadata `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutomationExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAutomationExecutions = "DescribeAutomationExecutions"

// DescribeAutomationExecutionsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Provides details about all active and terminated Automation executions.
//
//    // Example sending a request using DescribeAutomationExecutionsRequest.
//    req := client.DescribeAutomationExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeAutomationExecutions
func (c *Client) DescribeAutomationExecutionsRequest(input *DescribeAutomationExecutionsInput) DescribeAutomationExecutionsRequest {
	op := &aws.Operation{
		Name:       opDescribeAutomationExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAutomationExecutionsInput{}
	}

	req := c.newRequest(op, input, &DescribeAutomationExecutionsOutput{})

	return DescribeAutomationExecutionsRequest{Request: req, Input: input, Copy: c.DescribeAutomationExecutionsRequest}
}

// DescribeAutomationExecutionsRequest is the request type for the
// DescribeAutomationExecutions API operation.
type DescribeAutomationExecutionsRequest struct {
	*aws.Request
	Input *DescribeAutomationExecutionsInput
	Copy  func(*DescribeAutomationExecutionsInput) DescribeAutomationExecutionsRequest
}

// Send marshals and sends the DescribeAutomationExecutions API request.
func (r DescribeAutomationExecutionsRequest) Send(ctx context.Context) (*DescribeAutomationExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAutomationExecutionsResponse{
		DescribeAutomationExecutionsOutput: r.Request.Data.(*DescribeAutomationExecutionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAutomationExecutionsResponse is the response type for the
// DescribeAutomationExecutions API operation.
type DescribeAutomationExecutionsResponse struct {
	*DescribeAutomationExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAutomationExecutions request.
func (r *DescribeAutomationExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
