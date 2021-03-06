// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyWorkspaceAccessPropertiesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory.
	//
	// ResourceId is a required field
	ResourceId *string `min:"10" type:"string" required:"true"`

	// The device types and operating systems to enable or disable for access.
	//
	// WorkspaceAccessProperties is a required field
	WorkspaceAccessProperties *WorkspaceAccessProperties `type:"structure" required:"true"`
}

// String returns the string representation
func (s ModifyWorkspaceAccessPropertiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyWorkspaceAccessPropertiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyWorkspaceAccessPropertiesInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 10))
	}

	if s.WorkspaceAccessProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceAccessProperties"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyWorkspaceAccessPropertiesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyWorkspaceAccessPropertiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyWorkspaceAccessProperties = "ModifyWorkspaceAccessProperties"

// ModifyWorkspaceAccessPropertiesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Specifies which devices and operating systems users can use to access their
// WorkSpaces. For more information, see Control Device Access (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html#control-device-access).
//
//    // Example sending a request using ModifyWorkspaceAccessPropertiesRequest.
//    req := client.ModifyWorkspaceAccessPropertiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ModifyWorkspaceAccessProperties
func (c *Client) ModifyWorkspaceAccessPropertiesRequest(input *ModifyWorkspaceAccessPropertiesInput) ModifyWorkspaceAccessPropertiesRequest {
	op := &aws.Operation{
		Name:       opModifyWorkspaceAccessProperties,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyWorkspaceAccessPropertiesInput{}
	}

	req := c.newRequest(op, input, &ModifyWorkspaceAccessPropertiesOutput{})
	return ModifyWorkspaceAccessPropertiesRequest{Request: req, Input: input, Copy: c.ModifyWorkspaceAccessPropertiesRequest}
}

// ModifyWorkspaceAccessPropertiesRequest is the request type for the
// ModifyWorkspaceAccessProperties API operation.
type ModifyWorkspaceAccessPropertiesRequest struct {
	*aws.Request
	Input *ModifyWorkspaceAccessPropertiesInput
	Copy  func(*ModifyWorkspaceAccessPropertiesInput) ModifyWorkspaceAccessPropertiesRequest
}

// Send marshals and sends the ModifyWorkspaceAccessProperties API request.
func (r ModifyWorkspaceAccessPropertiesRequest) Send(ctx context.Context) (*ModifyWorkspaceAccessPropertiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyWorkspaceAccessPropertiesResponse{
		ModifyWorkspaceAccessPropertiesOutput: r.Request.Data.(*ModifyWorkspaceAccessPropertiesOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyWorkspaceAccessPropertiesResponse is the response type for the
// ModifyWorkspaceAccessProperties API operation.
type ModifyWorkspaceAccessPropertiesResponse struct {
	*ModifyWorkspaceAccessPropertiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyWorkspaceAccessProperties request.
func (r *ModifyWorkspaceAccessPropertiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
