// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateTrialComponentInput struct {
	_ struct{} `type:"structure"`

	// The name of the component as displayed. The name doesn't need to be unique.
	// If DisplayName isn't specified, TrialComponentName is displayed.
	DisplayName *string `min:"1" type:"string"`

	// When the component ended.
	EndTime *time.Time `type:"timestamp"`

	// Replaces all of the component's input artifacts with the specified artifacts.
	InputArtifacts map[string]TrialComponentArtifact `type:"map"`

	// The input artifacts to remove from the component.
	InputArtifactsToRemove []string `type:"list"`

	// Replaces all of the component's output artifacts with the specified artifacts.
	OutputArtifacts map[string]TrialComponentArtifact `type:"map"`

	// The output artifacts to remove from the component.
	OutputArtifactsToRemove []string `type:"list"`

	// Replaces all of the component's hyperparameters with the specified hyperparameters.
	Parameters map[string]TrialComponentParameterValue `type:"map"`

	// The hyperparameters to remove from the component.
	ParametersToRemove []string `type:"list"`

	// When the component started.
	StartTime *time.Time `type:"timestamp"`

	// The new status of the component.
	Status *TrialComponentStatus `type:"structure"`

	// The name of the component to update.
	//
	// TrialComponentName is a required field
	TrialComponentName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTrialComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTrialComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTrialComponentInput"}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

	if s.TrialComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrialComponentName"))
	}
	if s.TrialComponentName != nil && len(*s.TrialComponentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrialComponentName", 1))
	}
	if s.InputArtifacts != nil {
		for i, v := range s.InputArtifacts {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InputArtifacts", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.OutputArtifacts != nil {
		for i, v := range s.OutputArtifacts {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OutputArtifacts", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateTrialComponentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string `type:"string"`
}

// String returns the string representation
func (s UpdateTrialComponentOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTrialComponent = "UpdateTrialComponent"

// UpdateTrialComponentRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Updates one or more properties of a trial component.
//
//    // Example sending a request using UpdateTrialComponentRequest.
//    req := client.UpdateTrialComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateTrialComponent
func (c *Client) UpdateTrialComponentRequest(input *UpdateTrialComponentInput) UpdateTrialComponentRequest {
	op := &aws.Operation{
		Name:       opUpdateTrialComponent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTrialComponentInput{}
	}

	req := c.newRequest(op, input, &UpdateTrialComponentOutput{})

	return UpdateTrialComponentRequest{Request: req, Input: input, Copy: c.UpdateTrialComponentRequest}
}

// UpdateTrialComponentRequest is the request type for the
// UpdateTrialComponent API operation.
type UpdateTrialComponentRequest struct {
	*aws.Request
	Input *UpdateTrialComponentInput
	Copy  func(*UpdateTrialComponentInput) UpdateTrialComponentRequest
}

// Send marshals and sends the UpdateTrialComponent API request.
func (r UpdateTrialComponentRequest) Send(ctx context.Context) (*UpdateTrialComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTrialComponentResponse{
		UpdateTrialComponentOutput: r.Request.Data.(*UpdateTrialComponentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTrialComponentResponse is the response type for the
// UpdateTrialComponent API operation.
type UpdateTrialComponentResponse struct {
	*UpdateTrialComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrialComponent request.
func (r *UpdateTrialComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
