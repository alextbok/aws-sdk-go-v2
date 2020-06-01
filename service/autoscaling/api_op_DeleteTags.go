// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteTagsInput struct {
	_ struct{} `type:"structure"`

	// One or more tags.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTagsInput"}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTags = "DeleteTags"

// DeleteTagsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Deletes the specified tags.
//
//    // Example sending a request using DeleteTagsRequest.
//    req := client.DeleteTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DeleteTags
func (c *Client) DeleteTagsRequest(input *DeleteTagsInput) DeleteTagsRequest {
	op := &aws.Operation{
		Name:       opDeleteTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTagsInput{}
	}

	req := c.newRequest(op, input, &DeleteTagsOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteTagsRequest{Request: req, Input: input, Copy: c.DeleteTagsRequest}
}

// DeleteTagsRequest is the request type for the
// DeleteTags API operation.
type DeleteTagsRequest struct {
	*aws.Request
	Input *DeleteTagsInput
	Copy  func(*DeleteTagsInput) DeleteTagsRequest
}

// Send marshals and sends the DeleteTags API request.
func (r DeleteTagsRequest) Send(ctx context.Context) (*DeleteTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTagsResponse{
		DeleteTagsOutput: r.Request.Data.(*DeleteTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTagsResponse is the response type for the
// DeleteTags API operation.
type DeleteTagsResponse struct {
	*DeleteTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTags request.
func (r *DeleteTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
