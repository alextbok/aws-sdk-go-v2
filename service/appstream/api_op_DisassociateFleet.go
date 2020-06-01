// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateFleetInput struct {
	_ struct{} `type:"structure"`

	// The name of the fleet.
	//
	// FleetName is a required field
	FleetName *string `min:"1" type:"string" required:"true"`

	// The name of the stack.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateFleetInput"}

	if s.FleetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetName"))
	}
	if s.FleetName != nil && len(*s.FleetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetName", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateFleetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateFleet = "DisassociateFleet"

// DisassociateFleetRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Disassociates the specified fleet from the specified stack.
//
//    // Example sending a request using DisassociateFleetRequest.
//    req := client.DisassociateFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DisassociateFleet
func (c *Client) DisassociateFleetRequest(input *DisassociateFleetInput) DisassociateFleetRequest {
	op := &aws.Operation{
		Name:       opDisassociateFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateFleetInput{}
	}

	req := c.newRequest(op, input, &DisassociateFleetOutput{})

	return DisassociateFleetRequest{Request: req, Input: input, Copy: c.DisassociateFleetRequest}
}

// DisassociateFleetRequest is the request type for the
// DisassociateFleet API operation.
type DisassociateFleetRequest struct {
	*aws.Request
	Input *DisassociateFleetInput
	Copy  func(*DisassociateFleetInput) DisassociateFleetRequest
}

// Send marshals and sends the DisassociateFleet API request.
func (r DisassociateFleetRequest) Send(ctx context.Context) (*DisassociateFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateFleetResponse{
		DisassociateFleetOutput: r.Request.Data.(*DisassociateFleetOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateFleetResponse is the response type for the
// DisassociateFleet API operation.
type DisassociateFleetResponse struct {
	*DisassociateFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateFleet request.
func (r *DisassociateFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
