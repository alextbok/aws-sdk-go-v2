// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetCustomDataIdentifierInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCustomDataIdentifierInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCustomDataIdentifierInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCustomDataIdentifierInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCustomDataIdentifierInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about the criteria and other settings for a custom data
// identifier.
type GetCustomDataIdentifierOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"iso8601"`

	Deleted *bool `locationName:"deleted" type:"boolean"`

	Description *string `locationName:"description" type:"string"`

	Id *string `locationName:"id" type:"string"`

	IgnoreWords []string `locationName:"ignoreWords" type:"list"`

	Keywords []string `locationName:"keywords" type:"list"`

	MaximumMatchDistance *int64 `locationName:"maximumMatchDistance" type:"integer"`

	Name *string `locationName:"name" type:"string"`

	Regex *string `locationName:"regex" type:"string"`

	// A string-to-string map of key-value pairs that specifies the tags (keys and
	// values) for a classification job, custom data identifier, findings filter,
	// or member account.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetCustomDataIdentifierOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCustomDataIdentifierOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Deleted != nil {
		v := *s.Deleted

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deleted", protocol.BoolValue(v), metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IgnoreWords != nil {
		v := s.IgnoreWords

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ignoreWords", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Keywords != nil {
		v := s.Keywords

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "keywords", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.MaximumMatchDistance != nil {
		v := *s.MaximumMatchDistance

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maximumMatchDistance", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Regex != nil {
		v := *s.Regex

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "regex", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

const opGetCustomDataIdentifier = "GetCustomDataIdentifier"

// GetCustomDataIdentifierRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves information about the criteria and other settings for a custom
// data identifier.
//
//    // Example sending a request using GetCustomDataIdentifierRequest.
//    req := client.GetCustomDataIdentifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/GetCustomDataIdentifier
func (c *Client) GetCustomDataIdentifierRequest(input *GetCustomDataIdentifierInput) GetCustomDataIdentifierRequest {
	op := &aws.Operation{
		Name:       opGetCustomDataIdentifier,
		HTTPMethod: "GET",
		HTTPPath:   "/custom-data-identifiers/{id}",
	}

	if input == nil {
		input = &GetCustomDataIdentifierInput{}
	}

	req := c.newRequest(op, input, &GetCustomDataIdentifierOutput{})

	return GetCustomDataIdentifierRequest{Request: req, Input: input, Copy: c.GetCustomDataIdentifierRequest}
}

// GetCustomDataIdentifierRequest is the request type for the
// GetCustomDataIdentifier API operation.
type GetCustomDataIdentifierRequest struct {
	*aws.Request
	Input *GetCustomDataIdentifierInput
	Copy  func(*GetCustomDataIdentifierInput) GetCustomDataIdentifierRequest
}

// Send marshals and sends the GetCustomDataIdentifier API request.
func (r GetCustomDataIdentifierRequest) Send(ctx context.Context) (*GetCustomDataIdentifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCustomDataIdentifierResponse{
		GetCustomDataIdentifierOutput: r.Request.Data.(*GetCustomDataIdentifierOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCustomDataIdentifierResponse is the response type for the
// GetCustomDataIdentifier API operation.
type GetCustomDataIdentifierResponse struct {
	*GetCustomDataIdentifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCustomDataIdentifier request.
func (r *GetCustomDataIdentifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}