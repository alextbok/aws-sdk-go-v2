// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for update Domain Association request.
type UpdateDomainAssociationInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name of the domain.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`

	// Enables automated creation of Subdomains for branches. (Currently not supported)
	EnableAutoSubDomain *bool `locationName:"enableAutoSubDomain" type:"boolean"`

	// Setting structure for the Subdomain.
	//
	// SubDomainSettings is a required field
	SubDomainSettings []SubDomainSetting `locationName:"subDomainSettings" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateDomainAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDomainAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDomainAssociationInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.SubDomainSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubDomainSettings"))
	}
	if s.SubDomainSettings != nil {
		for i, v := range s.SubDomainSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SubDomainSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDomainAssociationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EnableAutoSubDomain != nil {
		v := *s.EnableAutoSubDomain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableAutoSubDomain", protocol.BoolValue(v), metadata)
	}
	if s.SubDomainSettings != nil {
		v := s.SubDomainSettings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subDomainSettings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the update Domain Association request.
type UpdateDomainAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Domain Association structure.
	//
	// DomainAssociation is a required field
	DomainAssociation *DomainAssociation `locationName:"domainAssociation" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateDomainAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDomainAssociationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainAssociation != nil {
		v := s.DomainAssociation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "domainAssociation", v, metadata)
	}
	return nil
}

const opUpdateDomainAssociation = "UpdateDomainAssociation"

// UpdateDomainAssociationRequest returns a request value for making API operation for
// AWS Amplify.
//
// Create a new DomainAssociation on an App
//
//    // Example sending a request using UpdateDomainAssociationRequest.
//    req := client.UpdateDomainAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateDomainAssociation
func (c *Client) UpdateDomainAssociationRequest(input *UpdateDomainAssociationInput) UpdateDomainAssociationRequest {
	op := &aws.Operation{
		Name:       opUpdateDomainAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}/domains/{domainName}",
	}

	if input == nil {
		input = &UpdateDomainAssociationInput{}
	}

	req := c.newRequest(op, input, &UpdateDomainAssociationOutput{})

	return UpdateDomainAssociationRequest{Request: req, Input: input, Copy: c.UpdateDomainAssociationRequest}
}

// UpdateDomainAssociationRequest is the request type for the
// UpdateDomainAssociation API operation.
type UpdateDomainAssociationRequest struct {
	*aws.Request
	Input *UpdateDomainAssociationInput
	Copy  func(*UpdateDomainAssociationInput) UpdateDomainAssociationRequest
}

// Send marshals and sends the UpdateDomainAssociation API request.
func (r UpdateDomainAssociationRequest) Send(ctx context.Context) (*UpdateDomainAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDomainAssociationResponse{
		UpdateDomainAssociationOutput: r.Request.Data.(*UpdateDomainAssociationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDomainAssociationResponse is the response type for the
// UpdateDomainAssociation API operation.
type UpdateDomainAssociationResponse struct {
	*UpdateDomainAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDomainAssociation request.
func (r *UpdateDomainAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
