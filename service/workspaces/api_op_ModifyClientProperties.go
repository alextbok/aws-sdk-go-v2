// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyClientPropertiesInput struct {
	_ struct{} `type:"structure"`

	// Information about the Amazon WorkSpaces client.
	//
	// ClientProperties is a required field
	ClientProperties *Properties `type:"structure" required:"true"`

	// The resource identifiers, in the form of directory IDs.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyClientPropertiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyClientPropertiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyClientPropertiesInput"}

	if s.ClientProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientProperties"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyClientPropertiesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyClientPropertiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyClientProperties = "ModifyClientProperties"

// ModifyClientPropertiesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Modifies the properties of the specified Amazon WorkSpaces clients.
//
//    // Example sending a request using ModifyClientPropertiesRequest.
//    req := client.ModifyClientPropertiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ModifyClientProperties
func (c *Client) ModifyClientPropertiesRequest(input *ModifyClientPropertiesInput) ModifyClientPropertiesRequest {
	op := &aws.Operation{
		Name:       opModifyClientProperties,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyClientPropertiesInput{}
	}

	req := c.newRequest(op, input, &ModifyClientPropertiesOutput{})

	return ModifyClientPropertiesRequest{Request: req, Input: input, Copy: c.ModifyClientPropertiesRequest}
}

// ModifyClientPropertiesRequest is the request type for the
// ModifyClientProperties API operation.
type ModifyClientPropertiesRequest struct {
	*aws.Request
	Input *ModifyClientPropertiesInput
	Copy  func(*ModifyClientPropertiesInput) ModifyClientPropertiesRequest
}

// Send marshals and sends the ModifyClientProperties API request.
func (r ModifyClientPropertiesRequest) Send(ctx context.Context) (*ModifyClientPropertiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyClientPropertiesResponse{
		ModifyClientPropertiesOutput: r.Request.Data.(*ModifyClientPropertiesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyClientPropertiesResponse is the response type for the
// ModifyClientProperties API operation.
type ModifyClientPropertiesResponse struct {
	*ModifyClientPropertiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyClientProperties request.
func (r *ModifyClientPropertiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
