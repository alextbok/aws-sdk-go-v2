// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package outposts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateOutpostInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone.
	//
	// You must specify AvailabilityZone or AvailabilityZoneId.
	AvailabilityZone *string `min:"1" type:"string"`

	// The ID of the Availability Zone.
	//
	// You must specify AvailabilityZone or AvailabilityZoneId.
	AvailabilityZoneId *string `min:"1" type:"string"`

	// The Outpost description.
	Description *string `min:"1" type:"string"`

	// The name of the Outpost.
	Name *string `min:"1" type:"string"`

	// The ID of the site.
	//
	// SiteId is a required field
	SiteId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateOutpostInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOutpostInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOutpostInput"}
	if s.AvailabilityZone != nil && len(*s.AvailabilityZone) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AvailabilityZone", 1))
	}
	if s.AvailabilityZoneId != nil && len(*s.AvailabilityZoneId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AvailabilityZoneId", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SiteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SiteId"))
	}
	if s.SiteId != nil && len(*s.SiteId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SiteId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOutpostInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AvailabilityZone != nil {
		v := *s.AvailabilityZone

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AvailabilityZone", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AvailabilityZoneId != nil {
		v := *s.AvailabilityZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AvailabilityZoneId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SiteId != nil {
		v := *s.SiteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SiteId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateOutpostOutput struct {
	_ struct{} `type:"structure"`

	// Information about an Outpost.
	Outpost *Outpost `type:"structure"`
}

// String returns the string representation
func (s CreateOutpostOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOutpostOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Outpost != nil {
		v := s.Outpost

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Outpost", v, metadata)
	}
	return nil
}

const opCreateOutpost = "CreateOutpost"

// CreateOutpostRequest returns a request value for making API operation for
// AWS Outposts.
//
// Creates an Outpost.
//
//    // Example sending a request using CreateOutpostRequest.
//    req := client.CreateOutpostRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/outposts-2019-12-03/CreateOutpost
func (c *Client) CreateOutpostRequest(input *CreateOutpostInput) CreateOutpostRequest {
	op := &aws.Operation{
		Name:       opCreateOutpost,
		HTTPMethod: "POST",
		HTTPPath:   "/outposts",
	}

	if input == nil {
		input = &CreateOutpostInput{}
	}

	req := c.newRequest(op, input, &CreateOutpostOutput{})
	return CreateOutpostRequest{Request: req, Input: input, Copy: c.CreateOutpostRequest}
}

// CreateOutpostRequest is the request type for the
// CreateOutpost API operation.
type CreateOutpostRequest struct {
	*aws.Request
	Input *CreateOutpostInput
	Copy  func(*CreateOutpostInput) CreateOutpostRequest
}

// Send marshals and sends the CreateOutpost API request.
func (r CreateOutpostRequest) Send(ctx context.Context) (*CreateOutpostResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOutpostResponse{
		CreateOutpostOutput: r.Request.Data.(*CreateOutpostOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOutpostResponse is the response type for the
// CreateOutpost API operation.
type CreateOutpostResponse struct {
	*CreateOutpostOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOutpost request.
func (r *CreateOutpostResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
