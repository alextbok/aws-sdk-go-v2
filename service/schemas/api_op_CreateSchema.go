// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateSchemaInput struct {
	_ struct{} `type:"structure"`

	// Content is a required field
	Content *string `min:"1" type:"string" required:"true"`

	Description *string `type:"string"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// SchemaName is a required field
	SchemaName *string `location:"uri" locationName:"schemaName" type:"string" required:"true"`

	// Key-value pairs associated with a resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// Type is a required field
	Type Type `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSchemaInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}
	if s.Content != nil && len(*s.Content) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Content", 1))
	}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if s.SchemaName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaName"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Content != nil {
		v := *s.Content

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Content", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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
	return nil
}

type CreateSchemaOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	LastModified *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	SchemaArn *string `type:"string"`

	SchemaName *string `type:"string"`

	SchemaVersion *string `type:"string"`

	// Key-value pairs associated with a resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	Type *string `type:"string"`

	VersionCreatedDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s CreateSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionCreatedDate != nil {
		v := *s.VersionCreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionCreatedDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

const opCreateSchema = "CreateSchema"

// CreateSchemaRequest returns a request value for making API operation for
// Schemas.
//
// Creates a schema definition.
//
// Inactive schemas will be deleted after two years.
//
//    // Example sending a request using CreateSchemaRequest.
//    req := client.CreateSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/CreateSchema
func (c *Client) CreateSchemaRequest(input *CreateSchemaInput) CreateSchemaRequest {
	op := &aws.Operation{
		Name:       opCreateSchema,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas/name/{schemaName}",
	}

	if input == nil {
		input = &CreateSchemaInput{}
	}

	req := c.newRequest(op, input, &CreateSchemaOutput{})

	return CreateSchemaRequest{Request: req, Input: input, Copy: c.CreateSchemaRequest}
}

// CreateSchemaRequest is the request type for the
// CreateSchema API operation.
type CreateSchemaRequest struct {
	*aws.Request
	Input *CreateSchemaInput
	Copy  func(*CreateSchemaInput) CreateSchemaRequest
}

// Send marshals and sends the CreateSchema API request.
func (r CreateSchemaRequest) Send(ctx context.Context) (*CreateSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSchemaResponse{
		CreateSchemaOutput: r.Request.Data.(*CreateSchemaOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSchemaResponse is the response type for the
// CreateSchema API operation.
type CreateSchemaResponse struct {
	*CreateSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSchema request.
func (r *CreateSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
