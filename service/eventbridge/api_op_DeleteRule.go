// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteRuleInput struct {
	_ struct{} `type:"structure"`

	// The event bus associated with the rule. If you omit this, the default event
	// bus is used.
	EventBusName *string `min:"1" type:"string"`

	// If this is a managed rule, created by an AWS service on your behalf, you
	// must specify Force as True to delete the rule. This parameter is ignored
	// for rules that are not managed rules. You can check whether a rule is a managed
	// rule by using DescribeRule or ListRules and checking the ManagedBy field
	// of the response.
	Force *bool `type:"boolean"`

	// The name of the rule.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRuleInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRule = "DeleteRule"

// DeleteRuleRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Deletes the specified rule.
//
// Before you can delete the rule, you must remove all targets, using RemoveTargets.
//
// When you delete a rule, incoming events might continue to match to the deleted
// rule. Allow a short period of time for changes to take effect.
//
// Managed rules are rules created and managed by another AWS service on your
// behalf. These rules are created by those other AWS services to support functionality
// in those services. You can delete these rules using the Force option, but
// you should do so only if you are sure the other service is not still using
// that rule.
//
//    // Example sending a request using DeleteRuleRequest.
//    req := client.DeleteRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/DeleteRule
func (c *Client) DeleteRuleRequest(input *DeleteRuleInput) DeleteRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteRuleOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteRuleRequest{Request: req, Input: input, Copy: c.DeleteRuleRequest}
}

// DeleteRuleRequest is the request type for the
// DeleteRule API operation.
type DeleteRuleRequest struct {
	*aws.Request
	Input *DeleteRuleInput
	Copy  func(*DeleteRuleInput) DeleteRuleRequest
}

// Send marshals and sends the DeleteRule API request.
func (r DeleteRuleRequest) Send(ctx context.Context) (*DeleteRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRuleResponse{
		DeleteRuleOutput: r.Request.Data.(*DeleteRuleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRuleResponse is the response type for the
// DeleteRule API operation.
type DeleteRuleResponse struct {
	*DeleteRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRule request.
func (r *DeleteRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
