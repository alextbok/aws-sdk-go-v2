// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteDashboardInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the dashboard that you're deleting.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" min:"1" type:"string" required:"true"`

	// The version number of the dashboard. If the version number property is provided,
	// only the specified version of the dashboard is deleted.
	VersionNumber *int64 `location:"querystring" locationName:"version-number" min:"1" type:"long"`
}

// String returns the string representation
func (s DeleteDashboardInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDashboardInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDashboardInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if s.DashboardId != nil && len(*s.DashboardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardId", 1))
	}
	if s.VersionNumber != nil && *s.VersionNumber < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("VersionNumber", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDashboardInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "version-number", protocol.Int64Value(v), metadata)
	}
	return nil
}

type DeleteDashboardOutput struct {
	_ struct{} `type:"structure"`

	// The Secure Socket Layer (SSL) properties that apply for the resource.
	Arn *string `type:"string"`

	// The ID of the dashboard.
	DashboardId *string `min:"1" type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DeleteDashboardOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDashboardOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDeleteDashboard = "DeleteDashboard"

// DeleteDashboardRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Deletes a dashboard.
//
//    // Example sending a request using DeleteDashboardRequest.
//    req := client.DeleteDashboardRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteDashboard
func (c *Client) DeleteDashboardRequest(input *DeleteDashboardInput) DeleteDashboardRequest {
	op := &aws.Operation{
		Name:       opDeleteDashboard,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}",
	}

	if input == nil {
		input = &DeleteDashboardInput{}
	}

	req := c.newRequest(op, input, &DeleteDashboardOutput{})

	return DeleteDashboardRequest{Request: req, Input: input, Copy: c.DeleteDashboardRequest}
}

// DeleteDashboardRequest is the request type for the
// DeleteDashboard API operation.
type DeleteDashboardRequest struct {
	*aws.Request
	Input *DeleteDashboardInput
	Copy  func(*DeleteDashboardInput) DeleteDashboardRequest
}

// Send marshals and sends the DeleteDashboard API request.
func (r DeleteDashboardRequest) Send(ctx context.Context) (*DeleteDashboardResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDashboardResponse{
		DeleteDashboardOutput: r.Request.Data.(*DeleteDashboardOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDashboardResponse is the response type for the
// DeleteDashboard API operation.
type DeleteDashboardResponse struct {
	*DeleteDashboardOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDashboard request.
func (r *DeleteDashboardResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
