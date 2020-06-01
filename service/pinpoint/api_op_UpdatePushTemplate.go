// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdatePushTemplateInput struct {
	_ struct{} `type:"structure" payload:"PushNotificationTemplateRequest"`

	CreateNewVersion *bool `location:"querystring" locationName:"create-new-version" type:"boolean"`

	// Specifies the content and settings for a message template that can be used
	// in messages that are sent through a push notification channel.
	//
	// PushNotificationTemplateRequest is a required field
	PushNotificationTemplateRequest *PushNotificationTemplateRequest `type:"structure" required:"true"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`

	Version *string `location:"querystring" locationName:"version" type:"string"`
}

// String returns the string representation
func (s UpdatePushTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePushTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePushTemplateInput"}

	if s.PushNotificationTemplateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("PushNotificationTemplateRequest"))
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
func (s UpdatePushTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PushNotificationTemplateRequest != nil {
		v := s.PushNotificationTemplateRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PushNotificationTemplateRequest", v, metadata)
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

type UpdatePushTemplateOutput struct {
	_ struct{} `type:"structure" payload:"MessageBody"`

	// Provides information about an API request or response.
	//
	// MessageBody is a required field
	MessageBody *MessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePushTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePushTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MessageBody != nil {
		v := s.MessageBody

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MessageBody", v, metadata)
	}
	return nil
}

const opUpdatePushTemplate = "UpdatePushTemplate"

// UpdatePushTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates an existing message template for messages that are sent through a
// push notification channel.
//
//    // Example sending a request using UpdatePushTemplateRequest.
//    req := client.UpdatePushTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdatePushTemplate
func (c *Client) UpdatePushTemplateRequest(input *UpdatePushTemplateInput) UpdatePushTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdatePushTemplate,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/templates/{template-name}/push",
	}

	if input == nil {
		input = &UpdatePushTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdatePushTemplateOutput{})

	return UpdatePushTemplateRequest{Request: req, Input: input, Copy: c.UpdatePushTemplateRequest}
}

// UpdatePushTemplateRequest is the request type for the
// UpdatePushTemplate API operation.
type UpdatePushTemplateRequest struct {
	*aws.Request
	Input *UpdatePushTemplateInput
	Copy  func(*UpdatePushTemplateInput) UpdatePushTemplateRequest
}

// Send marshals and sends the UpdatePushTemplate API request.
func (r UpdatePushTemplateRequest) Send(ctx context.Context) (*UpdatePushTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePushTemplateResponse{
		UpdatePushTemplateOutput: r.Request.Data.(*UpdatePushTemplateOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePushTemplateResponse is the response type for the
// UpdatePushTemplate API operation.
type UpdatePushTemplateResponse struct {
	*UpdatePushTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePushTemplate request.
func (r *UpdatePushTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
