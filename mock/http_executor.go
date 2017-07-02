//
// http_executor.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github_mock

import "net/http"

type HttpClient struct {
	ActualDo func(*http.Request) (*http.Response, error)
}

func (e *HttpClient) Do(r *http.Request) (*http.Response, error) {
	return e.ActualDo(r)
}
