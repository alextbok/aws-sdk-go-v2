// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The DB parameter group family name. A DB parameter group can be associated
	// with one and only one DB parameter group family, and can be applied only
	// to a DB instance running a database engine and engine version compatible
	// with that DB parameter group family.
	//
	// To list all of the available parameter group families, use the following
	// command:
	//
	// aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"
	//
	// The output contains duplicates.
	//
	// DBParameterGroupFamily is a required field
	DBParameterGroupFamily *string `type:"string" required:"true"`

	// The name of the DB parameter group.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// This value is stored as a lowercase string.
	//
	// DBParameterGroupName is a required field
	DBParameterGroupName *string `type:"string" required:"true"`

	// The description for the DB parameter group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// Tags to assign to the DB parameter group.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateDBParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBParameterGroupInput"}

	if s.DBParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupFamily"))
	}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB parameter group.
	//
	// This data type is used as a response element in the DescribeDBParameterGroups
	// action.
	DBParameterGroup *DBParameterGroup `type:"structure"`
}

// String returns the string representation
func (s CreateDBParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBParameterGroup = "CreateDBParameterGroup"

// CreateDBParameterGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new DB parameter group.
//
// A DB parameter group is initially created with the default parameters for
// the database engine used by the DB instance. To provide custom values for
// any of the parameters, you must modify the group after creating it using
// ModifyDBParameterGroup. Once you've created a DB parameter group, you need
// to associate it with your DB instance using ModifyDBInstance. When you associate
// a new DB parameter group with a running DB instance, you need to reboot the
// DB instance without failover for the new DB parameter group and associated
// settings to take effect.
//
// After you create a DB parameter group, you should wait at least 5 minutes
// before creating your first DB instance that uses that DB parameter group
// as the default parameter group. This allows Amazon RDS to fully complete
// the create action before the parameter group is used as the default for a
// new DB instance. This is especially important for parameters that are critical
// when creating the default database for a DB instance, such as the character
// set for the default database defined by the character_set_database parameter.
// You can use the Parameter Groups option of the Amazon RDS console (https://console.aws.amazon.com/rds/)
// or the DescribeDBParameters command to verify that your DB parameter group
// has been created or modified.
//
//    // Example sending a request using CreateDBParameterGroupRequest.
//    req := client.CreateDBParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBParameterGroup
func (c *Client) CreateDBParameterGroupRequest(input *CreateDBParameterGroupInput) CreateDBParameterGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDBParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBParameterGroupInput{}
	}

	req := c.newRequest(op, input, &CreateDBParameterGroupOutput{})

	return CreateDBParameterGroupRequest{Request: req, Input: input, Copy: c.CreateDBParameterGroupRequest}
}

// CreateDBParameterGroupRequest is the request type for the
// CreateDBParameterGroup API operation.
type CreateDBParameterGroupRequest struct {
	*aws.Request
	Input *CreateDBParameterGroupInput
	Copy  func(*CreateDBParameterGroupInput) CreateDBParameterGroupRequest
}

// Send marshals and sends the CreateDBParameterGroup API request.
func (r CreateDBParameterGroupRequest) Send(ctx context.Context) (*CreateDBParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBParameterGroupResponse{
		CreateDBParameterGroupOutput: r.Request.Data.(*CreateDBParameterGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBParameterGroupResponse is the response type for the
// CreateDBParameterGroup API operation.
type CreateDBParameterGroupResponse struct {
	*CreateDBParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBParameterGroup request.
func (r *CreateDBParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
