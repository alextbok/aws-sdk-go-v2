// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Represents a request to begin email address verification with Amazon SES.
// For information about email address verification, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html).
type VerifyEmailAddressInput struct {
	_ struct{} `type:"structure"`

	// The email address to be verified.
	//
	// EmailAddress is a required field
	EmailAddress *string `type:"string" required:"true"`
}

// String returns the string representation
func (s VerifyEmailAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifyEmailAddressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "VerifyEmailAddressInput"}

	if s.EmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type VerifyEmailAddressOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s VerifyEmailAddressOutput) String() string {
	return awsutil.Prettify(s)
}

const opVerifyEmailAddress = "VerifyEmailAddress"

// VerifyEmailAddressRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deprecated. Use the VerifyEmailIdentity operation to verify a new email address.
//
//    // Example sending a request using VerifyEmailAddressRequest.
//    req := client.VerifyEmailAddressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/VerifyEmailAddress
func (c *Client) VerifyEmailAddressRequest(input *VerifyEmailAddressInput) VerifyEmailAddressRequest {
	op := &aws.Operation{
		Name:       opVerifyEmailAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &VerifyEmailAddressInput{}
	}

	req := c.newRequest(op, input, &VerifyEmailAddressOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return VerifyEmailAddressRequest{Request: req, Input: input, Copy: c.VerifyEmailAddressRequest}
}

// VerifyEmailAddressRequest is the request type for the
// VerifyEmailAddress API operation.
type VerifyEmailAddressRequest struct {
	*aws.Request
	Input *VerifyEmailAddressInput
	Copy  func(*VerifyEmailAddressInput) VerifyEmailAddressRequest
}

// Send marshals and sends the VerifyEmailAddress API request.
func (r VerifyEmailAddressRequest) Send(ctx context.Context) (*VerifyEmailAddressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &VerifyEmailAddressResponse{
		VerifyEmailAddressOutput: r.Request.Data.(*VerifyEmailAddressOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// VerifyEmailAddressResponse is the response type for the
// VerifyEmailAddress API operation.
type VerifyEmailAddressResponse struct {
	*VerifyEmailAddressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// VerifyEmailAddress request.
func (r *VerifyEmailAddressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
