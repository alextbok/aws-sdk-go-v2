// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetApprovalRuleTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the approval rule template for which you want to get information.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApprovalRuleTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApprovalRuleTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApprovalRuleTemplateInput"}

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

type GetApprovalRuleTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The content and structure of the approval rule template.
	//
	// ApprovalRuleTemplate is a required field
	ApprovalRuleTemplate *ApprovalRuleTemplate `locationName:"approvalRuleTemplate" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetApprovalRuleTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetApprovalRuleTemplate = "GetApprovalRuleTemplate"

// GetApprovalRuleTemplateRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about a specified approval rule template.
//
//    // Example sending a request using GetApprovalRuleTemplateRequest.
//    req := client.GetApprovalRuleTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetApprovalRuleTemplate
func (c *Client) GetApprovalRuleTemplateRequest(input *GetApprovalRuleTemplateInput) GetApprovalRuleTemplateRequest {
	op := &aws.Operation{
		Name:       opGetApprovalRuleTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetApprovalRuleTemplateInput{}
	}

	req := c.newRequest(op, input, &GetApprovalRuleTemplateOutput{})

	return GetApprovalRuleTemplateRequest{Request: req, Input: input, Copy: c.GetApprovalRuleTemplateRequest}
}

// GetApprovalRuleTemplateRequest is the request type for the
// GetApprovalRuleTemplate API operation.
type GetApprovalRuleTemplateRequest struct {
	*aws.Request
	Input *GetApprovalRuleTemplateInput
	Copy  func(*GetApprovalRuleTemplateInput) GetApprovalRuleTemplateRequest
}

// Send marshals and sends the GetApprovalRuleTemplate API request.
func (r GetApprovalRuleTemplateRequest) Send(ctx context.Context) (*GetApprovalRuleTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApprovalRuleTemplateResponse{
		GetApprovalRuleTemplateOutput: r.Request.Data.(*GetApprovalRuleTemplateOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApprovalRuleTemplateResponse is the response type for the
// GetApprovalRuleTemplate API operation.
type GetApprovalRuleTemplateResponse struct {
	*GetApprovalRuleTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApprovalRuleTemplate request.
func (r *GetApprovalRuleTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
