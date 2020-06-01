// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the DataSource during creation.
	//
	// DataSourceId is a required field
	DataSourceId *string `min:"1" type:"string" required:"true"`

	// A new user-supplied name or description of the DataSource that will replace
	// the current description.
	//
	// DataSourceName is a required field
	DataSourceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDataSourceInput"}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}
	if s.DataSourceId != nil && len(*s.DataSourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataSourceId", 1))
	}

	if s.DataSourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an UpdateDataSource operation.
//
// You can see the updated content by using the GetBatchPrediction operation.
type UpdateDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the DataSource during creation. This value should be identical
	// to the value of the DataSourceID in the request.
	DataSourceId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDataSource = "UpdateDataSource"

// UpdateDataSourceRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Updates the DataSourceName of a DataSource.
//
// You can use the GetDataSource operation to view the contents of the updated
// data element.
//
//    // Example sending a request using UpdateDataSourceRequest.
//    req := client.UpdateDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateDataSourceRequest(input *UpdateDataSourceInput) UpdateDataSourceRequest {
	op := &aws.Operation{
		Name:       opUpdateDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDataSourceInput{}
	}

	req := c.newRequest(op, input, &UpdateDataSourceOutput{})

	return UpdateDataSourceRequest{Request: req, Input: input, Copy: c.UpdateDataSourceRequest}
}

// UpdateDataSourceRequest is the request type for the
// UpdateDataSource API operation.
type UpdateDataSourceRequest struct {
	*aws.Request
	Input *UpdateDataSourceInput
	Copy  func(*UpdateDataSourceInput) UpdateDataSourceRequest
}

// Send marshals and sends the UpdateDataSource API request.
func (r UpdateDataSourceRequest) Send(ctx context.Context) (*UpdateDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDataSourceResponse{
		UpdateDataSourceOutput: r.Request.Data.(*UpdateDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDataSourceResponse is the response type for the
// UpdateDataSource API operation.
type UpdateDataSourceResponse struct {
	*UpdateDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataSource request.
func (r *UpdateDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
