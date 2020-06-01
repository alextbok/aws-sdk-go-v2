// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to get information about a specified health check.
type GetHealthCheckInput struct {
	_ struct{} `type:"structure"`

	// The identifier that Amazon Route 53 assigned to the health check when you
	// created it. When you add or update a resource record set, you use this value
	// to specify which health check to use. The value can be up to 64 characters
	// long.
	//
	// HealthCheckId is a required field
	HealthCheckId *string `location:"uri" locationName:"HealthCheckId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHealthCheckInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHealthCheckInput"}

	if s.HealthCheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthCheckId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHealthCheckInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.HealthCheckId != nil {
		v := *s.HealthCheckId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "HealthCheckId", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response to a GetHealthCheck request.
type GetHealthCheckOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about one health check that is associated
	// with the current AWS account.
	//
	// HealthCheck is a required field
	HealthCheck *HealthCheck `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetHealthCheckOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HealthCheck != nil {
		v := s.HealthCheck

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HealthCheck", v, metadata)
	}
	return nil
}

const opGetHealthCheck = "GetHealthCheck"

// GetHealthCheckRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets information about a specified health check.
//
//    // Example sending a request using GetHealthCheckRequest.
//    req := client.GetHealthCheckRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetHealthCheck
func (c *Client) GetHealthCheckRequest(input *GetHealthCheckInput) GetHealthCheckRequest {
	op := &aws.Operation{
		Name:       opGetHealthCheck,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/healthcheck/{HealthCheckId}",
	}

	if input == nil {
		input = &GetHealthCheckInput{}
	}

	req := c.newRequest(op, input, &GetHealthCheckOutput{})

	return GetHealthCheckRequest{Request: req, Input: input, Copy: c.GetHealthCheckRequest}
}

// GetHealthCheckRequest is the request type for the
// GetHealthCheck API operation.
type GetHealthCheckRequest struct {
	*aws.Request
	Input *GetHealthCheckInput
	Copy  func(*GetHealthCheckInput) GetHealthCheckRequest
}

// Send marshals and sends the GetHealthCheck API request.
func (r GetHealthCheckRequest) Send(ctx context.Context) (*GetHealthCheckResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHealthCheckResponse{
		GetHealthCheckOutput: r.Request.Data.(*GetHealthCheckOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHealthCheckResponse is the response type for the
// GetHealthCheck API operation.
type GetHealthCheckResponse struct {
	*GetHealthCheckOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHealthCheck request.
func (r *GetHealthCheckResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
