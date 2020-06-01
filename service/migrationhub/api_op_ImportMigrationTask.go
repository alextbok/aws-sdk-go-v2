// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ImportMigrationTaskInput struct {
	_ struct{} `type:"structure"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// Unique identifier that references the migration task. Do not store personal
	// data in this field.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// The name of the ProgressUpdateStream. >
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ImportMigrationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportMigrationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportMigrationTaskInput"}

	if s.MigrationTaskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MigrationTaskName"))
	}
	if s.MigrationTaskName != nil && len(*s.MigrationTaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MigrationTaskName", 1))
	}

	if s.ProgressUpdateStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStream"))
	}
	if s.ProgressUpdateStream != nil && len(*s.ProgressUpdateStream) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStream", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ImportMigrationTaskOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ImportMigrationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportMigrationTask = "ImportMigrationTask"

// ImportMigrationTaskRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Registers a new migration task which represents a server, database, etc.,
// being migrated to AWS by a migration tool.
//
// This API is a prerequisite to calling the NotifyMigrationTaskState API as
// the migration tool must first register the migration task with Migration
// Hub.
//
//    // Example sending a request using ImportMigrationTaskRequest.
//    req := client.ImportMigrationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ImportMigrationTask
func (c *Client) ImportMigrationTaskRequest(input *ImportMigrationTaskInput) ImportMigrationTaskRequest {
	op := &aws.Operation{
		Name:       opImportMigrationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportMigrationTaskInput{}
	}

	req := c.newRequest(op, input, &ImportMigrationTaskOutput{})

	return ImportMigrationTaskRequest{Request: req, Input: input, Copy: c.ImportMigrationTaskRequest}
}

// ImportMigrationTaskRequest is the request type for the
// ImportMigrationTask API operation.
type ImportMigrationTaskRequest struct {
	*aws.Request
	Input *ImportMigrationTaskInput
	Copy  func(*ImportMigrationTaskInput) ImportMigrationTaskRequest
}

// Send marshals and sends the ImportMigrationTask API request.
func (r ImportMigrationTaskRequest) Send(ctx context.Context) (*ImportMigrationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportMigrationTaskResponse{
		ImportMigrationTaskOutput: r.Request.Data.(*ImportMigrationTaskOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportMigrationTaskResponse is the response type for the
// ImportMigrationTask API operation.
type ImportMigrationTaskResponse struct {
	*ImportMigrationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportMigrationTask request.
func (r *ImportMigrationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
