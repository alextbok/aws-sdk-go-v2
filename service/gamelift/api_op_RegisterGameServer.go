// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterGameServerInput struct {
	_ struct{} `type:"structure"`

	// Information needed to make inbound client connections to the game server.
	// This might include IP address and port, DNS name, etc.
	ConnectionInfo *string `min:"1" type:"string"`

	// A game server tag that can be used to request sorted lists of game servers
	// using ListGameServers. Custom sort keys are developer-defined based on how
	// you want to organize the retrieved game server information.
	CustomSortKey *string `min:"1" type:"string"`

	// A set of custom game server properties, formatted as a single string value.
	// This data is passed to a game client or service when it requests information
	// on a game servers using ListGameServers or ClaimGameServer.
	GameServerData *string `min:"1" type:"string"`

	// An identifier for the game server group where the game server is running.
	// You can use either the GameServerGroup name or ARN value.
	//
	// GameServerGroupName is a required field
	GameServerGroupName *string `min:"1" type:"string" required:"true"`

	// A custom string that uniquely identifies the new game server. Game server
	// IDs are developer-defined and must be unique across all game server groups
	// in your AWS account.
	//
	// GameServerId is a required field
	GameServerId *string `min:"3" type:"string" required:"true"`

	// The unique identifier for the instance where the game server is running.
	// This ID is available in the instance metadata.
	//
	// InstanceId is a required field
	InstanceId *string `min:"19" type:"string" required:"true"`

	// A list of labels to assign to the new game server resource. Tags are developer-defined
	// key-value pairs. Tagging AWS resources are useful for resource management,
	// access management, and cost allocation. For more information, see Tagging
	// AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the AWS General Reference. Once the resource is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference
	// for actual tagging limits.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s RegisterGameServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterGameServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterGameServerInput"}
	if s.ConnectionInfo != nil && len(*s.ConnectionInfo) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConnectionInfo", 1))
	}
	if s.CustomSortKey != nil && len(*s.CustomSortKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomSortKey", 1))
	}
	if s.GameServerData != nil && len(*s.GameServerData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameServerData", 1))
	}

	if s.GameServerGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameServerGroupName"))
	}
	if s.GameServerGroupName != nil && len(*s.GameServerGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameServerGroupName", 1))
	}

	if s.GameServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameServerId"))
	}
	if s.GameServerId != nil && len(*s.GameServerId) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("GameServerId", 3))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 19))
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

type RegisterGameServerOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the newly created game server resource.
	GameServer *GameServer `type:"structure"`
}

// String returns the string representation
func (s RegisterGameServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterGameServer = "RegisterGameServer"

// RegisterGameServerRequest returns a request value for making API operation for
// Amazon GameLift.
//
// This action is part of Amazon GameLift FleetIQ with game server groups, which
// is in preview release and is subject to change.
//
// Creates a new game server resource and notifies GameLift FleetIQ that the
// game server is ready to host gameplay and players. This action is called
// by a game server process that is running on an instance in a game server
// group. Registering game servers enables GameLift FleetIQ to track available
// game servers and enables game clients and services to claim a game server
// for a new game session.
//
// To register a game server, identify the game server group and instance where
// the game server is running, and provide a unique identifier for the game
// server. You can also include connection and game server data; when a game
// client or service requests a game server by calling ClaimGameServer, this
// information is returned in response.
//
// Once a game server is successfully registered, it is put in status AVAILABLE.
// A request to register a game server may fail if the instance it is in the
// process of shutting down as part of instance rebalancing or scale-down activity.
//
// Learn more
//
// GameLift FleetIQ Guide (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
//
// Related operations
//
//    * RegisterGameServer
//
//    * ListGameServers
//
//    * ClaimGameServer
//
//    * DescribeGameServer
//
//    * UpdateGameServer
//
//    * DeregisterGameServer
//
//    // Example sending a request using RegisterGameServerRequest.
//    req := client.RegisterGameServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/RegisterGameServer
func (c *Client) RegisterGameServerRequest(input *RegisterGameServerInput) RegisterGameServerRequest {
	op := &aws.Operation{
		Name:       opRegisterGameServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterGameServerInput{}
	}

	req := c.newRequest(op, input, &RegisterGameServerOutput{})

	return RegisterGameServerRequest{Request: req, Input: input, Copy: c.RegisterGameServerRequest}
}

// RegisterGameServerRequest is the request type for the
// RegisterGameServer API operation.
type RegisterGameServerRequest struct {
	*aws.Request
	Input *RegisterGameServerInput
	Copy  func(*RegisterGameServerInput) RegisterGameServerRequest
}

// Send marshals and sends the RegisterGameServer API request.
func (r RegisterGameServerRequest) Send(ctx context.Context) (*RegisterGameServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterGameServerResponse{
		RegisterGameServerOutput: r.Request.Data.(*RegisterGameServerOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterGameServerResponse is the response type for the
// RegisterGameServer API operation.
type RegisterGameServerResponse struct {
	*RegisterGameServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterGameServer request.
func (r *RegisterGameServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
