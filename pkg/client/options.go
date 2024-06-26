package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/charlesbases/crawler/pkg/ioutils"
)

const (
	defaultTimeout = 30 * time.Second

	userAgentGooglebot    = `Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; UserAgentGooglebot/2.1; +http://www.google.com/bot.html) Chrome/W.X.Y.Z Safari/537.36`
	userAgentGoogleChrome = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36`
)

// Options .
type Options struct {
	client  Client
	body    io.Reader
	method  string
	header  map[string]string
	timeout time.Duration
}

// defaultOptions .
func defaultOptions() *Options {
	return &Options{
		timeout: defaultTimeout,
		method:  http.MethodGet,
		client:  defaultClient,
		header: map[string]string{
			"User-Agent": userAgentGooglebot,
		},
	}
}

// NewOptions .
func NewOptions(opts ...Option) *Options {
	o := defaultOptions()
	for _, opt := range opts {
		opt.apply(o)
	}
	return o
}

// Run .
func (o *Options) Run(url ioutils.Line) (ioutils.Body, error) {
	return o.client.Request(string(url), o)
}

// Option .
type Option interface {
	apply(o *Options)
}

// optfn .
type optfn func(o *Options)

// apply .
func (fn optfn) apply(o *Options) {
	fn(o)
}

// Body .
func Body(b io.Reader) Option {
	return optfn(func(o *Options) {
		o.body = b
	})
}

// Params .
func Params(params map[string]interface{}) Option {
	return optfn(func(o *Options) {
		data, _ := json.Marshal(&params)
		o.body = bytes.NewReader(data)
	})
}

// Method .
func Method(m string) Option {
	return optfn(func(o *Options) {
		o.method = m
	})
}

// Header .
func Header(h map[string]string) Option {
	return optfn(func(o *Options) {
		for k := range h {
			o.header[k] = h[k]
		}
	})
}

// Timeout .
func Timeout(d time.Duration) Option {
	return optfn(func(o *Options) {
		o.timeout = d
	})
}
