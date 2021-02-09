// Package redis is a wrapper of the original "github.com/go-redis/redis" library.
//
// Only contains basic get, set, delete, and close methods.
// Data will be encoded to JSON before saving to cache.
package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rl404/mal-plugin/cache"
)

// Redis client implements Cacher interface.
var _ cache.Cacher = &Client{}

// Client is redis client.
type Client struct {
	client      *redis.Client
	expiredTime time.Duration
	ctx         context.Context
}

// New to create cache cache with default config.
func New(address, password string, expiredTime time.Duration) (*Client, error) {
	return NewWithConfig(redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	}, expiredTime)
}

// NewWithConfig to create cache from go-redis options.
func NewWithConfig(option redis.Options, expiredTime time.Duration) (*Client, error) {
	client := redis.NewClient(&option)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping test.
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return NewFromGoRedis(client, expiredTime), nil
}

// NewFromGoRedis to create cache from go-redis client.
func NewFromGoRedis(client *redis.Client, expiredTime time.Duration) *Client {
	return &Client{
		client:      client,
		expiredTime: expiredTime,
		ctx:         context.Background(),
	}
}

// Set to save data to cache,
func (c *Client) Set(key string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, key, d, c.expiredTime).Err()
}

// Get to get data from cache.
func (c *Client) Get(key string, data interface{}) error {
	d, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(d), &data)
}

// Delete to delete data from cache.
func (c *Client) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

// Close to close cache connection.
func (c *Client) Close() error {
	return c.client.Close()
}
