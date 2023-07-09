package request

import (
	"io"
	"net/http"
	"strings"
)

type Client struct {
	client  *http.Client
	url     string
	method  string
	headers map[string]string
	body    string
}

func (c *Client) SetHeaders(headers map[string]string) {
	c.headers = headers
}

func (c *Client) SetBody(body string) {
	c.body = body
}

func (c *Client) Get(url string) string {
	c.method = "GET"
	c.url = url
	return c.Send()
}
func (c *Client) Post(url string) string {
	c.method = "POST"
	c.url = url
	return c.Send()
}

func NewGetClient(url string, headers map[string]string) *Client {
	return NewClient(url, "GET", headers, "")
}
func NewPostClient(url string, headers map[string]string, body string) *Client {
	return NewClient(url, "POST", headers, body)
}

func SetHeaders(client *Client, headers map[string]string) {
	client.headers = headers
}
func SetBody(client *Client, body string) {
	client.body = body
}

func NewClient(url string, method string, headers map[string]string, body string) *Client {
	return &Client{
		client:  &http.Client{},
		url:     url,
		method:  method,
		headers: headers,
		body:    body,
	}
}

func (c *Client) Send() string {
	req, err := http.NewRequest(c.method, c.url, strings.NewReader(c.body))
	if err != nil {
		return err.Error()
	}
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	all, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(all)
}
