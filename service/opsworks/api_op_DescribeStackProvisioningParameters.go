// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStackProvisioningParametersInput struct {
	_ struct{} `type:"structure"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackProvisioningParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackProvisioningParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackProvisioningParametersInput"}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a DescribeStackProvisioningParameters request.
type DescribeStackProvisioningParametersOutput struct {
	_ struct{} `type:"structure"`

	// The AWS OpsWorks Stacks agent installer's URL.
	AgentInstallerUrl *string `type:"string"`

	// An embedded object that contains the provisioning parameters.
	Parameters map[string]string `type:"map"`
}

// String returns the string representation
func (s DescribeStackProvisioningParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackProvisioningParameters = "DescribeStackProvisioningParameters"

// DescribeStackProvisioningParametersRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Requests a description of a stack's provisioning parameters.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeStackProvisioningParametersRequest.
//    req := client.DescribeStackProvisioningParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeStackProvisioningParameters
func (c *Client) DescribeStackProvisioningParametersRequest(input *DescribeStackProvisioningParametersInput) DescribeStackProvisioningParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeStackProvisioningParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStackProvisioningParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeStackProvisioningParametersOutput{})

	return DescribeStackProvisioningParametersRequest{Request: req, Input: input, Copy: c.DescribeStackProvisioningParametersRequest}
}

// DescribeStackProvisioningParametersRequest is the request type for the
// DescribeStackProvisioningParameters API operation.
type DescribeStackProvisioningParametersRequest struct {
	*aws.Request
	Input *DescribeStackProvisioningParametersInput
	Copy  func(*DescribeStackProvisioningParametersInput) DescribeStackProvisioningParametersRequest
}

// Send marshals and sends the DescribeStackProvisioningParameters API request.
func (r DescribeStackProvisioningParametersRequest) Send(ctx context.Context) (*DescribeStackProvisioningParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackProvisioningParametersResponse{
		DescribeStackProvisioningParametersOutput: r.Request.Data.(*DescribeStackProvisioningParametersOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackProvisioningParametersResponse is the response type for the
// DescribeStackProvisioningParameters API operation.
type DescribeStackProvisioningParametersResponse struct {
	*DescribeStackProvisioningParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackProvisioningParameters request.
func (r *DescribeStackProvisioningParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
