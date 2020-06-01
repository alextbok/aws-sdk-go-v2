// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEffectiveInstanceAssociationsInput struct {
	_ struct{} `type:"structure"`

	// The instance ID for which you want to view all associations.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEffectiveInstanceAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEffectiveInstanceAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEffectiveInstanceAssociationsInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEffectiveInstanceAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// The associations for the requested instance.
	Associations []InstanceAssociation `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEffectiveInstanceAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEffectiveInstanceAssociations = "DescribeEffectiveInstanceAssociations"

// DescribeEffectiveInstanceAssociationsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// All associations for the instance(s).
//
//    // Example sending a request using DescribeEffectiveInstanceAssociationsRequest.
//    req := client.DescribeEffectiveInstanceAssociationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeEffectiveInstanceAssociations
func (c *Client) DescribeEffectiveInstanceAssociationsRequest(input *DescribeEffectiveInstanceAssociationsInput) DescribeEffectiveInstanceAssociationsRequest {
	op := &aws.Operation{
		Name:       opDescribeEffectiveInstanceAssociations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEffectiveInstanceAssociationsInput{}
	}

	req := c.newRequest(op, input, &DescribeEffectiveInstanceAssociationsOutput{})

	return DescribeEffectiveInstanceAssociationsRequest{Request: req, Input: input, Copy: c.DescribeEffectiveInstanceAssociationsRequest}
}

// DescribeEffectiveInstanceAssociationsRequest is the request type for the
// DescribeEffectiveInstanceAssociations API operation.
type DescribeEffectiveInstanceAssociationsRequest struct {
	*aws.Request
	Input *DescribeEffectiveInstanceAssociationsInput
	Copy  func(*DescribeEffectiveInstanceAssociationsInput) DescribeEffectiveInstanceAssociationsRequest
}

// Send marshals and sends the DescribeEffectiveInstanceAssociations API request.
func (r DescribeEffectiveInstanceAssociationsRequest) Send(ctx context.Context) (*DescribeEffectiveInstanceAssociationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEffectiveInstanceAssociationsResponse{
		DescribeEffectiveInstanceAssociationsOutput: r.Request.Data.(*DescribeEffectiveInstanceAssociationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEffectiveInstanceAssociationsResponse is the response type for the
// DescribeEffectiveInstanceAssociations API operation.
type DescribeEffectiveInstanceAssociationsResponse struct {
	*DescribeEffectiveInstanceAssociationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEffectiveInstanceAssociations request.
func (r *DescribeEffectiveInstanceAssociationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
