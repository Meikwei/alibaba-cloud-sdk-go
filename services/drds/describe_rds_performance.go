package drds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeRDSPerformance invokes the drds.DescribeRDSPerformance API synchronously
// api document: https://help.aliyun.com/api/drds/describerdsperformance.html
func (client *Client) DescribeRDSPerformance(request *DescribeRDSPerformanceRequest) (response *DescribeRDSPerformanceResponse, err error) {
	response = CreateDescribeRDSPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRDSPerformanceWithChan invokes the drds.DescribeRDSPerformance API asynchronously
// api document: https://help.aliyun.com/api/drds/describerdsperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRDSPerformanceWithChan(request *DescribeRDSPerformanceRequest) (<-chan *DescribeRDSPerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeRDSPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRDSPerformance(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeRDSPerformanceWithCallback invokes the drds.DescribeRDSPerformance API asynchronously
// api document: https://help.aliyun.com/api/drds/describerdsperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRDSPerformanceWithCallback(request *DescribeRDSPerformanceRequest, callback func(response *DescribeRDSPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRDSPerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeRDSPerformance(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeRDSPerformanceRequest is the request struct for api DescribeRDSPerformance
type DescribeRDSPerformanceRequest struct {
	*requests.RpcRequest
	Keys           string           `position:"Query" name:"Keys"`
	EndTime        requests.Integer `position:"Query" name:"EndTime"`
	StartTime      requests.Integer `position:"Query" name:"StartTime"`
	RdsInstanceId  string           `position:"Query" name:"RdsInstanceId"`
	DrdsInstanceId string           `position:"Query" name:"DrdsInstanceId"`
	DbInstType     string           `position:"Query" name:"DbInstType"`
}

// DescribeRDSPerformanceResponse is the response struct for api DescribeRDSPerformance
type DescribeRDSPerformanceResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Success   bool                     `json:"Success" xml:"Success"`
	Data      []PartialPerformanceData `json:"Data" xml:"Data"`
}

// CreateDescribeRDSPerformanceRequest creates a request to invoke DescribeRDSPerformance API
func CreateDescribeRDSPerformanceRequest() (request *DescribeRDSPerformanceRequest) {
	request = &DescribeRDSPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeRDSPerformance", "Drds", "openAPI")
	return
}

// CreateDescribeRDSPerformanceResponse creates a response to parse from DescribeRDSPerformance response
func CreateDescribeRDSPerformanceResponse() (response *DescribeRDSPerformanceResponse) {
	response = &DescribeRDSPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
