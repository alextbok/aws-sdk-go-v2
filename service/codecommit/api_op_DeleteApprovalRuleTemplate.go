// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApprovalRuleTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the approval rule template to delete.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApprovalRuleTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApprovalRuleTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApprovalRuleTemplateInput"}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApprovalRuleTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the deleted approval rule template. If the template
	// has been previously deleted, the only response is a 200 OK.
	//
	// ApprovalRuleTemplateId is a required field
	ApprovalRuleTemplateId *string `locationName:"approvalRuleTemplateId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApprovalRuleTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApprovalRuleTemplate = "DeleteApprovalRuleTemplate"

// DeleteApprovalRuleTemplateRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Deletes a specified approval rule template. Deleting a template does not
// remove approval rules on pull requests already created with the template.
//
//    // Example sending a request using DeleteApprovalRuleTemplateRequest.
//    req := client.DeleteApprovalRuleTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeleteApprovalRuleTemplate
func (c *Client) DeleteApprovalRuleTemplateRequest(input *DeleteApprovalRuleTemplateInput) DeleteApprovalRuleTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteApprovalRuleTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApprovalRuleTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteApprovalRuleTemplateOutput{})

	return DeleteApprovalRuleTemplateRequest{Request: req, Input: input, Copy: c.DeleteApprovalRuleTemplateRequest}
}

// DeleteApprovalRuleTemplateRequest is the request type for the
// DeleteApprovalRuleTemplate API operation.
type DeleteApprovalRuleTemplateRequest struct {
	*aws.Request
	Input *DeleteApprovalRuleTemplateInput
	Copy  func(*DeleteApprovalRuleTemplateInput) DeleteApprovalRuleTemplateRequest
}

// Send marshals and sends the DeleteApprovalRuleTemplate API request.
func (r DeleteApprovalRuleTemplateRequest) Send(ctx context.Context) (*DeleteApprovalRuleTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApprovalRuleTemplateResponse{
		DeleteApprovalRuleTemplateOutput: r.Request.Data.(*DeleteApprovalRuleTemplateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApprovalRuleTemplateResponse is the response type for the
// DeleteApprovalRuleTemplate API operation.
type DeleteApprovalRuleTemplateResponse struct {
	*DeleteApprovalRuleTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApprovalRuleTemplate request.
func (r *DeleteApprovalRuleTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
