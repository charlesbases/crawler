package sender

import (
	"bytes"
	"encoding/json"
	"io"
)

/*
UserAgentGooglebot: Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; UserAgentGooglebot/2.1; +http://www.google.com/bot.html) Chrome/W.X.Y.Z Safari/537.36
GoogleChrome: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
*/

// Processor .
type Processor struct {
	url    string
	sender Sender
	header map[string]string
	body   io.Reader
	expr   string
}

// Option .
type Option interface {
	apply(s *Processor)
}

// optfn .
type optfn func(s *Processor)

// apply .
func (f optfn) apply(s *Processor) {
	f(s)
}

// NewProcessor .
func NewProcessor(url string, opts ...Option) *Processor {
	s := newDefaultProcessor(url)
	for _, opt := range opts {
		opt.apply(s)
	}
	return s
}

// newDefaultProcessor .
func newDefaultProcessor(url string) *Processor {
	return &Processor{
		url:    url,
		sender: defaultClientSender,
		header: map[string]string{
			"User-Agent": userAgentGooglebot,
		},
	}
}

// Express regexp express
func Express(expr string) Option {
	return optfn(func(s *Processor) {

	})
}

// Header .
func Header(key string, val string) Option {
	return optfn(func(s *Processor) {
		s.header[key] = val
	})
}

// UserAgentGooglebot .
func UserAgentGooglebot() Option {
	return optfn(func(s *Processor) {
		s.header["User-Agent"] = userAgentGooglebot
	})
}

// UserAgentGoogleChrome .
func UserAgentGoogleChrome() Option {
	return optfn(func(s *Processor) {
		s.header["User-Agent"] = userAgentGoogleChrome
	})
}

// Params .
func Params(params map[string]interface{}) Option {
	return optfn(func(s *Processor) {
		data, _ := json.Marshal(&params)
		s.body = bytes.NewReader(data)
	})
}

// HTTPSender .
func HTTPSender() Option {
	return optfn(func(s *Processor) {
		s.sender = defaultClientSender
	})
}
