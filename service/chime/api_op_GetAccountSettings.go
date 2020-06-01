// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetAccountSettingsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccountSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccountSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccountSettingsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetAccountSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account settings.
	AccountSettings *AccountSettings `type:"structure"`
}

// String returns the string representation
func (s GetAccountSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccountSettings != nil {
		v := s.AccountSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "AccountSettings", v, metadata)
	}
	return nil
}

const opGetAccountSettings = "GetAccountSettings"

// GetAccountSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves account settings for the specified Amazon Chime account ID, such
// as remote control and dial out settings. For more information about these
// settings, see Use the Policies Page (https://docs.aws.amazon.com/chime/latest/ag/policies.html)
// in the Amazon Chime Administration Guide.
//
//    // Example sending a request using GetAccountSettingsRequest.
//    req := client.GetAccountSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetAccountSettings
func (c *Client) GetAccountSettingsRequest(input *GetAccountSettingsInput) GetAccountSettingsRequest {
	op := &aws.Operation{
		Name:       opGetAccountSettings,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}/settings",
	}

	if input == nil {
		input = &GetAccountSettingsInput{}
	}

	req := c.newRequest(op, input, &GetAccountSettingsOutput{})

	return GetAccountSettingsRequest{Request: req, Input: input, Copy: c.GetAccountSettingsRequest}
}

// GetAccountSettingsRequest is the request type for the
// GetAccountSettings API operation.
type GetAccountSettingsRequest struct {
	*aws.Request
	Input *GetAccountSettingsInput
	Copy  func(*GetAccountSettingsInput) GetAccountSettingsRequest
}

// Send marshals and sends the GetAccountSettings API request.
func (r GetAccountSettingsRequest) Send(ctx context.Context) (*GetAccountSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountSettingsResponse{
		GetAccountSettingsOutput: r.Request.Data.(*GetAccountSettingsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountSettingsResponse is the response type for the
// GetAccountSettings API operation.
type GetAccountSettingsResponse struct {
	*GetAccountSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccountSettings request.
func (r *GetAccountSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
