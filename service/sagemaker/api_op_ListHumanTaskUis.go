// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListHumanTaskUisInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only human task user interfaces with a creation time
	// greater than or equal to the specified timestamp.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only human task user interfaces that were created before
	// the specified timestamp.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// The total number of items to return. If the total number of available items
	// is more than the value specified in MaxResults, then a NextToken will be
	// provided in the output that you can use to resume pagination.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to resume pagination.
	NextToken *string `type:"string"`

	// An optional value that specifies whether you want the results sorted in Ascending
	// or Descending order.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListHumanTaskUisInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListHumanTaskUisInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListHumanTaskUisInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListHumanTaskUisOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects describing the human task user interfaces.
	//
	// HumanTaskUiSummaries is a required field
	HumanTaskUiSummaries []HumanTaskUiSummary `type:"list" required:"true"`

	// A token to resume pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHumanTaskUisOutput) String() string {
	return awsutil.Prettify(s)
}

const opListHumanTaskUis = "ListHumanTaskUis"

// ListHumanTaskUisRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns information about the human task user interfaces in your account.
//
//    // Example sending a request using ListHumanTaskUisRequest.
//    req := client.ListHumanTaskUisRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListHumanTaskUis
func (c *Client) ListHumanTaskUisRequest(input *ListHumanTaskUisInput) ListHumanTaskUisRequest {
	op := &aws.Operation{
		Name:       opListHumanTaskUis,
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
		input = &ListHumanTaskUisInput{}
	}

	req := c.newRequest(op, input, &ListHumanTaskUisOutput{})

	return ListHumanTaskUisRequest{Request: req, Input: input, Copy: c.ListHumanTaskUisRequest}
}

// ListHumanTaskUisRequest is the request type for the
// ListHumanTaskUis API operation.
type ListHumanTaskUisRequest struct {
	*aws.Request
	Input *ListHumanTaskUisInput
	Copy  func(*ListHumanTaskUisInput) ListHumanTaskUisRequest
}

// Send marshals and sends the ListHumanTaskUis API request.
func (r ListHumanTaskUisRequest) Send(ctx context.Context) (*ListHumanTaskUisResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHumanTaskUisResponse{
		ListHumanTaskUisOutput: r.Request.Data.(*ListHumanTaskUisOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHumanTaskUisRequestPaginator returns a paginator for ListHumanTaskUis.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHumanTaskUisRequest(input)
//   p := sagemaker.NewListHumanTaskUisRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHumanTaskUisPaginator(req ListHumanTaskUisRequest) ListHumanTaskUisPaginator {
	return ListHumanTaskUisPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListHumanTaskUisInput
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

// ListHumanTaskUisPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHumanTaskUisPaginator struct {
	aws.Pager
}

func (p *ListHumanTaskUisPaginator) CurrentPage() *ListHumanTaskUisOutput {
	return p.Pager.CurrentPage().(*ListHumanTaskUisOutput)
}

// ListHumanTaskUisResponse is the response type for the
// ListHumanTaskUis API operation.
type ListHumanTaskUisResponse struct {
	*ListHumanTaskUisOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHumanTaskUis request.
func (r *ListHumanTaskUisResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
