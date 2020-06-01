// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the GetPolicyVersion operation.
type GetPolicyVersionInput struct {
	_ struct{} `type:"structure"`

	// The name of the policy.
	//
	// PolicyName is a required field
	PolicyName *string `location:"uri" locationName:"policyName" min:"1" type:"string" required:"true"`

	// The policy version ID.
	//
	// PolicyVersionId is a required field
	PolicyVersionId *string `location:"uri" locationName:"policyVersionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPolicyVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPolicyVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPolicyVersionInput"}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if s.PolicyVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyVersionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPolicyVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PolicyName != nil {
		v := *s.PolicyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "policyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyVersionId != nil {
		v := *s.PolicyVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "policyVersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the GetPolicyVersion operation.
type GetPolicyVersionOutput struct {
	_ struct{} `type:"structure"`

	// The date the policy was created.
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp"`

	// The generation ID of the policy version.
	GenerationId *string `locationName:"generationId" type:"string"`

	// Specifies whether the policy version is the default.
	IsDefaultVersion *bool `locationName:"isDefaultVersion" type:"boolean"`

	// The date the policy was last modified.
	LastModifiedDate *time.Time `locationName:"lastModifiedDate" type:"timestamp"`

	// The policy ARN.
	PolicyArn *string `locationName:"policyArn" type:"string"`

	// The JSON document that describes the policy.
	PolicyDocument *string `locationName:"policyDocument" type:"string"`

	// The policy name.
	PolicyName *string `locationName:"policyName" min:"1" type:"string"`

	// The policy version ID.
	PolicyVersionId *string `locationName:"policyVersionId" type:"string"`
}

// String returns the string representation
func (s GetPolicyVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPolicyVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.GenerationId != nil {
		v := *s.GenerationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "generationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IsDefaultVersion != nil {
		v := *s.IsDefaultVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "isDefaultVersion", protocol.BoolValue(v), metadata)
	}
	if s.LastModifiedDate != nil {
		v := *s.LastModifiedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastModifiedDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.PolicyArn != nil {
		v := *s.PolicyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyDocument != nil {
		v := *s.PolicyDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyDocument", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyName != nil {
		v := *s.PolicyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyVersionId != nil {
		v := *s.PolicyVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyVersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetPolicyVersion = "GetPolicyVersion"

// GetPolicyVersionRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the specified policy version.
//
//    // Example sending a request using GetPolicyVersionRequest.
//    req := client.GetPolicyVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetPolicyVersionRequest(input *GetPolicyVersionInput) GetPolicyVersionRequest {
	op := &aws.Operation{
		Name:       opGetPolicyVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/policies/{policyName}/version/{policyVersionId}",
	}

	if input == nil {
		input = &GetPolicyVersionInput{}
	}

	req := c.newRequest(op, input, &GetPolicyVersionOutput{})

	return GetPolicyVersionRequest{Request: req, Input: input, Copy: c.GetPolicyVersionRequest}
}

// GetPolicyVersionRequest is the request type for the
// GetPolicyVersion API operation.
type GetPolicyVersionRequest struct {
	*aws.Request
	Input *GetPolicyVersionInput
	Copy  func(*GetPolicyVersionInput) GetPolicyVersionRequest
}

// Send marshals and sends the GetPolicyVersion API request.
func (r GetPolicyVersionRequest) Send(ctx context.Context) (*GetPolicyVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPolicyVersionResponse{
		GetPolicyVersionOutput: r.Request.Data.(*GetPolicyVersionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPolicyVersionResponse is the response type for the
// GetPolicyVersion API operation.
type GetPolicyVersionResponse struct {
	*GetPolicyVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPolicyVersion request.
func (r *GetPolicyVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
