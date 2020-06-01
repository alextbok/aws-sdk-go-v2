// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSubscribedRuleGroupsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of subscribed rule groups that you want AWS WAF to return
	// for this request. If you have more objects than the number you specify for
	// Limit, the response includes a NextMarker value that you can use to get another
	// batch of objects.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more ByteMatchSetssubscribed
	// rule groups than the value of Limit, AWS WAF returns a NextMarker value in
	// the response that allows you to list another group of subscribed rule groups.
	// For the second and subsequent ListSubscribedRuleGroupsRequest requests, specify
	// the value of NextMarker from the previous response to get information about
	// another batch of subscribed rule groups.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListSubscribedRuleGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSubscribedRuleGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSubscribedRuleGroupsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSubscribedRuleGroupsOutput struct {
	_ struct{} `type:"structure"`

	// If you have more objects than the number that you specified for Limit in
	// the request, the response includes a NextMarker value. To list more objects,
	// submit another ListSubscribedRuleGroups request, and specify the NextMarker
	// value from the response in the NextMarker value in the next request.
	NextMarker *string `min:"1" type:"string"`

	// An array of RuleGroup objects.
	RuleGroups []SubscribedRuleGroupSummary `type:"list"`
}

// String returns the string representation
func (s ListSubscribedRuleGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSubscribedRuleGroups = "ListSubscribedRuleGroups"

// ListSubscribedRuleGroupsRequest returns a request value for making API operation for
// AWS WAF.
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
// Returns an array of RuleGroup objects that you are subscribed to.
//
//    // Example sending a request using ListSubscribedRuleGroupsRequest.
//    req := client.ListSubscribedRuleGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/ListSubscribedRuleGroups
func (c *Client) ListSubscribedRuleGroupsRequest(input *ListSubscribedRuleGroupsInput) ListSubscribedRuleGroupsRequest {
	op := &aws.Operation{
		Name:       opListSubscribedRuleGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSubscribedRuleGroupsInput{}
	}

	req := c.newRequest(op, input, &ListSubscribedRuleGroupsOutput{})

	return ListSubscribedRuleGroupsRequest{Request: req, Input: input, Copy: c.ListSubscribedRuleGroupsRequest}
}

// ListSubscribedRuleGroupsRequest is the request type for the
// ListSubscribedRuleGroups API operation.
type ListSubscribedRuleGroupsRequest struct {
	*aws.Request
	Input *ListSubscribedRuleGroupsInput
	Copy  func(*ListSubscribedRuleGroupsInput) ListSubscribedRuleGroupsRequest
}

// Send marshals and sends the ListSubscribedRuleGroups API request.
func (r ListSubscribedRuleGroupsRequest) Send(ctx context.Context) (*ListSubscribedRuleGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSubscribedRuleGroupsResponse{
		ListSubscribedRuleGroupsOutput: r.Request.Data.(*ListSubscribedRuleGroupsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSubscribedRuleGroupsResponse is the response type for the
// ListSubscribedRuleGroups API operation.
type ListSubscribedRuleGroupsResponse struct {
	*ListSubscribedRuleGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSubscribedRuleGroups request.
func (r *ListSubscribedRuleGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
