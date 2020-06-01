// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AttachLoadBalancerTargetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Names (ARN) of the target groups. You can specify up
	// to 10 target groups.
	//
	// TargetGroupARNs is a required field
	TargetGroupARNs []string `type:"list" required:"true"`
}

// String returns the string representation
func (s AttachLoadBalancerTargetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachLoadBalancerTargetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachLoadBalancerTargetGroupsInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.TargetGroupARNs == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetGroupARNs"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachLoadBalancerTargetGroupsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachLoadBalancerTargetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachLoadBalancerTargetGroups = "AttachLoadBalancerTargetGroups"

// AttachLoadBalancerTargetGroupsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Attaches one or more target groups to the specified Auto Scaling group.
//
// To describe the target groups for an Auto Scaling group, call the DescribeLoadBalancerTargetGroups
// API. To detach the target group from the Auto Scaling group, call the DetachLoadBalancerTargetGroups
// API.
//
// With Application Load Balancers and Network Load Balancers, instances are
// registered as targets with a target group. With Classic Load Balancers, instances
// are registered with the load balancer. For more information, see Attaching
// a Load Balancer to Your Auto Scaling Group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using AttachLoadBalancerTargetGroupsRequest.
//    req := client.AttachLoadBalancerTargetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AttachLoadBalancerTargetGroups
func (c *Client) AttachLoadBalancerTargetGroupsRequest(input *AttachLoadBalancerTargetGroupsInput) AttachLoadBalancerTargetGroupsRequest {
	op := &aws.Operation{
		Name:       opAttachLoadBalancerTargetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachLoadBalancerTargetGroupsInput{}
	}

	req := c.newRequest(op, input, &AttachLoadBalancerTargetGroupsOutput{})

	return AttachLoadBalancerTargetGroupsRequest{Request: req, Input: input, Copy: c.AttachLoadBalancerTargetGroupsRequest}
}

// AttachLoadBalancerTargetGroupsRequest is the request type for the
// AttachLoadBalancerTargetGroups API operation.
type AttachLoadBalancerTargetGroupsRequest struct {
	*aws.Request
	Input *AttachLoadBalancerTargetGroupsInput
	Copy  func(*AttachLoadBalancerTargetGroupsInput) AttachLoadBalancerTargetGroupsRequest
}

// Send marshals and sends the AttachLoadBalancerTargetGroups API request.
func (r AttachLoadBalancerTargetGroupsRequest) Send(ctx context.Context) (*AttachLoadBalancerTargetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachLoadBalancerTargetGroupsResponse{
		AttachLoadBalancerTargetGroupsOutput: r.Request.Data.(*AttachLoadBalancerTargetGroupsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachLoadBalancerTargetGroupsResponse is the response type for the
// AttachLoadBalancerTargetGroups API operation.
type AttachLoadBalancerTargetGroupsResponse struct {
	*AttachLoadBalancerTargetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachLoadBalancerTargetGroups request.
func (r *AttachLoadBalancerTargetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
