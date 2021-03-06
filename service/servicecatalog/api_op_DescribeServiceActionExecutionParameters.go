// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeServiceActionExecutionParametersInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The identifier of the provisioned product.
	//
	// ProvisionedProductId is a required field
	ProvisionedProductId *string `min:"1" type:"string" required:"true"`

	// The self-service action identifier.
	//
	// ServiceActionId is a required field
	ServiceActionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeServiceActionExecutionParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServiceActionExecutionParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeServiceActionExecutionParametersInput"}

	if s.ProvisionedProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisionedProductId"))
	}
	if s.ProvisionedProductId != nil && len(*s.ProvisionedProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisionedProductId", 1))
	}

	if s.ServiceActionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceActionId"))
	}
	if s.ServiceActionId != nil && len(*s.ServiceActionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceActionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeServiceActionExecutionParametersOutput struct {
	_ struct{} `type:"structure"`

	// The parameters of the self-service action.
	ServiceActionParameters []ExecutionParameter `type:"list"`
}

// String returns the string representation
func (s DescribeServiceActionExecutionParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeServiceActionExecutionParameters = "DescribeServiceActionExecutionParameters"

// DescribeServiceActionExecutionParametersRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Finds the default parameters for a specific self-service action on a specific
// provisioned product and returns a map of the results to the user.
//
//    // Example sending a request using DescribeServiceActionExecutionParametersRequest.
//    req := client.DescribeServiceActionExecutionParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeServiceActionExecutionParameters
func (c *Client) DescribeServiceActionExecutionParametersRequest(input *DescribeServiceActionExecutionParametersInput) DescribeServiceActionExecutionParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeServiceActionExecutionParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeServiceActionExecutionParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeServiceActionExecutionParametersOutput{})
	return DescribeServiceActionExecutionParametersRequest{Request: req, Input: input, Copy: c.DescribeServiceActionExecutionParametersRequest}
}

// DescribeServiceActionExecutionParametersRequest is the request type for the
// DescribeServiceActionExecutionParameters API operation.
type DescribeServiceActionExecutionParametersRequest struct {
	*aws.Request
	Input *DescribeServiceActionExecutionParametersInput
	Copy  func(*DescribeServiceActionExecutionParametersInput) DescribeServiceActionExecutionParametersRequest
}

// Send marshals and sends the DescribeServiceActionExecutionParameters API request.
func (r DescribeServiceActionExecutionParametersRequest) Send(ctx context.Context) (*DescribeServiceActionExecutionParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeServiceActionExecutionParametersResponse{
		DescribeServiceActionExecutionParametersOutput: r.Request.Data.(*DescribeServiceActionExecutionParametersOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeServiceActionExecutionParametersResponse is the response type for the
// DescribeServiceActionExecutionParameters API operation.
type DescribeServiceActionExecutionParametersResponse struct {
	*DescribeServiceActionExecutionParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeServiceActionExecutionParameters request.
func (r *DescribeServiceActionExecutionParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
