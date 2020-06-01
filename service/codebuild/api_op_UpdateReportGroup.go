// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateReportGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the report group to update.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"1" type:"string" required:"true"`

	// Used to specify an updated export type. Valid values are:
	//
	//    * S3: The report results are exported to an S3 bucket.
	//
	//    * NO_EXPORT: The report results are not exported.
	ExportConfig *ReportExportConfig `locationName:"exportConfig" type:"structure"`

	// An updated list of tag key and value pairs associated with this report group.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild
	// report group tags.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s UpdateReportGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateReportGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateReportGroupInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 1))
	}
	if s.ExportConfig != nil {
		if err := s.ExportConfig.Validate(); err != nil {
			invalidParams.AddNested("ExportConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateReportGroupOutput struct {
	_ struct{} `type:"structure"`

	// Information about the updated report group.
	ReportGroup *ReportGroup `locationName:"reportGroup" type:"structure"`
}

// String returns the string representation
func (s UpdateReportGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateReportGroup = "UpdateReportGroup"

// UpdateReportGroupRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Updates a report group.
//
//    // Example sending a request using UpdateReportGroupRequest.
//    req := client.UpdateReportGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/UpdateReportGroup
func (c *Client) UpdateReportGroupRequest(input *UpdateReportGroupInput) UpdateReportGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateReportGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateReportGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateReportGroupOutput{})

	return UpdateReportGroupRequest{Request: req, Input: input, Copy: c.UpdateReportGroupRequest}
}

// UpdateReportGroupRequest is the request type for the
// UpdateReportGroup API operation.
type UpdateReportGroupRequest struct {
	*aws.Request
	Input *UpdateReportGroupInput
	Copy  func(*UpdateReportGroupInput) UpdateReportGroupRequest
}

// Send marshals and sends the UpdateReportGroup API request.
func (r UpdateReportGroupRequest) Send(ctx context.Context) (*UpdateReportGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateReportGroupResponse{
		UpdateReportGroupOutput: r.Request.Data.(*UpdateReportGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateReportGroupResponse is the response type for the
// UpdateReportGroup API operation.
type UpdateReportGroupResponse struct {
	*UpdateReportGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateReportGroup request.
func (r *UpdateReportGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
