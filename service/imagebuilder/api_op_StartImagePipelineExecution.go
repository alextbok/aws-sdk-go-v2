// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StartImagePipelineExecutionInput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the image pipeline that you want to manually
	// invoke.
	//
	// ImagePipelineArn is a required field
	ImagePipelineArn *string `locationName:"imagePipelineArn" type:"string" required:"true"`
}

// String returns the string representation
func (s StartImagePipelineExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartImagePipelineExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartImagePipelineExecutionInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.ImagePipelineArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImagePipelineArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartImagePipelineExecutionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImagePipelineArn != nil {
		v := *s.ImagePipelineArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imagePipelineArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartImagePipelineExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	ClientToken *string `locationName:"clientToken" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the image that was created by this request.
	ImageBuildVersionArn *string `locationName:"imageBuildVersionArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s StartImagePipelineExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartImagePipelineExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageBuildVersionArn != nil {
		v := *s.ImageBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opStartImagePipelineExecution = "StartImagePipelineExecution"

// StartImagePipelineExecutionRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Manually triggers a pipeline to create an image.
//
//    // Example sending a request using StartImagePipelineExecutionRequest.
//    req := client.StartImagePipelineExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/StartImagePipelineExecution
func (c *Client) StartImagePipelineExecutionRequest(input *StartImagePipelineExecutionInput) StartImagePipelineExecutionRequest {
	op := &aws.Operation{
		Name:       opStartImagePipelineExecution,
		HTTPMethod: "PUT",
		HTTPPath:   "/StartImagePipelineExecution",
	}

	if input == nil {
		input = &StartImagePipelineExecutionInput{}
	}

	req := c.newRequest(op, input, &StartImagePipelineExecutionOutput{})

	return StartImagePipelineExecutionRequest{Request: req, Input: input, Copy: c.StartImagePipelineExecutionRequest}
}

// StartImagePipelineExecutionRequest is the request type for the
// StartImagePipelineExecution API operation.
type StartImagePipelineExecutionRequest struct {
	*aws.Request
	Input *StartImagePipelineExecutionInput
	Copy  func(*StartImagePipelineExecutionInput) StartImagePipelineExecutionRequest
}

// Send marshals and sends the StartImagePipelineExecution API request.
func (r StartImagePipelineExecutionRequest) Send(ctx context.Context) (*StartImagePipelineExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartImagePipelineExecutionResponse{
		StartImagePipelineExecutionOutput: r.Request.Data.(*StartImagePipelineExecutionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartImagePipelineExecutionResponse is the response type for the
// StartImagePipelineExecution API operation.
type StartImagePipelineExecutionResponse struct {
	*StartImagePipelineExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartImagePipelineExecution request.
func (r *StartImagePipelineExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
