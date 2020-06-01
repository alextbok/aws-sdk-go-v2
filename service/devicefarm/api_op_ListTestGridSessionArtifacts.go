// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTestGridSessionArtifactsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned by a request.
	MaxResult *int64 `locationName:"maxResult" min:"1" type:"integer"`

	// Pagination token.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// The ARN of a TestGridSession.
	//
	// SessionArn is a required field
	SessionArn *string `locationName:"sessionArn" min:"32" type:"string" required:"true"`

	// Limit results to a specified type of artifact.
	Type TestGridSessionArtifactCategory `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTestGridSessionArtifactsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTestGridSessionArtifactsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTestGridSessionArtifactsInput"}
	if s.MaxResult != nil && *s.MaxResult < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResult", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if s.SessionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SessionArn"))
	}
	if s.SessionArn != nil && len(*s.SessionArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("SessionArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTestGridSessionArtifactsOutput struct {
	_ struct{} `type:"structure"`

	// A list of test grid session artifacts for a TestGridSession.
	Artifacts []TestGridSessionArtifact `locationName:"artifacts" type:"list"`

	// Pagination token.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListTestGridSessionArtifactsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTestGridSessionArtifacts = "ListTestGridSessionArtifacts"

// ListTestGridSessionArtifactsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Retrieves a list of artifacts created during the session.
//
//    // Example sending a request using ListTestGridSessionArtifactsRequest.
//    req := client.ListTestGridSessionArtifactsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListTestGridSessionArtifacts
func (c *Client) ListTestGridSessionArtifactsRequest(input *ListTestGridSessionArtifactsInput) ListTestGridSessionArtifactsRequest {
	op := &aws.Operation{
		Name:       opListTestGridSessionArtifacts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResult",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTestGridSessionArtifactsInput{}
	}

	req := c.newRequest(op, input, &ListTestGridSessionArtifactsOutput{})

	return ListTestGridSessionArtifactsRequest{Request: req, Input: input, Copy: c.ListTestGridSessionArtifactsRequest}
}

// ListTestGridSessionArtifactsRequest is the request type for the
// ListTestGridSessionArtifacts API operation.
type ListTestGridSessionArtifactsRequest struct {
	*aws.Request
	Input *ListTestGridSessionArtifactsInput
	Copy  func(*ListTestGridSessionArtifactsInput) ListTestGridSessionArtifactsRequest
}

// Send marshals and sends the ListTestGridSessionArtifacts API request.
func (r ListTestGridSessionArtifactsRequest) Send(ctx context.Context) (*ListTestGridSessionArtifactsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTestGridSessionArtifactsResponse{
		ListTestGridSessionArtifactsOutput: r.Request.Data.(*ListTestGridSessionArtifactsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTestGridSessionArtifactsRequestPaginator returns a paginator for ListTestGridSessionArtifacts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTestGridSessionArtifactsRequest(input)
//   p := devicefarm.NewListTestGridSessionArtifactsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTestGridSessionArtifactsPaginator(req ListTestGridSessionArtifactsRequest) ListTestGridSessionArtifactsPaginator {
	return ListTestGridSessionArtifactsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTestGridSessionArtifactsInput
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

// ListTestGridSessionArtifactsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTestGridSessionArtifactsPaginator struct {
	aws.Pager
}

func (p *ListTestGridSessionArtifactsPaginator) CurrentPage() *ListTestGridSessionArtifactsOutput {
	return p.Pager.CurrentPage().(*ListTestGridSessionArtifactsOutput)
}

// ListTestGridSessionArtifactsResponse is the response type for the
// ListTestGridSessionArtifacts API operation.
type ListTestGridSessionArtifactsResponse struct {
	*ListTestGridSessionArtifactsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTestGridSessionArtifacts request.
func (r *ListTestGridSessionArtifactsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
