// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetAppInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAppInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAppInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetAppOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationResponse"`

	// Provides information about an application.
	//
	// ApplicationResponse is a required field
	ApplicationResponse *ApplicationResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetAppOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAppOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationResponse != nil {
		v := s.ApplicationResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ApplicationResponse", v, metadata)
	}
	return nil
}

const opGetApp = "GetApp"

// GetAppRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about an application.
//
//    // Example sending a request using GetAppRequest.
//    req := client.GetAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetApp
func (c *Client) GetAppRequest(input *GetAppInput) GetAppRequest {
	op := &aws.Operation{
		Name:       opGetApp,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}",
	}

	if input == nil {
		input = &GetAppInput{}
	}

	req := c.newRequest(op, input, &GetAppOutput{})

	return GetAppRequest{Request: req, Input: input, Copy: c.GetAppRequest}
}

// GetAppRequest is the request type for the
// GetApp API operation.
type GetAppRequest struct {
	*aws.Request
	Input *GetAppInput
	Copy  func(*GetAppInput) GetAppRequest
}

// Send marshals and sends the GetApp API request.
func (r GetAppRequest) Send(ctx context.Context) (*GetAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAppResponse{
		GetAppOutput: r.Request.Data.(*GetAppOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAppResponse is the response type for the
// GetApp API operation.
type GetAppResponse struct {
	*GetAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApp request.
func (r *GetAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
