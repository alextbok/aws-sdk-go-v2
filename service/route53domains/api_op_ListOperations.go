// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The ListOperations request includes the following elements.
type ListOperationsInput struct {
	_ struct{} `type:"structure"`

	// For an initial request for a list of operations, omit this element. If the
	// number of operations that are not yet complete is greater than the value
	// that you specified for MaxItems, you can use Marker to return additional
	// operations. Get the value of NextPageMarker from the previous response, and
	// submit another request that includes the value of NextPageMarker in the Marker
	// element.
	Marker *string `type:"string"`

	// Number of domains to be returned.
	//
	// Default: 20
	MaxItems *int64 `type:"integer"`

	// An optional parameter that lets you get information about all the operations
	// that you submitted after a specified date and time. Specify the date and
	// time in Unix time format and Coordinated Universal time (UTC).
	SubmittedSince *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ListOperationsInput) String() string {
	return awsutil.Prettify(s)
}

// The ListOperations response includes the following elements.
type ListOperationsOutput struct {
	_ struct{} `type:"structure"`

	// If there are more operations than you specified for MaxItems in the request,
	// submit another request and include the value of NextPageMarker in the value
	// of Marker.
	NextPageMarker *string `type:"string"`

	// Lists summaries of the operations.
	//
	// Operations is a required field
	Operations []OperationSummary `type:"list" required:"true"`
}

// String returns the string representation
func (s ListOperationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListOperations = "ListOperations"

// ListOperationsRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// Returns information about all of the operations that return an operation
// ID and that have ever been performed on domains that were registered by the
// current account.
//
//    // Example sending a request using ListOperationsRequest.
//    req := client.ListOperationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ListOperations
func (c *Client) ListOperationsRequest(input *ListOperationsInput) ListOperationsRequest {
	op := &aws.Operation{
		Name:       opListOperations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextPageMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListOperationsInput{}
	}

	req := c.newRequest(op, input, &ListOperationsOutput{})

	return ListOperationsRequest{Request: req, Input: input, Copy: c.ListOperationsRequest}
}

// ListOperationsRequest is the request type for the
// ListOperations API operation.
type ListOperationsRequest struct {
	*aws.Request
	Input *ListOperationsInput
	Copy  func(*ListOperationsInput) ListOperationsRequest
}

// Send marshals and sends the ListOperations API request.
func (r ListOperationsRequest) Send(ctx context.Context) (*ListOperationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOperationsResponse{
		ListOperationsOutput: r.Request.Data.(*ListOperationsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListOperationsRequestPaginator returns a paginator for ListOperations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListOperationsRequest(input)
//   p := route53domains.NewListOperationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListOperationsPaginator(req ListOperationsRequest) ListOperationsPaginator {
	return ListOperationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListOperationsInput
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

// ListOperationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListOperationsPaginator struct {
	aws.Pager
}

func (p *ListOperationsPaginator) CurrentPage() *ListOperationsOutput {
	return p.Pager.CurrentPage().(*ListOperationsOutput)
}

// ListOperationsResponse is the response type for the
// ListOperations API operation.
type ListOperationsResponse struct {
	*ListOperationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOperations request.
func (r *ListOperationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
