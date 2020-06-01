// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateProjectVersionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 location to store the results of training.
	//
	// OutputConfig is a required field
	OutputConfig *OutputConfig `type:"structure" required:"true"`

	// The ARN of the Amazon Rekognition Custom Labels project that manages the
	// model that you want to train.
	//
	// ProjectArn is a required field
	ProjectArn *string `min:"20" type:"string" required:"true"`

	// The dataset to use for testing.
	//
	// TestingData is a required field
	TestingData *TestingData `type:"structure" required:"true"`

	// The dataset to use for training.
	//
	// TrainingData is a required field
	TrainingData *TrainingData `type:"structure" required:"true"`

	// A name for the version of the model. This value must be unique.
	//
	// VersionName is a required field
	VersionName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProjectVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProjectVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProjectVersionInput"}

	if s.OutputConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputConfig"))
	}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 20))
	}

	if s.TestingData == nil {
		invalidParams.Add(aws.NewErrParamRequired("TestingData"))
	}

	if s.TrainingData == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrainingData"))
	}

	if s.VersionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionName"))
	}
	if s.VersionName != nil && len(*s.VersionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionName", 1))
	}
	if s.OutputConfig != nil {
		if err := s.OutputConfig.Validate(); err != nil {
			invalidParams.AddNested("OutputConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.TestingData != nil {
		if err := s.TestingData.Validate(); err != nil {
			invalidParams.AddNested("TestingData", err.(aws.ErrInvalidParams))
		}
	}
	if s.TrainingData != nil {
		if err := s.TrainingData.Validate(); err != nil {
			invalidParams.AddNested("TrainingData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateProjectVersionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the model version that was created. Use DescribeProjectVersion
	// to get the current status of the training operation.
	ProjectVersionArn *string `min:"20" type:"string"`
}

// String returns the string representation
func (s CreateProjectVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProjectVersion = "CreateProjectVersion"

// CreateProjectVersionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Creates a new version of a model and begins training. Models are managed
// as part of an Amazon Rekognition Custom Labels project. You can specify one
// training dataset and one testing dataset. The response from CreateProjectVersion
// is an Amazon Resource Name (ARN) for the version of the model.
//
// Training takes a while to complete. You can get the current status by calling
// DescribeProjectVersions.
//
// Once training has successfully completed, call DescribeProjectVersions to
// get the training results and evaluate the model.
//
// After evaluating the model, you start the model by calling StartProjectVersion.
//
// This operation requires permissions to perform the rekognition:CreateProjectVersion
// action.
//
//    // Example sending a request using CreateProjectVersionRequest.
//    req := client.CreateProjectVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateProjectVersionRequest(input *CreateProjectVersionInput) CreateProjectVersionRequest {
	op := &aws.Operation{
		Name:       opCreateProjectVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProjectVersionInput{}
	}

	req := c.newRequest(op, input, &CreateProjectVersionOutput{})

	return CreateProjectVersionRequest{Request: req, Input: input, Copy: c.CreateProjectVersionRequest}
}

// CreateProjectVersionRequest is the request type for the
// CreateProjectVersion API operation.
type CreateProjectVersionRequest struct {
	*aws.Request
	Input *CreateProjectVersionInput
	Copy  func(*CreateProjectVersionInput) CreateProjectVersionRequest
}

// Send marshals and sends the CreateProjectVersion API request.
func (r CreateProjectVersionRequest) Send(ctx context.Context) (*CreateProjectVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectVersionResponse{
		CreateProjectVersionOutput: r.Request.Data.(*CreateProjectVersionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectVersionResponse is the response type for the
// CreateProjectVersion API operation.
type CreateProjectVersionResponse struct {
	*CreateProjectVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProjectVersion request.
func (r *CreateProjectVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
