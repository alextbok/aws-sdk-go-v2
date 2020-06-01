// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAggregateDiscoveredResourceCountsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration aggregator.
	//
	// ConfigurationAggregatorName is a required field
	ConfigurationAggregatorName *string `min:"1" type:"string" required:"true"`

	// Filters the results based on the ResourceCountFilters object.
	Filters *ResourceCountFilters `type:"structure"`

	// The key to group the resource counts.
	GroupByKey ResourceCountGroupKey `type:"string" enum:"true"`

	// The maximum number of GroupedResourceCount objects returned on each page.
	// The default is 1000. You cannot specify a number greater than 1000. If you
	// specify 0, AWS Config uses the default.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetAggregateDiscoveredResourceCountsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAggregateDiscoveredResourceCountsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAggregateDiscoveredResourceCountsInput"}

	if s.ConfigurationAggregatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationAggregatorName"))
	}
	if s.ConfigurationAggregatorName != nil && len(*s.ConfigurationAggregatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationAggregatorName", 1))
	}
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			invalidParams.AddNested("Filters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAggregateDiscoveredResourceCountsOutput struct {
	_ struct{} `type:"structure"`

	// The key passed into the request object. If GroupByKey is not provided, the
	// result will be empty.
	GroupByKey *string `min:"1" type:"string"`

	// Returns a list of GroupedResourceCount objects.
	GroupedResourceCounts []GroupedResourceCount `type:"list"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// The total number of resources that are present in an aggregator with the
	// filters that you provide.
	//
	// TotalDiscoveredResources is a required field
	TotalDiscoveredResources *int64 `type:"long" required:"true"`
}

// String returns the string representation
func (s GetAggregateDiscoveredResourceCountsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAggregateDiscoveredResourceCounts = "GetAggregateDiscoveredResourceCounts"

// GetAggregateDiscoveredResourceCountsRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the resource counts across accounts and regions that are present
// in your AWS Config aggregator. You can request the resource counts by providing
// filters and GroupByKey.
//
// For example, if the input contains accountID 12345678910 and region us-east-1
// in filters, the API returns the count of resources in account ID 12345678910
// and region us-east-1. If the input contains ACCOUNT_ID as a GroupByKey, the
// API returns resource counts for all source accounts that are present in your
// aggregator.
//
//    // Example sending a request using GetAggregateDiscoveredResourceCountsRequest.
//    req := client.GetAggregateDiscoveredResourceCountsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetAggregateDiscoveredResourceCounts
func (c *Client) GetAggregateDiscoveredResourceCountsRequest(input *GetAggregateDiscoveredResourceCountsInput) GetAggregateDiscoveredResourceCountsRequest {
	op := &aws.Operation{
		Name:       opGetAggregateDiscoveredResourceCounts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAggregateDiscoveredResourceCountsInput{}
	}

	req := c.newRequest(op, input, &GetAggregateDiscoveredResourceCountsOutput{})

	return GetAggregateDiscoveredResourceCountsRequest{Request: req, Input: input, Copy: c.GetAggregateDiscoveredResourceCountsRequest}
}

// GetAggregateDiscoveredResourceCountsRequest is the request type for the
// GetAggregateDiscoveredResourceCounts API operation.
type GetAggregateDiscoveredResourceCountsRequest struct {
	*aws.Request
	Input *GetAggregateDiscoveredResourceCountsInput
	Copy  func(*GetAggregateDiscoveredResourceCountsInput) GetAggregateDiscoveredResourceCountsRequest
}

// Send marshals and sends the GetAggregateDiscoveredResourceCounts API request.
func (r GetAggregateDiscoveredResourceCountsRequest) Send(ctx context.Context) (*GetAggregateDiscoveredResourceCountsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAggregateDiscoveredResourceCountsResponse{
		GetAggregateDiscoveredResourceCountsOutput: r.Request.Data.(*GetAggregateDiscoveredResourceCountsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAggregateDiscoveredResourceCountsResponse is the response type for the
// GetAggregateDiscoveredResourceCounts API operation.
type GetAggregateDiscoveredResourceCountsResponse struct {
	*GetAggregateDiscoveredResourceCountsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAggregateDiscoveredResourceCounts request.
func (r *GetAggregateDiscoveredResourceCountsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
