// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Deletes an archive rule.
type DeleteArchiveRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the analyzer that associated with the archive rule to delete.
	//
	// AnalyzerName is a required field
	AnalyzerName *string `location:"uri" locationName:"analyzerName" min:"1" type:"string" required:"true"`

	// A client token.
	ClientToken *string `location:"querystring" locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// The name of the rule to delete.
	//
	// RuleName is a required field
	RuleName *string `location:"uri" locationName:"ruleName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteArchiveRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteArchiveRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteArchiveRuleInput"}

	if s.AnalyzerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AnalyzerName"))
	}
	if s.AnalyzerName != nil && len(*s.AnalyzerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AnalyzerName", 1))
	}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteArchiveRuleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AnalyzerName != nil {
		v := *s.AnalyzerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "analyzerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RuleName != nil {
		v := *s.RuleName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ruleName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteArchiveRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteArchiveRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteArchiveRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteArchiveRule = "DeleteArchiveRule"

// DeleteArchiveRuleRequest returns a request value for making API operation for
// Access Analyzer.
//
// Deletes the specified archive rule.
//
//    // Example sending a request using DeleteArchiveRuleRequest.
//    req := client.DeleteArchiveRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/accessanalyzer-2019-11-01/DeleteArchiveRule
func (c *Client) DeleteArchiveRuleRequest(input *DeleteArchiveRuleInput) DeleteArchiveRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteArchiveRule,
		HTTPMethod: "DELETE",
		HTTPPath:   "/analyzer/{analyzerName}/archive-rule/{ruleName}",
	}

	if input == nil {
		input = &DeleteArchiveRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteArchiveRuleOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteArchiveRuleRequest{Request: req, Input: input, Copy: c.DeleteArchiveRuleRequest}
}

// DeleteArchiveRuleRequest is the request type for the
// DeleteArchiveRule API operation.
type DeleteArchiveRuleRequest struct {
	*aws.Request
	Input *DeleteArchiveRuleInput
	Copy  func(*DeleteArchiveRuleInput) DeleteArchiveRuleRequest
}

// Send marshals and sends the DeleteArchiveRule API request.
func (r DeleteArchiveRuleRequest) Send(ctx context.Context) (*DeleteArchiveRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteArchiveRuleResponse{
		DeleteArchiveRuleOutput: r.Request.Data.(*DeleteArchiveRuleOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteArchiveRuleResponse is the response type for the
// DeleteArchiveRule API operation.
type DeleteArchiveRuleResponse struct {
	*DeleteArchiveRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteArchiveRule request.
func (r *DeleteArchiveRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
