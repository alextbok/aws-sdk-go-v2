// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListAssetsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the asset model by which to filter the list of assets. This parameter
	// is required if you choose ALL for filter.
	AssetModelId *string `location:"querystring" locationName:"assetModelId" min:"36" type:"string"`

	// The filter for the requested list of assets. Choose one of the following
	// options. Defaults to ALL.
	//
	//    * ALL – The list includes all assets for a given asset model ID. The
	//    assetModelId parameter is required if you filter by ALL.
	//
	//    * TOP_LEVEL – The list includes only top-level assets in the asset hierarchy
	//    tree.
	Filter ListAssetsFilter `location:"querystring" locationName:"filter" type:"string" enum:"true"`

	// The maximum number of results to be returned per paginated request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to be used for the next set of paginated results.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListAssetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAssetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAssetsInput"}
	if s.AssetModelId != nil && len(*s.AssetModelId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AssetModelId", 36))
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
func (s ListAssetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssetModelId != nil {
		v := *s.AssetModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "assetModelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Filter) > 0 {
		v := s.Filter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "filter", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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

type ListAssetsOutput struct {
	_ struct{} `type:"structure"`

	// A list that summarizes each asset.
	//
	// AssetSummaries is a required field
	AssetSummaries []AssetSummary `locationName:"assetSummaries" type:"list" required:"true"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListAssetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAssetsOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opListAssets = "ListAssets"

// ListAssetsRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Retrieves a paginated list of asset summaries.
//
// You can use this operation to do the following:
//
//    * List assets based on a specific asset model.
//
//    * List top-level assets.
//
// You can't use this operation to list all assets. To retrieve summaries for
// all of your assets, use ListAssetModels (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_ListAssetModels.html)
// to get all of your asset model IDs. Then, use ListAssets to get all assets
// for each asset model.
//
//    // Example sending a request using ListAssetsRequest.
//    req := client.ListAssetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/ListAssets
func (c *Client) ListAssetsRequest(input *ListAssetsInput) ListAssetsRequest {
	op := &aws.Operation{
		Name:       opListAssets,
		HTTPMethod: "GET",
		HTTPPath:   "/assets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAssetsInput{}
	}

	req := c.newRequest(op, input, &ListAssetsOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("model.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return ListAssetsRequest{Request: req, Input: input, Copy: c.ListAssetsRequest}
}

// ListAssetsRequest is the request type for the
// ListAssets API operation.
type ListAssetsRequest struct {
	*aws.Request
	Input *ListAssetsInput
	Copy  func(*ListAssetsInput) ListAssetsRequest
}

// Send marshals and sends the ListAssets API request.
func (r ListAssetsRequest) Send(ctx context.Context) (*ListAssetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssetsResponse{
		ListAssetsOutput: r.Request.Data.(*ListAssetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAssetsRequestPaginator returns a paginator for ListAssets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAssetsRequest(input)
//   p := iotsitewise.NewListAssetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAssetsPaginator(req ListAssetsRequest) ListAssetsPaginator {
	return ListAssetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAssetsInput
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

// ListAssetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAssetsPaginator struct {
	aws.Pager
}

func (p *ListAssetsPaginator) CurrentPage() *ListAssetsOutput {
	return p.Pager.CurrentPage().(*ListAssetsOutput)
}

// ListAssetsResponse is the response type for the
// ListAssets API operation.
type ListAssetsResponse struct {
	*ListAssetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssets request.
func (r *ListAssetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
