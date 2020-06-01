// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CreateTaskRequest
type CreateTaskInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is
	// used to monitor and log events in the task.
	//
	// For more information on these groups, see Working with Log Groups and Log
	// Streams in the Amazon CloudWatch User Guide.
	//
	// For more information about how to use CloudWatch Logs with DataSync, see
	// Monitoring Your Task in the AWS DataSync User Guide.
	CloudWatchLogGroupArn *string `type:"string"`

	// The Amazon Resource Name (ARN) of an AWS storage resource's location.
	//
	// DestinationLocationArn is a required field
	DestinationLocationArn *string `type:"string" required:"true"`

	// A list of filter rules that determines which files to exclude from a task.
	// The list should contain a single filter string that consists of the patterns
	// to exclude. The patterns are delimited by "|" (that is, a pipe), for example,
	// "/folder1|/folder2"
	Excludes []FilterRule `type:"list"`

	// The name of a task. This value is a text reference that is used to identify
	// the task in the console.
	Name *string `min:"1" type:"string"`

	// The set of configuration options that control the behavior of a single execution
	// of the task that occurs when you call StartTaskExecution. You can configure
	// these options to preserve metadata such as user ID (UID) and group ID (GID),
	// file permissions, data integrity verification, and so on.
	//
	// For each individual task execution, you can override these options by specifying
	// the OverrideOptions before starting a the task execution. For more information,
	// see the operation.
	Options *Options `type:"structure"`

	// Specifies a schedule used to periodically transfer files from a source to
	// a destination location. The schedule should be specified in UTC time. For
	// more information, see task-scheduling.
	Schedule *TaskSchedule `type:"structure"`

	// The Amazon Resource Name (ARN) of the source location for the task.
	//
	// SourceLocationArn is a required field
	SourceLocationArn *string `type:"string" required:"true"`

	// The key-value pair that represents the tag that you want to add to the resource.
	// The value can be an empty string.
	Tags []TagListEntry `type:"list"`
}

// String returns the string representation
func (s CreateTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTaskInput"}

	if s.DestinationLocationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationLocationArn"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SourceLocationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceLocationArn"))
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			invalidParams.AddNested("Options", err.(aws.ErrInvalidParams))
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			invalidParams.AddNested("Schedule", err.(aws.ErrInvalidParams))
		}
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

// CreateTaskResponse
type CreateTaskOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string `type:"string"`
}

// String returns the string representation
func (s CreateTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTask = "CreateTask"

// CreateTaskRequest returns a request value for making API operation for
// AWS DataSync.
//
// Creates a task. A task is a set of two locations (source and destination)
// and a set of Options that you use to control the behavior of a task. If you
// don't specify Options when you create a task, AWS DataSync populates them
// with service defaults.
//
// When you create a task, it first enters the CREATING state. During CREATING
// AWS DataSync attempts to mount the on-premises Network File System (NFS)
// location. The task transitions to the AVAILABLE state without waiting for
// the AWS location to become mounted. If required, AWS DataSync mounts the
// AWS location before each task execution.
//
// If an agent that is associated with a source (NFS) location goes offline,
// the task transitions to the UNAVAILABLE status. If the status of the task
// remains in the CREATING status for more than a few minutes, it means that
// your agent might be having trouble mounting the source NFS file system. Check
// the task's ErrorCode and ErrorDetail. Mount issues are often caused by either
// a misconfigured firewall or a mistyped NFS server host name.
//
//    // Example sending a request using CreateTaskRequest.
//    req := client.CreateTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateTask
func (c *Client) CreateTaskRequest(input *CreateTaskInput) CreateTaskRequest {
	op := &aws.Operation{
		Name:       opCreateTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTaskInput{}
	}

	req := c.newRequest(op, input, &CreateTaskOutput{})

	return CreateTaskRequest{Request: req, Input: input, Copy: c.CreateTaskRequest}
}

// CreateTaskRequest is the request type for the
// CreateTask API operation.
type CreateTaskRequest struct {
	*aws.Request
	Input *CreateTaskInput
	Copy  func(*CreateTaskInput) CreateTaskRequest
}

// Send marshals and sends the CreateTask API request.
func (r CreateTaskRequest) Send(ctx context.Context) (*CreateTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTaskResponse{
		CreateTaskOutput: r.Request.Data.(*CreateTaskOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTaskResponse is the response type for the
// CreateTask API operation.
type CreateTaskResponse struct {
	*CreateTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTask request.
func (r *CreateTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
