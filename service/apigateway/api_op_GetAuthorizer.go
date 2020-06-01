// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to describe an existing Authorizer resource.
type GetAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the Authorizer resource.
	//
	// AuthorizerId is a required field
	AuthorizerId *string `location:"uri" locationName:"authorizer_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAuthorizerInput"}

	if s.AuthorizerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthorizerId != nil {
		v := *s.AuthorizerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "authorizer_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents an authorization layer for methods. If enabled on a method, API
// Gateway will activate the authorizer when a client calls the method.
//
// Use Lambda Function as Authorizer (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html)
// Use Cognito User Pool as Authorizer (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html)
type GetAuthorizerOutput struct {
	_ struct{} `type:"structure"`

	// Optional customer-defined field, used in OpenAPI imports and exports without
	// functional impact.
	AuthType *string `locationName:"authType" type:"string"`

	// Specifies the required credentials as an IAM role for API Gateway to invoke
	// the authorizer. To specify an IAM role for API Gateway to assume, use the
	// role's Amazon Resource Name (ARN). To use resource-based permissions on the
	// Lambda function, specify null.
	AuthorizerCredentials *string `locationName:"authorizerCredentials" type:"string"`

	// The TTL in seconds of cached authorizer results. If it equals 0, authorization
	// caching is disabled. If it is greater than 0, API Gateway will cache authorizer
	// responses. If this field is not set, the default value is 300. The maximum
	// value is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *int64 `locationName:"authorizerResultTtlInSeconds" type:"integer"`

	// Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or
	// REQUEST authorizers, this must be a well-formed Lambda function URI, for
	// example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form arn:aws:apigateway:{region}:lambda:path/{service_api},
	// where {region} is the same as the region hosting the Lambda function, path
	// indicates that the remaining substring in the URI should be treated as the
	// path to the resource, including the initial /. For Lambda functions, this
	// is usually of the form /2015-03-31/functions/[FunctionARN]/invocations.
	AuthorizerUri *string `locationName:"authorizerUri" type:"string"`

	// The identifier for the authorizer resource.
	Id *string `locationName:"id" type:"string"`

	// The identity source for which authorization is requested.
	//    * For a TOKEN or COGNITO_USER_POOLS authorizer, this is required and specifies
	//    the request header mapping expression for the custom header holding the
	//    authorization token submitted by the client. For example, if the token
	//    header name is Auth, the header mapping expression is method.request.header.Auth.
	//
	//    * For the REQUEST authorizer, this is required when authorization caching
	//    is enabled. The value is a comma-separated string of one or more mapping
	//    expressions of the specified request parameters. For example, if an Auth
	//    header, a Name query string parameter are defined as identity sources,
	//    this value is method.request.header.Auth, method.request.querystring.Name.
	//    These parameters will be used to derive the authorization caching key
	//    and to perform runtime validation of the REQUEST authorizer by verifying
	//    all of the identity-related request parameters are present, not null and
	//    non-empty. Only when this is true does the authorizer invoke the authorizer
	//    Lambda function, otherwise, it returns a 401 Unauthorized response without
	//    calling the Lambda function. The valid value is a string of comma-separated
	//    mapping expressions of the specified request parameters. When the authorization
	//    caching is not enabled, this property is optional.
	IdentitySource *string `locationName:"identitySource" type:"string"`

	// A validation expression for the incoming identity token. For TOKEN authorizers,
	// this value is a regular expression. For COGNITO_USER_POOLS authorizers, API
	// Gateway will match the aud field of the incoming token from the client against
	// the specified regular expression. It will invoke the authorizer's Lambda
	// function when there is a match. Otherwise, it will return a 401 Unauthorized
	// response without calling the Lambda function. The validation expression does
	// not apply to the REQUEST authorizer.
	IdentityValidationExpression *string `locationName:"identityValidationExpression" type:"string"`

	// [Required] The name of the authorizer.
	Name *string `locationName:"name" type:"string"`

	// A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.
	// Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}.
	// For a TOKEN or REQUEST authorizer, this is not defined.
	ProviderARNs []string `locationName:"providerARNs" type:"list"`

	// The authorizer type. Valid values are TOKEN for a Lambda function using a
	// single authorization token submitted in a custom header, REQUEST for a Lambda
	// function using incoming request parameters, and COGNITO_USER_POOLS for using
	// an Amazon Cognito user pool.
	Type AuthorizerType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuthType != nil {
		v := *s.AuthType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerCredentials != nil {
		v := *s.AuthorizerCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerResultTtlInSeconds != nil {
		v := *s.AuthorizerResultTtlInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerResultTtlInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.AuthorizerUri != nil {
		v := *s.AuthorizerUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentitySource != nil {
		v := *s.IdentitySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identitySource", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityValidationExpression != nil {
		v := *s.IdentityValidationExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identityValidationExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProviderARNs != nil {
		v := s.ProviderARNs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "providerARNs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opGetAuthorizer = "GetAuthorizer"

// GetAuthorizerRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Describe an existing Authorizer resource.
//
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-authorizer.html)
//
//    // Example sending a request using GetAuthorizerRequest.
//    req := client.GetAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetAuthorizerRequest(input *GetAuthorizerInput) GetAuthorizerRequest {
	op := &aws.Operation{
		Name:       opGetAuthorizer,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/authorizers/{authorizer_id}",
	}

	if input == nil {
		input = &GetAuthorizerInput{}
	}

	req := c.newRequest(op, input, &GetAuthorizerOutput{})

	return GetAuthorizerRequest{Request: req, Input: input, Copy: c.GetAuthorizerRequest}
}

// GetAuthorizerRequest is the request type for the
// GetAuthorizer API operation.
type GetAuthorizerRequest struct {
	*aws.Request
	Input *GetAuthorizerInput
	Copy  func(*GetAuthorizerInput) GetAuthorizerRequest
}

// Send marshals and sends the GetAuthorizer API request.
func (r GetAuthorizerRequest) Send(ctx context.Context) (*GetAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAuthorizerResponse{
		GetAuthorizerOutput: r.Request.Data.(*GetAuthorizerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAuthorizerResponse is the response type for the
// GetAuthorizer API operation.
type GetAuthorizerResponse struct {
	*GetAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAuthorizer request.
func (r *GetAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
