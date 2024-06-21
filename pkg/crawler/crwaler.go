package crawler

import (
	"time"

	"github.com/charlesbases/crawler/pkg/sender"
)

// Crwaler .
type Crwaler struct {
	*options

	cores []*sender.Processor
}

// options .
type options struct {
	concurrency int
	timeout     time.Duration
}

// Option .
type Option interface {
	apply(o *options)
}

// optfn .
type optfn func(o *options)

// apply .
func (fn optfn) apply(o *options) {
	fn(o)
}

// Concurrency .
func Concurrency(c int) Option {
	return optfn(func(o *options) {
		o.concurrency = c
	})
}

// Timeout .
func Timeout(d time.Duration) Option {
	return optfn(func(o *options) {
		o.timeout = d
	})
}

// New .
func New(opts ...Option) *Crwaler {
	c := newDefaultCrawler()
	for _, opt := range opts {
		opt.apply(c.options)
	}
	return c
}

// newDefaultCrawler .
func newDefaultCrawler() *Crwaler {
	return &Crwaler{
		options: &options{
			concurrency: 1,
			timeout:     3 * time.Second,
		},
	}
}

// Processor .
func (c *Crwaler) Processor(url string, opts ...sender.Option) *sender.Processor {
	p := sender.NewProcessor(url, opts...)
	c.cores = append(c.cores, p)
	return p
}

// Run .
func (c *Crwaler) Run() error {
	// TODO
	panic("todo")
}
