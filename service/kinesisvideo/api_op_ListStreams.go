// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListStreamsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of streams to return in the response. The default is 10,000.
	MaxResults *int64 `min:"1" type:"integer"`

	// If you specify this parameter, when the result of a ListStreams operation
	// is truncated, the call returns the NextToken in the response. To get another
	// batch of streams, provide this token in your next request.
	NextToken *string `type:"string"`

	// Optional: Returns only streams that satisfy a specific condition. Currently,
	// you can specify only the prefix of a stream name as a condition.
	StreamNameCondition *StreamNameCondition `type:"structure"`
}

// String returns the string representation
func (s ListStreamsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStreamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStreamsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.StreamNameCondition != nil {
		if err := s.StreamNameCondition.Validate(); err != nil {
			invalidParams.AddNested("StreamNameCondition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListStreamsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamNameCondition != nil {
		v := s.StreamNameCondition

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "StreamNameCondition", v, metadata)
	}
	return nil
}

type ListStreamsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, the call returns this element with a token.
	// To get the next batch of streams, use this token in your next request.
	NextToken *string `type:"string"`

	// An array of StreamInfo objects.
	StreamInfoList []StreamInfo `type:"list"`
}

// String returns the string representation
func (s ListStreamsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListStreamsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamInfoList != nil {
		v := s.StreamInfoList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "StreamInfoList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListStreams = "ListStreams"

// ListStreamsRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Returns an array of StreamInfo objects. Each object describes a stream. To
// retrieve only streams that satisfy a specific condition, you can specify
// a StreamNameCondition.
//
//    // Example sending a request using ListStreamsRequest.
//    req := client.ListStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/ListStreams
func (c *Client) ListStreamsRequest(input *ListStreamsInput) ListStreamsRequest {
	op := &aws.Operation{
		Name:       opListStreams,
		HTTPMethod: "POST",
		HTTPPath:   "/listStreams",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListStreamsInput{}
	}

	req := c.newRequest(op, input, &ListStreamsOutput{})

	return ListStreamsRequest{Request: req, Input: input, Copy: c.ListStreamsRequest}
}

// ListStreamsRequest is the request type for the
// ListStreams API operation.
type ListStreamsRequest struct {
	*aws.Request
	Input *ListStreamsInput
	Copy  func(*ListStreamsInput) ListStreamsRequest
}

// Send marshals and sends the ListStreams API request.
func (r ListStreamsRequest) Send(ctx context.Context) (*ListStreamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStreamsResponse{
		ListStreamsOutput: r.Request.Data.(*ListStreamsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListStreamsRequestPaginator returns a paginator for ListStreams.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListStreamsRequest(input)
//   p := kinesisvideo.NewListStreamsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListStreamsPaginator(req ListStreamsRequest) ListStreamsPaginator {
	return ListStreamsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListStreamsInput
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

// ListStreamsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListStreamsPaginator struct {
	aws.Pager
}

func (p *ListStreamsPaginator) CurrentPage() *ListStreamsOutput {
	return p.Pager.CurrentPage().(*ListStreamsOutput)
}

// ListStreamsResponse is the response type for the
// ListStreams API operation.
type ListStreamsResponse struct {
	*ListStreamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStreams request.
func (r *ListStreamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
