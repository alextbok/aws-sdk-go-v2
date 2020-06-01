// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetFindingsInput struct {
	_ struct{} `type:"structure"`

	// The finding attributes used to define a condition to filter the returned
	// findings.
	Filters *AwsSecurityFindingFilters `type:"structure"`

	// The maximum number of findings to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token that is required for pagination. On your first call to the GetFindings
	// operation, set the value of this parameter to NULL.
	//
	// For subsequent calls to the operation, to continue listing data, set the
	// value of this parameter to the value returned from the previous response.
	NextToken *string `type:"string"`

	// The finding attributes used to sort the list of returned findings.
	SortCriteria []SortCriterion `type:"list"`
}

// String returns the string representation
func (s GetFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFindingsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Filters", v, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SortCriteria != nil {
		v := s.SortCriteria

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SortCriteria", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type GetFindingsOutput struct {
	_ struct{} `type:"structure"`

	// The findings that matched the filters specified in the request.
	//
	// Findings is a required field
	Findings []AwsSecurityFinding `type:"list" required:"true"`

	// The pagination token to use to request the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Findings != nil {
		v := s.Findings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Findings", metadata)
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

const opGetFindings = "GetFindings"

// GetFindingsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Returns a list of findings that match the specified criteria.
//
//    // Example sending a request using GetFindingsRequest.
//    req := client.GetFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetFindings
func (c *Client) GetFindingsRequest(input *GetFindingsInput) GetFindingsRequest {
	op := &aws.Operation{
		Name:       opGetFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/findings",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetFindingsInput{}
	}

	req := c.newRequest(op, input, &GetFindingsOutput{})

	return GetFindingsRequest{Request: req, Input: input, Copy: c.GetFindingsRequest}
}

// GetFindingsRequest is the request type for the
// GetFindings API operation.
type GetFindingsRequest struct {
	*aws.Request
	Input *GetFindingsInput
	Copy  func(*GetFindingsInput) GetFindingsRequest
}

// Send marshals and sends the GetFindings API request.
func (r GetFindingsRequest) Send(ctx context.Context) (*GetFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFindingsResponse{
		GetFindingsOutput: r.Request.Data.(*GetFindingsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetFindingsRequestPaginator returns a paginator for GetFindings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetFindingsRequest(input)
//   p := securityhub.NewGetFindingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetFindingsPaginator(req GetFindingsRequest) GetFindingsPaginator {
	return GetFindingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetFindingsInput
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

// GetFindingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetFindingsPaginator struct {
	aws.Pager
}

func (p *GetFindingsPaginator) CurrentPage() *GetFindingsOutput {
	return p.Pager.CurrentPage().(*GetFindingsOutput)
}

// GetFindingsResponse is the response type for the
// GetFindings API operation.
type GetFindingsResponse struct {
	*GetFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFindings request.
func (r *GetFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
