package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charlesbases/crawler/pkg/internal/errors"
	"github.com/charlesbases/crawler/pkg/stream"
)

const (
	defaultTimeout = 3 * time.Second

	userAgentGooglebot    = `Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; UserAgentGooglebot/2.1; +http://www.google.com/bot.html) Chrome/W.X.Y.Z Safari/537.36`
	userAgentGoogleChrome = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36`
)

var defaultClient client = &httpClient{
	client: &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   defaultTimeout,
	},
}

// client .
type client interface {
	request(url string, o *Options) (stream.Body, error)
}

// httpClient .
type httpClient struct {
	client *http.Client
}

// request .
func (c *httpClient) request(url string, opts *Options) (stream.Body, error) {
	c.client.Timeout = opts.timeout

	req, err := http.NewRequest(opts.method, url, opts.body)
	if err != nil {
		return nil, errors.WebError(url, err)
	}

	for key := range opts.header {
		req.Header.Add(key, opts.header[key])
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WebError(url, err)
	}

	// Status: 2xx
	if resp.StatusCode/100 == 2 {
		return stream.Response(opts.hook(resp.Body)), nil
	}

	resp.Body.Close()
	return nil, errors.WebError(url, fmt.Errorf(http.StatusText(resp.StatusCode)))
}
