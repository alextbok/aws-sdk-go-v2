// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request a new export of a RestApi for a particular Stage.
type GetExportInput struct {
	_ struct{} `type:"structure"`

	// The content-type of the export, for example application/json. Currently application/json
	// and application/yaml are supported for exportType ofoas30 and swagger. This
	// should be specified in the Accept header for direct API requests.
	Accepts *string `location:"header" locationName:"Accept" type:"string"`

	// [Required] The type of export. Acceptable values are 'oas30' for OpenAPI
	// 3.0.x and 'swagger' for Swagger/OpenAPI 2.0.
	//
	// ExportType is a required field
	ExportType *string `location:"uri" locationName:"export_type" type:"string" required:"true"`

	// A key-value map of query string parameters that specify properties of the
	// export, depending on the requested exportType. For exportType oas30 and swagger,
	// any combination of the following parameters are supported: extensions='integrations'
	// or extensions='apigateway' will export the API with x-amazon-apigateway-integration
	// extensions. extensions='authorizers' will export the API with x-amazon-apigateway-authorizer
	// extensions. postman will export the API with Postman extensions, allowing
	// for import to the Postman tool
	Parameters map[string]string `location:"querystring" locationName:"parameters" type:"map"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The name of the Stage that will be exported.
	//
	// StageName is a required field
	StageName *string `location:"uri" locationName:"stage_name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetExportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetExportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetExportInput"}

	if s.ExportType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExportType"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetExportInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Accepts != nil {
		v := *s.Accepts

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Accept", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExportType != nil {
		v := *s.ExportType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "export_type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "stage_name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Parameters != nil {
		v := s.Parameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.QueryTarget, "parameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// The binary blob response to GetExport, which contains the generated SDK.
type GetExportOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The binary blob response to GetExport, which contains the export.
	Body []byte `locationName:"body" type:"blob"`

	// The content-disposition header value in the HTTP response.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// The content-type header value in the HTTP response. This will correspond
	// to a valid 'accept' type in the request.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`
}

// String returns the string representation
func (s GetExportOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetExportOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentDisposition != nil {
		v := *s.ContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Disposition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "body", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opGetExport = "GetExport"

// GetExportRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Exports a deployed version of a RestApi in a specified format.
//
//    // Example sending a request using GetExportRequest.
//    req := client.GetExportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetExportRequest(input *GetExportInput) GetExportRequest {
	op := &aws.Operation{
		Name:       opGetExport,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/stages/{stage_name}/exports/{export_type}",
	}

	if input == nil {
		input = &GetExportInput{}
	}

	req := c.newRequest(op, input, &GetExportOutput{})

	return GetExportRequest{Request: req, Input: input, Copy: c.GetExportRequest}
}

// GetExportRequest is the request type for the
// GetExport API operation.
type GetExportRequest struct {
	*aws.Request
	Input *GetExportInput
	Copy  func(*GetExportInput) GetExportRequest
}

// Send marshals and sends the GetExport API request.
func (r GetExportRequest) Send(ctx context.Context) (*GetExportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetExportResponse{
		GetExportOutput: r.Request.Data.(*GetExportOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetExportResponse is the response type for the
// GetExport API operation.
type GetExportResponse struct {
	*GetExportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetExport request.
func (r *GetExportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
