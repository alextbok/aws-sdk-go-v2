// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteProvisioningArtifactInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The product identifier.
	//
	// ProductId is a required field
	ProductId *string `min:"1" type:"string" required:"true"`

	// The identifier of the provisioning artifact.
	//
	// ProvisioningArtifactId is a required field
	ProvisioningArtifactId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProvisioningArtifactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProvisioningArtifactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProvisioningArtifactInput"}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if s.ProvisioningArtifactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisioningArtifactId"))
	}
	if s.ProvisioningArtifactId != nil && len(*s.ProvisioningArtifactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisioningArtifactId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteProvisioningArtifactOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProvisioningArtifactOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteProvisioningArtifact = "DeleteProvisioningArtifact"

// DeleteProvisioningArtifactRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Deletes the specified provisioning artifact (also known as a version) for
// the specified product.
//
// You cannot delete a provisioning artifact associated with a product that
// was shared with you. You cannot delete the last provisioning artifact for
// a product, because a product must have at least one provisioning artifact.
//
//    // Example sending a request using DeleteProvisioningArtifactRequest.
//    req := client.DeleteProvisioningArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteProvisioningArtifact
func (c *Client) DeleteProvisioningArtifactRequest(input *DeleteProvisioningArtifactInput) DeleteProvisioningArtifactRequest {
	op := &aws.Operation{
		Name:       opDeleteProvisioningArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProvisioningArtifactInput{}
	}

	req := c.newRequest(op, input, &DeleteProvisioningArtifactOutput{})

	return DeleteProvisioningArtifactRequest{Request: req, Input: input, Copy: c.DeleteProvisioningArtifactRequest}
}

// DeleteProvisioningArtifactRequest is the request type for the
// DeleteProvisioningArtifact API operation.
type DeleteProvisioningArtifactRequest struct {
	*aws.Request
	Input *DeleteProvisioningArtifactInput
	Copy  func(*DeleteProvisioningArtifactInput) DeleteProvisioningArtifactRequest
}

// Send marshals and sends the DeleteProvisioningArtifact API request.
func (r DeleteProvisioningArtifactRequest) Send(ctx context.Context) (*DeleteProvisioningArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProvisioningArtifactResponse{
		DeleteProvisioningArtifactOutput: r.Request.Data.(*DeleteProvisioningArtifactOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProvisioningArtifactResponse is the response type for the
// DeleteProvisioningArtifact API operation.
type DeleteProvisioningArtifactResponse struct {
	*DeleteProvisioningArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProvisioningArtifact request.
func (r *DeleteProvisioningArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
