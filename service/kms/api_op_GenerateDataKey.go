// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GenerateDataKeyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the encryption context that will be used when encrypting the data
	// key.
	//
	// An encryption context is a collection of non-secret key-value pairs that
	// represents additional authenticated data. When you use an encryption context
	// to encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is optional
	// when encrypting with a symmetric CMK, but it is highly recommended.
	//
	// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]string `type:"map"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// Identifies the symmetric CMK that encrypts the data key.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". To specify
	// a CMK in a different AWS account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To
	// get the alias name and alias ARN, use ListAliases.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// Specifies the length of the data key. Use AES_128 to generate a 128-bit symmetric
	// key, or AES_256 to generate a 256-bit symmetric key.
	//
	// You must specify either the KeySpec or the NumberOfBytes parameter (but not
	// both) in every GenerateDataKey request.
	KeySpec DataKeySpec `type:"string" enum:"true"`

	// Specifies the length of the data key in bytes. For example, use the value
	// 64 to generate a 512-bit data key (64 bytes is 512 bits). For 128-bit (16-byte)
	// and 256-bit (32-byte) data keys, use the KeySpec parameter.
	//
	// You must specify either the KeySpec or the NumberOfBytes parameter (but not
	// both) in every GenerateDataKey request.
	NumberOfBytes *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s GenerateDataKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateDataKeyInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.NumberOfBytes != nil && *s.NumberOfBytes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("NumberOfBytes", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GenerateDataKeyOutput struct {
	_ struct{} `type:"structure"`

	// The encrypted copy of the data key. When you use the HTTP API or the AWS
	// CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	CiphertextBlob []byte `min:"1" type:"blob"`

	// The identifier of the CMK that encrypted the data key.
	KeyId *string `min:"1" type:"string"`

	// The plaintext data key. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not Base64-encoded. Use this data key
	// to encrypt your data outside of KMS. Then, remove it from memory as soon
	// as possible.
	//
	// Plaintext is automatically base64 encoded/decoded by the SDK.
	Plaintext []byte `min:"1" type:"blob" sensitive:"true"`
}

// String returns the string representation
func (s GenerateDataKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGenerateDataKey = "GenerateDataKey"

// GenerateDataKeyRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Generates a unique symmetric data key. This operation returns a plaintext
// copy of the data key and a copy that is encrypted under a customer master
// key (CMK) that you specify. You can use the plaintext key to encrypt your
// data outside of AWS KMS and store the encrypted data key with the encrypted
// data.
//
// GenerateDataKey returns a unique data key for each request. The bytes in
// the key are not related to the caller or CMK that is used to encrypt the
// data key.
//
// To generate a data key, specify the symmetric CMK that will be used to encrypt
// the data key. You cannot use an asymmetric CMK to generate data keys. To
// get the type of your CMK, use the DescribeKey operation.
//
// You must also specify the length of the data key. Use either the KeySpec
// or NumberOfBytes parameters (but not both). For 128-bit and 256-bit data
// keys, use the KeySpec parameter.
//
// If the operation succeeds, the plaintext copy of the data key is in the Plaintext
// field of the response, and the encrypted copy of the data key in the CiphertextBlob
// field.
//
// To get only an encrypted copy of the data key, use GenerateDataKeyWithoutPlaintext.
// To generate an asymmetric data key pair, use the GenerateDataKeyPair or GenerateDataKeyPairWithoutPlaintext
// operation. To get a cryptographically secure random byte string, use GenerateRandom.
//
// You can use the optional encryption context to add additional security to
// the encryption operation. If you specify an EncryptionContext, you must specify
// the same encryption context (a case-sensitive exact match) when decrypting
// the encrypted data key. Otherwise, the request to decrypt fails with an InvalidCiphertextException.
// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the AWS Key Management Service Developer Guide.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
// We recommend that you use the following pattern to encrypt data locally in
// your application:
//
// Use the GenerateDataKey operation to get a data encryption key.
//
// Use the plaintext data key (returned in the Plaintext field of the response)
// to encrypt data locally, then erase the plaintext data key from memory.
//
// Store the encrypted data key (returned in the CiphertextBlob field of the
// response) alongside the locally encrypted data.
//
// To decrypt data locally:
//
// Use the Decrypt operation to decrypt the encrypted data key. The operation
// returns a plaintext copy of the data key.
//
// Use the plaintext data key to decrypt data locally, then erase the plaintext
// data key from memory.
//
//    // Example sending a request using GenerateDataKeyRequest.
//    req := client.GenerateDataKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GenerateDataKey
func (c *Client) GenerateDataKeyRequest(input *GenerateDataKeyInput) GenerateDataKeyRequest {
	op := &aws.Operation{
		Name:       opGenerateDataKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataKeyInput{}
	}

	req := c.newRequest(op, input, &GenerateDataKeyOutput{})

	return GenerateDataKeyRequest{Request: req, Input: input, Copy: c.GenerateDataKeyRequest}
}

// GenerateDataKeyRequest is the request type for the
// GenerateDataKey API operation.
type GenerateDataKeyRequest struct {
	*aws.Request
	Input *GenerateDataKeyInput
	Copy  func(*GenerateDataKeyInput) GenerateDataKeyRequest
}

// Send marshals and sends the GenerateDataKey API request.
func (r GenerateDataKeyRequest) Send(ctx context.Context) (*GenerateDataKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateDataKeyResponse{
		GenerateDataKeyOutput: r.Request.Data.(*GenerateDataKeyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateDataKeyResponse is the response type for the
// GenerateDataKey API operation.
type GenerateDataKeyResponse struct {
	*GenerateDataKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateDataKey request.
func (r *GenerateDataKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
