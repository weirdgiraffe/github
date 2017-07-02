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
	"strconv"
	"time"
)

const connectionTimeoutSec = 5
const defaultUserAgent = "star-giraffe-v0.0"
const defaultAccept = "application/vnd.github.v3+json"

type HttpRequestExecutor interface {
	Do(r *http.Request) (*http.Response, error)
}

type Logger interface {
	Printf(string, ...interface{})
}

type nullLogger struct{}

func (l *nullLogger) Printf(string, ...interface{}) {}

type Client struct {
	client HttpRequestExecutor
	auth   Authenticator
	rate   RateLimit
	log    Logger
}

func NewClient(auth Authenticator) *Client {
	return &Client{
		client: &http.Client{
			Timeout: connectionTimeoutSec * time.Second,
		},
		auth: auth,
		rate: RateLimit{RestLimit: -1},
		log:  &nullLogger{},
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
	if c.auth != nil {
		err = c.auth.AddAuth(req)
		if err != nil {
			return
		}
	}
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

type Authenticator interface {
	AddAuth(r *http.Request) error
}

type BasicAuth struct {
	user, pass string
}

func NewBasicAuth(user, pass string) *BasicAuth {
	return &BasicAuth{user, pass}
}

func (a *BasicAuth) AddAuth(r *http.Request) error {
	if a.user == "" {
		return fmt.Errorf("BasicAuth: user not set")
	}
	if a.pass == "" {
		return fmt.Errorf("BasicAuth: pass not set")
	}
	r.SetBasicAuth(a.user, a.pass)
	return nil
}
