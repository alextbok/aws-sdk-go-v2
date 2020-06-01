// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateEnvironmentInput struct {
	_ struct{} `type:"structure"`

	// The application ID.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// A description of the environment.
	Description *string `type:"string"`

	// The environment ID.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `location:"uri" locationName:"EnvironmentId" type:"string" required:"true"`

	// Amazon CloudWatch alarms to monitor during the deployment process.
	Monitors []Monitor `type:"list"`

	// The name of the environment.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateEnvironmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEnvironmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateEnvironmentInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Monitors != nil {
		for i, v := range s.Monitors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Monitors", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateEnvironmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Monitors != nil {
		v := s.Monitors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Monitors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnvironmentId != nil {
		v := *s.EnvironmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EnvironmentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateEnvironmentOutput struct {
	_ struct{} `type:"structure"`

	// The application ID.
	ApplicationId *string `type:"string"`

	// The description of the environment.
	Description *string `type:"string"`

	// The environment ID.
	Id *string `type:"string"`

	// Amazon CloudWatch alarms monitored during the deployment.
	Monitors []Monitor `type:"list"`

	// The name of the environment.
	Name *string `min:"1" type:"string"`

	// The state of the environment. An environment can be in one of the following
	// states: READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK, or ROLLED_BACK
	State EnvironmentState `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateEnvironmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateEnvironmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Monitors != nil {
		v := s.Monitors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Monitors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opUpdateEnvironment = "UpdateEnvironment"

// UpdateEnvironmentRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Updates an environment.
//
//    // Example sending a request using UpdateEnvironmentRequest.
//    req := client.UpdateEnvironmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/UpdateEnvironment
func (c *Client) UpdateEnvironmentRequest(input *UpdateEnvironmentInput) UpdateEnvironmentRequest {
	op := &aws.Operation{
		Name:       opUpdateEnvironment,
		HTTPMethod: "PATCH",
		HTTPPath:   "/applications/{ApplicationId}/environments/{EnvironmentId}",
	}

	if input == nil {
		input = &UpdateEnvironmentInput{}
	}

	req := c.newRequest(op, input, &UpdateEnvironmentOutput{})

	return UpdateEnvironmentRequest{Request: req, Input: input, Copy: c.UpdateEnvironmentRequest}
}

// UpdateEnvironmentRequest is the request type for the
// UpdateEnvironment API operation.
type UpdateEnvironmentRequest struct {
	*aws.Request
	Input *UpdateEnvironmentInput
	Copy  func(*UpdateEnvironmentInput) UpdateEnvironmentRequest
}

// Send marshals and sends the UpdateEnvironment API request.
func (r UpdateEnvironmentRequest) Send(ctx context.Context) (*UpdateEnvironmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEnvironmentResponse{
		UpdateEnvironmentOutput: r.Request.Data.(*UpdateEnvironmentOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEnvironmentResponse is the response type for the
// UpdateEnvironment API operation.
type UpdateEnvironmentResponse struct {
	*UpdateEnvironmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEnvironment request.
func (r *UpdateEnvironmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
