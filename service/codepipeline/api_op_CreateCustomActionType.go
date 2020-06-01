// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateCustomActionType operation.
type CreateCustomActionTypeInput struct {
	_ struct{} `type:"structure"`

	// The category of the custom action, such as a build action or a test action.
	//
	// Although Source and Approval are listed as valid values, they are not currently
	// functional. These values are reserved for future use.
	//
	// Category is a required field
	Category ActionCategory `locationName:"category" type:"string" required:"true" enum:"true"`

	// The configuration properties for the custom action.
	//
	// You can refer to a name in the configuration properties of the custom action
	// within the URL templates by following the format of {Config:name}, as long
	// as the configuration property is both required and not secret. For more information,
	// see Create a Custom Action for a Pipeline (https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html).
	ConfigurationProperties []ActionConfigurationProperty `locationName:"configurationProperties" type:"list"`

	// The details of the input artifact for the action, such as its commit ID.
	//
	// InputArtifactDetails is a required field
	InputArtifactDetails *ArtifactDetails `locationName:"inputArtifactDetails" type:"structure" required:"true"`

	// The details of the output artifact of the action, such as its commit ID.
	//
	// OutputArtifactDetails is a required field
	OutputArtifactDetails *ArtifactDetails `locationName:"outputArtifactDetails" type:"structure" required:"true"`

	// The provider of the service used in the custom action, such as AWS CodeDeploy.
	//
	// Provider is a required field
	Provider *string `locationName:"provider" min:"1" type:"string" required:"true"`

	// URLs that provide users information about this custom action.
	Settings *ActionTypeSettings `locationName:"settings" type:"structure"`

	// The tags for the custom action.
	Tags []Tag `locationName:"tags" type:"list"`

	// The version identifier of the custom action.
	//
	// Version is a required field
	Version *string `locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCustomActionTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCustomActionTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCustomActionTypeInput"}
	if len(s.Category) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Category"))
	}

	if s.InputArtifactDetails == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputArtifactDetails"))
	}

	if s.OutputArtifactDetails == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputArtifactDetails"))
	}

	if s.Provider == nil {
		invalidParams.Add(aws.NewErrParamRequired("Provider"))
	}
	if s.Provider != nil && len(*s.Provider) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Provider", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.ConfigurationProperties != nil {
		for i, v := range s.ConfigurationProperties {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ConfigurationProperties", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.InputArtifactDetails != nil {
		if err := s.InputArtifactDetails.Validate(); err != nil {
			invalidParams.AddNested("InputArtifactDetails", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputArtifactDetails != nil {
		if err := s.OutputArtifactDetails.Validate(); err != nil {
			invalidParams.AddNested("OutputArtifactDetails", err.(aws.ErrInvalidParams))
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			invalidParams.AddNested("Settings", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateCustomActionType operation.
type CreateCustomActionTypeOutput struct {
	_ struct{} `type:"structure"`

	// Returns information about the details of an action type.
	//
	// ActionType is a required field
	ActionType *ActionType `locationName:"actionType" type:"structure" required:"true"`

	// Specifies the tags applied to the custom action.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateCustomActionTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCustomActionType = "CreateCustomActionType"

// CreateCustomActionTypeRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Creates a new custom action that can be used in all pipelines associated
// with the AWS account. Only used for custom actions.
//
//    // Example sending a request using CreateCustomActionTypeRequest.
//    req := client.CreateCustomActionTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/CreateCustomActionType
func (c *Client) CreateCustomActionTypeRequest(input *CreateCustomActionTypeInput) CreateCustomActionTypeRequest {
	op := &aws.Operation{
		Name:       opCreateCustomActionType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCustomActionTypeInput{}
	}

	req := c.newRequest(op, input, &CreateCustomActionTypeOutput{})

	return CreateCustomActionTypeRequest{Request: req, Input: input, Copy: c.CreateCustomActionTypeRequest}
}

// CreateCustomActionTypeRequest is the request type for the
// CreateCustomActionType API operation.
type CreateCustomActionTypeRequest struct {
	*aws.Request
	Input *CreateCustomActionTypeInput
	Copy  func(*CreateCustomActionTypeInput) CreateCustomActionTypeRequest
}

// Send marshals and sends the CreateCustomActionType API request.
func (r CreateCustomActionTypeRequest) Send(ctx context.Context) (*CreateCustomActionTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCustomActionTypeResponse{
		CreateCustomActionTypeOutput: r.Request.Data.(*CreateCustomActionTypeOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCustomActionTypeResponse is the response type for the
// CreateCustomActionType API operation.
type CreateCustomActionTypeResponse struct {
	*CreateCustomActionTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCustomActionType request.
func (r *CreateCustomActionTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
