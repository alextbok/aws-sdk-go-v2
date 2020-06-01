// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDeleteScheduledActionInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The names of the scheduled actions to delete. The maximum number allowed
	// is 50.
	//
	// ScheduledActionNames is a required field
	ScheduledActionNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDeleteScheduledActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeleteScheduledActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeleteScheduledActionInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.ScheduledActionNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledActionNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDeleteScheduledActionOutput struct {
	_ struct{} `type:"structure"`

	// The names of the scheduled actions that could not be deleted, including an
	// error message.
	FailedScheduledActions []FailedScheduledUpdateGroupActionRequest `type:"list"`
}

// String returns the string representation
func (s BatchDeleteScheduledActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDeleteScheduledAction = "BatchDeleteScheduledAction"

// BatchDeleteScheduledActionRequest returns a request value for making API operation for
// Auto Scaling.
//
// Deletes one or more scheduled actions for the specified Auto Scaling group.
//
//    // Example sending a request using BatchDeleteScheduledActionRequest.
//    req := client.BatchDeleteScheduledActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/BatchDeleteScheduledAction
func (c *Client) BatchDeleteScheduledActionRequest(input *BatchDeleteScheduledActionInput) BatchDeleteScheduledActionRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteScheduledAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeleteScheduledActionInput{}
	}

	req := c.newRequest(op, input, &BatchDeleteScheduledActionOutput{})

	return BatchDeleteScheduledActionRequest{Request: req, Input: input, Copy: c.BatchDeleteScheduledActionRequest}
}

// BatchDeleteScheduledActionRequest is the request type for the
// BatchDeleteScheduledAction API operation.
type BatchDeleteScheduledActionRequest struct {
	*aws.Request
	Input *BatchDeleteScheduledActionInput
	Copy  func(*BatchDeleteScheduledActionInput) BatchDeleteScheduledActionRequest
}

// Send marshals and sends the BatchDeleteScheduledAction API request.
func (r BatchDeleteScheduledActionRequest) Send(ctx context.Context) (*BatchDeleteScheduledActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteScheduledActionResponse{
		BatchDeleteScheduledActionOutput: r.Request.Data.(*BatchDeleteScheduledActionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteScheduledActionResponse is the response type for the
// BatchDeleteScheduledAction API operation.
type BatchDeleteScheduledActionResponse struct {
	*BatchDeleteScheduledActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteScheduledAction request.
func (r *BatchDeleteScheduledActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
