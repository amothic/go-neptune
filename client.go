package neptune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	endpoint   string
}

func Open(urlString string) (*Client, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	return &Client{endpoint: url.String(), httpClient: http.DefaultClient}, nil
}

func (c *Client) do(method string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	return c.checkResponse(resp, err)
}

func (c *Client) post(payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do(http.MethodPost, bytes.NewReader(data))
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, fmt.Errorf("error calling the api endpoint: %w", err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusPartialContent {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return resp, fmt.Errorf("response error: status:%q, body:%q", resp.Status, body)
	}
	return resp, nil
}
