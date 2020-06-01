// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDeliveryStreamInput struct {
	_ struct{} `type:"structure"`

	// Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed
	// for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput *DeliveryStreamEncryptionConfigurationInput `type:"structure"`

	// The name of the delivery stream. This name must be unique per AWS account
	// in the same AWS Region. If the delivery streams are in different accounts
	// or different Regions, you can have multiple delivery streams with the same
	// name.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`

	// The delivery stream type. This parameter can be one of the following values:
	//
	//    * DirectPut: Provider applications access the delivery stream directly.
	//
	//    * KinesisStreamAsSource: The delivery stream uses a Kinesis data stream
	//    as a source.
	DeliveryStreamType DeliveryStreamType `type:"string" enum:"true"`

	// The destination in Amazon ES. You can specify only one destination.
	ElasticsearchDestinationConfiguration *ElasticsearchDestinationConfiguration `type:"structure"`

	// The destination in Amazon S3. You can specify only one destination.
	ExtendedS3DestinationConfiguration *ExtendedS3DestinationConfiguration `type:"structure"`

	// When a Kinesis data stream is used as the source for the delivery stream,
	// a KinesisStreamSourceConfiguration containing the Kinesis data stream Amazon
	// Resource Name (ARN) and the role ARN for the source stream.
	KinesisStreamSourceConfiguration *KinesisStreamSourceConfiguration `type:"structure"`

	// The destination in Amazon Redshift. You can specify only one destination.
	RedshiftDestinationConfiguration *RedshiftDestinationConfiguration `type:"structure"`

	// [Deprecated] The destination in Amazon S3. You can specify only one destination.
	S3DestinationConfiguration *S3DestinationConfiguration `deprecated:"true" type:"structure"`

	// The destination in Splunk. You can specify only one destination.
	SplunkDestinationConfiguration *SplunkDestinationConfiguration `type:"structure"`

	// A set of tags to assign to the delivery stream. A tag is a key-value pair
	// that you can define and assign to AWS resources. Tags are metadata. For example,
	// you can add friendly names and descriptions or other types of information
	// that can help you distinguish the delivery stream. For more information about
	// tags, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
	// in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a delivery stream.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateDeliveryStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeliveryStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeliveryStreamInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.DeliveryStreamEncryptionConfigurationInput != nil {
		if err := s.DeliveryStreamEncryptionConfigurationInput.Validate(); err != nil {
			invalidParams.AddNested("DeliveryStreamEncryptionConfigurationInput", err.(aws.ErrInvalidParams))
		}
	}
	if s.ElasticsearchDestinationConfiguration != nil {
		if err := s.ElasticsearchDestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ElasticsearchDestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.ExtendedS3DestinationConfiguration != nil {
		if err := s.ExtendedS3DestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ExtendedS3DestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.KinesisStreamSourceConfiguration != nil {
		if err := s.KinesisStreamSourceConfiguration.Validate(); err != nil {
			invalidParams.AddNested("KinesisStreamSourceConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.RedshiftDestinationConfiguration != nil {
		if err := s.RedshiftDestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RedshiftDestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.S3DestinationConfiguration != nil {
		if err := s.S3DestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("S3DestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.SplunkDestinationConfiguration != nil {
		if err := s.SplunkDestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SplunkDestinationConfiguration", err.(aws.ErrInvalidParams))
		}
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

type CreateDeliveryStreamOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the delivery stream.
	DeliveryStreamARN *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDeliveryStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDeliveryStream = "CreateDeliveryStream"

// CreateDeliveryStreamRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Creates a Kinesis Data Firehose delivery stream.
//
// By default, you can create up to 50 delivery streams per AWS Region.
//
// This is an asynchronous operation that immediately returns. The initial status
// of the delivery stream is CREATING. After the delivery stream is created,
// its status is ACTIVE and it now accepts data. If the delivery stream creation
// fails, the status transitions to CREATING_FAILED. Attempts to send data to
// a delivery stream that is not in the ACTIVE state cause an exception. To
// check the state of a delivery stream, use DescribeDeliveryStream.
//
// If the status of a delivery stream is CREATING_FAILED, this status doesn't
// change, and you can't invoke CreateDeliveryStream again on it. However, you
// can invoke the DeleteDeliveryStream operation to delete it.
//
// A Kinesis Data Firehose delivery stream can be configured to receive records
// directly from providers using PutRecord or PutRecordBatch, or it can be configured
// to use an existing Kinesis stream as its source. To specify a Kinesis data
// stream as input, set the DeliveryStreamType parameter to KinesisStreamAsSource,
// and provide the Kinesis stream Amazon Resource Name (ARN) and role ARN in
// the KinesisStreamSourceConfiguration parameter.
//
// To create a delivery stream with server-side encryption (SSE) enabled, include
// DeliveryStreamEncryptionConfigurationInput in your request. This is optional.
// You can also invoke StartDeliveryStreamEncryption to turn on SSE for an existing
// delivery stream that doesn't have SSE enabled.
//
// A delivery stream is configured with a single destination: Amazon S3, Amazon
// ES, Amazon Redshift, or Splunk. You must specify only one of the following
// destination configuration parameters: ExtendedS3DestinationConfiguration,
// S3DestinationConfiguration, ElasticsearchDestinationConfiguration, RedshiftDestinationConfiguration,
// or SplunkDestinationConfiguration.
//
// When you specify S3DestinationConfiguration, you can also provide the following
// optional values: BufferingHints, EncryptionConfiguration, and CompressionFormat.
// By default, if no BufferingHints value is provided, Kinesis Data Firehose
// buffers data up to 5 MB or for 5 minutes, whichever condition is satisfied
// first. BufferingHints is a hint, so there are some cases where the service
// cannot adhere to these conditions strictly. For example, record boundaries
// might be such that the size is a little over or under the configured buffering
// size. By default, no encryption is performed. We strongly recommend that
// you enable encryption to ensure secure data storage in Amazon S3.
//
// A few notes about Amazon Redshift as a destination:
//
//    * An Amazon Redshift destination requires an S3 bucket as intermediate
//    location. Kinesis Data Firehose first delivers data to Amazon S3 and then
//    uses COPY syntax to load data into an Amazon Redshift table. This is specified
//    in the RedshiftDestinationConfiguration.S3Configuration parameter.
//
//    * The compression formats SNAPPY or ZIP cannot be specified in RedshiftDestinationConfiguration.S3Configuration
//    because the Amazon Redshift COPY operation that reads from the S3 bucket
//    doesn't support these compression formats.
//
//    * We strongly recommend that you use the user name and password you provide
//    exclusively with Kinesis Data Firehose, and that the permissions for the
//    account are restricted for Amazon Redshift INSERT permissions.
//
// Kinesis Data Firehose assumes the IAM role that is configured as part of
// the destination. The role should allow the Kinesis Data Firehose principal
// to assume the role, and the role should have permissions that allow the service
// to deliver the data. For more information, see Grant Kinesis Data Firehose
// Access to an Amazon S3 Destination (https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-s3)
// in the Amazon Kinesis Data Firehose Developer Guide.
//
//    // Example sending a request using CreateDeliveryStreamRequest.
//    req := client.CreateDeliveryStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/CreateDeliveryStream
func (c *Client) CreateDeliveryStreamRequest(input *CreateDeliveryStreamInput) CreateDeliveryStreamRequest {
	op := &aws.Operation{
		Name:       opCreateDeliveryStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDeliveryStreamInput{}
	}

	req := c.newRequest(op, input, &CreateDeliveryStreamOutput{})

	return CreateDeliveryStreamRequest{Request: req, Input: input, Copy: c.CreateDeliveryStreamRequest}
}

// CreateDeliveryStreamRequest is the request type for the
// CreateDeliveryStream API operation.
type CreateDeliveryStreamRequest struct {
	*aws.Request
	Input *CreateDeliveryStreamInput
	Copy  func(*CreateDeliveryStreamInput) CreateDeliveryStreamRequest
}

// Send marshals and sends the CreateDeliveryStream API request.
func (r CreateDeliveryStreamRequest) Send(ctx context.Context) (*CreateDeliveryStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeliveryStreamResponse{
		CreateDeliveryStreamOutput: r.Request.Data.(*CreateDeliveryStreamOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeliveryStreamResponse is the response type for the
// CreateDeliveryStream API operation.
type CreateDeliveryStreamResponse struct {
	*CreateDeliveryStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeliveryStream request.
func (r *CreateDeliveryStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
