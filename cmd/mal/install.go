package main

import (
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-db/internal/config"
	"github.com/rl404/mal-db/internal/database"
	"github.com/rl404/mal-plugin/cache/nocache"
	"github.com/rl404/mal-plugin/log/mallogger"
)

func install() {
	// Get config.
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// Init logger.
	l := mallogger.New(cfg.Log.Level, cfg.Log.Color)
	l.Info("logger initialized")

	// Init malscraper.
	c, _ := nocache.New()
	mal, err := malscraper.New(malscraper.Config{
		Cacher:        c,
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

	// Migrate db.
	if err = database.Migrate(db, mal); err != nil {
		l.Fatal(err.Error())
	}
	l.Info("database migrated")
}
