// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request body for UpdateClusterConfiguration.
type UpdateClusterConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`

	// Represents the configuration that you want MSK to use for the cluster.
	//
	// ConfigurationInfo is a required field
	ConfigurationInfo *ConfigurationInfo `locationName:"configurationInfo" type:"structure" required:"true"`

	// The version of the cluster that you want to update.
	//
	// CurrentVersion is a required field
	CurrentVersion *string `locationName:"currentVersion" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateClusterConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClusterConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClusterConfigurationInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}

	if s.ConfigurationInfo == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationInfo"))
	}

	if s.CurrentVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentVersion"))
	}
	if s.ConfigurationInfo != nil {
		if err := s.ConfigurationInfo.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationInfo", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigurationInfo != nil {
		v := s.ConfigurationInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configurationInfo", v, metadata)
	}
	if s.CurrentVersion != nil {
		v := *s.CurrentVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response body for UpdateClusterConfiguration.
type UpdateClusterConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string `locationName:"clusterOperationArn" type:"string"`
}

// String returns the string representation
func (s UpdateClusterConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterOperationArn != nil {
		v := *s.ClusterOperationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterOperationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateClusterConfiguration = "UpdateClusterConfiguration"

// UpdateClusterConfigurationRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Updates the cluster with the configuration that is specified in the request
// body.
//
//    // Example sending a request using UpdateClusterConfigurationRequest.
//    req := client.UpdateClusterConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateClusterConfiguration
func (c *Client) UpdateClusterConfigurationRequest(input *UpdateClusterConfigurationInput) UpdateClusterConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateClusterConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/clusters/{clusterArn}/configuration",
	}

	if input == nil {
		input = &UpdateClusterConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateClusterConfigurationOutput{})

	return UpdateClusterConfigurationRequest{Request: req, Input: input, Copy: c.UpdateClusterConfigurationRequest}
}

// UpdateClusterConfigurationRequest is the request type for the
// UpdateClusterConfiguration API operation.
type UpdateClusterConfigurationRequest struct {
	*aws.Request
	Input *UpdateClusterConfigurationInput
	Copy  func(*UpdateClusterConfigurationInput) UpdateClusterConfigurationRequest
}

// Send marshals and sends the UpdateClusterConfiguration API request.
func (r UpdateClusterConfigurationRequest) Send(ctx context.Context) (*UpdateClusterConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClusterConfigurationResponse{
		UpdateClusterConfigurationOutput: r.Request.Data.(*UpdateClusterConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClusterConfigurationResponse is the response type for the
// UpdateClusterConfiguration API operation.
type UpdateClusterConfigurationResponse struct {
	*UpdateClusterConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClusterConfiguration request.
func (r *UpdateClusterConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
