// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetReservationPurchaseRecommendationInput struct {
	_ struct{} `type:"structure"`

	// The account ID that is associated with the recommendation.
	AccountId *string `type:"string"`

	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the payer account and linked accounts
	// if the value is set to PAYER. If the value is LINKED, recommendations are
	// calculated for individual linked accounts only.
	AccountScope AccountScope `type:"string" enum:"true"`

	// The number of previous days that you want AWS to consider when it calculates
	// your recommendations.
	LookbackPeriodInDays LookbackPeriodInDays `type:"string" enum:"true"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextPageToken *string `type:"string"`

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize *int64 `type:"integer"`

	// The reservation purchase option that you want recommendations for.
	PaymentOption PaymentOption `type:"string" enum:"true"`

	// The specific service that you want recommendations for.
	//
	// Service is a required field
	Service *string `type:"string" required:"true"`

	// The hardware specifications for the service instances that you want recommendations
	// for, such as standard or convertible Amazon EC2 instances.
	ServiceSpecification *ServiceSpecification `type:"structure"`

	// The reservation term that you want recommendations for.
	TermInYears TermInYears `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetReservationPurchaseRecommendationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReservationPurchaseRecommendationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReservationPurchaseRecommendationInput"}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetReservationPurchaseRecommendationOutput struct {
	_ struct{} `type:"structure"`

	// Information about this specific recommendation call, such as the time stamp
	// for when Cost Explorer generated this recommendation.
	Metadata *ReservationPurchaseRecommendationMetadata `type:"structure"`

	// The pagination token for the next set of retrievable results.
	NextPageToken *string `type:"string"`

	// Recommendations for reservations to purchase.
	Recommendations []ReservationPurchaseRecommendation `type:"list"`
}

// String returns the string representation
func (s GetReservationPurchaseRecommendationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetReservationPurchaseRecommendation = "GetReservationPurchaseRecommendation"

// GetReservationPurchaseRecommendationRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Gets recommendations for which reservations to purchase. These recommendations
// could help you reduce your costs. Reservations provide a discounted hourly
// rate (up to 75%) compared to On-Demand pricing.
//
// AWS generates your recommendations by identifying your On-Demand usage during
// a specific time period and collecting your usage into categories that are
// eligible for a reservation. After AWS has these categories, it simulates
// every combination of reservations in each category of usage to identify the
// best number of each type of RI to purchase to maximize your estimated savings.
//
// For example, AWS automatically aggregates your Amazon EC2 Linux, shared tenancy,
// and c4 family usage in the US West (Oregon) Region and recommends that you
// buy size-flexible regional reservations to apply to the c4 family usage.
// AWS recommends the smallest size instance in an instance family. This makes
// it easier to purchase a size-flexible RI. AWS also shows the equal number
// of normalized units so that you can purchase any instance size that you want.
// For this example, your RI recommendation would be for c4.large because that
// is the smallest size instance in the c4 instance family.
//
//    // Example sending a request using GetReservationPurchaseRecommendationRequest.
//    req := client.GetReservationPurchaseRecommendationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetReservationPurchaseRecommendation
func (c *Client) GetReservationPurchaseRecommendationRequest(input *GetReservationPurchaseRecommendationInput) GetReservationPurchaseRecommendationRequest {
	op := &aws.Operation{
		Name:       opGetReservationPurchaseRecommendation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetReservationPurchaseRecommendationInput{}
	}

	req := c.newRequest(op, input, &GetReservationPurchaseRecommendationOutput{})

	return GetReservationPurchaseRecommendationRequest{Request: req, Input: input, Copy: c.GetReservationPurchaseRecommendationRequest}
}

// GetReservationPurchaseRecommendationRequest is the request type for the
// GetReservationPurchaseRecommendation API operation.
type GetReservationPurchaseRecommendationRequest struct {
	*aws.Request
	Input *GetReservationPurchaseRecommendationInput
	Copy  func(*GetReservationPurchaseRecommendationInput) GetReservationPurchaseRecommendationRequest
}

// Send marshals and sends the GetReservationPurchaseRecommendation API request.
func (r GetReservationPurchaseRecommendationRequest) Send(ctx context.Context) (*GetReservationPurchaseRecommendationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReservationPurchaseRecommendationResponse{
		GetReservationPurchaseRecommendationOutput: r.Request.Data.(*GetReservationPurchaseRecommendationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetReservationPurchaseRecommendationResponse is the response type for the
// GetReservationPurchaseRecommendation API operation.
type GetReservationPurchaseRecommendationResponse struct {
	*GetReservationPurchaseRecommendationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReservationPurchaseRecommendation request.
func (r *GetReservationPurchaseRecommendationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
