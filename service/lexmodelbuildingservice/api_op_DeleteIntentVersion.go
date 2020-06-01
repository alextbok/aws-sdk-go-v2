// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteIntentVersionInput struct {
	_ struct{} `type:"structure"`

	// The name of the intent.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// The version of the intent to delete. You cannot delete the $LATEST version
	// of the intent. To delete the $LATEST version, use the DeleteIntent operation.
	//
	// Version is a required field
	Version *string `location:"uri" locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIntentVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIntentVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIntentVersionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteIntentVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteIntentVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteIntentVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteIntentVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteIntentVersion = "DeleteIntentVersion"

// DeleteIntentVersionRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Deletes a specific version of an intent. To delete all versions of a intent,
// use the DeleteIntent operation.
//
// This operation requires permissions for the lex:DeleteIntentVersion action.
//
//    // Example sending a request using DeleteIntentVersionRequest.
//    req := client.DeleteIntentVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteIntentVersion
func (c *Client) DeleteIntentVersionRequest(input *DeleteIntentVersionInput) DeleteIntentVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteIntentVersion,
		HTTPMethod: "DELETE",
		HTTPPath:   "/intents/{name}/versions/{version}",
	}

	if input == nil {
		input = &DeleteIntentVersionInput{}
	}

	req := c.newRequest(op, input, &DeleteIntentVersionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteIntentVersionRequest{Request: req, Input: input, Copy: c.DeleteIntentVersionRequest}
}

// DeleteIntentVersionRequest is the request type for the
// DeleteIntentVersion API operation.
type DeleteIntentVersionRequest struct {
	*aws.Request
	Input *DeleteIntentVersionInput
	Copy  func(*DeleteIntentVersionInput) DeleteIntentVersionRequest
}

// Send marshals and sends the DeleteIntentVersion API request.
func (r DeleteIntentVersionRequest) Send(ctx context.Context) (*DeleteIntentVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIntentVersionResponse{
		DeleteIntentVersionOutput: r.Request.Data.(*DeleteIntentVersionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIntentVersionResponse is the response type for the
// DeleteIntentVersion API operation.
type DeleteIntentVersionResponse struct {
	*DeleteIntentVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIntentVersion request.
func (r *DeleteIntentVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
