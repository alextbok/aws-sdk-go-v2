// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateBackupSelectionInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies the backup plan to be associated with the selection of
	// resources.
	//
	// BackupPlanId is a required field
	BackupPlanId *string `location:"uri" locationName:"backupPlanId" type:"string" required:"true"`

	// Specifies the body of a request to assign a set of resources to a backup
	// plan.
	//
	// BackupSelection is a required field
	BackupSelection *BackupSelection `type:"structure" required:"true"`

	// A unique string that identifies the request and allows failed requests to
	// be retried without the risk of executing the operation twice.
	CreatorRequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateBackupSelectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBackupSelectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBackupSelectionInput"}

	if s.BackupPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanId"))
	}

	if s.BackupSelection == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupSelection"))
	}
	if s.BackupSelection != nil {
		if err := s.BackupSelection.Validate(); err != nil {
			invalidParams.AddNested("BackupSelection", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBackupSelectionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupSelection != nil {
		v := s.BackupSelection

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "BackupSelection", v, metadata)
	}
	if s.CreatorRequestId != nil {
		v := *s.CreatorRequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatorRequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateBackupSelectionOutput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	BackupPlanId *string `type:"string"`

	// The date and time a backup selection is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// Uniquely identifies the body of a request to assign a set of resources to
	// a backup plan.
	SelectionId *string `type:"string"`
}

// String returns the string representation
func (s CreateBackupSelectionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBackupSelectionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.SelectionId != nil {
		v := *s.SelectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SelectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateBackupSelection = "CreateBackupSelection"

// CreateBackupSelectionRequest returns a request value for making API operation for
// AWS Backup.
//
// Creates a JSON document that specifies a set of resources to assign to a
// backup plan. Resources can be included by specifying patterns for a ListOfTags
// and selected Resources.
//
// For example, consider the following patterns:
//
//    * Resources: "arn:aws:ec2:region:account-id:volume/volume-id"
//
//    * ConditionKey:"department" ConditionValue:"finance" ConditionType:"STRINGEQUALS"
//
//    * ConditionKey:"importance" ConditionValue:"critical" ConditionType:"STRINGEQUALS"
//
// Using these patterns would back up all Amazon Elastic Block Store (Amazon
// EBS) volumes that are tagged as "department=finance", "importance=critical",
// in addition to an EBS volume with the specified volume Id.
//
// Resources and conditions are additive in that all resources that match the
// pattern are selected. This shouldn't be confused with a logical AND, where
// all conditions must match. The matching patterns are logically 'put together
// using the OR operator. In other words, all patterns that match are selected
// for backup.
//
//    // Example sending a request using CreateBackupSelectionRequest.
//    req := client.CreateBackupSelectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/CreateBackupSelection
func (c *Client) CreateBackupSelectionRequest(input *CreateBackupSelectionInput) CreateBackupSelectionRequest {
	op := &aws.Operation{
		Name:       opCreateBackupSelection,
		HTTPMethod: "PUT",
		HTTPPath:   "/backup/plans/{backupPlanId}/selections/",
	}

	if input == nil {
		input = &CreateBackupSelectionInput{}
	}

	req := c.newRequest(op, input, &CreateBackupSelectionOutput{})

	return CreateBackupSelectionRequest{Request: req, Input: input, Copy: c.CreateBackupSelectionRequest}
}

// CreateBackupSelectionRequest is the request type for the
// CreateBackupSelection API operation.
type CreateBackupSelectionRequest struct {
	*aws.Request
	Input *CreateBackupSelectionInput
	Copy  func(*CreateBackupSelectionInput) CreateBackupSelectionRequest
}

// Send marshals and sends the CreateBackupSelection API request.
func (r CreateBackupSelectionRequest) Send(ctx context.Context) (*CreateBackupSelectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBackupSelectionResponse{
		CreateBackupSelectionOutput: r.Request.Data.(*CreateBackupSelectionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBackupSelectionResponse is the response type for the
// CreateBackupSelection API operation.
type CreateBackupSelectionResponse struct {
	*CreateBackupSelectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBackupSelection request.
func (r *CreateBackupSelectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
