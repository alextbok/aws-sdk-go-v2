// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartRelationalDatabaseInput struct {
	_ struct{} `type:"structure"`

	// The name of your database to start.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`
}

// String returns the string representation
func (s StartRelationalDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartRelationalDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartRelationalDatabaseInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartRelationalDatabaseOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s StartRelationalDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartRelationalDatabase = "StartRelationalDatabase"

// StartRelationalDatabaseRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Starts a specific database from a stopped state in Amazon Lightsail. To restart
// a database, use the reboot relational database operation.
//
// The start relational database operation supports tag-based access control
// via resource tags applied to the resource identified by relationalDatabaseName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using StartRelationalDatabaseRequest.
//    req := client.StartRelationalDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/StartRelationalDatabase
func (c *Client) StartRelationalDatabaseRequest(input *StartRelationalDatabaseInput) StartRelationalDatabaseRequest {
	op := &aws.Operation{
		Name:       opStartRelationalDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartRelationalDatabaseInput{}
	}

	req := c.newRequest(op, input, &StartRelationalDatabaseOutput{})

	return StartRelationalDatabaseRequest{Request: req, Input: input, Copy: c.StartRelationalDatabaseRequest}
}

// StartRelationalDatabaseRequest is the request type for the
// StartRelationalDatabase API operation.
type StartRelationalDatabaseRequest struct {
	*aws.Request
	Input *StartRelationalDatabaseInput
	Copy  func(*StartRelationalDatabaseInput) StartRelationalDatabaseRequest
}

// Send marshals and sends the StartRelationalDatabase API request.
func (r StartRelationalDatabaseRequest) Send(ctx context.Context) (*StartRelationalDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartRelationalDatabaseResponse{
		StartRelationalDatabaseOutput: r.Request.Data.(*StartRelationalDatabaseOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartRelationalDatabaseResponse is the response type for the
// StartRelationalDatabase API operation.
type StartRelationalDatabaseResponse struct {
	*StartRelationalDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartRelationalDatabase request.
func (r *StartRelationalDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
