// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RevokeClusterSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	// The IP range for which to revoke access. This range must be a valid Classless
	// Inter-Domain Routing (CIDR) block of IP addresses. If CIDRIP is specified,
	// EC2SecurityGroupName and EC2SecurityGroupOwnerId cannot be provided.
	CIDRIP *string `type:"string"`

	// The name of the security Group from which to revoke the ingress rule.
	//
	// ClusterSecurityGroupName is a required field
	ClusterSecurityGroupName *string `type:"string" required:"true"`

	// The name of the EC2 Security Group whose access is to be revoked. If EC2SecurityGroupName
	// is specified, EC2SecurityGroupOwnerId must also be provided and CIDRIP cannot
	// be provided.
	EC2SecurityGroupName *string `type:"string"`

	// The AWS account number of the owner of the security group specified in the
	// EC2SecurityGroupName parameter. The AWS access key ID is not an acceptable
	// value. If EC2SecurityGroupOwnerId is specified, EC2SecurityGroupName must
	// also be provided. and CIDRIP cannot be provided.
	//
	// Example: 111122223333
	EC2SecurityGroupOwnerId *string `type:"string"`
}

// String returns the string representation
func (s RevokeClusterSecurityGroupIngressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeClusterSecurityGroupIngressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevokeClusterSecurityGroupIngressInput"}

	if s.ClusterSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterSecurityGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RevokeClusterSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	// Describes a security group.
	ClusterSecurityGroup *ClusterSecurityGroup `type:"structure"`
}

// String returns the string representation
func (s RevokeClusterSecurityGroupIngressOutput) String() string {
	return awsutil.Prettify(s)
}

const opRevokeClusterSecurityGroupIngress = "RevokeClusterSecurityGroupIngress"

// RevokeClusterSecurityGroupIngressRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Revokes an ingress rule in an Amazon Redshift security group for a previously
// authorized IP range or Amazon EC2 security group. To add an ingress rule,
// see AuthorizeClusterSecurityGroupIngress. For information about managing
// security groups, go to Amazon Redshift Cluster Security Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using RevokeClusterSecurityGroupIngressRequest.
//    req := client.RevokeClusterSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/RevokeClusterSecurityGroupIngress
func (c *Client) RevokeClusterSecurityGroupIngressRequest(input *RevokeClusterSecurityGroupIngressInput) RevokeClusterSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opRevokeClusterSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokeClusterSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &RevokeClusterSecurityGroupIngressOutput{})

	return RevokeClusterSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.RevokeClusterSecurityGroupIngressRequest}
}

// RevokeClusterSecurityGroupIngressRequest is the request type for the
// RevokeClusterSecurityGroupIngress API operation.
type RevokeClusterSecurityGroupIngressRequest struct {
	*aws.Request
	Input *RevokeClusterSecurityGroupIngressInput
	Copy  func(*RevokeClusterSecurityGroupIngressInput) RevokeClusterSecurityGroupIngressRequest
}

// Send marshals and sends the RevokeClusterSecurityGroupIngress API request.
func (r RevokeClusterSecurityGroupIngressRequest) Send(ctx context.Context) (*RevokeClusterSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeClusterSecurityGroupIngressResponse{
		RevokeClusterSecurityGroupIngressOutput: r.Request.Data.(*RevokeClusterSecurityGroupIngressOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeClusterSecurityGroupIngressResponse is the response type for the
// RevokeClusterSecurityGroupIngress API operation.
type RevokeClusterSecurityGroupIngressResponse struct {
	*RevokeClusterSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeClusterSecurityGroupIngress request.
func (r *RevokeClusterSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
