// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeBillingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the billing group.
	//
	// BillingGroupName is a required field
	BillingGroupName *string `location:"uri" locationName:"billingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBillingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBillingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBillingGroupInput"}

	if s.BillingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BillingGroupName"))
	}
	if s.BillingGroupName != nil && len(*s.BillingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BillingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBillingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeBillingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the billing group.
	BillingGroupArn *string `locationName:"billingGroupArn" type:"string"`

	// The ID of the billing group.
	BillingGroupId *string `locationName:"billingGroupId" min:"1" type:"string"`

	// Additional information about the billing group.
	BillingGroupMetadata *BillingGroupMetadata `locationName:"billingGroupMetadata" type:"structure"`

	// The name of the billing group.
	BillingGroupName *string `locationName:"billingGroupName" min:"1" type:"string"`

	// The properties of the billing group.
	BillingGroupProperties *BillingGroupProperties `locationName:"billingGroupProperties" type:"structure"`

	// The version of the billing group.
	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s DescribeBillingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBillingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BillingGroupArn != nil {
		v := *s.BillingGroupArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BillingGroupId != nil {
		v := *s.BillingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BillingGroupMetadata != nil {
		v := s.BillingGroupMetadata

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "billingGroupMetadata", v, metadata)
	}
	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BillingGroupProperties != nil {
		v := s.BillingGroupProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "billingGroupProperties", v, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opDescribeBillingGroup = "DescribeBillingGroup"

// DescribeBillingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Returns information about a billing group.
//
//    // Example sending a request using DescribeBillingGroupRequest.
//    req := client.DescribeBillingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeBillingGroupRequest(input *DescribeBillingGroupInput) DescribeBillingGroupRequest {
	op := &aws.Operation{
		Name:       opDescribeBillingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/billing-groups/{billingGroupName}",
	}

	if input == nil {
		input = &DescribeBillingGroupInput{}
	}

	req := c.newRequest(op, input, &DescribeBillingGroupOutput{})

	return DescribeBillingGroupRequest{Request: req, Input: input, Copy: c.DescribeBillingGroupRequest}
}

// DescribeBillingGroupRequest is the request type for the
// DescribeBillingGroup API operation.
type DescribeBillingGroupRequest struct {
	*aws.Request
	Input *DescribeBillingGroupInput
	Copy  func(*DescribeBillingGroupInput) DescribeBillingGroupRequest
}

// Send marshals and sends the DescribeBillingGroup API request.
func (r DescribeBillingGroupRequest) Send(ctx context.Context) (*DescribeBillingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBillingGroupResponse{
		DescribeBillingGroupOutput: r.Request.Data.(*DescribeBillingGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBillingGroupResponse is the response type for the
// DescribeBillingGroup API operation.
type DescribeBillingGroupResponse struct {
	*DescribeBillingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBillingGroup request.
func (r *DescribeBillingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
