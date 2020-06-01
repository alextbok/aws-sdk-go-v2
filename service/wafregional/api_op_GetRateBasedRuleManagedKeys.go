// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRateBasedRuleManagedKeysInput struct {
	_ struct{} `type:"structure"`

	// A null value and not currently used. Do not include this in your request.
	NextMarker *string `min:"1" type:"string"`

	// The RuleId of the RateBasedRule for which you want to get a list of ManagedKeys.
	// RuleId is returned by CreateRateBasedRule and by ListRateBasedRules.
	//
	// RuleId is a required field
	RuleId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRateBasedRuleManagedKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRateBasedRuleManagedKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRateBasedRuleManagedKeysInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if s.RuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleId"))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRateBasedRuleManagedKeysOutput struct {
	_ struct{} `type:"structure"`

	// An array of IP addresses that currently are blocked by the specified RateBasedRule.
	ManagedKeys []string `type:"list"`

	// A null value and not currently used.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetRateBasedRuleManagedKeysOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRateBasedRuleManagedKeys = "GetRateBasedRuleManagedKeys"

// GetRateBasedRuleManagedKeysRequest returns a request value for making API operation for
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
// Returns an array of IP addresses currently being blocked by the RateBasedRule
// that is specified by the RuleId. The maximum number of managed keys that
// will be blocked is 10,000. If more than 10,000 addresses exceed the rate
// limit, the 10,000 addresses with the highest rates will be blocked.
//
//    // Example sending a request using GetRateBasedRuleManagedKeysRequest.
//    req := client.GetRateBasedRuleManagedKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRateBasedRuleManagedKeys
func (c *Client) GetRateBasedRuleManagedKeysRequest(input *GetRateBasedRuleManagedKeysInput) GetRateBasedRuleManagedKeysRequest {
	op := &aws.Operation{
		Name:       opGetRateBasedRuleManagedKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRateBasedRuleManagedKeysInput{}
	}

	req := c.newRequest(op, input, &GetRateBasedRuleManagedKeysOutput{})

	return GetRateBasedRuleManagedKeysRequest{Request: req, Input: input, Copy: c.GetRateBasedRuleManagedKeysRequest}
}

// GetRateBasedRuleManagedKeysRequest is the request type for the
// GetRateBasedRuleManagedKeys API operation.
type GetRateBasedRuleManagedKeysRequest struct {
	*aws.Request
	Input *GetRateBasedRuleManagedKeysInput
	Copy  func(*GetRateBasedRuleManagedKeysInput) GetRateBasedRuleManagedKeysRequest
}

// Send marshals and sends the GetRateBasedRuleManagedKeys API request.
func (r GetRateBasedRuleManagedKeysRequest) Send(ctx context.Context) (*GetRateBasedRuleManagedKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRateBasedRuleManagedKeysResponse{
		GetRateBasedRuleManagedKeysOutput: r.Request.Data.(*GetRateBasedRuleManagedKeysOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRateBasedRuleManagedKeysResponse is the response type for the
// GetRateBasedRuleManagedKeys API operation.
type GetRateBasedRuleManagedKeysResponse struct {
	*GetRateBasedRuleManagedKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRateBasedRuleManagedKeys request.
func (r *GetRateBasedRuleManagedKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
