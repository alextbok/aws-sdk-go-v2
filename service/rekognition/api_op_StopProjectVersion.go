// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopProjectVersionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the model version that you want to delete.
	//
	// This operation requires permissions to perform the rekognition:StopProjectVersion
	// action.
	//
	// ProjectVersionArn is a required field
	ProjectVersionArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s StopProjectVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopProjectVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopProjectVersionInput"}

	if s.ProjectVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectVersionArn"))
	}
	if s.ProjectVersionArn != nil && len(*s.ProjectVersionArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectVersionArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopProjectVersionOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the stop operation.
	Status ProjectVersionStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s StopProjectVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopProjectVersion = "StopProjectVersion"

// StopProjectVersionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Stops a running model. The operation might take a while to complete. To check
// the current status, call DescribeProjectVersions.
//
//    // Example sending a request using StopProjectVersionRequest.
//    req := client.StopProjectVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StopProjectVersionRequest(input *StopProjectVersionInput) StopProjectVersionRequest {
	op := &aws.Operation{
		Name:       opStopProjectVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopProjectVersionInput{}
	}

	req := c.newRequest(op, input, &StopProjectVersionOutput{})

	return StopProjectVersionRequest{Request: req, Input: input, Copy: c.StopProjectVersionRequest}
}

// StopProjectVersionRequest is the request type for the
// StopProjectVersion API operation.
type StopProjectVersionRequest struct {
	*aws.Request
	Input *StopProjectVersionInput
	Copy  func(*StopProjectVersionInput) StopProjectVersionRequest
}

// Send marshals and sends the StopProjectVersion API request.
func (r StopProjectVersionRequest) Send(ctx context.Context) (*StopProjectVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopProjectVersionResponse{
		StopProjectVersionOutput: r.Request.Data.(*StopProjectVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopProjectVersionResponse is the response type for the
// StopProjectVersion API operation.
type StopProjectVersionResponse struct {
	*StopProjectVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopProjectVersion request.
func (r *StopProjectVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
