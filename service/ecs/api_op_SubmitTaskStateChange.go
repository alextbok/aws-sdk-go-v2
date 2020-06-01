// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SubmitTaskStateChangeInput struct {
	_ struct{} `type:"structure"`

	// Any attachments associated with the state change request.
	Attachments []AttachmentStateChange `locationName:"attachments" type:"list"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task.
	Cluster *string `locationName:"cluster" type:"string"`

	// Any containers associated with the state change request.
	Containers []ContainerStateChange `locationName:"containers" type:"list"`

	// The Unix timestamp for when the task execution stopped.
	ExecutionStoppedAt *time.Time `locationName:"executionStoppedAt" type:"timestamp"`

	// The Unix timestamp for when the container image pull began.
	PullStartedAt *time.Time `locationName:"pullStartedAt" type:"timestamp"`

	// The Unix timestamp for when the container image pull completed.
	PullStoppedAt *time.Time `locationName:"pullStoppedAt" type:"timestamp"`

	// The reason for the state change request.
	Reason *string `locationName:"reason" type:"string"`

	// The status of the state change request.
	Status *string `locationName:"status" type:"string"`

	// The task ID or full ARN of the task in the state change request.
	Task *string `locationName:"task" type:"string"`
}

// String returns the string representation
func (s SubmitTaskStateChangeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitTaskStateChangeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SubmitTaskStateChangeInput"}
	if s.Attachments != nil {
		for i, v := range s.Attachments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attachments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SubmitTaskStateChangeOutput struct {
	_ struct{} `type:"structure"`

	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`
}

// String returns the string representation
func (s SubmitTaskStateChangeOutput) String() string {
	return awsutil.Prettify(s)
}

const opSubmitTaskStateChange = "SubmitTaskStateChange"

// SubmitTaskStateChangeRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
//
// This action is only used by the Amazon ECS agent, and it is not intended
// for use outside of the agent.
//
// Sent to acknowledge that a task changed states.
//
//    // Example sending a request using SubmitTaskStateChangeRequest.
//    req := client.SubmitTaskStateChangeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/SubmitTaskStateChange
func (c *Client) SubmitTaskStateChangeRequest(input *SubmitTaskStateChangeInput) SubmitTaskStateChangeRequest {
	op := &aws.Operation{
		Name:       opSubmitTaskStateChange,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitTaskStateChangeInput{}
	}

	req := c.newRequest(op, input, &SubmitTaskStateChangeOutput{})

	return SubmitTaskStateChangeRequest{Request: req, Input: input, Copy: c.SubmitTaskStateChangeRequest}
}

// SubmitTaskStateChangeRequest is the request type for the
// SubmitTaskStateChange API operation.
type SubmitTaskStateChangeRequest struct {
	*aws.Request
	Input *SubmitTaskStateChangeInput
	Copy  func(*SubmitTaskStateChangeInput) SubmitTaskStateChangeRequest
}

// Send marshals and sends the SubmitTaskStateChange API request.
func (r SubmitTaskStateChangeRequest) Send(ctx context.Context) (*SubmitTaskStateChangeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SubmitTaskStateChangeResponse{
		SubmitTaskStateChangeOutput: r.Request.Data.(*SubmitTaskStateChangeOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SubmitTaskStateChangeResponse is the response type for the
// SubmitTaskStateChange API operation.
type SubmitTaskStateChangeResponse struct {
	*SubmitTaskStateChangeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SubmitTaskStateChange request.
func (r *SubmitTaskStateChangeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
