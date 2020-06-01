// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The ListTagsForDomainRequest includes the following elements.
type ListTagsForDomainInput struct {
	_ struct{} `type:"structure"`

	// The domain for which you want to get a list of tags.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The ListTagsForDomain response includes the following elements.
type ListTagsForDomainOutput struct {
	_ struct{} `type:"structure"`

	// A list of the tags that are associated with the specified domain.
	//
	// TagList is a required field
	TagList []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s ListTagsForDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForDomain = "ListTagsForDomain"

// ListTagsForDomainRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation returns all of the tags that are associated with the specified
// domain.
//
// All tag operations are eventually consistent; subsequent operations might
// not immediately represent all issued operations.
//
//    // Example sending a request using ListTagsForDomainRequest.
//    req := client.ListTagsForDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ListTagsForDomain
func (c *Client) ListTagsForDomainRequest(input *ListTagsForDomainInput) ListTagsForDomainRequest {
	op := &aws.Operation{
		Name:       opListTagsForDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForDomainInput{}
	}

	req := c.newRequest(op, input, &ListTagsForDomainOutput{})

	return ListTagsForDomainRequest{Request: req, Input: input, Copy: c.ListTagsForDomainRequest}
}

// ListTagsForDomainRequest is the request type for the
// ListTagsForDomain API operation.
type ListTagsForDomainRequest struct {
	*aws.Request
	Input *ListTagsForDomainInput
	Copy  func(*ListTagsForDomainInput) ListTagsForDomainRequest
}

// Send marshals and sends the ListTagsForDomain API request.
func (r ListTagsForDomainRequest) Send(ctx context.Context) (*ListTagsForDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForDomainResponse{
		ListTagsForDomainOutput: r.Request.Data.(*ListTagsForDomainOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForDomainResponse is the response type for the
// ListTagsForDomain API operation.
type ListTagsForDomainResponse struct {
	*ListTagsForDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForDomain request.
func (r *ListTagsForDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
