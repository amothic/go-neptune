package neptune

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	endpoint   string
	httpClient *http.Client
}

func Open(urlString string) (*Client, error) {
	url, err := getEndpointURL(urlString)
	if err != nil {
		return nil, err
	}
	return &Client{endpoint: url, httpClient: http.DefaultClient}, nil
}

func getEndpointURL(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, "gremlin")
	return u.String(), nil
}

func (c *Client) do(method string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return c.httpClient.Do(req)
}

func (c *Client) post(payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do(http.MethodPost, bytes.NewBuffer(data))
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}
