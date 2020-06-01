// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to describe one or more environments.
type DescribeEnvironmentsInput struct {
	_ struct{} `type:"structure"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that are associated with this application.
	ApplicationName *string `min:"1" type:"string"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that have the specified IDs.
	EnvironmentIds []string `type:"list"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that have the specified names.
	EnvironmentNames []string `type:"list"`

	// Indicates whether to include deleted environments:
	//
	// true: Environments that have been deleted after IncludedDeletedBackTo are
	// displayed.
	//
	// false: Do not include deleted environments.
	IncludeDeleted *bool `type:"boolean"`

	// If specified when IncludeDeleted is set to true, then environments deleted
	// after this date are displayed.
	IncludedDeletedBackTo *time.Time `type:"timestamp"`

	// For a paginated request. Specify a maximum number of environments to include
	// in each response.
	//
	// If no MaxRecords is specified, all available environments are retrieved in
	// a single response.
	MaxRecords *int64 `min:"1" type:"integer"`

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical
	// to the ones specified in the initial request.
	//
	// If no NextToken is specified, the first page is retrieved.
	NextToken *string `type:"string"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that are associated with this application version.
	VersionLabel *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeEnvironmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEnvironmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEnvironmentsInput"}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.MaxRecords != nil && *s.MaxRecords < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 1))
	}
	if s.VersionLabel != nil && len(*s.VersionLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionLabel", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result message containing a list of environment descriptions.
type DescribeEnvironmentsOutput struct {
	_ struct{} `type:"structure"`

	// Returns an EnvironmentDescription list.
	Environments []EnvironmentDescription `type:"list"`

	// In a paginated request, the token that you can pass in a subsequent request
	// to get the next response page.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEnvironmentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEnvironments = "DescribeEnvironments"

// DescribeEnvironmentsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Returns descriptions for existing environments.
//
//    // Example sending a request using DescribeEnvironmentsRequest.
//    req := client.DescribeEnvironmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeEnvironments
func (c *Client) DescribeEnvironmentsRequest(input *DescribeEnvironmentsInput) DescribeEnvironmentsRequest {
	op := &aws.Operation{
		Name:       opDescribeEnvironments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEnvironmentsInput{}
	}

	req := c.newRequest(op, input, &DescribeEnvironmentsOutput{})

	return DescribeEnvironmentsRequest{Request: req, Input: input, Copy: c.DescribeEnvironmentsRequest}
}

// DescribeEnvironmentsRequest is the request type for the
// DescribeEnvironments API operation.
type DescribeEnvironmentsRequest struct {
	*aws.Request
	Input *DescribeEnvironmentsInput
	Copy  func(*DescribeEnvironmentsInput) DescribeEnvironmentsRequest
}

// Send marshals and sends the DescribeEnvironments API request.
func (r DescribeEnvironmentsRequest) Send(ctx context.Context) (*DescribeEnvironmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEnvironmentsResponse{
		DescribeEnvironmentsOutput: r.Request.Data.(*DescribeEnvironmentsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEnvironmentsResponse is the response type for the
// DescribeEnvironments API operation.
type DescribeEnvironmentsResponse struct {
	*DescribeEnvironmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEnvironments request.
func (r *DescribeEnvironmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
