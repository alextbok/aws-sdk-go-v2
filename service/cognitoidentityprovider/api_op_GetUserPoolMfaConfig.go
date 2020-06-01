// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetUserPoolMfaConfigInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserPoolMfaConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserPoolMfaConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserPoolMfaConfigInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetUserPoolMfaConfigOutput struct {
	_ struct{} `type:"structure"`

	// The multi-factor (MFA) configuration. Valid values include:
	//
	//    * OFF MFA will not be used for any users.
	//
	//    * ON MFA is required for all users to sign in.
	//
	//    * OPTIONAL MFA will be required only for individual users who have an
	//    MFA factor enabled.
	MfaConfiguration UserPoolMfaType `type:"string" enum:"true"`

	// The SMS text message multi-factor (MFA) configuration.
	SmsMfaConfiguration *SmsMfaConfigType `type:"structure"`

	// The software token multi-factor (MFA) configuration.
	SoftwareTokenMfaConfiguration *SoftwareTokenMfaConfigType `type:"structure"`
}

// String returns the string representation
func (s GetUserPoolMfaConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUserPoolMfaConfig = "GetUserPoolMfaConfig"

// GetUserPoolMfaConfigRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Gets the user pool multi-factor authentication (MFA) configuration.
//
//    // Example sending a request using GetUserPoolMfaConfigRequest.
//    req := client.GetUserPoolMfaConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUserPoolMfaConfig
func (c *Client) GetUserPoolMfaConfigRequest(input *GetUserPoolMfaConfigInput) GetUserPoolMfaConfigRequest {
	op := &aws.Operation{
		Name:       opGetUserPoolMfaConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUserPoolMfaConfigInput{}
	}

	req := c.newRequest(op, input, &GetUserPoolMfaConfigOutput{})

	return GetUserPoolMfaConfigRequest{Request: req, Input: input, Copy: c.GetUserPoolMfaConfigRequest}
}

// GetUserPoolMfaConfigRequest is the request type for the
// GetUserPoolMfaConfig API operation.
type GetUserPoolMfaConfigRequest struct {
	*aws.Request
	Input *GetUserPoolMfaConfigInput
	Copy  func(*GetUserPoolMfaConfigInput) GetUserPoolMfaConfigRequest
}

// Send marshals and sends the GetUserPoolMfaConfig API request.
func (r GetUserPoolMfaConfigRequest) Send(ctx context.Context) (*GetUserPoolMfaConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserPoolMfaConfigResponse{
		GetUserPoolMfaConfigOutput: r.Request.Data.(*GetUserPoolMfaConfigOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUserPoolMfaConfigResponse is the response type for the
// GetUserPoolMfaConfig API operation.
type GetUserPoolMfaConfigResponse struct {
	*GetUserPoolMfaConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUserPoolMfaConfig request.
func (r *GetUserPoolMfaConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
