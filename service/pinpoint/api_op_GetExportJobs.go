// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetExportJobsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetExportJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetExportJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetExportJobsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetExportJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetExportJobsOutput struct {
	_ struct{} `type:"structure" payload:"ExportJobsResponse"`

	// Provides information about all the export jobs that are associated with an
	// application or segment. An export job is a job that exports endpoint definitions
	// to a file.
	//
	// ExportJobsResponse is a required field
	ExportJobsResponse *ExportJobsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetExportJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetExportJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExportJobsResponse != nil {
		v := s.ExportJobsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ExportJobsResponse", v, metadata)
	}
	return nil
}

const opGetExportJobs = "GetExportJobs"

// GetExportJobsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of all the export jobs
// for an application.
//
//    // Example sending a request using GetExportJobsRequest.
//    req := client.GetExportJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetExportJobs
func (c *Client) GetExportJobsRequest(input *GetExportJobsInput) GetExportJobsRequest {
	op := &aws.Operation{
		Name:       opGetExportJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/jobs/export",
	}

	if input == nil {
		input = &GetExportJobsInput{}
	}

	req := c.newRequest(op, input, &GetExportJobsOutput{})

	return GetExportJobsRequest{Request: req, Input: input, Copy: c.GetExportJobsRequest}
}

// GetExportJobsRequest is the request type for the
// GetExportJobs API operation.
type GetExportJobsRequest struct {
	*aws.Request
	Input *GetExportJobsInput
	Copy  func(*GetExportJobsInput) GetExportJobsRequest
}

// Send marshals and sends the GetExportJobs API request.
func (r GetExportJobsRequest) Send(ctx context.Context) (*GetExportJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetExportJobsResponse{
		GetExportJobsOutput: r.Request.Data.(*GetExportJobsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetExportJobsResponse is the response type for the
// GetExportJobs API operation.
type GetExportJobsResponse struct {
	*GetExportJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetExportJobs request.
func (r *GetExportJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
