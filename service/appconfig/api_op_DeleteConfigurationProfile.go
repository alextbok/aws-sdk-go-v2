// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteConfigurationProfileInput struct {
	_ struct{} `type:"structure"`

	// The application ID that includes the configuration profile you want to delete.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// The ID of the configuration profile you want to delete.
	//
	// ConfigurationProfileId is a required field
	ConfigurationProfileId *string `location:"uri" locationName:"ConfigurationProfileId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationProfileInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.ConfigurationProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationProfileId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationProfileId != nil {
		v := *s.ConfigurationProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteConfigurationProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteConfigurationProfile = "DeleteConfigurationProfile"

// DeleteConfigurationProfileRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Delete a configuration profile. Deleting a configuration profile does not
// delete a configuration from a host.
//
//    // Example sending a request using DeleteConfigurationProfileRequest.
//    req := client.DeleteConfigurationProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/DeleteConfigurationProfile
func (c *Client) DeleteConfigurationProfileRequest(input *DeleteConfigurationProfileInput) DeleteConfigurationProfileRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationProfile,
		HTTPMethod: "DELETE",
		HTTPPath:   "/applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}",
	}

	if input == nil {
		input = &DeleteConfigurationProfileInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationProfileOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteConfigurationProfileRequest{Request: req, Input: input, Copy: c.DeleteConfigurationProfileRequest}
}

// DeleteConfigurationProfileRequest is the request type for the
// DeleteConfigurationProfile API operation.
type DeleteConfigurationProfileRequest struct {
	*aws.Request
	Input *DeleteConfigurationProfileInput
	Copy  func(*DeleteConfigurationProfileInput) DeleteConfigurationProfileRequest
}

// Send marshals and sends the DeleteConfigurationProfile API request.
func (r DeleteConfigurationProfileRequest) Send(ctx context.Context) (*DeleteConfigurationProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationProfileResponse{
		DeleteConfigurationProfileOutput: r.Request.Data.(*DeleteConfigurationProfileOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationProfileResponse is the response type for the
// DeleteConfigurationProfile API operation.
type DeleteConfigurationProfileResponse struct {
	*DeleteConfigurationProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationProfile request.
func (r *DeleteConfigurationProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
