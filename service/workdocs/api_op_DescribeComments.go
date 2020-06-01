// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeCommentsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// The maximum number of items to return.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The marker for the next set of results. This marker was received from a previous
	// call.
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`

	// The ID of the document version.
	//
	// VersionId is a required field
	VersionId *string `location:"uri" locationName:"VersionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCommentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCommentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCommentsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}
	if s.VersionId != nil && len(*s.VersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCommentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentId != nil {
		v := *s.DocumentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DocumentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeCommentsOutput struct {
	_ struct{} `type:"structure"`

	// The list of comments for the specified document version.
	Comments []Comment `type:"list"`

	// The marker for the next set of results. This marker was received from a previous
	// call.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeCommentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCommentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Comments != nil {
		v := s.Comments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Comments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeComments = "DescribeComments"

// DescribeCommentsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// List all the comments for the specified document version.
//
//    // Example sending a request using DescribeCommentsRequest.
//    req := client.DescribeCommentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeComments
func (c *Client) DescribeCommentsRequest(input *DescribeCommentsInput) DescribeCommentsRequest {
	op := &aws.Operation{
		Name:       opDescribeComments,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/documents/{DocumentId}/versions/{VersionId}/comments",
	}

	if input == nil {
		input = &DescribeCommentsInput{}
	}

	req := c.newRequest(op, input, &DescribeCommentsOutput{})

	return DescribeCommentsRequest{Request: req, Input: input, Copy: c.DescribeCommentsRequest}
}

// DescribeCommentsRequest is the request type for the
// DescribeComments API operation.
type DescribeCommentsRequest struct {
	*aws.Request
	Input *DescribeCommentsInput
	Copy  func(*DescribeCommentsInput) DescribeCommentsRequest
}

// Send marshals and sends the DescribeComments API request.
func (r DescribeCommentsRequest) Send(ctx context.Context) (*DescribeCommentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCommentsResponse{
		DescribeCommentsOutput: r.Request.Data.(*DescribeCommentsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCommentsResponse is the response type for the
// DescribeComments API operation.
type DescribeCommentsResponse struct {
	*DescribeCommentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeComments request.
func (r *DescribeCommentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
