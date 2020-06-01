// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to update an existing MethodResponse resource.
type UpdateMethodResponseInput struct {
	_ struct{} `type:"structure"`

	// [Required] The HTTP verb of the Method resource.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] The Resource identifier for the MethodResponse resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The status code for the MethodResponse resource.
	//
	// StatusCode is a required field
	StatusCode *string `location:"uri" locationName:"status_code" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateMethodResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMethodResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMethodResponseInput"}

	if s.HttpMethod == nil {
		invalidParams.Add(aws.NewErrParamRequired("HttpMethod"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if s.StatusCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatusCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMethodResponseInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PatchOperations != nil {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.HttpMethod != nil {
		v := *s.HttpMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "http_method", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resource_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StatusCode != nil {
		v := *s.StatusCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "status_code", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a method response of a given HTTP status code returned to the
// client. The method response is passed from the back end through the associated
// integration response that can be transformed using a mapping template.
//
// Example: A MethodResponse instance of an API
//
// Request
//
// The example request retrieves a MethodResponse of the 200 status code.
//  GET /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200
//  HTTP/1.1 Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
//  X-Amz-Date: 20160603T222952Z Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160603/us-east-1/apigateway/aws4_request,
//  SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
// Response
//
// The successful response returns 200 OK status and a payload as follows:
//  { "_links": { "curies": { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
//  "name": "methodresponse", "templated": true }, "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
//  "title": "200" }, "methodresponse:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
//  }, "methodresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200"
//  } }, "responseModels": { "application/json": "Empty" }, "responseParameters":
//  { "method.response.header.Content-Type": false }, "statusCode": "200" }
//
// Method, IntegrationResponse, Integration Creating an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type UpdateMethodResponseOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the Model resources used for the response's content-type. Response
	// models are represented as a key/value map, with a content-type as the key
	// and a Model name as the value.
	ResponseModels map[string]string `locationName:"responseModels" type:"map"`

	// A key-value map specifying required or optional response parameters that
	// API Gateway can send back to the caller. A key defines a method response
	// header and the value specifies whether the associated method response header
	// is required or not. The expression of the key must match the pattern method.response.header.{name},
	// where name is a valid and unique header name. API Gateway passes certain
	// integration response data to the method response headers specified here according
	// to the mapping you prescribe in the API's IntegrationResponse. The integration
	// response data that can be mapped include an integration response header expressed
	// in integration.response.header.{name}, a static value enclosed within a pair
	// of single quotes (e.g., 'application/json'), or a JSON expression from the
	// back-end response payload in the form of integration.response.body.{JSON-expression},
	// where JSON-expression is a valid JSON expression without the $ prefix.)
	ResponseParameters map[string]bool `locationName:"responseParameters" type:"map"`

	// The method response's status code.
	StatusCode *string `locationName:"statusCode" type:"string"`
}

// String returns the string representation
func (s UpdateMethodResponseOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMethodResponseOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ResponseModels != nil {
		v := s.ResponseModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResponseParameters != nil {
		v := s.ResponseParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.BoolValue(v1))
		}
		ms0.End()

	}
	if s.StatusCode != nil {
		v := *s.StatusCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateMethodResponse = "UpdateMethodResponse"

// UpdateMethodResponseRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Updates an existing MethodResponse resource.
//
//    // Example sending a request using UpdateMethodResponseRequest.
//    req := client.UpdateMethodResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateMethodResponseRequest(input *UpdateMethodResponseInput) UpdateMethodResponseRequest {
	op := &aws.Operation{
		Name:       opUpdateMethodResponse,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}/responses/{status_code}",
	}

	if input == nil {
		input = &UpdateMethodResponseInput{}
	}

	req := c.newRequest(op, input, &UpdateMethodResponseOutput{})

	return UpdateMethodResponseRequest{Request: req, Input: input, Copy: c.UpdateMethodResponseRequest}
}

// UpdateMethodResponseRequest is the request type for the
// UpdateMethodResponse API operation.
type UpdateMethodResponseRequest struct {
	*aws.Request
	Input *UpdateMethodResponseInput
	Copy  func(*UpdateMethodResponseInput) UpdateMethodResponseRequest
}

// Send marshals and sends the UpdateMethodResponse API request.
func (r UpdateMethodResponseRequest) Send(ctx context.Context) (*UpdateMethodResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMethodResponseResponse{
		UpdateMethodResponseOutput: r.Request.Data.(*UpdateMethodResponseOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMethodResponseResponse is the response type for the
// UpdateMethodResponse API operation.
type UpdateMethodResponseResponse struct {
	*UpdateMethodResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMethodResponse request.
func (r *UpdateMethodResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
