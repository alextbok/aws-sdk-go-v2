// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListEndpointConfigsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only endpoint configurations with a creation time greater
	// than or equal to the specified time (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only endpoint configurations created before the specified
	// time (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of training jobs to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the endpoint configuration name. This filter returns only endpoint
	// configurations whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of the previous ListEndpointConfig request was truncated, the
	// response includes a NextToken. To retrieve the next set of endpoint configurations,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is CreationTime.
	SortBy EndpointConfigSortKey `type:"string" enum:"true"`

	// The sort order for results. The default is Descending.
	SortOrder OrderKey `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListEndpointConfigsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEndpointConfigsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEndpointConfigsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListEndpointConfigsOutput struct {
	_ struct{} `type:"structure"`

	// An array of endpoint configurations.
	//
	// EndpointConfigs is a required field
	EndpointConfigs []EndpointConfigSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of endpoint configurations, use it in the subsequent request
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListEndpointConfigsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEndpointConfigs = "ListEndpointConfigs"

// ListEndpointConfigsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists endpoint configurations.
//
//    // Example sending a request using ListEndpointConfigsRequest.
//    req := client.ListEndpointConfigsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListEndpointConfigs
func (c *Client) ListEndpointConfigsRequest(input *ListEndpointConfigsInput) ListEndpointConfigsRequest {
	op := &aws.Operation{
		Name:       opListEndpointConfigs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListEndpointConfigsInput{}
	}

	req := c.newRequest(op, input, &ListEndpointConfigsOutput{})

	return ListEndpointConfigsRequest{Request: req, Input: input, Copy: c.ListEndpointConfigsRequest}
}

// ListEndpointConfigsRequest is the request type for the
// ListEndpointConfigs API operation.
type ListEndpointConfigsRequest struct {
	*aws.Request
	Input *ListEndpointConfigsInput
	Copy  func(*ListEndpointConfigsInput) ListEndpointConfigsRequest
}

// Send marshals and sends the ListEndpointConfigs API request.
func (r ListEndpointConfigsRequest) Send(ctx context.Context) (*ListEndpointConfigsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEndpointConfigsResponse{
		ListEndpointConfigsOutput: r.Request.Data.(*ListEndpointConfigsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEndpointConfigsRequestPaginator returns a paginator for ListEndpointConfigs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEndpointConfigsRequest(input)
//   p := sagemaker.NewListEndpointConfigsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEndpointConfigsPaginator(req ListEndpointConfigsRequest) ListEndpointConfigsPaginator {
	return ListEndpointConfigsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEndpointConfigsInput
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

// ListEndpointConfigsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEndpointConfigsPaginator struct {
	aws.Pager
}

func (p *ListEndpointConfigsPaginator) CurrentPage() *ListEndpointConfigsOutput {
	return p.Pager.CurrentPage().(*ListEndpointConfigsOutput)
}

// ListEndpointConfigsResponse is the response type for the
// ListEndpointConfigs API operation.
type ListEndpointConfigsResponse struct {
	*ListEndpointConfigsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEndpointConfigs request.
func (r *ListEndpointConfigsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
