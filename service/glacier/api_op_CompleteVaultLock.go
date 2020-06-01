// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input values for CompleteVaultLock.
type CompleteVaultLockInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You
	// can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you specify your account ID, do
	// not include any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The lockId value is the lock ID obtained from a InitiateVaultLock request.
	//
	// LockId is a required field
	LockId *string `location:"uri" locationName:"lockId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s CompleteVaultLockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CompleteVaultLockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CompleteVaultLockInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.LockId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LockId"))
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
func (s CompleteVaultLockInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LockId != nil {
		v := *s.LockId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "lockId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CompleteVaultLockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CompleteVaultLockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CompleteVaultLockOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCompleteVaultLock = "CompleteVaultLock"

// CompleteVaultLockRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation completes the vault locking process by transitioning the vault
// lock from the InProgress state to the Locked state, which causes the vault
// lock policy to become unchangeable. A vault lock is put into the InProgress
// state by calling InitiateVaultLock. You can obtain the state of the vault
// lock by calling GetVaultLock. For more information about the vault locking
// process, Amazon Glacier Vault Lock (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html).
//
// This operation is idempotent. This request is always successful if the vault
// lock is in the Locked state and the provided lock ID matches the lock ID
// originally used to lock the vault.
//
// If an invalid lock ID is passed in the request when the vault lock is in
// the Locked state, the operation returns an AccessDeniedException error. If
// an invalid lock ID is passed in the request when the vault lock is in the
// InProgress state, the operation throws an InvalidParameter error.
//
//    // Example sending a request using CompleteVaultLockRequest.
//    req := client.CompleteVaultLockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CompleteVaultLockRequest(input *CompleteVaultLockInput) CompleteVaultLockRequest {
	op := &aws.Operation{
		Name:       opCompleteVaultLock,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/lock-policy/{lockId}",
	}

	if input == nil {
		input = &CompleteVaultLockInput{}
	}

	req := c.newRequest(op, input, &CompleteVaultLockOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return CompleteVaultLockRequest{Request: req, Input: input, Copy: c.CompleteVaultLockRequest}
}

// CompleteVaultLockRequest is the request type for the
// CompleteVaultLock API operation.
type CompleteVaultLockRequest struct {
	*aws.Request
	Input *CompleteVaultLockInput
	Copy  func(*CompleteVaultLockInput) CompleteVaultLockRequest
}

// Send marshals and sends the CompleteVaultLock API request.
func (r CompleteVaultLockRequest) Send(ctx context.Context) (*CompleteVaultLockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CompleteVaultLockResponse{
		CompleteVaultLockOutput: r.Request.Data.(*CompleteVaultLockOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CompleteVaultLockResponse is the response type for the
// CompleteVaultLock API operation.
type CompleteVaultLockResponse struct {
	*CompleteVaultLockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CompleteVaultLock request.
func (r *CompleteVaultLockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
