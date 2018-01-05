//
// paginator.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const defaultPerPage = 100

type Paginator struct {
	PerPage int

	req        *http.Request
	client     *Client
	next, last string
}

func NewPaginator(c *Client, req *http.Request) *Paginator {
	return &Paginator{
		PerPage: defaultPerPage,
		req:     req,
		client:  c,
	}
}

func (p *Paginator) Next() (res *http.Response, err error) {
	if p.req.URL.String() == p.last && p.last != "" {
		return nil, io.EOF
	}
	if p.next != "" {
		p.req.URL, err = url.Parse(p.next)
		if err != nil {
			return
		}
	}

	q := p.req.URL.Query()
	q.Set("per_page", strconv.Itoa(p.PerPage))
	p.req.URL.RawQuery = q.Encode()

	res, err = p.client.HTTP.Do(p.req)
	if err != nil {
		return nil, err
	}
	p.update(res)
	return res, nil
}

func (p *Paginator) update(res *http.Response) {
	link := strings.Split(res.Header.Get("Link"), ",")
	for i := range link {
		part := strings.Split(link[i], ";")
		if len(part) == 2 {
			if strings.TrimSpace(part[1]) == `rel="next"` {
				raw := strings.TrimSpace(part[0])
				p.next = raw[1 : len(raw)-1]
			}
			if strings.TrimSpace(part[1]) == `rel="last"` {
				raw := strings.TrimSpace(part[0])
				p.last = raw[1 : len(raw)-1]
			}
		}
	}
}
