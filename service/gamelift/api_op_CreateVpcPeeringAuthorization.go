// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type CreateVpcPeeringAuthorizationInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the AWS account that you use to manage your Amazon
	// GameLift fleet. You can find your Account ID in the AWS Management Console
	// under account settings.
	//
	// GameLiftAwsAccountId is a required field
	GameLiftAwsAccountId *string `min:"1" type:"string" required:"true"`

	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region where your fleet is deployed.
	// Look up a VPC ID using the VPC Dashboard (https://console.aws.amazon.com/vpc/)
	// in the AWS Management Console. Learn more about VPC peering in VPC Peering
	// with Amazon GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	//
	// PeerVpcId is a required field
	PeerVpcId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVpcPeeringAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcPeeringAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVpcPeeringAuthorizationInput"}

	if s.GameLiftAwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameLiftAwsAccountId"))
	}
	if s.GameLiftAwsAccountId != nil && len(*s.GameLiftAwsAccountId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameLiftAwsAccountId", 1))
	}

	if s.PeerVpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PeerVpcId"))
	}
	if s.PeerVpcId != nil && len(*s.PeerVpcId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PeerVpcId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type CreateVpcPeeringAuthorizationOutput struct {
	_ struct{} `type:"structure"`

	// Details on the requested VPC peering authorization, including expiration.
	VpcPeeringAuthorization *VpcPeeringAuthorization `type:"structure"`
}

// String returns the string representation
func (s CreateVpcPeeringAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVpcPeeringAuthorization = "CreateVpcPeeringAuthorization"

// CreateVpcPeeringAuthorizationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Requests authorization to create or delete a peer connection between the
// VPC for your Amazon GameLift fleet and a virtual private cloud (VPC) in your
// AWS account. VPC peering enables the game servers on your fleet to communicate
// directly with other AWS resources. Once you've received authorization, call
// CreateVpcPeeringConnection to establish the peering connection. For more
// information, see VPC Peering with Amazon GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
//
// You can peer with VPCs that are owned by any AWS account you have access
// to, including the account that you use to manage your Amazon GameLift fleets.
// You cannot peer with VPCs that are in different Regions.
//
// To request authorization to create a connection, call this operation from
// the AWS account with the VPC that you want to peer to your Amazon GameLift
// fleet. For example, to enable your game servers to retrieve data from a DynamoDB
// table, use the account that manages that DynamoDB resource. Identify the
// following values: (1) The ID of the VPC that you want to peer with, and (2)
// the ID of the AWS account that you use to manage Amazon GameLift. If successful,
// VPC peering is authorized for the specified VPC.
//
// To request authorization to delete a connection, call this operation from
// the AWS account with the VPC that is peered with your Amazon GameLift fleet.
// Identify the following values: (1) VPC ID that you want to delete the peering
// connection for, and (2) ID of the AWS account that you use to manage Amazon
// GameLift.
//
// The authorization remains valid for 24 hours unless it is canceled by a call
// to DeleteVpcPeeringAuthorization. You must create or delete the peering connection
// while the authorization is valid.
//
//    * CreateVpcPeeringAuthorization
//
//    * DescribeVpcPeeringAuthorizations
//
//    * DeleteVpcPeeringAuthorization
//
//    * CreateVpcPeeringConnection
//
//    * DescribeVpcPeeringConnections
//
//    * DeleteVpcPeeringConnection
//
//    // Example sending a request using CreateVpcPeeringAuthorizationRequest.
//    req := client.CreateVpcPeeringAuthorizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateVpcPeeringAuthorization
func (c *Client) CreateVpcPeeringAuthorizationRequest(input *CreateVpcPeeringAuthorizationInput) CreateVpcPeeringAuthorizationRequest {
	op := &aws.Operation{
		Name:       opCreateVpcPeeringAuthorization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcPeeringAuthorizationInput{}
	}

	req := c.newRequest(op, input, &CreateVpcPeeringAuthorizationOutput{})

	return CreateVpcPeeringAuthorizationRequest{Request: req, Input: input, Copy: c.CreateVpcPeeringAuthorizationRequest}
}

// CreateVpcPeeringAuthorizationRequest is the request type for the
// CreateVpcPeeringAuthorization API operation.
type CreateVpcPeeringAuthorizationRequest struct {
	*aws.Request
	Input *CreateVpcPeeringAuthorizationInput
	Copy  func(*CreateVpcPeeringAuthorizationInput) CreateVpcPeeringAuthorizationRequest
}

// Send marshals and sends the CreateVpcPeeringAuthorization API request.
func (r CreateVpcPeeringAuthorizationRequest) Send(ctx context.Context) (*CreateVpcPeeringAuthorizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcPeeringAuthorizationResponse{
		CreateVpcPeeringAuthorizationOutput: r.Request.Data.(*CreateVpcPeeringAuthorizationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcPeeringAuthorizationResponse is the response type for the
// CreateVpcPeeringAuthorization API operation.
type CreateVpcPeeringAuthorizationResponse struct {
	*CreateVpcPeeringAuthorizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcPeeringAuthorization request.
func (r *CreateVpcPeeringAuthorizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
