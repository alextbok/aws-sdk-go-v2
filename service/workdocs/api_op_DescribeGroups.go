// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeGroupsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The maximum number of items to return with this call.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`

	// The ID of the organization.
	OrganizationId *string `location:"querystring" locationName:"organizationId" min:"1" type:"string"`

	// A query to describe groups by group name.
	//
	// SearchQuery is a required field
	SearchQuery *string `location:"querystring" locationName:"searchQuery" min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s DescribeGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGroupsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.OrganizationId != nil && len(*s.OrganizationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationId", 1))
	}

	if s.SearchQuery == nil {
		invalidParams.Add(aws.NewErrParamRequired("SearchQuery"))
	}
	if s.SearchQuery != nil && len(*s.SearchQuery) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SearchQuery", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.OrganizationId != nil {
		v := *s.OrganizationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "organizationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SearchQuery != nil {
		v := *s.SearchQuery

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "searchQuery", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The list of groups.
	Groups []GroupMetadata `type:"list"`

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Groups != nil {
		v := s.Groups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Groups", metadata)
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

const opDescribeGroups = "DescribeGroups"

// DescribeGroupsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Describes the groups specified by the query. Groups are defined by the underlying
// Active Directory.
//
//    // Example sending a request using DescribeGroupsRequest.
//    req := client.DescribeGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeGroups
func (c *Client) DescribeGroupsRequest(input *DescribeGroupsInput) DescribeGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/groups",
	}

	if input == nil {
		input = &DescribeGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeGroupsOutput{})

	return DescribeGroupsRequest{Request: req, Input: input, Copy: c.DescribeGroupsRequest}
}

// DescribeGroupsRequest is the request type for the
// DescribeGroups API operation.
type DescribeGroupsRequest struct {
	*aws.Request
	Input *DescribeGroupsInput
	Copy  func(*DescribeGroupsInput) DescribeGroupsRequest
}

// Send marshals and sends the DescribeGroups API request.
func (r DescribeGroupsRequest) Send(ctx context.Context) (*DescribeGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGroupsResponse{
		DescribeGroupsOutput: r.Request.Data.(*DescribeGroupsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGroupsResponse is the response type for the
// DescribeGroups API operation.
type DescribeGroupsResponse struct {
	*DescribeGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGroups request.
func (r *DescribeGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
