// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopAccessLoggingInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that you want to stop access logging on.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopAccessLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopAccessLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopAccessLoggingInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopAccessLoggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopAccessLoggingOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopAccessLogging = "StopAccessLogging"

// StopAccessLoggingRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Stops access logging on the specified container. When you stop access logging
// on a container, MediaStore stops sending access logs to Amazon CloudWatch
// Logs. These access logs are not saved and are not retrievable.
//
//    // Example sending a request using StopAccessLoggingRequest.
//    req := client.StopAccessLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/StopAccessLogging
func (c *Client) StopAccessLoggingRequest(input *StopAccessLoggingInput) StopAccessLoggingRequest {
	op := &aws.Operation{
		Name:       opStopAccessLogging,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopAccessLoggingInput{}
	}

	req := c.newRequest(op, input, &StopAccessLoggingOutput{})

	return StopAccessLoggingRequest{Request: req, Input: input, Copy: c.StopAccessLoggingRequest}
}

// StopAccessLoggingRequest is the request type for the
// StopAccessLogging API operation.
type StopAccessLoggingRequest struct {
	*aws.Request
	Input *StopAccessLoggingInput
	Copy  func(*StopAccessLoggingInput) StopAccessLoggingRequest
}

// Send marshals and sends the StopAccessLogging API request.
func (r StopAccessLoggingRequest) Send(ctx context.Context) (*StopAccessLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopAccessLoggingResponse{
		StopAccessLoggingOutput: r.Request.Data.(*StopAccessLoggingOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopAccessLoggingResponse is the response type for the
// StopAccessLogging API operation.
type StopAccessLoggingResponse struct {
	*StopAccessLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopAccessLogging request.
func (r *StopAccessLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
