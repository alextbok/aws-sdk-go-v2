// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRightsizingRecommendationInput struct {
	_ struct{} `type:"structure"`

	// Enables you to customize recommendations across two attributes. You can choose
	// to view recommendations for instances within the same instance families or
	// across different instance families. You can also choose to view your estimated
	// savings associated with recommendations with consideration of existing Savings
	// Plans or RI benefits, or niether.
	Configuration *RightsizingRecommendationConfiguration `type:"structure"`

	// Use Expression to filter by cost or by usage. There are two patterns:
	//
	//    * Simple dimension values - You can set the dimension name and values
	//    for the filters that you plan to use. For example, you can filter for
	//    REGION==us-east-1 OR REGION==us-west-1. The Expression for that looks
	//    like this: { "Dimensions": { "Key": "REGION", "Values": [ "us-east-1",
	//    “us-west-1” ] } } The list of dimension values are OR'd together to
	//    retrieve cost or usage data. You can create Expression and DimensionValues
	//    objects using either with* methods or set* methods in multiple lines.
	//
	//    * Compound dimension values with logical operations - You can use multiple
	//    Expression types and the logical operators AND/OR/NOT to create a list
	//    of one or more Expression objects. This allows you to filter on more advanced
	//    options. For example, you can filter on ((REGION == us-east-1 OR REGION
	//    == us-west-1) OR (TAG.Type == Type1)) AND (USAGE_TYPE != DataTransfer).
	//    The Expression for that looks like this: { "And": [ {"Or": [ {"Dimensions":
	//    { "Key": "REGION", "Values": [ "us-east-1", "us-west-1" ] }}, {"Tags":
	//    { "Key": "TagName", "Values": ["Value1"] } } ]}, {"Not": {"Dimensions":
	//    { "Key": "USAGE_TYPE", "Values": ["DataTransfer"] }}} ] } Because each
	//    Expression can have only one operator, the service returns an error if
	//    more than one is specified. The following example shows an Expression
	//    object that creates an error. { "And": [ ... ], "DimensionValues": { "Dimension":
	//    "USAGE_TYPE", "Values": [ "DataTransfer" ] } }
	//
	// For GetRightsizingRecommendation action, a combination of OR and NOT is not
	// supported. OR is not supported between different dimensions, or dimensions
	// and tags. NOT operators aren't supported. Dimensions are also limited to
	// LINKED_ACCOUNT, REGION, or RIGHTSIZING_TYPE.
	Filter *Expression `type:"structure"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextPageToken *string `type:"string"`

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize *int64 `type:"integer"`

	// The specific service that you want recommendations for. The only valid value
	// for GetRightsizingRecommendation is "AmazonEC2".
	//
	// Service is a required field
	Service *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetRightsizingRecommendationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRightsizingRecommendationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRightsizingRecommendationInput"}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			invalidParams.AddNested("Configuration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRightsizingRecommendationOutput struct {
	_ struct{} `type:"structure"`

	// Enables you to customize recommendations across two attributes. You can choose
	// to view recommendations for instances within the same instance families or
	// across different instance families. You can also choose to view your estimated
	// savings associated with recommendations with consideration of existing Savings
	// Plans or RI benefits, or niether.
	Configuration *RightsizingRecommendationConfiguration `type:"structure"`

	// Information regarding this specific recommendation set.
	Metadata *RightsizingRecommendationMetadata `type:"structure"`

	// The token to retrieve the next set of results.
	NextPageToken *string `type:"string"`

	// Recommendations to rightsize resources.
	RightsizingRecommendations []RightsizingRecommendation `type:"list"`

	// Summary of this recommendation set.
	Summary *RightsizingRecommendationSummary `type:"structure"`
}

// String returns the string representation
func (s GetRightsizingRecommendationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRightsizingRecommendation = "GetRightsizingRecommendation"

// GetRightsizingRecommendationRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Creates recommendations that helps you save cost by identifying idle and
// underutilized Amazon EC2 instances.
//
// Recommendations are generated to either downsize or terminate instances,
// along with providing savings detail and metrics. For details on calculation
// and function, see Optimizing Your Cost with Rightsizing Recommendations (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-what-is.html).
//
//    // Example sending a request using GetRightsizingRecommendationRequest.
//    req := client.GetRightsizingRecommendationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetRightsizingRecommendation
func (c *Client) GetRightsizingRecommendationRequest(input *GetRightsizingRecommendationInput) GetRightsizingRecommendationRequest {
	op := &aws.Operation{
		Name:       opGetRightsizingRecommendation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRightsizingRecommendationInput{}
	}

	req := c.newRequest(op, input, &GetRightsizingRecommendationOutput{})

	return GetRightsizingRecommendationRequest{Request: req, Input: input, Copy: c.GetRightsizingRecommendationRequest}
}

// GetRightsizingRecommendationRequest is the request type for the
// GetRightsizingRecommendation API operation.
type GetRightsizingRecommendationRequest struct {
	*aws.Request
	Input *GetRightsizingRecommendationInput
	Copy  func(*GetRightsizingRecommendationInput) GetRightsizingRecommendationRequest
}

// Send marshals and sends the GetRightsizingRecommendation API request.
func (r GetRightsizingRecommendationRequest) Send(ctx context.Context) (*GetRightsizingRecommendationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRightsizingRecommendationResponse{
		GetRightsizingRecommendationOutput: r.Request.Data.(*GetRightsizingRecommendationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRightsizingRecommendationResponse is the response type for the
// GetRightsizingRecommendation API operation.
type GetRightsizingRecommendationResponse struct {
	*GetRightsizingRecommendationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRightsizingRecommendation request.
func (r *GetRightsizingRecommendationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
