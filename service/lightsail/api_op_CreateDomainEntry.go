// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDomainEntryInput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the domain entry
	// request.
	//
	// DomainEntry is a required field
	DomainEntry *DomainEntry `locationName:"domainEntry" type:"structure" required:"true"`

	// The domain name (e.g., example.com) for which you want to create the domain
	// entry.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDomainEntryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDomainEntryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDomainEntryInput"}

	if s.DomainEntry == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainEntry"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDomainEntryOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the time stamp of the request, and the resources affected
	// by the request.
	Operation *Operation `locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s CreateDomainEntryOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDomainEntry = "CreateDomainEntry"

// CreateDomainEntryRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates one of the following entry records associated with the domain: Address
// (A), canonical name (CNAME), mail exchanger (MX), name server (NS), start
// of authority (SOA), service locator (SRV), or text (TXT).
//
// The create domain entry operation supports tag-based access control via resource
// tags applied to the resource identified by domain name. For more information,
// see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateDomainEntryRequest.
//    req := client.CreateDomainEntryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateDomainEntry
func (c *Client) CreateDomainEntryRequest(input *CreateDomainEntryInput) CreateDomainEntryRequest {
	op := &aws.Operation{
		Name:       opCreateDomainEntry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDomainEntryInput{}
	}

	req := c.newRequest(op, input, &CreateDomainEntryOutput{})
	return CreateDomainEntryRequest{Request: req, Input: input, Copy: c.CreateDomainEntryRequest}
}

// CreateDomainEntryRequest is the request type for the
// CreateDomainEntry API operation.
type CreateDomainEntryRequest struct {
	*aws.Request
	Input *CreateDomainEntryInput
	Copy  func(*CreateDomainEntryInput) CreateDomainEntryRequest
}

// Send marshals and sends the CreateDomainEntry API request.
func (r CreateDomainEntryRequest) Send(ctx context.Context) (*CreateDomainEntryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainEntryResponse{
		CreateDomainEntryOutput: r.Request.Data.(*CreateDomainEntryOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDomainEntryResponse is the response type for the
// CreateDomainEntry API operation.
type CreateDomainEntryResponse struct {
	*CreateDomainEntryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDomainEntry request.
func (r *CreateDomainEntryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
