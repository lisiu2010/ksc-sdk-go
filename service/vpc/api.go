// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opDescribeSubnets = "DescribeSubnets"

// DescribeSubnetsRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeSubnets operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeSubnets for more information on using the DescribeSubnets
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeSubnetsRequest method.
//    req, resp := client.DescribeSubnetsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/vpc-2016-03-04/DescribeSubnets
func (c *Vpc) DescribeSubnetsRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSubnets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSubnets API operation for vpc.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for vpc's
// API operation DescribeSubnets for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/vpc-2016-03-04/DescribeSubnets
func (c *Vpc) DescribeSubnets(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSubnetsRequest(input)
	return out, req.Send()
}

// DescribeSubnetsWithContext is the same as DescribeSubnets with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSubnets for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Vpc) DescribeSubnetsWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSubnetsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcs = "DescribeVpcs"

// DescribeVpcsRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeVpcs operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeVpcs for more information on using the DescribeVpcs
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeVpcsRequest method.
//    req, resp := client.DescribeVpcsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/vpc-2016-03-04/DescribeVpcs
func (c *Vpc) DescribeVpcsRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcs,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpcs API operation for vpc.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for vpc's
// API operation DescribeVpcs for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/vpc-2016-03-04/DescribeVpcs
func (c *Vpc) DescribeVpcs(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcsRequest(input)
	return out, req.Send()
}

// DescribeVpcsWithContext is the same as DescribeVpcs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Vpc) DescribeVpcsWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}
