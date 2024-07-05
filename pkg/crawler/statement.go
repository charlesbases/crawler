package crawler

import (
	"fmt"

	"github.com/charlesbases/crawler/pkg/client"
	"github.com/charlesbases/crawler/pkg/internal/errors"
	"github.com/charlesbases/crawler/pkg/internal/logger"
	"github.com/charlesbases/crawler/pkg/internal/process"
	"github.com/charlesbases/crawler/pkg/stream"
)

// errEmpty .
var errEmpty = fmt.Errorf("empty data")

// Statement .
type Statement struct {
	url    stream.Line
	steps  []process.Process
	output process.Process
}

// newStatement .
func newStatement(url stream.Line) *Statement {
	return &Statement{url: url}
}

// Next .
func (s *Statement) Next(hook stream.Hook, opts ...client.Option) *Statement {
	s.steps = append(s.steps, process.Extract(client.Configuration(opts...), hook))
	return s
}

// Output .
func (s *Statement) Output(file string, opts ...client.Option) {
	s.steps = append(s.steps, process.Output(client.Configuration(opts...), file))
}

// run .
func (s *Statement) run() error {
	links := []stream.Line{s.url}

	for i, step := range s.steps {
		logger.Debugf(i, "%s [rows: %d]", string(links[0]), len(links))

		res, err := step.Run(links)
		if err != nil {
			return errors.StepError(i, err)
		}
		// empty to extract
		if len(res) == 0 {
			logger.Error(errors.StepError(i, errEmpty))
			break
		}
		links = res
	}

	return nil
}
