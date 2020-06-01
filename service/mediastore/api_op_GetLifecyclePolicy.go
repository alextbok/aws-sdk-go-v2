// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLifecyclePolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that the object lifecycle policy is assigned to.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLifecyclePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLifecyclePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLifecyclePolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLifecyclePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The object lifecycle policy that is assigned to the container.
	//
	// LifecyclePolicy is a required field
	LifecyclePolicy *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetLifecyclePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLifecyclePolicy = "GetLifecyclePolicy"

// GetLifecyclePolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Retrieves the object lifecycle policy that is assigned to a container.
//
//    // Example sending a request using GetLifecyclePolicyRequest.
//    req := client.GetLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/GetLifecyclePolicy
func (c *Client) GetLifecyclePolicyRequest(input *GetLifecyclePolicyInput) GetLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opGetLifecyclePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &GetLifecyclePolicyOutput{})

	return GetLifecyclePolicyRequest{Request: req, Input: input, Copy: c.GetLifecyclePolicyRequest}
}

// GetLifecyclePolicyRequest is the request type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyRequest struct {
	*aws.Request
	Input *GetLifecyclePolicyInput
	Copy  func(*GetLifecyclePolicyInput) GetLifecyclePolicyRequest
}

// Send marshals and sends the GetLifecyclePolicy API request.
func (r GetLifecyclePolicyRequest) Send(ctx context.Context) (*GetLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLifecyclePolicyResponse{
		GetLifecyclePolicyOutput: r.Request.Data.(*GetLifecyclePolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLifecyclePolicyResponse is the response type for the
// GetLifecyclePolicy API operation.
type GetLifecyclePolicyResponse struct {
	*GetLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLifecyclePolicy request.
func (r *GetLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
