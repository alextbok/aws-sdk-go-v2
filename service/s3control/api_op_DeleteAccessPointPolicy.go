// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type DeleteAccessPointPolicyInput struct {
	_ struct{} `type:"structure"`

	// The account ID for the account that owns the specified access point.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The name of the access point whose policy you want to delete.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAccessPointPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccessPointPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccessPointPolicyInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccessPointPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.StringValue(v), metadata)
	}
	return nil
}

type DeleteAccessPointPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccessPointPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccessPointPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAccessPointPolicy = "DeleteAccessPointPolicy"

// DeleteAccessPointPolicyRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Deletes the access point policy for the specified access point.
//
//    // Example sending a request using DeleteAccessPointPolicyRequest.
//    req := client.DeleteAccessPointPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/DeleteAccessPointPolicy
func (c *Client) DeleteAccessPointPolicyRequest(input *DeleteAccessPointPolicyInput) DeleteAccessPointPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteAccessPointPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20180820/accesspoint/{name}/policy",
	}

	if input == nil {
		input = &DeleteAccessPointPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteAccessPointPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))

	return DeleteAccessPointPolicyRequest{Request: req, Input: input, Copy: c.DeleteAccessPointPolicyRequest}
}

// DeleteAccessPointPolicyRequest is the request type for the
// DeleteAccessPointPolicy API operation.
type DeleteAccessPointPolicyRequest struct {
	*aws.Request
	Input *DeleteAccessPointPolicyInput
	Copy  func(*DeleteAccessPointPolicyInput) DeleteAccessPointPolicyRequest
}

// Send marshals and sends the DeleteAccessPointPolicy API request.
func (r DeleteAccessPointPolicyRequest) Send(ctx context.Context) (*DeleteAccessPointPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccessPointPolicyResponse{
		DeleteAccessPointPolicyOutput: r.Request.Data.(*DeleteAccessPointPolicyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccessPointPolicyResponse is the response type for the
// DeleteAccessPointPolicy API operation.
type DeleteAccessPointPolicyResponse struct {
	*DeleteAccessPointPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccessPointPolicy request.
func (r *DeleteAccessPointPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
