package main

import (
	"github.com/rl404/mal-db/internal/config"
	"github.com/rl404/mal-db/internal/database"
	"github.com/rl404/mal-db/internal/pubsub"
	"github.com/rl404/mal-db/internal/tool"
	"github.com/rl404/mal-plugin/log/mallogger"
)

func filler() {
	// Get config.
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// Init logger.
	l := mallogger.New(cfg.Log.Level, cfg.Log.Color)
	l.Info("logger initialized")

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

	// Init tools.
	f := tool.NewFiller(l, db, ps)
	l.Info("tool initialized")

	// Run tools.
	if err = f.Run(); err != nil {
		l.Error(err.Error())
	}
	l.Info("done")
}
