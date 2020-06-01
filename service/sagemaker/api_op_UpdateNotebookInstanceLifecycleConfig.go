// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateNotebookInstanceLifecycleConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the lifecycle configuration.
	//
	// NotebookInstanceLifecycleConfigName is a required field
	NotebookInstanceLifecycleConfigName *string `type:"string" required:"true"`

	// The shell script that runs only once, when you create a notebook instance.
	// The shell script must be a base64-encoded string.
	OnCreate []NotebookInstanceLifecycleHook `type:"list"`

	// The shell script that runs every time you start a notebook instance, including
	// when you create the notebook instance. The shell script must be a base64-encoded
	// string.
	OnStart []NotebookInstanceLifecycleHook `type:"list"`
}

// String returns the string representation
func (s UpdateNotebookInstanceLifecycleConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNotebookInstanceLifecycleConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNotebookInstanceLifecycleConfigInput"}

	if s.NotebookInstanceLifecycleConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceLifecycleConfigName"))
	}
	if s.OnCreate != nil {
		for i, v := range s.OnCreate {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OnCreate", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.OnStart != nil {
		for i, v := range s.OnStart {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OnStart", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateNotebookInstanceLifecycleConfigOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateNotebookInstanceLifecycleConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateNotebookInstanceLifecycleConfig = "UpdateNotebookInstanceLifecycleConfig"

// UpdateNotebookInstanceLifecycleConfigRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Updates a notebook instance lifecycle configuration created with the CreateNotebookInstanceLifecycleConfig
// API.
//
//    // Example sending a request using UpdateNotebookInstanceLifecycleConfigRequest.
//    req := client.UpdateNotebookInstanceLifecycleConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateNotebookInstanceLifecycleConfig
func (c *Client) UpdateNotebookInstanceLifecycleConfigRequest(input *UpdateNotebookInstanceLifecycleConfigInput) UpdateNotebookInstanceLifecycleConfigRequest {
	op := &aws.Operation{
		Name:       opUpdateNotebookInstanceLifecycleConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNotebookInstanceLifecycleConfigInput{}
	}

	req := c.newRequest(op, input, &UpdateNotebookInstanceLifecycleConfigOutput{})

	return UpdateNotebookInstanceLifecycleConfigRequest{Request: req, Input: input, Copy: c.UpdateNotebookInstanceLifecycleConfigRequest}
}

// UpdateNotebookInstanceLifecycleConfigRequest is the request type for the
// UpdateNotebookInstanceLifecycleConfig API operation.
type UpdateNotebookInstanceLifecycleConfigRequest struct {
	*aws.Request
	Input *UpdateNotebookInstanceLifecycleConfigInput
	Copy  func(*UpdateNotebookInstanceLifecycleConfigInput) UpdateNotebookInstanceLifecycleConfigRequest
}

// Send marshals and sends the UpdateNotebookInstanceLifecycleConfig API request.
func (r UpdateNotebookInstanceLifecycleConfigRequest) Send(ctx context.Context) (*UpdateNotebookInstanceLifecycleConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNotebookInstanceLifecycleConfigResponse{
		UpdateNotebookInstanceLifecycleConfigOutput: r.Request.Data.(*UpdateNotebookInstanceLifecycleConfigOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNotebookInstanceLifecycleConfigResponse is the response type for the
// UpdateNotebookInstanceLifecycleConfig API operation.
type UpdateNotebookInstanceLifecycleConfigResponse struct {
	*UpdateNotebookInstanceLifecycleConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNotebookInstanceLifecycleConfig request.
func (r *UpdateNotebookInstanceLifecycleConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
