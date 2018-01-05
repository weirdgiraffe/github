//
// requests.go
// Copyright (C) 2018 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const UserAgent = "star-giraffe-v0.0"

type DefaultRequestsCreator struct{}

func (c DefaultRequestsCreator) addHeaders(req *http.Request) (*http.Request, error) {
	req.Header.Set("User-Agent", "star-giraffe-v0.0")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	return req, nil
}

func (c DefaultRequestsCreator) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	return c.addHeaders(req)
}

type RateLimitedHTTPClient struct {
	c         *http.Client
	RestLimit int
	Reset     time.Time
}

func (r *RateLimitedHTTPClient) updateLimits(res *http.Response) error {
	var err error
	if s := res.Header.Get("X-RateLimit-Remaining"); s != "" {
		r.RestLimit, err = strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("update X-RateLimit-Remaining: %v", err)
		}
	}
	if s := res.Header.Get("X-RateLimit-Reset"); s != "" {
		var t int64
		t, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fmt.Errorf("update X-RateLimit-Reset %v", err)
		}
		r.Reset = time.Unix(t, 0)
	}
	return nil
}

func (r *RateLimitedHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if r.c == nil {
		r.c = &http.Client{Timeout: defaultHTTPClientTimeout}
	}
	if r.RestLimit == 0 && time.Now().Before(r.Reset) {
		return nil, RateLimitReached(r.Reset.Sub(time.Now()) + time.Second)
	}

	res, err := r.c.Do(req)
	if err != nil {
		return nil, err
	}

	err = r.updateLimits(res)
	if err != nil {
		res.Body.Close()
		return nil, err
	}
	return res, err
}
