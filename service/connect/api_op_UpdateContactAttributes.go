// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateContactAttributesInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Connect attributes. These attributes can be accessed in contact
	// flows just like any other contact attributes.
	//
	// You can have up to 32,768 UTF-8 bytes across all attributes for a contact.
	// Attribute keys can include only alphanumeric, dash, and underscore characters.
	//
	// Attributes is a required field
	Attributes map[string]string `type:"map" required:"true"`

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with the contact center.
	//
	// InitialContactId is a required field
	InitialContactId *string `min:"1" type:"string" required:"true"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateContactAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateContactAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateContactAttributesInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.InitialContactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InitialContactId"))
	}
	if s.InitialContactId != nil && len(*s.InitialContactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InitialContactId", 1))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateContactAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Attributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.InitialContactId != nil {
		v := *s.InitialContactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InitialContactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateContactAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateContactAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateContactAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateContactAttributes = "UpdateContactAttributes"

// UpdateContactAttributesRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Creates or updates the contact attributes associated with the specified contact.
//
// You can add or update attributes for both ongoing and completed contacts.
// For example, you can update the customer's name or the reason the customer
// called while the call is active, or add notes about steps that the agent
// took during the call that are displayed to the next agent that takes the
// call. You can also update attributes for a contact using data from your CRM
// application and save the data with the contact in Amazon Connect. You could
// also flag calls for additional analysis, such as legal review or identifying
// abusive callers.
//
// Contact attributes are available in Amazon Connect for 24 months, and are
// then deleted.
//
// Important: You cannot use the operation to update attributes for contacts
// that occurred prior to the release of the API, September 12, 2018. You can
// update attributes only for contacts that started after the release of the
// API. If you attempt to update attributes for a contact that occurred prior
// to the release of the API, a 400 error is returned. This applies also to
// queued callbacks that were initiated prior to the release of the API but
// are still active in your instance.
//
//    // Example sending a request using UpdateContactAttributesRequest.
//    req := client.UpdateContactAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/UpdateContactAttributes
func (c *Client) UpdateContactAttributesRequest(input *UpdateContactAttributesInput) UpdateContactAttributesRequest {
	op := &aws.Operation{
		Name:       opUpdateContactAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/contact/attributes",
	}

	if input == nil {
		input = &UpdateContactAttributesInput{}
	}

	req := c.newRequest(op, input, &UpdateContactAttributesOutput{})

	return UpdateContactAttributesRequest{Request: req, Input: input, Copy: c.UpdateContactAttributesRequest}
}

// UpdateContactAttributesRequest is the request type for the
// UpdateContactAttributes API operation.
type UpdateContactAttributesRequest struct {
	*aws.Request
	Input *UpdateContactAttributesInput
	Copy  func(*UpdateContactAttributesInput) UpdateContactAttributesRequest
}

// Send marshals and sends the UpdateContactAttributes API request.
func (r UpdateContactAttributesRequest) Send(ctx context.Context) (*UpdateContactAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateContactAttributesResponse{
		UpdateContactAttributesOutput: r.Request.Data.(*UpdateContactAttributesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateContactAttributesResponse is the response type for the
// UpdateContactAttributes API operation.
type UpdateContactAttributesResponse struct {
	*UpdateContactAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateContactAttributes request.
func (r *UpdateContactAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
