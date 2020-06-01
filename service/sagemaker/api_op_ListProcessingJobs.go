// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListProcessingJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns only processing jobs modified after the specified time.
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only processing jobs modified before the specified
	// time.
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of processing jobs to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the processing job name. This filter returns only processing
	// jobs whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of the previous ListProcessingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of processing jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is CreationTime.
	SortBy SortBy `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder SortOrder `type:"string" enum:"true"`

	// A filter that retrieves only processing jobs with a specific status.
	StatusEquals ProcessingJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListProcessingJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProcessingJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProcessingJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListProcessingJobsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of processing jobs, use it in the subsequent request.
	NextToken *string `type:"string"`

	// An array of ProcessingJobSummary objects, each listing a processing job.
	//
	// ProcessingJobSummaries is a required field
	ProcessingJobSummaries []ProcessingJobSummary `type:"list" required:"true"`
}

// String returns the string representation
func (s ListProcessingJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListProcessingJobs = "ListProcessingJobs"

// ListProcessingJobsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists processing jobs that satisfy various filters.
//
//    // Example sending a request using ListProcessingJobsRequest.
//    req := client.ListProcessingJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListProcessingJobs
func (c *Client) ListProcessingJobsRequest(input *ListProcessingJobsInput) ListProcessingJobsRequest {
	op := &aws.Operation{
		Name:       opListProcessingJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListProcessingJobsInput{}
	}

	req := c.newRequest(op, input, &ListProcessingJobsOutput{})

	return ListProcessingJobsRequest{Request: req, Input: input, Copy: c.ListProcessingJobsRequest}
}

// ListProcessingJobsRequest is the request type for the
// ListProcessingJobs API operation.
type ListProcessingJobsRequest struct {
	*aws.Request
	Input *ListProcessingJobsInput
	Copy  func(*ListProcessingJobsInput) ListProcessingJobsRequest
}

// Send marshals and sends the ListProcessingJobs API request.
func (r ListProcessingJobsRequest) Send(ctx context.Context) (*ListProcessingJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProcessingJobsResponse{
		ListProcessingJobsOutput: r.Request.Data.(*ListProcessingJobsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProcessingJobsRequestPaginator returns a paginator for ListProcessingJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProcessingJobsRequest(input)
//   p := sagemaker.NewListProcessingJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProcessingJobsPaginator(req ListProcessingJobsRequest) ListProcessingJobsPaginator {
	return ListProcessingJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProcessingJobsInput
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

// ListProcessingJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProcessingJobsPaginator struct {
	aws.Pager
}

func (p *ListProcessingJobsPaginator) CurrentPage() *ListProcessingJobsOutput {
	return p.Pager.CurrentPage().(*ListProcessingJobsOutput)
}

// ListProcessingJobsResponse is the response type for the
// ListProcessingJobs API operation.
type ListProcessingJobsResponse struct {
	*ListProcessingJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProcessingJobs request.
func (r *ListProcessingJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
