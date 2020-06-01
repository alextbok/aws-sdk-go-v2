// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDatasetInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset.
	//
	// DatasetArn is a required field
	DatasetArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDatasetInput"}

	if s.DatasetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDatasetOutput struct {
	_ struct{} `type:"structure"`

	// When the dataset was created.
	CreationTime *time.Time `type:"timestamp"`

	// The frequency of data collection.
	//
	// Valid intervals are Y (Year), M (Month), W (Week), D (Day), H (Hour), 30min
	// (30 minutes), 15min (15 minutes), 10min (10 minutes), 5min (5 minutes), and
	// 1min (1 minute). For example, "M" indicates every month and "30min" indicates
	// every 30 minutes.
	DataFrequency *string `type:"string"`

	// The Amazon Resource Name (ARN) of the dataset.
	DatasetArn *string `type:"string"`

	// The name of the dataset.
	DatasetName *string `min:"1" type:"string"`

	// The dataset type.
	DatasetType DatasetType `type:"string" enum:"true"`

	// The domain associated with the dataset.
	Domain Domain `type:"string" enum:"true"`

	// The AWS Key Management Service (KMS) key and the AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *EncryptionConfig `type:"structure"`

	// When you create a dataset, LastModificationTime is the same as CreationTime.
	// While data is being imported to the dataset, LastModificationTime is the
	// current time of the DescribeDataset call. After a CreateDatasetImportJob
	// operation has finished, LastModificationTime is when the import job completed
	// or failed.
	LastModificationTime *time.Time `type:"timestamp"`

	// An array of SchemaAttribute objects that specify the dataset fields. Each
	// SchemaAttribute specifies the name and data type of a field.
	Schema *Schema `type:"structure"`

	// The status of the dataset. States include:
	//
	//    * ACTIVE
	//
	//    * CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//    * DELETE_PENDING, DELETE_IN_PROGRESS, DELETE_FAILED
	//
	//    * UPDATE_PENDING, UPDATE_IN_PROGRESS, UPDATE_FAILED
	//
	// The UPDATE states apply while data is imported to the dataset from a call
	// to the CreateDatasetImportJob operation and reflect the status of the dataset
	// import job. For example, when the import job status is CREATE_IN_PROGRESS,
	// the status of the dataset is UPDATE_IN_PROGRESS.
	//
	// The Status of the dataset must be ACTIVE before you can import training data.
	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeDatasetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDataset = "DescribeDataset"

// DescribeDatasetRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Describes an Amazon Forecast dataset created using the CreateDataset operation.
//
// In addition to listing the parameters specified in the CreateDataset request,
// this operation includes the following dataset properties:
//
//    * CreationTime
//
//    * LastModificationTime
//
//    * Status
//
//    // Example sending a request using DescribeDatasetRequest.
//    req := client.DescribeDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/DescribeDataset
func (c *Client) DescribeDatasetRequest(input *DescribeDatasetInput) DescribeDatasetRequest {
	op := &aws.Operation{
		Name:       opDescribeDataset,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDatasetInput{}
	}

	req := c.newRequest(op, input, &DescribeDatasetOutput{})

	return DescribeDatasetRequest{Request: req, Input: input, Copy: c.DescribeDatasetRequest}
}

// DescribeDatasetRequest is the request type for the
// DescribeDataset API operation.
type DescribeDatasetRequest struct {
	*aws.Request
	Input *DescribeDatasetInput
	Copy  func(*DescribeDatasetInput) DescribeDatasetRequest
}

// Send marshals and sends the DescribeDataset API request.
func (r DescribeDatasetRequest) Send(ctx context.Context) (*DescribeDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatasetResponse{
		DescribeDatasetOutput: r.Request.Data.(*DescribeDatasetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatasetResponse is the response type for the
// DescribeDataset API operation.
type DescribeDatasetResponse struct {
	*DescribeDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataset request.
func (r *DescribeDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
