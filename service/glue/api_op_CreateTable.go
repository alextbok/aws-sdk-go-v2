// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTableInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which to create the Table. If none is supplied,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The catalog database in which to create the new table. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The TableInput object that defines the metadata table to create in the catalog.
	//
	// TableInput is a required field
	TableInput *TableInput `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTableInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.TableInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableInput"))
	}
	if s.TableInput != nil {
		if err := s.TableInput.Validate(); err != nil {
			invalidParams.AddNested("TableInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTableOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTable = "CreateTable"

// CreateTableRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new table definition in the Data Catalog.
//
//    // Example sending a request using CreateTableRequest.
//    req := client.CreateTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateTable
func (c *Client) CreateTableRequest(input *CreateTableInput) CreateTableRequest {
	op := &aws.Operation{
		Name:       opCreateTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTableInput{}
	}

	req := c.newRequest(op, input, &CreateTableOutput{})

	return CreateTableRequest{Request: req, Input: input, Copy: c.CreateTableRequest}
}

// CreateTableRequest is the request type for the
// CreateTable API operation.
type CreateTableRequest struct {
	*aws.Request
	Input *CreateTableInput
	Copy  func(*CreateTableInput) CreateTableRequest
}

// Send marshals and sends the CreateTable API request.
func (r CreateTableRequest) Send(ctx context.Context) (*CreateTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTableResponse{
		CreateTableOutput: r.Request.Data.(*CreateTableOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTableResponse is the response type for the
// CreateTable API operation.
type CreateTableResponse struct {
	*CreateTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTable request.
func (r *CreateTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
