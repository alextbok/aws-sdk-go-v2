// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateNodeInput struct {
	_ struct{} `type:"structure"`

	// Configuration properties for publishing to Amazon CloudWatch Logs.
	LogPublishingConfiguration *NodeLogPublishingConfiguration `type:"structure"`

	// The unique ID of the member that owns the node.
	//
	// MemberId is a required field
	MemberId *string `location:"uri" locationName:"memberId" min:"1" type:"string" required:"true"`

	// The unique ID of the Managed Blockchain network to which the node belongs.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The unique ID of the node.
	//
	// NodeId is a required field
	NodeId *string `location:"uri" locationName:"nodeId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNodeInput"}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if s.NodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeId"))
	}
	if s.NodeId != nil && len(*s.NodeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NodeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateNodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LogPublishingConfiguration != nil {
		v := s.LogPublishingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LogPublishingConfiguration", v, metadata)
	}
	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "memberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodeId != nil {
		v := *s.NodeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "nodeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateNodeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateNodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateNodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateNode = "UpdateNode"

// UpdateNodeRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Updates a node configuration with new parameters.
//
//    // Example sending a request using UpdateNodeRequest.
//    req := client.UpdateNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/UpdateNode
func (c *Client) UpdateNodeRequest(input *UpdateNodeInput) UpdateNodeRequest {
	op := &aws.Operation{
		Name:       opUpdateNode,
		HTTPMethod: "PATCH",
		HTTPPath:   "/networks/{networkId}/members/{memberId}/nodes/{nodeId}",
	}

	if input == nil {
		input = &UpdateNodeInput{}
	}

	req := c.newRequest(op, input, &UpdateNodeOutput{})

	return UpdateNodeRequest{Request: req, Input: input, Copy: c.UpdateNodeRequest}
}

// UpdateNodeRequest is the request type for the
// UpdateNode API operation.
type UpdateNodeRequest struct {
	*aws.Request
	Input *UpdateNodeInput
	Copy  func(*UpdateNodeInput) UpdateNodeRequest
}

// Send marshals and sends the UpdateNode API request.
func (r UpdateNodeRequest) Send(ctx context.Context) (*UpdateNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNodeResponse{
		UpdateNodeOutput: r.Request.Data.(*UpdateNodeOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNodeResponse is the response type for the
// UpdateNode API operation.
type UpdateNodeResponse struct {
	*UpdateNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNode request.
func (r *UpdateNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
