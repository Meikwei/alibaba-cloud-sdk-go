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

// ListRecordingOfDualTrack invokes the ccc.ListRecordingOfDualTrack API synchronously
// api document: https://help.aliyun.com/api/ccc/listrecordingofdualtrack.html
func (client *Client) ListRecordingOfDualTrack(request *ListRecordingOfDualTrackRequest) (response *ListRecordingOfDualTrackResponse, err error) {
	response = CreateListRecordingOfDualTrackResponse()
	err = client.DoAction(request, response)
	return
}

// ListRecordingOfDualTrackWithChan invokes the ccc.ListRecordingOfDualTrack API asynchronously
// api document: https://help.aliyun.com/api/ccc/listrecordingofdualtrack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRecordingOfDualTrackWithChan(request *ListRecordingOfDualTrackRequest) (<-chan *ListRecordingOfDualTrackResponse, <-chan error) {
	responseChan := make(chan *ListRecordingOfDualTrackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRecordingOfDualTrack(request)
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

// ListRecordingOfDualTrackWithCallback invokes the ccc.ListRecordingOfDualTrack API asynchronously
// api document: https://help.aliyun.com/api/ccc/listrecordingofdualtrack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRecordingOfDualTrackWithCallback(request *ListRecordingOfDualTrackRequest, callback func(response *ListRecordingOfDualTrackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRecordingOfDualTrackResponse
		var err error
		defer close(result)
		response, err = client.ListRecordingOfDualTrack(request)
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

// ListRecordingOfDualTrackRequest is the request struct for api ListRecordingOfDualTrack
type ListRecordingOfDualTrackRequest struct {
	*requests.RpcRequest
	CallingNumber string           `position:"Query" name:"CallingNumber"`
	AgentId       string           `position:"Query" name:"AgentId"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	CalledNumber  string           `position:"Query" name:"CalledNumber"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	StopTime      requests.Integer `position:"Query" name:"StopTime"`
	ConnectId     string           `position:"Query" name:"ConnectId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// ListRecordingOfDualTrackResponse is the response struct for api ListRecordingOfDualTrack
type ListRecordingOfDualTrackResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Recordings     Recordings `json:"Recordings" xml:"Recordings"`
}

// CreateListRecordingOfDualTrackRequest creates a request to invoke ListRecordingOfDualTrack API
func CreateListRecordingOfDualTrackRequest() (request *ListRecordingOfDualTrackRequest) {
	request = &ListRecordingOfDualTrackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListRecordingOfDualTrack", "", "")
	return
}

// CreateListRecordingOfDualTrackResponse creates a response to parse from ListRecordingOfDualTrack response
func CreateListRecordingOfDualTrackResponse() (response *ListRecordingOfDualTrackResponse) {
	response = &ListRecordingOfDualTrackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
