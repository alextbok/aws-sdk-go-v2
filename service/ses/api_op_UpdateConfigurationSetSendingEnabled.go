// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Represents a request to enable or disable the email sending capabilities
// for a specific configuration set.
type UpdateConfigurationSetSendingEnabledInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to update.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// Describes whether email sending is enabled or disabled for the configuration
	// set.
	//
	// Enabled is a required field
	Enabled *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s UpdateConfigurationSetSendingEnabledInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigurationSetSendingEnabledInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigurationSetSendingEnabledInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateConfigurationSetSendingEnabledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConfigurationSetSendingEnabledOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateConfigurationSetSendingEnabled = "UpdateConfigurationSetSendingEnabled"

// UpdateConfigurationSetSendingEnabledRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Enables or disables email sending for messages sent using a specific configuration
// set in a given AWS Region. You can use this operation in conjunction with
// Amazon CloudWatch alarms to temporarily pause email sending for a configuration
// set when the reputation metrics for that configuration set (such as your
// bounce on complaint rate) exceed certain thresholds.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using UpdateConfigurationSetSendingEnabledRequest.
//    req := client.UpdateConfigurationSetSendingEnabledRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/UpdateConfigurationSetSendingEnabled
func (c *Client) UpdateConfigurationSetSendingEnabledRequest(input *UpdateConfigurationSetSendingEnabledInput) UpdateConfigurationSetSendingEnabledRequest {
	op := &aws.Operation{
		Name:       opUpdateConfigurationSetSendingEnabled,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateConfigurationSetSendingEnabledInput{}
	}

	req := c.newRequest(op, input, &UpdateConfigurationSetSendingEnabledOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UpdateConfigurationSetSendingEnabledRequest{Request: req, Input: input, Copy: c.UpdateConfigurationSetSendingEnabledRequest}
}

// UpdateConfigurationSetSendingEnabledRequest is the request type for the
// UpdateConfigurationSetSendingEnabled API operation.
type UpdateConfigurationSetSendingEnabledRequest struct {
	*aws.Request
	Input *UpdateConfigurationSetSendingEnabledInput
	Copy  func(*UpdateConfigurationSetSendingEnabledInput) UpdateConfigurationSetSendingEnabledRequest
}

// Send marshals and sends the UpdateConfigurationSetSendingEnabled API request.
func (r UpdateConfigurationSetSendingEnabledRequest) Send(ctx context.Context) (*UpdateConfigurationSetSendingEnabledResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConfigurationSetSendingEnabledResponse{
		UpdateConfigurationSetSendingEnabledOutput: r.Request.Data.(*UpdateConfigurationSetSendingEnabledOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConfigurationSetSendingEnabledResponse is the response type for the
// UpdateConfigurationSetSendingEnabled API operation.
type UpdateConfigurationSetSendingEnabledResponse struct {
	*UpdateConfigurationSetSendingEnabledOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConfigurationSetSendingEnabled request.
func (r *UpdateConfigurationSetSendingEnabledResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
