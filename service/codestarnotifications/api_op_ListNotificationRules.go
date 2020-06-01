// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListNotificationRulesInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to return information by service or resource type. For
	// valid values, see ListNotificationRulesFilter.
	//
	// A filter with the same name can appear more than once when used with OR statements.
	// Filters with different names should be applied with AND statements.
	Filters []ListNotificationRulesFilter `type:"list"`

	// A non-negative integer used to limit the number of returned results. The
	// maximum number of results that can be returned is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListNotificationRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNotificationRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNotificationRulesInput"}
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNotificationRulesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

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
	return nil
}

type ListNotificationRulesOutput struct {
	_ struct{} `type:"structure"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `type:"string"`

	// The list of notification rules for the AWS account, by Amazon Resource Name
	// (ARN) and ID.
	NotificationRules []NotificationRuleSummary `type:"list"`
}

// String returns the string representation
func (s ListNotificationRulesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNotificationRulesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NotificationRules != nil {
		v := s.NotificationRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "NotificationRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListNotificationRules = "ListNotificationRules"

// ListNotificationRulesRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Returns a list of the notification rules for an AWS account.
//
//    // Example sending a request using ListNotificationRulesRequest.
//    req := client.ListNotificationRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/ListNotificationRules
func (c *Client) ListNotificationRulesRequest(input *ListNotificationRulesInput) ListNotificationRulesRequest {
	op := &aws.Operation{
		Name:       opListNotificationRules,
		HTTPMethod: "POST",
		HTTPPath:   "/listNotificationRules",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListNotificationRulesInput{}
	}

	req := c.newRequest(op, input, &ListNotificationRulesOutput{})

	return ListNotificationRulesRequest{Request: req, Input: input, Copy: c.ListNotificationRulesRequest}
}

// ListNotificationRulesRequest is the request type for the
// ListNotificationRules API operation.
type ListNotificationRulesRequest struct {
	*aws.Request
	Input *ListNotificationRulesInput
	Copy  func(*ListNotificationRulesInput) ListNotificationRulesRequest
}

// Send marshals and sends the ListNotificationRules API request.
func (r ListNotificationRulesRequest) Send(ctx context.Context) (*ListNotificationRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNotificationRulesResponse{
		ListNotificationRulesOutput: r.Request.Data.(*ListNotificationRulesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNotificationRulesRequestPaginator returns a paginator for ListNotificationRules.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNotificationRulesRequest(input)
//   p := codestarnotifications.NewListNotificationRulesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNotificationRulesPaginator(req ListNotificationRulesRequest) ListNotificationRulesPaginator {
	return ListNotificationRulesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListNotificationRulesInput
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

// ListNotificationRulesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNotificationRulesPaginator struct {
	aws.Pager
}

func (p *ListNotificationRulesPaginator) CurrentPage() *ListNotificationRulesOutput {
	return p.Pager.CurrentPage().(*ListNotificationRulesOutput)
}

// ListNotificationRulesResponse is the response type for the
// ListNotificationRules API operation.
type ListNotificationRulesResponse struct {
	*ListNotificationRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNotificationRules request.
func (r *ListNotificationRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
