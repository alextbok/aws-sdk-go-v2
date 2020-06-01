// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribePackagingGroupInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePackagingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePackagingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePackagingGroupInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePackagingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribePackagingGroupOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	DomainName *string `locationName:"domainName" type:"string"`

	Id *string `locationName:"id" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribePackagingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePackagingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opDescribePackagingGroup = "DescribePackagingGroup"

// DescribePackagingGroupRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a description of a MediaPackage VOD PackagingGroup resource.
//
//    // Example sending a request using DescribePackagingGroupRequest.
//    req := client.DescribePackagingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DescribePackagingGroup
func (c *Client) DescribePackagingGroupRequest(input *DescribePackagingGroupInput) DescribePackagingGroupRequest {
	op := &aws.Operation{
		Name:       opDescribePackagingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/packaging_groups/{id}",
	}

	if input == nil {
		input = &DescribePackagingGroupInput{}
	}

	req := c.newRequest(op, input, &DescribePackagingGroupOutput{})

	return DescribePackagingGroupRequest{Request: req, Input: input, Copy: c.DescribePackagingGroupRequest}
}

// DescribePackagingGroupRequest is the request type for the
// DescribePackagingGroup API operation.
type DescribePackagingGroupRequest struct {
	*aws.Request
	Input *DescribePackagingGroupInput
	Copy  func(*DescribePackagingGroupInput) DescribePackagingGroupRequest
}

// Send marshals and sends the DescribePackagingGroup API request.
func (r DescribePackagingGroupRequest) Send(ctx context.Context) (*DescribePackagingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePackagingGroupResponse{
		DescribePackagingGroupOutput: r.Request.Data.(*DescribePackagingGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePackagingGroupResponse is the response type for the
// DescribePackagingGroup API operation.
type DescribePackagingGroupResponse struct {
	*DescribePackagingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePackagingGroup request.
func (r *DescribePackagingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
