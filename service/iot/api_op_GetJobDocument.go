// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetJobDocumentInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier you assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobDocumentInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobDocumentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetJobDocumentOutput struct {
	_ struct{} `type:"structure"`

	// The job document content.
	Document *string `locationName:"document" type:"string"`
}

// String returns the string representation
func (s GetJobDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobDocumentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Document != nil {
		v := *s.Document

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "document", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetJobDocument = "GetJobDocument"

// GetJobDocumentRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets a job document.
//
//    // Example sending a request using GetJobDocumentRequest.
//    req := client.GetJobDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetJobDocumentRequest(input *GetJobDocumentInput) GetJobDocumentRequest {
	op := &aws.Operation{
		Name:       opGetJobDocument,
		HTTPMethod: "GET",
		HTTPPath:   "/jobs/{jobId}/job-document",
	}

	if input == nil {
		input = &GetJobDocumentInput{}
	}

	req := c.newRequest(op, input, &GetJobDocumentOutput{})

	return GetJobDocumentRequest{Request: req, Input: input, Copy: c.GetJobDocumentRequest}
}

// GetJobDocumentRequest is the request type for the
// GetJobDocument API operation.
type GetJobDocumentRequest struct {
	*aws.Request
	Input *GetJobDocumentInput
	Copy  func(*GetJobDocumentInput) GetJobDocumentRequest
}

// Send marshals and sends the GetJobDocument API request.
func (r GetJobDocumentRequest) Send(ctx context.Context) (*GetJobDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobDocumentResponse{
		GetJobDocumentOutput: r.Request.Data.(*GetJobDocumentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobDocumentResponse is the response type for the
// GetJobDocument API operation.
type GetJobDocumentResponse struct {
	*GetJobDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJobDocument request.
func (r *GetJobDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
