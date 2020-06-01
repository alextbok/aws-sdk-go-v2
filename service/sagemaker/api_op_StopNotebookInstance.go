// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StopNotebookInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the notebook instance to terminate.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopNotebookInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopNotebookInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopNotebookInstanceInput"}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopNotebookInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopNotebookInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopNotebookInstance = "StopNotebookInstance"

// StopNotebookInstanceRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Terminates the ML compute instance. Before terminating the instance, Amazon
// SageMaker disconnects the ML storage volume from it. Amazon SageMaker preserves
// the ML storage volume. Amazon SageMaker stops charging you for the ML compute
// instance when you call StopNotebookInstance.
//
// To access data on the ML storage volume for a notebook instance that has
// been terminated, call the StartNotebookInstance API. StartNotebookInstance
// launches another ML compute instance, configures it, and attaches the preserved
// ML storage volume so you can continue your work.
//
//    // Example sending a request using StopNotebookInstanceRequest.
//    req := client.StopNotebookInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopNotebookInstance
func (c *Client) StopNotebookInstanceRequest(input *StopNotebookInstanceInput) StopNotebookInstanceRequest {
	op := &aws.Operation{
		Name:       opStopNotebookInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopNotebookInstanceInput{}
	}

	req := c.newRequest(op, input, &StopNotebookInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return StopNotebookInstanceRequest{Request: req, Input: input, Copy: c.StopNotebookInstanceRequest}
}

// StopNotebookInstanceRequest is the request type for the
// StopNotebookInstance API operation.
type StopNotebookInstanceRequest struct {
	*aws.Request
	Input *StopNotebookInstanceInput
	Copy  func(*StopNotebookInstanceInput) StopNotebookInstanceRequest
}

// Send marshals and sends the StopNotebookInstance API request.
func (r StopNotebookInstanceRequest) Send(ctx context.Context) (*StopNotebookInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopNotebookInstanceResponse{
		StopNotebookInstanceOutput: r.Request.Data.(*StopNotebookInstanceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopNotebookInstanceResponse is the response type for the
// StopNotebookInstance API operation.
type StopNotebookInstanceResponse struct {
	*StopNotebookInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopNotebookInstance request.
func (r *StopNotebookInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
