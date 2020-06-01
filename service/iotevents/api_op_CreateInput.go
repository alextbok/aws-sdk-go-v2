// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateInputInput struct {
	_ struct{} `type:"structure"`

	// The definition of the input.
	//
	// InputDefinition is a required field
	InputDefinition *InputDefinition `locationName:"inputDefinition" type:"structure" required:"true"`

	// A brief description of the input.
	InputDescription *string `locationName:"inputDescription" type:"string"`

	// The name you want to give to the input.
	//
	// InputName is a required field
	InputName *string `locationName:"inputName" min:"1" type:"string" required:"true"`

	// Metadata that can be used to manage the input.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateInputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateInputInput"}

	if s.InputDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDefinition"))
	}

	if s.InputName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputName"))
	}
	if s.InputName != nil && len(*s.InputName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputName", 1))
	}
	if s.InputDefinition != nil {
		if err := s.InputDefinition.Validate(); err != nil {
			invalidParams.AddNested("InputDefinition", err.(aws.ErrInvalidParams))
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
func (s CreateInputInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InputDefinition != nil {
		v := s.InputDefinition

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "inputDefinition", v, metadata)
	}
	if s.InputDescription != nil {
		v := *s.InputDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "inputDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InputName != nil {
		v := *s.InputName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "inputName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type CreateInputOutput struct {
	_ struct{} `type:"structure"`

	// Information about the configuration of the input.
	InputConfiguration *InputConfiguration `locationName:"inputConfiguration" type:"structure"`
}

// String returns the string representation
func (s CreateInputOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateInputOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InputConfiguration != nil {
		v := s.InputConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "inputConfiguration", v, metadata)
	}
	return nil
}

const opCreateInput = "CreateInput"

// CreateInputRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Creates an input.
//
//    // Example sending a request using CreateInputRequest.
//    req := client.CreateInputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/CreateInput
func (c *Client) CreateInputRequest(input *CreateInputInput) CreateInputRequest {
	op := &aws.Operation{
		Name:       opCreateInput,
		HTTPMethod: "POST",
		HTTPPath:   "/inputs",
	}

	if input == nil {
		input = &CreateInputInput{}
	}

	req := c.newRequest(op, input, &CreateInputOutput{})

	return CreateInputRequest{Request: req, Input: input, Copy: c.CreateInputRequest}
}

// CreateInputRequest is the request type for the
// CreateInput API operation.
type CreateInputRequest struct {
	*aws.Request
	Input *CreateInputInput
	Copy  func(*CreateInputInput) CreateInputRequest
}

// Send marshals and sends the CreateInput API request.
func (r CreateInputRequest) Send(ctx context.Context) (*CreateInputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInputResponse{
		CreateInputOutput: r.Request.Data.(*CreateInputOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInputResponse is the response type for the
// CreateInput API operation.
type CreateInputResponse struct {
	*CreateInputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInput request.
func (r *CreateInputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
