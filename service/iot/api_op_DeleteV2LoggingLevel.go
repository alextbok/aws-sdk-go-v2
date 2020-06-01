// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteV2LoggingLevelInput struct {
	_ struct{} `type:"structure"`

	// The name of the resource for which you are configuring logging.
	//
	// TargetName is a required field
	TargetName *string `location:"querystring" locationName:"targetName" type:"string" required:"true"`

	// The type of resource for which you are configuring logging. Must be THING_Group.
	//
	// TargetType is a required field
	TargetType LogTargetType `location:"querystring" locationName:"targetType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DeleteV2LoggingLevelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteV2LoggingLevelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteV2LoggingLevelInput"}

	if s.TargetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetName"))
	}
	if len(s.TargetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TargetType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteV2LoggingLevelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TargetName != nil {
		v := *s.TargetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "targetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TargetType) > 0 {
		v := s.TargetType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "targetType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type DeleteV2LoggingLevelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteV2LoggingLevelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteV2LoggingLevelOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteV2LoggingLevel = "DeleteV2LoggingLevel"

// DeleteV2LoggingLevelRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a logging level.
//
//    // Example sending a request using DeleteV2LoggingLevelRequest.
//    req := client.DeleteV2LoggingLevelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteV2LoggingLevelRequest(input *DeleteV2LoggingLevelInput) DeleteV2LoggingLevelRequest {
	op := &aws.Operation{
		Name:       opDeleteV2LoggingLevel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2LoggingLevel",
	}

	if input == nil {
		input = &DeleteV2LoggingLevelInput{}
	}

	req := c.newRequest(op, input, &DeleteV2LoggingLevelOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteV2LoggingLevelRequest{Request: req, Input: input, Copy: c.DeleteV2LoggingLevelRequest}
}

// DeleteV2LoggingLevelRequest is the request type for the
// DeleteV2LoggingLevel API operation.
type DeleteV2LoggingLevelRequest struct {
	*aws.Request
	Input *DeleteV2LoggingLevelInput
	Copy  func(*DeleteV2LoggingLevelInput) DeleteV2LoggingLevelRequest
}

// Send marshals and sends the DeleteV2LoggingLevel API request.
func (r DeleteV2LoggingLevelRequest) Send(ctx context.Context) (*DeleteV2LoggingLevelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteV2LoggingLevelResponse{
		DeleteV2LoggingLevelOutput: r.Request.Data.(*DeleteV2LoggingLevelOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteV2LoggingLevelResponse is the response type for the
// DeleteV2LoggingLevel API operation.
type DeleteV2LoggingLevelResponse struct {
	*DeleteV2LoggingLevelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteV2LoggingLevel request.
func (r *DeleteV2LoggingLevelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
