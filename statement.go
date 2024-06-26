package crawler

import (
	"github.com/charlesbases/crawler/internal/logger"
	"github.com/charlesbases/crawler/internal/types"
	"github.com/charlesbases/crawler/pkg/client"
	"github.com/charlesbases/crawler/pkg/ioutils"
)

// Statement .
type Statement struct {
	url   string
	steps []process
}

// newStatement .
func newStatement(url string) *Statement {
	return &Statement{url: url}
}

// Next .
func (s *Statement) Next(a ioutils.Handle, opts ...client.Option) *Statement {
	s.steps = append(s.steps, newExtractProcess(client.NewOptions(opts...), a))
	return s
}

// Save .
func (s *Statement) Save(f string, opts ...client.Option) {
	s.steps = append(s.steps, newOutputProcess(client.NewOptions(opts...), f))
}

// run extract from http.Response
func (s *Statement) run() error {
	links := []ioutils.Line{ioutils.Line(s.url)}

	for i, step := range s.steps {
		logger.Debug(i, string(links[0]))

		res, err := step.run(links)
		if err != nil {
			return types.NewStepError(i, err)
		}
		links = res
	}

	return nil
}
