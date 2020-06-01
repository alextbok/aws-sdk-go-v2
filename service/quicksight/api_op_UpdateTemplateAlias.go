// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateTemplateAliasInput struct {
	_ struct{} `type:"structure"`

	// The alias of the template that you want to update. If you name a specific
	// alias, you update the version that the alias points to. You can specify the
	// latest version of the template by providing the keyword $LATEST in the AliasName
	// parameter. The keyword $PUBLISHED doesn't apply to templates.
	//
	// AliasName is a required field
	AliasName *string `location:"uri" locationName:"AliasName" min:"1" type:"string" required:"true"`

	// The ID of the AWS account that contains the template alias that you're updating.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the template.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`

	// The version number of the template.
	//
	// TemplateVersionNumber is a required field
	TemplateVersionNumber *int64 `min:"1" type:"long" required:"true"`
}

// String returns the string representation
func (s UpdateTemplateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTemplateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTemplateAliasInput"}

	if s.AliasName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasName"))
	}
	if s.AliasName != nil && len(*s.AliasName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AliasName", 1))
	}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}

	if s.TemplateVersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateVersionNumber"))
	}
	if s.TemplateVersionNumber != nil && *s.TemplateVersionNumber < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("TemplateVersionNumber", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTemplateAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateVersionNumber != nil {
		v := *s.TemplateVersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateVersionNumber", protocol.Int64Value(v), metadata)
	}
	if s.AliasName != nil {
		v := *s.AliasName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AliasName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateTemplateAliasOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The template alias.
	TemplateAlias *TemplateAlias `type:"structure"`
}

// String returns the string representation
func (s UpdateTemplateAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTemplateAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateAlias != nil {
		v := s.TemplateAlias

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TemplateAlias", v, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateTemplateAlias = "UpdateTemplateAlias"

// UpdateTemplateAliasRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates the template alias of a template.
//
//    // Example sending a request using UpdateTemplateAliasRequest.
//    req := client.UpdateTemplateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateTemplateAlias
func (c *Client) UpdateTemplateAliasRequest(input *UpdateTemplateAliasInput) UpdateTemplateAliasRequest {
	op := &aws.Operation{
		Name:       opUpdateTemplateAlias,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/aliases/{AliasName}",
	}

	if input == nil {
		input = &UpdateTemplateAliasInput{}
	}

	req := c.newRequest(op, input, &UpdateTemplateAliasOutput{})

	return UpdateTemplateAliasRequest{Request: req, Input: input, Copy: c.UpdateTemplateAliasRequest}
}

// UpdateTemplateAliasRequest is the request type for the
// UpdateTemplateAlias API operation.
type UpdateTemplateAliasRequest struct {
	*aws.Request
	Input *UpdateTemplateAliasInput
	Copy  func(*UpdateTemplateAliasInput) UpdateTemplateAliasRequest
}

// Send marshals and sends the UpdateTemplateAlias API request.
func (r UpdateTemplateAliasRequest) Send(ctx context.Context) (*UpdateTemplateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTemplateAliasResponse{
		UpdateTemplateAliasOutput: r.Request.Data.(*UpdateTemplateAliasOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTemplateAliasResponse is the response type for the
// UpdateTemplateAlias API operation.
type UpdateTemplateAliasResponse struct {
	*UpdateTemplateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTemplateAlias request.
func (r *UpdateTemplateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
