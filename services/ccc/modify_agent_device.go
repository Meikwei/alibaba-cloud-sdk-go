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

// ModifyAgentDevice invokes the ccc.ModifyAgentDevice API synchronously
// api document: https://help.aliyun.com/api/ccc/modifyagentdevice.html
func (client *Client) ModifyAgentDevice(request *ModifyAgentDeviceRequest) (response *ModifyAgentDeviceResponse, err error) {
	response = CreateModifyAgentDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAgentDeviceWithChan invokes the ccc.ModifyAgentDevice API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyagentdevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAgentDeviceWithChan(request *ModifyAgentDeviceRequest) (<-chan *ModifyAgentDeviceResponse, <-chan error) {
	responseChan := make(chan *ModifyAgentDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAgentDevice(request)
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

// ModifyAgentDeviceWithCallback invokes the ccc.ModifyAgentDevice API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyagentdevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAgentDeviceWithCallback(request *ModifyAgentDeviceRequest, callback func(response *ModifyAgentDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAgentDeviceResponse
		var err error
		defer close(result)
		response, err = client.ModifyAgentDevice(request)
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

// ModifyAgentDeviceRequest is the request struct for api ModifyAgentDevice
type ModifyAgentDeviceRequest struct {
	*requests.RpcRequest
	AgentDeviceId requests.Integer `position:"Query" name:"AgentDeviceId"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	IsLogin       requests.Integer `position:"Query" name:"IsLogin"`
}

// ModifyAgentDeviceResponse is the response struct for api ModifyAgentDevice
type ModifyAgentDeviceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifyAgentDeviceRequest creates a request to invoke ModifyAgentDevice API
func CreateModifyAgentDeviceRequest() (request *ModifyAgentDeviceRequest) {
	request = &ModifyAgentDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifyAgentDevice", "", "")
	return
}

// CreateModifyAgentDeviceResponse creates a response to parse from ModifyAgentDevice response
func CreateModifyAgentDeviceResponse() (response *ModifyAgentDeviceResponse) {
	response = &ModifyAgentDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
