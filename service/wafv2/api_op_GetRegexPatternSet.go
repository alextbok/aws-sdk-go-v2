// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRegexPatternSetInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the set. This ID is returned in the responses to
	// create and list commands. You provide it to operations like update and delete.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The name of the set. You cannot change the name after you create the set.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

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
}

// String returns the string representation
func (s GetRegexPatternSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRegexPatternSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRegexPatternSetInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
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

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRegexPatternSetOutput struct {
	_ struct{} `type:"structure"`

	// A token used for optimistic locking. AWS WAF returns a token to your get
	// and list requests, to mark the state of the entity at the time of the request.
	// To make changes to the entity associated with the token, you provide the
	// token to operations like update and delete. AWS WAF uses the token to ensure
	// that no changes have been made to the entity since you last retrieved it.
	// If a change has been made, the update fails with a WAFOptimisticLockException.
	// If this happens, perform another get, and use the new token returned by that
	// operation.
	LockToken *string `min:"1" type:"string"`

	//
	// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
	// 2019. For information, including how to migrate your AWS WAF resources from
	// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
	//
	// Contains one or more regular expressions.
	//
	// AWS WAF assigns an ARN to each RegexPatternSet that you create. To use a
	// set in a rule, you provide the ARN to the Rule statement RegexPatternSetReferenceStatement.
	RegexPatternSet *RegexPatternSet `type:"structure"`
}

// String returns the string representation
func (s GetRegexPatternSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRegexPatternSet = "GetRegexPatternSet"

// GetRegexPatternSetRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Retrieves the specified RegexPatternSet.
//
//    // Example sending a request using GetRegexPatternSetRequest.
//    req := client.GetRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/GetRegexPatternSet
func (c *Client) GetRegexPatternSetRequest(input *GetRegexPatternSetInput) GetRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opGetRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &GetRegexPatternSetOutput{})

	return GetRegexPatternSetRequest{Request: req, Input: input, Copy: c.GetRegexPatternSetRequest}
}

// GetRegexPatternSetRequest is the request type for the
// GetRegexPatternSet API operation.
type GetRegexPatternSetRequest struct {
	*aws.Request
	Input *GetRegexPatternSetInput
	Copy  func(*GetRegexPatternSetInput) GetRegexPatternSetRequest
}

// Send marshals and sends the GetRegexPatternSet API request.
func (r GetRegexPatternSetRequest) Send(ctx context.Context) (*GetRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRegexPatternSetResponse{
		GetRegexPatternSetOutput: r.Request.Data.(*GetRegexPatternSetOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRegexPatternSetResponse is the response type for the
// GetRegexPatternSet API operation.
type GetRegexPatternSetResponse struct {
	*GetRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRegexPatternSet request.
func (r *GetRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
