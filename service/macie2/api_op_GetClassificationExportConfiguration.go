// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetClassificationExportConfigurationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetClassificationExportConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClassificationExportConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

// Provides information about the current configuration settings for exporting
// data classification results.
type GetClassificationExportConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Specifies where to export data classification results to, and the encryption
	// settings to use when storing results in that location. Currently, you can
	// export classification results only to an S3 bucket.
	Configuration *ClassificationExportConfiguration `locationName:"configuration" type:"structure"`
}

// String returns the string representation
func (s GetClassificationExportConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClassificationExportConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Configuration != nil {
		v := s.Configuration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configuration", v, metadata)
	}
	return nil
}

const opGetClassificationExportConfiguration = "GetClassificationExportConfiguration"

// GetClassificationExportConfigurationRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves the configuration settings for exporting data classification results.
//
//    // Example sending a request using GetClassificationExportConfigurationRequest.
//    req := client.GetClassificationExportConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/GetClassificationExportConfiguration
func (c *Client) GetClassificationExportConfigurationRequest(input *GetClassificationExportConfigurationInput) GetClassificationExportConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetClassificationExportConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/classification-export-configuration",
	}

	if input == nil {
		input = &GetClassificationExportConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetClassificationExportConfigurationOutput{})

	return GetClassificationExportConfigurationRequest{Request: req, Input: input, Copy: c.GetClassificationExportConfigurationRequest}
}

// GetClassificationExportConfigurationRequest is the request type for the
// GetClassificationExportConfiguration API operation.
type GetClassificationExportConfigurationRequest struct {
	*aws.Request
	Input *GetClassificationExportConfigurationInput
	Copy  func(*GetClassificationExportConfigurationInput) GetClassificationExportConfigurationRequest
}

// Send marshals and sends the GetClassificationExportConfiguration API request.
func (r GetClassificationExportConfigurationRequest) Send(ctx context.Context) (*GetClassificationExportConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetClassificationExportConfigurationResponse{
		GetClassificationExportConfigurationOutput: r.Request.Data.(*GetClassificationExportConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetClassificationExportConfigurationResponse is the response type for the
// GetClassificationExportConfiguration API operation.
type GetClassificationExportConfigurationResponse struct {
	*GetClassificationExportConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetClassificationExportConfiguration request.
func (r *GetClassificationExportConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
