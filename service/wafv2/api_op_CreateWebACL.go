// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateWebACLInput struct {
	_ struct{} `type:"structure"`

	// The action to perform if none of the Rules contained in the WebACL match.
	//
	// DefaultAction is a required field
	DefaultAction *DefaultAction `type:"structure" required:"true"`

	// A description of the Web ACL that helps with identification. You cannot change
	// the description of a Web ACL after you create it.
	Description *string `min:"1" type:"string"`

	// The name of the Web ACL. You cannot change the name of a Web ACL after you
	// create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Rule statements used to identify the web requests that you want to allow,
	// block, or count. Each rule includes one top-level statement that AWS WAF
	// uses to identify matching web requests, and parameters that govern how AWS
	// WAF handles them.
	Rules []Rule `type:"list"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`

	// An array of key:value pairs to associate with the resource.
	Tags []Tag `min:"1" type:"list"`

	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// VisibilityConfig is a required field
	VisibilityConfig *VisibilityConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWebACLInput"}

	if s.DefaultAction == nil {
		invalidParams.Add(aws.NewErrParamRequired("DefaultAction"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if s.VisibilityConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("VisibilityConfig"))
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VisibilityConfig != nil {
		if err := s.VisibilityConfig.Validate(); err != nil {
			invalidParams.AddNested("VisibilityConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateWebACLOutput struct {
	_ struct{} `type:"structure"`

	// High-level information about a WebACL, returned by operations like create
	// and list. This provides information like the ID, that you can use to retrieve
	// and manage a WebACL, and the ARN, that you provide to operations like AssociateWebACL.
	Summary *WebACLSummary `type:"structure"`
}

// String returns the string representation
func (s CreateWebACLOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateWebACL = "CreateWebACL"

// CreateWebACLRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Creates a WebACL per the specifications provided.
//
// A Web ACL defines a collection of rules to use to inspect and control web
// requests. Each rule has an action defined (allow, block, or count) for requests
// that match the statement of the rule. In the Web ACL, you assign a default
// action to take (allow, block) for any request that does not match any of
// the rules. The rules in a Web ACL can be a combination of the types Rule,
// RuleGroup, and managed rule group. You can associate a Web ACL with one or
// more AWS resources to protect. The resources can be Amazon CloudFront, an
// Amazon API Gateway API, or an Application Load Balancer.
//
//    // Example sending a request using CreateWebACLRequest.
//    req := client.CreateWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/CreateWebACL
func (c *Client) CreateWebACLRequest(input *CreateWebACLInput) CreateWebACLRequest {
	op := &aws.Operation{
		Name:       opCreateWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWebACLInput{}
	}

	req := c.newRequest(op, input, &CreateWebACLOutput{})

	return CreateWebACLRequest{Request: req, Input: input, Copy: c.CreateWebACLRequest}
}

// CreateWebACLRequest is the request type for the
// CreateWebACL API operation.
type CreateWebACLRequest struct {
	*aws.Request
	Input *CreateWebACLInput
	Copy  func(*CreateWebACLInput) CreateWebACLRequest
}

// Send marshals and sends the CreateWebACL API request.
func (r CreateWebACLRequest) Send(ctx context.Context) (*CreateWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWebACLResponse{
		CreateWebACLOutput: r.Request.Data.(*CreateWebACLOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWebACLResponse is the response type for the
// CreateWebACL API operation.
type CreateWebACLResponse struct {
	*CreateWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWebACL request.
func (r *CreateWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
