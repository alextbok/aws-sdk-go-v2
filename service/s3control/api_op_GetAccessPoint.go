// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetAccessPointInput struct {
	_ struct{} `type:"structure"`

	// The account ID for the account that owns the specified access point.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The name of the access point whose configuration information you want to
	// retrieve.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccessPointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccessPointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccessPointInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccessPointInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetAccessPointOutput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket associated with the specified access point.
	Bucket *string `min:"3" type:"string"`

	// The date and time when the specified access point was created.
	CreationDate *time.Time `type:"timestamp"`

	// The name of the specified access point.
	Name *string `min:"3" type:"string"`

	// Indicates whether this access point allows access from the public internet.
	// If VpcConfiguration is specified for this access point, then NetworkOrigin
	// is VPC, and the access point doesn't allow access from the public internet.
	// Otherwise, NetworkOrigin is Internet, and the access point allows access
	// from the public internet, subject to the access point and bucket access policies.
	NetworkOrigin NetworkOrigin `type:"string" enum:"true"`

	// The PublicAccessBlock configuration that you want to apply to this Amazon
	// S3 bucket. You can enable the configuration options in any combination. For
	// more information about when Amazon S3 considers a bucket or object public,
	// see The Meaning of "Public" (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `type:"structure"`

	// Contains the virtual private cloud (VPC) configuration for the specified
	// access point.
	VpcConfiguration *VpcConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetAccessPointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccessPointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if len(s.NetworkOrigin) > 0 {
		v := s.NetworkOrigin

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NetworkOrigin", v, metadata)
	}
	if s.PublicAccessBlockConfiguration != nil {
		v := s.PublicAccessBlockConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PublicAccessBlockConfiguration", v, metadata)
	}
	if s.VpcConfiguration != nil {
		v := s.VpcConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConfiguration", v, metadata)
	}
	return nil
}

const opGetAccessPoint = "GetAccessPoint"

// GetAccessPointRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Returns configuration information about the specified access point.
//
//    // Example sending a request using GetAccessPointRequest.
//    req := client.GetAccessPointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/GetAccessPoint
func (c *Client) GetAccessPointRequest(input *GetAccessPointInput) GetAccessPointRequest {
	op := &aws.Operation{
		Name:       opGetAccessPoint,
		HTTPMethod: "GET",
		HTTPPath:   "/v20180820/accesspoint/{name}",
	}

	if input == nil {
		input = &GetAccessPointInput{}
	}

	req := c.newRequest(op, input, &GetAccessPointOutput{})
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))

	return GetAccessPointRequest{Request: req, Input: input, Copy: c.GetAccessPointRequest}
}

// GetAccessPointRequest is the request type for the
// GetAccessPoint API operation.
type GetAccessPointRequest struct {
	*aws.Request
	Input *GetAccessPointInput
	Copy  func(*GetAccessPointInput) GetAccessPointRequest
}

// Send marshals and sends the GetAccessPoint API request.
func (r GetAccessPointRequest) Send(ctx context.Context) (*GetAccessPointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccessPointResponse{
		GetAccessPointOutput: r.Request.Data.(*GetAccessPointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccessPointResponse is the response type for the
// GetAccessPoint API operation.
type GetAccessPointResponse struct {
	*GetAccessPointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccessPoint request.
func (r *GetAccessPointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
