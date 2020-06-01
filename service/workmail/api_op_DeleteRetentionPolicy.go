// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRetentionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The retention policy ID.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The organization ID.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRetentionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRetentionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRetentionPolicyInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRetentionPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRetentionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRetentionPolicy = "DeleteRetentionPolicy"

// DeleteRetentionPolicyRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Deletes the specified retention policy from the specified organization.
//
//    // Example sending a request using DeleteRetentionPolicyRequest.
//    req := client.DeleteRetentionPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DeleteRetentionPolicy
func (c *Client) DeleteRetentionPolicyRequest(input *DeleteRetentionPolicyInput) DeleteRetentionPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteRetentionPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRetentionPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteRetentionPolicyOutput{})

	return DeleteRetentionPolicyRequest{Request: req, Input: input, Copy: c.DeleteRetentionPolicyRequest}
}

// DeleteRetentionPolicyRequest is the request type for the
// DeleteRetentionPolicy API operation.
type DeleteRetentionPolicyRequest struct {
	*aws.Request
	Input *DeleteRetentionPolicyInput
	Copy  func(*DeleteRetentionPolicyInput) DeleteRetentionPolicyRequest
}

// Send marshals and sends the DeleteRetentionPolicy API request.
func (r DeleteRetentionPolicyRequest) Send(ctx context.Context) (*DeleteRetentionPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRetentionPolicyResponse{
		DeleteRetentionPolicyOutput: r.Request.Data.(*DeleteRetentionPolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRetentionPolicyResponse is the response type for the
// DeleteRetentionPolicy API operation.
type DeleteRetentionPolicyResponse struct {
	*DeleteRetentionPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRetentionPolicy request.
func (r *DeleteRetentionPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
