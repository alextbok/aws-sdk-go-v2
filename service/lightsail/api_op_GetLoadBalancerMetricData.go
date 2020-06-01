// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLoadBalancerMetricDataInput struct {
	_ struct{} `type:"structure"`

	// The end time of the period.
	//
	// EndTime is a required field
	EndTime *time.Time `locationName:"endTime" type:"timestamp" required:"true"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`

	// The metric for which you want to return information.
	//
	// Valid load balancer metric names are listed below, along with the most useful
	// statistics to include in your request, and the published unit value.
	//
	//    * ClientTLSNegotiationErrorCount - The number of TLS connections initiated
	//    by the client that did not establish a session with the load balancer
	//    due to a TLS error generated by the load balancer. Possible causes include
	//    a mismatch of ciphers or protocols. Statistics: The most useful statistic
	//    is Sum. Unit: The published unit is Count.
	//
	//    * HealthyHostCount - The number of target instances that are considered
	//    healthy. Statistics: The most useful statistic are Average, Minimum, and
	//    Maximum. Unit: The published unit is Count.
	//
	//    * HTTPCode_Instance_2XX_Count - The number of HTTP 2XX response codes
	//    generated by the target instances. This does not include any response
	//    codes generated by the load balancer. Statistics: The most useful statistic
	//    is Sum. Note that Minimum, Maximum, and Average all return 1. Unit: The
	//    published unit is Count.
	//
	//    * HTTPCode_Instance_3XX_Count - The number of HTTP 3XX response codes
	//    generated by the target instances. This does not include any response
	//    codes generated by the load balancer. Statistics: The most useful statistic
	//    is Sum. Note that Minimum, Maximum, and Average all return 1. Unit: The
	//    published unit is Count.
	//
	//    * HTTPCode_Instance_4XX_Count - The number of HTTP 4XX response codes
	//    generated by the target instances. This does not include any response
	//    codes generated by the load balancer. Statistics: The most useful statistic
	//    is Sum. Note that Minimum, Maximum, and Average all return 1. Unit: The
	//    published unit is Count.
	//
	//    * HTTPCode_Instance_5XX_Count - The number of HTTP 5XX response codes
	//    generated by the target instances. This does not include any response
	//    codes generated by the load balancer. Statistics: The most useful statistic
	//    is Sum. Note that Minimum, Maximum, and Average all return 1. Unit: The
	//    published unit is Count.
	//
	//    * HTTPCode_LB_4XX_Count - The number of HTTP 4XX client error codes that
	//    originated from the load balancer. Client errors are generated when requests
	//    are malformed or incomplete. These requests were not received by the target
	//    instance. This count does not include response codes generated by the
	//    target instances. Statistics: The most useful statistic is Sum. Note that
	//    Minimum, Maximum, and Average all return 1. Unit: The published unit is
	//    Count.
	//
	//    * HTTPCode_LB_5XX_Count - The number of HTTP 5XX server error codes that
	//    originated from the load balancer. This does not include any response
	//    codes generated by the target instance. This metric is reported if there
	//    are no healthy instances attached to the load balancer, or if the request
	//    rate exceeds the capacity of the instances (spillover) or the load balancer.
	//    Statistics: The most useful statistic is Sum. Note that Minimum, Maximum,
	//    and Average all return 1. Unit: The published unit is Count.
	//
	//    * InstanceResponseTime - The time elapsed, in seconds, after the request
	//    leaves the load balancer until a response from the target instance is
	//    received. Statistics: The most useful statistic is Average. Unit: The
	//    published unit is Seconds.
	//
	//    * RejectedConnectionCount - The number of connections that were rejected
	//    because the load balancer had reached its maximum number of connections.
	//    Statistics: The most useful statistic is Sum. Unit: The published unit
	//    is Count.
	//
	//    * RequestCount - The number of requests processed over IPv4. This count
	//    includes only the requests with a response generated by a target instance
	//    of the load balancer. Statistics: The most useful statistic is Sum. Note
	//    that Minimum, Maximum, and Average all return 1. Unit: The published unit
	//    is Count.
	//
	//    * UnhealthyHostCount - The number of target instances that are considered
	//    unhealthy. Statistics: The most useful statistic are Average, Minimum,
	//    and Maximum. Unit: The published unit is Count.
	//
	// MetricName is a required field
	MetricName LoadBalancerMetricName `locationName:"metricName" type:"string" required:"true" enum:"true"`

	// The granularity, in seconds, of the returned data points.
	//
	// Period is a required field
	Period *int64 `locationName:"period" min:"60" type:"integer" required:"true"`

	// The start time of the period.
	//
	// StartTime is a required field
	StartTime *time.Time `locationName:"startTime" type:"timestamp" required:"true"`

	// The statistic for the metric.
	//
	// The following statistics are available:
	//
	//    * Minimum - The lowest value observed during the specified period. Use
	//    this value to determine low volumes of activity for your application.
	//
	//    * Maximum - The highest value observed during the specified period. Use
	//    this value to determine high volumes of activity for your application.
	//
	//    * Sum - All values submitted for the matching metric added together. You
	//    can use this statistic to determine the total volume of a metric.
	//
	//    * Average - The value of Sum / SampleCount during the specified period.
	//    By comparing this statistic with the Minimum and Maximum values, you can
	//    determine the full scope of a metric and how close the average use is
	//    to the Minimum and Maximum values. This comparison helps you to know when
	//    to increase or decrease your resources.
	//
	//    * SampleCount - The count, or number, of data points used for the statistical
	//    calculation.
	//
	// Statistics is a required field
	Statistics []MetricStatistic `locationName:"statistics" type:"list" required:"true"`

	// The unit for the metric data request. Valid units depend on the metric data
	// being required. For the valid units with each available metric, see the metricName
	// parameter.
	//
	// Unit is a required field
	Unit MetricUnit `locationName:"unit" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetLoadBalancerMetricDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLoadBalancerMetricDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLoadBalancerMetricDataInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}
	if len(s.MetricName) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Period == nil {
		invalidParams.Add(aws.NewErrParamRequired("Period"))
	}
	if s.Period != nil && *s.Period < 60 {
		invalidParams.Add(aws.NewErrParamMinValue("Period", 60))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if s.Statistics == nil {
		invalidParams.Add(aws.NewErrParamRequired("Statistics"))
	}
	if len(s.Unit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Unit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLoadBalancerMetricDataOutput struct {
	_ struct{} `type:"structure"`

	// An array of metric datapoint objects.
	MetricData []MetricDatapoint `locationName:"metricData" type:"list"`

	// The metric about which you are receiving information. Valid values are listed
	// below, along with the most useful statistics to include in your request.
	//
	//    * ClientTLSNegotiationErrorCount - The number of TLS connections initiated
	//    by the client that did not establish a session with the load balancer.
	//    Possible causes include a mismatch of ciphers or protocols. Statistics:
	//    The most useful statistic is Sum.
	//
	//    * HealthyHostCount - The number of target instances that are considered
	//    healthy. Statistics: The most useful statistic are Average, Minimum, and
	//    Maximum.
	//
	//    * UnhealthyHostCount - The number of target instances that are considered
	//    unhealthy. Statistics: The most useful statistic are Average, Minimum,
	//    and Maximum.
	//
	//    * HTTPCode_LB_4XX_Count - The number of HTTP 4XX client error codes that
	//    originate from the load balancer. Client errors are generated when requests
	//    are malformed or incomplete. These requests have not been received by
	//    the target instance. This count does not include any response codes generated
	//    by the target instances. Statistics: The most useful statistic is Sum.
	//    Note that Minimum, Maximum, and Average all return 1.
	//
	//    * HTTPCode_LB_5XX_Count - The number of HTTP 5XX server error codes that
	//    originate from the load balancer. This count does not include any response
	//    codes generated by the target instances. Statistics: The most useful statistic
	//    is Sum. Note that Minimum, Maximum, and Average all return 1. Note that
	//    Minimum, Maximum, and Average all return 1.
	//
	//    * HTTPCode_Instance_2XX_Count - The number of HTTP response codes generated
	//    by the target instances. This does not include any response codes generated
	//    by the load balancer. Statistics: The most useful statistic is Sum. Note
	//    that Minimum, Maximum, and Average all return 1.
	//
	//    * HTTPCode_Instance_3XX_Count - The number of HTTP response codes generated
	//    by the target instances. This does not include any response codes generated
	//    by the load balancer. Statistics: The most useful statistic is Sum. Note
	//    that Minimum, Maximum, and Average all return 1.
	//
	//    * HTTPCode_Instance_4XX_Count - The number of HTTP response codes generated
	//    by the target instances. This does not include any response codes generated
	//    by the load balancer. Statistics: The most useful statistic is Sum. Note
	//    that Minimum, Maximum, and Average all return 1.
	//
	//    * HTTPCode_Instance_5XX_Count - The number of HTTP response codes generated
	//    by the target instances. This does not include any response codes generated
	//    by the load balancer. Statistics: The most useful statistic is Sum. Note
	//    that Minimum, Maximum, and Average all return 1.
	//
	//    * InstanceResponseTime - The time elapsed, in seconds, after the request
	//    leaves the load balancer until a response from the target instance is
	//    received. Statistics: The most useful statistic is Average.
	//
	//    * RejectedConnectionCount - The number of connections that were rejected
	//    because the load balancer had reached its maximum number of connections.
	//    Statistics: The most useful statistic is Sum.
	//
	//    * RequestCount - The number of requests processed over IPv4. This count
	//    includes only the requests with a response generated by a target instance
	//    of the load balancer. Statistics: The most useful statistic is Sum. Note
	//    that Minimum, Maximum, and Average all return 1.
	MetricName LoadBalancerMetricName `locationName:"metricName" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetLoadBalancerMetricDataOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLoadBalancerMetricData = "GetLoadBalancerMetricData"

// GetLoadBalancerMetricDataRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about health metrics for your Lightsail load balancer.
//
//    // Example sending a request using GetLoadBalancerMetricDataRequest.
//    req := client.GetLoadBalancerMetricDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetLoadBalancerMetricData
func (c *Client) GetLoadBalancerMetricDataRequest(input *GetLoadBalancerMetricDataInput) GetLoadBalancerMetricDataRequest {
	op := &aws.Operation{
		Name:       opGetLoadBalancerMetricData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLoadBalancerMetricDataInput{}
	}

	req := c.newRequest(op, input, &GetLoadBalancerMetricDataOutput{})

	return GetLoadBalancerMetricDataRequest{Request: req, Input: input, Copy: c.GetLoadBalancerMetricDataRequest}
}

// GetLoadBalancerMetricDataRequest is the request type for the
// GetLoadBalancerMetricData API operation.
type GetLoadBalancerMetricDataRequest struct {
	*aws.Request
	Input *GetLoadBalancerMetricDataInput
	Copy  func(*GetLoadBalancerMetricDataInput) GetLoadBalancerMetricDataRequest
}

// Send marshals and sends the GetLoadBalancerMetricData API request.
func (r GetLoadBalancerMetricDataRequest) Send(ctx context.Context) (*GetLoadBalancerMetricDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoadBalancerMetricDataResponse{
		GetLoadBalancerMetricDataOutput: r.Request.Data.(*GetLoadBalancerMetricDataOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoadBalancerMetricDataResponse is the response type for the
// GetLoadBalancerMetricData API operation.
type GetLoadBalancerMetricDataResponse struct {
	*GetLoadBalancerMetricDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoadBalancerMetricData request.
func (r *GetLoadBalancerMetricDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
