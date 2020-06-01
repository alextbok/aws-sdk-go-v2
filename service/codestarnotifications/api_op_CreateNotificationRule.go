// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateNotificationRuleInput struct {
	_ struct{} `type:"structure"`

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request
	// with the same parameters is received and a token is included, the request
	// returns information about the initial request that used that token.
	//
	// The AWS SDKs prepopulate client request tokens. If you are using an AWS SDK,
	// an idempotency token is created for you.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The level of detail to include in the notifications for this resource. BASIC
	// will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications
	// and/or the service for the resource for which the notification is created.
	//
	// DetailType is a required field
	DetailType DetailType `type:"string" required:"true" enum:"true"`

	// A list of event types associated with this notification rule. For a list
	// of allowed events, see EventTypeSummary.
	//
	// EventTypeIds is a required field
	EventTypeIds []string `type:"list" required:"true"`

	// The name for the notification rule. Notifictaion rule names must be unique
	// in your AWS account.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// The Amazon Resource Name (ARN) of the resource to associate with the notification
	// rule. Supported resources include pipelines in AWS CodePipeline, repositories
	// in AWS CodeCommit, and build projects in AWS CodeBuild.
	//
	// Resource is a required field
	Resource *string `type:"string" required:"true"`

	// The status of the notification rule. The default value is ENABLED. If the
	// status is set to DISABLED, notifications aren't sent for the notification
	// rule.
	Status NotificationRuleStatus `type:"string" enum:"true"`

	// A list of tags to apply to this notification rule. Key names cannot start
	// with "aws".
	Tags map[string]string `type:"map"`

	// A list of Amazon Resource Names (ARNs) of SNS topics to associate with the
	// notification rule.
	//
	// Targets is a required field
	Targets []Target `type:"list" required:"true"`
}

// String returns the string representation
func (s CreateNotificationRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNotificationRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNotificationRuleInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if len(s.DetailType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DetailType"))
	}

	if s.EventTypeIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventTypeIds"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Resource == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resource"))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateNotificationRuleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DetailType) > 0 {
		v := s.DetailType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DetailType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EventTypeIds != nil {
		v := s.EventTypeIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "EventTypeIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Resource != nil {
		v := *s.Resource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Resource", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Targets != nil {
		v := s.Targets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Targets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type CreateNotificationRuleOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string `type:"string"`
}

// String returns the string representation
func (s CreateNotificationRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateNotificationRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateNotificationRule = "CreateNotificationRule"

// CreateNotificationRuleRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Creates a notification rule for a resource. The rule specifies the events
// you want notifications about and the targets (such as SNS topics) where you
// want to receive them.
//
//    // Example sending a request using CreateNotificationRuleRequest.
//    req := client.CreateNotificationRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/CreateNotificationRule
func (c *Client) CreateNotificationRuleRequest(input *CreateNotificationRuleInput) CreateNotificationRuleRequest {
	op := &aws.Operation{
		Name:       opCreateNotificationRule,
		HTTPMethod: "POST",
		HTTPPath:   "/createNotificationRule",
	}

	if input == nil {
		input = &CreateNotificationRuleInput{}
	}

	req := c.newRequest(op, input, &CreateNotificationRuleOutput{})

	return CreateNotificationRuleRequest{Request: req, Input: input, Copy: c.CreateNotificationRuleRequest}
}

// CreateNotificationRuleRequest is the request type for the
// CreateNotificationRule API operation.
type CreateNotificationRuleRequest struct {
	*aws.Request
	Input *CreateNotificationRuleInput
	Copy  func(*CreateNotificationRuleInput) CreateNotificationRuleRequest
}

// Send marshals and sends the CreateNotificationRule API request.
func (r CreateNotificationRuleRequest) Send(ctx context.Context) (*CreateNotificationRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNotificationRuleResponse{
		CreateNotificationRuleOutput: r.Request.Data.(*CreateNotificationRuleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNotificationRuleResponse is the response type for the
// CreateNotificationRule API operation.
type CreateNotificationRuleResponse struct {
	*CreateNotificationRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNotificationRule request.
func (r *CreateNotificationRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
