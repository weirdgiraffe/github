//
// starred.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"fmt"
	"net/http"
)

type Logger interface {
	Printf(string, ...interface{})
}

type nullLogger struct{}

func (l *nullLogger) Printf(string, ...interface{}) {}

type BasicAuthClient struct {
	RateLimitedHTTPClient
	user, pass string
}

func NewBasicAuthClient(user, pass string) *BasicAuthClient {
	return &BasicAuthClient{
		user: user,
		pass: pass,
	}
}

func (a *BasicAuthClient) Do(r *http.Request) (res *http.Response, err error) {
	if a.user == "" {
		err = fmt.Errorf("BasicAuth: user not set")
		return
	}
	if a.pass == "" {
		err = fmt.Errorf("BasicAuth: pass not set")
		return
	}
	r.SetBasicAuth(a.user, a.pass)
	return a.RateLimitedHTTPClient.Do(r)
}
