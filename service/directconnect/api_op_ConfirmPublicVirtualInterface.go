// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ConfirmPublicVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the virtual interface.
	//
	// VirtualInterfaceId is a required field
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s ConfirmPublicVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmPublicVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmPublicVirtualInterfaceInput"}

	if s.VirtualInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ConfirmPublicVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// The state of the virtual interface. The following are the possible values:
	//
	//    * confirming: The creation of the virtual interface is pending confirmation
	//    from the virtual interface owner. If the owner of the virtual interface
	//    is different from the owner of the connection on which it is provisioned,
	//    then the virtual interface will remain in this state until it is confirmed
	//    by the virtual interface owner.
	//
	//    * verifying: This state only applies to public virtual interfaces. Each
	//    public virtual interface needs validation before the virtual interface
	//    can be created.
	//
	//    * pending: A virtual interface is in this state from the time that it
	//    is created until the virtual interface is ready to forward traffic.
	//
	//    * available: A virtual interface that is able to forward traffic.
	//
	//    * down: A virtual interface that is BGP down.
	//
	//    * deleting: A virtual interface is in this state immediately after calling
	//    DeleteVirtualInterface until it can no longer forward traffic.
	//
	//    * deleted: A virtual interface that cannot forward traffic.
	//
	//    * rejected: The virtual interface owner has declined creation of the virtual
	//    interface. If a virtual interface in the Confirming state is deleted by
	//    the virtual interface owner, the virtual interface enters the Rejected
	//    state.
	//
	//    * unknown: The state of the virtual interface is not available.
	VirtualInterfaceState VirtualInterfaceState `locationName:"virtualInterfaceState" type:"string" enum:"true"`
}

// String returns the string representation
func (s ConfirmPublicVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opConfirmPublicVirtualInterface = "ConfirmPublicVirtualInterface"

// ConfirmPublicVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Accepts ownership of a public virtual interface created by another AWS account.
//
// After the virtual interface owner makes this call, the specified virtual
// interface is created and made available to handle traffic.
//
//    // Example sending a request using ConfirmPublicVirtualInterfaceRequest.
//    req := client.ConfirmPublicVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/ConfirmPublicVirtualInterface
func (c *Client) ConfirmPublicVirtualInterfaceRequest(input *ConfirmPublicVirtualInterfaceInput) ConfirmPublicVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opConfirmPublicVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConfirmPublicVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &ConfirmPublicVirtualInterfaceOutput{})

	return ConfirmPublicVirtualInterfaceRequest{Request: req, Input: input, Copy: c.ConfirmPublicVirtualInterfaceRequest}
}

// ConfirmPublicVirtualInterfaceRequest is the request type for the
// ConfirmPublicVirtualInterface API operation.
type ConfirmPublicVirtualInterfaceRequest struct {
	*aws.Request
	Input *ConfirmPublicVirtualInterfaceInput
	Copy  func(*ConfirmPublicVirtualInterfaceInput) ConfirmPublicVirtualInterfaceRequest
}

// Send marshals and sends the ConfirmPublicVirtualInterface API request.
func (r ConfirmPublicVirtualInterfaceRequest) Send(ctx context.Context) (*ConfirmPublicVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConfirmPublicVirtualInterfaceResponse{
		ConfirmPublicVirtualInterfaceOutput: r.Request.Data.(*ConfirmPublicVirtualInterfaceOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConfirmPublicVirtualInterfaceResponse is the response type for the
// ConfirmPublicVirtualInterface API operation.
type ConfirmPublicVirtualInterfaceResponse struct {
	*ConfirmPublicVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConfirmPublicVirtualInterface request.
func (r *ConfirmPublicVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
