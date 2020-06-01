// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastoredata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListItemsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per API request. For example, you
	// submit a ListItems request with MaxResults set at 500. Although 2,000 items
	// match your request, the service returns no more than the first 500 items.
	// (The service also returns a NextToken value that you can use to fetch the
	// next batch of results.) The service might return fewer results than the MaxResults
	// value.
	//
	// If MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 1,000 results per page.
	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListItems request with MaxResults set at 500. The service
	// returns the first batch of results (up to 500) and a NextToken value. To
	// see the next batch of results, you can submit the ListItems request a second
	// time and specify the NextToken value.
	//
	// Tokens expire after 15 minutes.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// The path in the container from which to retrieve items. Format: <folder name>/<folder
	// name>/<file name>
	Path *string `location:"querystring" locationName:"Path" type:"string"`
}

// String returns the string representation
func (s ListItemsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListItemsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListItemsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListItemsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Path != nil {
		v := *s.Path

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Path", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListItemsOutput struct {
	_ struct{} `type:"structure"`

	// The metadata entries for the folders and objects at the requested path.
	Items []Item `type:"list"`

	// The token that can be used in a request to view the next set of results.
	// For example, you submit a ListItems request that matches 2,000 items with
	// MaxResults set at 500. The service returns the first batch of results (up
	// to 500) and a NextToken value that can be used to fetch the next batch of
	// results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListItemsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListItemsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListItems = "ListItems"

// ListItemsRequest returns a request value for making API operation for
// AWS Elemental MediaStore Data Plane.
//
// Provides a list of metadata entries about folders and objects in the specified
// folder.
//
//    // Example sending a request using ListItemsRequest.
//    req := client.ListItemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-data-2017-09-01/ListItems
func (c *Client) ListItemsRequest(input *ListItemsInput) ListItemsRequest {
	op := &aws.Operation{
		Name:       opListItems,
		HTTPMethod: "GET",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListItemsInput{}
	}

	req := c.newRequest(op, input, &ListItemsOutput{})

	return ListItemsRequest{Request: req, Input: input, Copy: c.ListItemsRequest}
}

// ListItemsRequest is the request type for the
// ListItems API operation.
type ListItemsRequest struct {
	*aws.Request
	Input *ListItemsInput
	Copy  func(*ListItemsInput) ListItemsRequest
}

// Send marshals and sends the ListItems API request.
func (r ListItemsRequest) Send(ctx context.Context) (*ListItemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListItemsResponse{
		ListItemsOutput: r.Request.Data.(*ListItemsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListItemsRequestPaginator returns a paginator for ListItems.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListItemsRequest(input)
//   p := mediastoredata.NewListItemsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListItemsPaginator(req ListItemsRequest) ListItemsPaginator {
	return ListItemsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListItemsInput
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

// ListItemsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListItemsPaginator struct {
	aws.Pager
}

func (p *ListItemsPaginator) CurrentPage() *ListItemsOutput {
	return p.Pager.CurrentPage().(*ListItemsOutput)
}

// ListItemsResponse is the response type for the
// ListItems API operation.
type ListItemsResponse struct {
	*ListItemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListItems request.
func (r *ListItemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
