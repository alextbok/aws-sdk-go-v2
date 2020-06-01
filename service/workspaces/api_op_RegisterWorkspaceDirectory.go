// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterWorkspaceDirectoryInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory. You cannot register a directory if it does
	// not have a status of Active. If the directory does not have a status of Active,
	// you will receive an InvalidResourceStateException error. If you have already
	// registered the maximum number of directories that you can register with Amazon
	// WorkSpaces, you will receive a ResourceLimitExceededException error. Deregister
	// directories that you are not using for WorkSpaces, and try again.
	//
	// DirectoryId is a required field
	DirectoryId *string `min:"10" type:"string" required:"true"`

	// Indicates whether self-service capabilities are enabled or disabled.
	EnableSelfService *bool `type:"boolean"`

	// Indicates whether Amazon WorkDocs is enabled or disabled. If you have enabled
	// this parameter and WorkDocs is not available in the Region, you will receive
	// an OperationNotSupportedException error. Set EnableWorkDocs to disabled,
	// and try again.
	//
	// EnableWorkDocs is a required field
	EnableWorkDocs *bool `type:"boolean" required:"true"`

	// The identifiers of the subnets for your virtual private cloud (VPC). Make
	// sure that the subnets are in supported Availability Zones. The subnets must
	// also be in separate Availability Zones. If these conditions are not met,
	// you will receive an OperationNotSupportedException error.
	SubnetIds []string `type:"list"`

	// The tags associated with the directory.
	Tags []Tag `type:"list"`

	// Indicates whether your WorkSpace directory is dedicated or shared. To use
	// Bring Your Own License (BYOL) images, this value must be set to DEDICATED
	// and your AWS account must be enabled for BYOL. If your account has not been
	// enabled for BYOL, you will receive an InvalidParameterValuesException error.
	// For more information about BYOL images, see Bring Your Own Windows Desktop
	// Images (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html).
	Tenancy Tenancy `type:"string" enum:"true"`
}

// String returns the string representation
func (s RegisterWorkspaceDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterWorkspaceDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterWorkspaceDirectoryInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}
	if s.DirectoryId != nil && len(*s.DirectoryId) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("DirectoryId", 10))
	}

	if s.EnableWorkDocs == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnableWorkDocs"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterWorkspaceDirectoryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterWorkspaceDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterWorkspaceDirectory = "RegisterWorkspaceDirectory"

// RegisterWorkspaceDirectoryRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Registers the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is registered. If this is the first time you
// are registering a directory, you will need to create the workspaces_DefaultRole
// role before you can register a directory. For more information, see Creating
// the workspaces_DefaultRole Role (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role).
//
//    // Example sending a request using RegisterWorkspaceDirectoryRequest.
//    req := client.RegisterWorkspaceDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/RegisterWorkspaceDirectory
func (c *Client) RegisterWorkspaceDirectoryRequest(input *RegisterWorkspaceDirectoryInput) RegisterWorkspaceDirectoryRequest {
	op := &aws.Operation{
		Name:       opRegisterWorkspaceDirectory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterWorkspaceDirectoryInput{}
	}

	req := c.newRequest(op, input, &RegisterWorkspaceDirectoryOutput{})

	return RegisterWorkspaceDirectoryRequest{Request: req, Input: input, Copy: c.RegisterWorkspaceDirectoryRequest}
}

// RegisterWorkspaceDirectoryRequest is the request type for the
// RegisterWorkspaceDirectory API operation.
type RegisterWorkspaceDirectoryRequest struct {
	*aws.Request
	Input *RegisterWorkspaceDirectoryInput
	Copy  func(*RegisterWorkspaceDirectoryInput) RegisterWorkspaceDirectoryRequest
}

// Send marshals and sends the RegisterWorkspaceDirectory API request.
func (r RegisterWorkspaceDirectoryRequest) Send(ctx context.Context) (*RegisterWorkspaceDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterWorkspaceDirectoryResponse{
		RegisterWorkspaceDirectoryOutput: r.Request.Data.(*RegisterWorkspaceDirectoryOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterWorkspaceDirectoryResponse is the response type for the
// RegisterWorkspaceDirectory API operation.
type RegisterWorkspaceDirectoryResponse struct {
	*RegisterWorkspaceDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterWorkspaceDirectory request.
func (r *RegisterWorkspaceDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
