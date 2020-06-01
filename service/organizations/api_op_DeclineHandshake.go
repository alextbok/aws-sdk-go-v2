// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeclineHandshakeInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the handshake that you want to decline. You
	// can get the ID from the ListHandshakesForAccount operation.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string
	// requires "h-" followed by from 8 to 32 lowercase letters or digits.
	//
	// HandshakeId is a required field
	HandshakeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeclineHandshakeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeclineHandshakeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeclineHandshakeInput"}

	if s.HandshakeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HandshakeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeclineHandshakeOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the declined handshake. The state
	// is updated to show the value DECLINED.
	Handshake *Handshake `type:"structure"`
}

// String returns the string representation
func (s DeclineHandshakeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeclineHandshake = "DeclineHandshake"

// DeclineHandshakeRequest returns a request value for making API operation for
// AWS Organizations.
//
// Declines a handshake request. This sets the handshake state to DECLINED and
// effectively deactivates the request.
//
// This operation can be called only from the account that received the handshake.
// The originator of the handshake can use CancelHandshake instead. The originator
// can't reactivate a declined request, but can reinitiate the process with
// a new handshake request.
//
// After you decline a handshake, it continues to appear in the results of relevant
// APIs for only 30 days. After that, it's deleted.
//
//    // Example sending a request using DeclineHandshakeRequest.
//    req := client.DeclineHandshakeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DeclineHandshake
func (c *Client) DeclineHandshakeRequest(input *DeclineHandshakeInput) DeclineHandshakeRequest {
	op := &aws.Operation{
		Name:       opDeclineHandshake,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeclineHandshakeInput{}
	}

	req := c.newRequest(op, input, &DeclineHandshakeOutput{})

	return DeclineHandshakeRequest{Request: req, Input: input, Copy: c.DeclineHandshakeRequest}
}

// DeclineHandshakeRequest is the request type for the
// DeclineHandshake API operation.
type DeclineHandshakeRequest struct {
	*aws.Request
	Input *DeclineHandshakeInput
	Copy  func(*DeclineHandshakeInput) DeclineHandshakeRequest
}

// Send marshals and sends the DeclineHandshake API request.
func (r DeclineHandshakeRequest) Send(ctx context.Context) (*DeclineHandshakeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeclineHandshakeResponse{
		DeclineHandshakeOutput: r.Request.Data.(*DeclineHandshakeOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeclineHandshakeResponse is the response type for the
// DeclineHandshake API operation.
type DeclineHandshakeResponse struct {
	*DeclineHandshakeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeclineHandshake request.
func (r *DeclineHandshakeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
