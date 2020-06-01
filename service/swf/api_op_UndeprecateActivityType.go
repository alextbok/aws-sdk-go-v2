// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type UndeprecateActivityTypeInput struct {
	_ struct{} `type:"structure"`

	// The activity type to undeprecate.
	//
	// ActivityType is a required field
	ActivityType *ActivityType `locationName:"activityType" type:"structure" required:"true"`

	// The name of the domain of the deprecated activity type.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UndeprecateActivityTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UndeprecateActivityTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UndeprecateActivityTypeInput"}

	if s.ActivityType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivityType"))
	}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}
	if s.ActivityType != nil {
		if err := s.ActivityType.Validate(); err != nil {
			invalidParams.AddNested("ActivityType", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UndeprecateActivityTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UndeprecateActivityTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opUndeprecateActivityType = "UndeprecateActivityType"

// UndeprecateActivityTypeRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Undeprecates a previously deprecated activity type. After an activity type
// has been undeprecated, you can create new tasks of that activity type.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * Constrain the following parameters by using a Condition element with
//    the appropriate keys. activityType.name: String constraint. The key is
//    swf:activityType.name. activityType.version: String constraint. The key
//    is swf:activityType.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using UndeprecateActivityTypeRequest.
//    req := client.UndeprecateActivityTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UndeprecateActivityTypeRequest(input *UndeprecateActivityTypeInput) UndeprecateActivityTypeRequest {
	op := &aws.Operation{
		Name:       opUndeprecateActivityType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UndeprecateActivityTypeInput{}
	}

	req := c.newRequest(op, input, &UndeprecateActivityTypeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UndeprecateActivityTypeRequest{Request: req, Input: input, Copy: c.UndeprecateActivityTypeRequest}
}

// UndeprecateActivityTypeRequest is the request type for the
// UndeprecateActivityType API operation.
type UndeprecateActivityTypeRequest struct {
	*aws.Request
	Input *UndeprecateActivityTypeInput
	Copy  func(*UndeprecateActivityTypeInput) UndeprecateActivityTypeRequest
}

// Send marshals and sends the UndeprecateActivityType API request.
func (r UndeprecateActivityTypeRequest) Send(ctx context.Context) (*UndeprecateActivityTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UndeprecateActivityTypeResponse{
		UndeprecateActivityTypeOutput: r.Request.Data.(*UndeprecateActivityTypeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UndeprecateActivityTypeResponse is the response type for the
// UndeprecateActivityType API operation.
type UndeprecateActivityTypeResponse struct {
	*UndeprecateActivityTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UndeprecateActivityType request.
func (r *UndeprecateActivityTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
