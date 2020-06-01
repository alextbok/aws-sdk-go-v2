// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResetDBClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster parameter group to reset.
	//
	// DBClusterParameterGroupName is a required field
	DBClusterParameterGroupName *string `type:"string" required:"true"`

	// A list of parameter names in the DB cluster parameter group to reset to the
	// default values. You can't use this parameter if the ResetAllParameters parameter
	// is enabled.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`

	// A value that indicates whether to reset all parameters in the DB cluster
	// parameter group to their default values. You can't use this parameter if
	// there is a list of parameter names specified for the Parameters parameter.
	ResetAllParameters *bool `type:"boolean"`
}

// String returns the string representation
func (s ResetDBClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetDBClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetDBClusterParameterGroupInput"}

	if s.DBClusterParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResetDBClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters or numbers.
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// This value is stored as a lowercase string.
	DBClusterParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ResetDBClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetDBClusterParameterGroup = "ResetDBClusterParameterGroup"

// ResetDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Modifies the parameters of a DB cluster parameter group to the default value.
// To reset specific parameters submit a list of the following: ParameterName
// and ApplyMethod. To reset the entire DB cluster parameter group, specify
// the DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When resetting the entire group, dynamic parameters are updated immediately
// and static parameters are set to pending-reboot to take effect on the next
// DB instance restart or RebootDBInstance request. You must call RebootDBInstance
// for every DB instance in your DB cluster that you want the updated static
// parameter to apply to.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using ResetDBClusterParameterGroupRequest.
//    req := client.ResetDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ResetDBClusterParameterGroup
func (c *Client) ResetDBClusterParameterGroupRequest(input *ResetDBClusterParameterGroupInput) ResetDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opResetDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &ResetDBClusterParameterGroupOutput{})

	return ResetDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.ResetDBClusterParameterGroupRequest}
}

// ResetDBClusterParameterGroupRequest is the request type for the
// ResetDBClusterParameterGroup API operation.
type ResetDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *ResetDBClusterParameterGroupInput
	Copy  func(*ResetDBClusterParameterGroupInput) ResetDBClusterParameterGroupRequest
}

// Send marshals and sends the ResetDBClusterParameterGroup API request.
func (r ResetDBClusterParameterGroupRequest) Send(ctx context.Context) (*ResetDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetDBClusterParameterGroupResponse{
		ResetDBClusterParameterGroupOutput: r.Request.Data.(*ResetDBClusterParameterGroupOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetDBClusterParameterGroupResponse is the response type for the
// ResetDBClusterParameterGroup API operation.
type ResetDBClusterParameterGroupResponse struct {
	*ResetDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetDBClusterParameterGroup request.
func (r *ResetDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
