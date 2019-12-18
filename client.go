package neptune

import (
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
