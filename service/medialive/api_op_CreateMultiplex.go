// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMultiplexInput struct {
	_ struct{} `type:"structure"`

	// AvailabilityZones is a required field
	AvailabilityZones []string `locationName:"availabilityZones" type:"list" required:"true"`

	// Contains configuration for a Multiplex event
	//
	// MultiplexSettings is a required field
	MultiplexSettings *MultiplexSettings `locationName:"multiplexSettings" type:"structure" required:"true"`

	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// RequestId is a required field
	RequestId *string `locationName:"requestId" type:"string" required:"true" idempotencyToken:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateMultiplexInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMultiplexInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMultiplexInput"}

	if s.AvailabilityZones == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZones"))
	}

	if s.MultiplexSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("MultiplexSettings"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestId"))
	}
	if s.MultiplexSettings != nil {
		if err := s.MultiplexSettings.Validate(); err != nil {
			invalidParams.AddNested("MultiplexSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMultiplexInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AvailabilityZones != nil {
		v := s.AvailabilityZones

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "availabilityZones", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.MultiplexSettings != nil {
		v := s.MultiplexSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "multiplexSettings", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	var RequestId string
	if s.RequestId != nil {
		RequestId = *s.RequestId
	} else {
		RequestId = protocol.GetIdempotencyToken()
	}
	{
		v := RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateMultiplexOutput struct {
	_ struct{} `type:"structure"`

	// The multiplex object.
	Multiplex *Multiplex `locationName:"multiplex" type:"structure"`
}

// String returns the string representation
func (s CreateMultiplexOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMultiplexOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Multiplex != nil {
		v := s.Multiplex

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "multiplex", v, metadata)
	}
	return nil
}

const opCreateMultiplex = "CreateMultiplex"

// CreateMultiplexRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Create a new multiplex.
//
//    // Example sending a request using CreateMultiplexRequest.
//    req := client.CreateMultiplexRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/CreateMultiplex
func (c *Client) CreateMultiplexRequest(input *CreateMultiplexInput) CreateMultiplexRequest {
	op := &aws.Operation{
		Name:       opCreateMultiplex,
		HTTPMethod: "POST",
		HTTPPath:   "/prod/multiplexes",
	}

	if input == nil {
		input = &CreateMultiplexInput{}
	}

	req := c.newRequest(op, input, &CreateMultiplexOutput{})

	return CreateMultiplexRequest{Request: req, Input: input, Copy: c.CreateMultiplexRequest}
}

// CreateMultiplexRequest is the request type for the
// CreateMultiplex API operation.
type CreateMultiplexRequest struct {
	*aws.Request
	Input *CreateMultiplexInput
	Copy  func(*CreateMultiplexInput) CreateMultiplexRequest
}

// Send marshals and sends the CreateMultiplex API request.
func (r CreateMultiplexRequest) Send(ctx context.Context) (*CreateMultiplexResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMultiplexResponse{
		CreateMultiplexOutput: r.Request.Data.(*CreateMultiplexOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMultiplexResponse is the response type for the
// CreateMultiplex API operation.
type CreateMultiplexResponse struct {
	*CreateMultiplexOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMultiplex request.
func (r *CreateMultiplexResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
