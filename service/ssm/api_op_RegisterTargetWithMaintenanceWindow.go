// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterTargetWithMaintenanceWindowInput struct {
	_ struct{} `type:"structure"`

	// User-provided idempotency token.
	ClientToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// An optional description for the target.
	Description *string `min:"1" type:"string" sensitive:"true"`

	// An optional name for the target.
	Name *string `min:"3" type:"string"`

	// User-provided value that will be included in any CloudWatch events raised
	// while running tasks for these targets in this maintenance window.
	OwnerInformation *string `min:"1" type:"string" sensitive:"true"`

	// The type of target being registered with the maintenance window.
	//
	// ResourceType is a required field
	ResourceType MaintenanceWindowResourceType `type:"string" required:"true" enum:"true"`

	// The targets to register with the maintenance window. In other words, the
	// instances to run commands on when the maintenance window runs.
	//
	// You can specify targets using instance IDs, resource group names, or tags
	// that have been applied to instances.
	//
	// Example 1: Specify instance IDs
	//
	// Key=InstanceIds,Values=instance-id-1,instance-id-2,instance-id-3
	//
	// Example 2: Use tag key-pairs applied to instances
	//
	// Key=tag:my-tag-key,Values=my-tag-value-1,my-tag-value-2
	//
	// Example 3: Use tag-keys applied to instances
	//
	// Key=tag-key,Values=my-tag-key-1,my-tag-key-2
	//
	// Example 4: Use resource group names
	//
	// Key=resource-groups:Name,Values=resource-group-name
	//
	// Example 5: Use filters for resource group types
	//
	// Key=resource-groups:ResourceTypeFilters,Values=resource-type-1,resource-type-2
	//
	// For Key=resource-groups:ResourceTypeFilters, specify resource types in the
	// following format
	//
	// Key=resource-groups:ResourceTypeFilters,Values=AWS::EC2::INSTANCE,AWS::EC2::VPC
	//
	// For more information about these examples formats, including the best use
	// case for each one, see Examples: Register Targets with a Maintenance Window
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	// in the AWS Systems Manager User Guide.
	//
	// Targets is a required field
	Targets []Target `type:"list" required:"true"`

	// The ID of the maintenance window the target should be registered with.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterTargetWithMaintenanceWindowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterTargetWithMaintenanceWindowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterTargetWithMaintenanceWindowInput"}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}
	if s.OwnerInformation != nil && len(*s.OwnerInformation) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OwnerInformation", 1))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterTargetWithMaintenanceWindowOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the target definition in this maintenance window.
	WindowTargetId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s RegisterTargetWithMaintenanceWindowOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterTargetWithMaintenanceWindow = "RegisterTargetWithMaintenanceWindow"

// RegisterTargetWithMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Registers a target with a maintenance window.
//
//    // Example sending a request using RegisterTargetWithMaintenanceWindowRequest.
//    req := client.RegisterTargetWithMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/RegisterTargetWithMaintenanceWindow
func (c *Client) RegisterTargetWithMaintenanceWindowRequest(input *RegisterTargetWithMaintenanceWindowInput) RegisterTargetWithMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opRegisterTargetWithMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterTargetWithMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &RegisterTargetWithMaintenanceWindowOutput{})
	return RegisterTargetWithMaintenanceWindowRequest{Request: req, Input: input, Copy: c.RegisterTargetWithMaintenanceWindowRequest}
}

// RegisterTargetWithMaintenanceWindowRequest is the request type for the
// RegisterTargetWithMaintenanceWindow API operation.
type RegisterTargetWithMaintenanceWindowRequest struct {
	*aws.Request
	Input *RegisterTargetWithMaintenanceWindowInput
	Copy  func(*RegisterTargetWithMaintenanceWindowInput) RegisterTargetWithMaintenanceWindowRequest
}

// Send marshals and sends the RegisterTargetWithMaintenanceWindow API request.
func (r RegisterTargetWithMaintenanceWindowRequest) Send(ctx context.Context) (*RegisterTargetWithMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterTargetWithMaintenanceWindowResponse{
		RegisterTargetWithMaintenanceWindowOutput: r.Request.Data.(*RegisterTargetWithMaintenanceWindowOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterTargetWithMaintenanceWindowResponse is the response type for the
// RegisterTargetWithMaintenanceWindow API operation.
type RegisterTargetWithMaintenanceWindowResponse struct {
	*RegisterTargetWithMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterTargetWithMaintenanceWindow request.
func (r *RegisterTargetWithMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
