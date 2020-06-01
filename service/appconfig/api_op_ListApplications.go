// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListApplicationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `location:"querystring" locationName:"max_results" min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `location:"querystring" locationName:"next_token" min:"1" type:"string"`
}

// String returns the string representation
func (s ListApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListApplicationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListApplicationsInput"}
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
func (s ListApplicationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max_results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next_token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// The elements from this collection.
	Items []Application `type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApplicationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListApplications = "ListApplications"

// ListApplicationsRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// List all applications in your AWS account.
//
//    // Example sending a request using ListApplicationsRequest.
//    req := client.ListApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/ListApplications
func (c *Client) ListApplicationsRequest(input *ListApplicationsInput) ListApplicationsRequest {
	op := &aws.Operation{
		Name:       opListApplications,
		HTTPMethod: "GET",
		HTTPPath:   "/applications",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListApplicationsInput{}
	}

	req := c.newRequest(op, input, &ListApplicationsOutput{})

	return ListApplicationsRequest{Request: req, Input: input, Copy: c.ListApplicationsRequest}
}

// ListApplicationsRequest is the request type for the
// ListApplications API operation.
type ListApplicationsRequest struct {
	*aws.Request
	Input *ListApplicationsInput
	Copy  func(*ListApplicationsInput) ListApplicationsRequest
}

// Send marshals and sends the ListApplications API request.
func (r ListApplicationsRequest) Send(ctx context.Context) (*ListApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApplicationsResponse{
		ListApplicationsOutput: r.Request.Data.(*ListApplicationsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListApplicationsRequestPaginator returns a paginator for ListApplications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListApplicationsRequest(input)
//   p := appconfig.NewListApplicationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListApplicationsPaginator(req ListApplicationsRequest) ListApplicationsPaginator {
	return ListApplicationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListApplicationsInput
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

// ListApplicationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListApplicationsPaginator struct {
	aws.Pager
}

func (p *ListApplicationsPaginator) CurrentPage() *ListApplicationsOutput {
	return p.Pager.CurrentPage().(*ListApplicationsOutput)
}

// ListApplicationsResponse is the response type for the
// ListApplications API operation.
type ListApplicationsResponse struct {
	*ListApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApplications request.
func (r *ListApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
