// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutPermissionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The policy to attach to the specified rule group.
	//
	// The policy specifications must conform to the following:
	//
	//    * The policy must be composed using IAM Policy version 2012-10-17 or version
	//    2015-01-01.
	//
	//    * The policy must include specifications for Effect, Action, and Principal.
	//
	//    * Effect must specify Allow.
	//
	//    * Action must specify wafv2:CreateWebACL, wafv2:UpdateWebACL, and wafv2:PutFirewallManagerRuleGroups.
	//    AWS WAF rejects any extra actions or wildcard actions in the policy.
	//
	//    * The policy must not include a Resource parameter.
	//
	// For more information, see IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html).
	//
	// Policy is a required field
	Policy *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the RuleGroup to which you want to attach
	// the policy.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s PutPermissionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPermissionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutPermissionPolicyInput"}

	if s.Policy == nil {
		invalidParams.Add(aws.NewErrParamRequired("Policy"))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutPermissionPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutPermissionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutPermissionPolicy = "PutPermissionPolicy"

// PutPermissionPolicyRequest returns a request value for making API operation for
// AWS WAFV2.
//
// Attaches an IAM policy to the specified resource. Use this to share a rule
// group across accounts.
//
// You must be the owner of the rule group to perform this operation.
//
// This action is subject to the following restrictions:
//
//    * You can attach only one policy with each PutPermissionPolicy request.
//
//    * The ARN in the request must be a valid WAF RuleGroup ARN and the rule
//    group must exist in the same region.
//
//    * The user making the request must be the owner of the rule group.
//
//    // Example sending a request using PutPermissionPolicyRequest.
//    req := client.PutPermissionPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/PutPermissionPolicy
func (c *Client) PutPermissionPolicyRequest(input *PutPermissionPolicyInput) PutPermissionPolicyRequest {
	op := &aws.Operation{
		Name:       opPutPermissionPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutPermissionPolicyInput{}
	}

	req := c.newRequest(op, input, &PutPermissionPolicyOutput{})
	return PutPermissionPolicyRequest{Request: req, Input: input, Copy: c.PutPermissionPolicyRequest}
}

// PutPermissionPolicyRequest is the request type for the
// PutPermissionPolicy API operation.
type PutPermissionPolicyRequest struct {
	*aws.Request
	Input *PutPermissionPolicyInput
	Copy  func(*PutPermissionPolicyInput) PutPermissionPolicyRequest
}

// Send marshals and sends the PutPermissionPolicy API request.
func (r PutPermissionPolicyRequest) Send(ctx context.Context) (*PutPermissionPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPermissionPolicyResponse{
		PutPermissionPolicyOutput: r.Request.Data.(*PutPermissionPolicyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPermissionPolicyResponse is the response type for the
// PutPermissionPolicy API operation.
type PutPermissionPolicyResponse struct {
	*PutPermissionPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPermissionPolicy request.
func (r *PutPermissionPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
