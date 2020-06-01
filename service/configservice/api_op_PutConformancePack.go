// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutConformancePackInput struct {
	_ struct{} `type:"structure"`

	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []ConformancePackInputParameter `type:"list"`

	// Name of the conformance pack you want to create.
	//
	// ConformancePackName is a required field
	ConformancePackName *string `min:"1" type:"string" required:"true"`

	// AWS Config stores intermediate files while processing conformance pack template.
	//
	// DeliveryS3Bucket is a required field
	DeliveryS3Bucket *string `min:"3" type:"string" required:"true"`

	// The prefix for the Amazon S3 bucket.
	DeliveryS3KeyPrefix *string `min:"1" type:"string"`

	// A string containing full conformance pack template body. Structure containing
	// the template body with a minimum length of 1 byte and a maximum length of
	// 51,200 bytes.
	//
	// You can only use a YAML template with one resource type, that is, config
	// rule and a remediation action.
	TemplateBody *string `min:"1" type:"string"`

	// Location of file containing the template body (s3://bucketname/prefix). The
	// uri must point to the conformance pack template (max size: 300 KB) that is
	// located in an Amazon S3 bucket in the same region as the conformance pack.
	//
	// You must have access to read Amazon S3 bucket.
	TemplateS3Uri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutConformancePackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConformancePackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConformancePackInput"}

	if s.ConformancePackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConformancePackName"))
	}
	if s.ConformancePackName != nil && len(*s.ConformancePackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConformancePackName", 1))
	}

	if s.DeliveryS3Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryS3Bucket"))
	}
	if s.DeliveryS3Bucket != nil && len(*s.DeliveryS3Bucket) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryS3Bucket", 3))
	}
	if s.DeliveryS3KeyPrefix != nil && len(*s.DeliveryS3KeyPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryS3KeyPrefix", 1))
	}
	if s.TemplateBody != nil && len(*s.TemplateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateBody", 1))
	}
	if s.TemplateS3Uri != nil && len(*s.TemplateS3Uri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateS3Uri", 1))
	}
	if s.ConformancePackInputParameters != nil {
		for i, v := range s.ConformancePackInputParameters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ConformancePackInputParameters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutConformancePackOutput struct {
	_ struct{} `type:"structure"`

	// ARN of the conformance pack.
	ConformancePackArn *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutConformancePackOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutConformancePack = "PutConformancePack"

// PutConformancePackRequest returns a request value for making API operation for
// AWS Config.
//
// Creates or updates a conformance pack. A conformance pack is a collection
// of AWS Config rules that can be easily deployed in an account and a region
// and across AWS Organization.
//
// This API creates a service linked role AWSServiceRoleForConfigConforms in
// your account. The service linked role is created only when the role does
// not exist in your account. AWS Config verifies the existence of role with
// GetRole action.
//
// You must specify either the TemplateS3Uri or the TemplateBody parameter,
// but not both. If you provide both AWS Config uses the TemplateS3Uri parameter
// and ignores the TemplateBody parameter.
//
//    // Example sending a request using PutConformancePackRequest.
//    req := client.PutConformancePackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutConformancePack
func (c *Client) PutConformancePackRequest(input *PutConformancePackInput) PutConformancePackRequest {
	op := &aws.Operation{
		Name:       opPutConformancePack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutConformancePackInput{}
	}

	req := c.newRequest(op, input, &PutConformancePackOutput{})

	return PutConformancePackRequest{Request: req, Input: input, Copy: c.PutConformancePackRequest}
}

// PutConformancePackRequest is the request type for the
// PutConformancePack API operation.
type PutConformancePackRequest struct {
	*aws.Request
	Input *PutConformancePackInput
	Copy  func(*PutConformancePackInput) PutConformancePackRequest
}

// Send marshals and sends the PutConformancePack API request.
func (r PutConformancePackRequest) Send(ctx context.Context) (*PutConformancePackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConformancePackResponse{
		PutConformancePackOutput: r.Request.Data.(*PutConformancePackOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConformancePackResponse is the response type for the
// PutConformancePack API operation.
type PutConformancePackResponse struct {
	*PutConformancePackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConformancePack request.
func (r *PutConformancePackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
