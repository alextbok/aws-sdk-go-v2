// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDestinationInput struct {
	_ struct{} `type:"structure"`

	// Obtain this value from the VersionId result of DeliveryStreamDescription.
	// This value is required, and helps the service perform conditional operations.
	// For example, if there is an interleaving update and this value is null, then
	// the update destination fails. After the update is successful, the VersionId
	// value is updated. The service then performs a merge of the old configuration
	// with the new configuration.
	//
	// CurrentDeliveryStreamVersionId is a required field
	CurrentDeliveryStreamVersionId *string `min:"1" type:"string" required:"true"`

	// The name of the delivery stream.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`

	// The ID of the destination.
	//
	// DestinationId is a required field
	DestinationId *string `min:"1" type:"string" required:"true"`

	// Describes an update for a destination in Amazon ES.
	ElasticsearchDestinationUpdate *ElasticsearchDestinationUpdate `type:"structure"`

	// Describes an update for a destination in Amazon S3.
	ExtendedS3DestinationUpdate *ExtendedS3DestinationUpdate `type:"structure"`

	// Describes an update for a destination in Amazon Redshift.
	RedshiftDestinationUpdate *RedshiftDestinationUpdate `type:"structure"`

	// [Deprecated] Describes an update for a destination in Amazon S3.
	S3DestinationUpdate *S3DestinationUpdate `deprecated:"true" type:"structure"`

	// Describes an update for a destination in Splunk.
	SplunkDestinationUpdate *SplunkDestinationUpdate `type:"structure"`
}

// String returns the string representation
func (s UpdateDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDestinationInput"}

	if s.CurrentDeliveryStreamVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentDeliveryStreamVersionId"))
	}
	if s.CurrentDeliveryStreamVersionId != nil && len(*s.CurrentDeliveryStreamVersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentDeliveryStreamVersionId", 1))
	}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}

	if s.DestinationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationId"))
	}
	if s.DestinationId != nil && len(*s.DestinationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationId", 1))
	}
	if s.ElasticsearchDestinationUpdate != nil {
		if err := s.ElasticsearchDestinationUpdate.Validate(); err != nil {
			invalidParams.AddNested("ElasticsearchDestinationUpdate", err.(aws.ErrInvalidParams))
		}
	}
	if s.ExtendedS3DestinationUpdate != nil {
		if err := s.ExtendedS3DestinationUpdate.Validate(); err != nil {
			invalidParams.AddNested("ExtendedS3DestinationUpdate", err.(aws.ErrInvalidParams))
		}
	}
	if s.RedshiftDestinationUpdate != nil {
		if err := s.RedshiftDestinationUpdate.Validate(); err != nil {
			invalidParams.AddNested("RedshiftDestinationUpdate", err.(aws.ErrInvalidParams))
		}
	}
	if s.S3DestinationUpdate != nil {
		if err := s.S3DestinationUpdate.Validate(); err != nil {
			invalidParams.AddNested("S3DestinationUpdate", err.(aws.ErrInvalidParams))
		}
	}
	if s.SplunkDestinationUpdate != nil {
		if err := s.SplunkDestinationUpdate.Validate(); err != nil {
			invalidParams.AddNested("SplunkDestinationUpdate", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDestination = "UpdateDestination"

// UpdateDestinationRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Updates the specified destination of the specified delivery stream.
//
// Use this operation to change the destination type (for example, to replace
// the Amazon S3 destination with Amazon Redshift) or change the parameters
// associated with a destination (for example, to change the bucket name of
// the Amazon S3 destination). The update might not occur immediately. The target
// delivery stream remains active while the configurations are updated, so data
// writes to the delivery stream can continue during this process. The updated
// configurations are usually effective within a few minutes.
//
// Switching between Amazon ES and other services is not supported. For an Amazon
// ES destination, you can only update to another Amazon ES destination.
//
// If the destination type is the same, Kinesis Data Firehose merges the configuration
// parameters specified with the destination configuration that already exists
// on the delivery stream. If any of the parameters are not specified in the
// call, the existing values are retained. For example, in the Amazon S3 destination,
// if EncryptionConfiguration is not specified, then the existing EncryptionConfiguration
// is maintained on the destination.
//
// If the destination type is not the same, for example, changing the destination
// from Amazon S3 to Amazon Redshift, Kinesis Data Firehose does not merge any
// parameters. In this case, all parameters must be specified.
//
// Kinesis Data Firehose uses CurrentDeliveryStreamVersionId to avoid race conditions
// and conflicting merges. This is a required field, and the service updates
// the configuration only if the existing configuration has a version ID that
// matches. After the update is applied successfully, the version ID is updated,
// and can be retrieved using DescribeDeliveryStream. Use the new version ID
// to set CurrentDeliveryStreamVersionId in the next call.
//
//    // Example sending a request using UpdateDestinationRequest.
//    req := client.UpdateDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/UpdateDestination
func (c *Client) UpdateDestinationRequest(input *UpdateDestinationInput) UpdateDestinationRequest {
	op := &aws.Operation{
		Name:       opUpdateDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDestinationInput{}
	}

	req := c.newRequest(op, input, &UpdateDestinationOutput{})

	return UpdateDestinationRequest{Request: req, Input: input, Copy: c.UpdateDestinationRequest}
}

// UpdateDestinationRequest is the request type for the
// UpdateDestination API operation.
type UpdateDestinationRequest struct {
	*aws.Request
	Input *UpdateDestinationInput
	Copy  func(*UpdateDestinationInput) UpdateDestinationRequest
}

// Send marshals and sends the UpdateDestination API request.
func (r UpdateDestinationRequest) Send(ctx context.Context) (*UpdateDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDestinationResponse{
		UpdateDestinationOutput: r.Request.Data.(*UpdateDestinationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDestinationResponse is the response type for the
// UpdateDestination API operation.
type UpdateDestinationResponse struct {
	*UpdateDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDestination request.
func (r *UpdateDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
