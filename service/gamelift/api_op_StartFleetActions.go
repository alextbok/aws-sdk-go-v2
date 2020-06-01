// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartFleetActionsInput struct {
	_ struct{} `type:"structure"`

	// List of actions to restart on the fleet.
	//
	// Actions is a required field
	Actions []FleetAction `min:"1" type:"list" required:"true"`

	// A unique identifier for a fleet to start actions on. You can use either the
	// fleet ID or ARN value.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartFleetActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartFleetActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartFleetActionsInput"}

	if s.Actions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Actions"))
	}
	if s.Actions != nil && len(s.Actions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Actions", 1))
	}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartFleetActionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartFleetActionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartFleetActions = "StartFleetActions"

// StartFleetActionsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Resumes activity on a fleet that was suspended with StopFleetActions. Currently,
// this operation is used to restart a fleet's auto-scaling activity.
//
// To start fleet actions, specify the fleet ID and the type of actions to restart.
// When auto-scaling fleet actions are restarted, Amazon GameLift once again
// initiates scaling events as triggered by the fleet's scaling policies. If
// actions on the fleet were never stopped, this operation will have no effect.
// You can view a fleet's stopped actions using DescribeFleetAttributes.
//
// Learn more
//
// Setting up GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * DescribeFleetAttributes
//
//    * UpdateFleetAttributes
//
//    * StartFleetActions or StopFleetActions
//
//    // Example sending a request using StartFleetActionsRequest.
//    req := client.StartFleetActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/StartFleetActions
func (c *Client) StartFleetActionsRequest(input *StartFleetActionsInput) StartFleetActionsRequest {
	op := &aws.Operation{
		Name:       opStartFleetActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartFleetActionsInput{}
	}

	req := c.newRequest(op, input, &StartFleetActionsOutput{})

	return StartFleetActionsRequest{Request: req, Input: input, Copy: c.StartFleetActionsRequest}
}

// StartFleetActionsRequest is the request type for the
// StartFleetActions API operation.
type StartFleetActionsRequest struct {
	*aws.Request
	Input *StartFleetActionsInput
	Copy  func(*StartFleetActionsInput) StartFleetActionsRequest
}

// Send marshals and sends the StartFleetActions API request.
func (r StartFleetActionsRequest) Send(ctx context.Context) (*StartFleetActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartFleetActionsResponse{
		StartFleetActionsOutput: r.Request.Data.(*StartFleetActionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartFleetActionsResponse is the response type for the
// StartFleetActions API operation.
type StartFleetActionsResponse struct {
	*StartFleetActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartFleetActions request.
func (r *StartFleetActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
