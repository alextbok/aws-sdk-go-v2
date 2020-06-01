// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeSignalingChannelInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the signaling channel that you want to describe.
	ChannelARN *string `min:"1" type:"string"`

	// The name of the signaling channel that you want to describe.
	ChannelName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeSignalingChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSignalingChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSignalingChannelInput"}
	if s.ChannelARN != nil && len(*s.ChannelARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelARN", 1))
	}
	if s.ChannelName != nil && len(*s.ChannelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSignalingChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelARN != nil {
		v := *s.ChannelARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ChannelName != nil {
		v := *s.ChannelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeSignalingChannelOutput struct {
	_ struct{} `type:"structure"`

	// A structure that encapsulates the specified signaling channel's metadata
	// and properties.
	ChannelInfo *ChannelInfo `type:"structure"`
}

// String returns the string representation
func (s DescribeSignalingChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSignalingChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ChannelInfo != nil {
		v := s.ChannelInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ChannelInfo", v, metadata)
	}
	return nil
}

const opDescribeSignalingChannel = "DescribeSignalingChannel"

// DescribeSignalingChannelRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Returns the most current information about the signaling channel. You must
// specify either the name or the Amazon Resource Name (ARN) of the channel
// that you want to describe.
//
//    // Example sending a request using DescribeSignalingChannelRequest.
//    req := client.DescribeSignalingChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/DescribeSignalingChannel
func (c *Client) DescribeSignalingChannelRequest(input *DescribeSignalingChannelInput) DescribeSignalingChannelRequest {
	op := &aws.Operation{
		Name:       opDescribeSignalingChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/describeSignalingChannel",
	}

	if input == nil {
		input = &DescribeSignalingChannelInput{}
	}

	req := c.newRequest(op, input, &DescribeSignalingChannelOutput{})

	return DescribeSignalingChannelRequest{Request: req, Input: input, Copy: c.DescribeSignalingChannelRequest}
}

// DescribeSignalingChannelRequest is the request type for the
// DescribeSignalingChannel API operation.
type DescribeSignalingChannelRequest struct {
	*aws.Request
	Input *DescribeSignalingChannelInput
	Copy  func(*DescribeSignalingChannelInput) DescribeSignalingChannelRequest
}

// Send marshals and sends the DescribeSignalingChannel API request.
func (r DescribeSignalingChannelRequest) Send(ctx context.Context) (*DescribeSignalingChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSignalingChannelResponse{
		DescribeSignalingChannelOutput: r.Request.Data.(*DescribeSignalingChannelOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSignalingChannelResponse is the response type for the
// DescribeSignalingChannel API operation.
type DescribeSignalingChannelResponse struct {
	*DescribeSignalingChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSignalingChannel request.
func (r *DescribeSignalingChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
