// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListInvitationsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListInvitationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInvitationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListInvitationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListInvitationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about all the Amazon Macie membership invitations that
// were received by an account.
type ListInvitationsOutput struct {
	_ struct{} `type:"structure"`

	Invitations []Invitation `locationName:"invitations" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListInvitationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListInvitationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Invitations != nil {
		v := s.Invitations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "invitations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListInvitations = "ListInvitations"

// ListInvitationsRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves information about all the Amazon Macie membership invitations that
// were received by an account.
//
//    // Example sending a request using ListInvitationsRequest.
//    req := client.ListInvitationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/ListInvitations
func (c *Client) ListInvitationsRequest(input *ListInvitationsInput) ListInvitationsRequest {
	op := &aws.Operation{
		Name:       opListInvitations,
		HTTPMethod: "GET",
		HTTPPath:   "/invitations",
	}

	if input == nil {
		input = &ListInvitationsInput{}
	}

	req := c.newRequest(op, input, &ListInvitationsOutput{})

	return ListInvitationsRequest{Request: req, Input: input, Copy: c.ListInvitationsRequest}
}

// ListInvitationsRequest is the request type for the
// ListInvitations API operation.
type ListInvitationsRequest struct {
	*aws.Request
	Input *ListInvitationsInput
	Copy  func(*ListInvitationsInput) ListInvitationsRequest
}

// Send marshals and sends the ListInvitations API request.
func (r ListInvitationsRequest) Send(ctx context.Context) (*ListInvitationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInvitationsResponse{
		ListInvitationsOutput: r.Request.Data.(*ListInvitationsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListInvitationsResponse is the response type for the
// ListInvitations API operation.
type ListInvitationsResponse struct {
	*ListInvitationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInvitations request.
func (r *ListInvitationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
