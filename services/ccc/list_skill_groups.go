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

// ListSkillGroups invokes the ccc.ListSkillGroups API synchronously
// api document: https://help.aliyun.com/api/ccc/listskillgroups.html
func (client *Client) ListSkillGroups(request *ListSkillGroupsRequest) (response *ListSkillGroupsResponse, err error) {
	response = CreateListSkillGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSkillGroupsWithChan invokes the ccc.ListSkillGroups API asynchronously
// api document: https://help.aliyun.com/api/ccc/listskillgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupsWithChan(request *ListSkillGroupsRequest) (<-chan *ListSkillGroupsResponse, <-chan error) {
	responseChan := make(chan *ListSkillGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSkillGroups(request)
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

// ListSkillGroupsWithCallback invokes the ccc.ListSkillGroups API asynchronously
// api document: https://help.aliyun.com/api/ccc/listskillgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupsWithCallback(request *ListSkillGroupsRequest, callback func(response *ListSkillGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSkillGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListSkillGroups(request)
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

// ListSkillGroupsRequest is the request struct for api ListSkillGroups
type ListSkillGroupsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListSkillGroupsResponse is the response struct for api ListSkillGroups
type ListSkillGroupsResponse struct {
	*responses.BaseResponse
	RequestId      string                       `json:"RequestId" xml:"RequestId"`
	Success        bool                         `json:"Success" xml:"Success"`
	Code           string                       `json:"Code" xml:"Code"`
	Message        string                       `json:"Message" xml:"Message"`
	HttpStatusCode int                          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	SkillGroups    SkillGroupsInListSkillGroups `json:"SkillGroups" xml:"SkillGroups"`
}

// CreateListSkillGroupsRequest creates a request to invoke ListSkillGroups API
func CreateListSkillGroupsRequest() (request *ListSkillGroupsRequest) {
	request = &ListSkillGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListSkillGroups", "", "")
	return
}

// CreateListSkillGroupsResponse creates a response to parse from ListSkillGroups response
func CreateListSkillGroupsResponse() (response *ListSkillGroupsResponse) {
	response = &ListSkillGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
