// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// ListVolumeInitiatorsInput
type ListVolumeInitiatorsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation
	// to return a list of gateway volumes for the gateway.
	//
	// VolumeARN is a required field
	VolumeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s ListVolumeInitiatorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVolumeInitiatorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVolumeInitiatorsInput"}

	if s.VolumeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeARN"))
	}
	if s.VolumeARN != nil && len(*s.VolumeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("VolumeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// ListVolumeInitiatorsOutput
type ListVolumeInitiatorsOutput struct {
	_ struct{} `type:"structure"`

	// The host names and port numbers of all iSCSI initiators that are connected
	// to the gateway.
	Initiators []string `type:"list"`
}

// String returns the string representation
func (s ListVolumeInitiatorsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListVolumeInitiators = "ListVolumeInitiators"

// ListVolumeInitiatorsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Lists iSCSI initiators that are connected to a volume. You can use this operation
// to determine whether a volume is being used or not. This operation is only
// supported in the cached volume and stored volume gateway types.
//
//    // Example sending a request using ListVolumeInitiatorsRequest.
//    req := client.ListVolumeInitiatorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ListVolumeInitiators
func (c *Client) ListVolumeInitiatorsRequest(input *ListVolumeInitiatorsInput) ListVolumeInitiatorsRequest {
	op := &aws.Operation{
		Name:       opListVolumeInitiators,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListVolumeInitiatorsInput{}
	}

	req := c.newRequest(op, input, &ListVolumeInitiatorsOutput{})

	return ListVolumeInitiatorsRequest{Request: req, Input: input, Copy: c.ListVolumeInitiatorsRequest}
}

// ListVolumeInitiatorsRequest is the request type for the
// ListVolumeInitiators API operation.
type ListVolumeInitiatorsRequest struct {
	*aws.Request
	Input *ListVolumeInitiatorsInput
	Copy  func(*ListVolumeInitiatorsInput) ListVolumeInitiatorsRequest
}

// Send marshals and sends the ListVolumeInitiators API request.
func (r ListVolumeInitiatorsRequest) Send(ctx context.Context) (*ListVolumeInitiatorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVolumeInitiatorsResponse{
		ListVolumeInitiatorsOutput: r.Request.Data.(*ListVolumeInitiatorsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListVolumeInitiatorsResponse is the response type for the
// ListVolumeInitiators API operation.
type ListVolumeInitiatorsResponse struct {
	*ListVolumeInitiatorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVolumeInitiators request.
func (r *ListVolumeInitiatorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
