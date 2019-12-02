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

// DescribeDrdsDbTestLink invokes the drds.DescribeDrdsDbTestLink API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbtestlink.html
func (client *Client) DescribeDrdsDbTestLink(request *DescribeDrdsDbTestLinkRequest) (response *DescribeDrdsDbTestLinkResponse, err error) {
	response = CreateDescribeDrdsDbTestLinkResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsDbTestLinkWithChan invokes the drds.DescribeDrdsDbTestLink API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbtestlink.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbTestLinkWithChan(request *DescribeDrdsDbTestLinkRequest) (<-chan *DescribeDrdsDbTestLinkResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsDbTestLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsDbTestLink(request)
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

// DescribeDrdsDbTestLinkWithCallback invokes the drds.DescribeDrdsDbTestLink API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbtestlink.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbTestLinkWithCallback(request *DescribeDrdsDbTestLinkRequest, callback func(response *DescribeDrdsDbTestLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsDbTestLinkResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsDbTestLink(request)
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

// DescribeDrdsDbTestLinkRequest is the request struct for api DescribeDrdsDbTestLink
type DescribeDrdsDbTestLinkRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// DescribeDrdsDbTestLinkResponse is the response struct for api DescribeDrdsDbTestLink
type DescribeDrdsDbTestLinkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateDescribeDrdsDbTestLinkRequest creates a request to invoke DescribeDrdsDbTestLink API
func CreateDescribeDrdsDbTestLinkRequest() (request *DescribeDrdsDbTestLinkRequest) {
	request = &DescribeDrdsDbTestLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsDbTestLink", "Drds", "openAPI")
	return
}

// CreateDescribeDrdsDbTestLinkResponse creates a response to parse from DescribeDrdsDbTestLink response
func CreateDescribeDrdsDbTestLinkResponse() (response *DescribeDrdsDbTestLinkResponse) {
	response = &DescribeDrdsDbTestLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
