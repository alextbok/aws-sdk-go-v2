// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type AttachPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// both objects reside. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The reference that identifies the object to which the policy will be attached.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`

	// The reference that is associated with the policy object.
	//
	// PolicyReference is a required field
	PolicyReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s AttachPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachPolicyInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if s.PolicyReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if s.PolicyReference != nil {
		v := s.PolicyReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PolicyReference", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AttachPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAttachPolicy = "AttachPolicy"

// AttachPolicyRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Attaches a policy object to a regular object. An object can have a limited
// number of attached policies.
//
//    // Example sending a request using AttachPolicyRequest.
//    req := client.AttachPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachPolicy
func (c *Client) AttachPolicyRequest(input *AttachPolicyInput) AttachPolicyRequest {
	op := &aws.Operation{
		Name:       opAttachPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/policy/attach",
	}

	if input == nil {
		input = &AttachPolicyInput{}
	}

	req := c.newRequest(op, input, &AttachPolicyOutput{})

	return AttachPolicyRequest{Request: req, Input: input, Copy: c.AttachPolicyRequest}
}

// AttachPolicyRequest is the request type for the
// AttachPolicy API operation.
type AttachPolicyRequest struct {
	*aws.Request
	Input *AttachPolicyInput
	Copy  func(*AttachPolicyInput) AttachPolicyRequest
}

// Send marshals and sends the AttachPolicy API request.
func (r AttachPolicyRequest) Send(ctx context.Context) (*AttachPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachPolicyResponse{
		AttachPolicyOutput: r.Request.Data.(*AttachPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachPolicyResponse is the response type for the
// AttachPolicy API operation.
type AttachPolicyResponse struct {
	*AttachPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachPolicy request.
func (r *AttachPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
