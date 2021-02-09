// Package elasticsearch is a wrapper of the original "github.com/elastic/go-elasticsearch".
package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// Client is elasticsearch client.
type Client struct {
	es *elasticsearch.Client
}

// New to create new elasticsearch client.
func New(addresses []string, username string, password string) (*Client, error) {
	return NewWithConfig(elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
	})
}

// NewWithConfig to create new elasticsearch client with config.
func NewWithConfig(cfg elasticsearch.Config) (*Client, error) {
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// Connection check.
	if err = isError(es.Info()); err != nil {
		return nil, err
	}
	return &Client{es: es}, nil
}

// NewDefault to create default elasticsearch client.
func NewDefault() (*Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &Client{es: es}, nil
}

// Send to send data to elasticsearch.
func (c *Client) Send(key string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index:   key,
		Body:    strings.NewReader(string(d)),
		Refresh: "true",
	}
	return isError(req.Do(context.Background(), c.es))
}

func isError(res *esapi.Response, err error) error {
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}
