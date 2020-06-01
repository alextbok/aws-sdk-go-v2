// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing one or more of the following fields:
//
//    * CreateSnapshotInput$SnapshotDescription
//
//    * CreateSnapshotInput$VolumeARN
type CreateSnapshotInput struct {
	_ struct{} `type:"structure"`

	// Textual description of the snapshot that appears in the Amazon EC2 console,
	// Elastic Block Store snapshots panel in the Description field, and in the
	// AWS Storage Gateway snapshot Details pane, Description field
	//
	// SnapshotDescription is a required field
	SnapshotDescription *string `min:"1" type:"string" required:"true"`

	// A list of up to 50 tags that can be assigned to a snapshot. Each tag is a
	// key-value pair.
	//
	// Valid characters for key and value are letters, spaces, and numbers representable
	// in UTF-8 format, and the following special characters: + - = . _ : / @. The
	// maximum length of a tag's key is 128 characters, and the maximum length for
	// a tag's value is 256.
	Tags []Tag `type:"list"`

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation
	// to return a list of gateway volumes.
	//
	// VolumeARN is a required field
	VolumeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSnapshotInput"}

	if s.SnapshotDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotDescription"))
	}
	if s.SnapshotDescription != nil && len(*s.SnapshotDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnapshotDescription", 1))
	}

	if s.VolumeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeARN"))
	}
	if s.VolumeARN != nil && len(*s.VolumeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("VolumeARN", 50))
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

// A JSON object containing the following fields:
type CreateSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// The snapshot ID that is used to refer to the snapshot in future operations
	// such as describing snapshots (Amazon Elastic Compute Cloud API DescribeSnapshots)
	// or creating a volume from a snapshot (CreateStorediSCSIVolume).
	SnapshotId *string `type:"string"`

	// The Amazon Resource Name (ARN) of the volume of which the snapshot was taken.
	VolumeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s CreateSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSnapshot = "CreateSnapshot"

// CreateSnapshotRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Initiates a snapshot of a volume.
//
// AWS Storage Gateway provides the ability to back up point-in-time snapshots
// of your data to Amazon Simple Storage Service (Amazon S3) for durable off-site
// recovery, as well as import the data to an Amazon Elastic Block Store (EBS)
// volume in Amazon Elastic Compute Cloud (EC2). You can take snapshots of your
// gateway volume on a scheduled or ad hoc basis. This API enables you to take
// an ad hoc snapshot. For more information, see Editing a Snapshot Schedule
// (https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-volumes.html#SchedulingSnapshot).
//
// In the CreateSnapshot request you identify the volume by providing its Amazon
// Resource Name (ARN). You must also provide description for the snapshot.
// When AWS Storage Gateway takes the snapshot of specified volume, the snapshot
// and description appears in the AWS Storage Gateway Console. In response,
// AWS Storage Gateway returns you a snapshot ID. You can use this snapshot
// ID to check the snapshot progress or later use it when you want to create
// a volume from a snapshot. This operation is only supported in stored and
// cached volume gateway type.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more information,
// see DescribeSnapshots or DeleteSnapshot in the EC2 API reference (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Operations.html).
//
// Volume and snapshot IDs are changing to a longer length ID format. For more
// information, see the important note on the Welcome (https://docs.aws.amazon.com/storagegateway/latest/APIReference/Welcome.html)
// page.
//
//    // Example sending a request using CreateSnapshotRequest.
//    req := client.CreateSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CreateSnapshot
func (c *Client) CreateSnapshotRequest(input *CreateSnapshotInput) CreateSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateSnapshotOutput{})

	return CreateSnapshotRequest{Request: req, Input: input, Copy: c.CreateSnapshotRequest}
}

// CreateSnapshotRequest is the request type for the
// CreateSnapshot API operation.
type CreateSnapshotRequest struct {
	*aws.Request
	Input *CreateSnapshotInput
	Copy  func(*CreateSnapshotInput) CreateSnapshotRequest
}

// Send marshals and sends the CreateSnapshot API request.
func (r CreateSnapshotRequest) Send(ctx context.Context) (*CreateSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSnapshotResponse{
		CreateSnapshotOutput: r.Request.Data.(*CreateSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSnapshotResponse is the response type for the
// CreateSnapshot API operation.
type CreateSnapshotResponse struct {
	*CreateSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSnapshot request.
func (r *CreateSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
