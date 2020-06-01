// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to create an XssMatchSet.
type CreateXssMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description for the XssMatchSet that you're creating.
	// You can't change Name after you create the XssMatchSet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateXssMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateXssMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateXssMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response to a CreateXssMatchSet request.
type CreateXssMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateXssMatchSet request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// An XssMatchSet.
	XssMatchSet *XssMatchSet `type:"structure"`
}

// String returns the string representation
func (s CreateXssMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateXssMatchSet = "CreateXssMatchSet"

// CreateXssMatchSetRequest returns a request value for making API operation for
// AWS WAF.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Creates an XssMatchSet, which you use to allow, block, or count requests
// that contain cross-site scripting attacks in the specified part of web requests.
// AWS WAF searches for character sequences that are likely to be malicious
// strings.
//
// To create and configure an XssMatchSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateXssMatchSet request.
//
// Submit a CreateXssMatchSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateXssMatchSet request.
//
// Submit an UpdateXssMatchSet request to specify the parts of web requests
// in which you want to allow, block, or count cross-site scripting attacks.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateXssMatchSetRequest.
//    req := client.CreateXssMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/CreateXssMatchSet
func (c *Client) CreateXssMatchSetRequest(input *CreateXssMatchSetInput) CreateXssMatchSetRequest {
	op := &aws.Operation{
		Name:       opCreateXssMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateXssMatchSetInput{}
	}

	req := c.newRequest(op, input, &CreateXssMatchSetOutput{})

	return CreateXssMatchSetRequest{Request: req, Input: input, Copy: c.CreateXssMatchSetRequest}
}

// CreateXssMatchSetRequest is the request type for the
// CreateXssMatchSet API operation.
type CreateXssMatchSetRequest struct {
	*aws.Request
	Input *CreateXssMatchSetInput
	Copy  func(*CreateXssMatchSetInput) CreateXssMatchSetRequest
}

// Send marshals and sends the CreateXssMatchSet API request.
func (r CreateXssMatchSetRequest) Send(ctx context.Context) (*CreateXssMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateXssMatchSetResponse{
		CreateXssMatchSetOutput: r.Request.Data.(*CreateXssMatchSetOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateXssMatchSetResponse is the response type for the
// CreateXssMatchSet API operation.
type CreateXssMatchSetResponse struct {
	*CreateXssMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateXssMatchSet request.
func (r *CreateXssMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
