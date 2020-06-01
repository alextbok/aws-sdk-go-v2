// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListAssociatedAssetsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the parent asset.
	//
	// AssetId is a required field
	AssetId *string `location:"uri" locationName:"assetId" min:"36" type:"string" required:"true"`

	// The hierarchy ID (of the parent asset model) whose associated assets are
	// returned. To find a hierarchy ID, use the DescribeAsset (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAsset.html)
	// or DescribeAssetModel (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAssetModel.html)
	// actions.
	//
	// For more information, see Asset Hierarchies (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// HierarchyId is a required field
	HierarchyId *string `location:"querystring" locationName:"hierarchyId" min:"36" type:"string" required:"true"`

	// The maximum number of results to be returned per paginated request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to be used for the next set of paginated results.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListAssociatedAssetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAssociatedAssetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAssociatedAssetsInput"}

	if s.AssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetId"))
	}
	if s.AssetId != nil && len(*s.AssetId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AssetId", 36))
	}

	if s.HierarchyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HierarchyId"))
	}
	if s.HierarchyId != nil && len(*s.HierarchyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("HierarchyId", 36))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAssociatedAssetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssetId != nil {
		v := *s.AssetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "assetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HierarchyId != nil {
		v := *s.HierarchyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "hierarchyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

type ListAssociatedAssetsOutput struct {
	_ struct{} `type:"structure"`

	// A list that summarizes the associated assets.
	//
	// AssetSummaries is a required field
	AssetSummaries []AssociatedAssetsSummary `locationName:"assetSummaries" type:"list" required:"true"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListAssociatedAssetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAssociatedAssetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssetSummaries != nil {
		v := s.AssetSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "assetSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListAssociatedAssets = "ListAssociatedAssets"

// ListAssociatedAssetsRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Retrieves a paginated list of the assets associated to a parent asset (assetId)
// by a given hierarchy (hierarchyId).
//
//    // Example sending a request using ListAssociatedAssetsRequest.
//    req := client.ListAssociatedAssetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/ListAssociatedAssets
func (c *Client) ListAssociatedAssetsRequest(input *ListAssociatedAssetsInput) ListAssociatedAssetsRequest {
	op := &aws.Operation{
		Name:       opListAssociatedAssets,
		HTTPMethod: "GET",
		HTTPPath:   "/assets/{assetId}/hierarchies",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAssociatedAssetsInput{}
	}

	req := c.newRequest(op, input, &ListAssociatedAssetsOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("model.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return ListAssociatedAssetsRequest{Request: req, Input: input, Copy: c.ListAssociatedAssetsRequest}
}

// ListAssociatedAssetsRequest is the request type for the
// ListAssociatedAssets API operation.
type ListAssociatedAssetsRequest struct {
	*aws.Request
	Input *ListAssociatedAssetsInput
	Copy  func(*ListAssociatedAssetsInput) ListAssociatedAssetsRequest
}

// Send marshals and sends the ListAssociatedAssets API request.
func (r ListAssociatedAssetsRequest) Send(ctx context.Context) (*ListAssociatedAssetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssociatedAssetsResponse{
		ListAssociatedAssetsOutput: r.Request.Data.(*ListAssociatedAssetsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAssociatedAssetsRequestPaginator returns a paginator for ListAssociatedAssets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAssociatedAssetsRequest(input)
//   p := iotsitewise.NewListAssociatedAssetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAssociatedAssetsPaginator(req ListAssociatedAssetsRequest) ListAssociatedAssetsPaginator {
	return ListAssociatedAssetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAssociatedAssetsInput
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

// ListAssociatedAssetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAssociatedAssetsPaginator struct {
	aws.Pager
}

func (p *ListAssociatedAssetsPaginator) CurrentPage() *ListAssociatedAssetsOutput {
	return p.Pager.CurrentPage().(*ListAssociatedAssetsOutput)
}

// ListAssociatedAssetsResponse is the response type for the
// ListAssociatedAssets API operation.
type ListAssociatedAssetsResponse struct {
	*ListAssociatedAssetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssociatedAssets request.
func (r *ListAssociatedAssetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
