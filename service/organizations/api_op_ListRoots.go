// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRootsInput struct {
	_ struct{} `type:"structure"`

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific
	// to the operation. If additional items exist beyond the maximum you specify,
	// the NextToken response element is present and has a value (is not null).
	// Include that value as the NextToken request parameter in the next call to
	// the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that
	// you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more
	// output is available. Set this parameter to the value of the previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListRootsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRootsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRootsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRootsOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You
	// should repeat this until the NextToken response element comes back as null.
	NextToken *string `type:"string"`

	// A list of roots that are defined in an organization.
	Roots []Root `type:"list"`
}

// String returns the string representation
func (s ListRootsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRoots = "ListRoots"

// ListRootsRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists the roots that are defined in the current organization.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's master account
// or by a member account that is a delegated administrator for an AWS service.
//
// Policy types can be enabled and disabled in roots. This is distinct from
// whether they're available in the organization. When you enable all features,
// you make policy types available for use in that organization. Individual
// policy types can then be enabled and disabled in a root. To see the availability
// of a policy type in an organization, use DescribeOrganization.
//
//    // Example sending a request using ListRootsRequest.
//    req := client.ListRootsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListRoots
func (c *Client) ListRootsRequest(input *ListRootsInput) ListRootsRequest {
	op := &aws.Operation{
		Name:       opListRoots,
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
		input = &ListRootsInput{}
	}

	req := c.newRequest(op, input, &ListRootsOutput{})
	return ListRootsRequest{Request: req, Input: input, Copy: c.ListRootsRequest}
}

// ListRootsRequest is the request type for the
// ListRoots API operation.
type ListRootsRequest struct {
	*aws.Request
	Input *ListRootsInput
	Copy  func(*ListRootsInput) ListRootsRequest
}

// Send marshals and sends the ListRoots API request.
func (r ListRootsRequest) Send(ctx context.Context) (*ListRootsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRootsResponse{
		ListRootsOutput: r.Request.Data.(*ListRootsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRootsRequestPaginator returns a paginator for ListRoots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRootsRequest(input)
//   p := organizations.NewListRootsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRootsPaginator(req ListRootsRequest) ListRootsPaginator {
	return ListRootsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRootsInput
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

// ListRootsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRootsPaginator struct {
	aws.Pager
}

func (p *ListRootsPaginator) CurrentPage() *ListRootsOutput {
	return p.Pager.CurrentPage().(*ListRootsOutput)
}

// ListRootsResponse is the response type for the
// ListRoots API operation.
type ListRootsResponse struct {
	*ListRootsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRoots request.
func (r *ListRootsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
