// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRulesOfIpGroupInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the group.
	//
	// GroupId is a required field
	GroupId *string `type:"string" required:"true"`

	// One or more rules.
	//
	// UserRules is a required field
	UserRules []IpRuleItem `type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateRulesOfIpGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRulesOfIpGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRulesOfIpGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if s.UserRules == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserRules"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRulesOfIpGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateRulesOfIpGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRulesOfIpGroup = "UpdateRulesOfIpGroup"

// UpdateRulesOfIpGroupRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Replaces the current rules of the specified IP access control group with
// the specified rules.
//
//    // Example sending a request using UpdateRulesOfIpGroupRequest.
//    req := client.UpdateRulesOfIpGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/UpdateRulesOfIpGroup
func (c *Client) UpdateRulesOfIpGroupRequest(input *UpdateRulesOfIpGroupInput) UpdateRulesOfIpGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateRulesOfIpGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRulesOfIpGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateRulesOfIpGroupOutput{})

	return UpdateRulesOfIpGroupRequest{Request: req, Input: input, Copy: c.UpdateRulesOfIpGroupRequest}
}

// UpdateRulesOfIpGroupRequest is the request type for the
// UpdateRulesOfIpGroup API operation.
type UpdateRulesOfIpGroupRequest struct {
	*aws.Request
	Input *UpdateRulesOfIpGroupInput
	Copy  func(*UpdateRulesOfIpGroupInput) UpdateRulesOfIpGroupRequest
}

// Send marshals and sends the UpdateRulesOfIpGroup API request.
func (r UpdateRulesOfIpGroupRequest) Send(ctx context.Context) (*UpdateRulesOfIpGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRulesOfIpGroupResponse{
		UpdateRulesOfIpGroupOutput: r.Request.Data.(*UpdateRulesOfIpGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRulesOfIpGroupResponse is the response type for the
// UpdateRulesOfIpGroup API operation.
type UpdateRulesOfIpGroupResponse struct {
	*UpdateRulesOfIpGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRulesOfIpGroup request.
func (r *UpdateRulesOfIpGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
