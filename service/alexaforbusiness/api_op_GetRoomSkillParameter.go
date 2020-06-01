// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRoomSkillParameterInput struct {
	_ struct{} `type:"structure"`

	// The room skill parameter key for which to get details. Required.
	//
	// ParameterKey is a required field
	ParameterKey *string `min:"1" type:"string" required:"true"`

	// The ARN of the room from which to get the room skill parameter details.
	RoomArn *string `type:"string"`

	// The ARN of the skill from which to get the room skill parameter details.
	// Required.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetRoomSkillParameterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRoomSkillParameterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRoomSkillParameterInput"}

	if s.ParameterKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterKey"))
	}
	if s.ParameterKey != nil && len(*s.ParameterKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParameterKey", 1))
	}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRoomSkillParameterOutput struct {
	_ struct{} `type:"structure"`

	// The details of the room skill parameter requested. Required.
	RoomSkillParameter *RoomSkillParameter `type:"structure"`
}

// String returns the string representation
func (s GetRoomSkillParameterOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRoomSkillParameter = "GetRoomSkillParameter"

// GetRoomSkillParameterRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets room skill parameter details by room, skill, and parameter key ARN.
//
//    // Example sending a request using GetRoomSkillParameterRequest.
//    req := client.GetRoomSkillParameterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetRoomSkillParameter
func (c *Client) GetRoomSkillParameterRequest(input *GetRoomSkillParameterInput) GetRoomSkillParameterRequest {
	op := &aws.Operation{
		Name:       opGetRoomSkillParameter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRoomSkillParameterInput{}
	}

	req := c.newRequest(op, input, &GetRoomSkillParameterOutput{})

	return GetRoomSkillParameterRequest{Request: req, Input: input, Copy: c.GetRoomSkillParameterRequest}
}

// GetRoomSkillParameterRequest is the request type for the
// GetRoomSkillParameter API operation.
type GetRoomSkillParameterRequest struct {
	*aws.Request
	Input *GetRoomSkillParameterInput
	Copy  func(*GetRoomSkillParameterInput) GetRoomSkillParameterRequest
}

// Send marshals and sends the GetRoomSkillParameter API request.
func (r GetRoomSkillParameterRequest) Send(ctx context.Context) (*GetRoomSkillParameterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRoomSkillParameterResponse{
		GetRoomSkillParameterOutput: r.Request.Data.(*GetRoomSkillParameterOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRoomSkillParameterResponse is the response type for the
// GetRoomSkillParameter API operation.
type GetRoomSkillParameterResponse struct {
	*GetRoomSkillParameterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRoomSkillParameter request.
func (r *GetRoomSkillParameterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
