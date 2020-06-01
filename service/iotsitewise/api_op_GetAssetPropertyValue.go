// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetAssetPropertyValueInput struct {
	_ struct{} `type:"structure"`

	// The ID of the asset.
	AssetId *string `location:"querystring" locationName:"assetId" min:"36" type:"string"`

	// The property alias that identifies the property, such as an OPC-UA server
	// data stream path (for example, /company/windfarm/3/turbine/7/temperature).
	// For more information, see Mapping Industrial Data Streams to Asset Properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide.
	PropertyAlias *string `location:"querystring" locationName:"propertyAlias" min:"1" type:"string"`

	// The ID of the asset property.
	PropertyId *string `location:"querystring" locationName:"propertyId" min:"36" type:"string"`
}

// String returns the string representation
func (s GetAssetPropertyValueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAssetPropertyValueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAssetPropertyValueInput"}
	if s.AssetId != nil && len(*s.AssetId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AssetId", 36))
	}
	if s.PropertyAlias != nil && len(*s.PropertyAlias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PropertyAlias", 1))
	}
	if s.PropertyId != nil && len(*s.PropertyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PropertyId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAssetPropertyValueInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssetId != nil {
		v := *s.AssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "assetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PropertyAlias != nil {
		v := *s.PropertyAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "propertyAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PropertyId != nil {
		v := *s.PropertyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "propertyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetAssetPropertyValueOutput struct {
	_ struct{} `type:"structure"`

	// The current asset property value.
	PropertyValue *AssetPropertyValue `locationName:"propertyValue" type:"structure"`
}

// String returns the string representation
func (s GetAssetPropertyValueOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAssetPropertyValueOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PropertyValue != nil {
		v := s.PropertyValue

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "propertyValue", v, metadata)
	}
	return nil
}

const opGetAssetPropertyValue = "GetAssetPropertyValue"

// GetAssetPropertyValueRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Gets an asset property's current value. For more information, see Querying
// Current Property Values (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#current-values)
// in the AWS IoT SiteWise User Guide.
//
// To identify an asset property, you must specify one of the following:
//
//    * The assetId and propertyId of an asset property.
//
//    * A propertyAlias, which is a data stream alias (for example, /company/windfarm/3/turbine/7/temperature).
//    To define an asset property's alias, see UpdateAssetProperty (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
//
//    // Example sending a request using GetAssetPropertyValueRequest.
//    req := client.GetAssetPropertyValueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/GetAssetPropertyValue
func (c *Client) GetAssetPropertyValueRequest(input *GetAssetPropertyValueInput) GetAssetPropertyValueRequest {
	op := &aws.Operation{
		Name:       opGetAssetPropertyValue,
		HTTPMethod: "GET",
		HTTPPath:   "/properties/latest",
	}

	if input == nil {
		input = &GetAssetPropertyValueInput{}
	}

	req := c.newRequest(op, input, &GetAssetPropertyValueOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("data.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return GetAssetPropertyValueRequest{Request: req, Input: input, Copy: c.GetAssetPropertyValueRequest}
}

// GetAssetPropertyValueRequest is the request type for the
// GetAssetPropertyValue API operation.
type GetAssetPropertyValueRequest struct {
	*aws.Request
	Input *GetAssetPropertyValueInput
	Copy  func(*GetAssetPropertyValueInput) GetAssetPropertyValueRequest
}

// Send marshals and sends the GetAssetPropertyValue API request.
func (r GetAssetPropertyValueRequest) Send(ctx context.Context) (*GetAssetPropertyValueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAssetPropertyValueResponse{
		GetAssetPropertyValueOutput: r.Request.Data.(*GetAssetPropertyValueOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAssetPropertyValueResponse is the response type for the
// GetAssetPropertyValue API operation.
type GetAssetPropertyValueResponse struct {
	*GetAssetPropertyValueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAssetPropertyValue request.
func (r *GetAssetPropertyValueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
