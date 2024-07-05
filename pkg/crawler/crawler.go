package crawler

import (
	"runtime"

	"github.com/charlesbases/crawler/pkg/internal/logger"
	"github.com/charlesbases/crawler/pkg/stream"
)

var defaultConcurrent = runtime.NumCPU()

// Crawler .
type Crawler struct {
	concurrent int

	cores []*Statement
}

// New .
func New(opts ...Option) *Crawler {
	c := &Crawler{concurrent: defaultConcurrent}
	for _, opt := range opts {
		opt.apply(c)
	}
	return new(Crawler)
}

// Website .
func (c *Crawler) Website(url string) *Statement {
	s := newStatement(stream.Line(url))
	c.cores = append(c.cores, s)
	return s
}

// Run .
func (c *Crawler) Run() error {
	for _, core := range c.cores {
		if err := core.run(); err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil
}
