/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"fmt"
	"net/http"

	httpUtils "inspur.com/cloudware/util/http"
)

func main() {
	DoPostTest()
	DoGetTest()
}

func DoPostTest() {
	client := httpUtils.NewHTTPClient()
	reqbody_str := "{\"age\":18,\"name\":\"Nicholas\"}"

	content := []byte(reqbody_str)
	req, err := httpUtils.BuildRequest(http.MethodPost, "http://127.0.0.1:3000/api/v1/ping", bytes.NewReader(content), "")
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := httpUtils.SendRequest(req, client)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func DoGetTest() {
	client := httpUtils.NewHTTPClient()

	req, err := httpUtils.BuildRequest(http.MethodGet, "http://127.0.0.1:3000/success", nil, "")
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := httpUtils.SendRequest(req, client)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}
