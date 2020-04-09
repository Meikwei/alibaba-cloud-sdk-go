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

package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeStoragePackages invokes the nas.DescribeStoragePackages API synchronously
// api document: https://help.aliyun.com/api/nas/describestoragepackages.html
func (client *Client) DescribeStoragePackages(request *DescribeStoragePackagesRequest) (response *DescribeStoragePackagesResponse, err error) {
	response = CreateDescribeStoragePackagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStoragePackagesWithChan invokes the nas.DescribeStoragePackages API asynchronously
// api document: https://help.aliyun.com/api/nas/describestoragepackages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStoragePackagesWithChan(request *DescribeStoragePackagesRequest) (<-chan *DescribeStoragePackagesResponse, <-chan error) {
	responseChan := make(chan *DescribeStoragePackagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStoragePackages(request)
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

// DescribeStoragePackagesWithCallback invokes the nas.DescribeStoragePackages API asynchronously
// api document: https://help.aliyun.com/api/nas/describestoragepackages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStoragePackagesWithCallback(request *DescribeStoragePackagesRequest, callback func(response *DescribeStoragePackagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStoragePackagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeStoragePackages(request)
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

// DescribeStoragePackagesRequest is the request struct for api DescribeStoragePackages
type DescribeStoragePackagesRequest struct {
	*requests.RpcRequest
	RegionId       string           `position:"Query" name:"RegionId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	UseUTCDateTime requests.Boolean `position:"Query" name:"UseUTCDateTime"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeStoragePackagesResponse is the response struct for api DescribeStoragePackages
type DescribeStoragePackagesResponse struct {
	*responses.BaseResponse
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	TotalCount int                               `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                               `json:"PageSize" xml:"PageSize"`
	PageNumber int                               `json:"PageNumber" xml:"PageNumber"`
	Packages   []DescribeStoragePackagesPackage0 `json:"Packages" xml:"Packages"`
}

type DescribeStoragePackagesPackage0 struct {
	StartTime    string `json:"StartTime" xml:"StartTime"`
	StorageType  string `json:"StorageType" xml:"StorageType"`
	Status       string `json:"Status" xml:"Status"`
	FileSystemId string `json:"FileSystemId" xml:"FileSystemId"`
	PackageId    string `json:"PackageId" xml:"PackageId"`
	ExpiredTime  string `json:"ExpiredTime" xml:"ExpiredTime"`
	Size         int64  `json:"Size" xml:"Size"`
}

// CreateDescribeStoragePackagesRequest creates a request to invoke DescribeStoragePackages API
func CreateDescribeStoragePackagesRequest() (request *DescribeStoragePackagesRequest) {
	request = &DescribeStoragePackagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeStoragePackages", "nas", "openAPI")
	return
}

// CreateDescribeStoragePackagesResponse creates a response to parse from DescribeStoragePackages response
func CreateDescribeStoragePackagesResponse() (response *DescribeStoragePackagesResponse) {
	response = &DescribeStoragePackagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
