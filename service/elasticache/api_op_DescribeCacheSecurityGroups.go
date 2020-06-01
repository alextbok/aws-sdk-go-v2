// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeCacheSecurityGroups operation.
type DescribeCacheSecurityGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache security group to return details for.
	CacheSecurityGroupName *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCacheSecurityGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeCacheSecurityGroups operation.
type DescribeCacheSecurityGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of cache security groups. Each element in the list contains detailed
	// information about one group.
	CacheSecurityGroups []CacheSecurityGroup `locationNameList:"CacheSecurityGroup" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCacheSecurityGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCacheSecurityGroups = "DescribeCacheSecurityGroups"

// DescribeCacheSecurityGroupsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns a list of cache security group descriptions. If a cache security
// group name is specified, the list contains only the description of that group.
// This applicable only when you have ElastiCache in Classic setup
//
//    // Example sending a request using DescribeCacheSecurityGroupsRequest.
//    req := client.DescribeCacheSecurityGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeCacheSecurityGroups
func (c *Client) DescribeCacheSecurityGroupsRequest(input *DescribeCacheSecurityGroupsInput) DescribeCacheSecurityGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeCacheSecurityGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeCacheSecurityGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeCacheSecurityGroupsOutput{})

	return DescribeCacheSecurityGroupsRequest{Request: req, Input: input, Copy: c.DescribeCacheSecurityGroupsRequest}
}

// DescribeCacheSecurityGroupsRequest is the request type for the
// DescribeCacheSecurityGroups API operation.
type DescribeCacheSecurityGroupsRequest struct {
	*aws.Request
	Input *DescribeCacheSecurityGroupsInput
	Copy  func(*DescribeCacheSecurityGroupsInput) DescribeCacheSecurityGroupsRequest
}

// Send marshals and sends the DescribeCacheSecurityGroups API request.
func (r DescribeCacheSecurityGroupsRequest) Send(ctx context.Context) (*DescribeCacheSecurityGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCacheSecurityGroupsResponse{
		DescribeCacheSecurityGroupsOutput: r.Request.Data.(*DescribeCacheSecurityGroupsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCacheSecurityGroupsRequestPaginator returns a paginator for DescribeCacheSecurityGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCacheSecurityGroupsRequest(input)
//   p := elasticache.NewDescribeCacheSecurityGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCacheSecurityGroupsPaginator(req DescribeCacheSecurityGroupsRequest) DescribeCacheSecurityGroupsPaginator {
	return DescribeCacheSecurityGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCacheSecurityGroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeCacheSecurityGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCacheSecurityGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeCacheSecurityGroupsPaginator) CurrentPage() *DescribeCacheSecurityGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeCacheSecurityGroupsOutput)
}

// DescribeCacheSecurityGroupsResponse is the response type for the
// DescribeCacheSecurityGroups API operation.
type DescribeCacheSecurityGroupsResponse struct {
	*DescribeCacheSecurityGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCacheSecurityGroups request.
func (r *DescribeCacheSecurityGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
