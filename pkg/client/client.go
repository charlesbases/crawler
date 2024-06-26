package client

import (
	"errors"
	"net/http"

	"github.com/charlesbases/crawler/internal/types"
	"github.com/charlesbases/crawler/pkg/ioutils"
)

var defaultClient = defaultHTTPClient()

// Client .
type Client interface {
	Request(url string, o *Options) (ioutils.Body, error)
}

// httpClient .
type httpClient struct {
	*http.Client
}

// defaultHTTPClient .
func defaultHTTPClient() Client {
	return &httpClient{
		Client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   defaultTimeout,
		},
	}
}

// Request .
func (c *httpClient) Request(url string, opts *Options) (ioutils.Body, error) {
	req, err := http.NewRequest(opts.method, url, opts.body)
	if err != nil {
		return nil, types.NewWebError(url, err)
	}

	for key := range opts.header {
		req.Header.Add(key, opts.header[key])
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, types.NewWebError(url, err)
	}

	// Status: 2xx
	if resp.StatusCode/100 == 2 {
		return ioutils.Response(resp.Body), nil
	}

	resp.Body.Close()
	return nil, types.NewWebError(url, errors.New(http.StatusText(resp.StatusCode)))
}
