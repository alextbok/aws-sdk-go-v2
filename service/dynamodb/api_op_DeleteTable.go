// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DeleteTable operation.
type DeleteTableInput struct {
	_ struct{} `type:"structure"`

	// The name of the table to delete.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTableInput"}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DeleteTable operation.
type DeleteTableOutput struct {
	_ struct{} `type:"structure"`

	// Represents the properties of a table.
	TableDescription *TableDescription `type:"structure"`
}

// String returns the string representation
func (s DeleteTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTable = "DeleteTable"

// DeleteTableRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// The DeleteTable operation deletes a table and all of its items. After a DeleteTable
// request, the specified table is in the DELETING state until DynamoDB completes
// the deletion. If the table is in the ACTIVE state, you can delete it. If
// a table is in CREATING or UPDATING states, then DynamoDB returns a ResourceInUseException.
// If the specified table does not exist, DynamoDB returns a ResourceNotFoundException.
// If table is already in the DELETING state, no error is returned.
//
// DynamoDB might continue to accept data read and write operations, such as
// GetItem and PutItem, on a table in the DELETING state until the table deletion
// is complete.
//
// When you delete a table, any indexes on that table are also deleted.
//
// If you have DynamoDB Streams enabled on the table, then the corresponding
// stream on that table goes into the DISABLED state, and the stream is automatically
// deleted after 24 hours.
//
// Use the DescribeTable action to check the status of the table.
//
//    // Example sending a request using DeleteTableRequest.
//    req := client.DeleteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DeleteTable
func (c *Client) DeleteTableRequest(input *DeleteTableInput) DeleteTableRequest {
	op := &aws.Operation{
		Name:       opDeleteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTableInput{}
	}

	req := c.newRequest(op, input, &DeleteTableOutput{})

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}

	return DeleteTableRequest{Request: req, Input: input, Copy: c.DeleteTableRequest}
}

// DeleteTableRequest is the request type for the
// DeleteTable API operation.
type DeleteTableRequest struct {
	*aws.Request
	Input *DeleteTableInput
	Copy  func(*DeleteTableInput) DeleteTableRequest
}

// Send marshals and sends the DeleteTable API request.
func (r DeleteTableRequest) Send(ctx context.Context) (*DeleteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTableResponse{
		DeleteTableOutput: r.Request.Data.(*DeleteTableOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTableResponse is the response type for the
// DeleteTable API operation.
type DeleteTableResponse struct {
	*DeleteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTable request.
func (r *DeleteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
