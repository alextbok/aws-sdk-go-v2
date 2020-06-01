// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListAuditFindingsInput struct {
	_ struct{} `type:"structure"`

	// A filter to limit results to the findings for the specified audit check.
	CheckName *string `locationName:"checkName" type:"string"`

	// A filter to limit results to those found before the specified time. You must
	// specify either the startTime and endTime or the taskId, but not both.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information identifying the noncompliant resource.
	ResourceIdentifier *ResourceIdentifier `locationName:"resourceIdentifier" type:"structure"`

	// A filter to limit results to those found after the specified time. You must
	// specify either the startTime and endTime or the taskId, but not both.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`

	// A filter to limit results to the audit with the specified ID. You must specify
	// either the taskId or the startTime and endTime, but not both.
	TaskId *string `locationName:"taskId" min:"1" type:"string"`
}

// String returns the string representation
func (s ListAuditFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAuditFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAuditFindingsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.TaskId != nil && len(*s.TaskId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskId", 1))
	}
	if s.ResourceIdentifier != nil {
		if err := s.ResourceIdentifier.Validate(); err != nil {
			invalidParams.AddNested("ResourceIdentifier", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAuditFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CheckName != nil {
		v := *s.CheckName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "checkName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceIdentifier != nil {
		v := s.ResourceIdentifier

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourceIdentifier", v, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "taskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListAuditFindingsOutput struct {
	_ struct{} `type:"structure"`

	// The findings (results) of the audit.
	Findings []AuditFinding `locationName:"findings" type:"list"`

	// A token that can be used to retrieve the next set of results, or null if
	// there are no additional results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAuditFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAuditFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Findings != nil {
		v := s.Findings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListAuditFindings = "ListAuditFindings"

// ListAuditFindingsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the findings (results) of a Device Defender audit or of the audits
// performed during a specified time period. (Findings are retained for 180
// days.)
//
//    // Example sending a request using ListAuditFindingsRequest.
//    req := client.ListAuditFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListAuditFindingsRequest(input *ListAuditFindingsInput) ListAuditFindingsRequest {
	op := &aws.Operation{
		Name:       opListAuditFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/audit/findings",
	}

	if input == nil {
		input = &ListAuditFindingsInput{}
	}

	req := c.newRequest(op, input, &ListAuditFindingsOutput{})

	return ListAuditFindingsRequest{Request: req, Input: input, Copy: c.ListAuditFindingsRequest}
}

// ListAuditFindingsRequest is the request type for the
// ListAuditFindings API operation.
type ListAuditFindingsRequest struct {
	*aws.Request
	Input *ListAuditFindingsInput
	Copy  func(*ListAuditFindingsInput) ListAuditFindingsRequest
}

// Send marshals and sends the ListAuditFindings API request.
func (r ListAuditFindingsRequest) Send(ctx context.Context) (*ListAuditFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAuditFindingsResponse{
		ListAuditFindingsOutput: r.Request.Data.(*ListAuditFindingsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAuditFindingsResponse is the response type for the
// ListAuditFindings API operation.
type ListAuditFindingsResponse struct {
	*ListAuditFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAuditFindings request.
func (r *ListAuditFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
