// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options to upload a part of an archive in a multipart upload operation.
type UploadMultipartPartInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The data to upload.
	Body io.ReadSeeker `locationName:"body" type:"blob"`

	// The SHA256 tree hash of the data being uploaded.
	Checksum *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string"`

	// Identifies the range of bytes in the assembled archive that will be uploaded
	// in this part. Amazon S3 Glacier uses this information to assemble the archive
	// in the proper sequence. The format of this header follows RFC 2616. An example
	// header is Content-Range:bytes 0-4194303/*.
	Range *string `location:"header" locationName:"Content-Range" type:"string"`

	// The upload ID of the multipart upload.
	//
	// UploadId is a required field
	UploadId *string `location:"uri" locationName:"uploadId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s UploadMultipartPartInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadMultipartPartInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadMultipartPartInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
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
func (s UploadMultipartPartInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-sha256-tree-hash", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Range != nil {
		v := *s.Range

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Range", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "uploadId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
type UploadMultipartPartOutput struct {
	_ struct{} `type:"structure"`

	// The SHA256 tree hash that Amazon S3 Glacier computed for the uploaded part.
	Checksum *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string"`
}

// String returns the string representation
func (s UploadMultipartPartOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadMultipartPartOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-sha256-tree-hash", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUploadMultipartPart = "UploadMultipartPart"

// UploadMultipartPartRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation uploads a part of an archive. You can upload archive parts
// in any order. You can also upload them in parallel. You can upload up to
// 10,000 parts for a multipart upload.
//
// Amazon Glacier rejects your upload part request if any of the following conditions
// is true:
//
//    * SHA256 tree hash does not matchTo ensure that part data is not corrupted
//    in transmission, you compute a SHA256 tree hash of the part and include
//    it in your request. Upon receiving the part data, Amazon S3 Glacier also
//    computes a SHA256 tree hash. If these hash values don't match, the operation
//    fails. For information about computing a SHA256 tree hash, see Computing
//    Checksums (https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html).
//
//    * Part size does not matchThe size of each part except the last must match
//    the size specified in the corresponding InitiateMultipartUpload request.
//    The size of the last part must be the same size as, or smaller than, the
//    specified size. If you upload a part whose size is smaller than the part
//    size you specified in your initiate multipart upload request and that
//    part is not the last part, then the upload part request will succeed.
//    However, the subsequent Complete Multipart Upload request will fail.
//
//    * Range does not alignThe byte range value in the request does not align
//    with the part size specified in the corresponding initiate request. For
//    example, if you specify a part size of 4194304 bytes (4 MB), then 0 to
//    4194303 bytes (4 MB - 1) and 4194304 (4 MB) to 8388607 (8 MB - 1) are
//    valid part ranges. However, if you set a range value of 2 MB to 6 MB,
//    the range does not align with the part size and the upload will fail.
//
// This operation is idempotent. If you upload the same part multiple times,
// the data included in the most recent request overwrites the previously uploaded
// data.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Uploading Large Archives
// in Parts (Multipart Upload) (https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html)
// and Upload Part (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-upload-part.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using UploadMultipartPartRequest.
//    req := client.UploadMultipartPartRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UploadMultipartPartRequest(input *UploadMultipartPartInput) UploadMultipartPartRequest {
	op := &aws.Operation{
		Name:       opUploadMultipartPart,
		HTTPMethod: "PUT",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/multipart-uploads/{uploadId}",
	}

	if input == nil {
		input = &UploadMultipartPartInput{}
	}

	req := c.newRequest(op, input, &UploadMultipartPartOutput{})

	return UploadMultipartPartRequest{Request: req, Input: input, Copy: c.UploadMultipartPartRequest}
}

// UploadMultipartPartRequest is the request type for the
// UploadMultipartPart API operation.
type UploadMultipartPartRequest struct {
	*aws.Request
	Input *UploadMultipartPartInput
	Copy  func(*UploadMultipartPartInput) UploadMultipartPartRequest
}

// Send marshals and sends the UploadMultipartPart API request.
func (r UploadMultipartPartRequest) Send(ctx context.Context) (*UploadMultipartPartResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadMultipartPartResponse{
		UploadMultipartPartOutput: r.Request.Data.(*UploadMultipartPartOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadMultipartPartResponse is the response type for the
// UploadMultipartPart API operation.
type UploadMultipartPartResponse struct {
	*UploadMultipartPartOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadMultipartPart request.
func (r *UploadMultipartPartResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
