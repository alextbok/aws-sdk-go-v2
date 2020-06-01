// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to retrieve the results of a predictive inbox placement test.
type GetDeliverabilityTestReportInput struct {
	_ struct{} `type:"structure"`

	// A unique string that identifies the predictive inbox placement test.
	//
	// ReportId is a required field
	ReportId *string `location:"uri" locationName:"ReportId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeliverabilityTestReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeliverabilityTestReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeliverabilityTestReportInput"}

	if s.ReportId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReportId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeliverabilityTestReportInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ReportId != nil {
		v := *s.ReportId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ReportId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The results of the predictive inbox placement test.
type GetDeliverabilityTestReportOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the results of the predictive inbox placement test.
	//
	// DeliverabilityTestReport is a required field
	DeliverabilityTestReport *DeliverabilityTestReport `type:"structure" required:"true"`

	// An object that describes how the test email was handled by several email
	// providers, including Gmail, Hotmail, Yahoo, AOL, and others.
	//
	// IspPlacements is a required field
	IspPlacements []IspPlacement `type:"list" required:"true"`

	// An object that contains the message that you sent when you performed this
	// predictive inbox placement test.
	Message *string `type:"string"`

	// An object that specifies how many test messages that were sent during the
	// predictive inbox placement test were delivered to recipients' inboxes, how
	// many were sent to recipients' spam folders, and how many weren't delivered.
	//
	// OverallPlacement is a required field
	OverallPlacement *PlacementStatistics `type:"structure" required:"true"`

	// An array of objects that define the tags (keys and values) that are associated
	// with the predictive inbox placement test.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s GetDeliverabilityTestReportOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeliverabilityTestReportOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeliverabilityTestReport != nil {
		v := s.DeliverabilityTestReport

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DeliverabilityTestReport", v, metadata)
	}
	if s.IspPlacements != nil {
		v := s.IspPlacements

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "IspPlacements", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OverallPlacement != nil {
		v := s.OverallPlacement

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "OverallPlacement", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetDeliverabilityTestReport = "GetDeliverabilityTestReport"

// GetDeliverabilityTestReportRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Retrieve the results of a predictive inbox placement test.
//
//    // Example sending a request using GetDeliverabilityTestReportRequest.
//    req := client.GetDeliverabilityTestReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/GetDeliverabilityTestReport
func (c *Client) GetDeliverabilityTestReportRequest(input *GetDeliverabilityTestReportInput) GetDeliverabilityTestReportRequest {
	op := &aws.Operation{
		Name:       opGetDeliverabilityTestReport,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/email/deliverability-dashboard/test-reports/{ReportId}",
	}

	if input == nil {
		input = &GetDeliverabilityTestReportInput{}
	}

	req := c.newRequest(op, input, &GetDeliverabilityTestReportOutput{})

	return GetDeliverabilityTestReportRequest{Request: req, Input: input, Copy: c.GetDeliverabilityTestReportRequest}
}

// GetDeliverabilityTestReportRequest is the request type for the
// GetDeliverabilityTestReport API operation.
type GetDeliverabilityTestReportRequest struct {
	*aws.Request
	Input *GetDeliverabilityTestReportInput
	Copy  func(*GetDeliverabilityTestReportInput) GetDeliverabilityTestReportRequest
}

// Send marshals and sends the GetDeliverabilityTestReport API request.
func (r GetDeliverabilityTestReportRequest) Send(ctx context.Context) (*GetDeliverabilityTestReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeliverabilityTestReportResponse{
		GetDeliverabilityTestReportOutput: r.Request.Data.(*GetDeliverabilityTestReportOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeliverabilityTestReportResponse is the response type for the
// GetDeliverabilityTestReport API operation.
type GetDeliverabilityTestReportResponse struct {
	*GetDeliverabilityTestReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeliverabilityTestReport request.
func (r *GetDeliverabilityTestReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
