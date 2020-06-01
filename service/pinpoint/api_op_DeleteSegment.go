// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteSegmentInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// SegmentId is a required field
	SegmentId *string `location:"uri" locationName:"segment-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSegmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSegmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSegmentInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.SegmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SegmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSegmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SegmentId != nil {
		v := *s.SegmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "segment-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSegmentOutput struct {
	_ struct{} `type:"structure" payload:"SegmentResponse"`

	// Provides information about the configuration, dimension, and other settings
	// for a segment.
	//
	// SegmentResponse is a required field
	SegmentResponse *SegmentResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteSegmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSegmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SegmentResponse != nil {
		v := s.SegmentResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SegmentResponse", v, metadata)
	}
	return nil
}

const opDeleteSegment = "DeleteSegment"

// DeleteSegmentRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes a segment from an application.
//
//    // Example sending a request using DeleteSegmentRequest.
//    req := client.DeleteSegmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteSegment
func (c *Client) DeleteSegmentRequest(input *DeleteSegmentInput) DeleteSegmentRequest {
	op := &aws.Operation{
		Name:       opDeleteSegment,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/segments/{segment-id}",
	}

	if input == nil {
		input = &DeleteSegmentInput{}
	}

	req := c.newRequest(op, input, &DeleteSegmentOutput{})

	return DeleteSegmentRequest{Request: req, Input: input, Copy: c.DeleteSegmentRequest}
}

// DeleteSegmentRequest is the request type for the
// DeleteSegment API operation.
type DeleteSegmentRequest struct {
	*aws.Request
	Input *DeleteSegmentInput
	Copy  func(*DeleteSegmentInput) DeleteSegmentRequest
}

// Send marshals and sends the DeleteSegment API request.
func (r DeleteSegmentRequest) Send(ctx context.Context) (*DeleteSegmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSegmentResponse{
		DeleteSegmentOutput: r.Request.Data.(*DeleteSegmentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSegmentResponse is the response type for the
// DeleteSegment API operation.
type DeleteSegmentResponse struct {
	*DeleteSegmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSegment request.
func (r *DeleteSegmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
