// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetApiInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApiInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApiInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetApiOutput struct {
	_ struct{} `type:"structure"`

	ApiEndpoint *string `locationName:"apiEndpoint" type:"string"`

	// The identifier.
	ApiId *string `locationName:"apiId" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiKeySelectionExpression *string `locationName:"apiKeySelectionExpression" type:"string"`

	// Represents a CORS configuration. Supported only for HTTP APIs. See Configuring
	// CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)
	// for more information.
	CorsConfiguration *Cors `locationName:"corsConfiguration" type:"structure"`

	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"iso8601"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	DisableSchemaValidation *bool `locationName:"disableSchemaValidation" type:"boolean"`

	ImportInfo []string `locationName:"importInfo" type:"list"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`

	// Represents a protocol type.
	ProtocolType ProtocolType `locationName:"protocolType" type:"string" enum:"true"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteSelectionExpression *string `locationName:"routeSelectionExpression" type:"string"`

	// Represents a collection of tags associated with the resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// A string with a length between [1-64].
	Version *string `locationName:"version" type:"string"`

	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s GetApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiEndpoint != nil {
		v := *s.ApiEndpoint

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiEndpoint", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiKeySelectionExpression != nil {
		v := *s.ApiKeySelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CorsConfiguration != nil {
		v := s.CorsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "corsConfiguration", v, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DisableSchemaValidation != nil {
		v := *s.DisableSchemaValidation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableSchemaValidation", protocol.BoolValue(v), metadata)
	}
	if s.ImportInfo != nil {
		v := s.ImportInfo

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "importInfo", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ProtocolType) > 0 {
		v := s.ProtocolType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "protocolType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RouteSelectionExpression != nil {
		v := *s.RouteSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opGetApi = "GetApi"

// GetApiRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets an Api resource.
//
//    // Example sending a request using GetApiRequest.
//    req := client.GetApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetApi
func (c *Client) GetApiRequest(input *GetApiInput) GetApiRequest {
	op := &aws.Operation{
		Name:       opGetApi,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}",
	}

	if input == nil {
		input = &GetApiInput{}
	}

	req := c.newRequest(op, input, &GetApiOutput{})

	return GetApiRequest{Request: req, Input: input, Copy: c.GetApiRequest}
}

// GetApiRequest is the request type for the
// GetApi API operation.
type GetApiRequest struct {
	*aws.Request
	Input *GetApiInput
	Copy  func(*GetApiInput) GetApiRequest
}

// Send marshals and sends the GetApi API request.
func (r GetApiRequest) Send(ctx context.Context) (*GetApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApiResponse{
		GetApiOutput: r.Request.Data.(*GetApiOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApiResponse is the response type for the
// GetApi API operation.
type GetApiResponse struct {
	*GetApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApi request.
func (r *GetApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
