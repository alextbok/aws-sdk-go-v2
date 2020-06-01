// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type AssociateAssetsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the parent asset.
	//
	// AssetId is a required field
	AssetId *string `location:"uri" locationName:"assetId" min:"36" type:"string" required:"true"`

	// The ID of the child asset to be associated.
	//
	// ChildAssetId is a required field
	ChildAssetId *string `locationName:"childAssetId" min:"36" type:"string" required:"true"`

	// A unique case-sensitive identifier that you can provide to ensure the idempotency
	// of the request. Don't reuse this client token if a new idempotent request
	// is required.
	ClientToken *string `locationName:"clientToken" min:"36" type:"string" idempotencyToken:"true"`

	// The ID of a hierarchy in the parent asset's model. Hierarchies allow different
	// groupings of assets to be formed that all come from the same asset model.
	// For more information, see Asset Hierarchies (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// HierarchyId is a required field
	HierarchyId *string `locationName:"hierarchyId" min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateAssetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateAssetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateAssetsInput"}

	if s.AssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetId"))
	}
	if s.AssetId != nil && len(*s.AssetId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AssetId", 36))
	}

	if s.ChildAssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChildAssetId"))
	}
	if s.ChildAssetId != nil && len(*s.ChildAssetId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ChildAssetId", 36))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}

	if s.HierarchyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HierarchyId"))
	}
	if s.HierarchyId != nil && len(*s.HierarchyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("HierarchyId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateAssetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChildAssetId != nil {
		v := *s.ChildAssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "childAssetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HierarchyId != nil {
		v := *s.HierarchyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "hierarchyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AssetId != nil {
		v := *s.AssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "assetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AssociateAssetsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateAssetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateAssetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAssociateAssets = "AssociateAssets"

// AssociateAssetsRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Associates a child asset with the given parent asset through a hierarchy
// defined in the parent asset's model. For more information, see Associating
// Assets (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/add-associated-assets.html)
// in the AWS IoT SiteWise User Guide.
//
//    // Example sending a request using AssociateAssetsRequest.
//    req := client.AssociateAssetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/AssociateAssets
func (c *Client) AssociateAssetsRequest(input *AssociateAssetsInput) AssociateAssetsRequest {
	op := &aws.Operation{
		Name:       opAssociateAssets,
		HTTPMethod: "POST",
		HTTPPath:   "/assets/{assetId}/associate",
	}

	if input == nil {
		input = &AssociateAssetsInput{}
	}

	req := c.newRequest(op, input, &AssociateAssetsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("model.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return AssociateAssetsRequest{Request: req, Input: input, Copy: c.AssociateAssetsRequest}
}

// AssociateAssetsRequest is the request type for the
// AssociateAssets API operation.
type AssociateAssetsRequest struct {
	*aws.Request
	Input *AssociateAssetsInput
	Copy  func(*AssociateAssetsInput) AssociateAssetsRequest
}

// Send marshals and sends the AssociateAssets API request.
func (r AssociateAssetsRequest) Send(ctx context.Context) (*AssociateAssetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateAssetsResponse{
		AssociateAssetsOutput: r.Request.Data.(*AssociateAssetsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateAssetsResponse is the response type for the
// AssociateAssets API operation.
type AssociateAssetsResponse struct {
	*AssociateAssetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateAssets request.
func (r *AssociateAssetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
