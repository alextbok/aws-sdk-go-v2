// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// The request to delete a streaming distribution.
type DeleteStreamingDistributionInput struct {
	_ struct{} `type:"structure"`

	// The distribution ID.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when you disabled the streaming
	// distribution. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s DeleteStreamingDistributionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteStreamingDistributionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteStreamingDistributionInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteStreamingDistributionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteStreamingDistributionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteStreamingDistributionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteStreamingDistributionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteStreamingDistribution = "DeleteStreamingDistribution2019_03_26"

// DeleteStreamingDistributionRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Delete a streaming distribution. To delete an RTMP distribution using the
// CloudFront API, perform the following steps.
//
// To delete an RTMP distribution using the CloudFront API:
//
// Disable the RTMP distribution.
//
// Submit a GET Streaming Distribution Config request to get the current configuration
// and the Etag header for the distribution.
//
// Update the XML document that was returned in the response to your GET Streaming
// Distribution Config request to change the value of Enabled to false.
//
// Submit a PUT Streaming Distribution Config request to update the configuration
// for your distribution. In the request body, include the XML document that
// you updated in Step 3. Then set the value of the HTTP If-Match header to
// the value of the ETag header that CloudFront returned when you submitted
// the GET Streaming Distribution Config request in Step 2.
//
// Review the response to the PUT Streaming Distribution Config request to confirm
// that the distribution was successfully disabled.
//
// Submit a GET Streaming Distribution Config request to confirm that your changes
// have propagated. When propagation is complete, the value of Status is Deployed.
//
// Submit a DELETE Streaming Distribution request. Set the value of the HTTP
// If-Match header to the value of the ETag header that CloudFront returned
// when you submitted the GET Streaming Distribution Config request in Step
// 2.
//
// Review the response to your DELETE Streaming Distribution request to confirm
// that the distribution was successfully deleted.
//
// For information about deleting a distribution using the CloudFront console,
// see Deleting a Distribution (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/HowToDeleteDistribution.html)
// in the Amazon CloudFront Developer Guide.
//
//    // Example sending a request using DeleteStreamingDistributionRequest.
//    req := client.DeleteStreamingDistributionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/DeleteStreamingDistribution
func (c *Client) DeleteStreamingDistributionRequest(input *DeleteStreamingDistributionInput) DeleteStreamingDistributionRequest {
	op := &aws.Operation{
		Name:       opDeleteStreamingDistribution,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2019-03-26/streaming-distribution/{Id}",
	}

	if input == nil {
		input = &DeleteStreamingDistributionInput{}
	}

	req := c.newRequest(op, input, &DeleteStreamingDistributionOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteStreamingDistributionRequest{Request: req, Input: input, Copy: c.DeleteStreamingDistributionRequest}
}

// DeleteStreamingDistributionRequest is the request type for the
// DeleteStreamingDistribution API operation.
type DeleteStreamingDistributionRequest struct {
	*aws.Request
	Input *DeleteStreamingDistributionInput
	Copy  func(*DeleteStreamingDistributionInput) DeleteStreamingDistributionRequest
}

// Send marshals and sends the DeleteStreamingDistribution API request.
func (r DeleteStreamingDistributionRequest) Send(ctx context.Context) (*DeleteStreamingDistributionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteStreamingDistributionResponse{
		DeleteStreamingDistributionOutput: r.Request.Data.(*DeleteStreamingDistributionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteStreamingDistributionResponse is the response type for the
// DeleteStreamingDistribution API operation.
type DeleteStreamingDistributionResponse struct {
	*DeleteStreamingDistributionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteStreamingDistribution request.
func (r *DeleteStreamingDistributionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
