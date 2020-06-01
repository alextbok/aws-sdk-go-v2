// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRuleInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RuleId of the Rule that you want to update. RuleId is returned by CreateRule
	// and by ListRules.
	//
	// RuleId is a required field
	RuleId *string `min:"1" type:"string" required:"true"`

	// An array of RuleUpdate objects that you want to insert into or delete from
	// a Rule. For more information, see the applicable data types:
	//
	//    * RuleUpdate: Contains Action and Predicate
	//
	//    * Predicate: Contains DataId, Negated, and Type
	//
	//    * FieldToMatch: Contains Data and Type
	//
	// Updates is a required field
	Updates []RuleUpdate `type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRuleInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleId"))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRuleOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateRule request. You can also
	// use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRule = "UpdateRule"

// UpdateRuleRequest returns a request value for making API operation for
// AWS WAF Regional.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Inserts or deletes Predicate objects in a Rule. Each Predicate object identifies
// a predicate, such as a ByteMatchSet or an IPSet, that specifies the web requests
// that you want to allow, block, or count. If you add more than one predicate
// to a Rule, a request must match all of the specifications to be allowed,
// blocked, or counted. For example, suppose that you add the following to a
// Rule:
//
//    * A ByteMatchSet that matches the value BadBot in the User-Agent header
//
//    * An IPSet that matches the IP address 192.0.2.44
//
// You then add the Rule to a WebACL and specify that you want to block requests
// that satisfy the Rule. For a request to be blocked, the User-Agent header
// in the request must contain the value BadBot and the request must originate
// from the IP address 192.0.2.44.
//
// To create and configure a Rule, perform the following steps:
//
// Create and update the predicates that you want to include in the Rule.
//
// Create the Rule. See CreateRule.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRule request.
//
// Submit an UpdateRule request to add predicates to the Rule.
//
// Create and update a WebACL that contains the Rule. See CreateWebACL.
//
// If you want to replace one ByteMatchSet or IPSet with another, you delete
// the existing one and add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateRuleRequest.
//    req := client.UpdateRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/UpdateRule
func (c *Client) UpdateRuleRequest(input *UpdateRuleInput) UpdateRuleRequest {
	op := &aws.Operation{
		Name:       opUpdateRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRuleInput{}
	}

	req := c.newRequest(op, input, &UpdateRuleOutput{})

	return UpdateRuleRequest{Request: req, Input: input, Copy: c.UpdateRuleRequest}
}

// UpdateRuleRequest is the request type for the
// UpdateRule API operation.
type UpdateRuleRequest struct {
	*aws.Request
	Input *UpdateRuleInput
	Copy  func(*UpdateRuleInput) UpdateRuleRequest
}

// Send marshals and sends the UpdateRule API request.
func (r UpdateRuleRequest) Send(ctx context.Context) (*UpdateRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRuleResponse{
		UpdateRuleOutput: r.Request.Data.(*UpdateRuleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRuleResponse is the response type for the
// UpdateRule API operation.
type UpdateRuleResponse struct {
	*UpdateRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRule request.
func (r *UpdateRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
