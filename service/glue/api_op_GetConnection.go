// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetConnectionRequest
type GetConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which the connection resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// Allows you to retrieve the connection metadata without returning the password.
	// For instance, the AWS Glue console uses this flag to retrieve the connection,
	// and does not display the password. Set this parameter when the caller might
	// not have permission to use the AWS KMS key to decrypt the password, but it
	// does have permission to access the rest of the connection properties.
	HidePassword *bool `type:"boolean"`

	// The name of the connection definition to retrieve.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConnectionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetConnectionResponse
type GetConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The requested connection definition.
	Connection *Connection `type:"structure"`
}

// String returns the string representation
func (s GetConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConnection = "GetConnection"

// GetConnectionRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a connection definition from the Data Catalog.
//
//    // Example sending a request using GetConnectionRequest.
//    req := client.GetConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetConnection
func (c *Client) GetConnectionRequest(input *GetConnectionInput) GetConnectionRequest {
	op := &aws.Operation{
		Name:       opGetConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConnectionInput{}
	}

	req := c.newRequest(op, input, &GetConnectionOutput{})
	return GetConnectionRequest{Request: req, Input: input, Copy: c.GetConnectionRequest}
}

// GetConnectionRequest is the request type for the
// GetConnection API operation.
type GetConnectionRequest struct {
	*aws.Request
	Input *GetConnectionInput
	Copy  func(*GetConnectionInput) GetConnectionRequest
}

// Send marshals and sends the GetConnection API request.
func (r GetConnectionRequest) Send(ctx context.Context) (*GetConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectionResponse{
		GetConnectionOutput: r.Request.Data.(*GetConnectionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConnectionResponse is the response type for the
// GetConnection API operation.
type GetConnectionResponse struct {
	*GetConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnection request.
func (r *GetConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
