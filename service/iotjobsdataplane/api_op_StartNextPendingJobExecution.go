// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StartNextPendingJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// A collection of name/value pairs that describe the status of the job execution.
	// If not specified, the statusDetails are unchanged.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// Specifies the amount of time this device has to finish execution of this
	// job. If the job execution status is not set to a terminal state before this
	// timer expires, or before the timer is reset (by calling UpdateJobExecution,
	// setting the status to IN_PROGRESS and specifying a new timeout value in field
	// stepTimeoutInMinutes) the job execution status will be automatically set
	// to TIMED_OUT. Note that setting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created (CreateJob
	// using field timeoutConfig).
	StepTimeoutInMinutes *int64 `locationName:"stepTimeoutInMinutes" type:"long"`

	// The name of the thing associated with the device.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartNextPendingJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartNextPendingJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartNextPendingJobExecutionInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartNextPendingJobExecutionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.StatusDetails != nil {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.StepTimeoutInMinutes != nil {
		v := *s.StepTimeoutInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stepTimeoutInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartNextPendingJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	// A JobExecution object.
	Execution *JobExecution `locationName:"execution" type:"structure"`
}

// String returns the string representation
func (s StartNextPendingJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartNextPendingJobExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Execution != nil {
		v := s.Execution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "execution", v, metadata)
	}
	return nil
}

const opStartNextPendingJobExecution = "StartNextPendingJobExecution"

// StartNextPendingJobExecutionRequest returns a request value for making API operation for
// AWS IoT Jobs Data Plane.
//
// Gets and starts the next pending (status IN_PROGRESS or QUEUED) job execution
// for a thing.
//
//    // Example sending a request using StartNextPendingJobExecutionRequest.
//    req := client.StartNextPendingJobExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/StartNextPendingJobExecution
func (c *Client) StartNextPendingJobExecutionRequest(input *StartNextPendingJobExecutionInput) StartNextPendingJobExecutionRequest {
	op := &aws.Operation{
		Name:       opStartNextPendingJobExecution,
		HTTPMethod: "PUT",
		HTTPPath:   "/things/{thingName}/jobs/$next",
	}

	if input == nil {
		input = &StartNextPendingJobExecutionInput{}
	}

	req := c.newRequest(op, input, &StartNextPendingJobExecutionOutput{})

	return StartNextPendingJobExecutionRequest{Request: req, Input: input, Copy: c.StartNextPendingJobExecutionRequest}
}

// StartNextPendingJobExecutionRequest is the request type for the
// StartNextPendingJobExecution API operation.
type StartNextPendingJobExecutionRequest struct {
	*aws.Request
	Input *StartNextPendingJobExecutionInput
	Copy  func(*StartNextPendingJobExecutionInput) StartNextPendingJobExecutionRequest
}

// Send marshals and sends the StartNextPendingJobExecution API request.
func (r StartNextPendingJobExecutionRequest) Send(ctx context.Context) (*StartNextPendingJobExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartNextPendingJobExecutionResponse{
		StartNextPendingJobExecutionOutput: r.Request.Data.(*StartNextPendingJobExecutionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartNextPendingJobExecutionResponse is the response type for the
// StartNextPendingJobExecution API operation.
type StartNextPendingJobExecutionResponse struct {
	*StartNextPendingJobExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartNextPendingJobExecution request.
func (r *StartNextPendingJobExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
