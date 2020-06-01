// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRateBasedRuleInput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateRateBasedRule request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description for the metrics for this RateBasedRule. The
	// name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum
	// length 128 and minimum length one. It can't contain whitespace or metric
	// names reserved for AWS WAF, including "All" and "Default_Action." You can't
	// change the name of the metric after you create the RateBasedRule.
	//
	// MetricName is a required field
	MetricName *string `min:"1" type:"string" required:"true"`

	// A friendly name or description of the RateBasedRule. You can't change the
	// name of a RateBasedRule after you create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The field that AWS WAF uses to determine if requests are likely arriving
	// from a single source and thus subject to rate monitoring. The only valid
	// value for RateKey is IP. IP indicates that requests that arrive from the
	// same IP address are subject to the RateLimit that is specified in the RateBasedRule.
	//
	// RateKey is a required field
	RateKey RateKey `type:"string" required:"true" enum:"true"`

	// The maximum number of requests, which have an identical value in the field
	// that is specified by RateKey, allowed in a five-minute period. If the number
	// of requests exceeds the RateLimit and the other predicates specified in the
	// rule are also met, AWS WAF triggers the action that is specified for this
	// rule.
	//
	// RateLimit is a required field
	RateLimit *int64 `min:"100" type:"long" required:"true"`

	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateRateBasedRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRateBasedRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRateBasedRuleInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.RateKey) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RateKey"))
	}

	if s.RateLimit == nil {
		invalidParams.Add(aws.NewErrParamRequired("RateLimit"))
	}
	if s.RateLimit != nil && *s.RateLimit < 100 {
		invalidParams.Add(aws.NewErrParamMinValue("RateLimit", 100))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRateBasedRuleOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateRateBasedRule request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// The RateBasedRule that is returned in the CreateRateBasedRule response.
	Rule *RateBasedRule `type:"structure"`
}

// String returns the string representation
func (s CreateRateBasedRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRateBasedRule = "CreateRateBasedRule"

// CreateRateBasedRuleRequest returns a request value for making API operation for
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
// Creates a RateBasedRule. The RateBasedRule contains a RateLimit, which specifies
// the maximum number of requests that AWS WAF allows from a specified IP address
// in a five-minute period. The RateBasedRule also contains the IPSet objects,
// ByteMatchSet objects, and other predicates that identify the requests that
// you want to count or block if these requests exceed the RateLimit.
//
// If you add more than one predicate to a RateBasedRule, a request not only
// must exceed the RateLimit, but it also must match all the conditions to be
// counted or blocked. For example, suppose you add the following to a RateBasedRule:
//
//    * An IPSet that matches the IP address 192.0.2.44/32
//
//    * A ByteMatchSet that matches BadBot in the User-Agent header
//
// Further, you specify a RateLimit of 1,000.
//
// You then add the RateBasedRule to a WebACL and specify that you want to block
// requests that meet the conditions in the rule. For a request to be blocked,
// it must come from the IP address 192.0.2.44 and the User-Agent header in
// the request must contain the value BadBot. Further, requests that match these
// two conditions must be received at a rate of more than 1,000 requests every
// five minutes. If both conditions are met and the rate is exceeded, AWS WAF
// blocks the requests. If the rate drops below 1,000 for a five-minute period,
// AWS WAF no longer blocks the requests.
//
// As a second example, suppose you want to limit requests to a particular page
// on your site. To do this, you could add the following to a RateBasedRule:
//
//    * A ByteMatchSet with FieldToMatch of URI
//
//    * A PositionalConstraint of STARTS_WITH
//
//    * A TargetString of login
//
// Further, you specify a RateLimit of 1,000.
//
// By adding this RateBasedRule to a WebACL, you could limit requests to your
// login page without affecting the rest of your site.
//
// To create and configure a RateBasedRule, perform the following steps:
//
// Create and update the predicates that you want to include in the rule. For
// more information, see CreateByteMatchSet, CreateIPSet, and CreateSqlInjectionMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateRule request.
//
// Submit a CreateRateBasedRule request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRule request.
//
// Submit an UpdateRateBasedRule request to specify the predicates that you
// want to include in the rule.
//
// Create and update a WebACL that contains the RateBasedRule. For more information,
// see CreateWebACL.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateRateBasedRuleRequest.
//    req := client.CreateRateBasedRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateRateBasedRule
func (c *Client) CreateRateBasedRuleRequest(input *CreateRateBasedRuleInput) CreateRateBasedRuleRequest {
	op := &aws.Operation{
		Name:       opCreateRateBasedRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRateBasedRuleInput{}
	}

	req := c.newRequest(op, input, &CreateRateBasedRuleOutput{})

	return CreateRateBasedRuleRequest{Request: req, Input: input, Copy: c.CreateRateBasedRuleRequest}
}

// CreateRateBasedRuleRequest is the request type for the
// CreateRateBasedRule API operation.
type CreateRateBasedRuleRequest struct {
	*aws.Request
	Input *CreateRateBasedRuleInput
	Copy  func(*CreateRateBasedRuleInput) CreateRateBasedRuleRequest
}

// Send marshals and sends the CreateRateBasedRule API request.
func (r CreateRateBasedRuleRequest) Send(ctx context.Context) (*CreateRateBasedRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRateBasedRuleResponse{
		CreateRateBasedRuleOutput: r.Request.Data.(*CreateRateBasedRuleOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRateBasedRuleResponse is the response type for the
// CreateRateBasedRule API operation.
type CreateRateBasedRuleResponse struct {
	*CreateRateBasedRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRateBasedRule request.
func (r *CreateRateBasedRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
