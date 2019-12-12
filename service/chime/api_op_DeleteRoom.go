// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteRoomInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The chat room ID.
	//
	// RoomId is a required field
	RoomId *string `location:"uri" locationName:"roomId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRoomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRoomInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.RoomId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRoomInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoomId != nil {
		v := *s.RoomId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "roomId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteRoomOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRoomOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRoomOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRoom = "DeleteRoom"

// DeleteRoomRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes a chat room.
//
//    // Example sending a request using DeleteRoomRequest.
//    req := client.DeleteRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteRoom
func (c *Client) DeleteRoomRequest(input *DeleteRoomInput) DeleteRoomRequest {
	op := &aws.Operation{
		Name:       opDeleteRoom,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{accountId}/rooms/{roomId}",
	}

	if input == nil {
		input = &DeleteRoomInput{}
	}

	req := c.newRequest(op, input, &DeleteRoomOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRoomRequest{Request: req, Input: input, Copy: c.DeleteRoomRequest}
}

// DeleteRoomRequest is the request type for the
// DeleteRoom API operation.
type DeleteRoomRequest struct {
	*aws.Request
	Input *DeleteRoomInput
	Copy  func(*DeleteRoomInput) DeleteRoomRequest
}

// Send marshals and sends the DeleteRoom API request.
func (r DeleteRoomRequest) Send(ctx context.Context) (*DeleteRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRoomResponse{
		DeleteRoomOutput: r.Request.Data.(*DeleteRoomOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRoomResponse is the response type for the
// DeleteRoom API operation.
type DeleteRoomResponse struct {
	*DeleteRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoom request.
func (r *DeleteRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}