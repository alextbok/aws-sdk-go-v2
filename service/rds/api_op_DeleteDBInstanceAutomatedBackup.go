// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Parameter input for the DeleteDBInstanceAutomatedBackup operation.
type DeleteDBInstanceAutomatedBackupInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the source DB instance, which can't be changed and which
	// is unique to an AWS Region.
	//
	// DbiResourceId is a required field
	DbiResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBInstanceAutomatedBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBInstanceAutomatedBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBInstanceAutomatedBackupInput"}

	if s.DbiResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DbiResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBInstanceAutomatedBackupOutput struct {
	_ struct{} `type:"structure"`

	// An automated backup of a DB instance. It it consists of system backups, transaction
	// logs, and the database instance properties that existed at the time you deleted
	// the source instance.
	DBInstanceAutomatedBackup *DBInstanceAutomatedBackup `type:"structure"`
}

// String returns the string representation
func (s DeleteDBInstanceAutomatedBackupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBInstanceAutomatedBackup = "DeleteDBInstanceAutomatedBackup"

// DeleteDBInstanceAutomatedBackupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes automated backups based on the source instance's DbiResourceId value
// or the restorable instance's resource ID.
//
//    // Example sending a request using DeleteDBInstanceAutomatedBackupRequest.
//    req := client.DeleteDBInstanceAutomatedBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBInstanceAutomatedBackup
func (c *Client) DeleteDBInstanceAutomatedBackupRequest(input *DeleteDBInstanceAutomatedBackupInput) DeleteDBInstanceAutomatedBackupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBInstanceAutomatedBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBInstanceAutomatedBackupInput{}
	}

	req := c.newRequest(op, input, &DeleteDBInstanceAutomatedBackupOutput{})

	return DeleteDBInstanceAutomatedBackupRequest{Request: req, Input: input, Copy: c.DeleteDBInstanceAutomatedBackupRequest}
}

// DeleteDBInstanceAutomatedBackupRequest is the request type for the
// DeleteDBInstanceAutomatedBackup API operation.
type DeleteDBInstanceAutomatedBackupRequest struct {
	*aws.Request
	Input *DeleteDBInstanceAutomatedBackupInput
	Copy  func(*DeleteDBInstanceAutomatedBackupInput) DeleteDBInstanceAutomatedBackupRequest
}

// Send marshals and sends the DeleteDBInstanceAutomatedBackup API request.
func (r DeleteDBInstanceAutomatedBackupRequest) Send(ctx context.Context) (*DeleteDBInstanceAutomatedBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBInstanceAutomatedBackupResponse{
		DeleteDBInstanceAutomatedBackupOutput: r.Request.Data.(*DeleteDBInstanceAutomatedBackupOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBInstanceAutomatedBackupResponse is the response type for the
// DeleteDBInstanceAutomatedBackup API operation.
type DeleteDBInstanceAutomatedBackupResponse struct {
	*DeleteDBInstanceAutomatedBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBInstanceAutomatedBackup request.
func (r *DeleteDBInstanceAutomatedBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
