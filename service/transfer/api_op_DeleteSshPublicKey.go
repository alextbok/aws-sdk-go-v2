// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteSshPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for a file transfer protocol-enabled
	// server instance that has the user assigned to it.
	//
	// ServerId is a required field
	ServerId *string `min:"19" type:"string" required:"true"`

	// A unique identifier used to reference your user's specific SSH key.
	//
	// SshPublicKeyId is a required field
	SshPublicKeyId *string `min:"21" type:"string" required:"true"`

	// A unique string that identifies a user whose public key is being deleted.
	//
	// UserName is a required field
	UserName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSshPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSshPublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSshPublicKeyInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}
	if s.ServerId != nil && len(*s.ServerId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerId", 19))
	}

	if s.SshPublicKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SshPublicKeyId"))
	}
	if s.SshPublicKeyId != nil && len(*s.SshPublicKeyId) < 21 {
		invalidParams.Add(aws.NewErrParamMinLen("SshPublicKeyId", 21))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSshPublicKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSshPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSshPublicKey = "DeleteSshPublicKey"

// DeleteSshPublicKeyRequest returns a request value for making API operation for
// AWS Transfer Family.
//
// Deletes a user's Secure Shell (SSH) public key.
//
// No response is returned from this operation.
//
//    // Example sending a request using DeleteSshPublicKeyRequest.
//    req := client.DeleteSshPublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DeleteSshPublicKey
func (c *Client) DeleteSshPublicKeyRequest(input *DeleteSshPublicKeyInput) DeleteSshPublicKeyRequest {
	op := &aws.Operation{
		Name:       opDeleteSshPublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSshPublicKeyInput{}
	}

	req := c.newRequest(op, input, &DeleteSshPublicKeyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteSshPublicKeyRequest{Request: req, Input: input, Copy: c.DeleteSshPublicKeyRequest}
}

// DeleteSshPublicKeyRequest is the request type for the
// DeleteSshPublicKey API operation.
type DeleteSshPublicKeyRequest struct {
	*aws.Request
	Input *DeleteSshPublicKeyInput
	Copy  func(*DeleteSshPublicKeyInput) DeleteSshPublicKeyRequest
}

// Send marshals and sends the DeleteSshPublicKey API request.
func (r DeleteSshPublicKeyRequest) Send(ctx context.Context) (*DeleteSshPublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSshPublicKeyResponse{
		DeleteSshPublicKeyOutput: r.Request.Data.(*DeleteSshPublicKeyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSshPublicKeyResponse is the response type for the
// DeleteSshPublicKey API operation.
type DeleteSshPublicKeyResponse struct {
	*DeleteSshPublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSshPublicKey request.
func (r *DeleteSshPublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
