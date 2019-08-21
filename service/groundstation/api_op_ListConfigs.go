// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListConfigsRequest
type ListConfigsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListConfigsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListConfigsResponse
type ListConfigsOutput struct {
	_ struct{} `type:"structure"`

	ConfigList []ConfigListItem `locationName:"configList" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListConfigsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConfigList != nil {
		v := s.ConfigList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "configList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListConfigs = "ListConfigs"

// ListConfigsRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a list of Config objects.
//
//    // Example sending a request using ListConfigsRequest.
//    req := client.ListConfigsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListConfigs
func (c *Client) ListConfigsRequest(input *ListConfigsInput) ListConfigsRequest {
	op := &aws.Operation{
		Name:       opListConfigs,
		HTTPMethod: "GET",
		HTTPPath:   "/config",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListConfigsInput{}
	}

	req := c.newRequest(op, input, &ListConfigsOutput{})
	return ListConfigsRequest{Request: req, Input: input, Copy: c.ListConfigsRequest}
}

// ListConfigsRequest is the request type for the
// ListConfigs API operation.
type ListConfigsRequest struct {
	*aws.Request
	Input *ListConfigsInput
	Copy  func(*ListConfigsInput) ListConfigsRequest
}

// Send marshals and sends the ListConfigs API request.
func (r ListConfigsRequest) Send(ctx context.Context) (*ListConfigsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConfigsResponse{
		ListConfigsOutput: r.Request.Data.(*ListConfigsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListConfigsRequestPaginator returns a paginator for ListConfigs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListConfigsRequest(input)
//   p := groundstation.NewListConfigsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListConfigsPaginator(req ListConfigsRequest) ListConfigsPaginator {
	return ListConfigsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListConfigsInput
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

// ListConfigsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListConfigsPaginator struct {
	aws.Pager
}

func (p *ListConfigsPaginator) CurrentPage() *ListConfigsOutput {
	return p.Pager.CurrentPage().(*ListConfigsOutput)
}

// ListConfigsResponse is the response type for the
// ListConfigs API operation.
type ListConfigsResponse struct {
	*ListConfigsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConfigs request.
func (r *ListConfigsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
