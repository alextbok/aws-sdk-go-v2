// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListMeetingsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`
}

// String returns the string representation
func (s ListMeetingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMeetingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMeetingsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMeetingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListMeetingsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting information.
	Meetings []Meeting `type:"list"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMeetingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMeetingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Meetings != nil {
		v := s.Meetings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Meetings", metadata)
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

const opListMeetings = "ListMeetings"

// ListMeetingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists up to 100 active Amazon Chime SDK meetings. For more information about
// the Amazon Chime SDK, see Using the Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
//
//    // Example sending a request using ListMeetingsRequest.
//    req := client.ListMeetingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListMeetings
func (c *Client) ListMeetingsRequest(input *ListMeetingsInput) ListMeetingsRequest {
	op := &aws.Operation{
		Name:       opListMeetings,
		HTTPMethod: "GET",
		HTTPPath:   "/meetings",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMeetingsInput{}
	}

	req := c.newRequest(op, input, &ListMeetingsOutput{})

	return ListMeetingsRequest{Request: req, Input: input, Copy: c.ListMeetingsRequest}
}

// ListMeetingsRequest is the request type for the
// ListMeetings API operation.
type ListMeetingsRequest struct {
	*aws.Request
	Input *ListMeetingsInput
	Copy  func(*ListMeetingsInput) ListMeetingsRequest
}

// Send marshals and sends the ListMeetings API request.
func (r ListMeetingsRequest) Send(ctx context.Context) (*ListMeetingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMeetingsResponse{
		ListMeetingsOutput: r.Request.Data.(*ListMeetingsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMeetingsRequestPaginator returns a paginator for ListMeetings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMeetingsRequest(input)
//   p := chime.NewListMeetingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMeetingsPaginator(req ListMeetingsRequest) ListMeetingsPaginator {
	return ListMeetingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMeetingsInput
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

// ListMeetingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMeetingsPaginator struct {
	aws.Pager
}

func (p *ListMeetingsPaginator) CurrentPage() *ListMeetingsOutput {
	return p.Pager.CurrentPage().(*ListMeetingsOutput)
}

// ListMeetingsResponse is the response type for the
// ListMeetings API operation.
type ListMeetingsResponse struct {
	*ListMeetingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMeetings request.
func (r *ListMeetingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
