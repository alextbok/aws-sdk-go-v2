// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computeoptimizer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRecommendationSummariesInput struct {
	_ struct{} `type:"structure"`

	// The AWS account IDs for which to return recommendation summaries.
	//
	// Only one account ID can be specified per request.
	AccountIds []string `locationName:"accountIds" type:"list"`

	// The maximum number of recommendation summaries to return with a single call.
	//
	// To retrieve the remaining results, make another call with the returned NextToken
	// value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token to advance to the next page of recommendation summaries.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetRecommendationSummariesInput) String() string {
	return awsutil.Prettify(s)
}

type GetRecommendationSummariesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to advance to the next page of recommendation summaries.
	//
	// This value is null when there are no more pages of recommendation summaries
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// An array of objects that summarize a recommendation.
	RecommendationSummaries []RecommendationSummary `locationName:"recommendationSummaries" type:"list"`
}

// String returns the string representation
func (s GetRecommendationSummariesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRecommendationSummaries = "GetRecommendationSummaries"

// GetRecommendationSummariesRequest returns a request value for making API operation for
// AWS Compute Optimizer.
//
// Returns the optimization findings for an account.
//
// For example, it returns the number of Amazon EC2 instances in an account
// that are under-provisioned, over-provisioned, or optimized. It also returns
// the number of Auto Scaling groups in an account that are not optimized, or
// optimized.
//
//    // Example sending a request using GetRecommendationSummariesRequest.
//    req := client.GetRecommendationSummariesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/compute-optimizer-2019-11-01/GetRecommendationSummaries
func (c *Client) GetRecommendationSummariesRequest(input *GetRecommendationSummariesInput) GetRecommendationSummariesRequest {
	op := &aws.Operation{
		Name:       opGetRecommendationSummaries,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRecommendationSummariesInput{}
	}

	req := c.newRequest(op, input, &GetRecommendationSummariesOutput{})

	return GetRecommendationSummariesRequest{Request: req, Input: input, Copy: c.GetRecommendationSummariesRequest}
}

// GetRecommendationSummariesRequest is the request type for the
// GetRecommendationSummaries API operation.
type GetRecommendationSummariesRequest struct {
	*aws.Request
	Input *GetRecommendationSummariesInput
	Copy  func(*GetRecommendationSummariesInput) GetRecommendationSummariesRequest
}

// Send marshals and sends the GetRecommendationSummaries API request.
func (r GetRecommendationSummariesRequest) Send(ctx context.Context) (*GetRecommendationSummariesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRecommendationSummariesResponse{
		GetRecommendationSummariesOutput: r.Request.Data.(*GetRecommendationSummariesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRecommendationSummariesResponse is the response type for the
// GetRecommendationSummaries API operation.
type GetRecommendationSummariesResponse struct {
	*GetRecommendationSummariesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRecommendationSummaries request.
func (r *GetRecommendationSummariesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
