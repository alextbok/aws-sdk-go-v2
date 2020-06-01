// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListCandidatesForAutoMLJobInput struct {
	_ struct{} `type:"structure"`

	// List the Candidates created for the job by providing the job's name.
	//
	// AutoMLJobName is a required field
	AutoMLJobName *string `min:"1" type:"string" required:"true"`

	// List the Candidates for the job and filter by candidate name.
	CandidateNameEquals *string `min:"1" type:"string"`

	// List the job's Candidates up to a specified limit.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was truncated, you will receive this token. Use
	// it in your next request to receive the next set of results.
	NextToken *string `type:"string"`

	// The parameter by which to sort the results. The default is Descending.
	SortBy CandidateSortBy `type:"string" enum:"true"`

	// The sort order for the results. The default is Ascending.
	SortOrder AutoMLSortOrder `type:"string" enum:"true"`

	// List the Candidates for the job and filter by status.
	StatusEquals CandidateStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListCandidatesForAutoMLJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCandidatesForAutoMLJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCandidatesForAutoMLJobInput"}

	if s.AutoMLJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoMLJobName"))
	}
	if s.AutoMLJobName != nil && len(*s.AutoMLJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoMLJobName", 1))
	}
	if s.CandidateNameEquals != nil && len(*s.CandidateNameEquals) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CandidateNameEquals", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCandidatesForAutoMLJobOutput struct {
	_ struct{} `type:"structure"`

	// Summaries about the Candidates.
	//
	// Candidates is a required field
	Candidates []AutoMLCandidate `type:"list" required:"true"`

	// If the previous response was truncated, you will receive this token. Use
	// it in your next request to receive the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCandidatesForAutoMLJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCandidatesForAutoMLJob = "ListCandidatesForAutoMLJob"

// ListCandidatesForAutoMLJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// List the Candidates created for the job.
//
//    // Example sending a request using ListCandidatesForAutoMLJobRequest.
//    req := client.ListCandidatesForAutoMLJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListCandidatesForAutoMLJob
func (c *Client) ListCandidatesForAutoMLJobRequest(input *ListCandidatesForAutoMLJobInput) ListCandidatesForAutoMLJobRequest {
	op := &aws.Operation{
		Name:       opListCandidatesForAutoMLJob,
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
		input = &ListCandidatesForAutoMLJobInput{}
	}

	req := c.newRequest(op, input, &ListCandidatesForAutoMLJobOutput{})

	return ListCandidatesForAutoMLJobRequest{Request: req, Input: input, Copy: c.ListCandidatesForAutoMLJobRequest}
}

// ListCandidatesForAutoMLJobRequest is the request type for the
// ListCandidatesForAutoMLJob API operation.
type ListCandidatesForAutoMLJobRequest struct {
	*aws.Request
	Input *ListCandidatesForAutoMLJobInput
	Copy  func(*ListCandidatesForAutoMLJobInput) ListCandidatesForAutoMLJobRequest
}

// Send marshals and sends the ListCandidatesForAutoMLJob API request.
func (r ListCandidatesForAutoMLJobRequest) Send(ctx context.Context) (*ListCandidatesForAutoMLJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCandidatesForAutoMLJobResponse{
		ListCandidatesForAutoMLJobOutput: r.Request.Data.(*ListCandidatesForAutoMLJobOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCandidatesForAutoMLJobRequestPaginator returns a paginator for ListCandidatesForAutoMLJob.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCandidatesForAutoMLJobRequest(input)
//   p := sagemaker.NewListCandidatesForAutoMLJobRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCandidatesForAutoMLJobPaginator(req ListCandidatesForAutoMLJobRequest) ListCandidatesForAutoMLJobPaginator {
	return ListCandidatesForAutoMLJobPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCandidatesForAutoMLJobInput
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

// ListCandidatesForAutoMLJobPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCandidatesForAutoMLJobPaginator struct {
	aws.Pager
}

func (p *ListCandidatesForAutoMLJobPaginator) CurrentPage() *ListCandidatesForAutoMLJobOutput {
	return p.Pager.CurrentPage().(*ListCandidatesForAutoMLJobOutput)
}

// ListCandidatesForAutoMLJobResponse is the response type for the
// ListCandidatesForAutoMLJob API operation.
type ListCandidatesForAutoMLJobResponse struct {
	*ListCandidatesForAutoMLJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCandidatesForAutoMLJob request.
func (r *ListCandidatesForAutoMLJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
