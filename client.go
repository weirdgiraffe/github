//
// starred.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const connectionTimeoutSec = 5
const defaultUserAgent = "star-giraffe-v0.0"
const defaultAccept = "application/vnd.github.v3+json"

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type Logger interface {
	Printf(string, ...interface{})
}

type nullLogger struct{}

func (l *nullLogger) Printf(string, ...interface{}) {}

type Client struct {
	client HttpClient
	rate   RateLimit
	log    Logger
}

func NewClient(client HttpClient) *Client {
	if client == nil {
		client = &http.Client{Timeout: connectionTimeoutSec * time.Second}
	}
	return &Client{
		client: client,
		rate:   RateLimit{RestLimit: -1},
		log:    &nullLogger{},
	}
}

func (c *Client) SetLogger(log Logger) {
	c.log = log
}

func (c *Client) NewRequest(method, url string) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set("Accept", defaultAccept)
	return req, nil

}

func (c *Client) Do(req *http.Request) (res *http.Response, err error) {
	err = c.rate.Check()
	if err != nil {
		return
	}
	c.log.Printf("(RateLimit remaining: %d) %s %s", c.rate.RestLimit, req.Method, req.URL.String())
	res, err = c.client.Do(req)
	if err != nil {
		return
	}
	err = c.rate.Update(res)
	if err != nil {
		return
	}
	return res, nil
}

func (c *Client) User() (user *User, err error) {
	var req *http.Request
	req, err = c.NewRequest("GET", "https://api.github.com/user")
	if err != nil {
		return
	}
	var res *http.Response
	res, err = c.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	user = new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	if err != nil {
		return
	}
	return user, nil
}

type RatelimitError struct {
	Timeout time.Duration
}

func (e *RatelimitError) Error() string {
	return fmt.Sprintf("X-RateLimit reached")
}

type RateLimit struct {
	RestLimit int
	Reset     time.Time
}

func (r *RateLimit) Check() error {
	if r.RestLimit == 0 && time.Now().Before(r.Reset) {
		return &RatelimitError{r.Reset.Sub(time.Now()) + time.Second}
	}
	return nil
}

func (r *RateLimit) Update(res *http.Response) (err error) {
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

type BasicAuthClient struct {
	client     HttpClient
	user, pass string
}

func NewBasicAuthClient(user, pass string) *BasicAuthClient {
	return &BasicAuthClient{
		client: &http.Client{Timeout: connectionTimeoutSec * time.Second},
		user:   user,
		pass:   pass,
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
	return a.client.Do(r)
}
