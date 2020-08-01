package hello

import (
	"net/http"
	"net/url"
)

type Client struct {
	config     Config
	httpClient *http.Client
}

type Config struct {
	Address    string
	HttpClient *http.Client
}

func NewClient(config *Config) (*Client, error) {
	defConfig := DefaultConfig()

	if config.Address == "" {
		config.Address = defConfig.Address
	}

	httpClient := config.HttpClient
	if httpClient == nil {
		httpClient = defConfig.HttpClient
	}

	client := &Client{
		config:     *config,
		httpClient: httpClient,
	}

	return client, nil
}

func DefaultConfig() *Config {
	config := &Config{
		Address:    "http://127.0.0.1:8080",
		HttpClient: &http.Client{},
	}

	return config
}

type request struct {
	config *Config
	method string
	url    *url.URL
}

func (c *Client) newRequest(method, path string) (*request, error) {
	base, _ := url.Parse(c.config.Address)
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	r := &request{
		config: &c.config,
		method: method,
		url: &url.URL{
			Scheme:  base.Scheme,
			Host:    base.Host,
			Path:    u.Path,
			RawPath: u.RawPath,
		},
	}

	return r, nil
}

func (r *request) toHTTP() (*http.Request, error) {
	req, err := http.NewRequest(r.method, r.url.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
