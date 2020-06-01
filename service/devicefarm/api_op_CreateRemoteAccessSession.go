// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Creates and submits a request to start a remote access session.
type CreateRemoteAccessSessionInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the client. If you want access to multiple devices
	// on the same client, you should pass the same clientId value in each call
	// to CreateRemoteAccessSession. This identifier is required only if remoteDebugEnabled
	// is set to true.
	//
	// Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html).
	ClientId *string `locationName:"clientId" type:"string"`

	// The configuration information for the remote access session request.
	Configuration *CreateRemoteAccessSessionConfiguration `locationName:"configuration" type:"structure"`

	// The ARN of the device for which you want to create a remote access session.
	//
	// DeviceArn is a required field
	DeviceArn *string `locationName:"deviceArn" min:"32" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the device instance for which you want
	// to create a remote access session.
	InstanceArn *string `locationName:"instanceArn" min:"32" type:"string"`

	// The interaction mode of the remote access session. Valid values are:
	//
	//    * INTERACTIVE: You can interact with the iOS device by viewing, touching,
	//    and rotating the screen. You cannot run XCUITest framework-based tests
	//    in this mode.
	//
	//    * NO_VIDEO: You are connected to the device, but cannot interact with
	//    it or view the screen. This mode has the fastest test execution speed.
	//    You can run XCUITest framework-based tests in this mode.
	//
	//    * VIDEO_ONLY: You can view the screen, but cannot touch or rotate it.
	//    You can run XCUITest framework-based tests and watch the screen in this
	//    mode.
	InteractionMode InteractionMode `locationName:"interactionMode" type:"string" enum:"true"`

	// The name of the remote access session to create.
	Name *string `locationName:"name" type:"string"`

	// The Amazon Resource Name (ARN) of the project for which you want to create
	// a remote access session.
	//
	// ProjectArn is a required field
	ProjectArn *string `locationName:"projectArn" min:"32" type:"string" required:"true"`

	// Set to true if you want to access devices remotely for debugging in your
	// remote access session.
	//
	// Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html).
	RemoteDebugEnabled *bool `locationName:"remoteDebugEnabled" type:"boolean"`

	// The Amazon Resource Name (ARN) for the app to be recorded in the remote access
	// session.
	RemoteRecordAppArn *string `locationName:"remoteRecordAppArn" min:"32" type:"string"`

	// Set to true to enable remote recording for the remote access session.
	RemoteRecordEnabled *bool `locationName:"remoteRecordEnabled" type:"boolean"`

	// When set to true, for private devices, Device Farm does not sign your app
	// again. For public devices, Device Farm always signs your apps again.
	//
	// For more information on how Device Farm modifies your uploads during tests,
	// see Do you modify my app? (https://aws.amazon.com/device-farm/faq/)
	SkipAppResign *bool `locationName:"skipAppResign" type:"boolean"`

	// Ignored. The public key of the ssh key pair you want to use for connecting
	// to remote devices in your remote debugging session. This key is required
	// only if remoteDebugEnabled is set to true.
	//
	// Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html).
	SshPublicKey *string `locationName:"sshPublicKey" type:"string"`
}

// String returns the string representation
func (s CreateRemoteAccessSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRemoteAccessSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRemoteAccessSessionInput"}

	if s.DeviceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceArn"))
	}
	if s.DeviceArn != nil && len(*s.DeviceArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceArn", 32))
	}
	if s.InstanceArn != nil && len(*s.InstanceArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceArn", 32))
	}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 32))
	}
	if s.RemoteRecordAppArn != nil && len(*s.RemoteRecordAppArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("RemoteRecordAppArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the server response from a request to create a remote access session.
type CreateRemoteAccessSessionOutput struct {
	_ struct{} `type:"structure"`

	// A container that describes the remote access session when the request to
	// create a remote access session is sent.
	RemoteAccessSession *RemoteAccessSession `locationName:"remoteAccessSession" type:"structure"`
}

// String returns the string representation
func (s CreateRemoteAccessSessionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRemoteAccessSession = "CreateRemoteAccessSession"

// CreateRemoteAccessSessionRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Specifies and starts a remote access session.
//
//    // Example sending a request using CreateRemoteAccessSessionRequest.
//    req := client.CreateRemoteAccessSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateRemoteAccessSession
func (c *Client) CreateRemoteAccessSessionRequest(input *CreateRemoteAccessSessionInput) CreateRemoteAccessSessionRequest {
	op := &aws.Operation{
		Name:       opCreateRemoteAccessSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRemoteAccessSessionInput{}
	}

	req := c.newRequest(op, input, &CreateRemoteAccessSessionOutput{})

	return CreateRemoteAccessSessionRequest{Request: req, Input: input, Copy: c.CreateRemoteAccessSessionRequest}
}

// CreateRemoteAccessSessionRequest is the request type for the
// CreateRemoteAccessSession API operation.
type CreateRemoteAccessSessionRequest struct {
	*aws.Request
	Input *CreateRemoteAccessSessionInput
	Copy  func(*CreateRemoteAccessSessionInput) CreateRemoteAccessSessionRequest
}

// Send marshals and sends the CreateRemoteAccessSession API request.
func (r CreateRemoteAccessSessionRequest) Send(ctx context.Context) (*CreateRemoteAccessSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRemoteAccessSessionResponse{
		CreateRemoteAccessSessionOutput: r.Request.Data.(*CreateRemoteAccessSessionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRemoteAccessSessionResponse is the response type for the
// CreateRemoteAccessSession API operation.
type CreateRemoteAccessSessionResponse struct {
	*CreateRemoteAccessSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRemoteAccessSession request.
func (r *CreateRemoteAccessSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
