//
// client_test.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/weirdgiraffe/github/mock"
)

func TestRequestDefaultHeaders(t *testing.T) {
	c := NewClient(nil)
	verb := []string{"GET", "PUT", "POST", "DELETE"}
	url := "https://api.github.com"
	for i := range verb {
		req, err := c.NewRequest(verb[i], url)
		if err != nil {
			t.Fatal("Failed NewRequest(): %v", err)
		}
		if req.Header.Get("User-Agent") != defaultUserAgent {
			t.Errorf(
				"Unexpected User-Agent header: %s != %s",
				defaultUserAgent,
				req.Header.Get("User-Agent"),
			)
		}
		if req.Header.Get("Accept") != defaultAccept {
			t.Errorf(
				"Unexpected Accept header: %s != %s",
				defaultAccept,
				req.Header.Get("Accept"),
			)
		}
	}
}

func TestBasicAuth(t *testing.T) {
	mockHTTP := &github_mock.HttpClient{
		ActualDo: func(r *http.Request) (*http.Response, error) {
			if r.Header.Get("Authorization") == "" {
				t.Error("BasicAuth not set")
			}
			return &http.Response{
				Header: make(http.Header),
				Body:   ioutil.NopCloser(bytes.NewBufferString("Hello World")),
			}, nil
		},
	}
	client := &BasicAuthClient{
		client: mockHTTP,
		user:   "hello",
		pass:   "world",
	}
	c := NewClient(client)
	verb := []string{"GET", "PUT", "POST", "DELETE"}
	url := "https://api.github.com"
	for i := range verb {
		req, err := c.NewRequest(verb[i], url)
		if err != nil {
			t.Fatal("Failed NewRequest(): %v", err)
		}
		_, _ = c.Do(req)
	}
}

func TestRateLimit(t *testing.T) {
	res := &http.Response{
		Header: make(http.Header),
	}
	var tc = []struct {
		restLimit  int
		reset      time.Time
		checkError bool
	}{
		{100, time.Now().Add(1 * time.Hour), false},
		{0, time.Now().Add(1 * time.Hour), true},
	}
	for i := range tc {
		r := RateLimit{}
		err := r.Check()
		if err != nil {
			t.Fatalf("testcase[%2d] New Ratelimit check: %v", i, err)
		}
		res.Header.Set(
			"X-RateLimit-Remaining",
			fmt.Sprintf("%d", tc[i].restLimit),
		)
		res.Header.Set(
			"X-RateLimit-Reset",
			fmt.Sprintf("%d", tc[i].reset.Unix()),
		)
		err = r.Update(res)
		if err != nil {
			t.Fatalf("testcase[%2d] Ratelimit update: %v", i, err)
		}
		err = r.Check()
		if tc[i].checkError {
			if err == nil {
				t.Errorf("testcase[%2d] No check error, but expected", i)
			} else {
				if e := err.(*RatelimitError); e == nil {
					t.Errorf("testcase[%2d] Unexpected error: %v", i, err)
				}
				if e := err.(*RatelimitError); e != nil {
					if time.Now().Add(e.Timeout).Before(tc[i].reset) {
						t.Errorf("testcase[%2d] Unexpected timeout: %v", i, e.Timeout)
					}
				}
			}
		} else {
			if err != nil {
				t.Errorf("testcase[%2d] Check error, but unexpected: %v", i, err)
			}
		}
		if r.RestLimit != tc[i].restLimit {
			t.Errorf(
				"testcase[%2d] Ratelimit.RestLimit: %d != %d",
				i, tc[i].restLimit, r.RestLimit,
			)
		}
		if r.Reset.Unix() != tc[i].reset.Unix() {
			t.Errorf(
				"testcase[%2d] Ratelimit.Reset %v != %v",
				i, tc[i].reset, r.Reset,
			)
		}
	}
}

func TestClientDo(t *testing.T) {
	mockHTTP := &github_mock.HttpClient{
		ActualDo: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{
				Header: make(http.Header),
				Body:   ioutil.NopCloser(bytes.NewBufferString("Hello World")),
			}
			res.Header.Set("X-RateLimit-Remaining", "1")
			res.Header.Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()))
			return res, nil
		},
	}
	c := NewClient(mockHTTP)
	req, err := c.NewRequest("GET", "https://api.github.com")
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	res, err := c.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	res.Body.Close()
}
