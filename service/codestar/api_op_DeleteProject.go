// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteProjectInput struct {
	_ struct{} `type:"structure"`

	// A user- or system-generated token that identifies the entity that requested
	// project deletion. This token can be used to repeat the request.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string"`

	// Whether to send a delete request for the primary stack in AWS CloudFormation
	// originally used to generate the project and its resources. This option will
	// delete all AWS resources for the project (except for any buckets in Amazon
	// S3) as well as deleting the project itself. Recommended for most use cases.
	DeleteStack *bool `locationName:"deleteStack" type:"boolean"`

	// The ID of the project to be deleted in AWS CodeStar.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProjectInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteProjectOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the deleted project.
	ProjectArn *string `locationName:"projectArn" type:"string"`

	// The ID of the primary stack in AWS CloudFormation that will be deleted as
	// part of deleting the project and its resources.
	StackId *string `locationName:"stackId" type:"string"`
}

// String returns the string representation
func (s DeleteProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteProject = "DeleteProject"

// DeleteProjectRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Deletes a project, including project resources. Does not delete users associated
// with the project, but does delete the IAM roles that allowed access to the
// project.
//
//    // Example sending a request using DeleteProjectRequest.
//    req := client.DeleteProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/DeleteProject
func (c *Client) DeleteProjectRequest(input *DeleteProjectInput) DeleteProjectRequest {
	op := &aws.Operation{
		Name:       opDeleteProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProjectInput{}
	}

	req := c.newRequest(op, input, &DeleteProjectOutput{})

	return DeleteProjectRequest{Request: req, Input: input, Copy: c.DeleteProjectRequest}
}

// DeleteProjectRequest is the request type for the
// DeleteProject API operation.
type DeleteProjectRequest struct {
	*aws.Request
	Input *DeleteProjectInput
	Copy  func(*DeleteProjectInput) DeleteProjectRequest
}

// Send marshals and sends the DeleteProject API request.
func (r DeleteProjectRequest) Send(ctx context.Context) (*DeleteProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProjectResponse{
		DeleteProjectOutput: r.Request.Data.(*DeleteProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProjectResponse is the response type for the
// DeleteProject API operation.
type DeleteProjectResponse struct {
	*DeleteProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProject request.
func (r *DeleteProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
