// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchApplyUpdateActionInput struct {
	_ struct{} `type:"structure"`

	// The cache cluster IDs
	CacheClusterIds []string `type:"list"`

	// The replication group IDs
	ReplicationGroupIds []string `type:"list"`

	// The unique ID of the service update
	//
	// ServiceUpdateName is a required field
	ServiceUpdateName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s BatchApplyUpdateActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchApplyUpdateActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchApplyUpdateActionInput"}

	if s.ServiceUpdateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceUpdateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchApplyUpdateActionOutput struct {
	_ struct{} `type:"structure"`

	// Update actions that have been processed successfully
	ProcessedUpdateActions []ProcessedUpdateAction `locationNameList:"ProcessedUpdateAction" type:"list"`

	// Update actions that haven't been processed successfully
	UnprocessedUpdateActions []UnprocessedUpdateAction `locationNameList:"UnprocessedUpdateAction" type:"list"`
}

// String returns the string representation
func (s BatchApplyUpdateActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchApplyUpdateAction = "BatchApplyUpdateAction"

// BatchApplyUpdateActionRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Apply the service update. For more information on service updates and applying
// them, see Applying Service Updates (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/applying-updates.html).
//
//    // Example sending a request using BatchApplyUpdateActionRequest.
//    req := client.BatchApplyUpdateActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/BatchApplyUpdateAction
func (c *Client) BatchApplyUpdateActionRequest(input *BatchApplyUpdateActionInput) BatchApplyUpdateActionRequest {
	op := &aws.Operation{
		Name:       opBatchApplyUpdateAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchApplyUpdateActionInput{}
	}

	req := c.newRequest(op, input, &BatchApplyUpdateActionOutput{})

	return BatchApplyUpdateActionRequest{Request: req, Input: input, Copy: c.BatchApplyUpdateActionRequest}
}

// BatchApplyUpdateActionRequest is the request type for the
// BatchApplyUpdateAction API operation.
type BatchApplyUpdateActionRequest struct {
	*aws.Request
	Input *BatchApplyUpdateActionInput
	Copy  func(*BatchApplyUpdateActionInput) BatchApplyUpdateActionRequest
}

// Send marshals and sends the BatchApplyUpdateAction API request.
func (r BatchApplyUpdateActionRequest) Send(ctx context.Context) (*BatchApplyUpdateActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchApplyUpdateActionResponse{
		BatchApplyUpdateActionOutput: r.Request.Data.(*BatchApplyUpdateActionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchApplyUpdateActionResponse is the response type for the
// BatchApplyUpdateAction API operation.
type BatchApplyUpdateActionResponse struct {
	*BatchApplyUpdateActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchApplyUpdateAction request.
func (r *BatchApplyUpdateActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
