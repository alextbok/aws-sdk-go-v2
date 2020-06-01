// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to update a data set.
type UpdateDataSetInput struct {
	_ struct{} `type:"structure"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// The description for the data set.
	Description *string `type:"string"`

	// The name of the data set.
	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDataSetInput"}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDataSetOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.
	Arn *string `type:"string"`

	// The type of file your data is stored in. Currently, the supported asset type
	// is S3_SNAPSHOT.
	AssetType AssetType `type:"string" enum:"true"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// A description of a resource.
	Description *string `type:"string"`

	// A unique identifier.
	Id *string `type:"string"`

	// The name of the model.
	Name *string `type:"string"`

	// A property that defines the data set as OWNED by the account (for providers)
	// or ENTITLED to the account (for subscribers). When an owned data set is published
	// in a product, AWS Data Exchange creates a copy of the data set. Subscribers
	// can access that copy of the data set as an entitled data set.
	Origin Origin `type:"string" enum:"true"`

	OriginDetails *OriginDetails `type:"structure"`

	// A unique identifier.
	SourceId *string `type:"string"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	UpdatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s UpdateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.AssetType) > 0 {
		v := s.AssetType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssetType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
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
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Origin) > 0 {
		v := s.Origin

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Origin", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.OriginDetails != nil {
		v := s.OriginDetails

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "OriginDetails", v, metadata)
	}
	if s.SourceId != nil {
		v := *s.SourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdatedAt != nil {
		v := *s.UpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpdatedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

const opUpdateDataSet = "UpdateDataSet"

// UpdateDataSetRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation updates a data set.
//
//    // Example sending a request using UpdateDataSetRequest.
//    req := client.UpdateDataSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/UpdateDataSet
func (c *Client) UpdateDataSetRequest(input *UpdateDataSetInput) UpdateDataSetRequest {
	op := &aws.Operation{
		Name:       opUpdateDataSet,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v1/data-sets/{DataSetId}",
	}

	if input == nil {
		input = &UpdateDataSetInput{}
	}

	req := c.newRequest(op, input, &UpdateDataSetOutput{})

	return UpdateDataSetRequest{Request: req, Input: input, Copy: c.UpdateDataSetRequest}
}

// UpdateDataSetRequest is the request type for the
// UpdateDataSet API operation.
type UpdateDataSetRequest struct {
	*aws.Request
	Input *UpdateDataSetInput
	Copy  func(*UpdateDataSetInput) UpdateDataSetRequest
}

// Send marshals and sends the UpdateDataSet API request.
func (r UpdateDataSetRequest) Send(ctx context.Context) (*UpdateDataSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDataSetResponse{
		UpdateDataSetOutput: r.Request.Data.(*UpdateDataSetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDataSetResponse is the response type for the
// UpdateDataSet API operation.
type UpdateDataSetResponse struct {
	*UpdateDataSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataSet request.
func (r *UpdateDataSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
