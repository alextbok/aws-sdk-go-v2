// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssooidc

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateTokenInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier string for each client. This value should come from
	// the persisted result of the RegisterClient API.
	//
	// ClientId is a required field
	ClientId *string `locationName:"clientId" type:"string" required:"true"`

	// A secret string generated for the client. This value should come from the
	// persisted result of the RegisterClient API.
	//
	// ClientSecret is a required field
	ClientSecret *string `locationName:"clientSecret" type:"string" required:"true"`

	// The authorization code received from the authorization service. This parameter
	// is required to perform an authorization grant request to get access to a
	// token.
	Code *string `locationName:"code" type:"string"`

	// Used only when calling this API for the device code grant type. This short-term
	// code is used to identify this authentication attempt. This should come from
	// an in-memory reference to the result of the StartDeviceAuthorization API.
	//
	// DeviceCode is a required field
	DeviceCode *string `locationName:"deviceCode" type:"string" required:"true"`

	// Supports grant types for authorization code, refresh token, and device code
	// request.
	//
	// GrantType is a required field
	GrantType *string `locationName:"grantType" type:"string" required:"true"`

	// The location of the application that will receive the authorization code.
	// Users authorize the service to send the request to this location.
	RedirectUri *string `locationName:"redirectUri" type:"string"`

	// The token used to obtain an access token in the event that the access token
	// is invalid or expired. This token is not issued by the service.
	RefreshToken *string `locationName:"refreshToken" type:"string"`

	// The list of scopes that is defined by the client. Upon authorization, this
	// list is used to restrict permissions when granting an access token.
	Scope []string `locationName:"scope" type:"list"`
}

// String returns the string representation
func (s CreateTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTokenInput"}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}

	if s.ClientSecret == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientSecret"))
	}

	if s.DeviceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceCode"))
	}

	if s.GrantType == nil {
		invalidParams.Add(aws.NewErrParamRequired("GrantType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTokenInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClientId != nil {
		v := *s.ClientId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientSecret != nil {
		v := *s.ClientSecret

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientSecret", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Code != nil {
		v := *s.Code

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "code", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeviceCode != nil {
		v := *s.DeviceCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GrantType != nil {
		v := *s.GrantType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "grantType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RedirectUri != nil {
		v := *s.RedirectUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "redirectUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RefreshToken != nil {
		v := *s.RefreshToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "refreshToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Scope != nil {
		v := s.Scope

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "scope", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type CreateTokenOutput struct {
	_ struct{} `type:"structure"`

	// An opaque token to access AWS SSO resources assigned to a user.
	AccessToken *string `locationName:"accessToken" type:"string"`

	// Indicates the time in seconds when an access token will expire.
	ExpiresIn *int64 `locationName:"expiresIn" type:"integer"`

	// The identifier of the user that associated with the access token, if present.
	IdToken *string `locationName:"idToken" type:"string"`

	// A token that, if present, can be used to refresh a previously issued access
	// token that might have expired.
	RefreshToken *string `locationName:"refreshToken" type:"string"`

	// Used to notify the client that the returned token is an access token. The
	// supported type is BearerToken.
	TokenType *string `locationName:"tokenType" type:"string"`
}

// String returns the string representation
func (s CreateTokenOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTokenOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccessToken != nil {
		v := *s.AccessToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "accessToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExpiresIn != nil {
		v := *s.ExpiresIn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expiresIn", protocol.Int64Value(v), metadata)
	}
	if s.IdToken != nil {
		v := *s.IdToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "idToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RefreshToken != nil {
		v := *s.RefreshToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "refreshToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TokenType != nil {
		v := *s.TokenType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "tokenType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateToken = "CreateToken"

// CreateTokenRequest returns a request value for making API operation for
// AWS SSO OIDC.
//
// Creates and returns an access token for the authorized client. The access
// token issued will be used to fetch short-term credentials for the assigned
// roles in the AWS account.
//
//    // Example sending a request using CreateTokenRequest.
//    req := client.CreateTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sso-oidc-2019-06-10/CreateToken
func (c *Client) CreateTokenRequest(input *CreateTokenInput) CreateTokenRequest {
	op := &aws.Operation{
		Name:       opCreateToken,
		HTTPMethod: "POST",
		HTTPPath:   "/token",
	}

	if input == nil {
		input = &CreateTokenInput{}
	}

	req := c.newRequest(op, input, &CreateTokenOutput{})
	req.Config.Credentials = aws.AnonymousCredentials

	return CreateTokenRequest{Request: req, Input: input, Copy: c.CreateTokenRequest}
}

// CreateTokenRequest is the request type for the
// CreateToken API operation.
type CreateTokenRequest struct {
	*aws.Request
	Input *CreateTokenInput
	Copy  func(*CreateTokenInput) CreateTokenRequest
}

// Send marshals and sends the CreateToken API request.
func (r CreateTokenRequest) Send(ctx context.Context) (*CreateTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTokenResponse{
		CreateTokenOutput: r.Request.Data.(*CreateTokenOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTokenResponse is the response type for the
// CreateToken API operation.
type CreateTokenResponse struct {
	*CreateTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateToken request.
func (r *CreateTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
