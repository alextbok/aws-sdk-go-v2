// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartDeliveryStreamEncryptionInput struct {
	_ struct{} `type:"structure"`

	// Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed
	// for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput *DeliveryStreamEncryptionConfigurationInput `type:"structure"`

	// The name of the delivery stream for which you want to enable server-side
	// encryption (SSE).
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartDeliveryStreamEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDeliveryStreamEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDeliveryStreamEncryptionInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}
	if s.DeliveryStreamEncryptionConfigurationInput != nil {
		if err := s.DeliveryStreamEncryptionConfigurationInput.Validate(); err != nil {
			invalidParams.AddNested("DeliveryStreamEncryptionConfigurationInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartDeliveryStreamEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartDeliveryStreamEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartDeliveryStreamEncryption = "StartDeliveryStreamEncryption"

// StartDeliveryStreamEncryptionRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Enables server-side encryption (SSE) for the delivery stream.
//
// This operation is asynchronous. It returns immediately. When you invoke it,
// Kinesis Data Firehose first sets the encryption status of the stream to ENABLING,
// and then to ENABLED. The encryption status of a delivery stream is the Status
// property in DeliveryStreamEncryptionConfiguration. If the operation fails,
// the encryption status changes to ENABLING_FAILED. You can continue to read
// and write data to your delivery stream while the encryption status is ENABLING,
// but the data is not encrypted. It can take up to 5 seconds after the encryption
// status changes to ENABLED before all records written to the delivery stream
// are encrypted. To find out whether a record or a batch of records was encrypted,
// check the response elements PutRecordOutput$Encrypted and PutRecordBatchOutput$Encrypted,
// respectively.
//
// To check the encryption status of a delivery stream, use DescribeDeliveryStream.
//
// Even if encryption is currently enabled for a delivery stream, you can still
// invoke this operation on it to change the ARN of the CMK or both its type
// and ARN. If you invoke this method to change the CMK, and the old CMK is
// of type CUSTOMER_MANAGED_CMK, Kinesis Data Firehose schedules the grant it
// had on the old CMK for retirement. If the new CMK is of type CUSTOMER_MANAGED_CMK,
// Kinesis Data Firehose creates a grant that enables it to use the new CMK
// to encrypt and decrypt data and to manage the grant.
//
// If a delivery stream already has encryption enabled and then you invoke this
// operation to change the ARN of the CMK or both its type and ARN and you get
// ENABLING_FAILED, this only means that the attempt to change the CMK failed.
// In this case, encryption remains enabled with the old CMK.
//
// If the encryption status of your delivery stream is ENABLING_FAILED, you
// can invoke this operation again with a valid CMK. The CMK must be enabled
// and the key policy mustn't explicitly deny the permission for Kinesis Data
// Firehose to invoke KMS encrypt and decrypt operations.
//
// You can enable SSE for a delivery stream only if it's a delivery stream that
// uses DirectPut as its source.
//
// The StartDeliveryStreamEncryption and StopDeliveryStreamEncryption operations
// have a combined limit of 25 calls per delivery stream per 24 hours. For example,
// you reach the limit if you call StartDeliveryStreamEncryption 13 times and
// StopDeliveryStreamEncryption 12 times for the same delivery stream in a 24-hour
// period.
//
//    // Example sending a request using StartDeliveryStreamEncryptionRequest.
//    req := client.StartDeliveryStreamEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/StartDeliveryStreamEncryption
func (c *Client) StartDeliveryStreamEncryptionRequest(input *StartDeliveryStreamEncryptionInput) StartDeliveryStreamEncryptionRequest {
	op := &aws.Operation{
		Name:       opStartDeliveryStreamEncryption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDeliveryStreamEncryptionInput{}
	}

	req := c.newRequest(op, input, &StartDeliveryStreamEncryptionOutput{})

	return StartDeliveryStreamEncryptionRequest{Request: req, Input: input, Copy: c.StartDeliveryStreamEncryptionRequest}
}

// StartDeliveryStreamEncryptionRequest is the request type for the
// StartDeliveryStreamEncryption API operation.
type StartDeliveryStreamEncryptionRequest struct {
	*aws.Request
	Input *StartDeliveryStreamEncryptionInput
	Copy  func(*StartDeliveryStreamEncryptionInput) StartDeliveryStreamEncryptionRequest
}

// Send marshals and sends the StartDeliveryStreamEncryption API request.
func (r StartDeliveryStreamEncryptionRequest) Send(ctx context.Context) (*StartDeliveryStreamEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDeliveryStreamEncryptionResponse{
		StartDeliveryStreamEncryptionOutput: r.Request.Data.(*StartDeliveryStreamEncryptionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDeliveryStreamEncryptionResponse is the response type for the
// StartDeliveryStreamEncryption API operation.
type StartDeliveryStreamEncryptionResponse struct {
	*StartDeliveryStreamEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDeliveryStreamEncryption request.
func (r *StartDeliveryStreamEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
