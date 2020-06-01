// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListIPSetsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of IPSet objects that you want AWS WAF to return for
	// this request. If you have more IPSet objects than the number you specify
	// for Limit, the response includes a NextMarker value that you can use to get
	// another batch of IPSet objects.
	Limit *int64 `type:"integer"`

	// AWS WAF returns a NextMarker value in the response that allows you to list
	// another group of IPSets. For the second and subsequent ListIPSets requests,
	// specify the value of NextMarker from the previous response to get information
	// about another batch of IPSets.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListIPSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIPSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIPSetsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListIPSetsOutput struct {
	_ struct{} `type:"structure"`

	// An array of IPSetSummary objects.
	IPSets []IPSetSummary `type:"list"`

	// To list more IPSet objects, submit another ListIPSets request, and in the
	// next request use the NextMarker response value as the NextMarker value.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListIPSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListIPSets = "ListIPSets"

// ListIPSetsRequest returns a request value for making API operation for
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
// Returns an array of IPSetSummary objects in the response.
//
//    // Example sending a request using ListIPSetsRequest.
//    req := client.ListIPSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListIPSets
func (c *Client) ListIPSetsRequest(input *ListIPSetsInput) ListIPSetsRequest {
	op := &aws.Operation{
		Name:       opListIPSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListIPSetsInput{}
	}

	req := c.newRequest(op, input, &ListIPSetsOutput{})

	return ListIPSetsRequest{Request: req, Input: input, Copy: c.ListIPSetsRequest}
}

// ListIPSetsRequest is the request type for the
// ListIPSets API operation.
type ListIPSetsRequest struct {
	*aws.Request
	Input *ListIPSetsInput
	Copy  func(*ListIPSetsInput) ListIPSetsRequest
}

// Send marshals and sends the ListIPSets API request.
func (r ListIPSetsRequest) Send(ctx context.Context) (*ListIPSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIPSetsResponse{
		ListIPSetsOutput: r.Request.Data.(*ListIPSetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIPSetsResponse is the response type for the
// ListIPSets API operation.
type ListIPSetsResponse struct {
	*ListIPSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIPSets request.
func (r *ListIPSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
