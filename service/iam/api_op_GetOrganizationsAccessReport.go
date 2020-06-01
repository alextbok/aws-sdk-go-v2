// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetOrganizationsAccessReportInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the request generated by the GenerateOrganizationsAccessReport
	// operation.
	//
	// JobId is a required field
	JobId *string `min:"36" type:"string" required:"true"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The key that is used to sort the results. If you choose the namespace key,
	// the results are returned in alphabetical order. If you choose the time key,
	// the results are sorted numerically by the date and time.
	SortKey SortKeyType `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetOrganizationsAccessReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOrganizationsAccessReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOrganizationsAccessReportInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 36))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetOrganizationsAccessReportOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains details about the most recent attempt to access the
	// service.
	AccessDetails []AccessDetail `type:"list"`

	// Contains information about the reason that the operation failed.
	//
	// This data type is used as a response element in the GetOrganizationsAccessReport,
	// GetServiceLastAccessedDetails, and GetServiceLastAccessedDetailsWithEntities
	// operations.
	ErrorDetails *ErrorDetails `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601),
	// when the generated report job was completed or failed.
	//
	// This field is null if the job is still in progress, as indicated by a job
	// status value of IN_PROGRESS.
	JobCompletionDate *time.Time `type:"timestamp"`

	// The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601),
	// when the report job was created.
	//
	// JobCreationDate is a required field
	JobCreationDate *time.Time `type:"timestamp" required:"true"`

	// The status of the job.
	//
	// JobStatus is a required field
	JobStatus JobStatusType `type:"string" required:"true" enum:"true"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `min:"1" type:"string"`

	// The number of services that the applicable SCPs allow account principals
	// to access.
	NumberOfServicesAccessible *int64 `type:"integer"`

	// The number of services that account principals are allowed but did not attempt
	// to access.
	NumberOfServicesNotAccessed *int64 `type:"integer"`
}

// String returns the string representation
func (s GetOrganizationsAccessReportOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOrganizationsAccessReport = "GetOrganizationsAccessReport"

// GetOrganizationsAccessReportRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves the service last accessed data report for AWS Organizations that
// was previously generated using the GenerateOrganizationsAccessReport operation.
// This operation retrieves the status of your report job and the report contents.
//
// Depending on the parameters that you passed when you generated the report,
// the data returned could include different information. For details, see GenerateOrganizationsAccessReport.
//
// To call this operation, you must be signed in to the master account in your
// organization. SCPs must be enabled for your organization root. You must have
// permissions to perform this operation. For more information, see Refining
// Permissions Using Service Last Accessed Data (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide.
//
// For each service that principals in an account (root users, IAM users, or
// IAM roles) could access using SCPs, the operation returns details about the
// most recent access attempt. If there was no attempt, the service is listed
// without details about the most recent attempt to access the service. If the
// operation fails, it returns the reason that it failed.
//
// By default, the list is sorted by service namespace.
//
//    // Example sending a request using GetOrganizationsAccessReportRequest.
//    req := client.GetOrganizationsAccessReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetOrganizationsAccessReport
func (c *Client) GetOrganizationsAccessReportRequest(input *GetOrganizationsAccessReportInput) GetOrganizationsAccessReportRequest {
	op := &aws.Operation{
		Name:       opGetOrganizationsAccessReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOrganizationsAccessReportInput{}
	}

	req := c.newRequest(op, input, &GetOrganizationsAccessReportOutput{})

	return GetOrganizationsAccessReportRequest{Request: req, Input: input, Copy: c.GetOrganizationsAccessReportRequest}
}

// GetOrganizationsAccessReportRequest is the request type for the
// GetOrganizationsAccessReport API operation.
type GetOrganizationsAccessReportRequest struct {
	*aws.Request
	Input *GetOrganizationsAccessReportInput
	Copy  func(*GetOrganizationsAccessReportInput) GetOrganizationsAccessReportRequest
}

// Send marshals and sends the GetOrganizationsAccessReport API request.
func (r GetOrganizationsAccessReportRequest) Send(ctx context.Context) (*GetOrganizationsAccessReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOrganizationsAccessReportResponse{
		GetOrganizationsAccessReportOutput: r.Request.Data.(*GetOrganizationsAccessReportOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOrganizationsAccessReportResponse is the response type for the
// GetOrganizationsAccessReport API operation.
type GetOrganizationsAccessReportResponse struct {
	*GetOrganizationsAccessReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOrganizationsAccessReport request.
func (r *GetOrganizationsAccessReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
