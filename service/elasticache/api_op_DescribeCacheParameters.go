// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeCacheParameters operation.
type DescribeCacheParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific cache parameter group to return details for.
	//
	// CacheParameterGroupName is a required field
	CacheParameterGroupName *string `type:"string" required:"true"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The parameter types to return.
	//
	// Valid values: user | system | engine-default
	Source *string `type:"string"`
}

// String returns the string representation
func (s DescribeCacheParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCacheParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCacheParametersInput"}

	if s.CacheParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DescribeCacheParameters operation.
type DescribeCacheParametersOutput struct {
	_ struct{} `type:"structure"`

	// A list of parameters specific to a particular cache node type. Each element
	// in the list contains detailed information about one parameter.
	CacheNodeTypeSpecificParameters []CacheNodeTypeSpecificParameter `locationNameList:"CacheNodeTypeSpecificParameter" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`

	// A list of Parameter instances.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`
}

// String returns the string representation
func (s DescribeCacheParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCacheParameters = "DescribeCacheParameters"

// DescribeCacheParametersRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns the detailed parameter list for a particular cache parameter group.
//
//    // Example sending a request using DescribeCacheParametersRequest.
//    req := client.DescribeCacheParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeCacheParameters
func (c *Client) DescribeCacheParametersRequest(input *DescribeCacheParametersInput) DescribeCacheParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeCacheParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeCacheParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeCacheParametersOutput{})

	return DescribeCacheParametersRequest{Request: req, Input: input, Copy: c.DescribeCacheParametersRequest}
}

// DescribeCacheParametersRequest is the request type for the
// DescribeCacheParameters API operation.
type DescribeCacheParametersRequest struct {
	*aws.Request
	Input *DescribeCacheParametersInput
	Copy  func(*DescribeCacheParametersInput) DescribeCacheParametersRequest
}

// Send marshals and sends the DescribeCacheParameters API request.
func (r DescribeCacheParametersRequest) Send(ctx context.Context) (*DescribeCacheParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCacheParametersResponse{
		DescribeCacheParametersOutput: r.Request.Data.(*DescribeCacheParametersOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCacheParametersRequestPaginator returns a paginator for DescribeCacheParameters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCacheParametersRequest(input)
//   p := elasticache.NewDescribeCacheParametersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCacheParametersPaginator(req DescribeCacheParametersRequest) DescribeCacheParametersPaginator {
	return DescribeCacheParametersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCacheParametersInput
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

// DescribeCacheParametersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCacheParametersPaginator struct {
	aws.Pager
}

func (p *DescribeCacheParametersPaginator) CurrentPage() *DescribeCacheParametersOutput {
	return p.Pager.CurrentPage().(*DescribeCacheParametersOutput)
}

// DescribeCacheParametersResponse is the response type for the
// DescribeCacheParameters API operation.
type DescribeCacheParametersResponse struct {
	*DescribeCacheParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCacheParameters request.
func (r *DescribeCacheParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
