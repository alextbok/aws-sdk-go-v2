// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateLoadBalancerPolicy.
type CreateLoadBalancerPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`

	// The policy attributes.
	PolicyAttributes []PolicyAttribute `type:"list"`

	// The name of the load balancer policy to be created. This name must be unique
	// within the set of policies for this load balancer.
	//
	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`

	// The name of the base policy type. To get the list of policy types, use DescribeLoadBalancerPolicyTypes.
	//
	// PolicyTypeName is a required field
	PolicyTypeName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLoadBalancerPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLoadBalancerPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLoadBalancerPolicyInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}

	if s.PolicyTypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyTypeName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CreateLoadBalancerPolicy.
type CreateLoadBalancerPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLoadBalancerPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLoadBalancerPolicy = "CreateLoadBalancerPolicy"

// CreateLoadBalancerPolicyRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates a policy with the specified attributes for the specified load balancer.
//
// Policies are settings that are saved for your load balancer and that can
// be applied to the listener or the application server, depending on the policy
// type.
//
//    // Example sending a request using CreateLoadBalancerPolicyRequest.
//    req := client.CreateLoadBalancerPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/CreateLoadBalancerPolicy
func (c *Client) CreateLoadBalancerPolicyRequest(input *CreateLoadBalancerPolicyInput) CreateLoadBalancerPolicyRequest {
	op := &aws.Operation{
		Name:       opCreateLoadBalancerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLoadBalancerPolicyInput{}
	}

	req := c.newRequest(op, input, &CreateLoadBalancerPolicyOutput{})

	return CreateLoadBalancerPolicyRequest{Request: req, Input: input, Copy: c.CreateLoadBalancerPolicyRequest}
}

// CreateLoadBalancerPolicyRequest is the request type for the
// CreateLoadBalancerPolicy API operation.
type CreateLoadBalancerPolicyRequest struct {
	*aws.Request
	Input *CreateLoadBalancerPolicyInput
	Copy  func(*CreateLoadBalancerPolicyInput) CreateLoadBalancerPolicyRequest
}

// Send marshals and sends the CreateLoadBalancerPolicy API request.
func (r CreateLoadBalancerPolicyRequest) Send(ctx context.Context) (*CreateLoadBalancerPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLoadBalancerPolicyResponse{
		CreateLoadBalancerPolicyOutput: r.Request.Data.(*CreateLoadBalancerPolicyOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLoadBalancerPolicyResponse is the response type for the
// CreateLoadBalancerPolicy API operation.
type CreateLoadBalancerPolicyResponse struct {
	*CreateLoadBalancerPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLoadBalancerPolicy request.
func (r *CreateLoadBalancerPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
