// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeclineInvitationsInput struct {
	_ struct{} `type:"structure"`

	// The list of account IDs for the accounts from which to decline the invitations
	// to Security Hub.
	//
	// AccountIds is a required field
	AccountIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DeclineInvitationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeclineInvitationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeclineInvitationsInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeclineInvitationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AccountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type DeclineInvitationsOutput struct {
	_ struct{} `type:"structure"`

	// The list of AWS accounts that were not processed. For each account, the list
	// includes the account ID and the email address.
	UnprocessedAccounts []Result `type:"list"`
}

// String returns the string representation
func (s DeclineInvitationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeclineInvitationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDeclineInvitations = "DeclineInvitations"

// DeclineInvitationsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Declines invitations to become a member account.
//
//    // Example sending a request using DeclineInvitationsRequest.
//    req := client.DeclineInvitationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeclineInvitations
func (c *Client) DeclineInvitationsRequest(input *DeclineInvitationsInput) DeclineInvitationsRequest {
	op := &aws.Operation{
		Name:       opDeclineInvitations,
		HTTPMethod: "POST",
		HTTPPath:   "/invitations/decline",
	}

	if input == nil {
		input = &DeclineInvitationsInput{}
	}

	req := c.newRequest(op, input, &DeclineInvitationsOutput{})

	return DeclineInvitationsRequest{Request: req, Input: input, Copy: c.DeclineInvitationsRequest}
}

// DeclineInvitationsRequest is the request type for the
// DeclineInvitations API operation.
type DeclineInvitationsRequest struct {
	*aws.Request
	Input *DeclineInvitationsInput
	Copy  func(*DeclineInvitationsInput) DeclineInvitationsRequest
}

// Send marshals and sends the DeclineInvitations API request.
func (r DeclineInvitationsRequest) Send(ctx context.Context) (*DeclineInvitationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeclineInvitationsResponse{
		DeclineInvitationsOutput: r.Request.Data.(*DeclineInvitationsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeclineInvitationsResponse is the response type for the
// DeclineInvitations API operation.
type DeclineInvitationsResponse struct {
	*DeclineInvitationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeclineInvitations request.
func (r *DeclineInvitationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
