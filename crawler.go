package crawler

import (
	"runtime"

	"github.com/charlesbases/crawler/internal/logger"
)

var (
	debugMode  = false
	concurrent = runtime.NumCPU()
)

// DebugMode .
func DebugMode() {
	debugMode = true
}

// Concurrent .
func Concurrent(n int) {
	concurrent = n
}

// Crawler .
type Crawler struct {
	cores []*Statement
}

// New .
func New() *Crawler {
	return new(Crawler)
}

// Website .
func (c *Crawler) Website(url string) *Statement {
	s := newStatement(url)
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
