// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListJobsInput struct {
	_ struct{} `type:"structure"`

	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The List Jobs request returns jobs that match the statuses listed in this
	// element.
	JobStatuses []JobStatus `location:"querystring" locationName:"jobStatuses" type:"list"`

	// The maximum number of jobs that Amazon S3 will include in the List Jobs response.
	// If there are more jobs than this number, the response will include a pagination
	// token in the NextToken field to enable you to retrieve the next page of results.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A pagination token to request the next page of results. Use the token that
	// Amazon S3 returned in the NextToken element of the ListJobsResult from the
	// previous List Jobs request.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	if s.JobStatuses != nil {
		v := s.JobStatuses

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "jobStatuses", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.StringValue(v1))
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.StringValue(v), metadata)
	}
	return nil
}

type ListJobsOutput struct {
	_ struct{} `type:"structure"`

	// The list of current jobs and jobs that have ended within the last 30 days.
	Jobs []JobListDescriptor `type:"list"`

	// If the List Jobs request produced more than the maximum number of results,
	// you can pass this value into a subsequent List Jobs request in order to retrieve
	// the next page of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Jobs != nil {
		v := s.Jobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Jobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListJobs = "ListJobs"

// ListJobsRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Lists current Amazon S3 Batch Operations jobs and jobs that have ended within
// the last 30 days for the AWS account making the request. For more information,
// see Amazon S3 Batch Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// Related actions include:
//
//    * CreateJob
//
//    * DescribeJob
//
//    * UpdateJobPriority
//
//    * UpdateJobStatus
//
//    // Example sending a request using ListJobsRequest.
//    req := client.ListJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/ListJobs
func (c *Client) ListJobsRequest(input *ListJobsInput) ListJobsRequest {
	op := &aws.Operation{
		Name:       opListJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/v20180820/jobs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListJobsInput{}
	}

	req := c.newRequest(op, input, &ListJobsOutput{})
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))

	return ListJobsRequest{Request: req, Input: input, Copy: c.ListJobsRequest}
}

// ListJobsRequest is the request type for the
// ListJobs API operation.
type ListJobsRequest struct {
	*aws.Request
	Input *ListJobsInput
	Copy  func(*ListJobsInput) ListJobsRequest
}

// Send marshals and sends the ListJobs API request.
func (r ListJobsRequest) Send(ctx context.Context) (*ListJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobsResponse{
		ListJobsOutput: r.Request.Data.(*ListJobsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListJobsRequestPaginator returns a paginator for ListJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListJobsRequest(input)
//   p := s3control.NewListJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListJobsPaginator(req ListJobsRequest) ListJobsPaginator {
	return ListJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListJobsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListJobsPaginator struct {
	aws.Pager
}

func (p *ListJobsPaginator) CurrentPage() *ListJobsOutput {
	return p.Pager.CurrentPage().(*ListJobsOutput)
}

// ListJobsResponse is the response type for the
// ListJobs API operation.
type ListJobsResponse struct {
	*ListJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobs request.
func (r *ListJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
