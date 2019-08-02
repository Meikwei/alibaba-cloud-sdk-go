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

// PredictiveRecordFailure invokes the ccc.PredictiveRecordFailure API synchronously
// api document: https://help.aliyun.com/api/ccc/predictiverecordfailure.html
func (client *Client) PredictiveRecordFailure(request *PredictiveRecordFailureRequest) (response *PredictiveRecordFailureResponse, err error) {
	response = CreatePredictiveRecordFailureResponse()
	err = client.DoAction(request, response)
	return
}

// PredictiveRecordFailureWithChan invokes the ccc.PredictiveRecordFailure API asynchronously
// api document: https://help.aliyun.com/api/ccc/predictiverecordfailure.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PredictiveRecordFailureWithChan(request *PredictiveRecordFailureRequest) (<-chan *PredictiveRecordFailureResponse, <-chan error) {
	responseChan := make(chan *PredictiveRecordFailureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PredictiveRecordFailure(request)
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

// PredictiveRecordFailureWithCallback invokes the ccc.PredictiveRecordFailure API asynchronously
// api document: https://help.aliyun.com/api/ccc/predictiverecordfailure.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PredictiveRecordFailureWithCallback(request *PredictiveRecordFailureRequest, callback func(response *PredictiveRecordFailureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PredictiveRecordFailureResponse
		var err error
		defer close(result)
		response, err = client.PredictiveRecordFailure(request)
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

// PredictiveRecordFailureRequest is the request struct for api PredictiveRecordFailure
type PredictiveRecordFailureRequest struct {
	*requests.RpcRequest
	CallId             string           `position:"Query" name:"CallId"`
	ActualTime         requests.Integer `position:"Query" name:"ActualTime"`
	CallingNumber      string           `position:"Query" name:"CallingNumber"`
	InstanceId         string           `position:"Query" name:"InstanceId"`
	DispositionCode    string           `position:"Query" name:"DispositionCode"`
	CalledNumber       string           `position:"Query" name:"CalledNumber"`
	TaskId             string           `position:"Query" name:"TaskId"`
	CabInstanceId      string           `position:"Query" name:"CabInstanceId"`
	CabInstanceOwnerId requests.Integer `position:"Query" name:"CabInstanceOwnerId"`
}

// PredictiveRecordFailureResponse is the response struct for api PredictiveRecordFailure
type PredictiveRecordFailureResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreatePredictiveRecordFailureRequest creates a request to invoke PredictiveRecordFailure API
func CreatePredictiveRecordFailureRequest() (request *PredictiveRecordFailureRequest) {
	request = &PredictiveRecordFailureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "PredictiveRecordFailure", "", "")
	return
}

// CreatePredictiveRecordFailureResponse creates a response to parse from PredictiveRecordFailure response
func CreatePredictiveRecordFailureResponse() (response *PredictiveRecordFailureResponse) {
	response = &PredictiveRecordFailureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
