// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetMLTransformRequest
type GetMLTransformInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the transform, generated at the time that the transform
	// was created.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMLTransformInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMLTransformInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMLTransformInput"}

	if s.TransformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransformId"))
	}
	if s.TransformId != nil && len(*s.TransformId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransformId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetMLTransformResponse
type GetMLTransformOutput struct {
	_ struct{} `type:"structure"`

	// The date and time when the transform was created.
	CreatedOn *time.Time `type:"timestamp"`

	// A description of the transform.
	Description *string `type:"string"`

	// The latest evaluation metrics.
	EvaluationMetrics *EvaluationMetrics `type:"structure"`

	// A list of AWS Glue table definitions used by the transform.
	InputRecordTables []GlueTable `type:"list"`

	// The number of labels available for this transform.
	LabelCount *int64 `type:"integer"`

	// The date and time when the transform was last modified.
	LastModifiedOn *time.Time `type:"timestamp"`

	// The number of AWS Glue data processing units (DPUs) that are allocated to
	// task runs for this transform. You can allocate from 2 to 100 DPUs; the default
	// is 10. A DPU is a relative measure of processing power that consists of 4
	// vCPUs of compute capacity and 16 GB of memory. For more information, see
	// the AWS Glue pricing page (https://aws.amazon.com/glue/pricing/).
	//
	// When the WorkerType field is set to a value other than Standard, the MaxCapacity
	// field is set automatically and becomes read-only.
	MaxCapacity *float64 `type:"double"`

	// The maximum number of times to retry a task for this transform after a task
	// run fails.
	MaxRetries *int64 `type:"integer"`

	// The unique name given to the transform when it was created.
	Name *string `min:"1" type:"string"`

	// The number of workers of a defined workerType that are allocated when this
	// task runs.
	NumberOfWorkers *int64 `type:"integer"`

	// The configuration parameters that are specific to the algorithm used.
	Parameters *TransformParameters `type:"structure"`

	// The name or Amazon Resource Name (ARN) of the IAM role with the required
	// permissions.
	Role *string `type:"string"`

	// The Map<Column, Type> object that represents the schema that this transform
	// accepts. Has an upper bound of 100 columns.
	Schema []SchemaColumn `type:"list"`

	// The last known status of the transform (to indicate whether it can be used
	// or not). One of "NOT_READY", "READY", or "DELETING".
	Status TransformStatusType `type:"string" enum:"true"`

	// The timeout for a task run for this transform in minutes. This is the maximum
	// time that a task run for this transform can consume resources before it is
	// terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout *int64 `min:"1" type:"integer"`

	// The unique identifier of the transform, generated at the time that the transform
	// was created.
	TransformId *string `min:"1" type:"string"`

	// The type of predefined worker that is allocated when this task runs. Accepts
	// a value of Standard, G.1X, or G.2X.
	//
	//    * For the Standard worker type, each worker provides 4 vCPU, 16 GB of
	//    memory and a 50GB disk, and 2 executors per worker.
	//
	//    * For the G.1X worker type, each worker provides 4 vCPU, 16 GB of memory
	//    and a 64GB disk, and 1 executor per worker.
	//
	//    * For the G.2X worker type, each worker provides 8 vCPU, 32 GB of memory
	//    and a 128GB disk, and 1 executor per worker.
	WorkerType WorkerType `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetMLTransformOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMLTransform = "GetMLTransform"

// GetMLTransformRequest returns a request value for making API operation for
// AWS Glue.
//
// Gets an AWS Glue machine learning transform artifact and all its corresponding
// metadata. Machine learning transforms are a special type of transform that
// use machine learning to learn the details of the transformation to be performed
// by learning from examples provided by humans. These transformations are then
// saved by AWS Glue. You can retrieve their metadata by calling GetMLTransform.
//
//    // Example sending a request using GetMLTransformRequest.
//    req := client.GetMLTransformRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetMLTransform
func (c *Client) GetMLTransformRequest(input *GetMLTransformInput) GetMLTransformRequest {
	op := &aws.Operation{
		Name:       opGetMLTransform,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMLTransformInput{}
	}

	req := c.newRequest(op, input, &GetMLTransformOutput{})
	return GetMLTransformRequest{Request: req, Input: input, Copy: c.GetMLTransformRequest}
}

// GetMLTransformRequest is the request type for the
// GetMLTransform API operation.
type GetMLTransformRequest struct {
	*aws.Request
	Input *GetMLTransformInput
	Copy  func(*GetMLTransformInput) GetMLTransformRequest
}

// Send marshals and sends the GetMLTransform API request.
func (r GetMLTransformRequest) Send(ctx context.Context) (*GetMLTransformResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMLTransformResponse{
		GetMLTransformOutput: r.Request.Data.(*GetMLTransformOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMLTransformResponse is the response type for the
// GetMLTransform API operation.
type GetMLTransformResponse struct {
	*GetMLTransformOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMLTransform request.
func (r *GetMLTransformResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
