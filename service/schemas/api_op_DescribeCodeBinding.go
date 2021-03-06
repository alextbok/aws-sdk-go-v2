// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeCodeBindingInput struct {
	_ struct{} `type:"structure"`

	// Language is a required field
	Language *string `location:"uri" locationName:"language" type:"string" required:"true"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// SchemaName is a required field
	SchemaName *string `location:"uri" locationName:"schemaName" type:"string" required:"true"`

	SchemaVersion *string `location:"querystring" locationName:"schemaVersion" type:"string"`
}

// String returns the string representation
func (s DescribeCodeBindingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCodeBindingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCodeBindingInput"}

	if s.Language == nil {
		invalidParams.Add(aws.NewErrParamRequired("Language"))
	}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if s.SchemaName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCodeBindingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Language != nil {
		v := *s.Language

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "language", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "schemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "schemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeCodeBindingOutput struct {
	_ struct{} `type:"structure"`

	CreationDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	LastModified *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	SchemaVersion *string `type:"string"`

	Status CodeGenerationStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeCodeBindingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCodeBindingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opDescribeCodeBinding = "DescribeCodeBinding"

// DescribeCodeBindingRequest returns a request value for making API operation for
// Schemas.
//
// Describe the code binding URI.
//
//    // Example sending a request using DescribeCodeBindingRequest.
//    req := client.DescribeCodeBindingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/DescribeCodeBinding
func (c *Client) DescribeCodeBindingRequest(input *DescribeCodeBindingInput) DescribeCodeBindingRequest {
	op := &aws.Operation{
		Name:       opDescribeCodeBinding,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas/name/{schemaName}/language/{language}",
	}

	if input == nil {
		input = &DescribeCodeBindingInput{}
	}

	req := c.newRequest(op, input, &DescribeCodeBindingOutput{})
	return DescribeCodeBindingRequest{Request: req, Input: input, Copy: c.DescribeCodeBindingRequest}
}

// DescribeCodeBindingRequest is the request type for the
// DescribeCodeBinding API operation.
type DescribeCodeBindingRequest struct {
	*aws.Request
	Input *DescribeCodeBindingInput
	Copy  func(*DescribeCodeBindingInput) DescribeCodeBindingRequest
}

// Send marshals and sends the DescribeCodeBinding API request.
func (r DescribeCodeBindingRequest) Send(ctx context.Context) (*DescribeCodeBindingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCodeBindingResponse{
		DescribeCodeBindingOutput: r.Request.Data.(*DescribeCodeBindingOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCodeBindingResponse is the response type for the
// DescribeCodeBinding API operation.
type DescribeCodeBindingResponse struct {
	*DescribeCodeBindingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCodeBinding request.
func (r *DescribeCodeBindingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
