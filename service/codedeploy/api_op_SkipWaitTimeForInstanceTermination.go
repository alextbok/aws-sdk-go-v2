// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type SkipWaitTimeForInstanceTerminationInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a blue/green deployment for which you want to skip the instance
	// termination wait time.
	DeploymentId *string `locationName:"deploymentId" type:"string"`
}

// String returns the string representation
func (s SkipWaitTimeForInstanceTerminationInput) String() string {
	return awsutil.Prettify(s)
}

type SkipWaitTimeForInstanceTerminationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SkipWaitTimeForInstanceTerminationOutput) String() string {
	return awsutil.Prettify(s)
}

const opSkipWaitTimeForInstanceTermination = "SkipWaitTimeForInstanceTermination"

// SkipWaitTimeForInstanceTerminationRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// In a blue/green deployment, overrides any specified wait time and starts
// terminating instances immediately after the traffic routing is complete.
//
//    // Example sending a request using SkipWaitTimeForInstanceTerminationRequest.
//    req := client.SkipWaitTimeForInstanceTerminationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/SkipWaitTimeForInstanceTermination
func (c *Client) SkipWaitTimeForInstanceTerminationRequest(input *SkipWaitTimeForInstanceTerminationInput) SkipWaitTimeForInstanceTerminationRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, SkipWaitTimeForInstanceTermination, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opSkipWaitTimeForInstanceTermination,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SkipWaitTimeForInstanceTerminationInput{}
	}

	req := c.newRequest(op, input, &SkipWaitTimeForInstanceTerminationOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return SkipWaitTimeForInstanceTerminationRequest{Request: req, Input: input, Copy: c.SkipWaitTimeForInstanceTerminationRequest}
}

// SkipWaitTimeForInstanceTerminationRequest is the request type for the
// SkipWaitTimeForInstanceTermination API operation.
type SkipWaitTimeForInstanceTerminationRequest struct {
	*aws.Request
	Input *SkipWaitTimeForInstanceTerminationInput
	Copy  func(*SkipWaitTimeForInstanceTerminationInput) SkipWaitTimeForInstanceTerminationRequest
}

// Send marshals and sends the SkipWaitTimeForInstanceTermination API request.
func (r SkipWaitTimeForInstanceTerminationRequest) Send(ctx context.Context) (*SkipWaitTimeForInstanceTerminationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SkipWaitTimeForInstanceTerminationResponse{
		SkipWaitTimeForInstanceTerminationOutput: r.Request.Data.(*SkipWaitTimeForInstanceTerminationOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SkipWaitTimeForInstanceTerminationResponse is the response type for the
// SkipWaitTimeForInstanceTermination API operation.
type SkipWaitTimeForInstanceTerminationResponse struct {
	*SkipWaitTimeForInstanceTerminationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SkipWaitTimeForInstanceTermination request.
func (r *SkipWaitTimeForInstanceTerminationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
