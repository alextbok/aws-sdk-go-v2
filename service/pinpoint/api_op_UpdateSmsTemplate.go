// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateSmsTemplateInput struct {
	_ struct{} `type:"structure" payload:"SMSTemplateRequest"`

	CreateNewVersion *bool `location:"querystring" locationName:"create-new-version" type:"boolean"`

	// Specifies the content and settings for a message template that can be used
	// in text messages that are sent through the SMS channel.
	//
	// SMSTemplateRequest is a required field
	SMSTemplateRequest *SMSTemplateRequest `type:"structure" required:"true"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`

	Version *string `location:"querystring" locationName:"version" type:"string"`
}

// String returns the string representation
func (s UpdateSmsTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSmsTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSmsTemplateInput"}

	if s.SMSTemplateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("SMSTemplateRequest"))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSmsTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SMSTemplateRequest != nil {
		v := s.SMSTemplateRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SMSTemplateRequest", v, metadata)
	}
	if s.CreateNewVersion != nil {
		v := *s.CreateNewVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "create-new-version", protocol.BoolValue(v), metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateSmsTemplateOutput struct {
	_ struct{} `type:"structure" payload:"MessageBody"`

	// Provides information about an API request or response.
	//
	// MessageBody is a required field
	MessageBody *MessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateSmsTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSmsTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MessageBody != nil {
		v := s.MessageBody

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MessageBody", v, metadata)
	}
	return nil
}

const opUpdateSmsTemplate = "UpdateSmsTemplate"

// UpdateSmsTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates an existing message template for messages that are sent through the
// SMS channel.
//
//    // Example sending a request using UpdateSmsTemplateRequest.
//    req := client.UpdateSmsTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateSmsTemplate
func (c *Client) UpdateSmsTemplateRequest(input *UpdateSmsTemplateInput) UpdateSmsTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateSmsTemplate,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/templates/{template-name}/sms",
	}

	if input == nil {
		input = &UpdateSmsTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateSmsTemplateOutput{})

	return UpdateSmsTemplateRequest{Request: req, Input: input, Copy: c.UpdateSmsTemplateRequest}
}

// UpdateSmsTemplateRequest is the request type for the
// UpdateSmsTemplate API operation.
type UpdateSmsTemplateRequest struct {
	*aws.Request
	Input *UpdateSmsTemplateInput
	Copy  func(*UpdateSmsTemplateInput) UpdateSmsTemplateRequest
}

// Send marshals and sends the UpdateSmsTemplate API request.
func (r UpdateSmsTemplateRequest) Send(ctx context.Context) (*UpdateSmsTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSmsTemplateResponse{
		UpdateSmsTemplateOutput: r.Request.Data.(*UpdateSmsTemplateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSmsTemplateResponse is the response type for the
// UpdateSmsTemplate API operation.
type UpdateSmsTemplateResponse struct {
	*UpdateSmsTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSmsTemplate request.
func (r *UpdateSmsTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
