// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Input for SetPlatformApplicationAttributes action.
type SetPlatformApplicationAttributesInput struct {
	_ struct{} `type:"structure"`

	// A map of the platform application attributes. Attributes in this map include
	// the following:
	//
	//    * PlatformCredential – The credential received from the notification
	//    service. For APNS/APNS_SANDBOX, PlatformCredential is private key. For
	//    FCM, PlatformCredential is "API key". For ADM, PlatformCredential is "client
	//    secret".
	//
	//    * PlatformPrincipal – The principal received from the notification service.
	//    For APNS/APNS_SANDBOX, PlatformPrincipal is SSL certificate. For FCM,
	//    PlatformPrincipal is not applicable. For ADM, PlatformPrincipal is "client
	//    id".
	//
	//    * EventEndpointCreated – Topic ARN to which EndpointCreated event notifications
	//    should be sent.
	//
	//    * EventEndpointDeleted – Topic ARN to which EndpointDeleted event notifications
	//    should be sent.
	//
	//    * EventEndpointUpdated – Topic ARN to which EndpointUpdate event notifications
	//    should be sent.
	//
	//    * EventDeliveryFailure – Topic ARN to which DeliveryFailure event notifications
	//    should be sent upon Direct Publish delivery failure (permanent) to one
	//    of the application's endpoints.
	//
	//    * SuccessFeedbackRoleArn – IAM role ARN used to give Amazon SNS write
	//    access to use CloudWatch Logs on your behalf.
	//
	//    * FailureFeedbackRoleArn – IAM role ARN used to give Amazon SNS write
	//    access to use CloudWatch Logs on your behalf.
	//
	//    * SuccessFeedbackSampleRate – Sample rate percentage (0-100) of successfully
	//    delivered messages.
	//
	// Attributes is a required field
	Attributes map[string]string `type:"map" required:"true"`

	// PlatformApplicationArn for SetPlatformApplicationAttributes action.
	//
	// PlatformApplicationArn is a required field
	PlatformApplicationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetPlatformApplicationAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetPlatformApplicationAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetPlatformApplicationAttributesInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.PlatformApplicationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformApplicationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetPlatformApplicationAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetPlatformApplicationAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetPlatformApplicationAttributes = "SetPlatformApplicationAttributes"

// SetPlatformApplicationAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Sets the attributes of the platform application object for the supported
// push notification services, such as APNS and FCM. For more information, see
// Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
// For information on configuring attributes for message delivery status, see
// Using Amazon SNS Application Attributes for Message Delivery Status (https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html).
//
//    // Example sending a request using SetPlatformApplicationAttributesRequest.
//    req := client.SetPlatformApplicationAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetPlatformApplicationAttributes
func (c *Client) SetPlatformApplicationAttributesRequest(input *SetPlatformApplicationAttributesInput) SetPlatformApplicationAttributesRequest {
	op := &aws.Operation{
		Name:       opSetPlatformApplicationAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetPlatformApplicationAttributesInput{}
	}

	req := c.newRequest(op, input, &SetPlatformApplicationAttributesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return SetPlatformApplicationAttributesRequest{Request: req, Input: input, Copy: c.SetPlatformApplicationAttributesRequest}
}

// SetPlatformApplicationAttributesRequest is the request type for the
// SetPlatformApplicationAttributes API operation.
type SetPlatformApplicationAttributesRequest struct {
	*aws.Request
	Input *SetPlatformApplicationAttributesInput
	Copy  func(*SetPlatformApplicationAttributesInput) SetPlatformApplicationAttributesRequest
}

// Send marshals and sends the SetPlatformApplicationAttributes API request.
func (r SetPlatformApplicationAttributesRequest) Send(ctx context.Context) (*SetPlatformApplicationAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetPlatformApplicationAttributesResponse{
		SetPlatformApplicationAttributesOutput: r.Request.Data.(*SetPlatformApplicationAttributesOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetPlatformApplicationAttributesResponse is the response type for the
// SetPlatformApplicationAttributes API operation.
type SetPlatformApplicationAttributesResponse struct {
	*SetPlatformApplicationAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetPlatformApplicationAttributes request.
func (r *SetPlatformApplicationAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
