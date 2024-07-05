package errors

import (
	"fmt"

	"github.com/charlesbases/crawler/pkg/internal/format"
)

// weberror .
type weberror struct {
	url string
	err error
}

// WebError .
func WebError(url string, err error) error {
	return &weberror{url: url, err: err}
}

// Error .
func (e *weberror) Error() string {
	return fmt.Sprintf("URL: %s, ERR: %s", e.url, e.err)
}

// steperror .
type steperror struct {
	step int
	err  error
}

// StepError .
func StepError(step int, err error) error {
	return &steperror{step: step, err: err}
}

func (s *steperror) Error() string {
	return format.Sprint(s.step, format.Space, s.err)
}
