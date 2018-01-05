//
// errors.go
// Copyright (C) 2018 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package github

import (
	"fmt"
	"time"
)

type ErrRatelimitReached struct {
	Timeout time.Duration
}

func (e *ErrRatelimitReached) Error() string {
	return fmt.Sprintf("RateLimit reached, will be restored in %v", e.Timeout)
}

func RateLimitReached(t time.Duration) *ErrRatelimitReached {
	return &ErrRatelimitReached{Timeout: t}
}

type ErrorUnexpectedHTTPStatus struct {
	ExpectedStatusCode int
	StatusCode         int
}

func (e *ErrorUnexpectedHTTPStatus) Error() string {
	return fmt.Sprintf("unexpected response status: expected %d got %d", e.ExpectedStatusCode, e.StatusCode)
}

func UnexpectedHTTPStatus(expected, got int) *ErrorUnexpectedHTTPStatus {
	return &ErrorUnexpectedHTTPStatus{
		ExpectedStatusCode: expected,
		StatusCode:         got,
	}
}
