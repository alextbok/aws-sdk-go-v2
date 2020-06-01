// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSolutionVersionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the solution containing the training configuration
	// information.
	//
	// SolutionArn is a required field
	SolutionArn *string `locationName:"solutionArn" type:"string" required:"true"`

	// The scope of training to be performed when creating the solution version.
	// The FULL option trains the solution version based on the entirety of the
	// input solution's training data, while the UPDATE option processes only the
	// data that has changed in comparison to the input solution. Choose UPDATE
	// when you want to incrementally update your solution version instead of creating
	// an entirely new one.
	//
	// The UPDATE option can only be used when you already have an active solution
	// version created from the input solution using the FULL option and the input
	// solution was trained with the native-recipe-hrnn-coldstart recipe.
	TrainingMode TrainingMode `locationName:"trainingMode" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateSolutionVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSolutionVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSolutionVersionInput"}

	if s.SolutionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SolutionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSolutionVersionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the new solution version.
	SolutionVersionArn *string `locationName:"solutionVersionArn" type:"string"`
}

// String returns the string representation
func (s CreateSolutionVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSolutionVersion = "CreateSolutionVersion"

// CreateSolutionVersionRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Trains or retrains an active solution. A solution is created using the CreateSolution
// operation and must be in the ACTIVE state before calling CreateSolutionVersion.
// A new version of the solution is created every time you call this operation.
//
// Status
//
// A solution version can be in one of the following states:
//
//    * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of the version, call DescribeSolutionVersion. Wait until
// the status shows as ACTIVE before calling CreateCampaign.
//
// If the status shows as CREATE FAILED, the response includes a failureReason
// key, which describes why the job failed.
//
// Related APIs
//
//    * ListSolutionVersions
//
//    * DescribeSolutionVersion
//
//    * ListSolutions
//
//    * CreateSolution
//
//    * DescribeSolution
//
//    * DeleteSolution
//
//    // Example sending a request using CreateSolutionVersionRequest.
//    req := client.CreateSolutionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/CreateSolutionVersion
func (c *Client) CreateSolutionVersionRequest(input *CreateSolutionVersionInput) CreateSolutionVersionRequest {
	op := &aws.Operation{
		Name:       opCreateSolutionVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSolutionVersionInput{}
	}

	req := c.newRequest(op, input, &CreateSolutionVersionOutput{})

	return CreateSolutionVersionRequest{Request: req, Input: input, Copy: c.CreateSolutionVersionRequest}
}

// CreateSolutionVersionRequest is the request type for the
// CreateSolutionVersion API operation.
type CreateSolutionVersionRequest struct {
	*aws.Request
	Input *CreateSolutionVersionInput
	Copy  func(*CreateSolutionVersionInput) CreateSolutionVersionRequest
}

// Send marshals and sends the CreateSolutionVersion API request.
func (r CreateSolutionVersionRequest) Send(ctx context.Context) (*CreateSolutionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSolutionVersionResponse{
		CreateSolutionVersionOutput: r.Request.Data.(*CreateSolutionVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSolutionVersionResponse is the response type for the
// CreateSolutionVersion API operation.
type CreateSolutionVersionResponse struct {
	*CreateSolutionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSolutionVersion request.
func (r *CreateSolutionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
