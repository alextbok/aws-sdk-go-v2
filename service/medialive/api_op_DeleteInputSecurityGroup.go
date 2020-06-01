// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteInputSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	// InputSecurityGroupId is a required field
	InputSecurityGroupId *string `location:"uri" locationName:"inputSecurityGroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInputSecurityGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInputSecurityGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInputSecurityGroupInput"}

	if s.InputSecurityGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputSecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInputSecurityGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InputSecurityGroupId != nil {
		v := *s.InputSecurityGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "inputSecurityGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteInputSecurityGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteInputSecurityGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInputSecurityGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteInputSecurityGroup = "DeleteInputSecurityGroup"

// DeleteInputSecurityGroupRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Deletes an Input Security Group
//
//    // Example sending a request using DeleteInputSecurityGroupRequest.
//    req := client.DeleteInputSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DeleteInputSecurityGroup
func (c *Client) DeleteInputSecurityGroupRequest(input *DeleteInputSecurityGroupInput) DeleteInputSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteInputSecurityGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/prod/inputSecurityGroups/{inputSecurityGroupId}",
	}

	if input == nil {
		input = &DeleteInputSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteInputSecurityGroupOutput{})

	return DeleteInputSecurityGroupRequest{Request: req, Input: input, Copy: c.DeleteInputSecurityGroupRequest}
}

// DeleteInputSecurityGroupRequest is the request type for the
// DeleteInputSecurityGroup API operation.
type DeleteInputSecurityGroupRequest struct {
	*aws.Request
	Input *DeleteInputSecurityGroupInput
	Copy  func(*DeleteInputSecurityGroupInput) DeleteInputSecurityGroupRequest
}

// Send marshals and sends the DeleteInputSecurityGroup API request.
func (r DeleteInputSecurityGroupRequest) Send(ctx context.Context) (*DeleteInputSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInputSecurityGroupResponse{
		DeleteInputSecurityGroupOutput: r.Request.Data.(*DeleteInputSecurityGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInputSecurityGroupResponse is the response type for the
// DeleteInputSecurityGroup API operation.
type DeleteInputSecurityGroupResponse struct {
	*DeleteInputSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInputSecurityGroup request.
func (r *DeleteInputSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
