// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSystemTemplateRevisionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the system template.
	//
	// The ID should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	//
	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetSystemTemplateRevisionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSystemTemplateRevisionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSystemTemplateRevisionsInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetSystemTemplateRevisionsOutput struct {
	_ struct{} `type:"structure"`

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// An array of objects that contain summary data about the system template revisions.
	Summaries []SystemTemplateSummary `locationName:"summaries" type:"list"`
}

// String returns the string representation
func (s GetSystemTemplateRevisionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSystemTemplateRevisions = "GetSystemTemplateRevisions"

// GetSystemTemplateRevisionsRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Gets revisions made to the specified system template. Only the previous 100
// revisions are stored. If the system has been deprecated, this action will
// return the revisions that occurred before its deprecation. This action won't
// work with systems that have been deleted.
//
//    // Example sending a request using GetSystemTemplateRevisionsRequest.
//    req := client.GetSystemTemplateRevisionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/GetSystemTemplateRevisions
func (c *Client) GetSystemTemplateRevisionsRequest(input *GetSystemTemplateRevisionsInput) GetSystemTemplateRevisionsRequest {
	op := &aws.Operation{
		Name:       opGetSystemTemplateRevisions,
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
		input = &GetSystemTemplateRevisionsInput{}
	}

	req := c.newRequest(op, input, &GetSystemTemplateRevisionsOutput{})

	return GetSystemTemplateRevisionsRequest{Request: req, Input: input, Copy: c.GetSystemTemplateRevisionsRequest}
}

// GetSystemTemplateRevisionsRequest is the request type for the
// GetSystemTemplateRevisions API operation.
type GetSystemTemplateRevisionsRequest struct {
	*aws.Request
	Input *GetSystemTemplateRevisionsInput
	Copy  func(*GetSystemTemplateRevisionsInput) GetSystemTemplateRevisionsRequest
}

// Send marshals and sends the GetSystemTemplateRevisions API request.
func (r GetSystemTemplateRevisionsRequest) Send(ctx context.Context) (*GetSystemTemplateRevisionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSystemTemplateRevisionsResponse{
		GetSystemTemplateRevisionsOutput: r.Request.Data.(*GetSystemTemplateRevisionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetSystemTemplateRevisionsRequestPaginator returns a paginator for GetSystemTemplateRevisions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetSystemTemplateRevisionsRequest(input)
//   p := iotthingsgraph.NewGetSystemTemplateRevisionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetSystemTemplateRevisionsPaginator(req GetSystemTemplateRevisionsRequest) GetSystemTemplateRevisionsPaginator {
	return GetSystemTemplateRevisionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetSystemTemplateRevisionsInput
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

// GetSystemTemplateRevisionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetSystemTemplateRevisionsPaginator struct {
	aws.Pager
}

func (p *GetSystemTemplateRevisionsPaginator) CurrentPage() *GetSystemTemplateRevisionsOutput {
	return p.Pager.CurrentPage().(*GetSystemTemplateRevisionsOutput)
}

// GetSystemTemplateRevisionsResponse is the response type for the
// GetSystemTemplateRevisions API operation.
type GetSystemTemplateRevisionsResponse struct {
	*GetSystemTemplateRevisionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSystemTemplateRevisions request.
func (r *GetSystemTemplateRevisionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
