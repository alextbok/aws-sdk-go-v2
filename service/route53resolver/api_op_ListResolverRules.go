// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListResolverRulesInput struct {
	_ struct{} `type:"structure"`

	// An optional specification to return a subset of resolver rules, such as all
	// resolver rules that are associated with the same resolver endpoint.
	//
	// If you submit a second or subsequent ListResolverRules request and specify
	// the NextToken parameter, you must use the same values for Filters, if any,
	// as in the previous request.
	Filters []Filter `type:"list"`

	// The maximum number of resolver rules that you want to return in the response
	// to a ListResolverRules request. If you don't specify a value for MaxResults,
	// Resolver returns up to 100 resolver rules.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first ListResolverRules request, omit this value.
	//
	// If you have more than MaxResults resolver rules, you can submit another ListResolverRules
	// request to get the next group of resolver rules. In the next request, specify
	// the value of NextToken from the previous response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListResolverRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResolverRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResolverRulesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResolverRulesOutput struct {
	_ struct{} `type:"structure"`

	// The value that you specified for MaxResults in the request.
	MaxResults *int64 `min:"1" type:"integer"`

	// If more than MaxResults resolver rules match the specified criteria, you
	// can submit another ListResolverRules request to get the next group of results.
	// In the next request, specify the value of NextToken from the previous response.
	NextToken *string `type:"string"`

	// The resolver rules that were created using the current AWS account and that
	// match the specified filters, if any.
	ResolverRules []ResolverRule `type:"list"`
}

// String returns the string representation
func (s ListResolverRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListResolverRules = "ListResolverRules"

// ListResolverRulesRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Lists the resolver rules that were created using the current AWS account.
//
//    // Example sending a request using ListResolverRulesRequest.
//    req := client.ListResolverRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ListResolverRules
func (c *Client) ListResolverRulesRequest(input *ListResolverRulesInput) ListResolverRulesRequest {
	op := &aws.Operation{
		Name:       opListResolverRules,
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
		input = &ListResolverRulesInput{}
	}

	req := c.newRequest(op, input, &ListResolverRulesOutput{})

	return ListResolverRulesRequest{Request: req, Input: input, Copy: c.ListResolverRulesRequest}
}

// ListResolverRulesRequest is the request type for the
// ListResolverRules API operation.
type ListResolverRulesRequest struct {
	*aws.Request
	Input *ListResolverRulesInput
	Copy  func(*ListResolverRulesInput) ListResolverRulesRequest
}

// Send marshals and sends the ListResolverRules API request.
func (r ListResolverRulesRequest) Send(ctx context.Context) (*ListResolverRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResolverRulesResponse{
		ListResolverRulesOutput: r.Request.Data.(*ListResolverRulesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResolverRulesRequestPaginator returns a paginator for ListResolverRules.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResolverRulesRequest(input)
//   p := route53resolver.NewListResolverRulesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResolverRulesPaginator(req ListResolverRulesRequest) ListResolverRulesPaginator {
	return ListResolverRulesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListResolverRulesInput
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

// ListResolverRulesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResolverRulesPaginator struct {
	aws.Pager
}

func (p *ListResolverRulesPaginator) CurrentPage() *ListResolverRulesOutput {
	return p.Pager.CurrentPage().(*ListResolverRulesOutput)
}

// ListResolverRulesResponse is the response type for the
// ListResolverRules API operation.
type ListResolverRulesResponse struct {
	*ListResolverRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResolverRules request.
func (r *ListResolverRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
