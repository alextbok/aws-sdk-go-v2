// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateSignalingChannelInput struct {
	_ struct{} `type:"structure"`

	// A name for the signaling channel that you are creating. It must be unique
	// for each AWS account and AWS Region.
	//
	// ChannelName is a required field
	ChannelName *string `min:"1" type:"string" required:"true"`

	// A type of the signaling channel that you are creating. Currently, SINGLE_MASTER
	// is the only supported channel type.
	ChannelType ChannelType `type:"string" enum:"true"`

	// A structure containing the configuration for the SINGLE_MASTER channel type.
	SingleMasterConfiguration *SingleMasterConfiguration `type:"structure"`

	// A set of tags (key-value pairs) that you want to associate with this channel.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateSignalingChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSignalingChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSignalingChannelInput"}

	if s.ChannelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelName"))
	}
	if s.ChannelName != nil && len(*s.ChannelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelName", 1))
	}
	if s.SingleMasterConfiguration != nil {
		if err := s.SingleMasterConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SingleMasterConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSignalingChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelName != nil {
		v := *s.ChannelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ChannelType) > 0 {
		v := s.ChannelType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SingleMasterConfiguration != nil {
		v := s.SingleMasterConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SingleMasterConfiguration", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type CreateSignalingChannelOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the created channel.
	ChannelARN *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateSignalingChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSignalingChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ChannelARN != nil {
		v := *s.ChannelARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateSignalingChannel = "CreateSignalingChannel"

// CreateSignalingChannelRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Creates a signaling channel.
//
// CreateSignalingChannel is an asynchronous operation.
//
//    // Example sending a request using CreateSignalingChannelRequest.
//    req := client.CreateSignalingChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/CreateSignalingChannel
func (c *Client) CreateSignalingChannelRequest(input *CreateSignalingChannelInput) CreateSignalingChannelRequest {
	op := &aws.Operation{
		Name:       opCreateSignalingChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/createSignalingChannel",
	}

	if input == nil {
		input = &CreateSignalingChannelInput{}
	}

	req := c.newRequest(op, input, &CreateSignalingChannelOutput{})

	return CreateSignalingChannelRequest{Request: req, Input: input, Copy: c.CreateSignalingChannelRequest}
}

// CreateSignalingChannelRequest is the request type for the
// CreateSignalingChannel API operation.
type CreateSignalingChannelRequest struct {
	*aws.Request
	Input *CreateSignalingChannelInput
	Copy  func(*CreateSignalingChannelInput) CreateSignalingChannelRequest
}

// Send marshals and sends the CreateSignalingChannel API request.
func (r CreateSignalingChannelRequest) Send(ctx context.Context) (*CreateSignalingChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSignalingChannelResponse{
		CreateSignalingChannelOutput: r.Request.Data.(*CreateSignalingChannelOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSignalingChannelResponse is the response type for the
// CreateSignalingChannel API operation.
type CreateSignalingChannelResponse struct {
	*CreateSignalingChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSignalingChannel request.
func (r *CreateSignalingChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
