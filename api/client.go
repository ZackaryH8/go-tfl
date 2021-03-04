package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	// RootURL is the start of the url as used in each API call
	RootURL = "https://api.tfl.gov.uk"
)

// Client is the main struct which contains all the methods for go-tfl
type Client struct {
	*Config
}

// Config contains all the configurable for the client
type Config struct {
	id         string
	key        string
	rootURL    string
	httpClient *http.Client
}

// Option defines a way of configuring the client
type Option func(*Config)

// NewClient configures and returns a new client for accessing the API
func NewClient(options ...Option) (*Client, error) {
	if len(options) < 1 {
		return nil, ErrNoOptionsPassed
	}

	var config Config
	for _, option := range options {
		option(&config)
	}

	return &Client{
		&config,
	}, nil
}

// DefaultOptions will set the registered id and key. For is the minimum that
// is required for regular operation
func DefaultOptions(id, key string) Option {
	return func(c *Config) {
		c.id = id
		c.key = key
		c.rootURL = RootURL
		c.httpClient = http.DefaultClient
	}
}

// CustomHTTPClientOption enables a custom http client to be used for all calls
// to the TFL api. This is only useful if you want to updating settings for the
// client
func CustomHTTPClientOption(client *http.Client) Option {
	return func(c *Config) {
		c.httpClient = client
	}
}

func (c *Client) getJSON(url string, dst interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dst)
}

// getAuthArgs returns the auth arguments to be appended to a URL
func (c *Client) getAuthArgs() string {
	return fmt.Sprintf("app_id=%s&app_key=%s", c.id, c.key)
}

func (c *Client) getEndpointURL(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.rootURL, endpoint)
}

func (c *Client) arrayToCSV(array []string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}

// Errors
var (
	ErrNoOptionsPassed = errors.New("no options passed, requires at least one option")
)
