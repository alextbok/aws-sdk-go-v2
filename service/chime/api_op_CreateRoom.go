// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateRoomInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The idempotency token for the request.
	ClientRequestToken *string `min:"2" type:"string" idempotencyToken:"true" sensitive:"true"`

	// The room name.
	//
	// Name is a required field
	Name *string `type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s CreateRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRoomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRoomInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 2))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRoomInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateRoomOutput struct {
	_ struct{} `type:"structure"`

	// The room details.
	Room *Room `type:"structure"`
}

// String returns the string representation
func (s CreateRoomOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRoomOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Room != nil {
		v := s.Room

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Room", v, metadata)
	}
	return nil
}

const opCreateRoom = "CreateRoom"

// CreateRoomRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates a chat room for the specified Amazon Chime Enterprise account.
//
//    // Example sending a request using CreateRoomRequest.
//    req := client.CreateRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreateRoom
func (c *Client) CreateRoomRequest(input *CreateRoomInput) CreateRoomRequest {
	op := &aws.Operation{
		Name:       opCreateRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/rooms",
	}

	if input == nil {
		input = &CreateRoomInput{}
	}

	req := c.newRequest(op, input, &CreateRoomOutput{})

	return CreateRoomRequest{Request: req, Input: input, Copy: c.CreateRoomRequest}
}

// CreateRoomRequest is the request type for the
// CreateRoom API operation.
type CreateRoomRequest struct {
	*aws.Request
	Input *CreateRoomInput
	Copy  func(*CreateRoomInput) CreateRoomRequest
}

// Send marshals and sends the CreateRoom API request.
func (r CreateRoomRequest) Send(ctx context.Context) (*CreateRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRoomResponse{
		CreateRoomOutput: r.Request.Data.(*CreateRoomOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRoomResponse is the response type for the
// CreateRoom API operation.
type CreateRoomResponse struct {
	*CreateRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRoom request.
func (r *CreateRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
