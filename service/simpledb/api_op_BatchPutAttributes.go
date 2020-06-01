// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type BatchPutAttributesInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which the attributes are being stored.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// A list of items on which to perform the operation.
	//
	// Items is a required field
	Items []ReplaceableItem `locationNameList:"Item" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s BatchPutAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchPutAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchPutAttributesInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.Items == nil {
		invalidParams.Add(aws.NewErrParamRequired("Items"))
	}
	if s.Items != nil {
		for i, v := range s.Items {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Items", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchPutAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s BatchPutAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchPutAttributes = "BatchPutAttributes"

// BatchPutAttributesRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// The BatchPutAttributes operation creates or replaces attributes within one
// or more items. By using this operation, the client can perform multiple PutAttribute
// operation with a single call. This helps yield savings in round trips and
// latencies, enabling Amazon SimpleDB to optimize requests and generally produce
// better throughput.
//
// The client may specify the item name with the Item.X.ItemName parameter.
// The client may specify new attributes using a combination of the Item.X.Attribute.Y.Name
// and Item.X.Attribute.Y.Value parameters. The client may specify the first
// attribute for the first item using the parameters Item.0.Attribute.0.Name
// and Item.0.Attribute.0.Value, and for the second attribute for the first
// item by the parameters Item.0.Attribute.1.Name and Item.0.Attribute.1.Value,
// and so on.
//
// Attributes are uniquely identified within an item by their name/value combination.
// For example, a single item can have the attributes { "first_name", "first_value"
// } and { "first_name", "second_value" }. However, it cannot have two attribute
// instances where both the Item.X.Attribute.Y.Name and Item.X.Attribute.Y.Value
// are the same.
//
// Optionally, the requester can supply the Replace parameter for each individual
// value. Setting this value to true will cause the new attribute values to
// replace the existing attribute values. For example, if an item I has the
// attributes { 'a', '1' }, { 'b', '2'} and { 'b', '3' } and the requester does
// a BatchPutAttributes of {'I', 'b', '4' } with the Replace parameter set to
// true, the final attributes of the item will be { 'a', '1' } and { 'b', '4'
// }, replacing the previous values of the 'b' attribute with the new value.
//   You cannot specify an empty string as an item or as an attribute name.
//   The
//    BatchPutAttributes
//  operation succeeds or fails in its entirety. There are no partial puts.
//  This operation is vulnerable to exceeding the maximum URL size when making
//  a REST request using the HTTP GET method. This operation does not support
//  conditions using
//    Expected.X.Name
// ,
//    Expected.X.Value
// , or
//    Expected.X.Exists
// .
// You can execute multiple BatchPutAttributes operations and other operations
// in parallel. However, large numbers of concurrent BatchPutAttributes calls
// can result in Service Unavailable (503) responses.
//
// The following limitations are enforced for this operation:
//    * 256 attribute name-value pairs per item
//
//    * 1 MB request size
//
//    * 1 billion attributes per domain
//
//    * 10 GB of total user data storage per domain
//
//    * 25 item limit per BatchPutAttributes operation
//
//    // Example sending a request using BatchPutAttributesRequest.
//    req := client.BatchPutAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) BatchPutAttributesRequest(input *BatchPutAttributesInput) BatchPutAttributesRequest {
	op := &aws.Operation{
		Name:       opBatchPutAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchPutAttributesInput{}
	}

	req := c.newRequest(op, input, &BatchPutAttributesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return BatchPutAttributesRequest{Request: req, Input: input, Copy: c.BatchPutAttributesRequest}
}

// BatchPutAttributesRequest is the request type for the
// BatchPutAttributes API operation.
type BatchPutAttributesRequest struct {
	*aws.Request
	Input *BatchPutAttributesInput
	Copy  func(*BatchPutAttributesInput) BatchPutAttributesRequest
}

// Send marshals and sends the BatchPutAttributes API request.
func (r BatchPutAttributesRequest) Send(ctx context.Context) (*BatchPutAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchPutAttributesResponse{
		BatchPutAttributesOutput: r.Request.Data.(*BatchPutAttributesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchPutAttributesResponse is the response type for the
// BatchPutAttributes API operation.
type BatchPutAttributesResponse struct {
	*BatchPutAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchPutAttributes request.
func (r *BatchPutAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
