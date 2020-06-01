// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeRootFoldersInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token.
	//
	// AuthenticationToken is a required field
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" required:"true" sensitive:"true"`

	// The maximum number of items to return.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeRootFoldersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRootFoldersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRootFoldersInput"}

	if s.AuthenticationToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationToken"))
	}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRootFoldersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

type DescribeRootFoldersOutput struct {
	_ struct{} `type:"structure"`

	// The user's special folders.
	Folders []FolderMetadata `type:"list"`

	// The marker for the next set of results.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeRootFoldersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRootFoldersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Folders != nil {
		v := s.Folders

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Folders", metadata)
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

const opDescribeRootFolders = "DescribeRootFolders"

// DescribeRootFoldersRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Describes the current user's special folders; the RootFolder and the RecycleBin.
// RootFolder is the root of user's files and folders and RecycleBin is the
// root of recycled items. This is not a valid action for SigV4 (administrative
// API) clients.
//
// This action requires an authentication token. To get an authentication token,
// register an application with Amazon WorkDocs. For more information, see Authentication
// and Access Control for User Applications (https://docs.aws.amazon.com/workdocs/latest/developerguide/wd-auth-user.html)
// in the Amazon WorkDocs Developer Guide.
//
//    // Example sending a request using DescribeRootFoldersRequest.
//    req := client.DescribeRootFoldersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeRootFolders
func (c *Client) DescribeRootFoldersRequest(input *DescribeRootFoldersInput) DescribeRootFoldersRequest {
	op := &aws.Operation{
		Name:       opDescribeRootFolders,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/me/root",
	}

	if input == nil {
		input = &DescribeRootFoldersInput{}
	}

	req := c.newRequest(op, input, &DescribeRootFoldersOutput{})

	return DescribeRootFoldersRequest{Request: req, Input: input, Copy: c.DescribeRootFoldersRequest}
}

// DescribeRootFoldersRequest is the request type for the
// DescribeRootFolders API operation.
type DescribeRootFoldersRequest struct {
	*aws.Request
	Input *DescribeRootFoldersInput
	Copy  func(*DescribeRootFoldersInput) DescribeRootFoldersRequest
}

// Send marshals and sends the DescribeRootFolders API request.
func (r DescribeRootFoldersRequest) Send(ctx context.Context) (*DescribeRootFoldersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRootFoldersResponse{
		DescribeRootFoldersOutput: r.Request.Data.(*DescribeRootFoldersOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRootFoldersResponse is the response type for the
// DescribeRootFolders API operation.
type DescribeRootFoldersResponse struct {
	*DescribeRootFoldersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRootFolders request.
func (r *DescribeRootFoldersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
