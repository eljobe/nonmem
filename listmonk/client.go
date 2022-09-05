package listmonk

import (
	"net/http"
	"time"
)

type client struct {
	hc     http.Client
	apiUrl string
}

func newClient() *client {
	return newClientAt("http://127.0.0.1:9000/api")
}

func newClientAt(url string) *client {
	return &client{
		hc: http.Client{
			Timeout: time.Second * 10,
		},
		apiUrl: url,
	}
}

func (c *client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.apiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
