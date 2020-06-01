// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetObjectInformationInput struct {
	_ struct{} `type:"structure"`

	// The consistency level at which to retrieve the object information.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The ARN of the directory being retrieved.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// A reference to the object.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetObjectInformationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectInformationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectInformationInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectInformationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetObjectInformationOutput struct {
	_ struct{} `type:"structure"`

	// The ObjectIdentifier of the specified object.
	ObjectIdentifier *string `type:"string"`

	// The facets attached to the specified object. Although the response does not
	// include minor version information, the most recently applied minor version
	// of each Facet is in effect. See GetAppliedSchemaVersion for details.
	SchemaFacets []SchemaFacet `type:"list"`
}

// String returns the string representation
func (s GetObjectInformationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectInformationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ObjectIdentifier != nil {
		v := *s.ObjectIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ObjectIdentifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaFacets != nil {
		v := s.SchemaFacets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SchemaFacets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetObjectInformation = "GetObjectInformation"

// GetObjectInformationRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Retrieves metadata about an object.
//
//    // Example sending a request using GetObjectInformationRequest.
//    req := client.GetObjectInformationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetObjectInformation
func (c *Client) GetObjectInformationRequest(input *GetObjectInformationInput) GetObjectInformationRequest {
	op := &aws.Operation{
		Name:       opGetObjectInformation,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/information",
	}

	if input == nil {
		input = &GetObjectInformationInput{}
	}

	req := c.newRequest(op, input, &GetObjectInformationOutput{})

	return GetObjectInformationRequest{Request: req, Input: input, Copy: c.GetObjectInformationRequest}
}

// GetObjectInformationRequest is the request type for the
// GetObjectInformation API operation.
type GetObjectInformationRequest struct {
	*aws.Request
	Input *GetObjectInformationInput
	Copy  func(*GetObjectInformationInput) GetObjectInformationRequest
}

// Send marshals and sends the GetObjectInformation API request.
func (r GetObjectInformationRequest) Send(ctx context.Context) (*GetObjectInformationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectInformationResponse{
		GetObjectInformationOutput: r.Request.Data.(*GetObjectInformationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectInformationResponse is the response type for the
// GetObjectInformation API operation.
type GetObjectInformationResponse struct {
	*GetObjectInformationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectInformation request.
func (r *GetObjectInformationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
