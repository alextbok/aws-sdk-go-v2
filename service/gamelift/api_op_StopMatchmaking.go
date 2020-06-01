// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type StopMatchmakingInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a matchmaking ticket.
	//
	// TicketId is a required field
	TicketId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopMatchmakingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopMatchmakingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopMatchmakingInput"}

	if s.TicketId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TicketId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopMatchmakingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopMatchmakingOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopMatchmaking = "StopMatchmaking"

// StopMatchmakingRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Cancels a matchmaking ticket or match backfill ticket that is currently being
// processed. To stop the matchmaking operation, specify the ticket ID. If successful,
// work on the ticket is stopped, and the ticket status is changed to CANCELLED.
//
// This call is also used to turn off automatic backfill for an individual game
// session. This is for game sessions that are created with a matchmaking configuration
// that has automatic backfill enabled. The ticket ID is included in the MatchmakerData
// of an updated game session object, which is provided to the game server.
//
// If the action is successful, the service sends back an empty JSON struct
// with the HTTP 200 response (not an empty HTTP body).
//
// Learn more
//
//  Add FlexMatch to a Game Client (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-client.html)
//
// Related operations
//
//    * StartMatchmaking
//
//    * DescribeMatchmaking
//
//    * StopMatchmaking
//
//    * AcceptMatch
//
//    * StartMatchBackfill
//
//    // Example sending a request using StopMatchmakingRequest.
//    req := client.StopMatchmakingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/StopMatchmaking
func (c *Client) StopMatchmakingRequest(input *StopMatchmakingInput) StopMatchmakingRequest {
	op := &aws.Operation{
		Name:       opStopMatchmaking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopMatchmakingInput{}
	}

	req := c.newRequest(op, input, &StopMatchmakingOutput{})

	return StopMatchmakingRequest{Request: req, Input: input, Copy: c.StopMatchmakingRequest}
}

// StopMatchmakingRequest is the request type for the
// StopMatchmaking API operation.
type StopMatchmakingRequest struct {
	*aws.Request
	Input *StopMatchmakingInput
	Copy  func(*StopMatchmakingInput) StopMatchmakingRequest
}

// Send marshals and sends the StopMatchmaking API request.
func (r StopMatchmakingRequest) Send(ctx context.Context) (*StopMatchmakingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopMatchmakingResponse{
		StopMatchmakingOutput: r.Request.Data.(*StopMatchmakingOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopMatchmakingResponse is the response type for the
// StopMatchmaking API operation.
type StopMatchmakingResponse struct {
	*StopMatchmakingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopMatchmaking request.
func (r *StopMatchmakingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
