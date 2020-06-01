// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsInput struct {
	_ struct{} `type:"structure"`

	// Maximum number of tags to return.
	MaxResults *int64 `min:"50" type:"integer"`

	// If the response to the previous ListTags request is truncated, Amazon SageMaker
	// returns this token. To retrieve the next set of tags, use it in the subsequent
	// request.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) of the resource whose tags you want to retrieve.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsInput"}
	if s.MaxResults != nil && *s.MaxResults < 50 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 50))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsOutput struct {
	_ struct{} `type:"structure"`

	// If response is truncated, Amazon SageMaker includes a token in the response.
	// You can use this token in your subsequent request to fetch next set of tokens.
	NextToken *string `type:"string"`

	// An array of Tag objects, each with a tag key and a value.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s ListTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTags = "ListTags"

// ListTagsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns the tags for the specified Amazon SageMaker resource.
//
//    // Example sending a request using ListTagsRequest.
//    req := client.ListTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTags
func (c *Client) ListTagsRequest(input *ListTagsInput) ListTagsRequest {
	op := &aws.Operation{
		Name:       opListTags,
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
		input = &ListTagsInput{}
	}

	req := c.newRequest(op, input, &ListTagsOutput{})

	return ListTagsRequest{Request: req, Input: input, Copy: c.ListTagsRequest}
}

// ListTagsRequest is the request type for the
// ListTags API operation.
type ListTagsRequest struct {
	*aws.Request
	Input *ListTagsInput
	Copy  func(*ListTagsInput) ListTagsRequest
}

// Send marshals and sends the ListTags API request.
func (r ListTagsRequest) Send(ctx context.Context) (*ListTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsResponse{
		ListTagsOutput: r.Request.Data.(*ListTagsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTagsRequestPaginator returns a paginator for ListTags.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTagsRequest(input)
//   p := sagemaker.NewListTagsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTagsPaginator(req ListTagsRequest) ListTagsPaginator {
	return ListTagsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTagsInput
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

// ListTagsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTagsPaginator struct {
	aws.Pager
}

func (p *ListTagsPaginator) CurrentPage() *ListTagsOutput {
	return p.Pager.CurrentPage().(*ListTagsOutput)
}

// ListTagsResponse is the response type for the
// ListTags API operation.
type ListTagsResponse struct {
	*ListTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTags request.
func (r *ListTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
