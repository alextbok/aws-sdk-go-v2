// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeNotebookInstanceInput
type DescribeNotebookInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the notebook instance that you want information about.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeNotebookInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNotebookInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNotebookInstanceInput"}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeNotebookInstanceOutput
type DescribeNotebookInstanceOutput struct {
	_ struct{} `type:"structure"`

	// A list of the Elastic Inference (EI) instance types associated with this
	// notebook instance. Currently only one EI instance type can be associated
	// with a notebook instance. For more information, see Using Elastic Inference
	// in Amazon SageMaker (sagemaker/latest/dg/ei.html).
	AcceleratorTypes []NotebookInstanceAcceleratorType `type:"list"`

	// An array of up to three Git repositories associated with the notebook instance.
	// These can be either the names of Git repositories stored as resources in
	// your account, or the URL of Git repositories in AWS CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. These repositories are cloned at the same
	// level as the default repository of your notebook instance. For more information,
	// see Associating Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	AdditionalCodeRepositories []string `type:"list"`

	// A timestamp. Use this parameter to return the time when the notebook instance
	// was created
	CreationTime *time.Time `type:"timestamp"`

	// The Git repository associated with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in AWS CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. When you open a notebook instance, it opens
	// in the directory that contains this repository. For more information, see
	// Associating Git Repositories with Amazon SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	DefaultCodeRepository *string `min:"1" type:"string"`

	// Describes whether Amazon SageMaker provides internet access to the notebook
	// instance. If this value is set to Disabled, the notebook instance does not
	// have internet access, and cannot connect to Amazon SageMaker training and
	// endpoint services.
	//
	// For more information, see Notebook Instances Are Internet-Enabled by Default
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access).
	DirectInternetAccess DirectInternetAccess `type:"string" enum:"true"`

	// If status is Failed, the reason it failed.
	FailureReason *string `type:"string"`

	// The type of ML compute instance running on the notebook instance.
	InstanceType InstanceType `type:"string" enum:"true"`

	// The AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it
	// on the ML storage volume attached to the instance.
	KmsKeyId *string `type:"string"`

	// A timestamp. Use this parameter to retrieve the time when the notebook instance
	// was last modified.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The network interface IDs that Amazon SageMaker created at the time of creating
	// the instance.
	NetworkInterfaceId *string `type:"string"`

	// The Amazon Resource Name (ARN) of the notebook instance.
	NotebookInstanceArn *string `type:"string"`

	// Returns the name of a notebook instance lifecycle configuration.
	//
	// For information about notebook instance lifestyle configurations, see Step
	// 2.1: (Optional) Customize a Notebook Instance (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html)
	NotebookInstanceLifecycleConfigName *string `type:"string"`

	// The name of the Amazon SageMaker notebook instance.
	NotebookInstanceName *string `type:"string"`

	// The status of the notebook instance.
	NotebookInstanceStatus NotebookInstanceStatus `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the IAM role associated with the instance.
	RoleArn *string `min:"20" type:"string"`

	// Whether root access is enabled or disabled for users of the notebook instance.
	//
	// Lifecycle configurations need root access to be able to set up a notebook
	// instance. Because of this, lifecycle configurations associated with a notebook
	// instance always run with root access even if you disable root access for
	// users.
	RootAccess RootAccess `type:"string" enum:"true"`

	// The IDs of the VPC security groups.
	SecurityGroups []string `type:"list"`

	// The ID of the VPC subnet.
	SubnetId *string `type:"string"`

	// The URL that you use to connect to the Jupyter notebook that is running in
	// your notebook instance.
	Url *string `type:"string"`

	// The size, in GB, of the ML storage volume attached to the notebook instance.
	VolumeSizeInGB *int64 `min:"5" type:"integer"`
}

// String returns the string representation
func (s DescribeNotebookInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeNotebookInstance = "DescribeNotebookInstance"

// DescribeNotebookInstanceRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns information about a notebook instance.
//
//    // Example sending a request using DescribeNotebookInstanceRequest.
//    req := client.DescribeNotebookInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeNotebookInstance
func (c *Client) DescribeNotebookInstanceRequest(input *DescribeNotebookInstanceInput) DescribeNotebookInstanceRequest {
	op := &aws.Operation{
		Name:       opDescribeNotebookInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNotebookInstanceInput{}
	}

	req := c.newRequest(op, input, &DescribeNotebookInstanceOutput{})
	return DescribeNotebookInstanceRequest{Request: req, Input: input, Copy: c.DescribeNotebookInstanceRequest}
}

// DescribeNotebookInstanceRequest is the request type for the
// DescribeNotebookInstance API operation.
type DescribeNotebookInstanceRequest struct {
	*aws.Request
	Input *DescribeNotebookInstanceInput
	Copy  func(*DescribeNotebookInstanceInput) DescribeNotebookInstanceRequest
}

// Send marshals and sends the DescribeNotebookInstance API request.
func (r DescribeNotebookInstanceRequest) Send(ctx context.Context) (*DescribeNotebookInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNotebookInstanceResponse{
		DescribeNotebookInstanceOutput: r.Request.Data.(*DescribeNotebookInstanceOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeNotebookInstanceResponse is the response type for the
// DescribeNotebookInstance API operation.
type DescribeNotebookInstanceResponse struct {
	*DescribeNotebookInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNotebookInstance request.
func (r *DescribeNotebookInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
