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

// GetCallMeasureSummaryReport invokes the ccc.GetCallMeasureSummaryReport API synchronously
// api document: https://help.aliyun.com/api/ccc/getcallmeasuresummaryreport.html
func (client *Client) GetCallMeasureSummaryReport(request *GetCallMeasureSummaryReportRequest) (response *GetCallMeasureSummaryReportResponse, err error) {
	response = CreateGetCallMeasureSummaryReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetCallMeasureSummaryReportWithChan invokes the ccc.GetCallMeasureSummaryReport API asynchronously
// api document: https://help.aliyun.com/api/ccc/getcallmeasuresummaryreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCallMeasureSummaryReportWithChan(request *GetCallMeasureSummaryReportRequest) (<-chan *GetCallMeasureSummaryReportResponse, <-chan error) {
	responseChan := make(chan *GetCallMeasureSummaryReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCallMeasureSummaryReport(request)
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

// GetCallMeasureSummaryReportWithCallback invokes the ccc.GetCallMeasureSummaryReport API asynchronously
// api document: https://help.aliyun.com/api/ccc/getcallmeasuresummaryreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCallMeasureSummaryReportWithCallback(request *GetCallMeasureSummaryReportRequest, callback func(response *GetCallMeasureSummaryReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCallMeasureSummaryReportResponse
		var err error
		defer close(result)
		response, err = client.GetCallMeasureSummaryReport(request)
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

// GetCallMeasureSummaryReportRequest is the request struct for api GetCallMeasureSummaryReport
type GetCallMeasureSummaryReportRequest struct {
	*requests.RpcRequest
	IntervalType string           `position:"Query" name:"IntervalType"`
	Month        requests.Integer `position:"Query" name:"Month"`
	Year         requests.Integer `position:"Query" name:"Year"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Day          requests.Integer `position:"Query" name:"Day"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
}

// GetCallMeasureSummaryReportResponse is the response struct for api GetCallMeasureSummaryReport
type GetCallMeasureSummaryReportResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Success        bool          `json:"Success" xml:"Success"`
	Code           string        `json:"Code" xml:"Code"`
	Message        string        `json:"Message" xml:"Message"`
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	SummaryReport  SummaryReport `json:"SummaryReport" xml:"SummaryReport"`
	NumberReports  NumberReports `json:"NumberReports" xml:"NumberReports"`
}

// CreateGetCallMeasureSummaryReportRequest creates a request to invoke GetCallMeasureSummaryReport API
func CreateGetCallMeasureSummaryReportRequest() (request *GetCallMeasureSummaryReportRequest) {
	request = &GetCallMeasureSummaryReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetCallMeasureSummaryReport", "", "")
	return
}

// CreateGetCallMeasureSummaryReportResponse creates a response to parse from GetCallMeasureSummaryReport response
func CreateGetCallMeasureSummaryReportResponse() (response *GetCallMeasureSummaryReportResponse) {
	response = &GetCallMeasureSummaryReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
