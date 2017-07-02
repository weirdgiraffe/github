//
// paginator_test.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/weirdgiraffe/github/mock"
)

const paginatorMaxPages = 10

var paginatorHttpMock = &github_mock.HttpRequestExecutor{
	ActualDo: func(req *http.Request) (res *http.Response, err error) {
		res = &http.Response{
			Header: make(http.Header),
			Body:   ioutil.NopCloser(bytes.NewBufferString("Hello World")),
		}
		u := *req.URL
		q := u.Query()
		p := 1
		if pstr := q.Get("page"); pstr != "" {
			p, err = strconv.Atoi(pstr)
			if err != nil {
				return
			}
		}
		if p < paginatorMaxPages {
			p = p + 1
		}
		q.Set("page", strconv.Itoa(p))
		u.RawQuery = q.Encode()
		lnext := fmt.Sprintf("<%s>; rel=\"next\"", u.String())
		q.Set("page", strconv.Itoa(paginatorMaxPages))
		u.RawQuery = q.Encode()
		llast := fmt.Sprintf("<%s>; rel=\"last\"", u.String())
		res.Header.Set("Link", fmt.Sprintf("%s, %s", lnext, llast))
		return res, nil
	},
}

var paginatorClient = func() *Client {
	c := NewClient(nil)
	c.client = paginatorHttpMock
	return c
}()

func TestPaginatorIteration(t *testing.T) {
	req, err := paginatorClient.NewRequest("GET", "https://api.github.com")
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}

	p := NewPaginator(paginatorClient, req)
	for i := 0; i < paginatorMaxPages; i++ {
		res, err := p.Next()
		if err != nil {
			t.Fatalf("Next(): %v", err)
		}

		if p.next == "" {
			t.Error("Link next not set")
		}
		unext, err := url.Parse(p.next)
		if err != nil {
			t.Fatalf("next page: %v", err)
		}
		if pstr := unext.Query().Get("page"); pstr != "" {
			pn, err := strconv.Atoi(pstr)
			if err != nil {
				t.Errorf("next page: %v", err)
			} else {
				ePage := i + 2
				if i == paginatorMaxPages-1 {
					// last page
					ePage = paginatorMaxPages
				}
				if pn != ePage {
					t.Errorf("unexpected next page: %d != %d", ePage, pn)
				}
			}
		}

		if p.last == "" {
			t.Error("Link last not set")
		}
		ulast, err := url.Parse(p.last)
		if err != nil {
			t.Fatalf("last page: %v", err)
		}
		if pstr := ulast.Query().Get("page"); pstr != "" {
			pn, err := strconv.Atoi(pstr)
			if err != nil {
				t.Errorf("last page: %v", err)
			} else {
				if pn != paginatorMaxPages {
					t.Errorf("unexpected last page: %d != %d", paginatorMaxPages, pn)
				}
			}
		}
		res.Body.Close()
	}

	_, err = p.Next()
	if err != io.EOF {
		t.Errorf("No io.EOF after last page reached")
	}
}
