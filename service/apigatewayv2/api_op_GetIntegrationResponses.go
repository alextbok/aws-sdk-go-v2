// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetIntegrationResponsesInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// IntegrationId is a required field
	IntegrationId *string `location:"uri" locationName:"integrationId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetIntegrationResponsesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIntegrationResponsesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIntegrationResponsesInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.IntegrationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IntegrationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntegrationResponsesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntegrationId != nil {
		v := *s.IntegrationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "integrationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetIntegrationResponsesOutput struct {
	_ struct{} `type:"structure"`

	Items []IntegrationResponse `locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetIntegrationResponsesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntegrationResponsesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetIntegrationResponses = "GetIntegrationResponses"

// GetIntegrationResponsesRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the IntegrationResponses for an Integration.
//
//    // Example sending a request using GetIntegrationResponsesRequest.
//    req := client.GetIntegrationResponsesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetIntegrationResponses
func (c *Client) GetIntegrationResponsesRequest(input *GetIntegrationResponsesInput) GetIntegrationResponsesRequest {
	op := &aws.Operation{
		Name:       opGetIntegrationResponses,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/integrations/{integrationId}/integrationresponses",
	}

	if input == nil {
		input = &GetIntegrationResponsesInput{}
	}

	req := c.newRequest(op, input, &GetIntegrationResponsesOutput{})

	return GetIntegrationResponsesRequest{Request: req, Input: input, Copy: c.GetIntegrationResponsesRequest}
}

// GetIntegrationResponsesRequest is the request type for the
// GetIntegrationResponses API operation.
type GetIntegrationResponsesRequest struct {
	*aws.Request
	Input *GetIntegrationResponsesInput
	Copy  func(*GetIntegrationResponsesInput) GetIntegrationResponsesRequest
}

// Send marshals and sends the GetIntegrationResponses API request.
func (r GetIntegrationResponsesRequest) Send(ctx context.Context) (*GetIntegrationResponsesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIntegrationResponsesResponse{
		GetIntegrationResponsesOutput: r.Request.Data.(*GetIntegrationResponsesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIntegrationResponsesResponse is the response type for the
// GetIntegrationResponses API operation.
type GetIntegrationResponsesResponse struct {
	*GetIntegrationResponsesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIntegrationResponses request.
func (r *GetIntegrationResponsesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
