package http_client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	//client3second is httpclient with 3 second timeout set.
	client3second *http.Client
)

const (
	defaultTimeout           = 3 * time.Second
	loggerNameRequest        = "http-request"
	loggerNameResponse       = "http-response"
	loggerNameResponseHeader = "http-response-header"
)

// Client http implementation wrapper
type Client struct {
	restyClient *resty.Client
}

func NewClient() *Client {
	return &Client{
		restyClient: resty.New(),
	}
}

func (c *Client) Get() *resty.Client {
	return c.restyClient
}

// Request ...
func Request(method, url string, queries, headers map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// set queries string
	q := req.URL.Query()
	for k, v := range queries {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// set headers
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := client3second.Do(req)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
