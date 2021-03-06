// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options to add an archive to a vault.
type UploadArchiveInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The optional description of the archive you are uploading.
	ArchiveDescription *string `location:"header" locationName:"x-amz-archive-description" type:"string"`

	// The data to upload.
	Body io.ReadSeeker `locationName:"body" type:"blob"`

	// The SHA256 tree hash of the data being uploaded.
	Checksum *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s UploadArchiveInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadArchiveInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadArchiveInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadArchiveInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ArchiveDescription != nil {
		v := *s.ArchiveDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-archive-description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-sha256-tree-hash", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "body", protocol.ReadSeekerStream{V: v}, metadata)
	}
	return nil
}

// Contains the Amazon S3 Glacier response to your request.
//
// For information about the underlying REST API, see Upload Archive (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html).
// For conceptual information, see Working with Archives in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html).
type UploadArchiveOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the archive. This value is also included as part of the location.
	ArchiveId *string `location:"header" locationName:"x-amz-archive-id" type:"string"`

	// The checksum of the archive computed by Amazon S3 Glacier.
	Checksum *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string"`

	// The relative URI path of the newly added archive resource.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s UploadArchiveOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadArchiveOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ArchiveId != nil {
		v := *s.ArchiveId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-archive-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-sha256-tree-hash", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUploadArchive = "UploadArchive"

// UploadArchiveRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation adds an archive to a vault. This is a synchronous operation,
// and for a successful upload, your data is durably persisted. Amazon S3 Glacier
// returns the archive ID in the x-amz-archive-id header of the response.
//
// You must use the archive ID to access your data in Amazon S3 Glacier. After
// you upload an archive, you should save the archive ID returned so that you
// can retrieve or delete the archive later. Besides saving the archive ID,
// you can also index it and give it a friendly name to allow for better searching.
// You can also use the optional archive description field to specify how the
// archive is referred to in an external index of archives, such as you might
// create in Amazon DynamoDB. You can also get the vault inventory to obtain
// a list of archive IDs in a vault. For more information, see InitiateJob.
//
// You must provide a SHA256 tree hash of the data you are uploading. For information
// about computing a SHA256 tree hash, see Computing Checksums (https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html).
//
// You can optionally specify an archive description of up to 1,024 printable
// ASCII characters. You can get the archive description when you either retrieve
// the archive or get the vault inventory. For more information, see InitiateJob.
// Amazon Glacier does not interpret the description in any way. An archive
// description does not need to be unique. You cannot use the description to
// retrieve or sort the archive list.
//
// Archives are immutable. After you upload an archive, you cannot edit the
// archive or its description.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Uploading an Archive
// in Amazon Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-an-archive.html)
// and Upload Archive (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using UploadArchiveRequest.
//    req := client.UploadArchiveRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UploadArchiveRequest(input *UploadArchiveInput) UploadArchiveRequest {
	op := &aws.Operation{
		Name:       opUploadArchive,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/archives",
	}

	if input == nil {
		input = &UploadArchiveInput{}
	}

	req := c.newRequest(op, input, &UploadArchiveOutput{})
	return UploadArchiveRequest{Request: req, Input: input, Copy: c.UploadArchiveRequest}
}

// UploadArchiveRequest is the request type for the
// UploadArchive API operation.
type UploadArchiveRequest struct {
	*aws.Request
	Input *UploadArchiveInput
	Copy  func(*UploadArchiveInput) UploadArchiveRequest
}

// Send marshals and sends the UploadArchive API request.
func (r UploadArchiveRequest) Send(ctx context.Context) (*UploadArchiveResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadArchiveResponse{
		UploadArchiveOutput: r.Request.Data.(*UploadArchiveOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadArchiveResponse is the response type for the
// UploadArchive API operation.
type UploadArchiveResponse struct {
	*UploadArchiveOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadArchive request.
func (r *UploadArchiveResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
