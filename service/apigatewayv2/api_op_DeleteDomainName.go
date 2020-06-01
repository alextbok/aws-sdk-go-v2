// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteDomainNameInput struct {
	_ struct{} `type:"structure"`

	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDomainNameInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDomainNameInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDomainNameInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDomainNameInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDomainNameOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDomainNameOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDomainNameOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDomainName = "DeleteDomainName"

// DeleteDomainNameRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes a domain name.
//
//    // Example sending a request using DeleteDomainNameRequest.
//    req := client.DeleteDomainNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteDomainName
func (c *Client) DeleteDomainNameRequest(input *DeleteDomainNameInput) DeleteDomainNameRequest {
	op := &aws.Operation{
		Name:       opDeleteDomainName,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/domainnames/{domainName}",
	}

	if input == nil {
		input = &DeleteDomainNameInput{}
	}

	req := c.newRequest(op, input, &DeleteDomainNameOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteDomainNameRequest{Request: req, Input: input, Copy: c.DeleteDomainNameRequest}
}

// DeleteDomainNameRequest is the request type for the
// DeleteDomainName API operation.
type DeleteDomainNameRequest struct {
	*aws.Request
	Input *DeleteDomainNameInput
	Copy  func(*DeleteDomainNameInput) DeleteDomainNameRequest
}

// Send marshals and sends the DeleteDomainName API request.
func (r DeleteDomainNameRequest) Send(ctx context.Context) (*DeleteDomainNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDomainNameResponse{
		DeleteDomainNameOutput: r.Request.Data.(*DeleteDomainNameOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDomainNameResponse is the response type for the
// DeleteDomainName API operation.
type DeleteDomainNameResponse struct {
	*DeleteDomainNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDomainName request.
func (r *DeleteDomainNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
