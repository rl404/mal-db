package cacher

import (
	"time"

	"github.com/rl404/go-malscraper/service"
)

// Cacher intercepts request to check the requested
// data to cache before actually access and parse
// MyAnimeList web.
type Cacher struct {
	api    service.API
	cacher service.Cacher
	logger service.Logger
}

// New to create new cacher.
func New(api service.API, c service.Cacher, l service.Logger) service.API {
	return &Cacher{
		api:    api,
		cacher: newCacherLog(c, l),
		logger: l,
	}
}

// Simple cacher wrapper with log to prevent writing
// repetitive log code.
type cacherLog struct {
	cacher service.Cacher
	logger service.Logger
}

// Testable time since func.
var timeSince = time.Since

func newCacherLog(c service.Cacher, l service.Logger) service.Cacher {
	return &cacherLog{
		cacher: c,
		logger: l,
	}
}

// Get to get data from cache with log.
func (c cacherLog) Get(key string, data interface{}) error {
	c.logger.Trace("[%s] retrieving cache...", key)
	t := time.Now()
	if err := c.cacher.Get(key, data); err != nil {
		c.logger.Warn("[%s] failed retrieving cache: %s", key, err.Error())
		return err
	}
	c.logger.Debug("[%s] cache found (%s)", key, timeSince(t).Truncate(time.Microsecond))
	return nil
}

// Set to save data to cache with log.
func (c cacherLog) Set(key string, data interface{}) error {
	c.logger.Trace("[%s] saving cache...", key)
	t := time.Now()
	if err := c.cacher.Set(key, data); err != nil {
		c.logger.Error("[%s] failed saving cache: %s", key, err.Error())
		return err
	}
	c.logger.Debug("[%s] cache saved (%s)", key, timeSince(t).Truncate(time.Microsecond))
	return nil
}

// Delete to delete data in cache with log.
func (c cacherLog) Delete(key string) error {
	c.logger.Trace("[%s] deleting cache...", key)
	t := time.Now()
	if err := c.cacher.Delete(key); err != nil {
		c.logger.Error("[%s] failed deleting cache: %s", key, err.Error())
		return err
	}
	c.logger.Debug("[%s] cache deleted (%s)", key, timeSince(t).Truncate(time.Microsecond))
	return nil
}

// Close to close cache connection.
func (c cacherLog) Close() error {
	return c.cacher.Close()
}
