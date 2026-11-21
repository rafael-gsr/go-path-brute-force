// Package transport implements all network layer of the application
package transport

import (
	"fmt"
	"net/http"
	"net/url"
)

type HTTPClient struct {
	host string
}

func (h *HTTPClient) GET(path string) (*http.Response, error) {
	url, err := url.JoinPath(h.host, path)
	if err != nil {
		return nil, err
	}
	fmt.Println("path, url, host", path, url, h.host)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewHTTPClient(host string) *HTTPClient {
	return &HTTPClient{host}
}
