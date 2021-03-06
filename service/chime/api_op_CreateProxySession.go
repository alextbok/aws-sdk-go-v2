// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateProxySessionInput struct {
	_ struct{} `type:"structure"`

	// The proxy session capabilities.
	//
	// Capabilities is a required field
	Capabilities []Capability `type:"list" required:"true"`

	// The number of minutes allowed for the proxy session.
	ExpiryMinutes *int64 `min:"1" type:"integer"`

	// The preference for matching the country or area code of the proxy phone number
	// with that of the first participant.
	GeoMatchLevel GeoMatchLevel `type:"string" enum:"true"`

	// The country and area code for the proxy phone number.
	GeoMatchParams *GeoMatchParams `type:"structure"`

	// The name of the proxy session.
	Name *string `type:"string" sensitive:"true"`

	// The preference for proxy phone number reuse, or stickiness, between the same
	// participants across sessions.
	NumberSelectionBehavior NumberSelectionBehavior `type:"string" enum:"true"`

	// The participant phone numbers.
	//
	// ParticipantPhoneNumbers is a required field
	ParticipantPhoneNumbers []string `min:"2" type:"list" required:"true"`

	// The Amazon Chime voice connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProxySessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProxySessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProxySessionInput"}

	if s.Capabilities == nil {
		invalidParams.Add(aws.NewErrParamRequired("Capabilities"))
	}
	if s.ExpiryMinutes != nil && *s.ExpiryMinutes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ExpiryMinutes", 1))
	}

	if s.ParticipantPhoneNumbers == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParticipantPhoneNumbers"))
	}
	if s.ParticipantPhoneNumbers != nil && len(s.ParticipantPhoneNumbers) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ParticipantPhoneNumbers", 2))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.VoiceConnectorId != nil && len(*s.VoiceConnectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VoiceConnectorId", 1))
	}
	if s.GeoMatchParams != nil {
		if err := s.GeoMatchParams.Validate(); err != nil {
			invalidParams.AddNested("GeoMatchParams", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProxySessionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Capabilities != nil {
		v := s.Capabilities

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Capabilities", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ExpiryMinutes != nil {
		v := *s.ExpiryMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExpiryMinutes", protocol.Int64Value(v), metadata)
	}
	if len(s.GeoMatchLevel) > 0 {
		v := s.GeoMatchLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GeoMatchLevel", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.GeoMatchParams != nil {
		v := s.GeoMatchParams

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "GeoMatchParams", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.NumberSelectionBehavior) > 0 {
		v := s.NumberSelectionBehavior

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NumberSelectionBehavior", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ParticipantPhoneNumbers != nil {
		v := s.ParticipantPhoneNumbers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ParticipantPhoneNumbers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateProxySessionOutput struct {
	_ struct{} `type:"structure"`

	// The proxy session details.
	ProxySession *ProxySession `type:"structure"`
}

// String returns the string representation
func (s CreateProxySessionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProxySessionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ProxySession != nil {
		v := s.ProxySession

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ProxySession", v, metadata)
	}
	return nil
}

const opCreateProxySession = "CreateProxySession"

// CreateProxySessionRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates a proxy session on the specified Amazon Chime Voice Connector for
// the specified participant phone numbers.
//
//    // Example sending a request using CreateProxySessionRequest.
//    req := client.CreateProxySessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreateProxySession
func (c *Client) CreateProxySessionRequest(input *CreateProxySessionInput) CreateProxySessionRequest {
	op := &aws.Operation{
		Name:       opCreateProxySession,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/proxy-sessions",
	}

	if input == nil {
		input = &CreateProxySessionInput{}
	}

	req := c.newRequest(op, input, &CreateProxySessionOutput{})
	return CreateProxySessionRequest{Request: req, Input: input, Copy: c.CreateProxySessionRequest}
}

// CreateProxySessionRequest is the request type for the
// CreateProxySession API operation.
type CreateProxySessionRequest struct {
	*aws.Request
	Input *CreateProxySessionInput
	Copy  func(*CreateProxySessionInput) CreateProxySessionRequest
}

// Send marshals and sends the CreateProxySession API request.
func (r CreateProxySessionRequest) Send(ctx context.Context) (*CreateProxySessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProxySessionResponse{
		CreateProxySessionOutput: r.Request.Data.(*CreateProxySessionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProxySessionResponse is the response type for the
// CreateProxySession API operation.
type CreateProxySessionResponse struct {
	*CreateProxySessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProxySession request.
func (r *CreateProxySessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
