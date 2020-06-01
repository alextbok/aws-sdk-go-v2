// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateFromMasterAccountInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateFromMasterAccountInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateFromMasterAccountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type DisassociateFromMasterAccountOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateFromMasterAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateFromMasterAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisassociateFromMasterAccount = "DisassociateFromMasterAccount"

// DisassociateFromMasterAccountRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Disassociates a member account from its Amazon Macie master account.
//
//    // Example sending a request using DisassociateFromMasterAccountRequest.
//    req := client.DisassociateFromMasterAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/DisassociateFromMasterAccount
func (c *Client) DisassociateFromMasterAccountRequest(input *DisassociateFromMasterAccountInput) DisassociateFromMasterAccountRequest {
	op := &aws.Operation{
		Name:       opDisassociateFromMasterAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/master/disassociate",
	}

	if input == nil {
		input = &DisassociateFromMasterAccountInput{}
	}

	req := c.newRequest(op, input, &DisassociateFromMasterAccountOutput{})

	return DisassociateFromMasterAccountRequest{Request: req, Input: input, Copy: c.DisassociateFromMasterAccountRequest}
}

// DisassociateFromMasterAccountRequest is the request type for the
// DisassociateFromMasterAccount API operation.
type DisassociateFromMasterAccountRequest struct {
	*aws.Request
	Input *DisassociateFromMasterAccountInput
	Copy  func(*DisassociateFromMasterAccountInput) DisassociateFromMasterAccountRequest
}

// Send marshals and sends the DisassociateFromMasterAccount API request.
func (r DisassociateFromMasterAccountRequest) Send(ctx context.Context) (*DisassociateFromMasterAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateFromMasterAccountResponse{
		DisassociateFromMasterAccountOutput: r.Request.Data.(*DisassociateFromMasterAccountOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateFromMasterAccountResponse is the response type for the
// DisassociateFromMasterAccount API operation.
type DisassociateFromMasterAccountResponse struct {
	*DisassociateFromMasterAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateFromMasterAccount request.
func (r *DisassociateFromMasterAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
