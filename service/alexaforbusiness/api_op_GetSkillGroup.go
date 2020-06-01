// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSkillGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the skill group for which to get details. Required.
	SkillGroupArn *string `type:"string"`
}

// String returns the string representation
func (s GetSkillGroupInput) String() string {
	return awsutil.Prettify(s)
}

type GetSkillGroupOutput struct {
	_ struct{} `type:"structure"`

	// The details of the skill group requested. Required.
	SkillGroup *SkillGroup `type:"structure"`
}

// String returns the string representation
func (s GetSkillGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSkillGroup = "GetSkillGroup"

// GetSkillGroupRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets skill group details by skill group ARN.
//
//    // Example sending a request using GetSkillGroupRequest.
//    req := client.GetSkillGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetSkillGroup
func (c *Client) GetSkillGroupRequest(input *GetSkillGroupInput) GetSkillGroupRequest {
	op := &aws.Operation{
		Name:       opGetSkillGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSkillGroupInput{}
	}

	req := c.newRequest(op, input, &GetSkillGroupOutput{})

	return GetSkillGroupRequest{Request: req, Input: input, Copy: c.GetSkillGroupRequest}
}

// GetSkillGroupRequest is the request type for the
// GetSkillGroup API operation.
type GetSkillGroupRequest struct {
	*aws.Request
	Input *GetSkillGroupInput
	Copy  func(*GetSkillGroupInput) GetSkillGroupRequest
}

// Send marshals and sends the GetSkillGroup API request.
func (r GetSkillGroupRequest) Send(ctx context.Context) (*GetSkillGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSkillGroupResponse{
		GetSkillGroupOutput: r.Request.Data.(*GetSkillGroupOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSkillGroupResponse is the response type for the
// GetSkillGroup API operation.
type GetSkillGroupResponse struct {
	*GetSkillGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSkillGroup request.
func (r *GetSkillGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
