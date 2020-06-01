// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListFunctionsInput struct {
	_ struct{} `type:"structure"`

	// Set to ALL to include entries for all published versions of each function.
	FunctionVersion FunctionVersion `location:"querystring" locationName:"FunctionVersion" type:"string" enum:"true"`

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// For Lambda@Edge functions, the AWS Region of the master function. For example,
	// us-east-1 filters the list of functions to only include Lambda@Edge functions
	// replicated from a master function in US East (N. Virginia). If specified,
	// you must set FunctionVersion to ALL.
	MasterRegion *string `location:"querystring" locationName:"MasterRegion" type:"string"`

	// The maximum number of functions to return.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListFunctionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFunctionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFunctionsInput"}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.FunctionVersion) > 0 {
		v := s.FunctionVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "FunctionVersion", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MasterRegion != nil {
		v := *s.MasterRegion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MasterRegion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// A list of Lambda functions.
type ListFunctionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of Lambda functions.
	Functions []FunctionConfiguration `type:"list"`

	// The pagination token that's included if more results are available.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s ListFunctionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFunctionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Functions != nil {
		v := s.Functions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Functions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFunctions = "ListFunctions"

// ListFunctionsRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns a list of Lambda functions, with the version-specific configuration
// of each. Lambda returns up to 50 functions per call.
//
// Set FunctionVersion to ALL to include all published versions of each function
// in addition to the unpublished version. To get more information about a function
// or version, use GetFunction.
//
//    // Example sending a request using ListFunctionsRequest.
//    req := client.ListFunctionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListFunctions
func (c *Client) ListFunctionsRequest(input *ListFunctionsInput) ListFunctionsRequest {
	op := &aws.Operation{
		Name:       opListFunctions,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/functions/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFunctionsInput{}
	}

	req := c.newRequest(op, input, &ListFunctionsOutput{})

	return ListFunctionsRequest{Request: req, Input: input, Copy: c.ListFunctionsRequest}
}

// ListFunctionsRequest is the request type for the
// ListFunctions API operation.
type ListFunctionsRequest struct {
	*aws.Request
	Input *ListFunctionsInput
	Copy  func(*ListFunctionsInput) ListFunctionsRequest
}

// Send marshals and sends the ListFunctions API request.
func (r ListFunctionsRequest) Send(ctx context.Context) (*ListFunctionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFunctionsResponse{
		ListFunctionsOutput: r.Request.Data.(*ListFunctionsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFunctionsRequestPaginator returns a paginator for ListFunctions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFunctionsRequest(input)
//   p := lambda.NewListFunctionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFunctionsPaginator(req ListFunctionsRequest) ListFunctionsPaginator {
	return ListFunctionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFunctionsInput
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

// ListFunctionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFunctionsPaginator struct {
	aws.Pager
}

func (p *ListFunctionsPaginator) CurrentPage() *ListFunctionsOutput {
	return p.Pager.CurrentPage().(*ListFunctionsOutput)
}

// ListFunctionsResponse is the response type for the
// ListFunctions API operation.
type ListFunctionsResponse struct {
	*ListFunctionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFunctions request.
func (r *ListFunctionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
