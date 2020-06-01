// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Import documentation parts from an external (e.g., OpenAPI) definition file.
type ImportDocumentationPartsInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// [Required] Raw byte array representing the to-be-imported documentation parts.
	// To import from an OpenAPI file, this is a JSON object.
	//
	// Body is a required field
	Body []byte `locationName:"body" type:"blob" required:"true"`

	// A query parameter to specify whether to rollback the documentation importation
	// (true) or not (false) when a warning is encountered. The default value is
	// false.
	FailOnWarnings *bool `location:"querystring" locationName:"failonwarnings" type:"boolean"`

	// A query parameter to indicate whether to overwrite (OVERWRITE) any existing
	// DocumentationParts definition or to merge (MERGE) the new definition into
	// the existing one. The default value is MERGE.
	Mode PutMode `location:"querystring" locationName:"mode" type:"string" enum:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s ImportDocumentationPartsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportDocumentationPartsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportDocumentationPartsInput"}

	if s.Body == nil {
		invalidParams.Add(aws.NewErrParamRequired("Body"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportDocumentationPartsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "body", protocol.BytesStream(v), metadata)
	}
	if s.FailOnWarnings != nil {
		v := *s.FailOnWarnings

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "failonwarnings", protocol.BoolValue(v), metadata)
	}
	if len(s.Mode) > 0 {
		v := s.Mode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "mode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// A collection of the imported DocumentationPart identifiers.
//
// This is used to return the result when documentation parts in an external
// (e.g., OpenAPI) file are imported into API Gateway
//
// Documenting an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// documentationpart:import (https://docs.aws.amazon.com/apigateway/api-reference/link-relation/documentationpart-import/),
// DocumentationPart
type ImportDocumentationPartsOutput struct {
	_ struct{} `type:"structure"`

	// A list of the returned documentation part identifiers.
	Ids []string `locationName:"ids" type:"list"`

	// A list of warning messages reported during import of documentation parts.
	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s ImportDocumentationPartsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportDocumentationPartsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Ids != nil {
		v := s.Ids

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ids", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opImportDocumentationParts = "ImportDocumentationParts"

// ImportDocumentationPartsRequest returns a request value for making API operation for
// Amazon API Gateway.
//
//    // Example sending a request using ImportDocumentationPartsRequest.
//    req := client.ImportDocumentationPartsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ImportDocumentationPartsRequest(input *ImportDocumentationPartsInput) ImportDocumentationPartsRequest {
	op := &aws.Operation{
		Name:       opImportDocumentationParts,
		HTTPMethod: "PUT",
		HTTPPath:   "/restapis/{restapi_id}/documentation/parts",
	}

	if input == nil {
		input = &ImportDocumentationPartsInput{}
	}

	req := c.newRequest(op, input, &ImportDocumentationPartsOutput{})

	return ImportDocumentationPartsRequest{Request: req, Input: input, Copy: c.ImportDocumentationPartsRequest}
}

// ImportDocumentationPartsRequest is the request type for the
// ImportDocumentationParts API operation.
type ImportDocumentationPartsRequest struct {
	*aws.Request
	Input *ImportDocumentationPartsInput
	Copy  func(*ImportDocumentationPartsInput) ImportDocumentationPartsRequest
}

// Send marshals and sends the ImportDocumentationParts API request.
func (r ImportDocumentationPartsRequest) Send(ctx context.Context) (*ImportDocumentationPartsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportDocumentationPartsResponse{
		ImportDocumentationPartsOutput: r.Request.Data.(*ImportDocumentationPartsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportDocumentationPartsResponse is the response type for the
// ImportDocumentationParts API operation.
type ImportDocumentationPartsResponse struct {
	*ImportDocumentationPartsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportDocumentationParts request.
func (r *ImportDocumentationPartsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
