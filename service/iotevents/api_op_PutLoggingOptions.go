// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type PutLoggingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The new values of the AWS IoT Events logging options.
	//
	// LoggingOptions is a required field
	LoggingOptions *LoggingOptions `locationName:"loggingOptions" type:"structure" required:"true"`
}

// String returns the string representation
func (s PutLoggingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLoggingOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLoggingOptionsInput"}

	if s.LoggingOptions == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggingOptions"))
	}
	if s.LoggingOptions != nil {
		if err := s.LoggingOptions.Validate(); err != nil {
			invalidParams.AddNested("LoggingOptions", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutLoggingOptionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LoggingOptions != nil {
		v := s.LoggingOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "loggingOptions", v, metadata)
	}
	return nil
}

type PutLoggingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutLoggingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutLoggingOptionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutLoggingOptions = "PutLoggingOptions"

// PutLoggingOptionsRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Sets or updates the AWS IoT Events logging options.
//
// If you update the value of any loggingOptions field, it takes up to one minute
// for the change to take effect. If you change the policy attached to the role
// you specified in the roleArn field (for example, to correct an invalid policy),
// it takes up to five minutes for that change to take effect.
//
//    // Example sending a request using PutLoggingOptionsRequest.
//    req := client.PutLoggingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/PutLoggingOptions
func (c *Client) PutLoggingOptionsRequest(input *PutLoggingOptionsInput) PutLoggingOptionsRequest {
	op := &aws.Operation{
		Name:       opPutLoggingOptions,
		HTTPMethod: "PUT",
		HTTPPath:   "/logging",
	}

	if input == nil {
		input = &PutLoggingOptionsInput{}
	}

	req := c.newRequest(op, input, &PutLoggingOptionsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return PutLoggingOptionsRequest{Request: req, Input: input, Copy: c.PutLoggingOptionsRequest}
}

// PutLoggingOptionsRequest is the request type for the
// PutLoggingOptions API operation.
type PutLoggingOptionsRequest struct {
	*aws.Request
	Input *PutLoggingOptionsInput
	Copy  func(*PutLoggingOptionsInput) PutLoggingOptionsRequest
}

// Send marshals and sends the PutLoggingOptions API request.
func (r PutLoggingOptionsRequest) Send(ctx context.Context) (*PutLoggingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLoggingOptionsResponse{
		PutLoggingOptionsOutput: r.Request.Data.(*PutLoggingOptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLoggingOptionsResponse is the response type for the
// PutLoggingOptions API operation.
type PutLoggingOptionsResponse struct {
	*PutLoggingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLoggingOptions request.
func (r *PutLoggingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
