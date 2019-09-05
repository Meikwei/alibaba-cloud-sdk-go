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

// ModifyPrivacyNumberCallDetail invokes the ccc.ModifyPrivacyNumberCallDetail API synchronously
// api document: https://help.aliyun.com/api/ccc/modifyprivacynumbercalldetail.html
func (client *Client) ModifyPrivacyNumberCallDetail(request *ModifyPrivacyNumberCallDetailRequest) (response *ModifyPrivacyNumberCallDetailResponse, err error) {
	response = CreateModifyPrivacyNumberCallDetailResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPrivacyNumberCallDetailWithChan invokes the ccc.ModifyPrivacyNumberCallDetail API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyprivacynumbercalldetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPrivacyNumberCallDetailWithChan(request *ModifyPrivacyNumberCallDetailRequest) (<-chan *ModifyPrivacyNumberCallDetailResponse, <-chan error) {
	responseChan := make(chan *ModifyPrivacyNumberCallDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPrivacyNumberCallDetail(request)
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

// ModifyPrivacyNumberCallDetailWithCallback invokes the ccc.ModifyPrivacyNumberCallDetail API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyprivacynumbercalldetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPrivacyNumberCallDetailWithCallback(request *ModifyPrivacyNumberCallDetailRequest, callback func(response *ModifyPrivacyNumberCallDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPrivacyNumberCallDetailResponse
		var err error
		defer close(result)
		response, err = client.ModifyPrivacyNumberCallDetail(request)
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

// ModifyPrivacyNumberCallDetailRequest is the request struct for api ModifyPrivacyNumberCallDetail
type ModifyPrivacyNumberCallDetailRequest struct {
	*requests.RpcRequest
	CallId     string `position:"Query" name:"CallId"`
	ContactId  string `position:"Query" name:"ContactId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ModifyPrivacyNumberCallDetailResponse is the response struct for api ModifyPrivacyNumberCallDetail
type ModifyPrivacyNumberCallDetailResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifyPrivacyNumberCallDetailRequest creates a request to invoke ModifyPrivacyNumberCallDetail API
func CreateModifyPrivacyNumberCallDetailRequest() (request *ModifyPrivacyNumberCallDetailRequest) {
	request = &ModifyPrivacyNumberCallDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifyPrivacyNumberCallDetail", "", "")
	return
}

// CreateModifyPrivacyNumberCallDetailResponse creates a response to parse from ModifyPrivacyNumberCallDetail response
func CreateModifyPrivacyNumberCallDetailResponse() (response *ModifyPrivacyNumberCallDetailResponse) {
	response = &ModifyPrivacyNumberCallDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
