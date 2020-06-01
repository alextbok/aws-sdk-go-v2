// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeDiscovererInput struct {
	_ struct{} `type:"structure"`

	// DiscovererId is a required field
	DiscovererId *string `location:"uri" locationName:"discovererId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDiscovererInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDiscovererInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDiscovererInput"}

	if s.DiscovererId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiscovererId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDiscovererInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DiscovererId != nil {
		v := *s.DiscovererId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "discovererId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeDiscovererOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DiscovererArn *string `type:"string"`

	DiscovererId *string `type:"string"`

	SourceArn *string `type:"string"`

	State DiscovererState `type:"string" enum:"true"`

	// Key-value pairs associated with a resource.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribeDiscovererOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDiscovererOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DiscovererArn != nil {
		v := *s.DiscovererArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DiscovererArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DiscovererId != nil {
		v := *s.DiscovererId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DiscovererId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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

const opDescribeDiscoverer = "DescribeDiscoverer"

// DescribeDiscovererRequest returns a request value for making API operation for
// Schemas.
//
// Describes the discoverer.
//
//    // Example sending a request using DescribeDiscovererRequest.
//    req := client.DescribeDiscovererRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/DescribeDiscoverer
func (c *Client) DescribeDiscovererRequest(input *DescribeDiscovererInput) DescribeDiscovererRequest {
	op := &aws.Operation{
		Name:       opDescribeDiscoverer,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/discoverers/id/{discovererId}",
	}

	if input == nil {
		input = &DescribeDiscovererInput{}
	}

	req := c.newRequest(op, input, &DescribeDiscovererOutput{})

	return DescribeDiscovererRequest{Request: req, Input: input, Copy: c.DescribeDiscovererRequest}
}

// DescribeDiscovererRequest is the request type for the
// DescribeDiscoverer API operation.
type DescribeDiscovererRequest struct {
	*aws.Request
	Input *DescribeDiscovererInput
	Copy  func(*DescribeDiscovererInput) DescribeDiscovererRequest
}

// Send marshals and sends the DescribeDiscoverer API request.
func (r DescribeDiscovererRequest) Send(ctx context.Context) (*DescribeDiscovererResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDiscovererResponse{
		DescribeDiscovererOutput: r.Request.Data.(*DescribeDiscovererOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDiscovererResponse is the response type for the
// DescribeDiscoverer API operation.
type DescribeDiscovererResponse struct {
	*DescribeDiscovererOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDiscoverer request.
func (r *DescribeDiscovererResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
