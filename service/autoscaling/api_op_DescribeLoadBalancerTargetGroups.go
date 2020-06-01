// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLoadBalancerTargetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The maximum number of items to return with this call. The default value is
	// 100 and the maximum value is 100.
	MaxRecords *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeLoadBalancerTargetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBalancerTargetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLoadBalancerTargetGroupsInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLoadBalancerTargetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the target groups.
	LoadBalancerTargetGroups []LoadBalancerTargetGroupState `type:"list"`

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this
	// string for the NextToken value when requesting the next set of items. This
	// value is null when there are no more items to return.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeLoadBalancerTargetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLoadBalancerTargetGroups = "DescribeLoadBalancerTargetGroups"

// DescribeLoadBalancerTargetGroupsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes the target groups for the specified Auto Scaling group.
//
//    // Example sending a request using DescribeLoadBalancerTargetGroupsRequest.
//    req := client.DescribeLoadBalancerTargetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeLoadBalancerTargetGroups
func (c *Client) DescribeLoadBalancerTargetGroupsRequest(input *DescribeLoadBalancerTargetGroupsInput) DescribeLoadBalancerTargetGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeLoadBalancerTargetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBalancerTargetGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeLoadBalancerTargetGroupsOutput{})

	return DescribeLoadBalancerTargetGroupsRequest{Request: req, Input: input, Copy: c.DescribeLoadBalancerTargetGroupsRequest}
}

// DescribeLoadBalancerTargetGroupsRequest is the request type for the
// DescribeLoadBalancerTargetGroups API operation.
type DescribeLoadBalancerTargetGroupsRequest struct {
	*aws.Request
	Input *DescribeLoadBalancerTargetGroupsInput
	Copy  func(*DescribeLoadBalancerTargetGroupsInput) DescribeLoadBalancerTargetGroupsRequest
}

// Send marshals and sends the DescribeLoadBalancerTargetGroups API request.
func (r DescribeLoadBalancerTargetGroupsRequest) Send(ctx context.Context) (*DescribeLoadBalancerTargetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoadBalancerTargetGroupsResponse{
		DescribeLoadBalancerTargetGroupsOutput: r.Request.Data.(*DescribeLoadBalancerTargetGroupsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLoadBalancerTargetGroupsResponse is the response type for the
// DescribeLoadBalancerTargetGroups API operation.
type DescribeLoadBalancerTargetGroupsResponse struct {
	*DescribeLoadBalancerTargetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoadBalancerTargetGroups request.
func (r *DescribeLoadBalancerTargetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
