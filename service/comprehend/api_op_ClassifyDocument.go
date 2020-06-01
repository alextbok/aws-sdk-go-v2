// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ClassifyDocumentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Number (ARN) of the endpoint.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`

	// The document text to be analyzed.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ClassifyDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ClassifyDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ClassifyDocumentInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ClassifyDocumentOutput struct {
	_ struct{} `type:"structure"`

	// The classes used by the document being analyzed. These are used for multi-class
	// trained models. Individual classes are mutually exclusive and each document
	// is expected to have only a single class assigned to it. For example, an animal
	// can be a dog or a cat, but not both at the same time.
	Classes []DocumentClass `type:"list"`

	// The labels used the document being analyzed. These are used for multi-label
	// trained models. Individual labels represent different categories that are
	// related in some manner and are not multually exclusive. For example, a movie
	// can be just an action movie, or it can be an action movie, a science fiction
	// movie, and a comedy, all at the same time.
	Labels []DocumentLabel `type:"list"`
}

// String returns the string representation
func (s ClassifyDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

const opClassifyDocument = "ClassifyDocument"

// ClassifyDocumentRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Creates a new document classification request to analyze a single document
// in real-time, using a previously created and trained custom model and an
// endpoint.
//
//    // Example sending a request using ClassifyDocumentRequest.
//    req := client.ClassifyDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ClassifyDocument
func (c *Client) ClassifyDocumentRequest(input *ClassifyDocumentInput) ClassifyDocumentRequest {
	op := &aws.Operation{
		Name:       opClassifyDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ClassifyDocumentInput{}
	}

	req := c.newRequest(op, input, &ClassifyDocumentOutput{})

	return ClassifyDocumentRequest{Request: req, Input: input, Copy: c.ClassifyDocumentRequest}
}

// ClassifyDocumentRequest is the request type for the
// ClassifyDocument API operation.
type ClassifyDocumentRequest struct {
	*aws.Request
	Input *ClassifyDocumentInput
	Copy  func(*ClassifyDocumentInput) ClassifyDocumentRequest
}

// Send marshals and sends the ClassifyDocument API request.
func (r ClassifyDocumentRequest) Send(ctx context.Context) (*ClassifyDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ClassifyDocumentResponse{
		ClassifyDocumentOutput: r.Request.Data.(*ClassifyDocumentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ClassifyDocumentResponse is the response type for the
// ClassifyDocument API operation.
type ClassifyDocumentResponse struct {
	*ClassifyDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ClassifyDocument request.
func (r *ClassifyDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
