// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DeleteDeploymentGroup operation.
type DeleteDeploymentGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of an AWS CodeDeploy application associated with the IAM user or
	// AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// The name of a deployment group for the specified application.
	//
	// DeploymentGroupName is a required field
	DeploymentGroupName *string `locationName:"deploymentGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDeploymentGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDeploymentGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDeploymentGroupInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.DeploymentGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentGroupName"))
	}
	if s.DeploymentGroupName != nil && len(*s.DeploymentGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DeleteDeploymentGroup operation.
type DeleteDeploymentGroupOutput struct {
	_ struct{} `type:"structure"`

	// If the output contains no data, and the corresponding deployment group contained
	// at least one Auto Scaling group, AWS CodeDeploy successfully removed all
	// corresponding Auto Scaling lifecycle event hooks from the Amazon EC2 instances
	// in the Auto Scaling group. If the output contains data, AWS CodeDeploy could
	// not remove some Auto Scaling lifecycle event hooks from the Amazon EC2 instances
	// in the Auto Scaling group.
	HooksNotCleanedUp []AutoScalingGroup `locationName:"hooksNotCleanedUp" type:"list"`
}

// String returns the string representation
func (s DeleteDeploymentGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDeploymentGroup = "DeleteDeploymentGroup"

// DeleteDeploymentGroupRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Deletes a deployment group.
//
//    // Example sending a request using DeleteDeploymentGroupRequest.
//    req := client.DeleteDeploymentGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeleteDeploymentGroup
func (c *Client) DeleteDeploymentGroupRequest(input *DeleteDeploymentGroupInput) DeleteDeploymentGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDeploymentGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDeploymentGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDeploymentGroupOutput{})

	return DeleteDeploymentGroupRequest{Request: req, Input: input, Copy: c.DeleteDeploymentGroupRequest}
}

// DeleteDeploymentGroupRequest is the request type for the
// DeleteDeploymentGroup API operation.
type DeleteDeploymentGroupRequest struct {
	*aws.Request
	Input *DeleteDeploymentGroupInput
	Copy  func(*DeleteDeploymentGroupInput) DeleteDeploymentGroupRequest
}

// Send marshals and sends the DeleteDeploymentGroup API request.
func (r DeleteDeploymentGroupRequest) Send(ctx context.Context) (*DeleteDeploymentGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDeploymentGroupResponse{
		DeleteDeploymentGroupOutput: r.Request.Data.(*DeleteDeploymentGroupOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDeploymentGroupResponse is the response type for the
// DeleteDeploymentGroup API operation.
type DeleteDeploymentGroupResponse struct {
	*DeleteDeploymentGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDeploymentGroup request.
func (r *DeleteDeploymentGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
