//
// http.go
// Copyright (C) 2018 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const defaultHTTPClientTimeout = 5 * time.Second

type RequestCreator interface {
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Req  RequestCreator
	HTTP HTTPClient
}

func (c Client) User() (*User, error) {
	req, err := c.Req.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, UnexpectedHTTPStatus(http.StatusOK, res.StatusCode)
	}

	var user *User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
