// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLicenseConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// LicenseConfigurationArn is a required field
	LicenseConfigurationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetLicenseConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLicenseConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLicenseConfigurationInput"}

	if s.LicenseConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LicenseConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLicenseConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Automated discovery information.
	AutomatedDiscoveryInformation *AutomatedDiscoveryInformation `type:"structure"`

	// Summaries of the licenses consumed by resources.
	ConsumedLicenseSummaryList []ConsumedLicenseSummary `type:"list"`

	// Number of licenses assigned to resources.
	ConsumedLicenses *int64 `type:"long"`

	// Description of the license configuration.
	Description *string `type:"string"`

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string `type:"string"`

	// Unique ID for the license configuration.
	LicenseConfigurationId *string `type:"string"`

	// Number of available licenses.
	LicenseCount *int64 `type:"long"`

	// Sets the number of available licenses as a hard limit.
	LicenseCountHardLimit *bool `type:"boolean"`

	// Dimension on which the licenses are counted.
	LicenseCountingType LicenseCountingType `type:"string" enum:"true"`

	// License rules.
	LicenseRules []string `type:"list"`

	// Summaries of the managed resources.
	ManagedResourceSummaryList []ManagedResourceSummary `type:"list"`

	// Name of the license configuration.
	Name *string `type:"string"`

	// Account ID of the owner of the license configuration.
	OwnerAccountId *string `type:"string"`

	// Product information.
	ProductInformationList []ProductInformation `type:"list"`

	// License configuration status.
	Status *string `type:"string"`

	// Tags for the license configuration.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s GetLicenseConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLicenseConfiguration = "GetLicenseConfiguration"

// GetLicenseConfigurationRequest returns a request value for making API operation for
// AWS License Manager.
//
// Gets detailed information about the specified license configuration.
//
//    // Example sending a request using GetLicenseConfigurationRequest.
//    req := client.GetLicenseConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/GetLicenseConfiguration
func (c *Client) GetLicenseConfigurationRequest(input *GetLicenseConfigurationInput) GetLicenseConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetLicenseConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLicenseConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetLicenseConfigurationOutput{})

	return GetLicenseConfigurationRequest{Request: req, Input: input, Copy: c.GetLicenseConfigurationRequest}
}

// GetLicenseConfigurationRequest is the request type for the
// GetLicenseConfiguration API operation.
type GetLicenseConfigurationRequest struct {
	*aws.Request
	Input *GetLicenseConfigurationInput
	Copy  func(*GetLicenseConfigurationInput) GetLicenseConfigurationRequest
}

// Send marshals and sends the GetLicenseConfiguration API request.
func (r GetLicenseConfigurationRequest) Send(ctx context.Context) (*GetLicenseConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLicenseConfigurationResponse{
		GetLicenseConfigurationOutput: r.Request.Data.(*GetLicenseConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLicenseConfigurationResponse is the response type for the
// GetLicenseConfiguration API operation.
type GetLicenseConfigurationResponse struct {
	*GetLicenseConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLicenseConfiguration request.
func (r *GetLicenseConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
