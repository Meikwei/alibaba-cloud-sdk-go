package ccc

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

// ListJobStatus invokes the ccc.ListJobStatus API synchronously
// api document: https://help.aliyun.com/api/ccc/listjobstatus.html
func (client *Client) ListJobStatus(request *ListJobStatusRequest) (response *ListJobStatusResponse, err error) {
	response = CreateListJobStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobStatusWithChan invokes the ccc.ListJobStatus API asynchronously
// api document: https://help.aliyun.com/api/ccc/listjobstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobStatusWithChan(request *ListJobStatusRequest) (<-chan *ListJobStatusResponse, <-chan error) {
	responseChan := make(chan *ListJobStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobStatus(request)
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

// ListJobStatusWithCallback invokes the ccc.ListJobStatus API asynchronously
// api document: https://help.aliyun.com/api/ccc/listjobstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobStatusWithCallback(request *ListJobStatusRequest, callback func(response *ListJobStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobStatusResponse
		var err error
		defer close(result)
		response, err = client.ListJobStatus(request)
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

// ListJobStatusRequest is the request struct for api ListJobStatus
type ListJobStatusRequest struct {
	*requests.RpcRequest
	ContactName   string           `position:"Query" name:"ContactName"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	TimeAlignment string           `position:"Query" name:"TimeAlignment"`
	GroupId       string           `position:"Query" name:"GroupId"`
	PhoneNumber   string           `position:"Query" name:"PhoneNumber"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	ScenarioId    string           `position:"Query" name:"ScenarioId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// ListJobStatusResponse is the response struct for api ListJobStatus
type ListJobStatusResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Jobs           Jobs   `json:"Jobs" xml:"Jobs"`
}

// CreateListJobStatusRequest creates a request to invoke ListJobStatus API
func CreateListJobStatusRequest() (request *ListJobStatusRequest) {
	request = &ListJobStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListJobStatus", "", "")
	return
}

// CreateListJobStatusResponse creates a response to parse from ListJobStatus response
func CreateListJobStatusResponse() (response *ListJobStatusResponse) {
	response = &ListJobStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
