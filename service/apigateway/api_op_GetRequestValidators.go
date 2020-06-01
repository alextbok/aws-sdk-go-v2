// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Gets the RequestValidators collection of a given RestApi.
type GetRequestValidatorsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRequestValidatorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRequestValidatorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRequestValidatorsInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRequestValidatorsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A collection of RequestValidator resources of a given RestApi.
//
// In OpenAPI, the RequestValidators of an API is defined by the x-amazon-apigateway-request-validators
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validators.html)
// extension.
//
// Enable Basic Request Validation in API Gateway (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html)
type GetRequestValidatorsOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []RequestValidator `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetRequestValidatorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRequestValidatorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetRequestValidators = "GetRequestValidators"

// GetRequestValidatorsRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets the RequestValidators collection of a given RestApi.
//
//    // Example sending a request using GetRequestValidatorsRequest.
//    req := client.GetRequestValidatorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetRequestValidatorsRequest(input *GetRequestValidatorsInput) GetRequestValidatorsRequest {
	op := &aws.Operation{
		Name:       opGetRequestValidators,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/requestvalidators",
	}

	if input == nil {
		input = &GetRequestValidatorsInput{}
	}

	req := c.newRequest(op, input, &GetRequestValidatorsOutput{})

	return GetRequestValidatorsRequest{Request: req, Input: input, Copy: c.GetRequestValidatorsRequest}
}

// GetRequestValidatorsRequest is the request type for the
// GetRequestValidators API operation.
type GetRequestValidatorsRequest struct {
	*aws.Request
	Input *GetRequestValidatorsInput
	Copy  func(*GetRequestValidatorsInput) GetRequestValidatorsRequest
}

// Send marshals and sends the GetRequestValidators API request.
func (r GetRequestValidatorsRequest) Send(ctx context.Context) (*GetRequestValidatorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRequestValidatorsResponse{
		GetRequestValidatorsOutput: r.Request.Data.(*GetRequestValidatorsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRequestValidatorsResponse is the response type for the
// GetRequestValidators API operation.
type GetRequestValidatorsResponse struct {
	*GetRequestValidatorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRequestValidators request.
func (r *GetRequestValidatorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
