// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDiskInput struct {
	_ struct{} `type:"structure"`

	// The unique name of the disk you want to delete (e.g., my-disk).
	//
	// DiskName is a required field
	DiskName *string `locationName:"diskName" type:"string" required:"true"`

	// A Boolean value to indicate whether to delete the enabled add-ons for the
	// disk.
	ForceDeleteAddOns *bool `locationName:"forceDeleteAddOns" type:"boolean"`
}

// String returns the string representation
func (s DeleteDiskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDiskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDiskInput"}

	if s.DiskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDiskOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteDiskOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDisk = "DeleteDisk"

// DeleteDiskRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes the specified block storage disk. The disk must be in the available
// state (not attached to a Lightsail instance).
//
// The disk may remain in the deleting state for several minutes.
//
// The delete disk operation supports tag-based access control via resource
// tags applied to the resource identified by disk name. For more information,
// see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteDiskRequest.
//    req := client.DeleteDiskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteDisk
func (c *Client) DeleteDiskRequest(input *DeleteDiskInput) DeleteDiskRequest {
	op := &aws.Operation{
		Name:       opDeleteDisk,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDiskInput{}
	}

	req := c.newRequest(op, input, &DeleteDiskOutput{})

	return DeleteDiskRequest{Request: req, Input: input, Copy: c.DeleteDiskRequest}
}

// DeleteDiskRequest is the request type for the
// DeleteDisk API operation.
type DeleteDiskRequest struct {
	*aws.Request
	Input *DeleteDiskInput
	Copy  func(*DeleteDiskInput) DeleteDiskRequest
}

// Send marshals and sends the DeleteDisk API request.
func (r DeleteDiskRequest) Send(ctx context.Context) (*DeleteDiskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDiskResponse{
		DeleteDiskOutput: r.Request.Data.(*DeleteDiskOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDiskResponse is the response type for the
// DeleteDisk API operation.
type DeleteDiskResponse struct {
	*DeleteDiskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDisk request.
func (r *DeleteDiskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
