// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetReplicationRunsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call. The default value
	// is 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The identifier of the replication job.
	//
	// ReplicationJobId is a required field
	ReplicationJobId *string `locationName:"replicationJobId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetReplicationRunsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReplicationRunsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReplicationRunsInput"}

	if s.ReplicationJobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationJobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetReplicationRunsOutput struct {
	_ struct{} `type:"structure"`

	// The token required to retrieve the next set of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the replication job.
	ReplicationJob *ReplicationJob `locationName:"replicationJob" type:"structure"`

	// Information about the replication runs.
	ReplicationRunList []ReplicationRun `locationName:"replicationRunList" type:"list"`
}

// String returns the string representation
func (s GetReplicationRunsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetReplicationRuns = "GetReplicationRuns"

// GetReplicationRunsRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Describes the replication runs for the specified replication job.
//
//    // Example sending a request using GetReplicationRunsRequest.
//    req := client.GetReplicationRunsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/GetReplicationRuns
func (c *Client) GetReplicationRunsRequest(input *GetReplicationRunsInput) GetReplicationRunsRequest {
	op := &aws.Operation{
		Name:       opGetReplicationRuns,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetReplicationRunsInput{}
	}

	req := c.newRequest(op, input, &GetReplicationRunsOutput{})

	return GetReplicationRunsRequest{Request: req, Input: input, Copy: c.GetReplicationRunsRequest}
}

// GetReplicationRunsRequest is the request type for the
// GetReplicationRuns API operation.
type GetReplicationRunsRequest struct {
	*aws.Request
	Input *GetReplicationRunsInput
	Copy  func(*GetReplicationRunsInput) GetReplicationRunsRequest
}

// Send marshals and sends the GetReplicationRuns API request.
func (r GetReplicationRunsRequest) Send(ctx context.Context) (*GetReplicationRunsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReplicationRunsResponse{
		GetReplicationRunsOutput: r.Request.Data.(*GetReplicationRunsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetReplicationRunsRequestPaginator returns a paginator for GetReplicationRuns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetReplicationRunsRequest(input)
//   p := sms.NewGetReplicationRunsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetReplicationRunsPaginator(req GetReplicationRunsRequest) GetReplicationRunsPaginator {
	return GetReplicationRunsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetReplicationRunsInput
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

// GetReplicationRunsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetReplicationRunsPaginator struct {
	aws.Pager
}

func (p *GetReplicationRunsPaginator) CurrentPage() *GetReplicationRunsOutput {
	return p.Pager.CurrentPage().(*GetReplicationRunsOutput)
}

// GetReplicationRunsResponse is the response type for the
// GetReplicationRuns API operation.
type GetReplicationRunsResponse struct {
	*GetReplicationRunsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReplicationRuns request.
func (r *GetReplicationRunsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
