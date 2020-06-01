// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DeprecateThingType operation.
type DeprecateThingTypeInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing type to deprecate.
	//
	// ThingTypeName is a required field
	ThingTypeName *string `location:"uri" locationName:"thingTypeName" min:"1" type:"string" required:"true"`

	// Whether to undeprecate a deprecated thing type. If true, the thing type will
	// not be deprecated anymore and you can associate it with things.
	UndoDeprecate *bool `locationName:"undoDeprecate" type:"boolean"`
}

// String returns the string representation
func (s DeprecateThingTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeprecateThingTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeprecateThingTypeInput"}

	if s.ThingTypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingTypeName"))
	}
	if s.ThingTypeName != nil && len(*s.ThingTypeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingTypeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeprecateThingTypeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.UndoDeprecate != nil {
		v := *s.UndoDeprecate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "undoDeprecate", protocol.BoolValue(v), metadata)
	}
	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output for the DeprecateThingType operation.
type DeprecateThingTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeprecateThingTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeprecateThingTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeprecateThingType = "DeprecateThingType"

// DeprecateThingTypeRequest returns a request value for making API operation for
// AWS IoT.
//
// Deprecates a thing type. You can not associate new things with deprecated
// thing type.
//
//    // Example sending a request using DeprecateThingTypeRequest.
//    req := client.DeprecateThingTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeprecateThingTypeRequest(input *DeprecateThingTypeInput) DeprecateThingTypeRequest {
	op := &aws.Operation{
		Name:       opDeprecateThingType,
		HTTPMethod: "POST",
		HTTPPath:   "/thing-types/{thingTypeName}/deprecate",
	}

	if input == nil {
		input = &DeprecateThingTypeInput{}
	}

	req := c.newRequest(op, input, &DeprecateThingTypeOutput{})

	return DeprecateThingTypeRequest{Request: req, Input: input, Copy: c.DeprecateThingTypeRequest}
}

// DeprecateThingTypeRequest is the request type for the
// DeprecateThingType API operation.
type DeprecateThingTypeRequest struct {
	*aws.Request
	Input *DeprecateThingTypeInput
	Copy  func(*DeprecateThingTypeInput) DeprecateThingTypeRequest
}

// Send marshals and sends the DeprecateThingType API request.
func (r DeprecateThingTypeRequest) Send(ctx context.Context) (*DeprecateThingTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeprecateThingTypeResponse{
		DeprecateThingTypeOutput: r.Request.Data.(*DeprecateThingTypeOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeprecateThingTypeResponse is the response type for the
// DeprecateThingType API operation.
type DeprecateThingTypeResponse struct {
	*DeprecateThingTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeprecateThingType request.
func (r *DeprecateThingTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
