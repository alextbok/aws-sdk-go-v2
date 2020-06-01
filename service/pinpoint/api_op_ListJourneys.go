// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListJourneysInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s ListJourneysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJourneysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJourneysInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJourneysInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListJourneysOutput struct {
	_ struct{} `type:"structure" payload:"JourneysResponse"`

	// Provides information about the status, configuration, and other settings
	// for all the journeys that are associated with an application.
	//
	// JourneysResponse is a required field
	JourneysResponse *JourneysResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListJourneysOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJourneysOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JourneysResponse != nil {
		v := s.JourneysResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "JourneysResponse", v, metadata)
	}
	return nil
}

const opListJourneys = "ListJourneys"

// ListJourneysRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status, configuration, and other settings
// for all the journeys that are associated with an application.
//
//    // Example sending a request using ListJourneysRequest.
//    req := client.ListJourneysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/ListJourneys
func (c *Client) ListJourneysRequest(input *ListJourneysInput) ListJourneysRequest {
	op := &aws.Operation{
		Name:       opListJourneys,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/journeys",
	}

	if input == nil {
		input = &ListJourneysInput{}
	}

	req := c.newRequest(op, input, &ListJourneysOutput{})

	return ListJourneysRequest{Request: req, Input: input, Copy: c.ListJourneysRequest}
}

// ListJourneysRequest is the request type for the
// ListJourneys API operation.
type ListJourneysRequest struct {
	*aws.Request
	Input *ListJourneysInput
	Copy  func(*ListJourneysInput) ListJourneysRequest
}

// Send marshals and sends the ListJourneys API request.
func (r ListJourneysRequest) Send(ctx context.Context) (*ListJourneysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJourneysResponse{
		ListJourneysOutput: r.Request.Data.(*ListJourneysOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListJourneysResponse is the response type for the
// ListJourneys API operation.
type ListJourneysResponse struct {
	*ListJourneysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJourneys request.
func (r *ListJourneysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
