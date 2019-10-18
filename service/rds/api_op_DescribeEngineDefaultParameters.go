// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEngineDefaultParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB parameter group family.
	//
	// DBParameterGroupFamily is a required field
	DBParameterGroupFamily *string `type:"string" required:"true"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeEngineDefaultParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEngineDefaultParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEngineDefaultParametersInput"}

	if s.DBParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupFamily"))
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

type DescribeEngineDefaultParametersOutput struct {
	_ struct{} `type:"structure"`

	// Contains the result of a successful invocation of the DescribeEngineDefaultParameters
	// action.
	EngineDefaults *EngineDefaults `type:"structure"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEngineDefaultParameters = "DescribeEngineDefaultParameters"

// DescribeEngineDefaultParametersRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns the default engine and system parameter information for the specified
// database engine.
//
//    // Example sending a request using DescribeEngineDefaultParametersRequest.
//    req := client.DescribeEngineDefaultParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeEngineDefaultParameters
func (c *Client) DescribeEngineDefaultParametersRequest(input *DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeEngineDefaultParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"EngineDefaults.Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEngineDefaultParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeEngineDefaultParametersOutput{})
	return DescribeEngineDefaultParametersRequest{Request: req, Input: input, Copy: c.DescribeEngineDefaultParametersRequest}
}

// DescribeEngineDefaultParametersRequest is the request type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersRequest struct {
	*aws.Request
	Input *DescribeEngineDefaultParametersInput
	Copy  func(*DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest
}

// Send marshals and sends the DescribeEngineDefaultParameters API request.
func (r DescribeEngineDefaultParametersRequest) Send(ctx context.Context) (*DescribeEngineDefaultParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEngineDefaultParametersResponse{
		DescribeEngineDefaultParametersOutput: r.Request.Data.(*DescribeEngineDefaultParametersOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEngineDefaultParametersRequestPaginator returns a paginator for DescribeEngineDefaultParameters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEngineDefaultParametersRequest(input)
//   p := rds.NewDescribeEngineDefaultParametersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEngineDefaultParametersPaginator(req DescribeEngineDefaultParametersRequest) DescribeEngineDefaultParametersPaginator {
	return DescribeEngineDefaultParametersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEngineDefaultParametersInput
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

// DescribeEngineDefaultParametersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEngineDefaultParametersPaginator struct {
	aws.Pager
}

func (p *DescribeEngineDefaultParametersPaginator) CurrentPage() *DescribeEngineDefaultParametersOutput {
	return p.Pager.CurrentPage().(*DescribeEngineDefaultParametersOutput)
}

// DescribeEngineDefaultParametersResponse is the response type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersResponse struct {
	*DescribeEngineDefaultParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEngineDefaultParameters request.
func (r *DescribeEngineDefaultParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
