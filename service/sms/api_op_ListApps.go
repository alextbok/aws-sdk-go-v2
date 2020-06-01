// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListAppsInput struct {
	_ struct{} `type:"structure"`

	AppIds []string `locationName:"appIds" type:"list"`

	// The maximum number of results to return in a single call. The default value
	// is 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAppsInput) String() string {
	return awsutil.Prettify(s)
}

type ListAppsOutput struct {
	_ struct{} `type:"structure"`

	// A list of application summaries.
	Apps []AppSummary `locationName:"apps" type:"list"`

	// The token required to retrieve the next set of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAppsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListApps = "ListApps"

// ListAppsRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Returns a list of summaries for all applications.
//
//    // Example sending a request using ListAppsRequest.
//    req := client.ListAppsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/ListApps
func (c *Client) ListAppsRequest(input *ListAppsInput) ListAppsRequest {
	op := &aws.Operation{
		Name:       opListApps,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAppsInput{}
	}

	req := c.newRequest(op, input, &ListAppsOutput{})

	return ListAppsRequest{Request: req, Input: input, Copy: c.ListAppsRequest}
}

// ListAppsRequest is the request type for the
// ListApps API operation.
type ListAppsRequest struct {
	*aws.Request
	Input *ListAppsInput
	Copy  func(*ListAppsInput) ListAppsRequest
}

// Send marshals and sends the ListApps API request.
func (r ListAppsRequest) Send(ctx context.Context) (*ListAppsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAppsResponse{
		ListAppsOutput: r.Request.Data.(*ListAppsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAppsResponse is the response type for the
// ListApps API operation.
type ListAppsResponse struct {
	*ListAppsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApps request.
func (r *ListAppsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
