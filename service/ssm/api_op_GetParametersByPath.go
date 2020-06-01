// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetParametersByPathInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`

	// Filters to limit the request results.
	ParameterFilters []ParameterStringFilter `type:"list"`

	// The hierarchy for the parameter. Hierarchies start with a forward slash (/)
	// and end with the parameter name. A parameter name hierarchy can have a maximum
	// of 15 levels. Here is an example of a hierarchy: /Finance/Prod/IAD/WinServ2016/license33
	//
	// Path is a required field
	Path *string `min:"1" type:"string" required:"true"`

	// Retrieve all parameters within a hierarchy.
	//
	// If a user has access to a path, then the user can access all levels of that
	// path. For example, if a user has permission to access path /a, then the user
	// can also access /a/b. Even if a user has explicitly been denied access in
	// IAM for parameter /a/b, they can still call the GetParametersByPath API action
	// recursively for /a and view /a/b.
	Recursive *bool `type:"boolean"`

	// Retrieve all parameters in a hierarchy with their value decrypted.
	WithDecryption *bool `type:"boolean"`
}

// String returns the string representation
func (s GetParametersByPathInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetParametersByPathInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetParametersByPathInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Path == nil {
		invalidParams.Add(aws.NewErrParamRequired("Path"))
	}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}
	if s.ParameterFilters != nil {
		for i, v := range s.ParameterFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ParameterFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetParametersByPathOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`

	// A list of parameters found in the specified hierarchy.
	Parameters []Parameter `type:"list"`
}

// String returns the string representation
func (s GetParametersByPathOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetParametersByPath = "GetParametersByPath"

// GetParametersByPathRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieve information about one or more parameters in a specific hierarchy.
//
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified.
// The number of items returned, however, can be between zero and the value
// of MaxResults. If the service reaches an internal limit while processing
// the results, it stops the operation and returns the matching values up to
// that point and a NextToken. You can specify the NextToken in a subsequent
// call to get the next set of results.
//
//    // Example sending a request using GetParametersByPathRequest.
//    req := client.GetParametersByPathRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetParametersByPath
func (c *Client) GetParametersByPathRequest(input *GetParametersByPathInput) GetParametersByPathRequest {
	op := &aws.Operation{
		Name:       opGetParametersByPath,
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
		input = &GetParametersByPathInput{}
	}

	req := c.newRequest(op, input, &GetParametersByPathOutput{})

	return GetParametersByPathRequest{Request: req, Input: input, Copy: c.GetParametersByPathRequest}
}

// GetParametersByPathRequest is the request type for the
// GetParametersByPath API operation.
type GetParametersByPathRequest struct {
	*aws.Request
	Input *GetParametersByPathInput
	Copy  func(*GetParametersByPathInput) GetParametersByPathRequest
}

// Send marshals and sends the GetParametersByPath API request.
func (r GetParametersByPathRequest) Send(ctx context.Context) (*GetParametersByPathResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetParametersByPathResponse{
		GetParametersByPathOutput: r.Request.Data.(*GetParametersByPathOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetParametersByPathRequestPaginator returns a paginator for GetParametersByPath.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetParametersByPathRequest(input)
//   p := ssm.NewGetParametersByPathRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetParametersByPathPaginator(req GetParametersByPathRequest) GetParametersByPathPaginator {
	return GetParametersByPathPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetParametersByPathInput
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

// GetParametersByPathPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetParametersByPathPaginator struct {
	aws.Pager
}

func (p *GetParametersByPathPaginator) CurrentPage() *GetParametersByPathOutput {
	return p.Pager.CurrentPage().(*GetParametersByPathOutput)
}

// GetParametersByPathResponse is the response type for the
// GetParametersByPath API operation.
type GetParametersByPathResponse struct {
	*GetParametersByPathOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetParametersByPath request.
func (r *GetParametersByPathResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
