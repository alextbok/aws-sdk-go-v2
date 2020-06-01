// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkflowTypeInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which this workflow type is registered.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The workflow type to describe.
	//
	// WorkflowType is a required field
	WorkflowType *WorkflowType `locationName:"workflowType" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeWorkflowTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkflowTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkflowTypeInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.WorkflowType == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowType"))
	}
	if s.WorkflowType != nil {
		if err := s.WorkflowType.Validate(); err != nil {
			invalidParams.AddNested("WorkflowType", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains details about a workflow type.
type DescribeWorkflowTypeOutput struct {
	_ struct{} `type:"structure"`

	// Configuration settings of the workflow type registered through RegisterWorkflowType
	//
	// Configuration is a required field
	Configuration *WorkflowTypeConfiguration `locationName:"configuration" type:"structure" required:"true"`

	// General information about the workflow type.
	//
	// The status of the workflow type (returned in the WorkflowTypeInfo structure)
	// can be one of the following.
	//
	//    * REGISTERED – The type is registered and available. Workers supporting
	//    this type should be running.
	//
	//    * DEPRECATED – The type was deprecated using DeprecateWorkflowType,
	//    but is still in use. You should keep workers supporting this type running.
	//    You cannot create new workflow executions of this type.
	//
	// TypeInfo is a required field
	TypeInfo *WorkflowTypeInfo `locationName:"typeInfo" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeWorkflowTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkflowType = "DescribeWorkflowType"

// DescribeWorkflowTypeRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns information about the specified workflow type. This includes configuration
// settings specified when the type was registered and other information such
// as creation date, current status, etc.
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
//    the appropriate keys. workflowType.name: String constraint. The key is
//    swf:workflowType.name. workflowType.version: String constraint. The key
//    is swf:workflowType.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using DescribeWorkflowTypeRequest.
//    req := client.DescribeWorkflowTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeWorkflowTypeRequest(input *DescribeWorkflowTypeInput) DescribeWorkflowTypeRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkflowType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeWorkflowTypeInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkflowTypeOutput{})

	return DescribeWorkflowTypeRequest{Request: req, Input: input, Copy: c.DescribeWorkflowTypeRequest}
}

// DescribeWorkflowTypeRequest is the request type for the
// DescribeWorkflowType API operation.
type DescribeWorkflowTypeRequest struct {
	*aws.Request
	Input *DescribeWorkflowTypeInput
	Copy  func(*DescribeWorkflowTypeInput) DescribeWorkflowTypeRequest
}

// Send marshals and sends the DescribeWorkflowType API request.
func (r DescribeWorkflowTypeRequest) Send(ctx context.Context) (*DescribeWorkflowTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkflowTypeResponse{
		DescribeWorkflowTypeOutput: r.Request.Data.(*DescribeWorkflowTypeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWorkflowTypeResponse is the response type for the
// DescribeWorkflowType API operation.
type DescribeWorkflowTypeResponse struct {
	*DescribeWorkflowTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkflowType request.
func (r *DescribeWorkflowTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
