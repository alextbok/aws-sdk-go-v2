// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing zero or more of the following fields:
//
//    * ListGatewaysInput$Limit
//
//    * ListGatewaysInput$Marker
type ListGatewaysInput struct {
	_ struct{} `type:"structure"`

	// Specifies that the list of gateways returned be limited to the specified
	// number of items.
	Limit *int64 `min:"1" type:"integer"`

	// An opaque string that indicates the position at which to begin the returned
	// list of gateways.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGatewaysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGatewaysInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListGatewaysOutput struct {
	_ struct{} `type:"structure"`

	// An array of GatewayInfo objects.
	Gateways []GatewayInfo `type:"list"`

	// Use the marker in your next request to fetch the next set of gateways in
	// the list. If there are no more gateways to list, this field does not appear
	// in the response.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}

const opListGateways = "ListGateways"

// ListGatewaysRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Lists gateways owned by an AWS account in an AWS Region specified in the
// request. The returned list is ordered by gateway Amazon Resource Name (ARN).
//
// By default, the operation returns a maximum of 100 gateways. This operation
// supports pagination that allows you to optionally reduce the number of gateways
// returned in a response.
//
// If you have more gateways than are returned in a response (that is, the response
// returns only a truncated list of your gateways), the response contains a
// marker that you can specify in your next request to fetch the next page of
// gateways.
//
//    // Example sending a request using ListGatewaysRequest.
//    req := client.ListGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ListGateways
func (c *Client) ListGatewaysRequest(input *ListGatewaysInput) ListGatewaysRequest {
	op := &aws.Operation{
		Name:       opListGateways,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListGatewaysInput{}
	}

	req := c.newRequest(op, input, &ListGatewaysOutput{})

	return ListGatewaysRequest{Request: req, Input: input, Copy: c.ListGatewaysRequest}
}

// ListGatewaysRequest is the request type for the
// ListGateways API operation.
type ListGatewaysRequest struct {
	*aws.Request
	Input *ListGatewaysInput
	Copy  func(*ListGatewaysInput) ListGatewaysRequest
}

// Send marshals and sends the ListGateways API request.
func (r ListGatewaysRequest) Send(ctx context.Context) (*ListGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGatewaysResponse{
		ListGatewaysOutput: r.Request.Data.(*ListGatewaysOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGatewaysRequestPaginator returns a paginator for ListGateways.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGatewaysRequest(input)
//   p := storagegateway.NewListGatewaysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGatewaysPaginator(req ListGatewaysRequest) ListGatewaysPaginator {
	return ListGatewaysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListGatewaysInput
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

// ListGatewaysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGatewaysPaginator struct {
	aws.Pager
}

func (p *ListGatewaysPaginator) CurrentPage() *ListGatewaysOutput {
	return p.Pager.CurrentPage().(*ListGatewaysOutput)
}

// ListGatewaysResponse is the response type for the
// ListGateways API operation.
type ListGatewaysResponse struct {
	*ListGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGateways request.
func (r *ListGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
