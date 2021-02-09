package database

import (
	"errors"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/loader/api"
	"github.com/rl404/mal-db/internal/logger"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pubsub"
	"gorm.io/gorm"
)

// Database implements API interface.
type Database struct {
	log logger.Logger
	db  *gorm.DB
	ps  pubsub.PubSub
}

// New to create new database methods.
func New(l logger.Logger, db *gorm.DB, ps pubsub.PubSub) api.API {
	return &Database{
		log: l,
		db:  db,
		ps:  ps,
	}
}

func (d *Database) enqueue(t string, id int) error {
	return d.ps.Publish(constant.PubSubTopic, pubsub.Message{
		Type: t,
		ID:   id,
	})
}

func (d *Database) isEntryEmpty(t string, id int) bool {
	return !errors.Is(d.db.Where("type = ? and id = ?", t, id).First(&raw.EmptyID{}).Error, gorm.ErrRecordNotFound)
}
