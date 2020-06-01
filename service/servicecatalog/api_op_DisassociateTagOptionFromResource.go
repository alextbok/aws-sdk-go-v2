// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateTagOptionFromResourceInput struct {
	_ struct{} `type:"structure"`

	// The resource identifier.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// The TagOption identifier.
	//
	// TagOptionId is a required field
	TagOptionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateTagOptionFromResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateTagOptionFromResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateTagOptionFromResourceInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.TagOptionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagOptionId"))
	}
	if s.TagOptionId != nil && len(*s.TagOptionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagOptionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateTagOptionFromResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateTagOptionFromResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateTagOptionFromResource = "DisassociateTagOptionFromResource"

// DisassociateTagOptionFromResourceRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Disassociates the specified TagOption from the specified resource.
//
//    // Example sending a request using DisassociateTagOptionFromResourceRequest.
//    req := client.DisassociateTagOptionFromResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DisassociateTagOptionFromResource
func (c *Client) DisassociateTagOptionFromResourceRequest(input *DisassociateTagOptionFromResourceInput) DisassociateTagOptionFromResourceRequest {
	op := &aws.Operation{
		Name:       opDisassociateTagOptionFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateTagOptionFromResourceInput{}
	}

	req := c.newRequest(op, input, &DisassociateTagOptionFromResourceOutput{})

	return DisassociateTagOptionFromResourceRequest{Request: req, Input: input, Copy: c.DisassociateTagOptionFromResourceRequest}
}

// DisassociateTagOptionFromResourceRequest is the request type for the
// DisassociateTagOptionFromResource API operation.
type DisassociateTagOptionFromResourceRequest struct {
	*aws.Request
	Input *DisassociateTagOptionFromResourceInput
	Copy  func(*DisassociateTagOptionFromResourceInput) DisassociateTagOptionFromResourceRequest
}

// Send marshals and sends the DisassociateTagOptionFromResource API request.
func (r DisassociateTagOptionFromResourceRequest) Send(ctx context.Context) (*DisassociateTagOptionFromResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateTagOptionFromResourceResponse{
		DisassociateTagOptionFromResourceOutput: r.Request.Data.(*DisassociateTagOptionFromResourceOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateTagOptionFromResourceResponse is the response type for the
// DisassociateTagOptionFromResource API operation.
type DisassociateTagOptionFromResourceResponse struct {
	*DisassociateTagOptionFromResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateTagOptionFromResource request.
func (r *DisassociateTagOptionFromResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
