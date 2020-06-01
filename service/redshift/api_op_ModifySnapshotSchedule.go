// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifySnapshotScheduleInput struct {
	_ struct{} `type:"structure"`

	// An updated list of schedule definitions. A schedule definition is made up
	// of schedule expressions, for example, "cron(30 12 *)" or "rate(12 hours)".
	//
	// ScheduleDefinitions is a required field
	ScheduleDefinitions []string `locationNameList:"ScheduleDefinition" type:"list" required:"true"`

	// A unique alphanumeric identifier of the schedule to modify.
	//
	// ScheduleIdentifier is a required field
	ScheduleIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifySnapshotScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifySnapshotScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifySnapshotScheduleInput"}

	if s.ScheduleDefinitions == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduleDefinitions"))
	}

	if s.ScheduleIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduleIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a snapshot schedule. You can set a regular interval for creating
// snapshots of a cluster. You can also schedule snapshots for specific dates.
type ModifySnapshotScheduleOutput struct {
	_ struct{} `type:"structure"`

	// The number of clusters associated with the schedule.
	AssociatedClusterCount *int64 `type:"integer"`

	// A list of clusters associated with the schedule. A maximum of 100 clusters
	// is returned.
	AssociatedClusters []ClusterAssociatedToSchedule `locationNameList:"ClusterAssociatedToSchedule" type:"list"`

	NextInvocations []time.Time `locationNameList:"SnapshotTime" type:"list"`

	// A list of ScheduleDefinitions.
	ScheduleDefinitions []string `locationNameList:"ScheduleDefinition" type:"list"`

	// The description of the schedule.
	ScheduleDescription *string `type:"string"`

	// A unique identifier for the schedule.
	ScheduleIdentifier *string `type:"string"`

	// An optional set of tags describing the schedule.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s ModifySnapshotScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifySnapshotSchedule = "ModifySnapshotSchedule"

// ModifySnapshotScheduleRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies a snapshot schedule. Any schedule associated with a cluster is modified
// asynchronously.
//
//    // Example sending a request using ModifySnapshotScheduleRequest.
//    req := client.ModifySnapshotScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifySnapshotSchedule
func (c *Client) ModifySnapshotScheduleRequest(input *ModifySnapshotScheduleInput) ModifySnapshotScheduleRequest {
	op := &aws.Operation{
		Name:       opModifySnapshotSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySnapshotScheduleInput{}
	}

	req := c.newRequest(op, input, &ModifySnapshotScheduleOutput{})

	return ModifySnapshotScheduleRequest{Request: req, Input: input, Copy: c.ModifySnapshotScheduleRequest}
}

// ModifySnapshotScheduleRequest is the request type for the
// ModifySnapshotSchedule API operation.
type ModifySnapshotScheduleRequest struct {
	*aws.Request
	Input *ModifySnapshotScheduleInput
	Copy  func(*ModifySnapshotScheduleInput) ModifySnapshotScheduleRequest
}

// Send marshals and sends the ModifySnapshotSchedule API request.
func (r ModifySnapshotScheduleRequest) Send(ctx context.Context) (*ModifySnapshotScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifySnapshotScheduleResponse{
		ModifySnapshotScheduleOutput: r.Request.Data.(*ModifySnapshotScheduleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifySnapshotScheduleResponse is the response type for the
// ModifySnapshotSchedule API operation.
type ModifySnapshotScheduleResponse struct {
	*ModifySnapshotScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifySnapshotSchedule request.
func (r *ModifySnapshotScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
