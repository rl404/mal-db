package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rl404/mal-db/internal/cacher"
	"github.com/rl404/mal-db/internal/config"
	"github.com/rl404/mal-db/internal/database"
	"github.com/rl404/mal-db/internal/loader"
	"github.com/rl404/mal-db/internal/pkg/http"
	"github.com/rl404/mal-db/internal/pkg/middleware"
	"github.com/rl404/mal-db/internal/pubsub"
	"github.com/rl404/mal-db/internal/router/api"
	"github.com/rl404/mal-db/internal/router/ping"
	"github.com/rl404/mal-db/internal/router/swagger"
	"github.com/rl404/mal-plugin/log/elasticsearch"
	"github.com/rl404/mal-plugin/log/mallogger"
	"github.com/rs/cors"
)

func server() {
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

	// Init loader.
	service := loader.New(l, db, c, ps, time.Duration(cfg.Worker.AgeLimit)*time.Second)
	l.Info("loader initialized")

	// Init web server.
	server := http.New(http.Config{
		Port:            cfg.Web.Port,
		ReadTimeout:     cfg.Web.ReadTimeout,
		WriteTimeout:    cfg.Web.WriteTimeout,
		GracefulTimeout: cfg.Web.GracefulTimeout,
	})
	r := server.Router()

	// Init web router middleware.
	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.RealIP)
	if cfg.ES.Address != "" {
		// Init elasticsearch.
		es, err := elasticsearch.New(strings.Split(cfg.ES.Address, ","), cfg.ES.User, cfg.ES.Password)
		if err != nil {
			l.Fatal(err.Error())
		}
		r.Use(middleware.Logger(l, es))
		l.Info("elasticsearch initialized")
	}
	r.Use(middleware.Recoverer)
	l.Info("middleware initialized")

	// Register ping route.
	ping.New().Register(r)
	l.Info("base routes initialized")

	// Register swagger route.
	swagger.New().Register(r)
	l.Info("swagger routes initialized")

	// Register api routes.
	api.New(service).Register(r)
	l.Info("api routes initialized")

	// Run web server.
	serverChan := server.Run()
	l.Info("web server initialized")
	l.Info("server listening at %s", cfg.Web.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-serverChan:
		if err != nil {
			l.Fatal(err.Error())
		}
	case <-sigChan:
	}
}
