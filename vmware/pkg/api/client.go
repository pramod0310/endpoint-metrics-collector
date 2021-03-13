package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client interface {
	Get(url string) (*Response, error)
	GetURL(path string) string
	GetTimeout() time.Duration
}

type Response struct {
	URL          string
	ResponseTime time.Duration
	StatusCode   int
}

type client struct {
	baseURL string
	scheme  string //either http or https
	timeout time.Duration
}

func NewHttpClient(baseURL, scheme string, timeout time.Duration) Client {
	return &client{baseURL: baseURL, scheme: scheme, timeout: timeout}
}

func (c *client) GetURL(targetPath string) string {
	return fmt.Sprintf("%s://%s/%s", c.scheme, c.baseURL, targetPath)
}

func (c *client) GetTimeout() time.Duration {
	return c.timeout
}

//methodType can be GET,POST,DELETE
func (c *client) getHttpResponse(methodType, url string) (*Response, error) {

	request, err := http.NewRequest(methodType, url, strings.NewReader(""))
	httpClient := &http.Client{Timeout: c.timeout}
	start := time.Now()
	resp, err := httpClient.Do(request)
	timeTaken := time.Since(start)

	if err != nil {
		return nil, err
	}

	return &Response{url, timeTaken, resp.StatusCode}, nil
}

func (c *client) Get(path string) (*Response, error) {
	return c.getHttpResponse("GET", path)
}
