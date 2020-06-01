// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to delete the specific dataset.
type DeleteDatasetInput struct {
	_ struct{} `type:"structure"`

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// DatasetName is a required field
	DatasetName *string `location:"uri" locationName:"DatasetName" min:"1" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityId is a required field
	IdentityId *string `location:"uri" locationName:"IdentityId" min:"1" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `location:"uri" locationName:"IdentityPoolId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDatasetInput"}

	if s.DatasetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetName"))
	}
	if s.DatasetName != nil && len(*s.DatasetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetName", 1))
	}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDatasetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DatasetName != nil {
		v := *s.DatasetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DatasetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityId != nil {
		v := *s.IdentityId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response to a successful DeleteDataset request.
type DeleteDatasetOutput struct {
	_ struct{} `type:"structure"`

	// A collection of data for an identity pool. An identity pool can have multiple
	// datasets. A dataset is per identity and can be general or associated with
	// a particular entity in an application (like a saved game). Datasets are automatically
	// created if they don't exist. Data is synced by dataset, and a dataset can
	// hold up to 1MB of key-value pairs.
	Dataset *Dataset `type:"structure"`
}

// String returns the string representation
func (s DeleteDatasetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDatasetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Dataset != nil {
		v := s.Dataset

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Dataset", v, metadata)
	}
	return nil
}

const opDeleteDataset = "DeleteDataset"

// DeleteDatasetRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Deletes the specific dataset. The dataset will be deleted permanently, and
// the action can't be undone. Datasets that this dataset was merged with will
// no longer report the merge. Any subsequent operation on this dataset will
// result in a ResourceNotFoundException.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
//
//    // Example sending a request using DeleteDatasetRequest.
//    req := client.DeleteDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DeleteDataset
func (c *Client) DeleteDatasetRequest(input *DeleteDatasetInput) DeleteDatasetRequest {
	op := &aws.Operation{
		Name:       opDeleteDataset,
		HTTPMethod: "DELETE",
		HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
	}

	if input == nil {
		input = &DeleteDatasetInput{}
	}

	req := c.newRequest(op, input, &DeleteDatasetOutput{})

	return DeleteDatasetRequest{Request: req, Input: input, Copy: c.DeleteDatasetRequest}
}

// DeleteDatasetRequest is the request type for the
// DeleteDataset API operation.
type DeleteDatasetRequest struct {
	*aws.Request
	Input *DeleteDatasetInput
	Copy  func(*DeleteDatasetInput) DeleteDatasetRequest
}

// Send marshals and sends the DeleteDataset API request.
func (r DeleteDatasetRequest) Send(ctx context.Context) (*DeleteDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDatasetResponse{
		DeleteDatasetOutput: r.Request.Data.(*DeleteDatasetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDatasetResponse is the response type for the
// DeleteDataset API operation.
type DeleteDatasetResponse struct {
	*DeleteDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDataset request.
func (r *DeleteDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
