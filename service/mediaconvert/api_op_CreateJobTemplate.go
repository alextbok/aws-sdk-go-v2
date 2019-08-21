// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Send your create job template request with the name of the template and the
// JSON for the template. The template JSON should include everything in a valid
// job, except for input location and filename, IAM role, and user metadata.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/CreateJobTemplateRequest
type CreateJobTemplateInput struct {
	_ struct{} `type:"structure"`

	// Accelerated transcoding can significantly speed up jobs with long, visually
	// complex content. Outputs that use this feature incur pro-tier pricing. For
	// information about feature limitations, see the AWS Elemental MediaConvert
	// User Guide.
	AccelerationSettings *AccelerationSettings `locationName:"accelerationSettings" type:"structure"`

	// Optional. A category for the job template you are creating
	Category *string `locationName:"category" type:"string"`

	// Optional. A description of the job template you are creating.
	Description *string `locationName:"description" type:"string"`

	// The name of the job template you are creating.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// Specify the relative priority for this job. In any given queue, the service
	// begins processing the job with the highest value first. When more than one
	// job has the same priority, the service begins processing the job that you
	// submitted first. If you don't specify a priority, the service uses the default
	// value 0.
	Priority *int64 `locationName:"priority" type:"integer"`

	// Optional. The queue that jobs created from this template are assigned to.
	// If you don't specify this, jobs will go to the default queue.
	Queue *string `locationName:"queue" type:"string"`

	// JobTemplateSettings contains all the transcode settings saved in the template
	// that will be applied to jobs created from it.
	//
	// Settings is a required field
	Settings *JobTemplateSettings `locationName:"settings" type:"structure" required:"true"`

	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch
	// Events. Set the interval, in seconds, between status updates. MediaConvert
	// sends an update at this interval from the time the service begins processing
	// your job to the time it completes the transcode or encounters an error.
	StatusUpdateInterval StatusUpdateInterval `locationName:"statusUpdateInterval" type:"string" enum:"true"`

	// The tags that you want to add to the resource. You can tag resources with
	// a key-value pair or with only a key.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateJobTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateJobTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateJobTemplateInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Priority != nil && *s.Priority < -50 {
		invalidParams.Add(aws.NewErrParamMinValue("Priority", -50))
	}

	if s.Settings == nil {
		invalidParams.Add(aws.NewErrParamRequired("Settings"))
	}
	if s.AccelerationSettings != nil {
		if err := s.AccelerationSettings.Validate(); err != nil {
			invalidParams.AddNested("AccelerationSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			invalidParams.AddNested("Settings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccelerationSettings != nil {
		v := s.AccelerationSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "accelerationSettings", v, metadata)
	}
	if s.Category != nil {
		v := *s.Category

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "category", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Priority != nil {
		v := *s.Priority

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "priority", protocol.Int64Value(v), metadata)
	}
	if s.Queue != nil {
		v := *s.Queue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Settings != nil {
		v := s.Settings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "settings", v, metadata)
	}
	if len(s.StatusUpdateInterval) > 0 {
		v := s.StatusUpdateInterval

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusUpdateInterval", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Successful create job template requests will return the template JSON.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/CreateJobTemplateResponse
type CreateJobTemplateOutput struct {
	_ struct{} `type:"structure"`

	// A job template is a pre-made set of encoding instructions that you can use
	// to quickly create a job.
	JobTemplate *JobTemplate `locationName:"jobTemplate" type:"structure"`
}

// String returns the string representation
func (s CreateJobTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobTemplate != nil {
		v := s.JobTemplate

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jobTemplate", v, metadata)
	}
	return nil
}

const opCreateJobTemplate = "CreateJobTemplate"

// CreateJobTemplateRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Create a new job template. For information about job templates see the User
// Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
//
//    // Example sending a request using CreateJobTemplateRequest.
//    req := client.CreateJobTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/CreateJobTemplate
func (c *Client) CreateJobTemplateRequest(input *CreateJobTemplateInput) CreateJobTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateJobTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/2017-08-29/jobTemplates",
	}

	if input == nil {
		input = &CreateJobTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateJobTemplateOutput{})
	return CreateJobTemplateRequest{Request: req, Input: input, Copy: c.CreateJobTemplateRequest}
}

// CreateJobTemplateRequest is the request type for the
// CreateJobTemplate API operation.
type CreateJobTemplateRequest struct {
	*aws.Request
	Input *CreateJobTemplateInput
	Copy  func(*CreateJobTemplateInput) CreateJobTemplateRequest
}

// Send marshals and sends the CreateJobTemplate API request.
func (r CreateJobTemplateRequest) Send(ctx context.Context) (*CreateJobTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJobTemplateResponse{
		CreateJobTemplateOutput: r.Request.Data.(*CreateJobTemplateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJobTemplateResponse is the response type for the
// CreateJobTemplate API operation.
type CreateJobTemplateResponse struct {
	*CreateJobTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJobTemplate request.
func (r *CreateJobTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
