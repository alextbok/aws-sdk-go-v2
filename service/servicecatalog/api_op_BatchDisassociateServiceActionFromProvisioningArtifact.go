// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDisassociateServiceActionFromProvisioningArtifactInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// One or more associations, each consisting of the Action ID, the Product ID,
	// and the Provisioning Artifact ID.
	//
	// ServiceActionAssociations is a required field
	ServiceActionAssociations []ServiceActionAssociation `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisassociateServiceActionFromProvisioningArtifactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDisassociateServiceActionFromProvisioningArtifactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDisassociateServiceActionFromProvisioningArtifactInput"}

	if s.ServiceActionAssociations == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceActionAssociations"))
	}
	if s.ServiceActionAssociations != nil && len(s.ServiceActionAssociations) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceActionAssociations", 1))
	}
	if s.ServiceActionAssociations != nil {
		for i, v := range s.ServiceActionAssociations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ServiceActionAssociations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDisassociateServiceActionFromProvisioningArtifactOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains a list of errors, along with information to help
	// you identify the self-service action.
	FailedServiceActionAssociations []FailedServiceActionAssociation `type:"list"`
}

// String returns the string representation
func (s BatchDisassociateServiceActionFromProvisioningArtifactOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDisassociateServiceActionFromProvisioningArtifact = "BatchDisassociateServiceActionFromProvisioningArtifact"

// BatchDisassociateServiceActionFromProvisioningArtifactRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Disassociates a batch of self-service actions from the specified provisioning
// artifact.
//
//    // Example sending a request using BatchDisassociateServiceActionFromProvisioningArtifactRequest.
//    req := client.BatchDisassociateServiceActionFromProvisioningArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/BatchDisassociateServiceActionFromProvisioningArtifact
func (c *Client) BatchDisassociateServiceActionFromProvisioningArtifactRequest(input *BatchDisassociateServiceActionFromProvisioningArtifactInput) BatchDisassociateServiceActionFromProvisioningArtifactRequest {
	op := &aws.Operation{
		Name:       opBatchDisassociateServiceActionFromProvisioningArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDisassociateServiceActionFromProvisioningArtifactInput{}
	}

	req := c.newRequest(op, input, &BatchDisassociateServiceActionFromProvisioningArtifactOutput{})

	return BatchDisassociateServiceActionFromProvisioningArtifactRequest{Request: req, Input: input, Copy: c.BatchDisassociateServiceActionFromProvisioningArtifactRequest}
}

// BatchDisassociateServiceActionFromProvisioningArtifactRequest is the request type for the
// BatchDisassociateServiceActionFromProvisioningArtifact API operation.
type BatchDisassociateServiceActionFromProvisioningArtifactRequest struct {
	*aws.Request
	Input *BatchDisassociateServiceActionFromProvisioningArtifactInput
	Copy  func(*BatchDisassociateServiceActionFromProvisioningArtifactInput) BatchDisassociateServiceActionFromProvisioningArtifactRequest
}

// Send marshals and sends the BatchDisassociateServiceActionFromProvisioningArtifact API request.
func (r BatchDisassociateServiceActionFromProvisioningArtifactRequest) Send(ctx context.Context) (*BatchDisassociateServiceActionFromProvisioningArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDisassociateServiceActionFromProvisioningArtifactResponse{
		BatchDisassociateServiceActionFromProvisioningArtifactOutput: r.Request.Data.(*BatchDisassociateServiceActionFromProvisioningArtifactOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDisassociateServiceActionFromProvisioningArtifactResponse is the response type for the
// BatchDisassociateServiceActionFromProvisioningArtifact API operation.
type BatchDisassociateServiceActionFromProvisioningArtifactResponse struct {
	*BatchDisassociateServiceActionFromProvisioningArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDisassociateServiceActionFromProvisioningArtifact request.
func (r *BatchDisassociateServiceActionFromProvisioningArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
