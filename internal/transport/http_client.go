// Package transport implements all network layer of the application
package transport

import (
	"errors"
	"net/http"
	"net/url"

	globaltypes "github.com/rafael-gsr/go-path-brute-force/internal/global_types"
)

type HTTPClient struct {
	host    string
	cookies map[string]string
	headers map[string]string
}

func (h *HTTPClient) GET(path string) (*http.Response, error) {
	url, err := url.JoinPath(h.host, path)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range h.headers {
		request.Header.Add(key, value)
	}

	for key, value := range h.cookies {
		cookie := &http.Cookie{}
		cookie.Name = key
		cookie.Value = value
		request.AddCookie(cookie)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewHTTPClient(request globaltypes.BruteForceRequest) (*HTTPClient, error) {
	if request.Host() == "" {
		return nil, errors.New("no host found")
	}

	return &HTTPClient{request.Host(), request.Cookies(), request.Headers()}, nil
}
