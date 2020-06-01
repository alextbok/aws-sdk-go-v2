// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSnapshotSchedulesInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the cluster whose snapshot schedules you want to
	// view.
	ClusterIdentifier *string `type:"string"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// The maximum number or response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	MaxRecords *int64 `type:"integer"`

	// A unique identifier for a snapshot schedule.
	ScheduleIdentifier *string `type:"string"`

	// The key value for a snapshot schedule tag.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// The value corresponding to the key of the snapshot schedule tag.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotSchedulesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeSnapshotSchedulesOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// A list of SnapshotSchedules.
	SnapshotSchedules []SnapshotSchedule `locationNameList:"SnapshotSchedule" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotSchedulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSnapshotSchedules = "DescribeSnapshotSchedules"

// DescribeSnapshotSchedulesRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns a list of snapshot schedules.
//
//    // Example sending a request using DescribeSnapshotSchedulesRequest.
//    req := client.DescribeSnapshotSchedulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeSnapshotSchedules
func (c *Client) DescribeSnapshotSchedulesRequest(input *DescribeSnapshotSchedulesInput) DescribeSnapshotSchedulesRequest {
	op := &aws.Operation{
		Name:       opDescribeSnapshotSchedules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSnapshotSchedulesInput{}
	}

	req := c.newRequest(op, input, &DescribeSnapshotSchedulesOutput{})

	return DescribeSnapshotSchedulesRequest{Request: req, Input: input, Copy: c.DescribeSnapshotSchedulesRequest}
}

// DescribeSnapshotSchedulesRequest is the request type for the
// DescribeSnapshotSchedules API operation.
type DescribeSnapshotSchedulesRequest struct {
	*aws.Request
	Input *DescribeSnapshotSchedulesInput
	Copy  func(*DescribeSnapshotSchedulesInput) DescribeSnapshotSchedulesRequest
}

// Send marshals and sends the DescribeSnapshotSchedules API request.
func (r DescribeSnapshotSchedulesRequest) Send(ctx context.Context) (*DescribeSnapshotSchedulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSnapshotSchedulesResponse{
		DescribeSnapshotSchedulesOutput: r.Request.Data.(*DescribeSnapshotSchedulesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSnapshotSchedulesResponse is the response type for the
// DescribeSnapshotSchedules API operation.
type DescribeSnapshotSchedulesResponse struct {
	*DescribeSnapshotSchedulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSnapshotSchedules request.
func (r *DescribeSnapshotSchedulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
