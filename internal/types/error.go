package types

import (
	"fmt"

	"github.com/charlesbases/crawler/internal/display"
)

// WebError .
type WebError struct {
	url string
	err error
}

// NewWebError .
func NewWebError(url string, err error) error {
	return &WebError{url: url, err: err}
}

// Error .
func (e *WebError) Error() string {
	return fmt.Sprintf("URL: %s, ERR: %s", e.url, e.err)
}

// StepError .
type StepError struct {
	step int
	err  error
}

// NewStepError .
func NewStepError(step int, err error) error {
	return &StepError{step: step, err: err}
}

func (s *StepError) Error() string {
	return display.Sprint(s.step, display.Tabulator, s.err)
}
