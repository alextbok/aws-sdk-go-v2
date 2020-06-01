// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StopDataSourceSyncJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the data source for which to stop the synchronization jobs.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The identifier of the index that contains the data source.
	//
	// IndexId is a required field
	IndexId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s StopDataSourceSyncJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopDataSourceSyncJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopDataSourceSyncJobInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.IndexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexId"))
	}
	if s.IndexId != nil && len(*s.IndexId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopDataSourceSyncJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopDataSourceSyncJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopDataSourceSyncJob = "StopDataSourceSyncJob"

// StopDataSourceSyncJobRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Stops a running synchronization job. You can't stop a scheduled synchronization
// job.
//
//    // Example sending a request using StopDataSourceSyncJobRequest.
//    req := client.StopDataSourceSyncJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/StopDataSourceSyncJob
func (c *Client) StopDataSourceSyncJobRequest(input *StopDataSourceSyncJobInput) StopDataSourceSyncJobRequest {
	op := &aws.Operation{
		Name:       opStopDataSourceSyncJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopDataSourceSyncJobInput{}
	}

	req := c.newRequest(op, input, &StopDataSourceSyncJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return StopDataSourceSyncJobRequest{Request: req, Input: input, Copy: c.StopDataSourceSyncJobRequest}
}

// StopDataSourceSyncJobRequest is the request type for the
// StopDataSourceSyncJob API operation.
type StopDataSourceSyncJobRequest struct {
	*aws.Request
	Input *StopDataSourceSyncJobInput
	Copy  func(*StopDataSourceSyncJobInput) StopDataSourceSyncJobRequest
}

// Send marshals and sends the StopDataSourceSyncJob API request.
func (r StopDataSourceSyncJobRequest) Send(ctx context.Context) (*StopDataSourceSyncJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopDataSourceSyncJobResponse{
		StopDataSourceSyncJobOutput: r.Request.Data.(*StopDataSourceSyncJobOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopDataSourceSyncJobResponse is the response type for the
// StopDataSourceSyncJob API operation.
type StopDataSourceSyncJobResponse struct {
	*StopDataSourceSyncJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopDataSourceSyncJob request.
func (r *StopDataSourceSyncJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
