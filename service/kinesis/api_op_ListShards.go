// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListShardsInput struct {
	_ struct{} `type:"structure"`

	// Specify this parameter to indicate that you want to list the shards starting
	// with the shard whose ID immediately follows ExclusiveStartShardId.
	//
	// If you don't specify this parameter, the default behavior is for ListShards
	// to list the shards starting with the first one in the stream.
	//
	// You cannot specify this parameter if you specify NextToken.
	ExclusiveStartShardId *string `min:"1" type:"string"`

	// The maximum number of shards to return in a single call to ListShards. The
	// minimum value you can specify for this parameter is 1, and the maximum is
	// 1,000, which is also the default.
	//
	// When the number of shards to be listed is greater than the value of MaxResults,
	// the response contains a NextToken value that you can use in a subsequent
	// call to ListShards to list the next set of shards.
	MaxResults *int64 `min:"1" type:"integer"`

	// When the number of shards in the data stream is greater than the default
	// value for the MaxResults parameter, or if you explicitly specify a value
	// for MaxResults that is less than the number of shards in the data stream,
	// the response includes a pagination token named NextToken. You can specify
	// this NextToken value in a subsequent call to ListShards to list the next
	// set of shards.
	//
	// Don't specify StreamName or StreamCreationTimestamp if you specify NextToken
	// because the latter unambiguously identifies the stream.
	//
	// You can optionally specify a value for the MaxResults parameter when you
	// specify NextToken. If you specify a MaxResults value that is less than the
	// number of shards that the operation returns if you don't specify MaxResults,
	// the response will contain a new NextToken value. You can use the new NextToken
	// value in a subsequent call to the ListShards operation.
	//
	// Tokens expire after 300 seconds. When you obtain a value for NextToken in
	// the response to a call to ListShards, you have 300 seconds to use that value.
	// If you specify an expired token in a call to ListShards, you get ExpiredNextTokenException.
	NextToken *string `min:"1" type:"string"`

	// Specify this input parameter to distinguish data streams that have the same
	// name. For example, if you create a data stream and then delete it, and you
	// later create another data stream with the same name, you can use this input
	// parameter to specify which of the two streams you want to list the shards
	// for.
	//
	// You cannot specify this parameter if you specify the NextToken parameter.
	StreamCreationTimestamp *time.Time `type:"timestamp"`

	// The name of the data stream whose shards you want to list.
	//
	// You cannot specify this parameter if you specify the NextToken parameter.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListShardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListShardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListShardsInput"}
	if s.ExclusiveStartShardId != nil && len(*s.ExclusiveStartShardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartShardId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListShardsOutput struct {
	_ struct{} `type:"structure"`

	// When the number of shards in the data stream is greater than the default
	// value for the MaxResults parameter, or if you explicitly specify a value
	// for MaxResults that is less than the number of shards in the data stream,
	// the response includes a pagination token named NextToken. You can specify
	// this NextToken value in a subsequent call to ListShards to list the next
	// set of shards. For more information about the use of this pagination token
	// when calling the ListShards operation, see ListShardsInput$NextToken.
	//
	// Tokens expire after 300 seconds. When you obtain a value for NextToken in
	// the response to a call to ListShards, you have 300 seconds to use that value.
	// If you specify an expired token in a call to ListShards, you get ExpiredNextTokenException.
	NextToken *string `min:"1" type:"string"`

	// An array of JSON objects. Each object represents one shard and specifies
	// the IDs of the shard, the shard's parent, and the shard that's adjacent to
	// the shard's parent. Each object also contains the starting and ending hash
	// keys and the starting and ending sequence numbers for the shard.
	Shards []Shard `type:"list"`
}

// String returns the string representation
func (s ListShardsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListShards = "ListShards"

// ListShardsRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Lists the shards in a stream and provides information about each shard. This
// operation has a limit of 100 transactions per second per data stream.
//
// This API is a new operation that is used by the Amazon Kinesis Client Library
// (KCL). If you have a fine-grained IAM policy that only allows specific operations,
// you must update your policy to allow calls to this API. For more information,
// see Controlling Access to Amazon Kinesis Data Streams Resources Using IAM
// (https://docs.aws.amazon.com/streams/latest/dev/controlling-access.html).
//
//    // Example sending a request using ListShardsRequest.
//    req := client.ListShardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/ListShards
func (c *Client) ListShardsRequest(input *ListShardsInput) ListShardsRequest {
	op := &aws.Operation{
		Name:       opListShards,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListShardsInput{}
	}

	req := c.newRequest(op, input, &ListShardsOutput{})

	return ListShardsRequest{Request: req, Input: input, Copy: c.ListShardsRequest}
}

// ListShardsRequest is the request type for the
// ListShards API operation.
type ListShardsRequest struct {
	*aws.Request
	Input *ListShardsInput
	Copy  func(*ListShardsInput) ListShardsRequest
}

// Send marshals and sends the ListShards API request.
func (r ListShardsRequest) Send(ctx context.Context) (*ListShardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListShardsResponse{
		ListShardsOutput: r.Request.Data.(*ListShardsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListShardsResponse is the response type for the
// ListShards API operation.
type ListShardsResponse struct {
	*ListShardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListShards request.
func (r *ListShardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
