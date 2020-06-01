// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDocumentClassifierInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the request. If you don't set the client request
	// token, Amazon Comprehend generates one.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Management (IAM) role
	// that grants Amazon Comprehend read access to your input data.
	//
	// DataAccessRoleArn is a required field
	DataAccessRoleArn *string `min:"20" type:"string" required:"true"`

	// The name of the document classifier.
	//
	// DocumentClassifierName is a required field
	DocumentClassifierName *string `type:"string" required:"true"`

	// Specifies the format and location of the input data for the job.
	//
	// InputDataConfig is a required field
	InputDataConfig *DocumentClassifierInputDataConfig `type:"structure" required:"true"`

	// The language of the input documents. You can specify any of the following
	// languages supported by Amazon Comprehend: German ("de"), English ("en"),
	// Spanish ("es"), French ("fr"), Italian ("it"), or Portuguese ("pt"). All
	// documents must be in the same language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// Indicates the mode in which the classifier will be trained. The classifier
	// can be trained in multi-class mode, which identifies one and only one class
	// for each document, or multi-label mode, which identifies one or more labels
	// for each document. In multi-label mode, multiple labels for an individual
	// document are separated by a delimiter. The default delimiter between labels
	// is a pipe (|).
	Mode DocumentClassifierMode `type:"string" enum:"true"`

	// Enables the addition of output results configuration parameters for custom
	// classifier jobs.
	OutputDataConfig *DocumentClassifierOutputDataConfig `type:"structure"`

	// Tags to be associated with the document classifier being created. A tag is
	// a key-value pair that adds as a metadata to a resource used by Amazon Comprehend.
	// For example, a tag with "Sales" as the key might be added to a resource to
	// indicate its use by the sales department.
	Tags []Tag `type:"list"`

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses
	// to encrypt data on the storage volume attached to the ML compute instance(s)
	// that process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	//    * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	//    * Amazon Resource Name (ARN) of a KMS Key: "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string `type:"string"`

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your custom classifier. For more
	// information, see Amazon VPC (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s CreateDocumentClassifierInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDocumentClassifierInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDocumentClassifierInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DataAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataAccessRoleArn"))
	}
	if s.DataAccessRoleArn != nil && len(*s.DataAccessRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("DataAccessRoleArn", 20))
	}

	if s.DocumentClassifierName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentClassifierName"))
	}

	if s.InputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDataConfig"))
	}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}
	if s.InputDataConfig != nil {
		if err := s.InputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			invalidParams.AddNested("VpcConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDocumentClassifierOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the document classifier.
	DocumentClassifierArn *string `type:"string"`
}

// String returns the string representation
func (s CreateDocumentClassifierOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDocumentClassifier = "CreateDocumentClassifier"

// CreateDocumentClassifierRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Creates a new document classifier that you can use to categorize documents.
// To create a classifier you provide a set of training documents that labeled
// with the categories that you want to use. After the classifier is trained
// you can use it to categorize a set of labeled documents into the categories.
// For more information, see how-document-classification.
//
//    // Example sending a request using CreateDocumentClassifierRequest.
//    req := client.CreateDocumentClassifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/CreateDocumentClassifier
func (c *Client) CreateDocumentClassifierRequest(input *CreateDocumentClassifierInput) CreateDocumentClassifierRequest {
	op := &aws.Operation{
		Name:       opCreateDocumentClassifier,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDocumentClassifierInput{}
	}

	req := c.newRequest(op, input, &CreateDocumentClassifierOutput{})

	return CreateDocumentClassifierRequest{Request: req, Input: input, Copy: c.CreateDocumentClassifierRequest}
}

// CreateDocumentClassifierRequest is the request type for the
// CreateDocumentClassifier API operation.
type CreateDocumentClassifierRequest struct {
	*aws.Request
	Input *CreateDocumentClassifierInput
	Copy  func(*CreateDocumentClassifierInput) CreateDocumentClassifierRequest
}

// Send marshals and sends the CreateDocumentClassifier API request.
func (r CreateDocumentClassifierRequest) Send(ctx context.Context) (*CreateDocumentClassifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDocumentClassifierResponse{
		CreateDocumentClassifierOutput: r.Request.Data.(*CreateDocumentClassifierOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDocumentClassifierResponse is the response type for the
// CreateDocumentClassifier API operation.
type CreateDocumentClassifierResponse struct {
	*CreateDocumentClassifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDocumentClassifier request.
func (r *CreateDocumentClassifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
