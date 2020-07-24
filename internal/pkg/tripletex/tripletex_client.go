// Copyright 2020 Simen A. W. Olsen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tripletex

import (
	apiclient "github.com/bjerkio/tripletex-go/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New returns a authenticated Tripletex client
func New() (*apiclient.Tripletex, runtime.ClientAuthInfoWriter, error) {

	token, err := GetToken()
	if err != nil {
		return nil, nil, err
	}

	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.BasicAuth("0", token)

	// Fix "application/json; charset=utf-8" issue
	r.Producers["application/json; charset=utf-8"] = runtime.JSONProducer()

	return apiclient.New(r, strfmt.Default), r.DefaultAuthentication, nil
}
