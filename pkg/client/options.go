package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/charlesbases/crawler/pkg/internal/util"
	"github.com/charlesbases/crawler/pkg/stream"
)

var defaultHook = func(r io.ReadCloser) io.ReadCloser {
	return r
}

// Options .
type Options struct {
	client  client
	body    io.Reader
	method  string
	header  map[string]string
	timeout time.Duration
	// handle with response
	hook func(r io.ReadCloser) io.ReadCloser
}

// Run .
func (o *Options) Run(url stream.Line) (stream.Body, error) {
	return o.client.request(string(url), o)
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

// Configuration .
func Configuration(opts ...Option) *Options {
	o := defaultOptions()
	for _, opt := range opts {
		opt.apply(o)
	}
	return o
}

// defaultOptions .
func defaultOptions() *Options {
	return &Options{
		hook:    defaultHook,
		timeout: defaultTimeout,
		method:  http.MethodGet,
		client:  defaultClient,
		header: map[string]string{
			"User-Agent": userAgentGooglebot,
		},
	}
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

// Debug write http.Response.Body to file
func Debug(file string) Option {
	return optfn(func(o *Options) {
		if absfile, err := util.CreateDir(file); err != nil {
			return
		} else {
			file, err := os.Create(absfile)
			if err != nil {
				return
			}

			o.hook = func(r io.ReadCloser) io.ReadCloser {
				io.Copy(file, r)
				return file
			}
		}
	})
}
