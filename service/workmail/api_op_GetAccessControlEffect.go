// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccessControlEffectInput struct {
	_ struct{} `type:"structure"`

	// The access protocol action. Valid values include ActiveSync, AutoDiscover,
	// EWS, IMAP, SMTP, WindowsOutlook, and WebMail.
	//
	// Action is a required field
	Action *string `min:"1" type:"string" required:"true"`

	// The IPv4 address.
	//
	// IpAddress is a required field
	IpAddress *string `min:"1" type:"string" required:"true"`

	// The identifier for the organization.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The user ID.
	//
	// UserId is a required field
	UserId *string `min:"12" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccessControlEffectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccessControlEffectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccessControlEffectInput"}

	if s.Action == nil {
		invalidParams.Add(aws.NewErrParamRequired("Action"))
	}
	if s.Action != nil && len(*s.Action) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Action", 1))
	}

	if s.IpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpAddress"))
	}
	if s.IpAddress != nil && len(*s.IpAddress) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IpAddress", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAccessControlEffectOutput struct {
	_ struct{} `type:"structure"`

	// The rule effect.
	Effect AccessControlRuleEffect `type:"string" enum:"true"`

	// The rules that match the given parameters, resulting in an effect.
	MatchedRules []string `type:"list"`
}

// String returns the string representation
func (s GetAccessControlEffectOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAccessControlEffect = "GetAccessControlEffect"

// GetAccessControlEffectRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Gets the effects of an organization's access control rules as they apply
// to a specified IPv4 address, access protocol action, or user ID.
//
//    // Example sending a request using GetAccessControlEffectRequest.
//    req := client.GetAccessControlEffectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/GetAccessControlEffect
func (c *Client) GetAccessControlEffectRequest(input *GetAccessControlEffectInput) GetAccessControlEffectRequest {
	op := &aws.Operation{
		Name:       opGetAccessControlEffect,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAccessControlEffectInput{}
	}

	req := c.newRequest(op, input, &GetAccessControlEffectOutput{})

	return GetAccessControlEffectRequest{Request: req, Input: input, Copy: c.GetAccessControlEffectRequest}
}

// GetAccessControlEffectRequest is the request type for the
// GetAccessControlEffect API operation.
type GetAccessControlEffectRequest struct {
	*aws.Request
	Input *GetAccessControlEffectInput
	Copy  func(*GetAccessControlEffectInput) GetAccessControlEffectRequest
}

// Send marshals and sends the GetAccessControlEffect API request.
func (r GetAccessControlEffectRequest) Send(ctx context.Context) (*GetAccessControlEffectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccessControlEffectResponse{
		GetAccessControlEffectOutput: r.Request.Data.(*GetAccessControlEffectOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccessControlEffectResponse is the response type for the
// GetAccessControlEffect API operation.
type GetAccessControlEffectResponse struct {
	*GetAccessControlEffectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccessControlEffect request.
func (r *GetAccessControlEffectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
