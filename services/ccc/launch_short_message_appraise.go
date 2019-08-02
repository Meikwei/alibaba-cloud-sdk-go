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

// LaunchShortMessageAppraise invokes the ccc.LaunchShortMessageAppraise API synchronously
// api document: https://help.aliyun.com/api/ccc/launchshortmessageappraise.html
func (client *Client) LaunchShortMessageAppraise(request *LaunchShortMessageAppraiseRequest) (response *LaunchShortMessageAppraiseResponse, err error) {
	response = CreateLaunchShortMessageAppraiseResponse()
	err = client.DoAction(request, response)
	return
}

// LaunchShortMessageAppraiseWithChan invokes the ccc.LaunchShortMessageAppraise API asynchronously
// api document: https://help.aliyun.com/api/ccc/launchshortmessageappraise.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LaunchShortMessageAppraiseWithChan(request *LaunchShortMessageAppraiseRequest) (<-chan *LaunchShortMessageAppraiseResponse, <-chan error) {
	responseChan := make(chan *LaunchShortMessageAppraiseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LaunchShortMessageAppraise(request)
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

// LaunchShortMessageAppraiseWithCallback invokes the ccc.LaunchShortMessageAppraise API asynchronously
// api document: https://help.aliyun.com/api/ccc/launchshortmessageappraise.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LaunchShortMessageAppraiseWithCallback(request *LaunchShortMessageAppraiseRequest, callback func(response *LaunchShortMessageAppraiseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LaunchShortMessageAppraiseResponse
		var err error
		defer close(result)
		response, err = client.LaunchShortMessageAppraise(request)
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

// LaunchShortMessageAppraiseRequest is the request struct for api LaunchShortMessageAppraise
type LaunchShortMessageAppraiseRequest struct {
	*requests.RpcRequest
	Acid         string           `position:"Query" name:"Acid"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	ContactType  requests.Integer `position:"Query" name:"ContactType"`
	PhoneNumbers string           `position:"Query" name:"PhoneNumbers"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
}

// LaunchShortMessageAppraiseResponse is the response struct for api LaunchShortMessageAppraise
type LaunchShortMessageAppraiseResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateLaunchShortMessageAppraiseRequest creates a request to invoke LaunchShortMessageAppraise API
func CreateLaunchShortMessageAppraiseRequest() (request *LaunchShortMessageAppraiseRequest) {
	request = &LaunchShortMessageAppraiseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "LaunchShortMessageAppraise", "", "")
	return
}

// CreateLaunchShortMessageAppraiseResponse creates a response to parse from LaunchShortMessageAppraise response
func CreateLaunchShortMessageAppraiseResponse() (response *LaunchShortMessageAppraiseResponse) {
	response = &LaunchShortMessageAppraiseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
