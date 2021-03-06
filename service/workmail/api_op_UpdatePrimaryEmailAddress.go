// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdatePrimaryEmailAddressInput struct {
	_ struct{} `type:"structure"`

	// The value of the email to be updated as primary.
	//
	// Email is a required field
	Email *string `min:"1" type:"string" required:"true"`

	// The user, group, or resource to update.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The organization that contains the user, group, or resource to update.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePrimaryEmailAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePrimaryEmailAddressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePrimaryEmailAddressInput"}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}
	if s.Email != nil && len(*s.Email) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Email", 1))
	}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdatePrimaryEmailAddressOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdatePrimaryEmailAddressOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdatePrimaryEmailAddress = "UpdatePrimaryEmailAddress"

// UpdatePrimaryEmailAddressRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Updates the primary email for a user, group, or resource. The current email
// is moved into the list of aliases (or swapped between an existing alias and
// the current primary email), and the email provided in the input is promoted
// as the primary.
//
//    // Example sending a request using UpdatePrimaryEmailAddressRequest.
//    req := client.UpdatePrimaryEmailAddressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/UpdatePrimaryEmailAddress
func (c *Client) UpdatePrimaryEmailAddressRequest(input *UpdatePrimaryEmailAddressInput) UpdatePrimaryEmailAddressRequest {
	op := &aws.Operation{
		Name:       opUpdatePrimaryEmailAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePrimaryEmailAddressInput{}
	}

	req := c.newRequest(op, input, &UpdatePrimaryEmailAddressOutput{})
	return UpdatePrimaryEmailAddressRequest{Request: req, Input: input, Copy: c.UpdatePrimaryEmailAddressRequest}
}

// UpdatePrimaryEmailAddressRequest is the request type for the
// UpdatePrimaryEmailAddress API operation.
type UpdatePrimaryEmailAddressRequest struct {
	*aws.Request
	Input *UpdatePrimaryEmailAddressInput
	Copy  func(*UpdatePrimaryEmailAddressInput) UpdatePrimaryEmailAddressRequest
}

// Send marshals and sends the UpdatePrimaryEmailAddress API request.
func (r UpdatePrimaryEmailAddressRequest) Send(ctx context.Context) (*UpdatePrimaryEmailAddressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePrimaryEmailAddressResponse{
		UpdatePrimaryEmailAddressOutput: r.Request.Data.(*UpdatePrimaryEmailAddressOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePrimaryEmailAddressResponse is the response type for the
// UpdatePrimaryEmailAddress API operation.
type UpdatePrimaryEmailAddressResponse struct {
	*UpdatePrimaryEmailAddressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePrimaryEmailAddress request.
func (r *UpdatePrimaryEmailAddressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
