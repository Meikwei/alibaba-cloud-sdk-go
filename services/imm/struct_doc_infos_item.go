package imm

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

// DocInfosItem is a nested struct in imm response
type DocInfosItem struct {
	UniqueId     string `json:"UniqueId" xml:"UniqueId"`
	SrcUri       string `json:"SrcUri" xml:"SrcUri"`
	Name         string `json:"Name" xml:"Name"`
	ContentType  string `json:"ContentType" xml:"ContentType"`
	LastModified string `json:"LastModified" xml:"LastModified"`
	Size         int    `json:"Size" xml:"Size"`
	PageNum      int    `json:"PageNum" xml:"PageNum"`
	CustomKey1   string `json:"CustomKey1" xml:"CustomKey1"`
	CustomKey2   string `json:"CustomKey2" xml:"CustomKey2"`
	CustomKey3   string `json:"CustomKey3" xml:"CustomKey3"`
	CustomKey4   string `json:"CustomKey4" xml:"CustomKey4"`
	CustomKey5   string `json:"CustomKey5" xml:"CustomKey5"`
	CustomKey6   string `json:"CustomKey6" xml:"CustomKey6"`
}
