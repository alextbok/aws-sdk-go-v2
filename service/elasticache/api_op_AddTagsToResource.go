// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an AddTagsToResource operation.
type AddTagsToResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource to which the tags are to be
	// added, for example arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster
	// or arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot. ElastiCache
	// resources are cluster and snapshot.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// ResourceName is a required field
	ResourceName *string `type:"string" required:"true"`

	// A list of cost allocation tags to be added to this resource. A tag is a key-value
	// pair. A tag key must be accompanied by a tag value.
	//
	// Tags is a required field
	Tags []Tag `locationNameList:"Tag" type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsToResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToResourceInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output from the AddTagsToResource, ListTagsForResource, and
// RemoveTagsFromResource operations.
type AddTagsToResourceOutput struct {
	_ struct{} `type:"structure"`

	// A list of cost allocation tags as key-value pairs.
	TagList []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s AddTagsToResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddTagsToResource = "AddTagsToResource"

// AddTagsToResourceRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Adds up to 50 cost allocation tags to the named resource. A cost allocation
// tag is a key-value pair where the key and value are case-sensitive. You can
// use cost allocation tags to categorize and track your AWS costs.
//
// When you apply tags to your ElastiCache resources, AWS generates a cost allocation
// report as a comma-separated value (CSV) file with your usage and costs aggregated
// by your tags. You can apply tags that represent business categories (such
// as cost centers, application names, or owners) to organize your costs across
// multiple services. For more information, see Using Cost Allocation Tags in
// Amazon ElastiCache (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Tagging.html)
// in the ElastiCache User Guide.
//
//    // Example sending a request using AddTagsToResourceRequest.
//    req := client.AddTagsToResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/AddTagsToResource
func (c *Client) AddTagsToResourceRequest(input *AddTagsToResourceInput) AddTagsToResourceRequest {
	op := &aws.Operation{
		Name:       opAddTagsToResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsToResourceInput{}
	}

	req := c.newRequest(op, input, &AddTagsToResourceOutput{})

	return AddTagsToResourceRequest{Request: req, Input: input, Copy: c.AddTagsToResourceRequest}
}

// AddTagsToResourceRequest is the request type for the
// AddTagsToResource API operation.
type AddTagsToResourceRequest struct {
	*aws.Request
	Input *AddTagsToResourceInput
	Copy  func(*AddTagsToResourceInput) AddTagsToResourceRequest
}

// Send marshals and sends the AddTagsToResource API request.
func (r AddTagsToResourceRequest) Send(ctx context.Context) (*AddTagsToResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsToResourceResponse{
		AddTagsToResourceOutput: r.Request.Data.(*AddTagsToResourceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsToResourceResponse is the response type for the
// AddTagsToResource API operation.
type AddTagsToResourceResponse struct {
	*AddTagsToResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTagsToResource request.
func (r *AddTagsToResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
