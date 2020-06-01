// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DisableAWSServiceAccessInput struct {
	_ struct{} `type:"structure"`

	// The service principal name of the AWS service for which you want to disable
	// integration with your organization. This is typically in the form of a URL,
	// such as service-abbreviation.amazonaws.com.
	//
	// ServicePrincipal is a required field
	ServicePrincipal *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisableAWSServiceAccessInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableAWSServiceAccessInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableAWSServiceAccessInput"}

	if s.ServicePrincipal == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServicePrincipal"))
	}
	if s.ServicePrincipal != nil && len(*s.ServicePrincipal) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServicePrincipal", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisableAWSServiceAccessOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableAWSServiceAccessOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableAWSServiceAccess = "DisableAWSServiceAccess"

// DisableAWSServiceAccessRequest returns a request value for making API operation for
// AWS Organizations.
//
// Disables the integration of an AWS service (the service that is specified
// by ServicePrincipal) with AWS Organizations. When you disable integration,
// the specified service no longer can create a service-linked role (http://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in new accounts in your organization. This means the service can't perform
// operations on your behalf on any new accounts in your organization. The service
// can still perform operations in older accounts until the service completes
// its clean-up from AWS Organizations.
//
// We recommend that you disable integration between AWS Organizations and the
// specified AWS service by using the console or commands that are provided
// by the specified service. Doing so ensures that the other service is aware
// that it can clean up any resources that are required only for the integration.
// How the service cleans up its resources in the organization's accounts depends
// on that service. For more information, see the documentation for the other
// AWS service.
//
// After you perform the DisableAWSServiceAccess operation, the specified service
// can no longer perform operations in your organization's accounts unless the
// operations are explicitly permitted by the IAM policies that are attached
// to your roles.
//
// For more information about integrating other services with AWS Organizations,
// including the list of services that work with Organizations, see Integrating
// AWS Organizations with Other AWS Services (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the AWS Organizations User Guide.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using DisableAWSServiceAccessRequest.
//    req := client.DisableAWSServiceAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DisableAWSServiceAccess
func (c *Client) DisableAWSServiceAccessRequest(input *DisableAWSServiceAccessInput) DisableAWSServiceAccessRequest {
	op := &aws.Operation{
		Name:       opDisableAWSServiceAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableAWSServiceAccessInput{}
	}

	req := c.newRequest(op, input, &DisableAWSServiceAccessOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DisableAWSServiceAccessRequest{Request: req, Input: input, Copy: c.DisableAWSServiceAccessRequest}
}

// DisableAWSServiceAccessRequest is the request type for the
// DisableAWSServiceAccess API operation.
type DisableAWSServiceAccessRequest struct {
	*aws.Request
	Input *DisableAWSServiceAccessInput
	Copy  func(*DisableAWSServiceAccessInput) DisableAWSServiceAccessRequest
}

// Send marshals and sends the DisableAWSServiceAccess API request.
func (r DisableAWSServiceAccessRequest) Send(ctx context.Context) (*DisableAWSServiceAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableAWSServiceAccessResponse{
		DisableAWSServiceAccessOutput: r.Request.Data.(*DisableAWSServiceAccessOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableAWSServiceAccessResponse is the response type for the
// DisableAWSServiceAccess API operation.
type DisableAWSServiceAccessResponse struct {
	*DisableAWSServiceAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableAWSServiceAccess request.
func (r *DisableAWSServiceAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
