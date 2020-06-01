// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteGatewayInput struct {
	_ struct{} `type:"structure"`

	// The ID of the gateway to delete.
	//
	// GatewayId is a required field
	GatewayId *string `location:"uri" locationName:"gatewayId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGatewayInput"}

	if s.GatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayId"))
	}
	if s.GatewayId != nil && len(*s.GatewayId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGatewayInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GatewayId != nil {
		v := *s.GatewayId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "gatewayId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteGatewayOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGatewayOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteGateway = "DeleteGateway"

// DeleteGatewayRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Deletes a gateway from AWS IoT SiteWise. When you delete a gateway, some
// of the gateway's files remain in your gateway's file system. For more information,
// see Data retention (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/data-retention.html)
// in the AWS IoT SiteWise User Guide.
//
//    // Example sending a request using DeleteGatewayRequest.
//    req := client.DeleteGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/DeleteGateway
func (c *Client) DeleteGatewayRequest(input *DeleteGatewayInput) DeleteGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteGateway,
		HTTPMethod: "DELETE",
		HTTPPath:   "/20200301/gateways/{gatewayId}",
	}

	if input == nil {
		input = &DeleteGatewayInput{}
	}

	req := c.newRequest(op, input, &DeleteGatewayOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("edge.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return DeleteGatewayRequest{Request: req, Input: input, Copy: c.DeleteGatewayRequest}
}

// DeleteGatewayRequest is the request type for the
// DeleteGateway API operation.
type DeleteGatewayRequest struct {
	*aws.Request
	Input *DeleteGatewayInput
	Copy  func(*DeleteGatewayInput) DeleteGatewayRequest
}

// Send marshals and sends the DeleteGateway API request.
func (r DeleteGatewayRequest) Send(ctx context.Context) (*DeleteGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGatewayResponse{
		DeleteGatewayOutput: r.Request.Data.(*DeleteGatewayOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGatewayResponse is the response type for the
// DeleteGateway API operation.
type DeleteGatewayResponse struct {
	*DeleteGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGateway request.
func (r *DeleteGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
