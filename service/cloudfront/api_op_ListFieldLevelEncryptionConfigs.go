// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListFieldLevelEncryptionConfigsInput struct {
	_ struct{} `type:"structure"`

	// Use this when paginating results to indicate where to begin in your list
	// of configurations. The results include configurations in the list that occur
	// after the marker. To get the next page of results, set the Marker to the
	// value of the NextMarker from the current page's response (which is also the
	// ID of the last configuration on that page).
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of field-level encryption configurations you want in the
	// response body.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListFieldLevelEncryptionConfigsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFieldLevelEncryptionConfigsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

type ListFieldLevelEncryptionConfigsOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionList"`

	// Returns a list of all field-level encryption configurations that have been
	// created in CloudFront for this account.
	FieldLevelEncryptionList *FieldLevelEncryptionList `type:"structure"`
}

// String returns the string representation
func (s ListFieldLevelEncryptionConfigsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFieldLevelEncryptionConfigsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FieldLevelEncryptionList != nil {
		v := s.FieldLevelEncryptionList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionList", v, metadata)
	}
	return nil
}

const opListFieldLevelEncryptionConfigs = "ListFieldLevelEncryptionConfigs2019_03_26"

// ListFieldLevelEncryptionConfigsRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// List all field-level encryption configurations that have been created in
// CloudFront for this account.
//
//    // Example sending a request using ListFieldLevelEncryptionConfigsRequest.
//    req := client.ListFieldLevelEncryptionConfigsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListFieldLevelEncryptionConfigs
func (c *Client) ListFieldLevelEncryptionConfigsRequest(input *ListFieldLevelEncryptionConfigsInput) ListFieldLevelEncryptionConfigsRequest {
	op := &aws.Operation{
		Name:       opListFieldLevelEncryptionConfigs,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/field-level-encryption",
	}

	if input == nil {
		input = &ListFieldLevelEncryptionConfigsInput{}
	}

	req := c.newRequest(op, input, &ListFieldLevelEncryptionConfigsOutput{})

	return ListFieldLevelEncryptionConfigsRequest{Request: req, Input: input, Copy: c.ListFieldLevelEncryptionConfigsRequest}
}

// ListFieldLevelEncryptionConfigsRequest is the request type for the
// ListFieldLevelEncryptionConfigs API operation.
type ListFieldLevelEncryptionConfigsRequest struct {
	*aws.Request
	Input *ListFieldLevelEncryptionConfigsInput
	Copy  func(*ListFieldLevelEncryptionConfigsInput) ListFieldLevelEncryptionConfigsRequest
}

// Send marshals and sends the ListFieldLevelEncryptionConfigs API request.
func (r ListFieldLevelEncryptionConfigsRequest) Send(ctx context.Context) (*ListFieldLevelEncryptionConfigsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFieldLevelEncryptionConfigsResponse{
		ListFieldLevelEncryptionConfigsOutput: r.Request.Data.(*ListFieldLevelEncryptionConfigsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListFieldLevelEncryptionConfigsResponse is the response type for the
// ListFieldLevelEncryptionConfigs API operation.
type ListFieldLevelEncryptionConfigsResponse struct {
	*ListFieldLevelEncryptionConfigsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFieldLevelEncryptionConfigs request.
func (r *ListFieldLevelEncryptionConfigsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
