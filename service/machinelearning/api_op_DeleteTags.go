// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTagsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the tagged ML object. For example, exampleModelId.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The type of the tagged ML object.
	//
	// ResourceType is a required field
	ResourceType TaggableResourceType `type:"string" required:"true" enum:"true"`

	// One or more tags to delete.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTagsInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Amazon ML returns the following elements.
type DeleteTagsOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the ML object from which tags were deleted.
	ResourceId *string `min:"1" type:"string"`

	// The type of the ML object from which tags were deleted.
	ResourceType TaggableResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTags = "DeleteTags"

// DeleteTagsRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Deletes the specified tags associated with an ML object. After this operation
// is complete, you can't recover deleted tags.
//
// If you specify a tag that doesn't exist, Amazon ML ignores it.
//
//    // Example sending a request using DeleteTagsRequest.
//    req := client.DeleteTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
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
