// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// Represents an Amazon Resource Name (ARN).
	AuthorizerCredentialsArn *string `locationName:"authorizerCredentialsArn" type:"string"`

	// An integer with a value between [0-3600].
	AuthorizerResultTtlInSeconds *int64 `locationName:"authorizerResultTtlInSeconds" type:"integer"`

	// The authorizer type. For WebSocket APIs, specify REQUEST for a Lambda function
	// using incoming request parameters. For HTTP APIs, specify JWT to use JSON
	// Web Tokens.
	//
	// AuthorizerType is a required field
	AuthorizerType AuthorizerType `locationName:"authorizerType" type:"string" required:"true" enum:"true"`

	// A string representation of a URI with a length between [1-2048].
	AuthorizerUri *string `locationName:"authorizerUri" type:"string"`

	// The identity source for which authorization is requested. For the REQUEST
	// authorizer, this is required when authorization caching is enabled. The value
	// is a comma-separated string of one or more mapping expressions of the specified
	// request parameters. For example, if an Auth header, a Name query string parameter
	// are defined as identity sources, this value is $method.request.header.Auth,
	// $method.request.querystring.Name. These parameters will be used to derive
	// the authorization caching key and to perform runtime validation of the REQUEST
	// authorizer by verifying all of the identity-related request parameters are
	// present, not null and non-empty. Only when this is true does the authorizer
	// invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized
	// response without calling the Lambda function. The valid value is a string
	// of comma-separated mapping expressions of the specified request parameters.
	// When the authorization caching is not enabled, this property is optional.
	//
	// IdentitySource is a required field
	IdentitySource []string `locationName:"identitySource" type:"list" required:"true"`

	// A string with a length between [0-1024].
	IdentityValidationExpression *string `locationName:"identityValidationExpression" type:"string"`

	// Represents the configuration of a JWT authorizer. Required for the JWT authorizer
	// type. Supported only for HTTP APIs.
	JwtConfiguration *JWTConfiguration `locationName:"jwtConfiguration" type:"structure"`

	// A string with a length between [1-128].
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAuthorizerInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}
	if len(s.AuthorizerType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerType"))
	}

	if s.IdentitySource == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentitySource"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthorizerCredentialsArn != nil {
		v := *s.AuthorizerCredentialsArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerCredentialsArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerResultTtlInSeconds != nil {
		v := *s.AuthorizerResultTtlInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerResultTtlInSeconds", protocol.Int64Value(v), metadata)
	}
	if len(s.AuthorizerType) > 0 {
		v := s.AuthorizerType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AuthorizerUri != nil {
		v := *s.AuthorizerUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentitySource != nil {
		v := s.IdentitySource

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "identitySource", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.IdentityValidationExpression != nil {
		v := *s.IdentityValidationExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identityValidationExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JwtConfiguration != nil {
		v := s.JwtConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jwtConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateAuthorizerOutput struct {
	_ struct{} `type:"structure"`

	// Represents an Amazon Resource Name (ARN).
	AuthorizerCredentialsArn *string `locationName:"authorizerCredentialsArn" type:"string"`

	// The identifier.
	AuthorizerId *string `locationName:"authorizerId" type:"string"`

	// An integer with a value between [0-3600].
	AuthorizerResultTtlInSeconds *int64 `locationName:"authorizerResultTtlInSeconds" type:"integer"`

	// The authorizer type. For WebSocket APIs, specify REQUEST for a Lambda function
	// using incoming request parameters. For HTTP APIs, specify JWT to use JSON
	// Web Tokens.
	AuthorizerType AuthorizerType `locationName:"authorizerType" type:"string" enum:"true"`

	// A string representation of a URI with a length between [1-2048].
	AuthorizerUri *string `locationName:"authorizerUri" type:"string"`

	// The identity source for which authorization is requested. For the REQUEST
	// authorizer, this is required when authorization caching is enabled. The value
	// is a comma-separated string of one or more mapping expressions of the specified
	// request parameters. For example, if an Auth header, a Name query string parameter
	// are defined as identity sources, this value is $method.request.header.Auth,
	// $method.request.querystring.Name. These parameters will be used to derive
	// the authorization caching key and to perform runtime validation of the REQUEST
	// authorizer by verifying all of the identity-related request parameters are
	// present, not null and non-empty. Only when this is true does the authorizer
	// invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized
	// response without calling the Lambda function. The valid value is a string
	// of comma-separated mapping expressions of the specified request parameters.
	// When the authorization caching is not enabled, this property is optional.
	IdentitySource []string `locationName:"identitySource" type:"list"`

	// A string with a length between [0-1024].
	IdentityValidationExpression *string `locationName:"identityValidationExpression" type:"string"`

	// Represents the configuration of a JWT authorizer. Required for the JWT authorizer
	// type. Supported only for HTTP APIs.
	JwtConfiguration *JWTConfiguration `locationName:"jwtConfiguration" type:"structure"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s CreateAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuthorizerCredentialsArn != nil {
		v := *s.AuthorizerCredentialsArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerCredentialsArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerId != nil {
		v := *s.AuthorizerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerResultTtlInSeconds != nil {
		v := *s.AuthorizerResultTtlInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerResultTtlInSeconds", protocol.Int64Value(v), metadata)
	}
	if len(s.AuthorizerType) > 0 {
		v := s.AuthorizerType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AuthorizerUri != nil {
		v := *s.AuthorizerUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentitySource != nil {
		v := s.IdentitySource

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "identitySource", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.IdentityValidationExpression != nil {
		v := *s.IdentityValidationExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identityValidationExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JwtConfiguration != nil {
		v := s.JwtConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jwtConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateAuthorizer = "CreateAuthorizer"

// CreateAuthorizerRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Creates an Authorizer for an API.
//
//    // Example sending a request using CreateAuthorizerRequest.
//    req := client.CreateAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateAuthorizer
func (c *Client) CreateAuthorizerRequest(input *CreateAuthorizerInput) CreateAuthorizerRequest {
	op := &aws.Operation{
		Name:       opCreateAuthorizer,
		HTTPMethod: "POST",
		HTTPPath:   "/v2/apis/{apiId}/authorizers",
	}

	if input == nil {
		input = &CreateAuthorizerInput{}
	}

	req := c.newRequest(op, input, &CreateAuthorizerOutput{})

	return CreateAuthorizerRequest{Request: req, Input: input, Copy: c.CreateAuthorizerRequest}
}

// CreateAuthorizerRequest is the request type for the
// CreateAuthorizer API operation.
type CreateAuthorizerRequest struct {
	*aws.Request
	Input *CreateAuthorizerInput
	Copy  func(*CreateAuthorizerInput) CreateAuthorizerRequest
}

// Send marshals and sends the CreateAuthorizer API request.
func (r CreateAuthorizerRequest) Send(ctx context.Context) (*CreateAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAuthorizerResponse{
		CreateAuthorizerOutput: r.Request.Data.(*CreateAuthorizerOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAuthorizerResponse is the response type for the
// CreateAuthorizer API operation.
type CreateAuthorizerResponse struct {
	*CreateAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAuthorizer request.
func (r *CreateAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
