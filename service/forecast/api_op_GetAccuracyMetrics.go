// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccuracyMetricsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the predictor to get metrics for.
	//
	// PredictorArn is a required field
	PredictorArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccuracyMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccuracyMetricsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccuracyMetricsInput"}

	if s.PredictorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PredictorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAccuracyMetricsOutput struct {
	_ struct{} `type:"structure"`

	// An array of results from evaluating the predictor.
	PredictorEvaluationResults []EvaluationResult `type:"list"`
}

// String returns the string representation
func (s GetAccuracyMetricsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAccuracyMetrics = "GetAccuracyMetrics"

// GetAccuracyMetricsRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Provides metrics on the accuracy of the models that were trained by the CreatePredictor
// operation. Use metrics to see how well the model performed and to decide
// whether to use the predictor to generate a forecast. For more information,
// see metrics.
//
// This operation generates metrics for each backtest window that was evaluated.
// The number of backtest windows (NumberOfBacktestWindows) is specified using
// the EvaluationParameters object, which is optionally included in the CreatePredictor
// request. If NumberOfBacktestWindows isn't specified, the number defaults
// to one.
//
// The parameters of the filling method determine which items contribute to
// the metrics. If you want all items to contribute, specify zero. If you want
// only those items that have complete data in the range being evaluated to
// contribute, specify nan. For more information, see FeaturizationMethod.
//
// Before you can get accuracy metrics, the Status of the predictor must be
// ACTIVE, signifying that training has completed. To get the status, use the
// DescribePredictor operation.
//
//    // Example sending a request using GetAccuracyMetricsRequest.
//    req := client.GetAccuracyMetricsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/GetAccuracyMetrics
func (c *Client) GetAccuracyMetricsRequest(input *GetAccuracyMetricsInput) GetAccuracyMetricsRequest {
	op := &aws.Operation{
		Name:       opGetAccuracyMetrics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAccuracyMetricsInput{}
	}

	req := c.newRequest(op, input, &GetAccuracyMetricsOutput{})

	return GetAccuracyMetricsRequest{Request: req, Input: input, Copy: c.GetAccuracyMetricsRequest}
}

// GetAccuracyMetricsRequest is the request type for the
// GetAccuracyMetrics API operation.
type GetAccuracyMetricsRequest struct {
	*aws.Request
	Input *GetAccuracyMetricsInput
	Copy  func(*GetAccuracyMetricsInput) GetAccuracyMetricsRequest
}

// Send marshals and sends the GetAccuracyMetrics API request.
func (r GetAccuracyMetricsRequest) Send(ctx context.Context) (*GetAccuracyMetricsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccuracyMetricsResponse{
		GetAccuracyMetricsOutput: r.Request.Data.(*GetAccuracyMetricsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccuracyMetricsResponse is the response type for the
// GetAccuracyMetrics API operation.
type GetAccuracyMetricsResponse struct {
	*GetAccuracyMetricsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccuracyMetrics request.
func (r *GetAccuracyMetricsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
