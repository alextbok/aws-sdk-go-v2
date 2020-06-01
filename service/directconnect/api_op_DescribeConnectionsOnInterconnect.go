// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConnectionsOnInterconnectInput struct {
	_ struct{} `type:"structure"`

	// The ID of the interconnect.
	//
	// InterconnectId is a required field
	InterconnectId *string `locationName:"interconnectId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeConnectionsOnInterconnectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConnectionsOnInterconnectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConnectionsOnInterconnectInput"}

	if s.InterconnectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InterconnectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeConnectionsOnInterconnectOutput struct {
	_ struct{} `type:"structure"`

	// The connections.
	Connections []Connection `locationName:"connections" type:"list"`
}

// String returns the string representation
func (s DescribeConnectionsOnInterconnectOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConnectionsOnInterconnect = "DescribeConnectionsOnInterconnect"

// DescribeConnectionsOnInterconnectRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deprecated. Use DescribeHostedConnections instead.
//
// Lists the connections that have been provisioned on the specified interconnect.
//
// Intended for use by AWS Direct Connect Partners only.
//
//    // Example sending a request using DescribeConnectionsOnInterconnectRequest.
//    req := client.DescribeConnectionsOnInterconnectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeConnectionsOnInterconnect
func (c *Client) DescribeConnectionsOnInterconnectRequest(input *DescribeConnectionsOnInterconnectInput) DescribeConnectionsOnInterconnectRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, DescribeConnectionsOnInterconnect, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opDescribeConnectionsOnInterconnect,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConnectionsOnInterconnectInput{}
	}

	req := c.newRequest(op, input, &DescribeConnectionsOnInterconnectOutput{})

	return DescribeConnectionsOnInterconnectRequest{Request: req, Input: input, Copy: c.DescribeConnectionsOnInterconnectRequest}
}

// DescribeConnectionsOnInterconnectRequest is the request type for the
// DescribeConnectionsOnInterconnect API operation.
type DescribeConnectionsOnInterconnectRequest struct {
	*aws.Request
	Input *DescribeConnectionsOnInterconnectInput
	Copy  func(*DescribeConnectionsOnInterconnectInput) DescribeConnectionsOnInterconnectRequest
}

// Send marshals and sends the DescribeConnectionsOnInterconnect API request.
func (r DescribeConnectionsOnInterconnectRequest) Send(ctx context.Context) (*DescribeConnectionsOnInterconnectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConnectionsOnInterconnectResponse{
		DescribeConnectionsOnInterconnectOutput: r.Request.Data.(*DescribeConnectionsOnInterconnectOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConnectionsOnInterconnectResponse is the response type for the
// DescribeConnectionsOnInterconnect API operation.
type DescribeConnectionsOnInterconnectResponse struct {
	*DescribeConnectionsOnInterconnectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConnectionsOnInterconnect request.
func (r *DescribeConnectionsOnInterconnectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
