// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutDashboardInput struct {
	_ struct{} `type:"structure"`

	// The detailed information about the dashboard in JSON format, including the
	// widgets to include and their location on the dashboard. This parameter is
	// required.
	//
	// For more information about the syntax, see Dashboard Body Structure and Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	//
	// DashboardBody is a required field
	DashboardBody *string `type:"string" required:"true"`

	// The name of the dashboard. If a dashboard with this name already exists,
	// this call modifies that dashboard, replacing its current contents. Otherwise,
	// a new dashboard is created. The maximum length is 255, and valid characters
	// are A-Z, a-z, 0-9, "-", and "_". This parameter is required.
	//
	// DashboardName is a required field
	DashboardName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutDashboardInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDashboardInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDashboardInput"}

	if s.DashboardBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardBody"))
	}

	if s.DashboardName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutDashboardOutput struct {
	_ struct{} `type:"structure"`

	// If the input for PutDashboard was correct and the dashboard was successfully
	// created or modified, this result is empty.
	//
	// If this result includes only warning messages, then the input was valid enough
	// for the dashboard to be created or modified, but some elements of the dashboard
	// may not render.
	//
	// If this result includes error messages, the input was not valid and the operation
	// failed.
	DashboardValidationMessages []DashboardValidationMessage `type:"list"`
}

// String returns the string representation
func (s PutDashboardOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutDashboard = "PutDashboard"

// PutDashboardRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Creates a dashboard if it does not already exist, or updates an existing
// dashboard. If you update a dashboard, the entire contents are replaced with
// what you specify here.
//
// All dashboards in your account are global, not region-specific.
//
// A simple way to create a dashboard using PutDashboard is to copy an existing
// dashboard. To copy an existing dashboard using the console, you can load
// the dashboard and then use the View/edit source command in the Actions menu
// to display the JSON block for that dashboard. Another way to copy a dashboard
// is to use GetDashboard, and then use the data returned within DashboardBody
// as the template for the new dashboard when you call PutDashboard.
//
// When you create a dashboard with PutDashboard, a good practice is to add
// a text widget at the top of the dashboard with a message that the dashboard
// was created by script and should not be changed in the console. This message
// could also point console users to the location of the DashboardBody script
// or the CloudFormation template used to create the dashboard.
//
//    // Example sending a request using PutDashboardRequest.
//    req := client.PutDashboardRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutDashboard
func (c *Client) PutDashboardRequest(input *PutDashboardInput) PutDashboardRequest {
	op := &aws.Operation{
		Name:       opPutDashboard,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutDashboardInput{}
	}

	req := c.newRequest(op, input, &PutDashboardOutput{})

	return PutDashboardRequest{Request: req, Input: input, Copy: c.PutDashboardRequest}
}

// PutDashboardRequest is the request type for the
// PutDashboard API operation.
type PutDashboardRequest struct {
	*aws.Request
	Input *PutDashboardInput
	Copy  func(*PutDashboardInput) PutDashboardRequest
}

// Send marshals and sends the PutDashboard API request.
func (r PutDashboardRequest) Send(ctx context.Context) (*PutDashboardResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDashboardResponse{
		PutDashboardOutput: r.Request.Data.(*PutDashboardOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDashboardResponse is the response type for the
// PutDashboard API operation.
type PutDashboardResponse struct {
	*PutDashboardOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDashboard request.
func (r *PutDashboardResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
