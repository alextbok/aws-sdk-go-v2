// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type AttachToIndexInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the directory where the object and index
	// exist.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// A reference to the index that you are attaching the object to.
	//
	// IndexReference is a required field
	IndexReference *ObjectReference `type:"structure" required:"true"`

	// A reference to the object that you are attaching to the index.
	//
	// TargetReference is a required field
	TargetReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s AttachToIndexInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachToIndexInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachToIndexInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.IndexReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexReference"))
	}

	if s.TargetReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachToIndexInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IndexReference != nil {
		v := s.IndexReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "IndexReference", v, metadata)
	}
	if s.TargetReference != nil {
		v := s.TargetReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TargetReference", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AttachToIndexOutput struct {
	_ struct{} `type:"structure"`

	// The ObjectIdentifier of the object that was attached to the index.
	AttachedObjectIdentifier *string `type:"string"`
}

// String returns the string representation
func (s AttachToIndexOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachToIndexOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AttachedObjectIdentifier != nil {
		v := *s.AttachedObjectIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AttachedObjectIdentifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opAttachToIndex = "AttachToIndex"

// AttachToIndexRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Attaches the specified object to the specified index.
//
//    // Example sending a request using AttachToIndexRequest.
//    req := client.AttachToIndexRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachToIndex
func (c *Client) AttachToIndexRequest(input *AttachToIndexInput) AttachToIndexRequest {
	op := &aws.Operation{
		Name:       opAttachToIndex,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/index/attach",
	}

	if input == nil {
		input = &AttachToIndexInput{}
	}

	req := c.newRequest(op, input, &AttachToIndexOutput{})

	return AttachToIndexRequest{Request: req, Input: input, Copy: c.AttachToIndexRequest}
}

// AttachToIndexRequest is the request type for the
// AttachToIndex API operation.
type AttachToIndexRequest struct {
	*aws.Request
	Input *AttachToIndexInput
	Copy  func(*AttachToIndexInput) AttachToIndexRequest
}

// Send marshals and sends the AttachToIndex API request.
func (r AttachToIndexRequest) Send(ctx context.Context) (*AttachToIndexResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachToIndexResponse{
		AttachToIndexOutput: r.Request.Data.(*AttachToIndexOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachToIndexResponse is the response type for the
// AttachToIndex API operation.
type AttachToIndexResponse struct {
	*AttachToIndexOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachToIndex request.
func (r *AttachToIndexResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
