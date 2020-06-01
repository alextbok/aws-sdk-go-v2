// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DeleteTapeInput
type DeleteTapeInput struct {
	_ struct{} `type:"structure"`

	// The unique Amazon Resource Name (ARN) of the gateway that the virtual tape
	// to delete is associated with. Use the ListGateways operation to return a
	// list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the virtual tape to delete.
	//
	// TapeARN is a required field
	TapeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTapeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTapeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTapeInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if s.TapeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TapeARN"))
	}
	if s.TapeARN != nil && len(*s.TapeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TapeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DeleteTapeOutput
type DeleteTapeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deleted virtual tape.
	TapeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeleteTapeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTape = "DeleteTape"

// DeleteTapeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes the specified virtual tape. This operation is only supported in the
// tape gateway type.
//
//    // Example sending a request using DeleteTapeRequest.
//    req := client.DeleteTapeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteTape
func (c *Client) DeleteTapeRequest(input *DeleteTapeInput) DeleteTapeRequest {
	op := &aws.Operation{
		Name:       opDeleteTape,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTapeInput{}
	}

	req := c.newRequest(op, input, &DeleteTapeOutput{})

	return DeleteTapeRequest{Request: req, Input: input, Copy: c.DeleteTapeRequest}
}

// DeleteTapeRequest is the request type for the
// DeleteTape API operation.
type DeleteTapeRequest struct {
	*aws.Request
	Input *DeleteTapeInput
	Copy  func(*DeleteTapeInput) DeleteTapeRequest
}

// Send marshals and sends the DeleteTape API request.
func (r DeleteTapeRequest) Send(ctx context.Context) (*DeleteTapeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTapeResponse{
		DeleteTapeOutput: r.Request.Data.(*DeleteTapeOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTapeResponse is the response type for the
// DeleteTape API operation.
type DeleteTapeResponse struct {
	*DeleteTapeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTape request.
func (r *DeleteTapeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
