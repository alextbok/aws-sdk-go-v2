// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeModelPackageInput struct {
	_ struct{} `type:"structure"`

	// The name of the model package to describe.
	//
	// ModelPackageName is a required field
	ModelPackageName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeModelPackageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeModelPackageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeModelPackageInput"}

	if s.ModelPackageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelPackageName"))
	}
	if s.ModelPackageName != nil && len(*s.ModelPackageName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ModelPackageName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeModelPackageOutput struct {
	_ struct{} `type:"structure"`

	// Whether the model package is certified for listing on AWS Marketplace.
	CertifyForMarketplace *bool `type:"boolean"`

	// A timestamp specifying when the model package was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// Details about inference jobs that can be run with models based on this model
	// package.
	InferenceSpecification *InferenceSpecification `type:"structure"`

	// The Amazon Resource Name (ARN) of the model package.
	//
	// ModelPackageArn is a required field
	ModelPackageArn *string `min:"1" type:"string" required:"true"`

	// A brief summary of the model package.
	ModelPackageDescription *string `type:"string"`

	// The name of the model package being described.
	//
	// ModelPackageName is a required field
	ModelPackageName *string `min:"1" type:"string" required:"true"`

	// The current status of the model package.
	//
	// ModelPackageStatus is a required field
	ModelPackageStatus ModelPackageStatus `type:"string" required:"true" enum:"true"`

	// Details about the current status of the model package.
	//
	// ModelPackageStatusDetails is a required field
	ModelPackageStatusDetails *ModelPackageStatusDetails `type:"structure" required:"true"`

	// Details about the algorithm that was used to create the model package.
	SourceAlgorithmSpecification *SourceAlgorithmSpecification `type:"structure"`

	// Configurations for one or more transform jobs that Amazon SageMaker runs
	// to test the model package.
	ValidationSpecification *ModelPackageValidationSpecification `type:"structure"`
}

// String returns the string representation
func (s DescribeModelPackageOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeModelPackage = "DescribeModelPackage"

// DescribeModelPackageRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns a description of the specified model package, which is used to create
// Amazon SageMaker models or list them on AWS Marketplace.
//
// To create models in Amazon SageMaker, buyers can subscribe to model packages
// listed on AWS Marketplace.
//
//    // Example sending a request using DescribeModelPackageRequest.
//    req := client.DescribeModelPackageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeModelPackage
func (c *Client) DescribeModelPackageRequest(input *DescribeModelPackageInput) DescribeModelPackageRequest {
	op := &aws.Operation{
		Name:       opDescribeModelPackage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeModelPackageInput{}
	}

	req := c.newRequest(op, input, &DescribeModelPackageOutput{})

	return DescribeModelPackageRequest{Request: req, Input: input, Copy: c.DescribeModelPackageRequest}
}

// DescribeModelPackageRequest is the request type for the
// DescribeModelPackage API operation.
type DescribeModelPackageRequest struct {
	*aws.Request
	Input *DescribeModelPackageInput
	Copy  func(*DescribeModelPackageInput) DescribeModelPackageRequest
}

// Send marshals and sends the DescribeModelPackage API request.
func (r DescribeModelPackageRequest) Send(ctx context.Context) (*DescribeModelPackageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeModelPackageResponse{
		DescribeModelPackageOutput: r.Request.Data.(*DescribeModelPackageOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeModelPackageResponse is the response type for the
// DescribeModelPackage API operation.
type DescribeModelPackageResponse struct {
	*DescribeModelPackageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeModelPackage request.
func (r *DescribeModelPackageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
