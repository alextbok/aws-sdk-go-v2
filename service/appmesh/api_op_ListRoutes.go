// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListRoutesInput
type ListRoutesInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// VirtualRouterName is a required field
	VirtualRouterName *string `location:"uri" locationName:"virtualRouterName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRoutesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.VirtualRouterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualRouterName"))
	}
	if s.VirtualRouterName != nil && len(*s.VirtualRouterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualRouterName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRoutesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualRouterName != nil {
		v := *s.VirtualRouterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualRouterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListRoutesOutput
type ListRoutesOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	// Routes is a required field
	Routes []RouteRef `locationName:"routes" type:"list" required:"true"`
}

// String returns the string representation
func (s ListRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRoutesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Routes != nil {
		v := s.Routes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "routes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListRoutes = "ListRoutes"

// ListRoutesRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Returns a list of existing routes in a service mesh.
//
//    // Example sending a request using ListRoutesRequest.
//    req := client.ListRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListRoutes
func (c *Client) ListRoutesRequest(input *ListRoutesInput) ListRoutesRequest {
	op := &aws.Operation{
		Name:       opListRoutes,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouter/{virtualRouterName}/routes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListRoutesInput{}
	}

	req := c.newRequest(op, input, &ListRoutesOutput{})
	return ListRoutesRequest{Request: req, Input: input, Copy: c.ListRoutesRequest}
}

// ListRoutesRequest is the request type for the
// ListRoutes API operation.
type ListRoutesRequest struct {
	*aws.Request
	Input *ListRoutesInput
	Copy  func(*ListRoutesInput) ListRoutesRequest
}

// Send marshals and sends the ListRoutes API request.
func (r ListRoutesRequest) Send(ctx context.Context) (*ListRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRoutesResponse{
		ListRoutesOutput: r.Request.Data.(*ListRoutesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRoutesRequestPaginator returns a paginator for ListRoutes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRoutesRequest(input)
//   p := appmesh.NewListRoutesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRoutesPaginator(req ListRoutesRequest) ListRoutesPaginator {
	return ListRoutesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRoutesInput
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

// ListRoutesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRoutesPaginator struct {
	aws.Pager
}

func (p *ListRoutesPaginator) CurrentPage() *ListRoutesOutput {
	return p.Pager.CurrentPage().(*ListRoutesOutput)
}

// ListRoutesResponse is the response type for the
// ListRoutes API operation.
type ListRoutesResponse struct {
	*ListRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRoutes request.
func (r *ListRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
