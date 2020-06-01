// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetWorkGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the workgroup.
	//
	// WorkGroup is a required field
	WorkGroup *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetWorkGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWorkGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWorkGroupInput"}

	if s.WorkGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkGroup"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetWorkGroupOutput struct {
	_ struct{} `type:"structure"`

	// Information about the workgroup.
	WorkGroup *WorkGroup `type:"structure"`
}

// String returns the string representation
func (s GetWorkGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetWorkGroup = "GetWorkGroup"

// GetWorkGroupRequest returns a request value for making API operation for
// Amazon Athena.
//
// Returns information about the workgroup with the specified name.
//
//    // Example sending a request using GetWorkGroupRequest.
//    req := client.GetWorkGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/GetWorkGroup
func (c *Client) GetWorkGroupRequest(input *GetWorkGroupInput) GetWorkGroupRequest {
	op := &aws.Operation{
		Name:       opGetWorkGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetWorkGroupInput{}
	}

	req := c.newRequest(op, input, &GetWorkGroupOutput{})

	return GetWorkGroupRequest{Request: req, Input: input, Copy: c.GetWorkGroupRequest}
}

// GetWorkGroupRequest is the request type for the
// GetWorkGroup API operation.
type GetWorkGroupRequest struct {
	*aws.Request
	Input *GetWorkGroupInput
	Copy  func(*GetWorkGroupInput) GetWorkGroupRequest
}

// Send marshals and sends the GetWorkGroup API request.
func (r GetWorkGroupRequest) Send(ctx context.Context) (*GetWorkGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWorkGroupResponse{
		GetWorkGroupOutput: r.Request.Data.(*GetWorkGroupOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetWorkGroupResponse is the response type for the
// GetWorkGroup API operation.
type GetWorkGroupResponse struct {
	*GetWorkGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWorkGroup request.
func (r *GetWorkGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
