// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateNotebookInstanceInput struct {
	_ struct{} `type:"structure"`

	// A list of the Elastic Inference (EI) instance types to associate with this
	// notebook instance. Currently only one EI instance type can be associated
	// with a notebook instance. For more information, see Using Elastic Inference
	// in Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html).
	AcceleratorTypes []NotebookInstanceAcceleratorType `type:"list"`

	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in
	// your account, or the URL of Git repositories in AWS CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. These repositories are cloned at the same
	// level as the default repository of your notebook instance. For more information,
	// see Associating Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	AdditionalCodeRepositories []string `type:"list"`

	// The Git repository to associate with the notebook instance as its default
	// code repository. This can be either the name of a Git repository stored as
	// a resource in your account, or the URL of a Git repository in AWS CodeCommit
	// (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or
	// in any other Git repository. When you open a notebook instance, it opens
	// in the directory that contains this repository. For more information, see
	// Associating Git Repositories with Amazon SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	DefaultCodeRepository *string `min:"1" type:"string"`

	// A list of the Elastic Inference (EI) instance types to remove from this notebook
	// instance. This operation is idempotent. If you specify an accelerator type
	// that is not associated with the notebook instance when you call this method,
	// it does not throw an error.
	DisassociateAcceleratorTypes *bool `type:"boolean"`

	// A list of names or URLs of the default Git repositories to remove from this
	// notebook instance. This operation is idempotent. If you specify a Git repository
	// that is not associated with the notebook instance when you call this method,
	// it does not throw an error.
	DisassociateAdditionalCodeRepositories *bool `type:"boolean"`

	// The name or URL of the default Git repository to remove from this notebook
	// instance. This operation is idempotent. If you specify a Git repository that
	// is not associated with the notebook instance when you call this method, it
	// does not throw an error.
	DisassociateDefaultCodeRepository *bool `type:"boolean"`

	// Set to true to remove the notebook instance lifecycle configuration currently
	// associated with the notebook instance. This operation is idempotent. If you
	// specify a lifecycle configuration that is not associated with the notebook
	// instance when you call this method, it does not throw an error.
	DisassociateLifecycleConfig *bool `type:"boolean"`

	// The Amazon ML compute instance type.
	InstanceType InstanceType `type:"string" enum:"true"`

	// The name of a lifecycle configuration to associate with the notebook instance.
	// For information about lifestyle configurations, see Step 2.1: (Optional)
	// Customize a Notebook Instance (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html).
	LifecycleConfigName *string `type:"string"`

	// The name of the notebook instance to update.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role that Amazon SageMaker can
	// assume to access the notebook instance. For more information, see Amazon
	// SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).
	//
	// To be able to pass this role to Amazon SageMaker, the caller of this API
	// must have the iam:PassRole permission.
	RoleArn *string `min:"20" type:"string"`

	// Whether root access is enabled or disabled for users of the notebook instance.
	// The default value is Enabled.
	//
	// If you set this to Disabled, users don't have root access on the notebook
	// instance, but lifecycle configuration scripts still run with root permissions.
	RootAccess RootAccess `type:"string" enum:"true"`

	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	// The default value is 5 GB. ML storage volumes are encrypted, so Amazon SageMaker
	// can't determine the amount of available free space on the volume. Because
	// of this, you can increase the volume size when you update a notebook instance,
	// but you can't decrease the volume size. If you want to decrease the size
	// of the ML storage volume in use, create a new notebook instance with the
	// desired size.
	VolumeSizeInGB *int64 `min:"5" type:"integer"`
}

// String returns the string representation
func (s UpdateNotebookInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNotebookInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNotebookInstanceInput"}
	if s.DefaultCodeRepository != nil && len(*s.DefaultCodeRepository) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultCodeRepository", 1))
	}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}
	if s.VolumeSizeInGB != nil && *s.VolumeSizeInGB < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("VolumeSizeInGB", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateNotebookInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateNotebookInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateNotebookInstance = "UpdateNotebookInstance"

// UpdateNotebookInstanceRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Updates a notebook instance. NotebookInstance updates include upgrading or
// downgrading the ML compute instance used for your notebook instance to accommodate
// changes in your workload requirements.
//
//    // Example sending a request using UpdateNotebookInstanceRequest.
//    req := client.UpdateNotebookInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateNotebookInstance
func (c *Client) UpdateNotebookInstanceRequest(input *UpdateNotebookInstanceInput) UpdateNotebookInstanceRequest {
	op := &aws.Operation{
		Name:       opUpdateNotebookInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNotebookInstanceInput{}
	}

	req := c.newRequest(op, input, &UpdateNotebookInstanceOutput{})

	return UpdateNotebookInstanceRequest{Request: req, Input: input, Copy: c.UpdateNotebookInstanceRequest}
}

// UpdateNotebookInstanceRequest is the request type for the
// UpdateNotebookInstance API operation.
type UpdateNotebookInstanceRequest struct {
	*aws.Request
	Input *UpdateNotebookInstanceInput
	Copy  func(*UpdateNotebookInstanceInput) UpdateNotebookInstanceRequest
}

// Send marshals and sends the UpdateNotebookInstance API request.
func (r UpdateNotebookInstanceRequest) Send(ctx context.Context) (*UpdateNotebookInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNotebookInstanceResponse{
		UpdateNotebookInstanceOutput: r.Request.Data.(*UpdateNotebookInstanceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNotebookInstanceResponse is the response type for the
// UpdateNotebookInstance API operation.
type UpdateNotebookInstanceResponse struct {
	*UpdateNotebookInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNotebookInstance request.
func (r *UpdateNotebookInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
