// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteTranscriptionJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the transcription job to be deleted.
	//
	// TranscriptionJobName is a required field
	TranscriptionJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTranscriptionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTranscriptionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTranscriptionJobInput"}

	if s.TranscriptionJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TranscriptionJobName"))
	}
	if s.TranscriptionJobName != nil && len(*s.TranscriptionJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TranscriptionJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTranscriptionJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTranscriptionJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTranscriptionJob = "DeleteTranscriptionJob"

// DeleteTranscriptionJobRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Deletes a previously submitted transcription job along with any other generated
// results such as the transcription, models, and so on.
//
//    // Example sending a request using DeleteTranscriptionJobRequest.
//    req := client.DeleteTranscriptionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/DeleteTranscriptionJob
func (c *Client) DeleteTranscriptionJobRequest(input *DeleteTranscriptionJobInput) DeleteTranscriptionJobRequest {
	op := &aws.Operation{
		Name:       opDeleteTranscriptionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTranscriptionJobInput{}
	}

	req := c.newRequest(op, input, &DeleteTranscriptionJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteTranscriptionJobRequest{Request: req, Input: input, Copy: c.DeleteTranscriptionJobRequest}
}

// DeleteTranscriptionJobRequest is the request type for the
// DeleteTranscriptionJob API operation.
type DeleteTranscriptionJobRequest struct {
	*aws.Request
	Input *DeleteTranscriptionJobInput
	Copy  func(*DeleteTranscriptionJobInput) DeleteTranscriptionJobRequest
}

// Send marshals and sends the DeleteTranscriptionJob API request.
func (r DeleteTranscriptionJobRequest) Send(ctx context.Context) (*DeleteTranscriptionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTranscriptionJobResponse{
		DeleteTranscriptionJobOutput: r.Request.Data.(*DeleteTranscriptionJobOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTranscriptionJobResponse is the response type for the
// DeleteTranscriptionJob API operation.
type DeleteTranscriptionJobResponse struct {
	*DeleteTranscriptionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTranscriptionJob request.
func (r *DeleteTranscriptionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
