// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateMembersInput struct {
	_ struct{} `type:"structure"`

	// The account IDs of the member accounts to disassociate from the master account.
	AccountIds []string `type:"list"`
}

// String returns the string representation
func (s DisassociateMembersInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateMembersInput) MarshalFields(e protocol.FieldEncoder) error {
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

type DisassociateMembersOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisassociateMembers = "DisassociateMembers"

// DisassociateMembersRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Disassociates the specified member accounts from the associated master account.
//
//    // Example sending a request using DisassociateMembersRequest.
//    req := client.DisassociateMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DisassociateMembers
func (c *Client) DisassociateMembersRequest(input *DisassociateMembersInput) DisassociateMembersRequest {
	op := &aws.Operation{
		Name:       opDisassociateMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/members/disassociate",
	}

	if input == nil {
		input = &DisassociateMembersInput{}
	}

	req := c.newRequest(op, input, &DisassociateMembersOutput{})

	return DisassociateMembersRequest{Request: req, Input: input, Copy: c.DisassociateMembersRequest}
}

// DisassociateMembersRequest is the request type for the
// DisassociateMembers API operation.
type DisassociateMembersRequest struct {
	*aws.Request
	Input *DisassociateMembersInput
	Copy  func(*DisassociateMembersInput) DisassociateMembersRequest
}

// Send marshals and sends the DisassociateMembers API request.
func (r DisassociateMembersRequest) Send(ctx context.Context) (*DisassociateMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateMembersResponse{
		DisassociateMembersOutput: r.Request.Data.(*DisassociateMembersOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateMembersResponse is the response type for the
// DisassociateMembers API operation.
type DisassociateMembersResponse struct {
	*DisassociateMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateMembers request.
func (r *DisassociateMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
