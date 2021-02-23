package main

import (
	"strings"
	"time"

	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-db/internal/cacher"
	"github.com/rl404/mal-db/internal/config"
	"github.com/rl404/mal-db/internal/database"
	"github.com/rl404/mal-db/internal/pubsub"
	"github.com/rl404/mal-db/internal/saver"
	"github.com/rl404/mal-db/internal/tool"
	"github.com/rl404/mal-plugin/cache/nocache"
	"github.com/rl404/mal-plugin/log/elasticsearch"
	"github.com/rl404/mal-plugin/log/mallogger"
)

func updater() {
	// Get config.
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// Init logger.
	l := mallogger.New(cfg.Log.Level, cfg.Log.Color)
	l.Info("logger initialized")

	// Init cache.
	c, err := cacher.New(l, cfg.Cache.Dialect, cfg.Cache.Address, cfg.Cache.Password, time.Duration(cfg.Cache.Time)*time.Second)
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("cache initialized")
	defer func() {
		if err = c.Close(); err != nil {
			l.Error(err.Error())
		} else {
			l.Info("cache stopped")
		}
	}()

	// Init malscraper without caching.
	nc, _ := nocache.New()
	mal, err := malscraper.New(malscraper.Config{
		Cacher:        nc,
		CleanImageURL: true,
		CleanVideoURL: true,
		Logger:        l,
	})
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("malscraper initialized")

	// Init db.
	db, err := database.New(cfg.DB.Address, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, cfg.DB.MaxConnOpen, cfg.DB.MaxConnIdle, cfg.DB.MaxConnLifetime)
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("database initialized")
	defer func() {
		if tmp, err := db.DB(); err != nil {
			l.Error(err.Error())
		} else {
			if err = tmp.Close(); err != nil {
				l.Error(err.Error())
			} else {
				l.Info("database stopped")
			}
		}
	}()

	var es *elasticsearch.Client
	if cfg.ES.Address != "" {
		// Init elasticsearch.
		es, err = elasticsearch.New(strings.Split(cfg.ES.Address, ","), cfg.ES.User, cfg.ES.Password)
		if err != nil {
			l.Fatal(err.Error())
		}
		l.Info("elasticsearch initialized")
	}

	// Init pubsub.
	ps, err := pubsub.New(l, cfg.PubSub.Dialect, cfg.PubSub.Address, cfg.PubSub.User, cfg.PubSub.Password)
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("pubsub initialized")
	defer func() {
		if err = ps.Close(); err != nil {
			l.Error(err.Error())
		} else {
			l.Info("pubsub stopped")
		}
	}()

	// Init saver.
	svr := saver.New(l, c, db, mal, ps, es)
	l.Info("saver initialized")

	// Init tools.
	t := tool.NewUpdater(l, svr, db, ps)
	l.Info("tool initialized")

	// Run tools.
	for i := 0; i < 5; i++ {
		if err = t.Run(); err != nil {
			l.Error(err.Error())
		}
		l.Trace("break time...")
		time.Sleep(time.Duration(cfg.Worker.BreakTime) * time.Second)
	}
	l.Info("done")
}
