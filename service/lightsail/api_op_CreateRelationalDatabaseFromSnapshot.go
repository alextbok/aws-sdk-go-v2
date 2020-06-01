// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRelationalDatabaseFromSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone in which to create your new database. Use the us-east-2a
	// case-sensitive format.
	//
	// You can get a list of Availability Zones by using the get regions operation.
	// Be sure to add the include relational database Availability Zones parameter
	// to your request.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// Specifies the accessibility options for your new database. A value of true
	// specifies a database that is available to resources outside of your Lightsail
	// account. A value of false specifies a database that is available only to
	// your Lightsail resources in the same region as your database.
	PubliclyAccessible *bool `locationName:"publiclyAccessible" type:"boolean"`

	// The bundle ID for your new database. A bundle describes the performance specifications
	// for your database.
	//
	// You can get a list of database bundle IDs by using the get relational database
	// bundles operation.
	//
	// When creating a new database from a snapshot, you cannot choose a bundle
	// that is smaller than the bundle of the source database.
	RelationalDatabaseBundleId *string `locationName:"relationalDatabaseBundleId" type:"string"`

	// The name to use for your new database.
	//
	// Constraints:
	//
	//    * Must contain from 2 to 255 alphanumeric characters, or hyphens.
	//
	//    * The first and last character must be a letter or number.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// The name of the database snapshot from which to create your new database.
	RelationalDatabaseSnapshotName *string `locationName:"relationalDatabaseSnapshotName" type:"string"`

	// The date and time to restore your database from.
	//
	// Constraints:
	//
	//    * Must be before the latest restorable time for the database.
	//
	//    * Cannot be specified if the use latest restorable time parameter is true.
	//
	//    * Specified in Coordinated Universal Time (UTC).
	//
	//    * Specified in the Unix time format. For example, if you wish to use a
	//    restore time of October 1, 2018, at 8 PM UTC, then you input 1538424000
	//    as the restore time.
	RestoreTime *time.Time `locationName:"restoreTime" type:"timestamp"`

	// The name of the source database.
	SourceRelationalDatabaseName *string `locationName:"sourceRelationalDatabaseName" type:"string"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`

	// Specifies whether your database is restored from the latest backup time.
	// A value of true restores from the latest backup time.
	//
	// Default: false
	//
	// Constraints: Cannot be specified if the restore time parameter is provided.
	UseLatestRestorableTime *bool `locationName:"useLatestRestorableTime" type:"boolean"`
}

// String returns the string representation
func (s CreateRelationalDatabaseFromSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRelationalDatabaseFromSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRelationalDatabaseFromSnapshotInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRelationalDatabaseFromSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CreateRelationalDatabaseFromSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRelationalDatabaseFromSnapshot = "CreateRelationalDatabaseFromSnapshot"

// CreateRelationalDatabaseFromSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates a new database from an existing database snapshot in Amazon Lightsail.
//
// You can create a new database from a snapshot in if something goes wrong
// with your original database, or to change it to a different plan, such as
// a high availability or standard plan.
//
// The create relational database from snapshot operation supports tag-based
// access control via request tags and resource tags applied to the resource
// identified by relationalDatabaseSnapshotName. For more information, see the
// Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateRelationalDatabaseFromSnapshotRequest.
//    req := client.CreateRelationalDatabaseFromSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateRelationalDatabaseFromSnapshot
func (c *Client) CreateRelationalDatabaseFromSnapshotRequest(input *CreateRelationalDatabaseFromSnapshotInput) CreateRelationalDatabaseFromSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateRelationalDatabaseFromSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRelationalDatabaseFromSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateRelationalDatabaseFromSnapshotOutput{})

	return CreateRelationalDatabaseFromSnapshotRequest{Request: req, Input: input, Copy: c.CreateRelationalDatabaseFromSnapshotRequest}
}

// CreateRelationalDatabaseFromSnapshotRequest is the request type for the
// CreateRelationalDatabaseFromSnapshot API operation.
type CreateRelationalDatabaseFromSnapshotRequest struct {
	*aws.Request
	Input *CreateRelationalDatabaseFromSnapshotInput
	Copy  func(*CreateRelationalDatabaseFromSnapshotInput) CreateRelationalDatabaseFromSnapshotRequest
}

// Send marshals and sends the CreateRelationalDatabaseFromSnapshot API request.
func (r CreateRelationalDatabaseFromSnapshotRequest) Send(ctx context.Context) (*CreateRelationalDatabaseFromSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRelationalDatabaseFromSnapshotResponse{
		CreateRelationalDatabaseFromSnapshotOutput: r.Request.Data.(*CreateRelationalDatabaseFromSnapshotOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRelationalDatabaseFromSnapshotResponse is the response type for the
// CreateRelationalDatabaseFromSnapshot API operation.
type CreateRelationalDatabaseFromSnapshotResponse struct {
	*CreateRelationalDatabaseFromSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRelationalDatabaseFromSnapshot request.
func (r *CreateRelationalDatabaseFromSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
