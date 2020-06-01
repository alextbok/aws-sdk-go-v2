// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendTaskFailureInput struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`

	// The token that represents this task. Task tokens are generated by Step Functions
	// when tasks are assigned to a worker, or in the context object (https://docs.aws.amazon.com/step-functions/latest/dg/input-output-contextobject.html)
	// when a workflow enters a task state. See GetActivityTaskOutput$taskToken.
	//
	// TaskToken is a required field
	TaskToken *string `locationName:"taskToken" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SendTaskFailureInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendTaskFailureInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendTaskFailureInput"}

	if s.TaskToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SendTaskFailureOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SendTaskFailureOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendTaskFailure = "SendTaskFailure"

// SendTaskFailureRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Used by activity workers and task states using the callback (https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token)
// pattern to report that the task identified by the taskToken failed.
//
//    // Example sending a request using SendTaskFailureRequest.
//    req := client.SendTaskFailureRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/SendTaskFailure
func (c *Client) SendTaskFailureRequest(input *SendTaskFailureInput) SendTaskFailureRequest {
	op := &aws.Operation{
		Name:       opSendTaskFailure,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendTaskFailureInput{}
	}

	req := c.newRequest(op, input, &SendTaskFailureOutput{})

	return SendTaskFailureRequest{Request: req, Input: input, Copy: c.SendTaskFailureRequest}
}

// SendTaskFailureRequest is the request type for the
// SendTaskFailure API operation.
type SendTaskFailureRequest struct {
	*aws.Request
	Input *SendTaskFailureInput
	Copy  func(*SendTaskFailureInput) SendTaskFailureRequest
}

// Send marshals and sends the SendTaskFailure API request.
func (r SendTaskFailureRequest) Send(ctx context.Context) (*SendTaskFailureResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendTaskFailureResponse{
		SendTaskFailureOutput: r.Request.Data.(*SendTaskFailureOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendTaskFailureResponse is the response type for the
// SendTaskFailure API operation.
type SendTaskFailureResponse struct {
	*SendTaskFailureOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendTaskFailure request.
func (r *SendTaskFailureResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
