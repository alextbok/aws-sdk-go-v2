// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListNodesInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListNodesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNodesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNodesInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNodesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about nodes in the cluster.
type ListNodesOutput struct {
	_ struct{} `type:"structure"`

	// The paginated results marker. When the result of a ListNodes operation is
	// truncated, the call returns NextToken in the response. To get another batch
	// of nodes, provide this token in your next request.
	NextToken *string `locationName:"nextToken" type:"string"`

	// List containing a NodeInfo object.
	NodeInfoList []NodeInfo `locationName:"nodeInfoList" type:"list"`
}

// String returns the string representation
func (s ListNodesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNodesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodeInfoList != nil {
		v := s.NodeInfoList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "nodeInfoList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListNodes = "ListNodes"

// ListNodesRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a list of the broker nodes in the cluster.
//
//    // Example sending a request using ListNodesRequest.
//    req := client.ListNodesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListNodes
func (c *Client) ListNodesRequest(input *ListNodesInput) ListNodesRequest {
	op := &aws.Operation{
		Name:       opListNodes,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{clusterArn}/nodes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListNodesInput{}
	}

	req := c.newRequest(op, input, &ListNodesOutput{})

	return ListNodesRequest{Request: req, Input: input, Copy: c.ListNodesRequest}
}

// ListNodesRequest is the request type for the
// ListNodes API operation.
type ListNodesRequest struct {
	*aws.Request
	Input *ListNodesInput
	Copy  func(*ListNodesInput) ListNodesRequest
}

// Send marshals and sends the ListNodes API request.
func (r ListNodesRequest) Send(ctx context.Context) (*ListNodesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNodesResponse{
		ListNodesOutput: r.Request.Data.(*ListNodesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNodesRequestPaginator returns a paginator for ListNodes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNodesRequest(input)
//   p := kafka.NewListNodesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNodesPaginator(req ListNodesRequest) ListNodesPaginator {
	return ListNodesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListNodesInput
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

// ListNodesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNodesPaginator struct {
	aws.Pager
}

func (p *ListNodesPaginator) CurrentPage() *ListNodesOutput {
	return p.Pager.CurrentPage().(*ListNodesOutput)
}

// ListNodesResponse is the response type for the
// ListNodes API operation.
type ListNodesResponse struct {
	*ListNodesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNodes request.
func (r *ListNodesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
