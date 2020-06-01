// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateTableReplicaAutoScalingInput struct {
	_ struct{} `type:"structure"`

	// Represents the auto scaling settings of the global secondary indexes of the
	// replica to be updated.
	GlobalSecondaryIndexUpdates []GlobalSecondaryIndexAutoScalingUpdate `min:"1" type:"list"`

	// Represents the auto scaling settings to be modified for a global table or
	// global secondary index.
	ProvisionedWriteCapacityAutoScalingUpdate *AutoScalingSettingsUpdate `type:"structure"`

	// Represents the auto scaling settings of replicas of the table that will be
	// modified.
	ReplicaUpdates []ReplicaAutoScalingUpdate `min:"1" type:"list"`

	// The name of the global table to be updated.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTableReplicaAutoScalingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTableReplicaAutoScalingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTableReplicaAutoScalingInput"}
	if s.GlobalSecondaryIndexUpdates != nil && len(s.GlobalSecondaryIndexUpdates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GlobalSecondaryIndexUpdates", 1))
	}
	if s.ReplicaUpdates != nil && len(s.ReplicaUpdates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ReplicaUpdates", 1))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}
	if s.GlobalSecondaryIndexUpdates != nil {
		for i, v := range s.GlobalSecondaryIndexUpdates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GlobalSecondaryIndexUpdates", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ProvisionedWriteCapacityAutoScalingUpdate != nil {
		if err := s.ProvisionedWriteCapacityAutoScalingUpdate.Validate(); err != nil {
			invalidParams.AddNested("ProvisionedWriteCapacityAutoScalingUpdate", err.(aws.ErrInvalidParams))
		}
	}
	if s.ReplicaUpdates != nil {
		for i, v := range s.ReplicaUpdates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ReplicaUpdates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateTableReplicaAutoScalingOutput struct {
	_ struct{} `type:"structure"`

	// Returns information about the auto scaling settings of a table with replicas.
	TableAutoScalingDescription *TableAutoScalingDescription `type:"structure"`
}

// String returns the string representation
func (s UpdateTableReplicaAutoScalingOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTableReplicaAutoScaling = "UpdateTableReplicaAutoScaling"

// UpdateTableReplicaAutoScalingRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Updates auto scaling settings on your global tables at once.
//
// This operation only applies to Version 2019.11.21 (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
// of global tables.
//
//    // Example sending a request using UpdateTableReplicaAutoScalingRequest.
//    req := client.UpdateTableReplicaAutoScalingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/UpdateTableReplicaAutoScaling
func (c *Client) UpdateTableReplicaAutoScalingRequest(input *UpdateTableReplicaAutoScalingInput) UpdateTableReplicaAutoScalingRequest {
	op := &aws.Operation{
		Name:       opUpdateTableReplicaAutoScaling,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTableReplicaAutoScalingInput{}
	}

	req := c.newRequest(op, input, &UpdateTableReplicaAutoScalingOutput{})

	return UpdateTableReplicaAutoScalingRequest{Request: req, Input: input, Copy: c.UpdateTableReplicaAutoScalingRequest}
}

// UpdateTableReplicaAutoScalingRequest is the request type for the
// UpdateTableReplicaAutoScaling API operation.
type UpdateTableReplicaAutoScalingRequest struct {
	*aws.Request
	Input *UpdateTableReplicaAutoScalingInput
	Copy  func(*UpdateTableReplicaAutoScalingInput) UpdateTableReplicaAutoScalingRequest
}

// Send marshals and sends the UpdateTableReplicaAutoScaling API request.
func (r UpdateTableReplicaAutoScalingRequest) Send(ctx context.Context) (*UpdateTableReplicaAutoScalingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTableReplicaAutoScalingResponse{
		UpdateTableReplicaAutoScalingOutput: r.Request.Data.(*UpdateTableReplicaAutoScalingOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTableReplicaAutoScalingResponse is the response type for the
// UpdateTableReplicaAutoScaling API operation.
type UpdateTableReplicaAutoScalingResponse struct {
	*UpdateTableReplicaAutoScalingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTableReplicaAutoScaling request.
func (r *UpdateTableReplicaAutoScalingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
