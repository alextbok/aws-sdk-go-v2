// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AdminListGroupsForUserInput struct {
	_ struct{} `type:"structure"`

	// The limit of the request to list groups.
	Limit *int64 `type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `min:"1" type:"string"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The username for the user.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminListGroupsForUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminListGroupsForUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminListGroupsForUserInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AdminListGroupsForUserOutput struct {
	_ struct{} `type:"structure"`

	// The groups that the user belongs to.
	Groups []GroupType `type:"list"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s AdminListGroupsForUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminListGroupsForUser = "AdminListGroupsForUser"

// AdminListGroupsForUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the groups that the user belongs to.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminListGroupsForUserRequest.
//    req := client.AdminListGroupsForUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminListGroupsForUser
func (c *Client) AdminListGroupsForUserRequest(input *AdminListGroupsForUserInput) AdminListGroupsForUserRequest {
	op := &aws.Operation{
		Name:       opAdminListGroupsForUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &AdminListGroupsForUserInput{}
	}

	req := c.newRequest(op, input, &AdminListGroupsForUserOutput{})

	return AdminListGroupsForUserRequest{Request: req, Input: input, Copy: c.AdminListGroupsForUserRequest}
}

// AdminListGroupsForUserRequest is the request type for the
// AdminListGroupsForUser API operation.
type AdminListGroupsForUserRequest struct {
	*aws.Request
	Input *AdminListGroupsForUserInput
	Copy  func(*AdminListGroupsForUserInput) AdminListGroupsForUserRequest
}

// Send marshals and sends the AdminListGroupsForUser API request.
func (r AdminListGroupsForUserRequest) Send(ctx context.Context) (*AdminListGroupsForUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminListGroupsForUserResponse{
		AdminListGroupsForUserOutput: r.Request.Data.(*AdminListGroupsForUserOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewAdminListGroupsForUserRequestPaginator returns a paginator for AdminListGroupsForUser.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.AdminListGroupsForUserRequest(input)
//   p := cognitoidentityprovider.NewAdminListGroupsForUserRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewAdminListGroupsForUserPaginator(req AdminListGroupsForUserRequest) AdminListGroupsForUserPaginator {
	return AdminListGroupsForUserPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *AdminListGroupsForUserInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// AdminListGroupsForUserPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type AdminListGroupsForUserPaginator struct {
	aws.Pager
}

func (p *AdminListGroupsForUserPaginator) CurrentPage() *AdminListGroupsForUserOutput {
	return p.Pager.CurrentPage().(*AdminListGroupsForUserOutput)
}

// AdminListGroupsForUserResponse is the response type for the
// AdminListGroupsForUser API operation.
type AdminListGroupsForUserResponse struct {
	*AdminListGroupsForUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminListGroupsForUser request.
func (r *AdminListGroupsForUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
