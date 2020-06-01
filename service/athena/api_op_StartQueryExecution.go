// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartQueryExecutionInput struct {
	_ struct{} `type:"structure"`

	// A unique case-sensitive string used to ensure the request to create the query
	// is idempotent (executes only once). If another StartQueryExecution request
	// is received, the same response is returned and another query is not created.
	// If a parameter has changed, for example, the QueryString, an error is returned.
	//
	// This token is listed as not required because AWS SDKs (for example the AWS
	// SDK for Java) auto-generate the token for users. If you are not using the
	// AWS SDK or the AWS CLI, you must provide this token or the action will fail.
	ClientRequestToken *string `min:"32" type:"string" idempotencyToken:"true"`

	// The database within which the query executes.
	QueryExecutionContext *QueryExecutionContext `type:"structure"`

	// The SQL query statements to be executed.
	//
	// QueryString is a required field
	QueryString *string `min:"1" type:"string" required:"true"`

	// Specifies information about where and how to save the results of the query
	// execution. If the query runs in a workgroup, then workgroup's settings may
	// override query settings. This affects the query results location. The workgroup
	// settings override is specified in EnforceWorkGroupConfiguration (true/false)
	// in the WorkGroupConfiguration. See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
	ResultConfiguration *ResultConfiguration `type:"structure"`

	// The name of the workgroup in which the query is being started.
	WorkGroup *string `type:"string"`
}

// String returns the string representation
func (s StartQueryExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartQueryExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartQueryExecutionInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 32))
	}

	if s.QueryString == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryString"))
	}
	if s.QueryString != nil && len(*s.QueryString) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QueryString", 1))
	}
	if s.QueryExecutionContext != nil {
		if err := s.QueryExecutionContext.Validate(); err != nil {
			invalidParams.AddNested("QueryExecutionContext", err.(aws.ErrInvalidParams))
		}
	}
	if s.ResultConfiguration != nil {
		if err := s.ResultConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ResultConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartQueryExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the query that ran as a result of this request.
	QueryExecutionId *string `type:"string"`
}

// String returns the string representation
func (s StartQueryExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartQueryExecution = "StartQueryExecution"

// StartQueryExecutionRequest returns a request value for making API operation for
// Amazon Athena.
//
// Runs the SQL query statements contained in the Query. Requires you to have
// access to the workgroup in which the query ran.
//
// For code samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
//
//    // Example sending a request using StartQueryExecutionRequest.
//    req := client.StartQueryExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/StartQueryExecution
func (c *Client) StartQueryExecutionRequest(input *StartQueryExecutionInput) StartQueryExecutionRequest {
	op := &aws.Operation{
		Name:       opStartQueryExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartQueryExecutionInput{}
	}

	req := c.newRequest(op, input, &StartQueryExecutionOutput{})

	return StartQueryExecutionRequest{Request: req, Input: input, Copy: c.StartQueryExecutionRequest}
}

// StartQueryExecutionRequest is the request type for the
// StartQueryExecution API operation.
type StartQueryExecutionRequest struct {
	*aws.Request
	Input *StartQueryExecutionInput
	Copy  func(*StartQueryExecutionInput) StartQueryExecutionRequest
}

// Send marshals and sends the StartQueryExecution API request.
func (r StartQueryExecutionRequest) Send(ctx context.Context) (*StartQueryExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartQueryExecutionResponse{
		StartQueryExecutionOutput: r.Request.Data.(*StartQueryExecutionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartQueryExecutionResponse is the response type for the
// StartQueryExecution API operation.
type StartQueryExecutionResponse struct {
	*StartQueryExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartQueryExecution request.
func (r *StartQueryExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
