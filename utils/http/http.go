package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HttpClient is the interface that wraps the basic Get method.
type HttpClient interface {
	Get(c context.Context, url string) (*http.Response, error)
}

// httpClient is an implementation of HttpClient interface.
type httpClient struct {
	Client     *http.Client
	MaxRetries int
	RetryDelay time.Duration
	TimeOut    time.Duration
}

// NewHTTPClient returns a new HttpClient.
func NewHTTPClient(maxRetries int, retryInterval time.Duration, timeout time.Duration) HttpClient {
	return &httpClient{
		Client:     &http.Client{},
		MaxRetries: maxRetries,
		RetryDelay: retryInterval,
		TimeOut:    timeout,
	}
}

// doRequest is a helper function to do http request.
func (hc *httpClient) doRequest(req *http.Request) (resp *http.Response, err error) {
	i := 0
	for ; i < hc.MaxRetries; i++ {
		resp, err = hc.Client.Do(req)
		if err != nil {
			time.Sleep(hc.RetryDelay)
			continue
		}
		return
	}
	if i == hc.MaxRetries {
		err = fmt.Errorf("max retries exceeded")
	}

	return
}

// Get is a helper function to do http get request.
func (hc *httpClient) Get(c context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(c, hc.TimeOut)
	defer cancel()
	req = req.WithContext(ctx)
	return hc.doRequest(req)
}
