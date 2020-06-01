// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetResourceShareAssociationsInput struct {
	_ struct{} `type:"structure"`

	// The association status.
	AssociationStatus ResourceShareAssociationStatus `locationName:"associationStatus" type:"string" enum:"true"`

	// The association type. Specify PRINCIPAL to list the principals that are associated
	// with the specified resource share. Specify RESOURCE to list the resources
	// that are associated with the specified resource share.
	//
	// AssociationType is a required field
	AssociationType ResourceShareAssociationType `locationName:"associationType" type:"string" required:"true" enum:"true"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The principal. You cannot specify this parameter if the association type
	// is RESOURCE.
	Principal *string `locationName:"principal" type:"string"`

	// The Amazon Resource Name (ARN) of the resource. You cannot specify this parameter
	// if the association type is PRINCIPAL.
	ResourceArn *string `locationName:"resourceArn" type:"string"`

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string `locationName:"resourceShareArns" type:"list"`
}

// String returns the string representation
func (s GetResourceShareAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceShareAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourceShareAssociationsInput"}
	if len(s.AssociationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AssociationType"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourceShareAssociationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.AssociationStatus) > 0 {
		v := s.AssociationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "associationStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.AssociationType) > 0 {
		v := s.AssociationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "associationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareArns != nil {
		v := s.ResourceShareArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceShareArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type GetResourceShareAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the associations.
	ResourceShareAssociations []ResourceShareAssociation `locationName:"resourceShareAssociations" type:"list"`
}

// String returns the string representation
func (s GetResourceShareAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourceShareAssociationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareAssociations != nil {
		v := s.ResourceShareAssociations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceShareAssociations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetResourceShareAssociations = "GetResourceShareAssociations"

// GetResourceShareAssociationsRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Gets the resources or principals for the resource shares that you own.
//
//    // Example sending a request using GetResourceShareAssociationsRequest.
//    req := client.GetResourceShareAssociationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/GetResourceShareAssociations
func (c *Client) GetResourceShareAssociationsRequest(input *GetResourceShareAssociationsInput) GetResourceShareAssociationsRequest {
	op := &aws.Operation{
		Name:       opGetResourceShareAssociations,
		HTTPMethod: "POST",
		HTTPPath:   "/getresourceshareassociations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetResourceShareAssociationsInput{}
	}

	req := c.newRequest(op, input, &GetResourceShareAssociationsOutput{})

	return GetResourceShareAssociationsRequest{Request: req, Input: input, Copy: c.GetResourceShareAssociationsRequest}
}

// GetResourceShareAssociationsRequest is the request type for the
// GetResourceShareAssociations API operation.
type GetResourceShareAssociationsRequest struct {
	*aws.Request
	Input *GetResourceShareAssociationsInput
	Copy  func(*GetResourceShareAssociationsInput) GetResourceShareAssociationsRequest
}

// Send marshals and sends the GetResourceShareAssociations API request.
func (r GetResourceShareAssociationsRequest) Send(ctx context.Context) (*GetResourceShareAssociationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourceShareAssociationsResponse{
		GetResourceShareAssociationsOutput: r.Request.Data.(*GetResourceShareAssociationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetResourceShareAssociationsRequestPaginator returns a paginator for GetResourceShareAssociations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetResourceShareAssociationsRequest(input)
//   p := ram.NewGetResourceShareAssociationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetResourceShareAssociationsPaginator(req GetResourceShareAssociationsRequest) GetResourceShareAssociationsPaginator {
	return GetResourceShareAssociationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetResourceShareAssociationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetResourceShareAssociationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetResourceShareAssociationsPaginator struct {
	aws.Pager
}

func (p *GetResourceShareAssociationsPaginator) CurrentPage() *GetResourceShareAssociationsOutput {
	return p.Pager.CurrentPage().(*GetResourceShareAssociationsOutput)
}

// GetResourceShareAssociationsResponse is the response type for the
// GetResourceShareAssociations API operation.
type GetResourceShareAssociationsResponse struct {
	*GetResourceShareAssociationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourceShareAssociations request.
func (r *GetResourceShareAssociationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
