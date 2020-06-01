// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RemoveFlowSourceInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// SourceArn is a required field
	SourceArn *string `location:"uri" locationName:"sourceArn" type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveFlowSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveFlowSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveFlowSourceInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if s.SourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveFlowSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "sourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful RemoveFlowSource request including the flow ARN
// and the source ARN that was removed.
type RemoveFlowSourceOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that is associated with the source you removed.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The ARN of the source that was removed.
	SourceArn *string `locationName:"sourceArn" type:"string"`
}

// String returns the string representation
func (s RemoveFlowSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveFlowSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceArn != nil {
		v := *s.SourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRemoveFlowSource = "RemoveFlowSource"

// RemoveFlowSourceRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Removes a source from an existing flow. This request can be made only if
// there is more than one source on the flow.
//
//    // Example sending a request using RemoveFlowSourceRequest.
//    req := client.RemoveFlowSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/RemoveFlowSource
func (c *Client) RemoveFlowSourceRequest(input *RemoveFlowSourceInput) RemoveFlowSourceRequest {
	op := &aws.Operation{
		Name:       opRemoveFlowSource,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/flows/{flowArn}/source/{sourceArn}",
	}

	if input == nil {
		input = &RemoveFlowSourceInput{}
	}

	req := c.newRequest(op, input, &RemoveFlowSourceOutput{})

	return RemoveFlowSourceRequest{Request: req, Input: input, Copy: c.RemoveFlowSourceRequest}
}

// RemoveFlowSourceRequest is the request type for the
// RemoveFlowSource API operation.
type RemoveFlowSourceRequest struct {
	*aws.Request
	Input *RemoveFlowSourceInput
	Copy  func(*RemoveFlowSourceInput) RemoveFlowSourceRequest
}

// Send marshals and sends the RemoveFlowSource API request.
func (r RemoveFlowSourceRequest) Send(ctx context.Context) (*RemoveFlowSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveFlowSourceResponse{
		RemoveFlowSourceOutput: r.Request.Data.(*RemoveFlowSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveFlowSourceResponse is the response type for the
// RemoveFlowSource API operation.
type RemoveFlowSourceResponse struct {
	*RemoveFlowSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveFlowSource request.
func (r *RemoveFlowSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
